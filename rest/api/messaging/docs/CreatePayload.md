# CreatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dest** | **string** | The destination phone number in E.164 format of the message. |
**Src** | **string** | The phone number (in E.164 format), or the alphanumeric sender ID that initiated the message. |
**Body** | **string** | The text of the message to send. It can be up to 1,600 GSM-7 characters in length. That limit varies if your message is not made of only GSM-7 characters. More generally, the message body should not exceed 10 segments. |
**BulkIdentifier** | **string** | The identifier of the bulk operation this message belongs to. |[optional] 
**ScheduledTime** | [**time.Time**](time.Time.md) | The datetime at which the message will be sent. Must be in ISO 8601 format. A message must be scheduled at least 15 min in advance of message send time and cannot be scheduled more than 5 days in advance of the request. |[optional] 
**StatusCallback** | **string** | The URL that will be called to send status information of your message. If provided, the API POST these message status changes to the URL: `queued`, `failed`, `sent`, `canceled`, `delivered`, or `undelivered`. URLs must contain a valid hostname and underscores are not allowed.  |[optional] 
**MaxPrice** | **float32** | The maximum total price in the applet currency that should be paid for the message to be delivered. If the cost exceeds `maxPrice`, the message will fail and a status of `failed` is sent to the status callback.  |[optional] 
**ValidityPeriod** | **int** | It represents how long, in seconds, the message can remain in the queue. After this period elapses, the message fails and the status callback is called. After a message has been accepted by a carrier, however, there is no guarantee that the message will not be queued after this period. It is recommended that this value be at least 5 seconds. The maximum allowed value is 14,400 which corresponds to 4 hours.  |[optional] [default to 14400]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


