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
	"github.com/talkylabs/reach-go/client"
	"time"
)
// MessageItem struct for MessageItem
type MessageItem struct {
		// The identifier of the applet sending the message.
	AppletId string `json:"appletId,omitempty"`
		// The API version used to process the message.
	ApiVersion string `json:"apiVersion,omitempty"`
		// The message text.
	Body string `json:"body,omitempty"`
		// The phone number in E.164 format that received the message.
	Dest string `json:"dest,omitempty"`
		// The phone number (in E.164 format), or the alphanumeric sender ID that initiated the message.
	Src string `json:"src,omitempty"`
		// The bulk identifier allowing to group messages together and have corresponding statistics.
	BulkId string `json:"bulkId,omitempty"`
		// The number of segments associated to the message. A message body that is too large to be sent in a single SMS is segmented and charged as multiple messages. The segments are reassembled once received by the destination phone. A message can have a maximum of 10 segments. 
	NumSegments int `json:"numSegments,omitempty"`
		// The number of media files included in the message
	NumMedia int `json:"numMedia,omitempty"`
		// The cost billed for the message, in the currency specified by `priceUnit`.
	Price float32 `json:"price,omitempty"`
		// The currency, in ISO 4127 format, in which price is measured. for example, usd, xaf, eur, cad.
	PriceUnit string `json:"priceUnit,omitempty"`
		// The identifier of the message
	MessageId string `json:"messageId,omitempty"`
		// The status of the message. Can be: `sent`, `scheduled`, `failed`, `delivered`, `undelivered`, `canceled`, `accepted`, `queued`, `sending`, `received`, `receiving`.  
	Status string `json:"status,omitempty"`
		// The type of the message. Can be: `inbound` for incoming messages, `outbound` for messages initiated by a REST API. 
	MessageType string `json:"messageType,omitempty"`
		// The error code returned if the message status is `failed` or `undelivered`. The errorMessage provides more information about the failure. The value is null if the message is successful. 
	ErrorCode int `json:"errorCode,omitempty"`
		// The error message returned if the message status is `failed` or `undelivered`.  The value is null if the message is successful. 
	ErrorMessage string `json:"errorMessage,omitempty"`
		// The date and time in GMT that the message was created. 
	DateCreated time.Time `json:"dateCreated,omitempty"`
		// The date and time in GMT that the message was sent. 
	DateSent time.Time `json:"dateSent,omitempty"`
		// The date and time in GMT that the message status was last updated. 
	DateUpdated time.Time `json:"dateUpdated,omitempty"`
}

func (response *MessageItem) UnmarshalJSON(bytes []byte) (err error) {
	raw := struct {
		AppletId string `json:"appletId"`
		ApiVersion string `json:"apiVersion"`
		Body string `json:"body"`
		Dest string `json:"dest"`
		Src string `json:"src"`
		BulkId string `json:"bulkId"`
		NumSegments int `json:"numSegments"`
		NumMedia int `json:"numMedia"`
		Price interface{} `json:"price"`
		PriceUnit string `json:"priceUnit"`
		MessageId string `json:"messageId"`
		Status string `json:"status"`
		MessageType string `json:"messageType"`
		ErrorCode int `json:"errorCode"`
		ErrorMessage string `json:"errorMessage"`
		DateCreated time.Time `json:"dateCreated"`
		DateSent time.Time `json:"dateSent"`
		DateUpdated time.Time `json:"dateUpdated"`
	}{}

	if err = json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	*response = MessageItem{
		AppletId: raw.AppletId,
		ApiVersion: raw.ApiVersion,
		Body: raw.Body,
		Dest: raw.Dest,
		Src: raw.Src,
		BulkId: raw.BulkId,
		NumSegments: raw.NumSegments,
		NumMedia: raw.NumMedia,
		PriceUnit: raw.PriceUnit,
		MessageId: raw.MessageId,
		Status: raw.Status,
		MessageType: raw.MessageType,
		ErrorCode: raw.ErrorCode,
		ErrorMessage: raw.ErrorMessage,
		DateCreated: raw.DateCreated,
		DateSent: raw.DateSent,
		DateUpdated: raw.DateUpdated,
	}

	responsePrice, err := client.UnmarshalFloat32(&raw.Price)
	if err != nil {
		return err
	}
	response.Price = *responsePrice

	return
}

