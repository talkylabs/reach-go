/*
 * This code was generated by
 *  ___ ___   _   ___ _  _    _____ _   _    _  ___   ___      _   ___ ___      ___   _   ___     ___ ___ _  _ ___ ___    _ _____ ___  ___ 
 * | _ \ __| /_\ / __| || |__|_   _/_\ | |  | |/ | \ / / |    /_\ | _ ) __|___ / _ \ /_\ |_ _|__ / __| __| \| | __| _ \  /_\_   _/ _ \| _ \
 * |   / _| / _ \ (__| __ |___|| |/ _ \| |__| ' < \ V /| |__ / _ \| _ \__ \___| (_) / _ \ | |___| (_ | _|| .` | _||   / / _ \| || (_) |   /
 * |_|_\___/_/ \_\___|_||_|    |_/_/ \_\____|_|\_\ |_| |____/_/ \_\___/___/    \___/_/ \_\___|   \___|___|_|\_|___|_|_\/_/ \_\_| \___/|_|_\
 * 
 * Reach Authentix API
 *  Reach Authentix API helps you easily integrate user authentification in your application. The authentification allows to verify that a user is indeed at the origin of a request from your application.  At the moment, the Reach Authentix API supports the following channels:    * SMS      * Email   We are continuously working to add additionnal channels. ## Base URL All endpoints described in this documentation are relative to the following base URL: ``` https://api.reach.talkylabs.com/rest/authentix/v1/ ```  The API is provided over HTTPS protocol to ensure data privacy.  ## API Authentication Requests made to the API must be authenticated. You need to provide the `ApiUser` and `ApiKey` associated with your applet. This information could be found in the settings of the applet. ```curl curl -X GET [BASE_URL]/configurations -H \"ApiUser:[Your_Api_User]\" -H \"ApiKey:[Your_Api_Key]\" ``` ## Reach Authentix API Workflow Three steps are needed in order to authenticate a given user using the Reach Authentix API. ### Step 1: Create an Authentix configuration A configuration is a set of settings used to define and send an authentication code to a user. This includes, for example: ```   - the length of the authentication code,    - the message template,    - and so on... ``` A configuaration could be created via the web application or directly using the Reach Authentix API. This step does not need to be performed every time one wants to use the Reach Authentix API. Indeed, once created, a configuartion could be used to authenticate several users in the future.    ### Step 2: Send an authentication code A configuration is used to send an authentication code via a selected channel to a user. For now, the supported channels are `sms`, and `email`. We are working hard to support additional channels. Newly created authentications will have a status of `awaiting`. ### Step 3: Verify the authentication code This step allows to verify that the code submitted by the user matched the one sent previously. If, there is a match, then the status of the authentication changes from `awaiting` to `passed`. Otherwise, the status remains `awaiting` until either it is verified or it expires. In the latter case, the status becomes `expired`. 
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi
import (
	"time"
)
// AuthenticationItem struct for AuthenticationItem
type AuthenticationItem struct {
		// The identifier of the applet.
	AppletId string `json:"appletId,omitempty"`
		// The API version.
	ApiVersion string `json:"apiVersion,omitempty"`
		// The identifier of the configuration.
	ConfigurationId string `json:"configurationId,omitempty"`
		// The identifier of the authentication.
	AuthenticationId string `json:"authenticationId,omitempty"`
		// The status of the authentication.
	Status string `json:"status,omitempty"`
		// The destination of the authentication code. Phone numbers must be in E.164 format.
	Dest string `json:"dest,omitempty"`
		// The channel used.
	Channel string `json:"channel,omitempty"`
		// An expiry time in minutes. 
	ExpiryTime int `json:"expiryTime,omitempty"`
		// The maximum number of trials. 
	MaxTrials int `json:"maxTrials,omitempty"`
		// The maximum number of code controls. 
	MaxControls int `json:"maxControls,omitempty"`
	PaymentInfo PaymentInfo `json:"paymentInfo,omitempty"`
		// An array of authentication trials containing channel-specific information about each trial.
	Trials []TrialQuickInfo `json:"trials,omitempty"`
		// The date and time in GMT that the authentication was created. 
	DateCreated time.Time `json:"dateCreated,omitempty"`
		// The date and time in GMT that the authentication was last updated. 
	DateUpdated time.Time `json:"dateUpdated,omitempty"`
}


