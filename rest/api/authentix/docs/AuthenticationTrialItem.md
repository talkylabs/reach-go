# AuthenticationTrialItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppletId** | **string** | The identifier of the applet. |[optional] 
**ApiVersion** | **string** | The API version. |[optional] 
**ConfigurationId** | **string** | The identifier of the configuration. |[optional] 
**AuthenticationId** | **string** | The identifier of the authentication. |[optional] 
**TrialId** | **string** | The identifier of the authentication trial. |[optional] 
**TrialStatus** | **string** | The status of the authentication. |[optional] 
**Channel** | **string** | The channel used. |[optional] 
**PaymentInfo** | [**PaymentInfo**](PaymentInfo.md) |  |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date and time in GMT that the authentication trial was created.  |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date and time in GMT that the authentication trial was last updated.  |[optional] 
**Price** | **float32** | The cost billed for the authentication trial, in the currency specified by `priceUnit`. This cost does not include the cost for checking the correctness of the code. |[optional] 
**PriceUnit** | **string** | The currency, in ISO 4127 format, in which price is measured. for example, usd, xaf, eur, cad. |[optional] 
**ChannelInfo** | **map[string]interface{}** | channel specific information related to a trial. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


