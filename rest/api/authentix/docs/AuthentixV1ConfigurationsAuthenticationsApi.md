# AuthentixV1ConfigurationsAuthenticationsApi

All URIs are relative to *https://api.reach.talkylabs.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAuthentication**](AuthentixV1ConfigurationsAuthenticationsApi.md#FetchAuthentication) | **Get** /rest/authentix/v1/configurations/{configurationId}/authentications/{authenticationId} | fetch an authentication
[**ListAuthentications**](AuthentixV1ConfigurationsAuthenticationsApi.md#ListAuthentications) | **Get** /rest/authentix/v1/configurations/{configurationId}/authentications | list all available authentications
[**StartAuthentication**](AuthentixV1ConfigurationsAuthenticationsApi.md#StartAuthentication) | **Post** /rest/authentix/v1/configurations/{configurationId}/authentications | start a new authentication process by sending a code via a channel
[**UpdateAuthenticationStatus**](AuthentixV1ConfigurationsAuthenticationsApi.md#UpdateAuthenticationStatus) | **Post** /rest/authentix/v1/configurations/{configurationId}/authentications/{authenticationId} | update an authentication status



## FetchAuthentication

> AuthenticationItem FetchAuthentication(ctx, ConfigurationIdAuthenticationId)

fetch an authentication

This operation allows to fetch the API record associated to an authentication.   

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConfigurationId** | **string** | The identifier of the configuration being used.
**AuthenticationId** | **string** | The identifier of the authentication to be fetched.

### Other Parameters

Other parameters are passed through a pointer to a FetchAuthenticationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AuthenticationItem**](AuthenticationItem.md)

### Authorization

[ApiKey](../README.md#ApiKey), [ApiUser](../README.md#ApiUser)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthentications

> []AuthenticationItem ListAuthentications(ctx, ConfigurationIdoptional)

list all available authentications

This operation allows to retrieve all authentications generated from a given configuration that are not expired and for which the number of maximum trials/controls is not exceeded.  When getting the authentication list, results will be sorted based  on the `dateCreated` field with the most recent record appearing first. 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConfigurationId** | **string** | The identifier of the configuration being used.

### Other Parameters

Other parameters are passed through a pointer to a ListAuthenticationsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | Maximum number of records to return per call.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]AuthenticationItem**](AuthenticationItem.md)

### Authorization

[ApiKey](../README.md#ApiKey), [ApiUser](../README.md#ApiUser)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartAuthentication

> AuthenticationItem StartAuthentication(ctx, ConfigurationIdoptional)

start a new authentication process by sending a code via a channel

This operation allows to start a new authentication process by sending a code via a specific channel. 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConfigurationId** | **string** | The identifier of the configuration being used.

### Other Parameters

Other parameters are passed through a pointer to a StartAuthenticationParams struct


Name | Type | Description
------------- | ------------- | -------------
**Dest** | **string** | The phone number or email where to send the authentication code. Phone numbers must be in E.164 format.
**Channel** | **string** | The channel by which the authentication code is sent.
**ServiceName** | **string** | a service name overwriting the one defined in the configuration.
**CustomCode** | **string** | the pre-generated code to be sent. Its length should be between 4 and 10 inclusive.
**PaymentInfo** | **string** | Information related to the digital payment to authenticate. It is required when `usedForDigitalPayment` is true. It is ignored otherwise. It is a stringfied JSON map where keys are `payee`, `amount`, and `currency` and the associated values are respectively the payee, the amount, and the currency of a financial transaction. 
**TemplateId** | **string** | This is the ID of the message template to use for sending the authenetication code. It could be an sms or email template depending on the channel being used. It overwirites the template ID defined in the configuration if any. 
**TemplateDataMap** | **string** | A stringfied JSON map where keys are message template parameters and the values are the parameter values to be used when sending the authentication code. It may also be used to provide additional parameters for sending email based authentications such as the email used for sending the code.

### Return type

[**AuthenticationItem**](AuthenticationItem.md)

### Authorization

[ApiKey](../README.md#ApiKey), [ApiUser](../README.md#ApiUser)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthenticationStatus

> AuthenticationItem UpdateAuthenticationStatus(ctx, ConfigurationIdAuthenticationIdoptional)

update an authentication status

This operation allows to manually update the status of an authentication. This should only be used with custom code authentications.  Set the authentication status to \"passed\" after you validated the authentication code. Similarly, set the authentication status to \"canceled\" if you want to restart start a new authentication with a different code.  When not using custom codes, there is no need to used this method as the REACH Authentix API can manage the whole life cycle of an authentication. 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConfigurationId** | **string** | The identifier of the configuration being used.
**AuthenticationId** | **string** | The identifier of the authentication to be updated.

### Other Parameters

Other parameters are passed through a pointer to a UpdateAuthenticationStatusParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | **string** | The new status of the authentication.

### Return type

[**AuthenticationItem**](AuthenticationItem.md)

### Authorization

[ApiKey](../README.md#ApiKey), [ApiUser](../README.md#ApiUser)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

