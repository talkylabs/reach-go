# MessagingV1FetchApi

All URIs are relative to *https://api.reach.talkylabs.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchMessage**](MessagingV1FetchApi.md#FetchMessage) | **Get** /rest/messaging/v1/fetch | fetch a message record



## FetchMessage

> MessageItem FetchMessage(ctx, optional)

fetch a message record

This operation allows to fetch the API record associated to a message.  This operation needs the `messageId` of the message to be fetched.   

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**MessageId** | **string** | The identifier of the message to be updated.

### Return type

[**MessageItem**](MessageItem.md)

### Authorization

[ApiKey](../README.md#ApiKey), [ApiUser](../README.md#ApiUser)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

