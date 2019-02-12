package balance

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/openbnb/stripe-go/v55"
	_ "github.com/openbnb/stripe-go/v55/testing"
)

func TestBalanceGet(t *testing.T) {
	balance, err := Get(nil)
	assert.Nil(t, err)
	assert.NotNil(t, balance)
}

func TestBalanceTransactionList(t *testing.T) {
	i := List(&stripe.BalanceTransactionListParams{})

	// Verify that we can get at least one transaction
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.BalanceTransaction())
}

func TestBalanceBalanceTransactionGet(t *testing.T) {
	transaction, err := GetBalanceTransaction("txn_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
}
