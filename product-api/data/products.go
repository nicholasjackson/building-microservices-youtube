package data

import (
	"context"
	"fmt"

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
}

// NewProductsDB returns a new ProductDB
func NewProductsDB(c protos.CurrencyClient) *ProductsDB {
	return &ProductsDB{c}
}

// GetProducts returns all products from the database and convert
// the price for the given currency
func (p *ProductsDB) GetProducts(currency string) Products {
	if currency == "" {
		return productList
	}

	fmt.Println(baseCurrency, currency)
	// get the exchange rate from the currency service
	rr, err := p.currency.GetRate(
		context.Background(),
		&protos.RateRequest{
			Base:        protos.Currencies(protos.Currencies_value[baseCurrency]),
			Destination: protos.Currencies(protos.Currencies_value[currency]),
		},
	)

	if err != nil {
		return nil
	}

	// modify the price for each product
	rpl := []*Product{}
	for _, p := range productList {
		// copy the product
		np := *p

		// convert the price to the destination currency
		np.Price = np.Price * rr.GetRate()

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

	return productList[i], nil
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
		Name:        "Esspresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
	},
}
