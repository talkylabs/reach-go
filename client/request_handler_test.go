package client_test

import (
	"errors"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/talkylabs/reach-go/client"
)

func NewRequestHandler(username string, authToken string) *client.RequestHandler {
	c := NewClient(username, authToken)
	return client.NewRequestHandler(c)
}

func TestRequestHandler_BuildHostRawHostWithoutPeriods(t *testing.T) {
	requestHandler := NewRequestHandler("user", "pass")
	assert.Equal(t, "https://prism_mywebsite:4010", assertAndGetURL(t, requestHandler, "https://prism_mywebsite:4010"))
}

func TestRequestHandler_BuildUrlInvalidCTLCharacter(t *testing.T) {
	requestHandler := NewRequestHandler("user", "pass")
	rawURL := "https://api.mywebsite.com/ServiceId\n"
	parsedURL, err := requestHandler.BuildUrl(rawURL)

	expectedErr := url.Error{Op: "parse", URL: rawURL, Err: errors.New("net/url: invalid control character in URL")}
	assert.Equal(t, &expectedErr, err)
	assert.Equal(t, parsedURL, "")
}

func assertAndGetURL(t *testing.T, requestHandler *client.RequestHandler, rawURL string) string {
	parsedURL, err := requestHandler.BuildUrl(rawURL)
	assert.Nil(t, err)
	return parsedURL
}
