package bank

import (
	"time"

	"github.com/google/uuid"
)

const AccountOpenEvent = "account-open"
const AccountUpdatedEvent = "account-updated"
const AccountClosedEvent = "account-closed"

type AccountOpen struct {
	ID         uuid.UUID
	Provider   string
	Name       string
	Owner      string
	Identifier string
	Currencies []string
	Timestamp  time.Time
}

type AccountUpdated struct {
	ID         uuid.UUID
	Provider   string
	Name       string
	Owner      string
	Identifier string
	Currencies []string
	Timestamp  time.Time
}

type AccountClosed struct {
	ID        uuid.UUID
	Provider  string
	Reason    string
	Timestamp time.Time
}
