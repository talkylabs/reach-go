// Package client provides internal utilities for the reach-go client library.
package client

import (
	"net/http"
	"net/url"
)

type RequestHandler struct {
	Client BaseClient
}

func NewRequestHandler(client BaseClient) *RequestHandler {
	return &RequestHandler{
		Client: client,
	}
}

func (c *RequestHandler) sendRequest(method string, rawURL string, data url.Values,
	headers map[string]interface{}) (*http.Response, error) {
	parsedURL, err := c.BuildUrl(rawURL)
	if err != nil {
		return nil, err
	}

	return c.Client.SendRequest(method, parsedURL, data, headers)
}

// BuildUrl builds the target host string taking into account region and edge configurations.
func (c *RequestHandler) BuildUrl(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	return u.String(), nil
}

func (c *RequestHandler) Post(path string, bodyData url.Values, headers map[string]interface{}) (*http.Response, error) {
	return c.sendRequest(http.MethodPost, path, bodyData, headers)
}

func (c *RequestHandler) Get(path string, queryData url.Values, headers map[string]interface{}) (*http.Response, error) {
	return c.sendRequest(http.MethodGet, path, queryData, headers)
}

func (c *RequestHandler) Delete(path string, nothing url.Values, headers map[string]interface{}) (*http.Response, error) {
	return c.sendRequest(http.MethodDelete, path, nil, headers)
}
