package auth_test

import (
	"encoding/json"
	"testing"

	"github.com/glocurrency/events/auth"
	"github.com/stretchr/testify/require"
)

func TestAuthEvents_BasicMarshalling(t *testing.T) {
	events := []any{
		auth.UserRegistered{},
	}

	for _, event := range events {
		_, err := json.Marshal(event)
		require.NoError(t, err, "Failed to marshal struct of type %T", event)
	}
}
