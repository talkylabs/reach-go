# AuthentixV1AuthenticationTrialStatsApi

All URIs are relative to *https://api.reach.talkylabs.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAuthenticationTrialStats**](AuthentixV1AuthenticationTrialStatsApi.md#FetchAuthenticationTrialStats) | **Get** /rest/authentix/v1/authenticationTrialStats | get the success rate statistics of the authentication trials matching some criteria



## FetchAuthenticationTrialStats

> AuthenticationTrialStatItem FetchAuthenticationTrialStats(ctx, optional)

get the success rate statistics of the authentication trials matching some criteria

This operation allows to retrieve from the API success rate statistics of authentication trial records that satisfied specified criteria.   

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchAuthenticationTrialStatsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Dest** | **string** | Filter authentication trials sent only to this phone number or email. The phone number must be in the E.164 format.
**TrialStatus** | **string** | Filter authentication trials with the specified status.
**Channel** | **string** | Filter authentication trials sent via the specified channel.
**ConfigurationId** | **string** | Filter authentication trials from the configuration whose ID matches the specified one.
**AuthenticationId** | **string** | Filter authentication trials from the authentication whose ID matches the specified one.
**Country** | **string** | Filter authentication trials sent to the specified destination country (in ISO 3166-1 alpha-2). Only possible when `dest` is a phone number.
**SentAt** | **time.Time** | Filter authentication trials created at the specified date. Must be in ISO 8601 format.
**SentAfter** | **time.Time** | Filter authentication trials created after the specified datetime. Must be in ISO 8601 format.
**SentBefore** | **time.Time** | Filter authentication trials created before the specified datetime. Must be in ISO 8601 format.

### Return type

[**AuthenticationTrialStatItem**](AuthenticationTrialStatItem.md)

### Authorization

[ApiKey](../README.md#ApiKey), [ApiUser](../README.md#ApiUser)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

