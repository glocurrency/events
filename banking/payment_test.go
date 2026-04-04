package banking_test

import (
	"encoding/json"
	"testing"

	"github.com/glocurrency/events/banking"
	"github.com/stretchr/testify/require"
)

func TestBankingEvents_BasicMarshalling(t *testing.T) {
	events := []any{
		banking.PaymentComplianceRequired{},
		banking.PaymentProcessing{},
		banking.PaymentCompleted{},
		banking.PaymentFailed{},
		banking.IdentificationDocument{},
		banking.DocumentMedia{},
	}

	for _, event := range events {
		_, err := json.Marshal(event)
		require.NoError(t, err, "Failed to marshal struct of type %T", event)
	}
}
