/*
 * This code was generated by
 *  ___ ___   _   ___ _  _    _____ _   _    _  ___   ___      _   ___ ___      ___   _   ___     ___ ___ _  _ ___ ___    _ _____ ___  ___ 
 * | _ \ __| /_\ / __| || |__|_   _/_\ | |  | |/ | \ / / |    /_\ | _ ) __|___ / _ \ /_\ |_ _|__ / __| __| \| | __| _ \  /_\_   _/ _ \| _ \
 * |   / _| / _ \ (__| __ |___|| |/ _ \| |__| ' < \ V /| |__ / _ \| _ \__ \___| (_) / _ \ | |___| (_ | _|| .` | _||   / / _ \| || (_) |   /
 * |_|_\___/_/ \_\___|_||_|    |_/_/ \_\____|_|\_\ |_| |____/_/ \_\___/___/    \___/_/ \_\___|   \___|___|_|\_|___|_|_\/_/ \_\_| \___/|_|_\
 * 
 * Reach Messaging API
 * Reach SMS API helps you add robust messaging capabilities to your applications.  Using this REST API, you can * send SMS messages * track the delivery of sent messages * schedule SMS messages to send at a later time * retrieve and modify message history
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

    "time"
    "github.com/talkylabs/reach-go/client"
)


// Optional parameters for the method 'ListMessage'
type ListMessageParams struct {
    // Retrieve messages sent to only this phone number. The phone number in E.164 format of the message.
    Dest *string `json:"dest,omitempty"`
    // Retrieve messages sent from only this phone number, in E.164 format, or alphanumeric sender ID.
    Src *string `json:"src,omitempty"`
    // Retrieve only messages that are assocaited with this `bulkIdentifier`.
    BulkIdentifier *string `json:"bulkIdentifier,omitempty"`
    // Retrieve only messages sent at the specified date. Must be in ISO 8601 format.
    SentAt *time.Time `json:"sentAt,omitempty"`
    // Retrieve only messages sent after the specified datetime. Must be in ISO 8601 format.
    SentAfter *time.Time `json:"sentAfter,omitempty"`
    // Retrieve only messages sent before the specified datetime. Must be in ISO 8601 format.
    SentBefore *time.Time `json:"sentBefore,omitempty"`
    // Maximum number of records to return per call.
    PageSize *int `json:"pageSize,omitempty"`
    // Max number of records to return.
    Limit *int `json:"limit,omitempty"`
}

func (params *ListMessageParams) SetDest(Dest string) (*ListMessageParams){
    params.Dest = &Dest
    return params
}
func (params *ListMessageParams) SetSrc(Src string) (*ListMessageParams){
    params.Src = &Src
    return params
}
func (params *ListMessageParams) SetBulkIdentifier(BulkIdentifier string) (*ListMessageParams){
    params.BulkIdentifier = &BulkIdentifier
    return params
}
func (params *ListMessageParams) SetSentAt(SentAt time.Time) (*ListMessageParams){
    params.SentAt = &SentAt
    return params
}
func (params *ListMessageParams) SetSentAfter(SentAfter time.Time) (*ListMessageParams){
    params.SentAfter = &SentAfter
    return params
}
func (params *ListMessageParams) SetSentBefore(SentBefore time.Time) (*ListMessageParams){
    params.SentBefore = &SentBefore
    return params
}
func (params *ListMessageParams) SetPageSize(PageSize int) (*ListMessageParams){
    params.PageSize = &PageSize
    return params
}
func (params *ListMessageParams) SetLimit(Limit int) (*ListMessageParams){
    params.Limit = &Limit
    return params
}

// Retrieve a single page of Message records from the API. Request is executed immediately.
func (c *ApiService) PageMessage(params *ListMessageParams, pageNumber string) (*PaginatedMessageItemList, error) {
    path := "/rest/messaging/v1/list"

    
data := url.Values{}
headers := make(map[string]interface{})

if params != nil && params.Dest != nil {
    data.Set("dest", *params.Dest)
}
if params != nil && params.Src != nil {
    data.Set("src", *params.Src)
}
if params != nil && params.BulkIdentifier != nil {
    data.Set("bulkIdentifier", *params.BulkIdentifier)
}
if params != nil && params.SentAt != nil {
    data.Set("sentAt", fmt.Sprint((*params.SentAt).Format(time.RFC3339)))
}
if params != nil && params.SentAfter != nil {
    data.Set("sentAfter", fmt.Sprint((*params.SentAfter).Format(time.RFC3339)))
}
if params != nil && params.SentBefore != nil {
    data.Set("sentBefore", fmt.Sprint((*params.SentBefore).Format(time.RFC3339)))
}
if params != nil && params.PageSize != nil {
    data.Set("pageSize", fmt.Sprint(*params.PageSize))
}

    if pageNumber != "" {
        data.Set("page", pageNumber)
    }

    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &PaginatedMessageItemList{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Lists Message records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListMessage(params *ListMessageParams) ([]MessageItem, error) {
	response, errors := c.StreamMessage(params)

	records := make([]MessageItem, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Message records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamMessage(params *ListMessageParams) (chan MessageItem, chan error) {
	if params == nil {
		params = &ListMessageParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan MessageItem, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageMessage(params, "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamMessage(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}


func (c *ApiService) streamMessage(response *PaginatedMessageItemList, params *ListMessageParams, recordChannel chan MessageItem, errorChannel chan error) {
	curRecord := 1
	
	path := "/rest/messaging/v1/list"

    
data := url.Values{}

if params != nil && params.Dest != nil {
    data.Set("dest", *params.Dest)
}
if params != nil && params.Src != nil {
    data.Set("src", *params.Src)
}
if params != nil && params.BulkIdentifier != nil {
    data.Set("bulkIdentifier", *params.BulkIdentifier)
}
if params != nil && params.SentAt != nil {
    data.Set("sentAt", fmt.Sprint((*params.SentAt).Format(time.RFC3339)))
}
if params != nil && params.SentAfter != nil {
    data.Set("sentAfter", fmt.Sprint((*params.SentAfter).Format(time.RFC3339)))
}
if params != nil && params.SentBefore != nil {
    data.Set("sentBefore", fmt.Sprint((*params.SentBefore).Format(time.RFC3339)))
}
if params != nil && params.PageSize != nil {
    data.Set("pageSize", fmt.Sprint(*params.PageSize))
}

    
    baseURL, err0 := client.UrlWithoutPaginationInfo(c.baseURL+path, data)
    if err0 != nil {
		errorChannel <- err0
		return
	}

	for response != nil {
		responseRecords := response.Messages
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(baseURL, response, c.getNextPaginatedMessageItemList)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*PaginatedMessageItemList)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextPaginatedMessageItemList(nextPageUrl string) (interface{}, error) {
    if nextPageUrl == "" {
        return nil, nil
    }
    resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &PaginatedMessageItemList{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }
    return ps, nil
}
