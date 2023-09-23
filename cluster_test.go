//go:build cluster
// +build cluster

package reach

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	Messaging "github.com/talkylabs/reach-go/rest/api/messaging"
)

var from string
var to string
var testClient *RestClient

func TestMain(m *testing.M) {
	from = os.Getenv("REACH_TALKYLABS_FROM_NUMBER")
	to = os.Getenv("REACH_TALKYLABS_TO_NUMBER")
	var apiUser = os.Getenv("REACH_TALKYLABS_API_USER")
	var apiKey = os.Getenv("REACH_TALKYLABS_API_KEY")

	testClient = NewRestClientWithParams(ClientParams{apiUser, apiKey, nil})
	ret := m.Run()
	os.Exit(ret)
}

func TestSendingAText(t *testing.T) {
	params := &Messaging.SendMessageParams{}
	params.SetDest(to)
	params.SetSrc(from)
	params.SetBody("Hello there")

	resp, err := testClient.Messaging.SendMessage(params)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Hello there", *resp.Body)
	assert.Equal(t, from, *resp.Dest)
	assert.Equal(t, to, *resp.Src)
}
