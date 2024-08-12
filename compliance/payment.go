package compliance

import (
	"time"

	"github.com/google/uuid"
)

const PaymentAcceptedEvent = "compliance-payment-accepted"
const PaymentPendingEvent = "compliance-payment-pending"
const PaymentApprovedEvent = "compliance-payment-approved"
const PaymentDeclinedEvent = "compliance-payment-declined"

type PaymentPending struct {
	PaymentID          uuid.NullUUID
	ComplianceProvider string
	Reason             string
	Timestamp          time.Time
}

type PaymentApproved struct {
	PaymentID          uuid.NullUUID
	ComplianceProvider string
	Reason             string
	Timestamp          time.Time
}

type PaymentDeclined struct {
	PaymentID          uuid.NullUUID
	ComplianceProvider string
	Reason             string
	Timestamp          time.Time
}
