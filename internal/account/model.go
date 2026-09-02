package account

import "time"

type Account struct {
	AccountID string    `json:"account_id"`
	UserID    string    `json:"user_id"`
	EventType EventType `json:"event_type"`
	SessionID string    `json:"session_id"`
	Timestamp time.Time `json:"timestamp"`
}
