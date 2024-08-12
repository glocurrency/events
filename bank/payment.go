package bank

import (
	"time"

	"github.com/google/uuid"
)

const PaymentComplianceRequiredEvent = "payment-compliance-required"
const PaymentProcessingEvent = "payment-processing"
const PaymentCompletedEvent = "payment-completed"
const PaymentFailedEvent = "payment-failed"

type PaymentCreated struct {
	PaymentID uuid.NullUUID
	Currency  string
	Amount    float64
	Reference string
	Timestamp time.Time
}

type PaymentComplianceRequired struct {
	PaymentID uuid.NullUUID
	Currency  string
	Amount    float64
	Reference string
	Timestamp time.Time
}

type PaymentProcessing struct {
	PaymentID uuid.NullUUID
	Reason    string
	Timestamp time.Time
}

type PaymentCompleted struct {
	PaymentID uuid.NullUUID
	Reason    string
	Timestamp time.Time
}

type PaymentFailed struct {
	PaymentID uuid.NullUUID
	Reason    string
	Timestamp time.Time
}
