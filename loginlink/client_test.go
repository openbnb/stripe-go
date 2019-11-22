package loginlink

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/openbnb/stripe-go"
	_ "github.com/openbnb/stripe-go/testing"
)

func TestLoginLinkNew(t *testing.T) {
	link, err := New(&stripe.LoginLinkParams{
		Account: stripe.String("acct_EXPRESS"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, link)
}
