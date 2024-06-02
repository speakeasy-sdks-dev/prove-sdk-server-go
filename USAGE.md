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
<!-- End SDK Example Usage [usage] -->