package transaction

import "math/rand"

type Currency string

const (
	PLN Currency = "PLN"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

var allCurrencies = []Currency{PLN, USD, EUR}

func RandomCurrency() Currency {
	return allCurrencies[rand.Intn(len(allCurrencies))]
}
