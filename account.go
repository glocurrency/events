package events

import (
	"time"

	"github.com/google/uuid"
)

type AccountOpen struct {
	ID         uuid.UUID
	Source     string
	Name       string
	Owner      string
	Identifier string
	Date       time.Time
}

type AccountUpdated struct {
	ID         uuid.UUID
	Source     string
	Name       string
	Owner      string
	Identifier string
	Date       time.Time
}

type AccountClosed struct {
	ID     uuid.UUID
	Source string
	Reason string
	Date   time.Time
}
