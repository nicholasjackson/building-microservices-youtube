package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/nicholasjackson/building-microservices-youtube/currency/data"
	protos "github.com/nicholasjackson/building-microservices-youtube/currency/protos/currency"
	"github.com/prometheus/common/log"
)

// Currency is a gRPC server it implements the methods defined by the CurrencyServer interface
type Currency struct {
	log   hclog.Logger
	rates *data.ExchangeRates
	// map holding the client connection and the list of rates they are subscribed to
	subscriptions map[protos.Currency_StreamRatesServer][]*protos.RateRequest
}

// NewCurrency creates a new Currency server
func NewCurrency(er *data.ExchangeRates, l hclog.Logger) *Currency {
	return &Currency{rates: er, log: l, subscriptions: make(map[protos.Currency_StreamRatesServer][]*protos.RateRequest)}
}

// GetRate implements the CurrencyServer GetRate method and returns the currency exchange rate
// for the two given currencies.
func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())

	// get the rate
	r, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}

	return &protos.RateResponse{Rate: r}, nil
}

// SubscribeRates allows a client to subscribe for changes in exchange rates
// this implements gRPC bidirectional streaming
//
// The client sends a protos.RateRequest to the stream and this is logged on the server
// whenever the rates change for that currency the rate is sent as a response
func (c *Currency) SubscribeRates(srs protos.Currency_StreamRatesServer) error {
	// setup the collection to store subscribed rates
	c.subscriptions[srs] = []*protos.RateRequest{}

	// connection has closed cleanup
	defer delete(c.subscriptions, srs)

	// listen for the stream, this is a blocking call
	for {
		msg, err := srs.Recv()
		if err != nil {
			log.Error("Unable to read message from client", "erroro", err)

			return err
		}

		// add the rate to the subscription
		log.Info("Received rate subscription from client", "message", msg)
		c.subscriptions[srs] = append(c.subscriptions[srs], msg)
	}
}
