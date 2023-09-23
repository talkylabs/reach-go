package reach

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_WithNoAccountSid(t *testing.T) {
	client := NewRestClientWithParams(ClientParams{
		ApiUser: "parentSid",
		ApiKey: "authToken",
	})
	assert.Equal(t, client.RequestHandler.Client.AccountSid(), "parentSid")
}
