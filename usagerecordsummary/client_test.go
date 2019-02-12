package usagerecordsummary

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/openbnb/stripe-go/v55"
	_ "github.com/openbnb/stripe-go/v55/testing"
)

func TestUsageRecordSummaryList(t *testing.T) {
	params := &stripe.UsageRecordSummaryListParams{
		SubscriptionItem: stripe.String("si_123"),
	}
	i := List(params)

	// Verify that we can get at least one usage record summary
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.UsageRecordSummary())
}
