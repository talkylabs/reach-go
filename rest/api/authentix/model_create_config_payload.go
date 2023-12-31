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
)
// CreateConfigPayload struct for CreateConfigPayload
type CreateConfigPayload struct {
		// The name of the authentication service attached to this configuration. It can be up to 40 characters long.
	ServiceName string `json:"serviceName"`
		// The length of the code to be generated. It must be a value between 4 and 10, inclusive. If not specified, the default value is 5.
	CodeLength int `json:"codeLength,omitempty"`
		// A flag indicating if the configuration should allow sending custom and non-generated code.
	AllowCustomCode bool `json:"allowCustomCode,omitempty"`
		// A flag indicating if the configuration is used to authenticate digital payments. In such a case, additional information such as the amount and the payee of the financial transaction should be sent to when starting the authentication.
	UsedForDigitalPayment bool `json:"usedForDigitalPayment,omitempty"`
		// It represents how long, in minutes, an authentication process will remained in the `awaiting` status before moving to `expired` in the case no valid matching is performed in between.   It also means that the code sent for the autentication remains the same during its validity period until the autentication is successful. In other words, if another authentication request is asked within that period, the same code will be sent.  If not specified, the default value is 15 minutes. It must be any value between 1 and 1440 which represents 24 hours. 
	DefaultExpiryTime int `json:"defaultExpiryTime,omitempty"`
		// It represents the maximum number of trials per authentication. The default value is 5. 
	DefaultMaxTrials int `json:"defaultMaxTrials,omitempty"`
		// It represents the maximum number of code controls per authentication. It must be between 1 and 6 inclusive. The default value is 3. 
	DefaultMaxControls int `json:"defaultMaxControls,omitempty"`
		// This is the ID of the SMTP settings used by this configuration. It is mandatory to provide this parameter in order to send the authentication code via email. An SMTPSetting can be created via the web application in an easy way.
	SmtpSettingId string `json:"smtpSettingId,omitempty"`
		// This is the ID of the default email template to use for sending authenetication codes via email. If not provided, the message used will be:   ```    ${SERVICE_NAME}: your authentication code is ${CODE}.  ```   
	EmailTemplateId string `json:"emailTemplateId,omitempty"`
		// This is the ID of the default sms template to use for sending authenetication codes via sms. If not provided, the message used will be:   ```    ${SERVICE_NAME}: your authentication code is ${CODE}.  ``` 
	SmsTemplateId string `json:"smsTemplateId,omitempty"`
}


