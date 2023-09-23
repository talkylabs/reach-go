// Package client provides internal utilities for the reach-go client library.
package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/talkylabs/reach-go/client/form"
)

// Credentials store user authentication credentials.
type Credentials struct {
	ApiUser string
	ApiKey  string
}

func NewCredentials(apiUser string, apiKey string) *Credentials {
	return &Credentials{ApiUser: apiUser, ApiKey: apiKey}
}

// Client encapsulates a standard HTTP backend with authorization.
type Client struct {
	*Credentials
	HTTPClient          *http.Client
	UserAgentExtensions []string
}

// default http Client should not follow redirects and return the most recent response.
func defaultHTTPClient() *http.Client {
	return &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: time.Second * 30,
	}
}

func (c *Client) basicAuth() (string, string) {
	return c.Credentials.ApiUser, c.Credentials.ApiKey
}

// SetTimeout sets the Timeout for HTTP requests.
func (c *Client) SetTimeout(timeout time.Duration) {
	if c.HTTPClient == nil {
		c.HTTPClient = defaultHTTPClient()
	}
	c.HTTPClient.Timeout = timeout
}

const (
	keepZeros = true
	delimiter = '.'
	escapee   = '\\'
)

func (c *Client) doWithErr(req *http.Request) (*http.Response, error) {
	client := c.HTTPClient

	if client == nil {
		client = defaultHTTPClient()
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Note that 3XX response codes are allowed for fetches
	if res.StatusCode < 200 || res.StatusCode >= 400 {
		err = &ReachRestError{}
		if decodeErr := json.NewDecoder(res.Body).Decode(err); decodeErr != nil {
			err = errors.Wrap(decodeErr, "error decoding the response for an HTTP error code: "+strconv.Itoa(res.StatusCode))
			return nil, err
		}

		return nil, err
	}
	return res, nil
}

// SendRequest verifies, constructs, and authorizes an HTTP request.
func (c *Client) SendRequest(method string, rawURL string, data url.Values,
	headers map[string]interface{}) (*http.Response, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	valueReader := &strings.Reader{}
	goVersion := runtime.Version()

	if method == http.MethodGet {
		if data != nil {
			v, _ := form.EncodeToStringWith(data, delimiter, escapee, keepZeros)
			regex := regexp.MustCompile(`\.\d+`)
			s := regex.ReplaceAllString(v, "")

			u.RawQuery = s
		}
	}

	if method == http.MethodPost {
		valueReader = strings.NewReader(data.Encode())
	}

	req, err := http.NewRequest(method, u.String(), valueReader)
	if err != nil {
		return nil, err
	}

	//req.SetBasicAuth(c.basicAuth())
	apiUser, apiKey := c.basicAuth()
	req.Header.Add("ApiUser", apiUser)
	req.Header.Add("ApiKey", apiKey)

	// E.g. "User-Agent": "reach-go/1.0.0 (darwin amd64) go/go1.17.8"
	userAgent := fmt.Sprintf("reach-go/%s (%s %s) go/%s", LibraryVersion, runtime.GOOS, runtime.GOARCH, goVersion)

	if len(c.UserAgentExtensions) > 0 {
		userAgent += " " + strings.Join(c.UserAgentExtensions, " ")
	}

	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Accept", "application/json")

	if method == http.MethodPost {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	for k, v := range headers {
		req.Header.Add(k, fmt.Sprint(v))
	}

	return c.doWithErr(req)
}

// Returns the Account SID.
func (c *Client) AccountSid() string {
	return c.ApiUser
}
