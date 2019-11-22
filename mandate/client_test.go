package mandate

import (
	"testing"

	_ "github.com/openbnb/stripe-go/testing"
	assert "github.com/stretchr/testify/require"
)

func TestMandateMethodGet(t *testing.T) {
	pm, err := Get("mandate_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, pm)
}
