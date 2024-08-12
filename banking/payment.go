package banking

import (
	"time"

	"github.com/google/uuid"
)

const PaymentComplianceRequiredEvent = "payment-compliance-required"
const PaymentProcessingEvent = "payment-processing"
const PaymentCompletedEvent = "payment-completed"
const PaymentFailedEvent = "payment-failed"

type PaymentComplianceRequired struct {
	PaymentID uuid.NullUUID
	Sender    struct {
		FirstName   string
		LastName    string
		DOB         time.Time
		Country     string
		County      string
		PostCode    string
		City        string
		StreetLine1 string
		StreetLine2 string
	}
	Recipient struct {
		FirstName     string
		LastName      string
		Country       string
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
