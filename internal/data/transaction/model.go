package transaction

import (
	"time"
)

type Transaction struct {
	TransactionID string      `json:"transaction_id"`
	UserID        string      `json:"user_id"`
	Amount        float64     `json:"amount"`
	Currency      Currency    `json:"currency"`
	Status        Status      `json:"status"`
	PaymentType   PaymentType `json:"payment_type"`
	Timestamp     time.Time   `json:"timestamp"`
}
