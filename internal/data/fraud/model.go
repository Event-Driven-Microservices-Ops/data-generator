package fraud

import "time"

type Fraud struct {
	EventID      string      `json:"event_id"`
	UserID       string      `json:"user_id"`
	FraudScore   int         `json:"fraud_score"` // 0-100
	IPAddress    string      `json:"ip_address"`
	CountryCode  CountryCode `json:"country_code"`
	DeviceType   DeviceType  `json:"device_type"`
	IsSuspicious bool        `json:"is_suspicious"`
	Timestamp    time.Time   `json:"timestamp"`
}
