# MessageItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppletId** | **string** | The identifier of the applet sending the message. |[optional] 
**ApiVersion** | **string** | The API version used to process the message. |[optional] 
**Body** | **string** | The message text. |[optional] 
**Dest** | **string** | The phone number in E.164 format that received the message. |[optional] 
**Src** | **string** | The phone number (in E.164 format), or the alphanumeric sender ID that initiated the message. |[optional] 
**BulkId** | **string** | The bulk identifier allowing to group messages together and have corresponding statistics. |[optional] 
**NumSegments** | **int** | The number of segments associated to the message. A message body that is too large to be sent in a single SMS is segmented and charged as multiple messages. The segments are reassembled once received by the destination phone. A message can have a maximum of 10 segments.  |[optional] 
**NumMedia** | **int** | The number of media files included in the message |[optional] 
**Price** | **float32** | The cost billed for the message, in the currency specified by `priceUnit`. |[optional] 
**PriceUnit** | **string** | The currency, in ISO 4127 format, in which price is measured. for example, usd, xaf, eur, cad. |[optional] 
**MessageId** | **string** | The identifier of the message |[optional] 
**Status** | **string** | The status of the message. Can be: `sent`, `scheduled`, `failed`, `delivered`, `undelivered`, `canceled`, `accepted`, `queued`, `sending`, `received`, `receiving`.   |[optional] 
**MessageType** | **string** | The type of the message. Can be: `inbound` for incoming messages, `outbound` for messages initiated by a REST API.  |[optional] 
**ErrorCode** | **int** | The error code returned if the message status is `failed` or `undelivered`. The errorMessage provides more information about the failure. The value is null if the message is successful.  |[optional] 
**ErrorMessage** | **string** | The error message returned if the message status is `failed` or `undelivered`.  The value is null if the message is successful.  |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date and time in GMT that the message was created.  |[optional] 
**DateSent** | [**time.Time**](time.Time.md) | The date and time in GMT that the message was sent.  |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date and time in GMT that the message status was last updated.  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


