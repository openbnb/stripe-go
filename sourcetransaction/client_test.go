package sourcetransaction

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/openbnb/stripe-go/v55"
	_ "github.com/openbnb/stripe-go/v55/testing"
)

func TestSourceTransactionList(t *testing.T) {
	i := List(&stripe.SourceTransactionListParams{
		Source: stripe.String("src_123"),
	})

	// Verify that we can get at least one transaction
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.SourceTransaction())
}
