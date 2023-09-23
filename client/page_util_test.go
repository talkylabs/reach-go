package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestPageUtil_ReadLimits(t *testing.T) {
	assert.Equal(t, 5, ReadLimits(nil, setLimit(5)))
	assert.Equal(t, 5, ReadLimits(setPageSize(10), setLimit(5)))
	assert.Equal(t, 1000, ReadLimits(nil, setLimit(5000)))
	assert.Equal(t, 10, ReadLimits(setPageSize(10), nil))
	assert.Equal(t, 20, ReadLimits(nil, nil))
}

func setLimit(limit int) *int {
	return &limit
}

func setPageSize(pageSize int) *int {
	return &pageSize
}

func TestPageUtil_GetNextPageUri(t *testing.T) {
	payload := map[string]interface{}{
		"page":           0,
		"pageSize":       2,
		"totalPages":     3,
		"outOfPageRange": false,
		"totalMessages":  5,
	}
	baseUrl := "https://api.mywebsite.com/resource"
	nextPageUrl, err := getNextPageUrl(baseUrl, payload)
	assert.Nil(t, err)
	assert.Equal(t, "https://api.mywebsite.com/resource?pageSize=2&page=1", nextPageUrl)

	payload["page"] = 1
	baseUrl = "https://api.mywebsite.com/resource"
	nextPageUrl, err = getNextPageUrl(baseUrl, payload)
	assert.Nil(t, err)
	assert.Equal(t, "https://api.mywebsite.com/resource?pageSize=2&page=2", nextPageUrl)

	payload = map[string]interface{}{}
	nextPageUrl, err = getNextPageUrl(baseUrl, payload)
	assert.Nil(t, err)
	assert.Equal(t, "", nextPageUrl)
}

func getTestClient(t *testing.T) *MockBaseClient {
	mockCtrl := gomock.NewController(t)
	testClient := NewMockBaseClient(mockCtrl)
	testClient.EXPECT().AccountSid().DoAndReturn(func() string {
		return "AC222222222222222222222222222222"
	}).AnyTimes()

	testClient.EXPECT().SendRequest(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any()).
		DoAndReturn(func(method string, rawURL string, data url.Values,
			headers map[string]interface{}) (*http.Response, error) {
			response := map[string]interface{}{
				"messages": []map[string]interface{}{
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Message 0",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Message 1",
						"status":    "delivered",
					},
				},
				"page":           0,
				"pageSize":       2,
				"totalPages":     3,
				"outOfPageRange": false,
				"totalMessages":  5,
			}

			resp, _ := json.Marshal(response)

			return &http.Response{
				Body: io.NopCloser(bytes.NewReader(resp)),
			}, nil
		},
		)

	return testClient
}

type testResponse struct {
	Messages       []testMessage `json:"messages,omitempty"`
	Page           int           `json:"page,omitempty"`
	PageSize       int           `json:"pageSize,omitempty"`
	TotalPages     int           `json:"totalPages,omitempty"`
	OutOfPageRange bool          `json:"outOfPageRange,omitempty"`
	TotalMessages  int           `json:"totalMessages,omitempty"`
}

type testMessage struct {
	// The message text
	Body *string `json:"body,omitempty"`
	// The direction of the message
	Direction *string `json:"direction,omitempty"`
	// The phone number that initiated the message
	From *string `json:"from,omitempty"`
	// The status of the message
	Status *string `json:"status,omitempty"`
	// The phone number that received the message
	To *string `json:"to,omitempty"`
}

func getSomething(nextPageUrl string) (interface{}, error) {
	return nextPageUrl, nil
}

func TestPageUtil_GetNext(t *testing.T) {
	testClient := getTestClient(t)
	baseUrl := "https://api.mywebsite.com/myresourse?k=v"
	response, _ := testClient.SendRequest("get", "", nil, nil) //nolint:bodyclose
	ps := &testResponse{}
	_ = json.NewDecoder(response.Body).Decode(ps)

	nextPageUrl, err := GetNext(baseUrl, ps, getSomething)
	assert.Equal(t, "https://api.mywebsite.com/myresourse?k=v&pageSize=2&page=1", nextPageUrl)
	assert.Nil(t, err)

	nextPageUrl, err = GetNext(baseUrl, nil, getSomething)
	assert.Empty(t, nextPageUrl)
	assert.Nil(t, err)
}

func TestPageUtil_ToMap(t *testing.T) {
	testMap, err := toMap("invalid")
	assert.NotNil(t, err)
	assert.Nil(t, testMap)

	valid := testResponse{
		Messages:       nil,
		TotalPages:     0,
		Page:           0,
		PageSize:       1,
		TotalMessages:  0,
		OutOfPageRange: true,
	}
	testMap, err = toMap(valid)
	assert.Nil(t, err)
	assert.NotNil(t, testMap)
}
