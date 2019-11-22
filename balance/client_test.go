package balance

import (
	"testing"

	_ "github.com/openbnb/stripe-go/testing"
	assert "github.com/stretchr/testify/require"
)

func TestBalanceGet(t *testing.T) {
	balance, err := Get(nil)
	assert.Nil(t, err)
	assert.NotNil(t, balance)
}
