# CreateConfigPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceName** | **string** | The name of the authentication service attached to this configuration. It can be up to 40 characters long. |
**CodeLength** | **int** | The length of the code to be generated. It must be a value between 4 and 10, inclusive. If not specified, the default value is 5. |[optional] [default to 5]
**AllowCustomCode** | **bool** | A flag indicating if the configuration should allow sending custom and non-generated code. |[optional] [default to false]
**UsedForDigitalPayment** | **bool** | A flag indicating if the configuration is used to authenticate digital payments. In such a case, additional information such as the amount and the payee of the financial transaction should be sent to when starting the authentication. |[optional] [default to false]
**DefaultExpiryTime** | **int** | It represents how long, in minutes, an authentication process will remained in the `awaiting` status before moving to `expired` in the case no valid matching is performed in between.   It also means that the code sent for the autentication remains the same during its validity period until the autentication is successful. In other words, if another authentication request is asked within that period, the same code will be sent.  If not specified, the default value is 15 minutes. It must be any value between 1 and 1440 which represents 24 hours.  |[optional] [default to 15]
**DefaultMaxTrials** | **int** | It represents the maximum number of trials per authentication. The default value is 5.  |[optional] [default to 5]
**DefaultMaxControls** | **int** | It represents the maximum number of code controls per authentication. It must be between 1 and 6 inclusive. The default value is 3.  |[optional] [default to 3]
**SmtpSettingId** | **string** | This is the ID of the SMTP settings used by this configuration. It is mandatory to provide this parameter in order to send the authentication code via email. An SMTPSetting can be created via the web application in an easy way. |[optional] 
**EmailTemplateId** | **string** | This is the ID of the default email template to use for sending authenetication codes via email. If not provided, the message used will be:   ```    ${SERVICE_NAME}: your authentication code is ${CODE}.  ```    |[optional] 
**SmsTemplateId** | **string** | This is the ID of the default sms template to use for sending authenetication codes via sms. If not provided, the message used will be:   ```    ${SERVICE_NAME}: your authentication code is ${CODE}.  ```  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


