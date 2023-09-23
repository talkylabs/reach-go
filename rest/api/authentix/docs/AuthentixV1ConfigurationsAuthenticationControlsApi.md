# AuthentixV1ConfigurationsAuthenticationControlsApi

All URIs are relative to *https://api.reach.talkylabs.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckAuthenticationCode**](AuthentixV1ConfigurationsAuthenticationControlsApi.md#CheckAuthenticationCode) | **Post** /rest/authentix/v1/configurations/{configurationId}/authentication-controls | check the correctness of a code provided by the end user



## CheckAuthenticationCode

> AuthenticationControlItem CheckAuthenticationCode(ctx, ConfigurationIdoptional)

check the correctness of a code provided by the end user

This operation allows check the correctness of a code provided by the end user. This operations only concerns authentication in `awaiting` state.   In case a check is performs on an authentication whose status is different from `awaiting`, an 404 HTTP response is returned. 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConfigurationId** | **string** | The identifier of the configuration being used.

### Other Parameters

Other parameters are passed through a pointer to a CheckAuthenticationCodeParams struct


Name | Type | Description
------------- | ------------- | -------------
**Dest** | **string** | The phone number or email being authenticated. Phone numbers must be in E.164 format. Either this parameter or the `authenticationId` must be specified.
**Code** | **string** | The 4-10 character string being verified. This is required for `sms` and `email` channels.
**AuthenticationId** | **string** | The ID of the authentication being checked. Either this parameter or the to `dest` must be specified.
**PaymentInfo** | **string** | Information related to the digital payment to authenticate. It is required when `usedForDigitalPayment` is true. It is ignored otherwise. It is a stringfied JSON map where keys are `payee`, `amount`, and `currency` and the associated values are respectively the payee, the amount, and the currency of a financial transaction. 

### Return type

[**AuthenticationControlItem**](AuthenticationControlItem.md)

### Authorization

[ApiKey](../README.md#ApiKey), [ApiUser](../README.md#ApiUser)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

