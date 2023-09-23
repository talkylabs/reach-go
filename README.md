# reach-go

## Documentation

The documentation for the Reach API can be found [here][apidocs].

The Go library documentation can be found [here][libdocs].

### Supported Go Versions

This library supports the following Go implementations:

- Go 1.15
- Go 1.16
- Go 1.17
- Go 1.18
- Go 1.19
- Go 1.20

## Installation

The recommended way to install `reach-go` is by using [Go modules](https://go.dev/ref/mod#go-get).

If you already have an initialized project, you can run the command below from your terminal in the project directory to install the library:

```shell
go get github.com/talkylabs/reach-go
```

If you are starting from scratch in a new directory, you will first need to create a go.mod file for tracking dependencies such as reach-go. This is similar to using package.json in a Node.js project or requirements.txt in a Python project. [You can read more about mod files in the Go documentation](https://golang.org/doc/modules/managing-dependencies). To create the file, run the following command in your terminal:

```shell
go mod init reach-example
```

Once the module is initialized, you may run the installation command from above, which will update your go.mod file to include reach-go.

### Test your installation

Try sending yourself an SMS message by pasting the following code example into a sendsms.go file in the same directory where you installed reach-go. Be sure to update the apiUser, apiKey, and from phone number with values from your Reach account. The to phone number can be your own mobile phone number.

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/talkylabs/reach-go"
	reachMessaging "github.com/talkylabs/reach-go/rest/api/messaging"
)

func main() {
	apiUser := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	apiKey := "f2xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

	client := reach.NewRestClientWithParams(reach.ClientParams{
		ApiUser: apiUser,
		ApiKey: apiKey,
	})

	params := &reachMessaging.SendMessageParams{}
	params.SetDest("+15558675309")
	params.SetSrc("+15017250604")
	params.SetBody("Hello from Go!")

	resp, err := client.Messaging.SendMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}
```

Save `sendsms.go`. In your terminal from the same directory, run:

```shell
go run sendsms.go
```

After a brief delay, you will receive the text message on your phone.

> **Warning**
> It's okay to hardcode your credentials when testing locally, but you should use environment variables to keep them secret before committing any code or deploying to production.

## Use the helper library

### API credentials

The Reach `RestClient` needs your Reach credentials. We recommend storing them as environment variables, so that you don't have to worry about committing and accidentally posting them somewhere public.

```go
package main

import "github.com/talkylabs/reach-go"

func main() {
	// This will look for `REACH_TALKYLABS_API_USER` and `REACH_TALKYLABS_API_KEY` variables inside the current environment to initialize the constructor
	// You can find your Account SID and Auth Token at the web console
	client := reach.NewRestClient()
}
```

If you don't want to use environment variables, you can also pass the credentials directly to the constructor as below.

```go
package main

import "github.com/talkylabs/reach-go"

func main() {
	apiUser := "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	apiKey := "YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY"
	client := reach.NewRestClientWithParams(reach.ClientParams{
		ApiUser: apiUser,
		ApiKey: apiKey,
	})
}
```


### Get data about an existing Messaging Item

```go
package main

import (
	"fmt"
	"os"

	"github.com/talkylabs/reach-go"
	reachMessaging "github.com/talkylabs/reach-go/rest/api/messaging"
)

func main() {
	apiUser := os.Getenv("REACH_TALKYLABS_API_USER")
	apiKey := os.Getenv("REACH_TALKYLABS_API_KEY")

	client := reach.NewRestClientWithParams(reach.ClientParams{
		ApiUser:   apiUser,
		ApiKey:   apiKey,
	})

	params := &reachMessaging.FetchMessageParams{}
	params.SetMessageId("CA42ed11f93dc08b952027ffbc406d0868")

	resp, err := client.Messaging.FetchMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Item Status: " + *resp.Status)
		fmt.Println("Item Sid: " + *resp.MessageId)
		fmt.Println("Item Destination: " + *resp.Dest)
	}
}
```

### Iterate through records

This library also offers paging functionality. Collections such as messages have `ListXxx` and `StreamXxx`
functions that page under the hood. With both list and stream, you can specify the number of records you want to
receive (limit) and the maximum size you want each page fetch to be (pageSize). The library will then handle the task
for you.

`List` eagerly fetches all records and returns them as a list, whereas `Stream` streams the records and lazily retrieves
the pages as you iterate over the collection. Also, `List` returns no records if any errors are encountered while paging,
whereas `Stream` returns all records up until encountering an error. You can also page manually using the `PageXxx`
function in each of the apis.

#### Use `ListXxx` or `StreamXxx`

```go
package main

