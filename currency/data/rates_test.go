package data

import (
	"testing"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
)

func TestRatesRetrievedOnNew(t *testing.T) {
	t.Skip()
	er, err := NewRates(hclog.Default())
	assert.NoError(t, err)
	assert.NotNil(t, er)
}

func TestReturnsCorrectRate(t *testing.T) {
	rates := map[string]float64{"GBP": 0.89, "USD": 1.097}
	er := ExchangeRates{log: hclog.NewNullLogger(), rates: rates}

	r, err := er.GetRate("GBP", "USD")
	assert.NoError(t, err)
	assert.Equal(t, 1.2325842696629212, r)
}

func TestReturnsCorrectRateWhenBaseEuro(t *testing.T) {
	rates := map[string]float64{"GBP": 0.89, "USD": 1.097, "EUR": 1.0}
	er := ExchangeRates{log: hclog.NewNullLogger(), rates: rates}

	r, err := er.GetRate("EUR", "USD")
	assert.NoError(t, err)
	assert.Equal(t, 1.097, r)
}

func TestReturnsErrorWhenBaseCurrencyNotExist(t *testing.T) {
	rates := map[string]float64{"GBP": 0.89, "USD": 1.097, "EUR": 1.0}
	er := ExchangeRates{log: hclog.NewNullLogger(), rates: rates}

	r, err := er.GetRate("HKD", "USD")
	assert.Error(t, err)
	assert.Equal(t, 0.0, r)
}

func TestReturnsErrorWhenDestCurrencyNotExist(t *testing.T) {
	rates := map[string]float64{"GBP": 0.89, "USD": 1.097, "EUR": 1.0}
	er := ExchangeRates{log: hclog.NewNullLogger(), rates: rates}

	r, err := er.GetRate("GBP", "HKD")
	assert.Error(t, err)
	assert.Equal(t, 0.0, r)
}

func TestMonitorUpdatesRates(t *testing.T) {
	rates := map[string]float64{"GBP": 0.89, "USD": 1.097, "EUR": 1.0}
	er := ExchangeRates{log: hclog.NewNullLogger(), rates: rates}

	r, err := er.GetRate("GBP", "USD")
	assert.NoError(t, err)

	// monitor for changes
	rc := er.MonitorRates(1 * time.Millisecond)

	// check multiple times to ensure that multiple updates
	// occur
	for i := 0; i < 2; i++ {
		// block for updates
		<-rc

		// get the rate again
		r2, err := er.GetRate("GBP", "USD")
		assert.NoError(t, err)

		// check is different
		assert.NotEqual(t, r, r2)

		r = r2
	}
}
