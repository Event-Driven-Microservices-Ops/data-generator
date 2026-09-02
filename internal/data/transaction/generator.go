package transaction

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func GenerateNewTransactionData(forcedAmount *float64, forcedCurrency *string, forcedStatus *string, forcedPaymentType *string) Transaction {
	var amountToUse float64
	if forcedAmount != nil {
		amountToUse = *forcedAmount
	} else {
		amountToUse = (rand.Float64() * 4990.0) + 10.0 //  10.00 - 5000.00
	}

	var currencyToUse Currency
	if forcedCurrency != nil && *forcedCurrency != "" {
		currencyToUse = Currency(*forcedCurrency)
	} else {
		currencyToUse = RandomCurrency()
	}

	var statusToUse Status
	if forcedStatus != nil && *forcedStatus != "" {
		statusToUse = Status(*forcedStatus)
	} else {
		statusToUse = RandomStatus()
	}

	var paymentTypeToUse PaymentType
	if forcedPaymentType != nil && *forcedPaymentType != "" {
		paymentTypeToUse = PaymentType(*forcedPaymentType)
	} else {
		paymentTypeToUse = RandomPaymentType()
	}

	return Transaction{
		TransactionID: uuid.New().String(),
		UserID:        uuid.New().String(),
		Amount:        amountToUse,
		Currency:      currencyToUse,
		Status:        statusToUse,
		PaymentType:   paymentTypeToUse,
		Timestamp:     time.Now().UTC(),
	}
}
