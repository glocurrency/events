package bank

import (
	"time"

	"github.com/google/uuid"
)

const FxCreatedEvent = "fx-created"
const FxSettledEvent = "fx-settled"
const FxFailedEvent = "fx-failed"

type FxCreated struct {
	ID       uuid.UUID
	Provider string

	BuyCurrency    string
	BuyAmount      float64
	SellCurrency   string
	SellAmount     float64
	MarginCurrency string
	MarginAmount   float64

	Rate float64

	SellAccountIdentifier   string
	BuyAccountIdentifier    string
	MarginAccountIdentifier string

	ExternalIdentifier string
	Timestamp          time.Time
}

type FxSettled struct {
	ID       uuid.UUID
	Provider string

	BuyCurrency    string
	BuyAmount      float64
	SellCurrency   string
	SellAmount     float64
	MarginCurrency string
	MarginAmount   float64

	Rate float64

	SellAccountIdentifier   string
	BuyAccountIdentifier    string
	MarginAccountIdentifier string

	ExternalIdentifier string
	Timestamp          time.Time
}

type FxFailed struct {
	ID                 uuid.UUID
	Provider           string
	Reason             string
	ExternalIdentifier string
	Timestamp          time.Time
}
