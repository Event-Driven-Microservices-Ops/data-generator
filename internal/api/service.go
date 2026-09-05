package api

import (
	"time"

	"github.com/google/uuid"
	"github.com/urbaniakmichal/data-generator/internal/config"
	"github.com/urbaniakmichal/data-generator/internal/data/account"
	"github.com/urbaniakmichal/data-generator/internal/data/fraud"
	"github.com/urbaniakmichal/data-generator/internal/data/transaction"
)

type Service struct {
	// for future use e.g. logger
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GenerateNewData(f config.Flags) []Payload {
	var payloads []Payload

	count := 1
	if f.BatchSizeFlag != nil && *f.BatchSizeFlag > 0 {
		count = *f.BatchSizeFlag
	}

	for i := 0; i < count; i++ {
		acc := account.GenerateNewEAccountData(f.EventTypeFlag)
		fr := fraud.GenerateNewFraudData(f.FraudScoreFlag, f.CountryCodeFlag, f.DeviceTypeFlag)
		tr := transaction.GenerateNewTransactionData(f.AmountFlag, f.CountryCodeFlag, f.StatusFlag, f.PaymentTypeFlag)

		payload := Payload{
			MetaData: MetaData{
				UUID:      uuid.New().String(),
				Timestamp: time.Now().UTC(),
			},
			Data: []Data{
				{
					Account:     acc,
					Fraud:       fr,
					Transaction: tr,
				},
			},
		}

		payloads = append(payloads, payload)
	}

	return payloads
}
