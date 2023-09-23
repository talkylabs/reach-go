# ConfigurationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppletId** | **string** | The identifier of the applet creating the configuration. |[optional] 
**ApiVersion** | **string** | The API version used to create the configuration. |[optional] 
**ConfigurationId** | **string** | The identifier of the configuration. |[optional] 
**ServiceName** | **string** | The name of the authentication service. |[optional] 
**CodeLength** | **int** | The length of the code to be generated. |[optional] 
**AllowCustomCode** | **bool** | A flag indicating if the configuration allows sending custom and non-generated code. |[optional] 
**UsedForDigitalPayment** | **bool** | A flag indicating if the configuration is used to authenticate digital payments. |[optional] 
**DefaultExpiryTime** | **int** | the default expiry time of the authentication code. |[optional] 
**DefaultMaxTrials** | **int** | the default maximum number of trials per authentication. |[optional] 
**DefaultMaxControls** | **int** | the default maximum number of code controls per authentication. |[optional] 
**SmtpSettingId** | **string** | The ID of the SMTP settings used by the configuration. |[optional] 
**EmailTemplateId** | **string** | The default email template ID used by this configuration.  |[optional] 
**SmsTemplateId** | **string** | The default sms template ID used by this configuration.  |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date and time in GMT that the configuration was created.  |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date and time in GMT that the configuration was last updated.  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


