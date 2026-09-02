package account

import (
	"time"

	"github.com/google/uuid"
)

func GenerateNewEAccountData(forcedEventType *string) Account {
	var eventToUse EventType

	if forcedEventType != nil && *forcedEventType != "" {
		eventToUse = EventType(*forcedEventType)
	} else {
		eventToUse = RandomEvent()
	}

	return Account{
		AccountID: uuid.New().String(),
		UserID:    uuid.New().String(),
		EventType: eventToUse,
		SessionID: uuid.New().String(),
		Timestamp: time.Now().UTC(),
	}
}
