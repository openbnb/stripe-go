package connectiontoken

import (
	"testing"

	stripe "github.com/openbnb/stripe-go"
	_ "github.com/openbnb/stripe-go/testing"
	assert "github.com/stretchr/testify/require"
)

func TestTerminalConnectionTokenNew(t *testing.T) {
	connectiontoken, err := New(&stripe.TerminalConnectionTokenParams{})
	assert.Nil(t, err)
	assert.NotNil(t, connectiontoken)
}
