package config

type Flags struct {
	BatchSizeFlag *int
	IntervalFlag  *int

	AccountFlag   *bool
	EventTypeFlag *string

	FraudFlag       *bool
	FraudScoreFlag  *int
	CountryCodeFlag *string
	DeviceTypeFlag  *string

	TransactionFlag *bool
	AmountFlag      *float64
	CurrencyFlag    *string
	StatusFlag      *string
	PaymentTypeFlag *string
}
