# StartAuthPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dest** | **string** | The phone number or email where to send the authentication code. Phone numbers must be in E.164 format. |
**Channel** | **string** | The channel by which the authentication code is sent. |
**ServiceName** | **string** | a service name overwriting the one defined in the configuration. |[optional] 
**CustomCode** | **string** | the pre-generated code to be sent. Its length should be between 4 and 10 inclusive. |[optional] 
**PaymentInfo** | **string** | Information related to the digital payment to authenticate. It is required when `usedForDigitalPayment` is true. It is ignored otherwise. It is a stringfied JSON map where keys are `payee`, `amount`, and `currency` and the associated values are respectively the payee, the amount, and the currency of a financial transaction.  |[optional] 
**TemplateId** | **string** | This is the ID of the message template to use for sending the authenetication code. It could be an sms or email template depending on the channel being used. It overwirites the template ID defined in the configuration if any.  |[optional] 
**TemplateDataMap** | **string** | A stringfied JSON map where keys are message template parameters and the values are the parameter values to be used when sending the authentication code. It may also be used to provide additional parameters for sending email based authentications such as the email used for sending the code. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


