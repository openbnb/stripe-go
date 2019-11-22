package accountlink

import (
	"testing"

	stripe "github.com/openbnb/stripe-go"
	_ "github.com/openbnb/stripe-go/testing"
	assert "github.com/stretchr/testify/require"
)

func TestAccountLinkNew(t *testing.T) {
	params := &stripe.AccountLinkParams{
		Account:    stripe.String("acct_123"),
		Collect:    stripe.String(string(stripe.AccountLinkCollectCurrentlyDue)),
		FailureURL: stripe.String("https://stripe.com/failure"),
		SuccessURL: stripe.String("https://stripe.com/success"),
		Type:       stripe.String(string(stripe.AccountLinkTypeCustomAccountVerification)),
	}
	link, err := New(params)
	assert.Nil(t, err)
	assert.NotNil(t, link)
}
