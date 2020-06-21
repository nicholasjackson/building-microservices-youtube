package data

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-hclog"
	protos "github.com/nicholasjackson/building-microservices-youtube/currency/protos/currency"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

type ProductsDB struct {
	currency protos.CurrencyClient
	log      hclog.Logger
	rates    map[string]float64
	client   protos.Currency_SubscribeRatesClient
}

// NewProductsDB returns a Data object for CRUD operations on
// Products data.
// This type also handles conversion of currencies through integraiton with the
// currency service.
func NewProductsDB(c protos.CurrencyClient, l hclog.Logger) *ProductsDB {
	pb := &ProductsDB{c, l, make(map[string]float64), nil}

	go pb.handleUpdates()

	return pb
}

func (p *ProductsDB) handleUpdates() {
	sub, err := p.currency.SubscribeRates(context.Background())
	if err != nil {
		p.log.Error("Unable to subscribe for rates", "error", err)
		return
	}

	p.client = sub

	for {
		// Recv returns a StreamingRateResponse which can contain one of two messages
		// RateResponse or an Error.
		// We need to handle each case separately
		srr, err := sub.Recv()

		// handle connection errors
		// this is normally terminal requires a reconnect
		if err != nil {
			p.log.Error("Error while waiting for message", "error", err)
			return
		}

		// handle a returned error message
		if ge := srr.GetError(); ge != nil {
			sre := status.FromProto(ge)

			if sre.Code() == codes.InvalidArgument {
				errDetails := ""
				// get the RateRequest serialized in the error response
				// Details is a collection but we are only returning a single item
				if d := sre.Details(); len(d) > 0 {
					p.log.Error("Deets", "d", d)
					if rr, ok := d[0].(*protos.RateRequest); ok {
						errDetails = fmt.Sprintf("base: %s destination: %s", rr.GetBase().String(), rr.GetDestination().String())
					}
				}

				p.log.Error("Received error from currency service rate subscription", "error", ge.GetMessage(), "details", errDetails)
			}
		}

		// handle a rate response
		if rr := srr.GetRateResponse(); rr != nil {
			p.log.Info("Recieved updated rate from server", "dest", rr.GetDestination().String())
			p.rates[rr.Destination.String()] = rr.Rate
		}
	}
}

// GetProducts returns all products from the database
func (p *ProductsDB) GetProducts(currency string) (Products, error) {
	if currency == "" {
		return productList, nil
	}

	rate, err := p.getRate(currency)
	if err != nil {
		p.log.Error("Unable to get rate", "currency", currency, "error", err)
		return nil, err
	}

	pr := Products{}
	for _, p := range productList {
		np := *p
		np.Price = np.Price * rate
		pr = append(pr, &np)
	}

	return pr, nil
}

// GetProductByID returns a single product which matches the id from the
// database.
// If a product is not found this function returns a ProductNotFound error
func (p *ProductsDB) GetProductByID(id int, currency string) (*Product, error) {
	i := findIndexByProductID(id)
	if id == -1 {
		return nil, ErrProductNotFound
	}

	if currency == "" {
		return productList[i], nil
	}

	rate, err := p.getRate(currency)
	if err != nil {
		p.log.Error("Unable to get rate", "currency", currency, "error", err)
		return nil, err
	}

	np := *productList[i]
	np.Price = np.Price * rate

	return &np, nil
}

// UpdateProduct replaces a product in the database with the given
// item.
// If a product with the given id does not exist in the database
// this function returns a ProductNotFound error
func (p *ProductsDB) UpdateProduct(pr Product) error {
	i := findIndexByProductID(pr.ID)
	if i == -1 {
		return ErrProductNotFound
	}

	// update the product in the DB
	productList[i] = &pr

	return nil
}

// AddProduct adds a new product to the database
func (p *ProductsDB) AddProduct(pr Product) {
	// get the next id in sequence
	maxID := productList[len(productList)-1].ID
	pr.ID = maxID + 1
	productList = append(productList, &pr)
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

func (p *ProductsDB) getRate(destination string) (float64, error) {
	// if cached return
	/*
		if r, ok := p.rates[destination]; ok {
			return r, nil
		}
	*/

	rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["EUR"]),
		Destination: protos.Currencies(protos.Currencies_value[destination]),
	}

	// get initial rate
	resp, err := p.currency.GetRate(context.Background(), rr)
	if err != nil {
		// convert the GRPC error message
		grpcError, ok := status.FromError(err)
		if !ok {
			// unable to convert grpc error
			return -1, err
		}

		// if this is an Invalid Arguments exception santise the message before returning
		if grpcError.Code() == codes.InvalidArgument {
			return -1, fmt.Errorf("Unable to retreive exchange rate from currency service: %s", grpcError.Message())
		}
	}

	// update cachce
	p.rates[destination] = resp.Rate

	// subscribe for updates
	err = p.client.Send(rr)
	if err != nil {
		return -1, err
	}

	return resp.Rate, err
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
