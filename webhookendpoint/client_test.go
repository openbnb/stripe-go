package webhookendpoint

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/openbnb/stripe-go/v55"
	_ "github.com/openbnb/stripe-go/v55/testing"
)

func TestWebhookEndpointDel(t *testing.T) {
	endpoint, err := Del("we_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, endpoint)
}

func TestWebhookEndpointGet(t *testing.T) {
	endpoint, err := Get("we_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, endpoint)
}

func TestWebhookEndpointList(t *testing.T) {
	i := List(&stripe.WebhookEndpointListParams{})

	// Verify that we can get at least one webhook endpoint
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.WebhookEndpoint())
}

func TestWebhookEndpointNew(t *testing.T) {
	endpoint, err := New(&stripe.WebhookEndpointParams{
		EnabledEvents: []*string{
			stripe.String("charge.succeeded"),
		},
		URL: stripe.String("https://stripe.com"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, endpoint)
}

func TestWebhookEndpointUpdate(t *testing.T) {
	endpoint, err := Update("we_123", &stripe.WebhookEndpointParams{
		EnabledEvents: []*string{
			stripe.String("charge.succeeded"),
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, endpoint)
}
