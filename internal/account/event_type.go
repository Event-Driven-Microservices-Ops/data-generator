package account

import "math/rand"

type EventType string

const (
	ACCOUNT_CREATED  EventType = "ACCOUNT_CREATED"
	KYC_VERIFIED     EventType = "KYC_VERIFIED"
	PASSWORD_CHANGED EventType = "PASSWORD_CHANGED"
	BALANCE_UPDATED  EventType = "BALANCE_UPDATED"
)

var allEvents = []EventType{ACCOUNT_CREATED, KYC_VERIFIED, PASSWORD_CHANGED, BALANCE_UPDATED}

func RandomEvent() EventType {
	return allEvents[rand.Intn(len(allEvents))]
}
