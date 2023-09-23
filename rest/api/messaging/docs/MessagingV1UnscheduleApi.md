# MessagingV1UnscheduleApi

All URIs are relative to *https://api.reach.talkylabs.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UnscheduleMessage**](MessagingV1UnscheduleApi.md#UnscheduleMessage) | **Post** /rest/messaging/v1/unschedule | cancel a previously scheduled message



## UnscheduleMessage

> MessageItem UnscheduleMessage(ctx, optional)

cancel a previously scheduled message

This operation allows to cancel a previously scheduled message.  This operation needs the `messageId` of the message to be unscheduled. To unschedule multiple messages, this operation should be called as many times needed since it can only unschedule one message at a time.  Note: The system will make the best attempt to cancel a scheduled message when a request is received. 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UnscheduleMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**MessageId** | **string** | The identifier of the message to be unscheduled.

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

