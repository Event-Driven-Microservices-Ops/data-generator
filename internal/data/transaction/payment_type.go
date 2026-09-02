package transaction

import "math/rand"

type PaymentType string

const (
	CARD          PaymentType = "CARD"
	CASH          PaymentType = "CASH"
	BANK_TRANSFER PaymentType = "BANK_TRANSFER"
	PAYPAL        PaymentType = "PAYPAL"
	APPLE_PAY     PaymentType = "APPLE_PAY"
	GOOGLE_PAY    PaymentType = "GOOGLE_PAY"
	BLIK          PaymentType = "BLIK"
)

var allPaymentTypes = []PaymentType{CARD, CASH, BANK_TRANSFER, PAYPAL, APPLE_PAY, GOOGLE_PAY, BLIK}

func RandomPaymentType() PaymentType {
	return allPaymentTypes[rand.Intn(len(allPaymentTypes))]
}
