# V3
(*V3*)

### Available Operations

* [V3TokenRequest](#v3tokenrequest) - Request OAuth token.
* [V3ChallengeRequest](#v3challengerequest) - Submit challenge.
* [V3CompleteRequest](#v3completerequest) - Complete flow.
* [V3StartRequest](#v3startrequest) - Start flow.
* [V3ValidateRequest](#v3validaterequest) - Validate phone number.

## V3TokenRequest

Send this request to request the OAuth token.

### Example Usage

```go
package main

import(
	provesdkservergo "github.com/prove-identity/prove-sdk-server-go"
	"github.com/prove-identity/prove-sdk-server-go/models/components"
	"context"
	"log"
)

func main() {
    s := provesdkservergo.New()
    var request *components.V3TokenRequest = &components.V3TokenRequest{
        ClientID: "customer_id",
        ClientSecret: "secret",
        GrantType: "client_credentials",
    }
    ctx := context.Background()
    res, err := s.V3.V3TokenRequest(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.V3TokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [components.V3TokenRequest](../../models/components/v3tokenrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.V3TokenRequestResponse](../../models/operations/v3tokenrequestresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 400,500            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## V3ChallengeRequest

Send this request to submit challenge information. Either a DOB or last 4 of SSN needs to be submitted if neither was submitted to the /start endpoint. It will return a correlation ID, user information, and the next step to call in the flow.

### Example Usage

```go
package main

import(
	"github.com/prove-identity/prove-sdk-server-go/models/components"
	provesdkservergo "github.com/prove-identity/prove-sdk-server-go"
	"context"
	"log"
)

func main() {
    s := provesdkservergo.New(
        provesdkservergo.WithSecurity(components.Security{
            ClientID: "<YOUR_CLIENT_ID_HERE>",
            ClientSecret: "<YOUR_CLIENT_SECRET_HERE>",
        }),
    )
    var request *components.V3ChallengeRequest = &components.V3ChallengeRequest{
        CorrelationID: "713189b8-5555-4b08-83ba-75d08780aebd",
        Dob: provesdkservergo.String("2024-05-02T00:00:00Z"),
        Last4SSN: provesdkservergo.String("1234"),
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

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.V3ChallengeRequest](../../models/components/v3challengerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.V3ChallengeRequestResponse](../../models/operations/v3challengerequestresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 400,500            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## V3CompleteRequest

Send this request to verify the user and complete the flow. It will return a correlation ID, user information, and the next step to call in the flow.

### Example Usage

```go
package main

import(
	"github.com/prove-identity/prove-sdk-server-go/models/components"
	provesdkservergo "github.com/prove-identity/prove-sdk-server-go"
	"context"
	"log"
)

func main() {
    s := provesdkservergo.New(
        provesdkservergo.WithSecurity(components.Security{
            ClientID: "<YOUR_CLIENT_ID_HERE>",
            ClientSecret: "<YOUR_CLIENT_SECRET_HERE>",
        }),
    )
    var request *components.V3CompleteRequest = &components.V3CompleteRequest{
        CorrelationID: "713189b8-5555-4b08-83ba-75d08780aebd",
        Individual: components.V3CompleteIndividualRequest{
            Addresses: []components.V3CompleteAddressEntryRequest{
                components.V3CompleteAddressEntryRequest{
                    Address: provesdkservergo.String("39 South Trail"),
                    City: provesdkservergo.String("San Antonio"),
                    ExtendedAddress: provesdkservergo.String("Apt 23"),
                    PostalCode: provesdkservergo.String("78285"),
                    Region: provesdkservergo.String("TX"),
                },
                components.V3CompleteAddressEntryRequest{
                    Address: provesdkservergo.String("4861 Jay Junction"),
                    City: provesdkservergo.String("Boston"),
                    ExtendedAddress: provesdkservergo.String("Apt 78"),
                    PostalCode: provesdkservergo.String("02208"),
                    Region: provesdkservergo.String("MS"),
                },
            },
            Dob: provesdkservergo.String("2024-05-02T00:00:00Z"),
            EmailAddresses: []string{
                "jdoe@example.com",
                "dsmith@example.com",
            },
            FirstName: provesdkservergo.String("Tod"),
            Last4SSN: provesdkservergo.String("1234"),
            LastName: provesdkservergo.String("Weedall"),
            Ssn: provesdkservergo.String("265228370"),
        },
    }
    ctx := context.Background()
    res, err := s.V3.V3CompleteRequest(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.V3CompleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [components.V3CompleteRequest](../../models/components/v3completerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.V3CompleteRequestResponse](../../models/operations/v3completerequestresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 400,500            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## V3StartRequest

Send this request to start a Prove flow. It will return a correlation ID and an authToken for the client SDK.

### Example Usage

```go
package main

import(
	"github.com/prove-identity/prove-sdk-server-go/models/components"
	provesdkservergo "github.com/prove-identity/prove-sdk-server-go"
	"context"
	"log"
)

func main() {
    s := provesdkservergo.New(
        provesdkservergo.WithSecurity(components.Security{
            ClientID: "<YOUR_CLIENT_ID_HERE>",
            ClientSecret: "<YOUR_CLIENT_SECRET_HERE>",
        }),
    )
    var request *components.V3StartRequest = &components.V3StartRequest{
        DeviceID: provesdkservergo.String("713189b8-5555-4b08-83ba-75d08780aebd"),
        Dob: provesdkservergo.String("2024-05-02T00:00:00Z"),
        EmailAddress: provesdkservergo.String("jdoe@example.com"),
        FinalTargetURL: provesdkservergo.String("https://www.example.com/landing-page"),
        FlowID: provesdkservergo.String("prove-standard-prefill-i1"),
        FlowType: "mobile",
        IPAddress: provesdkservergo.String("10.0.0.1"),
        Last4SSN: provesdkservergo.String("1234"),
        PhoneNumber: provesdkservergo.String("12065550100"),
    }
    ctx := context.Background()
    res, err := s.V3.V3StartRequest(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.V3StartResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [components.V3StartRequest](../../models/components/v3startrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.V3StartRequestResponse](../../models/operations/v3startrequestresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 400,500            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## V3ValidateRequest

Send this request to check the phone number entered/discovered earlier in the flow is validated. It will return a correlation ID and the next step.

### Example Usage

```go
package main

import(
	"github.com/prove-identity/prove-sdk-server-go/models/components"
	provesdkservergo "github.com/prove-identity/prove-sdk-server-go"
	"context"
	"log"
)

func main() {
    s := provesdkservergo.New(
        provesdkservergo.WithSecurity(components.Security{
            ClientID: "<YOUR_CLIENT_ID_HERE>",
            ClientSecret: "<YOUR_CLIENT_SECRET_HERE>",
        }),
    )
    var request *components.V3ValidateRequest = &components.V3ValidateRequest{
        CorrelationID: "713189b8-5555-4b08-83ba-75d08780aebd",
    }
    ctx := context.Background()
    res, err := s.V3.V3ValidateRequest(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.V3ValidateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [components.V3ValidateRequest](../../models/components/v3validaterequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.V3ValidateRequestResponse](../../models/operations/v3validaterequestresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 400,500            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
