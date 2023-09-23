# MessagingV1CreateApi

All URIs are relative to *https://api.reach.talkylabs.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendMessage**](MessagingV1CreateApi.md#SendMessage) | **Post** /rest/messaging/v1/create | send a message



## SendMessage

> MessageItem SendMessage(ctx, optional)

send a message

This operation allows to send or schedule a message.  When sending a new message via the API, you must include the `dest` parameter.          This value should be the destination phone number. You must also include the `body` parameter containing the message's content as well as the `src` parameter   containing the sender alphanumeric Id or number.  To schedule a message, you must additionally pass the following parameter:  * `scheduledTime`: the date and time at which the sms will be sent in the ISO-8601 format. 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a SendMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**Dest** | **string** | The destination phone number in E.164 format of the message.
**Src** | **string** | The phone number (in E.164 format), or the alphanumeric sender ID that initiated the message.
**Body** | **string** | The text of the message to send. It can be up to 1,600 GSM-7 characters in length. That limit varies if your message is not made of only GSM-7 characters. More generally, the message body should not exceed 10 segments.
**BulkIdentifier** | **string** | The identifier of the bulk operation this message belongs to.
**ScheduledTime** | **time.Time** | The datetime at which the message will be sent. Must be in ISO 8601 format. A message must be scheduled at least 15 min in advance of message send time and cannot be scheduled more than 5 days in advance of the request.
**StatusCallback** | **string** | The URL that will be called to send status information of your message. If provided, the API POST these message status changes to the URL: `queued`, `failed`, `sent`, `canceled`, `delivered`, or `undelivered`. URLs must contain a valid hostname and underscores are not allowed. 
**MaxPrice** | **float32** | The maximum total price in the applet currency that should be paid for the message to be delivered. If the cost exceeds `maxPrice`, the message will fail and a status of `failed` is sent to the status callback. 
**ValidityPeriod** | **int** | It represents how long, in seconds, the message can remain in the queue. After this period elapses, the message fails and the status callback is called. After a message has been accepted by a carrier, however, there is no guarantee that the message will not be queued after this period. It is recommended that this value be at least 5 seconds. The maximum allowed value is 14,400 which corresponds to 4 hours. 

### Return type

[**MessageItem**](MessageItem.md)

### Authorization

[ApiKey](../README.md#ApiKey), [ApiUser](../README.md#ApiUser)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

