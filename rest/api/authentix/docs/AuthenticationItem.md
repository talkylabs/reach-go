# AuthenticationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppletId** | **string** | The identifier of the applet. |[optional] 
**ApiVersion** | **string** | The API version. |[optional] 
**ConfigurationId** | **string** | The identifier of the configuration. |[optional] 
**AuthenticationId** | **string** | The identifier of the authentication. |[optional] 
**Status** | **string** | The status of the authentication. |[optional] 
**Dest** | **string** | The destination of the authentication code. Phone numbers must be in E.164 format. |[optional] 
**Channel** | **string** | The channel used. |[optional] 
**ExpiryTime** | **int** | An expiry time in minutes.  |[optional] 
**MaxTrials** | **int** | The maximum number of trials.  |[optional] 
**MaxControls** | **int** | The maximum number of code controls.  |[optional] 
**PaymentInfo** | [**PaymentInfo**](PaymentInfo.md) |  |[optional] 
**Trials** | [**[]TrialQuickInfo**](TrialQuickInfo.md) | An array of authentication trials containing channel-specific information about each trial. |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date and time in GMT that the authentication was created.  |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date and time in GMT that the authentication was last updated.  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


