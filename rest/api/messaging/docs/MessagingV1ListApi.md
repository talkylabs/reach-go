# MessagingV1ListApi

All URIs are relative to *https://api.reach.talkylabs.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMessage**](MessagingV1ListApi.md#ListMessage) | **Get** /rest/messaging/v1/list | list message that match some criteria



## ListMessage

> []MessageItem ListMessage(ctx, optional)

list message that match some criteria

This operation allows to retrieve from the API message records that satisfied specified criteria.  When getting the message record list, results will be sorted on the `dateSent` field with the most recent message records appearing first.   

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**Dest** | **string** | Retrieve messages sent to only this phone number. The phone number in E.164 format of the message.
**Src** | **string** | Retrieve messages sent from only this phone number, in E.164 format, or alphanumeric sender ID.
**BulkIdentifier** | **string** | Retrieve only messages that are assocaited with this `bulkIdentifier`.
**SentAt** | **time.Time** | Retrieve only messages sent at the specified date. Must be in ISO 8601 format.
**SentAfter** | **time.Time** | Retrieve only messages sent after the specified datetime. Must be in ISO 8601 format.
**SentBefore** | **time.Time** | Retrieve only messages sent before the specified datetime. Must be in ISO 8601 format.
**PageSize** | **int** | Maximum number of records to return per call.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]MessageItem**](MessageItem.md)

### Authorization

[ApiKey](../README.md#ApiKey), [ApiUser](../README.md#ApiUser)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

