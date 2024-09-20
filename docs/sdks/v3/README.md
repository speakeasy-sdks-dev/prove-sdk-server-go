# V3
(*V3*)

## Overview

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
	"context"
	"github.com/prove-identity/prove-sdk-server-go/models/components"
	"log"
)

func main() {
    s := provesdkservergo.New()

    ctx := context.Background()
    res, err := s.V3.V3TokenRequest(ctx, &components.V3TokenRequest{
        ClientID: "customer_id",
        ClientSecret: "secret",
        GrantType: "client_credentials",
    })
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
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.V3TokenRequestResponse](../../models/operations/v3tokenrequestresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## V3ChallengeRequest

Send this request to submit challenge information. Either a DOB or last 4 of SSN needs to be submitted if neither was submitted to the /start endpoint (challenge fields submitted to this endpoint will overwrite the /start endpoint fields submitted). It will return a correlation ID, user information, and the next step to call in the flow. This capability is only available in Pre-Fill®, it's not available in Prove Identity®. You'll notice that when using Prove Identity®, if /validate is successful, it will then return `v3-complete` as one of the keys in the `Next` field map instead of `v3-challenge`.

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
            ClientID: provesdkservergo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: provesdkservergo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.V3.V3ChallengeRequest(ctx, &components.V3ChallengeRequest{
        CorrelationID: "713189b8-5555-4b08-83ba-75d08780aebd",
        Dob: provesdkservergo.String("1981-01"),
        Ssn: provesdkservergo.String("0596"),
    })
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
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.V3ChallengeRequestResponse](../../models/operations/v3challengerequestresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## V3CompleteRequest

Send this request to verify the user and complete the flow. It will return a correlation ID, user information, and the next step to call in the flow. There is a validation check that requires at least first + last name or SSN passed in, else an HTTP 400 is returned. Additionally, specific to the Pre-Fill® or Prove Identity® with KYC use case, you need to pass in first name, last name, DOB and SSN (or address) to ensure you receive back the KYC elements and correct CIP values.

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
            ClientID: provesdkservergo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: provesdkservergo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.V3.V3CompleteRequest(ctx, &components.V3CompleteRequest{
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
            Dob: provesdkservergo.String("1981-01"),
            EmailAddresses: []string{
                "jdoe@example.com",
                "dsmith@example.com",
            },
            FirstName: provesdkservergo.String("Tod"),
            LastName: provesdkservergo.String("Weedall"),
            Ssn: provesdkservergo.String("265228370"),
        },
    })
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
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.V3CompleteRequestResponse](../../models/operations/v3completerequestresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
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
            ClientID: provesdkservergo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: provesdkservergo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.V3.V3StartRequest(ctx, &components.V3StartRequest{
        Dob: provesdkservergo.String("1981-01"),
        EmailAddress: provesdkservergo.String("mpinsonm@dyndns.org"),
        FinalTargetURL: provesdkservergo.String("https://www.example.com/landing-page"),
        FlowType: "mobile",
        IPAddress: provesdkservergo.String("10.0.0.1"),
        PhoneNumber: provesdkservergo.String("2001001695"),
        SmsMessage: provesdkservergo.String("\"Your code is: ####.\""),
        Ssn: provesdkservergo.String("0596"),
    })
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
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.V3StartRequestResponse](../../models/operations/v3startrequestresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
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
            ClientID: provesdkservergo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: provesdkservergo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.V3.V3ValidateRequest(ctx, &components.V3ValidateRequest{
        CorrelationID: "713189b8-5555-4b08-83ba-75d08780aebd",
    })
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
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.V3ValidateRequestResponse](../../models/operations/v3validaterequestresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
