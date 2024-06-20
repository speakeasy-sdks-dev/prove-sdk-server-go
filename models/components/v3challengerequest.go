// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type V3ChallengeRequest struct {
	// Correlation ID is the unique ID of the flow. To continue the flow, the field will also be used for each of the subsequent API calls in the same flow.
	CorrelationID string `json:"correlationId"`
	// DOB is the date of birth in this format: YYYY-MM-DD. Acceptable characters are: numeric with symbol '-'.
	Dob *string `json:"dob,omitempty"`
	// SSN is either the full or last 4 numbers of the social security number. Acceptable characters are: numeric.
	Ssn *string `json:"ssn,omitempty"`
}

func (o *V3ChallengeRequest) GetCorrelationID() string {
	if o == nil {
		return ""
	}
	return o.CorrelationID
}

func (o *V3ChallengeRequest) GetDob() *string {
	if o == nil {
		return nil
	}
	return o.Dob
}

func (o *V3ChallengeRequest) GetSsn() *string {
	if o == nil {
		return nil
	}
	return o.Ssn
}
