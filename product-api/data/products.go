package data

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-hclog"
	protos "github.com/nicholasjackson/building-microservices-youtube/currency/protos/currency"
)

var baseCurrency = "EUR"

// ErrProductNotFound is an error raised when a product can not be found in the database
var ErrProductNotFound = fmt.Errorf("Product not found")

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the product

	// the name for this poduct
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this poduct
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float64 `json:"price" validate:"required,gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"sku"`
}

// Products defines a slice of Product
type Products []*Product

// ProductsDB is a structure for retriving products from the datastore
type ProductsDB struct {
	// client for the currency service
	currency protos.CurrencyClient
	// ratesClient is the bidiectional streaming client for the currency service
	ratesClient protos.Currency_SubscribeRatesClient
	// exchange rate cache
	rates map[string]float64
	// logger
	log hclog.Logger
}

// NewProductsDB returns a new ProductDB
func NewProductsDB(c protos.CurrencyClient, l hclog.Logger) *ProductsDB {

	p := &ProductsDB{currency: c, log: l, rates: map[string]float64{}}

	// subscribe for rate changes
	go p.subscribeRates()

	return p
}

func (p *ProductsDB) subscribeRates() {
	p.log.Info("Listening for rate updates")

	// create an instance of the bidirectional streaming API to
	// subscribe for updates
	sr, err := p.currency.SubscribeRates(context.Background())
	if err != nil {
		p.log.Error("Unable to subscribe for rate updates", "error", err)
		return
	}

	p.ratesClient = sr

	// listen for updated rates on the API
	for {
		rr, err := p.ratesClient.Recv()
		if err != nil {
			p.log.Error("Unable to receive streaming message from currency server", "error", err)
			return
		}

		p.log.Info("Received update rate from currency server", "currency", rr.GetDestination().String(), "rate", rr.GetRate())

		// update the cached value
		p.rates[rr.GetDestination().String()] = rr.GetRate()
	}
}

func (p *ProductsDB) getRate(currency string) (float64, error) {
	// check the cache to see if the rate is in there
	// if so return it
	if r, ok := p.rates[currency]; ok {
		p.log.Info("Found rate in the cache", "currency", currency)
		return r, nil
	}

	p.log.Info("Fetching rate from currency server", "currency", currency)
	// get the exchange rate from the currency service
	rr, err := p.currency.GetRate(
		context.Background(),
		&protos.RateRequest{
			Base:        protos.Currencies(protos.Currencies_value[baseCurrency]),
			Destination: protos.Currencies(protos.Currencies_value[currency]),
		},
	)

	if err != nil {
		p.log.Error("Error fetching rate from currency server", "error", err)
		return -1, err
	}

	// cache the result
	p.rates[currency] = rr.GetRate()

	// subscribe for updates on this rate
	p.log.Info("Subscribing for changes in the rate", "currency", currency)
	p.ratesClient.Send(
		&protos.RateRequest{
			Base:        protos.Currencies(protos.Currencies_value[baseCurrency]),
			Destination: protos.Currencies(protos.Currencies_value[currency]),
		},
	)

	return rr.GetRate(), nil
}

// GetProducts returns all products from the database and convert
// the price for the given currency
func (p *ProductsDB) GetProducts(currency string) Products {
	if currency == "" {
		return productList
	}

	// modify the price for each product
	rpl := []*Product{}
	for _, pr := range productList {
		// copy the product
		np := *pr

		// convert the price to the destination currency
		rate, err := p.getRate(currency)
		if err != nil {
			return nil
		}

		np.Price = np.Price * rate

		rpl = append(rpl, &np)
	}

	return rpl
}

// GetProductByID returns a single product which matches the id from the
// database. Product price will be returned converted for the given currency.
// If a product is not found this function returns a ProductNotFound error
func (p *ProductsDB) GetProductByID(id int, currency string) (*Product, error) {
	i := findIndexByProductID(id)
	if id == -1 {
		return nil, ErrProductNotFound
	}

	np := *productList[i]

	if currency == "" {
		return &np, nil
	}

	// convert the price to the destination currency
	rate, err := p.getRate(currency)
	if err != nil {
		return nil, err
	}

	np.Price = np.Price * rate

	return &np, nil
}

// UpdateProduct replaces a product in the database with the given
// item.
// If a product with the given id does not exist in the database
// this function returns a ProductNotFound error
func (p *ProductsDB) UpdateProduct(prod Product) error {
	i := findIndexByProductID(prod.ID)
	if i == -1 {
		return ErrProductNotFound
	}

	// update the product in the DB
	productList[i] = &prod

	return nil
}

// AddProduct adds a new product to the database
func (p *ProductsDB) AddProduct(prod Product) {
	// get the next id in sequence
	maxID := productList[len(productList)-1].ID
	prod.ID = maxID + 1
	productList = append(productList, &prod)
}

// DeleteProduct deletes a product from the database
func (p *ProductsDB) DeleteProduct(id int) error {
	i := findIndexByProductID(id)
	if i == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:i], productList[i+1])

	return nil
}

// findIndex finds the index of a product in the database
// returns -1 when no product can be found
func findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}

	return -1
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
	},
	&Product{
		ID:          3,
		Name:        "Frapuccinio",
		Description: "Blended Ice Coffee",
		Price:       3.99,
		SKU:         "34v9d",
	},
	&Product{
		ID:          4,
		Name:        "Tea",
		Description: "Classic black tea",
		Price:       1.50,
		SKU:         "3rt33",
	},
}
