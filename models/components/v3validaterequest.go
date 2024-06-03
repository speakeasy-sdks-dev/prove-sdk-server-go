// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type V3ValidateRequest struct {
	// Correlation ID is the unique ID of the flow. To continue the flow, the field will also be used for each of the subsequent API calls in the same flow.
	CorrelationID string `json:"correlationId"`
}

func (o *V3ValidateRequest) GetCorrelationID() string {
	if o == nil {
		return ""
	}
	return o.CorrelationID
}
