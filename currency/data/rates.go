package data

import (
	"encoding/xml"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/go-hclog"
)

// ExchangeRates allows access to the Reference echange rates provided by the European central
// bank
type ExchangeRates struct {
	log   hclog.Logger
	rates map[string]float64
}

// NewRates creates a new ExchangeRates
func NewRates(l hclog.Logger) (*ExchangeRates, error) {
	er := &ExchangeRates{log: l, rates: map[string]float64{}}

	// get the rates from the European Central Bank API
	err := er.getRates()
	if err != nil {
		return nil, err
	}

	return er, nil
}

// GetRate returns the exchange rate between the base currency and the destination.
// If either base or destination currency is not in the database, an error is returned.
func (e *ExchangeRates) GetRate(base string, dest string) (float64, error) {
	// find the rate for both currencies
	br, ok := e.rates[base]
	if !ok {
		return 0, fmt.Errorf("Rate not found for currency %s", base)
	}

	dr, ok := e.rates[dest]
	if !ok {
		return 0, fmt.Errorf("Rate not found for currency %s", dest)
	}

	// the base rate is Euro to obtain the rate with a different base we need to divide the desination and the base
	return dr / br, nil
}

// MonitorRates checks the rates in the ECB API every interval and sends a message to the
// returned channel when there are changes
//
// Note: the ECB API only returns data once a day, this function only simulates the changes
// in rates for demonstration purposes
func (e *ExchangeRates) MonitorRates(interval time.Duration) chan map[string]float64 {
	ret := make(chan map[string]float64)

	go func() {
		ticker := time.NewTicker(interval)
		for {
			select {
			case <-ticker.C:
				// just add a random difference to the rate and return it
				// this simulates the fluctuations in currency rates
				for k, v := range e.rates {
					// change can be 10% of original value
					change := (rand.Float64() / 10)
					// is this a postive or negative change
					direction := rand.Intn(1)

					if direction == 0 {
						// new value with be min 90% of old
						change = 1 - change
					} else {
						// new value will be 110% of old
						change = 1 + change
					}

					// modify the rate
					e.rates[k] = v * change
				}

				// notify updates, this will block unless there is a listener on the other end
				ret <- e.rates
			}
		}
	}()

	return ret
}

// getRates fetches the reference rates from the Europen central banks
// API.
func (e *ExchangeRates) getRates() error {
	resp, err := http.DefaultClient.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Expected status code 200 got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	// parse the data
	md := &Cubes{}
	xml.NewDecoder(resp.Body).Decode(&md)

	// store rates in the cache
	for _, c := range md.CubeData {
		r, err := strconv.ParseFloat(c.Rate, 64)
		if err != nil {
			return err
		}

		e.rates[c.Currency] = r
	}

	// add the base currency EUR
	e.rates["EUR"] = 1

	e.log.Info("Got data", "rates", e.rates)
	return nil
}

// Cubes is the holding data type corresponding to the data returned by the ECB
//	<gesmes:Envelope xmlns:gesmes="http://www.gesmes.org/xml/2002-08-01" xmlns="http://www.ecb.int/vocabulary/2002-08-01/eurofxref">
//		<gesmes:subject>Reference rates</gesmes:subject>
//		<gesmes:Sender>...</gesmes:Sender>
//		<Cube>
//		<Cube time="2020-03-27">
//			<Cube currency="USD" rate="1.0977"/>
//			</Cube>
//		</Cube>
//	</gesmes>
type Cubes struct {
	CubeData []Cube `xml:"Cube>Cube>Cube"`
}

// Cube represents an individual line item for a currency in the returned
// data from the ECB API
type Cube struct {
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}
