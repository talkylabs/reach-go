# AuthentixV1ConfigurationsApi

All URIs are relative to *https://api.reach.talkylabs.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfiguration**](AuthentixV1ConfigurationsApi.md#CreateConfiguration) | **Post** /rest/authentix/v1/configurations | create a configuration
[**DeleteConfiguration**](AuthentixV1ConfigurationsApi.md#DeleteConfiguration) | **Delete** /rest/authentix/v1/configurations/{configurationId} | delete a configuration
[**FetchConfiguration**](AuthentixV1ConfigurationsApi.md#FetchConfiguration) | **Get** /rest/authentix/v1/configurations/{configurationId} | fetch a configuration
[**ListConfigurations**](AuthentixV1ConfigurationsApi.md#ListConfigurations) | **Get** /rest/authentix/v1/configurations | list all configurations
[**UpdateConfiguration**](AuthentixV1ConfigurationsApi.md#UpdateConfiguration) | **Post** /rest/authentix/v1/configurations/{configurationId} | update a configuration



## CreateConfiguration

> ConfigurationItem CreateConfiguration(ctx, optional)

create a configuration

This operation allows to create a Reach Authentix configuration. 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**ServiceName** | **string** | The name of the authentication service attached to this configuration. It can be up to 40 characters long.
**CodeLength** | **int** | The length of the code to be generated. It must be a value between 4 and 10, inclusive. If not specified, the default value is 5.
**AllowCustomCode** | **bool** | A flag indicating if the configuration should allow sending custom and non-generated code.
**UsedForDigitalPayment** | **bool** | A flag indicating if the configuration is used to authenticate digital payments. In such a case, additional information such as the amount and the payee of the financial transaction should be sent to when starting the authentication.
**DefaultExpiryTime** | **int** | It represents how long, in minutes, an authentication process will remained in the `awaiting` status before moving to `expired` in the case no valid matching is performed in between.   It also means that the code sent for the autentication remains the same during its validity period until the autentication is successful. In other words, if another authentication request is asked within that period, the same code will be sent.  If not specified, the default value is 15 minutes. It must be any value between 1 and 1440 which represents 24 hours. 
**DefaultMaxTrials** | **int** | It represents the maximum number of trials per authentication. The default value is 5. 
**DefaultMaxControls** | **int** | It represents the maximum number of code controls per authentication. It must be between 1 and 6 inclusive. The default value is 3. 
**SmtpSettingId** | **string** | This is the ID of the SMTP settings used by this configuration. It is mandatory to provide this parameter in order to send the authentication code via email. An SMTPSetting can be created via the web application in an easy way.
**EmailTemplateId** | **string** | This is the ID of the default email template to use for sending authenetication codes via email. If not provided, the message used will be:   ```    ${SERVICE_NAME}: your authentication code is ${CODE}.  ```   
**SmsTemplateId** | **string** | This is the ID of the default sms template to use for sending authenetication codes via sms. If not provided, the message used will be:   ```    ${SERVICE_NAME}: your authentication code is ${CODE}.  ``` 

### Return type

[**ConfigurationItem**](ConfigurationItem.md)

### Authorization

[ApiKey](../README.md#ApiKey), [ApiUser](../README.md#ApiUser)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfiguration

> DeleteConfiguration(ctx, ConfigurationId)

delete a configuration

This operation allows to delete a configuration from the applet account. Once the record is deleted, it will no longer appear in the API and the applet portal.   

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConfigurationId** | **string** | The identifier of the configuration to be deleted.

### Other Parameters

Other parameters are passed through a pointer to a DeleteConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------

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


## FetchConfiguration

> ConfigurationItem FetchConfiguration(ctx, ConfigurationId)

fetch a configuration

This operation allows to fetch the API record associated with a configuration.  This operation needs the `configurationId` of the configuration to be fetched.   

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConfigurationId** | **string** | The identifier of the configuration to be fetched.

### Other Parameters

Other parameters are passed through a pointer to a FetchConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConfigurationItem**](ConfigurationItem.md)

### Authorization

[ApiKey](../README.md#ApiKey), [ApiUser](../README.md#ApiUser)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConfigurations

> []ConfigurationItem ListConfigurations(ctx, optional)

list all configurations

This operation allows to retrieve all configurations defined in your applet.  When getting the configuration record list, results will be sorted based on the `dateCreated` field with the most recent record appearing first. 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListConfigurationsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | Maximum number of records to return per call.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ConfigurationItem**](ConfigurationItem.md)

### Authorization

[ApiKey](../README.md#ApiKey), [ApiUser](../README.md#ApiUser)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration

> ConfigurationItem UpdateConfiguration(ctx, ConfigurationIdoptional)

update a configuration

This operation allows to update the attributes of a configuration. 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConfigurationId** | **string** | The identifier of the configuration to be updated.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**ServiceName** | **string** | The name of the authentication service attached to this configuration. It can be up to 40 characters long.
**CodeLength** | **int** | The length of the code to be generated. It must be a value between 4 and 10, inclusive.
**AllowCustomCode** | **bool** | A flag indicating if the configuration should allow sending custom and non-generated code.
**UsedForDigitalPayment** | **bool** | A flag indicating if the configuration is used to authenticate digital payments. In such a case, additional information such as the amount and the payee of the financial transaction should be sent to when starting the authentication.
**DefaultExpiryTime** | **int** | It represents how long, in minutes, an authentication process will remained in the `awaiting` status before moving to `expired` in the case no valid matching is performed in between. It must be any value between 1 and 1440 which represents 24 hours.
**DefaultMaxTrials** | **int** | It represents the maximum number of trials per authentication. 
**DefaultMaxControls** | **int** | It represents the maximum number of code controls per authentication. It must be between 1 and 6 inclusive. 
**SmtpSettingId** | **string** | This is the ID of the SMTP settings used by this configuration. It is mandatory for sending authentication codes via email.
**EmailTemplateId** | **string** | This is the ID of the default email template to use for sending authenetication codes via email. 
**SmsTemplateId** | **string** | This is the ID of the default sms template to use for sending authenetication codes via sms. 

### Return type

[**ConfigurationItem**](ConfigurationItem.md)

### Authorization

[ApiKey](../README.md#ApiKey), [ApiUser](../README.md#ApiUser)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

