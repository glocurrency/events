package banking

import (
	"time"

	"github.com/google/uuid"
)

const PaymentComplianceRequiredEvent = "payment-compliance-required"
const PaymentProcessingEvent = "payment-processing"
const PaymentCompletedEvent = "payment-completed"
const PaymentFailedEvent = "payment-failed"

type DocumentMedia struct {
	Mimetype string
	Content  []byte
}

type IdentificationDocument struct {
	Type       string
	Number     string
	ValidUntil time.Time
	Media      *DocumentMedia
}

type PaymentComplianceRequired struct {
	PaymentID   uuid.UUID
	PaymentType string
	Sender      struct {
		FirstName string
		LastName  string
		DOB       time.Time
		Address   struct {
			Country     string
			County      string
			PostCode    string
			City        string
			StreetLine1 string
			StreetLine2 string
		}
		Document *IdentificationDocument
	}
	Recipient struct {
		FirstName string
		LastName  string
		Address   struct {
			Country string
		}
		Details struct {
			AccountNumber string
			BankCode      string
			SortCode      string
			RoutingNumber string
			IBAN          string
			CLABE         string
			BIC           string
			Phone         string
			Provider      string
			Email         string
		}
	}
	Currency  string
	Amount    float64
	Reference string
	Timestamp time.Time
}

type PaymentProcessing struct {
	PaymentID uuid.UUID
	Reason    string
	Timestamp time.Time
}

type PaymentCompleted struct {
	PaymentID uuid.UUID
	Reason    string
	Timestamp time.Time
}

type PaymentFailed struct {
	PaymentID uuid.UUID
	Reason    string
	Timestamp time.Time
}
