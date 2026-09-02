package api

import (
	"time"

	"github.com/urbaniakmichal/data-generator/internal/data/account"
	"github.com/urbaniakmichal/data-generator/internal/data/fraud"
	"github.com/urbaniakmichal/data-generator/internal/data/transaction"
)

type Payload struct {
	MetaData MetaData `json:"metadata"`
	Data     []Data   `json:"data"`
}

type MetaData struct {
	UUID      string    `json:"uuid"`
	Timestamp time.Time `json:"timestamp"`
}

type Data struct {
	Account     account.Account         `json:"account"`
	Fraud       fraud.Fraud             `json:"fraud"`
	Transaction transaction.Transaction `json:"transaction"`
}
