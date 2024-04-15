package bank

import (
	"time"

	"github.com/google/uuid"
)

const TransactionCreatedEvent = "transaction-created"
const TransactionSettledEvent = "transaction-settled"
const TransactionFailedEvent = "transaction-failed"

type TransactionCreated struct {
	ID                        uuid.UUID
	Provider                  string
	Currency                  string
	Amount                    float64
	AccountIdentifier         string
	AccountOwner              string
	OppositeAccountIdentifier string
	OppositeAccountOwner      string
	Reference                 string
	ExternalIdentifier        string
	Timestamp                 time.Time
}

type TransactionSettled struct {
	ID                        uuid.UUID
	Provider                  string
	Currency                  string
	Amount                    float64
	AccountIdentifier         string
	AccountOwner              string
	OppositeAccountIdentifier string
	OppositeAccountOwner      string
	Reference                 string
	ExternalIdentifier        string
	Timestamp                 time.Time
}

type TransactionFailed struct {
	ID                 uuid.UUID
	Provider           string
	Reason             string
	ExternalIdentifier string
	Timestamp          time.Time
}
