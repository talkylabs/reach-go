# AuthenticationControlItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppletId** | **string** | The identifier of the applet. |[optional] 
**ApiVersion** | **string** | The API version. |[optional] 
**ConfigurationId** | **string** | The identifier of the configuration. |[optional] 
**AuthenticationId** | **string** | The identifier of the authentication. |[optional] 
**Status** | **string** | The outcome of the authentication control. |[optional] 
**Dest** | **string** | The phone number or email being verified. Phone numbers must be in E.164 format. |[optional] 
**Channel** | **string** | The channel used. |[optional] 
**PaymentInfo** | [**PaymentInfo**](PaymentInfo.md) |  |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date and time in GMT that the authentication was created.  |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date and time in GMT that the authentication was last updated.  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


