# AuthentixV1AuthenticationTrialsApi

All URIs are relative to *https://api.reach.talkylabs.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAthenticationTrial**](AuthentixV1AuthenticationTrialsApi.md#FetchAthenticationTrial) | **Get** /rest/authentix/v1/authenticationTrials/{trialId} | fetch a authentication trial
[**ListAuthenticationTrials**](AuthentixV1AuthenticationTrialsApi.md#ListAuthenticationTrials) | **Get** /rest/authentix/v1/authenticationTrials | list message that match some criteria



## FetchAthenticationTrial

> AuthenticationTrialItem FetchAthenticationTrial(ctx, TrialId)

fetch a authentication trial

This operation allows to fetch the API record associated with an authentication trial.   

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrialId** | **string** | The identifier of the authentication trial to be fetched.

### Other Parameters

Other parameters are passed through a pointer to a FetchAthenticationTrialParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AuthenticationTrialItem**](AuthenticationTrialItem.md)

### Authorization

[ApiKey](../README.md#ApiKey), [ApiUser](../README.md#ApiUser)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthenticationTrials

> []AuthenticationTrialItem ListAuthenticationTrials(ctx, optional)

list message that match some criteria

This operation allows to retrieve from the API authentication trial records that satisfied specified criteria.  When getting the record list, results will be sorted based on the `dateCreated` field with the most recent record appearing first.   

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAuthenticationTrialsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Dest** | **string** | Retrieve authentication trials sent only to this phone number or email. The phone number must be in the E.164 format.
**TrialStatus** | **string** | Retrieve authentication trials with the specified status.
**Channel** | **string** | Retrieve authentication trials sent via the specified channel.
**ConfigurationId** | **string** | Retrieve authentication trials from the configuration whose ID matches the specified one.
**AuthenticationId** | **string** | Retrieve authentication trials from the authentication whose ID matches the specified one.
**Country** | **string** | Retrieve authentication trials sent to the specified destination country (in ISO 3166-1 alpha-2). Only possible when `dest` is a phone number.
**SentAt** | **time.Time** | Retrieve only authentication trials created at the specified date. Must be in ISO 8601 format.
**SentAfter** | **time.Time** | Retrieve only authentication trials created after the specified datetime. Must be in ISO 8601 format.
**SentBefore** | **time.Time** | Retrieve only authentication trials created before the specified datetime. Must be in ISO 8601 format.
**PageSize** | **int** | Maximum number of records to return per call.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]AuthenticationTrialItem**](AuthenticationTrialItem.md)

### Authorization

[ApiKey](../README.md#ApiKey), [ApiUser](../README.md#ApiUser)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

