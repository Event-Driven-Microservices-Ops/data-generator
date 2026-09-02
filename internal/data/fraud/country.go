package fraud

import "math/rand"

type CountryCode string

const (
	PL  CountryCode = "POLAND"
	USA CountryCode = "USA"
	EU  CountryCode = "EUROPE"
	IR  CountryCode = "IRAN"
	RU  CountryCode = "RUSSIA"
)

var allCountries = []CountryCode{PL, USA, EU, IR, RU}

func RandomCountry() CountryCode {
	return allCountries[rand.Intn(len(allCountries))]
}
