package bank_test

import (
	"encoding/json"
	"testing"

	"github.com/glocurrency/events/bank"
	"github.com/stretchr/testify/require"
)

func TestTransactionEvents_BasicMarshalling(t *testing.T) {
	events := []any{
		bank.TransactionCreated{},
		bank.TransactionSettled{},
		bank.TransactionFailed{},
	}

	for _, event := range events {
		_, err := json.Marshal(event)
		require.NoError(t, err, "Failed to marshal struct of type %T", event)
	}
}
