package bank_test

import (
	"encoding/json"
	"testing"

	"github.com/glocurrency/events/bank"
	"github.com/stretchr/testify/require"
)

func TestAccountEvents_BasicMarshalling(t *testing.T) {
	events := []any{
		bank.FxCreated{},
		bank.FxSettled{},
		bank.FxFailed{},
	}

	for _, event := range events {
		_, err := json.Marshal(event)
		require.NoError(t, err, "Failed to marshal struct of type %T", event)
	}
}
