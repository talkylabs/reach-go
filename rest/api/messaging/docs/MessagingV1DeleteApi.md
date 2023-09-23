# MessagingV1DeleteApi

All URIs are relative to *https://api.reach.talkylabs.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMessage**](MessagingV1DeleteApi.md#DeleteMessage) | **Delete** /rest/messaging/v1/delete | delete a message



## DeleteMessage

> DeleteMessage(ctx, optional)

delete a message

This operation allows to delete a message record from the applet account. Once the record is deleted, it will no longer appear in the API and the applet portal.  This operation needs the `messageId` of the message to be deleted. To delete multiple messages, this operation should be called as many times as needed since it can only delete one message at a time. Note: Attempting to delete an in-progress message record, i.e. a message whose status is not `delivered`, `failed`, `canceled`, `undelivered`, will result in an error.   

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a DeleteMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**MessageId** | **string** | The identifier of the message to be updated.

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey), [ApiUser](../README.md#ApiUser)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

