# ControlAuthPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dest** | **string** | The phone number or email being authenticated. Phone numbers must be in E.164 format. Either this parameter or the `authenticationId` must be specified. |[optional] 
**Code** | **string** | The 4-10 character string being verified. This is required for `sms` and `email` channels. |[optional] 
**AuthenticationId** | **string** | The ID of the authentication being checked. Either this parameter or the to `dest` must be specified. |[optional] 
**PaymentInfo** | **string** | Information related to the digital payment to authenticate. It is required when `usedForDigitalPayment` is true. It is ignored otherwise. It is a stringfied JSON map where keys are `payee`, `amount`, and `currency` and the associated values are respectively the payee, the amount, and the currency of a financial transaction.  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


