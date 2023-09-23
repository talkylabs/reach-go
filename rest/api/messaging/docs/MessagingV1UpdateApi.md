# MessagingV1UpdateApi

All URIs are relative to *https://api.reach.talkylabs.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateMessage**](MessagingV1UpdateApi.md#UpdateMessage) | **Post** /rest/messaging/v1/update | update the content of a message



## UpdateMessage

> MessageItem UpdateMessage(ctx, optional)

update the content of a message

This operation allows to update the body of a message. It is primarily used to redact the content of message while leaving all the other properties untouched.  This operation needs the `messageId` of the message to be updated. It also requires the `body` that will be newly associated with the message. To update multiple messages, this operation should be called as many times as needed since it can only update one message at a time.  Note: The previous body of the message is the one that is sent to the destination phone number. This operation just update the `body` in the API platform. 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**MessageId** | **string** | The identifier of the message to be updated.
**Body** | **string** | The text to be newly associated with the message.

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

