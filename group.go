package events

// Events posted by the auth services.
const (
	AuthUserEvents = "auth-events-user"
)

// Events posted by the compliance services.
const (
	CompliancePaymentEvents = "compliance-events-payment"
)

// Events posted by the bank services.
const (
	BankAccountEvents     = "bank-events-account"
	BankTransactionEvents = "bank-events-transaction"
	BankFxEvents          = "bank-events-fx"
)

// Events posted by the banking service.
const (
	BankingPaymentEvents = "banking-events-payment"
)
