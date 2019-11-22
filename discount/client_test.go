package discount

import (
	"testing"

	_ "github.com/openbnb/stripe-go/testing"
	assert "github.com/stretchr/testify/require"
)

func TestDiscountDel(t *testing.T) {
	discount, err := Del("cus_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, discount)
}

func TestDiscountDelSub(t *testing.T) {
	discount, err := DelSubscription("sub_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, discount)
}
