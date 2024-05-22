# prove-sdk-server-go
Go SDK for Prove APIs - Customer Access

<!-- No SDK Installation -->
<!-- No SDK Example Usage -->
<!-- No SDK Available Operations -->
<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code | Content Type     |
| ------------------ | ----------- | ---------------- |
| sdkerrors.Error    | 400,500     | application/json |
| sdkerrors.SDKError | 4xx-5xx     | */*              |

### Example

```go
package main

import (
	"context"
	"errors"
	provesdkservergo "github.com/payfone/prove-sdk-server-go"
	"github.com/payfone/prove-sdk-server-go/models/components"
	"github.com/payfone/prove-sdk-server-go/models/sdkerrors"
	"log"
)

func main() {
	s := provesdkservergo.New(
		provesdkservergo.WithSecurity("<YOUR_AUTH_HERE>"),
	)

	var request *components.V3ChallengeRequest = &components.V3ChallengeRequest{
		CorrelationID: provesdkservergo.String("713189b8-5555-4b08-83ba-75d08780aebd"),
		Dob:           provesdkservergo.String("2024-05-02T00:00:00Z"),
		Last4SSN:      provesdkservergo.String("1234"),
	}

	ctx := context.Background()
	res, err := s.V3.V3ChallengeRequest(ctx, request)
	if err != nil {

		var e *sdkerrors.Error
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| #   | Server                           | Variables |
| --- | -------------------------------- | --------- |
| 0   | `https://api.uat.proveapis.com/` | None      |

#### Example

```go
package main

import (
	"context"
	provesdkservergo "github.com/payfone/prove-sdk-server-go"
	"github.com/payfone/prove-sdk-server-go/models/components"
	"log"
)

func main() {
	s := provesdkservergo.New(
		provesdkservergo.WithServerIndex(0),
		provesdkservergo.WithSecurity("<YOUR_AUTH_HERE>"),
	)

	var request *components.V3ChallengeRequest = &components.V3ChallengeRequest{
		CorrelationID: provesdkservergo.String("713189b8-5555-4b08-83ba-75d08780aebd"),
		Dob:           provesdkservergo.String("2024-05-02T00:00:00Z"),
		Last4SSN:      provesdkservergo.String("1234"),
	}

	ctx := context.Background()
	res, err := s.V3.V3ChallengeRequest(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.V3ChallengeResponse != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	provesdkservergo "github.com/payfone/prove-sdk-server-go"
	"github.com/payfone/prove-sdk-server-go/models/components"
	"log"
)

func main() {
	s := provesdkservergo.New(
		provesdkservergo.WithServerURL("https://api.uat.proveapis.com/"),
		provesdkservergo.WithSecurity("<YOUR_AUTH_HERE>"),
	)

	var request *components.V3ChallengeRequest = &components.V3ChallengeRequest{
		CorrelationID: provesdkservergo.String("713189b8-5555-4b08-83ba-75d08780aebd"),
		Dob:           provesdkservergo.String("2024-05-02T00:00:00Z"),
		Last4SSN:      provesdkservergo.String("1234"),
	}

	ctx := context.Background()
	res, err := s.V3.V3ChallengeRequest(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.V3ChallengeResponse != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name   | Type   | Scheme       |
| ------ | ------ | ------------ |
| `Auth` | oauth2 | OAuth2 token |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	provesdkservergo "github.com/payfone/prove-sdk-server-go"
	"github.com/payfone/prove-sdk-server-go/models/components"
	"log"
)

func main() {
	s := provesdkservergo.New(
		provesdkservergo.WithSecurity("<YOUR_AUTH_HERE>"),
	)

	var request *components.V3ChallengeRequest = &components.V3ChallengeRequest{
		CorrelationID: provesdkservergo.String("713189b8-5555-4b08-83ba-75d08780aebd"),
		Dob:           provesdkservergo.String("2024-05-02T00:00:00Z"),
		Last4SSN:      provesdkservergo.String("1234"),
	}

	ctx := context.Background()
	res, err := s.V3.V3ChallengeRequest(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.V3ChallengeResponse != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->


