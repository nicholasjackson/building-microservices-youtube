package data

import (
	"testing"

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
