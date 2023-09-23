package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"

	"github.com/talkylabs/reach-go"
	reachMessaging "github.com/talkylabs/reach-go/rest/api/messaging"
	reachAuthentix "github.com/talkylabs/reach-go/rest/api/authentix"
)

func main() {
	apiUser := os.Getenv("REACH_TALKYLABS_API_USER")
	apiKey := os.Getenv("REACH_TALKYLABS_API_KEY")

	client := reach.NewRestClientWithParams(reach.ClientParams{
		ApiUser: apiUser,
		ApiKey:  apiKey,
	})

	
	params1 := &reachMessaging.SendMessageParams{}
	params1.SetDest("+23767xxxxxx")
	params1.SetSrc("+23769xxxxxx")
	params1.SetBody("Hello World! This is a sms message.")

	sms, err := client.Messaging.SendMessage(params1)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	} else {
		fmt.Println(sms.MessageId)
	}
	

	params2 := &reachMessaging.ListMessageParams{}
	params2.SetLimit(25)

	resp, err := client.Messaging.ListMessage(params2)
	num := 0
	if err != nil {
		fmt.Println("Error: " + err.Error())
	} else {
		for message := range resp {
		        num = num + 1
			fmt.Println(resp[message].Dest)
		}
	}
	fmt.Println(num)
	
	
	
	/*
	
	params := &reachAuthentix.CreateConfigurationParams{}
        params.SetServiceName("AuthServiceGo2")
        params.SetCodeLength(6)
        params.SetAllowCustomCode(true)
        confId := ""

        config, err := client.Authentix.CreateConfiguration(params)
        if err != nil {
            fmt.Println("Error: " + err.Error())
        } else {
          confId = config.ConfigurationId
          fmt.Println(config.ConfigurationId)
        }
        
        paramsL := &reachAuthentix.ListConfigurationsParams{}
        paramsL.SetPageSize(20)

        resp, err := client.Authentix.ListConfigurations(paramsL)
        if err != nil {
            fmt.Println("Error: " + err.Error())
        } else {
          for configuration := range resp {
            fmt.Println(resp[configuration].ConfigurationId)
          }
        }
        
        config2, err := client.Authentix.FetchConfiguration(confId)
	if err != nil {
	    fmt.Println("Error: " + err.Error())
	} else {
	  fmt.Println(config2.ServiceName)
	}
	
	err = client.Authentix.DeleteConfiguration(confId)
        if err != nil {
            fmt.Println("Error: " + err.Error())
        }else{
            fmt.Println("true")
        }
        */
        
        paramsStart := &reachAuthentix.StartAuthenticationParams{}
        paramsStart.SetDest("custom@domain.com")
        paramsStart.SetChannel("email")
    

        auth1, err := client.Authentix.StartAuthentication("CIDXXXXXXXXXXXX", paramsStart)
        if err != nil {
            fmt.Println("Error: " + err.Error())
        } else {
            fmt.Println(auth1.AuthenticationId)
        }
        
        auth2, err := client.Authentix.FetchAuthentication("CIDXXXXXXXXXXXX", auth1.AuthenticationId)
        if err != nil {
            fmt.Println("Error: " + err.Error())
        } else {
            fmt.Println(auth2.AuthenticationId)
        }
        
        reader := bufio.NewReader(os.Stdin)
        
        fmt.Print("Enter received Code: ")
    	codeStr, _ := reader.ReadString('\n')
    	codeStr = strings.Replace(codeStr, "\n", "", -1)
    	
    	paramsControl := &reachAuthentix.CheckAuthenticationCodeParams{}
        paramsControl.SetCode(codeStr)
        paramsControl.SetDest("custom@domain.com")

        authControl, err := client.Authentix.CheckAuthenticationCode("CIDXXXXXXXXXXXX", paramsControl)
        if err != nil {
            fmt.Println("Error: " + err.Error())
        } else {
            fmt.Println(authControl.Status)
        }
        
        paramsTrialStats := &reachAuthentix.FetchAuthenticationTrialStatsParams{}
        
        authStats, err := client.Authentix.FetchAuthenticationTrialStats(paramsTrialStats)
        if err != nil {
            fmt.Println("Error: " + err.Error())
        } else {
            fmt.Println(authStats.SuccessRate)
        }
        
        paramsTrialList := &reachAuthentix.ListAuthenticationTrialsParams{}
        paramsTrialList.SetPageSize(20)

        trialList, err := client.Authentix.ListAuthenticationTrials(paramsTrialList)
        if err != nil {
            fmt.Println("Error: " + err.Error())
        } else {
            for authTrial := range trialList {
                fmt.Println(trialList[authTrial].TrialId)
            }
        }
        
        fmt.Print("Enter Trial Id: ")
    	trialIdStr, _ := reader.ReadString('\n')
    	trialIdStr = strings.Replace(trialIdStr, "\n", "", -1)
    	
    	trial, err := client.Authentix.FetchAthenticationTrial(trialIdStr)
        if err != nil {
            fmt.Println("Error: " + err.Error())
        } else {
            fmt.Println(trial.TrialId)
            fmt.Println(trial.TrialStatus)
        }
}
