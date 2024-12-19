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
	Country    string
	Type       string
	Number     string
	ValidFrom  *time.Time
	ValidUntil *time.Time
	Media      *DocumentMedia
}

type Gender string

const (
	GenderMale   Gender = "MALE"
	GenderFemale Gender = "FEMALE"
)

type Profession string

const (
	ProfessionAgriculture Profession = "AGRICULTURE"
	ProfessionRetired     Profession = "RETURED"
	// TODO: add more professions
)

type PaymentComplianceRequired struct {
	PaymentID   uuid.UUID
	PaymentType string
	Sender      struct {
		FirstName    string
		LastName     string
		DOB          time.Time
		Phone        string
		Email        string
		Gender       Gender
		Profession   Profession
		PEP          bool
		Citizenchip  string
		BirthCountry string
		BirthCity    string
		Address      struct {
			Country     string
			County      string
			PostCode    string
			City        string
			StreetLine1 string
			StreetLine2 string
		}
		Currency string
		Amount   float64
		Document *IdentificationDocument
	}
	Recipient struct {
		FirstName string
		LastName  string
		Address   struct {
			Country     string
			County      string
			PostCode    string
			City        string
			StreetLine1 string
			StreetLine2 string
		}
		PayoutDetails struct {
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
		Currency string
		Amount   float64
	}
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
