<!-- Start SDK Example Usage [usage] -->
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
			ClientID:     "<YOUR_CLIENT_ID_HERE>",
			ClientSecret: "<YOUR_CLIENT_SECRET_HERE>",
		}),
	)
	var request *components.V3StartRequest = &components.V3StartRequest{
		DeviceID:       provesdkservergo.String("713189b8-5555-4b08-83ba-75d08780aebd"),
		Dob:            provesdkservergo.String("2024-05-02T00:00:00Z"),
		EmailAddress:   provesdkservergo.String("jdoe@example.com"),
		FinalTargetURL: provesdkservergo.String("https://www.example.com/landing-page"),
		FlowID:         provesdkservergo.String("prove-standard-prefill-i1"),
		FlowType:       "mobile",
		IPAddress:      provesdkservergo.String("10.0.0.1"),
		Last4SSN:       provesdkservergo.String("1234"),
		PhoneNumber:    provesdkservergo.String("12065550100"),
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
<!-- End SDK Example Usage [usage] -->