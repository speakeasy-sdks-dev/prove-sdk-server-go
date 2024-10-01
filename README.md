# github.com/prove-identity/prove-sdk-server-go

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
</div>

<!-- Start Summary [summary] -->
## Summary

Prove APIs: This specification describes the Prove API.

OpenAPI Spec - generated.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents

* [SDK Installation](#sdk-installation)
* [SDK Example Usage](#sdk-example-usage)
* [Available Resources and Operations](#available-resources-and-operations)
* [Retries](#retries)
* [Error Handling](#error-handling)
* [Server Selection](#server-selection)
* [Custom HTTP Client](#custom-http-client)
* [Authentication](#authentication)
* [Special Types](#special-types)
<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/prove-identity/prove-sdk-server-go
```
<!-- End SDK Installation [installation] -->

<!-- No SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	provesdkservergo "github.com/prove-identity/prove-sdk-server-go"
	"github.com/prove-identity/prove-sdk-server-go/models/components"
)

func main() {
	// Perform a Prove flow.
	err := flow()
	if err != nil {
		log.Fatal(err)
	}
}

func flow() error {
	// Get environment variables.
	clientID := os.Getenv("PROVE_CLIENT_ID")
	if len(clientID) == 0 {
		return fmt.Errorf("missing env variable: %s", "PROVE_CLIENT_ID")
	}

	clientSecret := os.Getenv("PROVE_CLIENT_SECRET")
	if len(clientSecret) == 0 {
		return fmt.Errorf("missing env variable: %s", "PROVE_CLIENT_SECRET")
	}

	proveEnv := "uat-us" // Use UAT in US region.

	// Create client for Prove API.
	client := provesdkservergo.New(
		provesdkservergo.WithServer(proveEnv),
		provesdkservergo.WithSecurity(components.Security{
			ClientID:     provesdkservergo.String(clientID),
			ClientSecret: provesdkservergo.String(clientSecret),
		}),
	)

	// Send the start request.
	rspStart, err := client.V3.V3StartRequest(context.TODO(), &components.V3StartRequest{
		FlowType:       "desktop",
		FinalTargetURL: provesdkservergo.String("https://example.com"),
	})
	if err != nil {
		return fmt.Errorf("error on Start: %w", err)
	}

	// Get the authToken for the client SDK.
	// authToken := rspStart.V3StartResponse.AuthToken

	//
	// Client SDK work happens here.
	//

	// Validate the phone number.
	rspValidate, err := client.V3.V3ValidateRequest(context.TODO(), &components.V3ValidateRequest{
		CorrelationID: rspStart.V3StartResponse.CorrelationID,
	})
	if err != nil {
		return fmt.Errorf("error on Validate: %w", err)
	}

	// If challenge is the next step, send request.
	if _, ok := rspValidate.V3ValidateResponse.Next["v3-challenge"]; ok {
		rspChallenge, err := client.V3.V3ChallengeRequest(context.TODO(), &components.V3ChallengeRequest{
			CorrelationID: rspStart.V3StartResponse.CorrelationID,
			Dob:           provesdkservergo.String("2024-01-15"),
		})
		if err != nil {
			return fmt.Errorf("error on Challenge: %w", err)
		}

		fmt.Printf("Challenge: %#v\n", rspChallenge.V3ChallengeResponse.Individual)

		// Send individual information to the front end for them to verify.
		// individual := rspChallenge.V3ChallengeResponse.Individual
	}

	// Finish with the complete request.
	rspComplete, err := client.V3.V3CompleteRequest(context.TODO(), &components.V3CompleteRequest{
		CorrelationID: rspStart.V3StartResponse.CorrelationID,
		Individual: components.V3CompleteIndividualRequest{
			FirstName: provesdkservergo.String("Tod"),
			LastName:  provesdkservergo.String("Weedall"),
			Addresses: []components.V3CompleteAddressEntryRequest{
				{
					Address:    provesdkservergo.String("39 South Trail"),
					City:       provesdkservergo.String("San Antonio"),
					Region:     provesdkservergo.String("TX"),
					PostalCode: provesdkservergo.String("78285"),
				},
			},
			Ssn: provesdkservergo.String("565228370"),
			Dob: provesdkservergo.String("1984-12-10"),
			EmailAddresses: []string{
				"tweedalld@ehow.com",
			},
		},
	})
	if err != nil {
		return fmt.Errorf("error on Complete: %w", err)
	}

	if !rspComplete.V3CompleteResponse.Success {
		return fmt.Errorf("user could not be validated")
	}

	return nil
}

```

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>


### [V3](docs/sdks/v3/README.md)

* [V3TokenRequest](docs/sdks/v3/README.md#v3tokenrequest) - Request OAuth token.
* [V3ChallengeRequest](docs/sdks/v3/README.md#v3challengerequest) - Submit challenge.
* [V3CompleteRequest](docs/sdks/v3/README.md#v3completerequest) - Complete flow.
* [V3StartRequest](docs/sdks/v3/README.md#v3startrequest) - Start flow.
* [V3ValidateRequest](docs/sdks/v3/README.md#v3validaterequest) - Validate phone number.

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `V3TokenRequest` function may return the following errors:

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

### Example

```go
package main

import (
	"context"
	"errors"
	provesdkservergo "github.com/prove-identity/prove-sdk-server-go"
	"github.com/prove-identity/prove-sdk-server-go/models/components"
	"github.com/prove-identity/prove-sdk-server-go/models/sdkerrors"
	"log"
)

func main() {
	s := provesdkservergo.New()

	ctx := context.Background()
	res, err := s.V3.V3TokenRequest(ctx, &components.V3TokenRequest{
		ClientID:     "customer_id",
		ClientSecret: "secret",
		GrantType:    "client_credentials",
	})
	if err != nil {

		var e *sdkerrors.Error400
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

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

### Select Server by Name

You can override the default server globally using the `WithServer` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the names associated with the available servers:

| Name | Server | Variables |
| ----- | ------ | --------- |
| `uat-us` | `https://platform.uat.proveapis.com` | None |
| `prod-us` | `https://platform.proveapis.com` | None |

#### Example

```go
package main

import (
	"context"
	provesdkservergo "github.com/prove-identity/prove-sdk-server-go"
	"github.com/prove-identity/prove-sdk-server-go/models/components"
	"log"
)

func main() {
	s := provesdkservergo.New(
		provesdkservergo.WithServer("prod-us"),
	)

	ctx := context.Background()
	res, err := s.V3.V3TokenRequest(ctx, &components.V3TokenRequest{
		ClientID:     "customer_id",
		ClientSecret: "secret",
		GrantType:    "client_credentials",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.V3TokenResponse != nil {
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
	provesdkservergo "github.com/prove-identity/prove-sdk-server-go"
	"github.com/prove-identity/prove-sdk-server-go/models/components"
	"log"
)

func main() {
	s := provesdkservergo.New(
		provesdkservergo.WithServerURL("https://platform.uat.proveapis.com"),
	)

	ctx := context.Background()
	res, err := s.V3.V3TokenRequest(ctx, &components.V3TokenRequest{
		ClientID:     "customer_id",
		ClientSecret: "secret",
		GrantType:    "client_credentials",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.V3TokenResponse != nil {
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

| Name                           | Type                           | Scheme                         |
| ------------------------------ | ------------------------------ | ------------------------------ |
| `ClientID` `ClientSecret`      | oauth2                         | OAuth2 Client Credentials Flow |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	provesdkservergo "github.com/prove-identity/prove-sdk-server-go"
	"github.com/prove-identity/prove-sdk-server-go/models/components"
	"log"
)

func main() {
	s := provesdkservergo.New(
		provesdkservergo.WithSecurity(components.Security{
			ClientID:     provesdkservergo.String("<YOUR_CLIENT_ID_HERE>"),
			ClientSecret: provesdkservergo.String("<YOUR_CLIENT_SECRET_HERE>"),
		}),
	)

	ctx := context.Background()
	res, err := s.V3.V3TokenRequest(ctx, &components.V3TokenRequest{
		ClientID:     "customer_id",
		ClientSecret: "secret",
		GrantType:    "client_credentials",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.V3TokenResponse != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- No Special Types [types] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	provesdkservergo "github.com/prove-identity/prove-sdk-server-go"
	"github.com/prove-identity/prove-sdk-server-go/models/components"
	"github.com/prove-identity/prove-sdk-server-go/retry"
	"log"
	"models/operations"
)

func main() {
	s := provesdkservergo.New()

	ctx := context.Background()
	res, err := s.V3.V3TokenRequest(ctx, &components.V3TokenRequest{
		ClientID:     "customer_id",
		ClientSecret: "secret",
		GrantType:    "client_credentials",
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.V3TokenResponse != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	provesdkservergo "github.com/prove-identity/prove-sdk-server-go"
	"github.com/prove-identity/prove-sdk-server-go/models/components"
	"github.com/prove-identity/prove-sdk-server-go/retry"
	"log"
)

func main() {
	s := provesdkservergo.New(
		provesdkservergo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
	)

	ctx := context.Background()
	res, err := s.V3.V3TokenRequest(ctx, &components.V3TokenRequest{
		ClientID:     "customer_id",
		ClientSecret: "secret",
		GrantType:    "client_credentials",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.V3TokenResponse != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