import (
	"fmt"
	"github.com/talkylabs/reach-go"
	reachMessaging "github.com/talkylabs/reach-go/rest/api/messaging"
	"os"
)

func main() {
	from := os.Getenv("REACH_TALKYLABS_FROM_PHONE_NUMBER")

	client := reach.NewRestClient()

	params := &reachMessaging.ListMessageParams{}
	params.SetSrc(from)
	params.SetPageSize(20)
	params.SetLimit(100)

	resp, _ := client.Messaging.ListMessage(params)
	for record := range resp {
		fmt.Println("Body: ", *resp[record].Body)
	}

	channel, _ := client.Messaging.StreamMessage(params)
	for record := range channel {
		fmt.Println("Body: ", *record.Body)
	}
}
```

#### Use `PageXxx`

```go
package main

import (
	"fmt"
	"github.com/talkylabs/reach-go"
	reachMessaging "github.com/talkylabs/reach-go/rest/api/messaging"
	"net/url"
	"os"
)

func main() {
	from := os.Getenv("REACH_TALKYLABS_FROM_PHONE_NUMBER")

	client := reach.NewRestClient()

	params := &reachMessaging.ListMessageParams{}
	params.SetSrc(from)
	params.SetPageSize(20)

	var pageNumber string
	var outOfPageRangeFlag bool
	resp, err = client.Messaging.PageMessage(params, "")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp.Page)
		fmt.Println(resp.OutOfPageRange)
		outOfPageRangeFlag = resp.OutOfPageRange
		pageNumber = strconv.Itoa(resp.Page + 1)
	}

	if !outOfPageRangeFlag{
		resp, err := client.Messaging.PageMessage(params, pageNumber)
		if err != nil {
			fmt.Println(err)
		} else {
			if resp != nil {
				fmt.Println(*resp.Messages[0].Body)
			}
		}
	}
	
}
```

### Handle Exceptions

If the Reach API returns a 400 or a 500 level HTTP response, the reach-go library will include information in the returned err value. 400-level errors are normal during API operation ("Invalid number", "Cannot deliver SMS to that number", for example) and should be handled appropriately.

```go
package main

import (
	"fmt"
	"os"

	"github.com/talkylabs/reach-go"
	reachclient "github.com/talkylabs/reach-go/client"
	reachMessaging "github.com/talkylabs/reach-go/rest/api/messaging"
)

func main() {
	phoneNumber := os.Getenv("REACH_TALKYLABS_PHONE_NUMBER")

	client := reach.NewRestClient()

	params := &reachMessaging.SendMessageParams{}
	params.SetDest("+15558675309")
	params.SetSrc("+15017250604")
	params.SetBody("Hello from Go!")

	resp, err := client.Messaging.SendMessage(params)
	if err != nil {
		reachError := err.(*reachclient.ReachRestError)
		fmt.Println(reachError.Error())
	}
}
```

## Advanced Usage


### Use standalone products

Don't want to import the top-level Reach RestClient with access to the full suite of Reach products? Use standalone product services instead:

```go
package main

import (
	"github.com/talkylabs/reach-go/client"
	reachMessaging "github.com/talkylabs/reach-go/rest/api/messaging"
	"os"
)

func main() {
	apiUser := os.Getenv("REACH_TALKYLABS_API_USER")
	apiKey := os.Getenv("REACH_TALKYLABS_API_KEY")

	// Create an instance of our default BaseClient implementation
	// You will need to provide your API credentials to the Client manually
	defaultClient := &client.Client{
		Credentials: client.NewCredentials(apiUser, apiKey),
	}

	messagingService := reachMessaging.NewApiServiceWithClient(defaultClient)
}
```

### Other advanced examples

- [Learn how to create your own custom HTTP client](./advanced-examples/custom-http-client.md)



## Local Usage

### Building

To build _reach-go_ run:

```shell
go build ./...
```

### Testing

To execute the test suite run:

```shell
go test ./...
```

### Generating Local Documentation

To generate documentation, from the root directory:

```shell
godoc -http=localhost:{port number}
```

Then, navigate to `http://localhost:{port number}/pkg/github.com/talkylabs/reach-go` in your local browser.

Example:

```shell
godoc -http=localhost:6060
```

http://localhost:6060/pkg/github.com/talkylabs/reach-go

## Docker Image

The `Dockerfile` present in this repository and its respective `talkylabs/reach-go` Docker image are currently used by TalkyLabs for testing purposes only.

## Getting help

If you've found a bug in the library or would like new features added, go ahead and open issues or pull requests against this repo!

[apidocs]: https://www.reach.talkylabs.com/docs/api
[libdocs]: https://pkg.go.dev/github.com/talkylabs/reach-go?tab=versions
