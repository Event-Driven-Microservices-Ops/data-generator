package fraud

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func GenerateNewFraudData(forcedFraudScore *int, forcedCountryCode *string, forcedDeviceType *string) Fraud {
	var scoreToUse int
	if forcedFraudScore != nil {
		scoreToUse = *forcedFraudScore
	} else {
		scoreToUse = rand.Intn(101)
	}

	var countryCodeToUse CountryCode
	if forcedCountryCode != nil && *forcedCountryCode != "" {
		countryCodeToUse = CountryCode(*forcedCountryCode)
	} else {
		countryCodeToUse = RandomCountry()
	}

	var deviceTypeToUse DeviceType
	if forcedDeviceType != nil && *forcedDeviceType != "" {
		deviceTypeToUse = DeviceType(*forcedDeviceType)
	} else {
		deviceTypeToUse = RandomDevice()
	}

	return Fraud{
		EventID:      uuid.New().String(),
		UserID:       uuid.New().String(),
		FraudScore:   scoreToUse,
		IPAddress:    fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255)),
		CountryCode:  countryCodeToUse,
		DeviceType:   deviceTypeToUse,
		IsSuspicious: scoreToUse > 70,
		Timestamp:    time.Now().UTC(),
	}
}
