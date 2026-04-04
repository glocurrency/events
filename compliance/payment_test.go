package compliance_test

import (
	"encoding/json"
	"testing"

	"github.com/glocurrency/events/compliance"
	"github.com/stretchr/testify/require"
)

func TestMarshalling_Basic(t *testing.T) {
	events := []any{
		compliance.PaymentAccepted{},
		compliance.PaymentPending{},
		compliance.PaymentApproved{},
		compliance.PaymentDeclined{},
	}

	for _, event := range events {
		_, err := json.Marshal(event)
		require.NoError(t, err, "Failed to marshal struct of type %T", event)
	}
}
