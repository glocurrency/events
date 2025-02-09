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

type DocumentType string

const (
	DocumentTypePassport        DocumentType = "PASSPORT"
	DocumentTypeIDCard          DocumentType = "ID_CARD"
	DocumentTypeDriversLicense  DocumentType = "DRIVERS_LICENSE"
	DocumentTypeResidencePermit DocumentType = "RESIDENCE_PERMIT"
)

type IdentificationDocument struct {
	Country    string
	Type       DocumentType
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
	ProfessionAgriculture    Profession = "AGRICULTURE"
	ProfessionArchitect      Profession = "ARCHITECT"
	ProfessionDefence        Profession = "DEFENCE"
	ProfessionArtist         Profession = "ARTIST"
	ProfessionAntiquesTrader Profession = "ANTIQUES_TRADER"
	ProfessionAssistantNurse Profession = "ASSISTANT_NURSE"
	ProfessionBanking        Profession = "BANKING"
	ProfessionBusinessAdmin  Profession = "BUSINESS_ADMIN"
	ProfessionCarDealer      Profession = "CAR_DEALER"
	ProfessionCivilServant   Profession = "CIVIL_SERVANT"
	ProfessionCleaning       Profession = "CLEANING"
	ProfessionAdminOfficer   Profession = "ADMIN_OFFICER"
	ProfessionConstruction   Profession = "CONSTRUCTION"
	ProfessionDiplomat       Profession = "DIPLOMAT"
	ProfessionFactoryWorker  Profession = "FACTORY_WORKER"
	ProfessionTransport      Profession = "TRANSPORT"
	ProfessionGambling       Profession = "GAMBLING"
	ProfessionFitness        Profession = "FITNESS"
	ProfessionHealthcare     Profession = "HEALTHCARE"
	ProfessionHospitality    Profession = "HOSPITALITY"
	ProfessionJewelry        Profession = "JEWELRY"
	ProfessionJudge          Profession = "JUDGE"
	ProfessionLegal          Profession = "LEGAL"
	ProfessionMerchant       Profession = "MERCHANT"
	ProfessionMilitary       Profession = "MILITARY"
	ProfessionMining         Profession = "MINING"
	ProfessionRetired        Profession = "RETIRED"
	ProfessionPharma         Profession = "PHARMA"
	ProfessionPolitics       Profession = "POLITICS"
	ProfessionSecurity       Profession = "SECURITY"
	ProfessionAthlete        Profession = "ATHLETE"
	ProfessionRealEstate     Profession = "REAL_ESTATE"
	ProfessionReligious      Profession = "RELIGIOUS"
	ProfessionTaxiDriver     Profession = "TAXI_DRIVER"
	ProfessionTeaching       Profession = "TEACHING"
	ProfessionTobacco        Profession = "TOBACCO"
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
	Checkout struct {
		Provider          string
		ProviderReference string
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
