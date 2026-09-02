package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/urbaniakmichal/data-generator/internal/config"
)

type RestHandler struct {
	service *Service
}

func NewRestHandler(s *Service) *RestHandler {
	return &RestHandler{
		service: s,
	}
}

func (rh *RestHandler) GetData(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()

	flags := &config.Flags{
		CountFlag:       parseIntPtr(query.Get("count")),
		IntervalFlag:    parseIntPtr(query.Get("interval")),
		AccountFlag:     parseBoolPtr(query.Get("account")),
		FraudFlag:       parseBoolPtr(query.Get("fraud")),
		TransactionFlag: parseBoolPtr(query.Get("transaction")),
		AmountFlag:      parseFloatPtr(query.Get("amount")),
		CurrencyFlag:    parseStringPtr(query.Get("currency")),
		StatusFlag:      parseStringPtr(query.Get("status")),
		PaymentTypeFlag: parseStringPtr(query.Get("payment_type")),
		EventTypeFlag:   parseStringPtr(query.Get("event_type")),
		FraudScoreFlag:  parseIntPtr(query.Get("fraud_score")),
		CountryCodeFlag: parseStringPtr(query.Get("country")),
		DeviceTypeFlag:  parseStringPtr(query.Get("device")),
	}

	resp := rh.service.GenerateNewData(*flags)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(res).Encode(resp)
}

func parseIntPtr(s string) *int {
	if s == "" {
		return nil
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	return &v
}

func parseFloatPtr(s string) *float64 {
	if s == "" {
		return nil
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil
	}
	return &v
}

func parseStringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func parseBoolPtr(s string) *bool {
	if s == "" {
		return nil
	}
	v, err := strconv.ParseBool(s)
	if err != nil {
		return nil
	}
	return &v
}
