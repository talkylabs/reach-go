# Custom HTTP Clients for the Reach Go Helper Library

If you are working with the [Reach Go Helper Library](../README.md), and you need to be able to modify the HTTP requests that the library makes to the Reach servers, you’re in the right place. The most common reason for altering the HTTP request is to connect and authenticate with an enterprise’s proxy server. We’ll provide sample code that you can drop into your app to handle this use case.

## Connect and authenticate with a proxy server

To connect and provide credentials to a proxy server that may be between your app and Reach, you need a way to modify the HTTP requests that the Reach helper library makes on your behalf when invoking Reach's REST API.

In Go, the Reach helper library uses the native [net/http package](https://pkg.go.dev/net/http) under the hood to make the HTTP requests. The Reach Helper Library allows you to provide your own `Client` for making API requests.

The following example shows a typical request without a custom `Client`.

```go
reachClient := reach.NewRestClient()

params := &reachMessaging.SendMessageParams{}
params.SetDest("+15558675309")
params.SetSrc("+15017250604")
params.SetBody("Hey there!")

resp, err := reachClient.Messaging.SendMessage(params)
```

Out of the box, the helper library creates a default `Client` for you, using the Reach credentials from your environment variables or that you pass in directly. However, there’s nothing stopping you from creating your own client and using that.

Once you have your own `Client`, you can pass it to any Reach REST API resource action you want.

## Create and use your custom Client

When you take a closer look at the input parameters for `reach.RestClient`, you see that the `Client` parameter is actually of type `client.BaseClient`.

`client.BaseClient` is an abstraction that allows plugging in any implementation of an HTTP client you want (or even creating a mocking layer for unit testing).

Now that you understand how all the components fit together, you can create your own `Client` that can connect through a proxy server. To make this reusable, here’s a class that you can use to create this `HttpClient` whenever you need one.

Here’s an example of sending an SMS message with a custom client:

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/talkylabs/reach-go"
	"github.com/talkylabs/reach-go/client"
	reachMessaging "github.com/talkylabs/reach-go/rest/api/messaging"
)

func main() {
	apiUser := os.Getenv("REACH_TALKYLABS_API_USER")
	apiKey := os.Getenv("REACH_TALKYLABS_API_KEY")

	// Add proxy settings to a http Transport object
	transport := &http.Transport{
		// https://pkg.go.dev/net/http#ProxyFromEnvironment
		Proxy: http.ProxyFromEnvironment,
	}

	// Add the Transport to an http Client
	httpClient := &http.Client{
		Transport: transport,
	}

	// Create your custom Reach client using the http client and your credentials
	reachHttpClient := client.Client{
		Credentials: client.NewCredentials(apiUser, apiKey),
		HTTPClient:  httpClient,
	}
	reachClient := reach.NewRestClientWithParams(reach.ClientParams{Client: &reachHttpClient})

	params := &reachMessaging.SendMessageParams{}
	params.SetDest("+15558675310")
	params.SetSrc("+15017122661")
	params.SetBody("Hey there!")

	resp, err := reachClient.Messaging.SendMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}
```

In this example, you use environment variables loaded at the program startup to retrieve various configuration settings:

- Your Reach Api user and key found in the Reach web application
- A proxy address in the form of `http://127.0.0.1:8888`

These settings are either exported manually by yourself in the terminal, or located in a file such as `.env`, like so:

```text
REACH_TALKYLABS_API_USER=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
REACH_TALKYLABS_API_KEY=your_auth_token

HTTPS_PROXY=https://127.0.0.1:8888
HTTP_PROXY=http://127.0.0.1:8888
```

## What else can this technique be used for?

Now that you know how to inject your own `Client` into the Reach API request pipeline, you could use this technique to add custom HTTP headers and authorization to the requests (perhaps as required by an upstream proxy server). You could do so by overriding the `SendRequest` method, and adding any desired pre-processing to your requests and responses, like so:

```go
package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/talkylabs/reach-go"
	"github.com/talkylabs/reach-go/client"
	openapi "github.com/talkylabs/reach-go/rest/api/messaging"
)

type MyClient struct {
	client.Client
}

func (c *MyClient) SendRequest(method string, rawURL string, data url.Values, headers map[string]interface{}) (*http.Response, error) {
	// Custom code to pre-process request here
	resp, err := c.Client.SendRequest(method, rawURL, data, headers)
	// Custom code to pre-process response here
	fmt.Println(resp.StatusCode)
	return resp, err
}

func main() {
	apiUser := os.Getenv("REACH_TALKYLABS_API_USER")
	apiKey := os.Getenv("REACH_TALKYLABS_API_KEY")

	customClient := &MyClient{
		Client: client.Client{
			Credentials: client.NewCredentials(apiUser, apiKey),
		},
	}

	reachClient := reach.NewRestClientWithParams(reach.ClientParams{Client: customClient})

	// You may also use custom clients with standalone product services
	reachMessaging := openapi.NewApiServiceWithClient(customClient)
}
```

You could also implement your own `Client` to mock the Reach API responses so your unit and integration tests can run quickly without needing to make a connection to Reach.

We can’t wait to see what you build!
