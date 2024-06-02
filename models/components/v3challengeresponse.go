// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type V3ChallengeResponse struct {
	Individual Individual `json:"individual"`
	// Next contains the next set of allowed calls in the same flow.
	Next map[string]string `json:"next"`
	// Success returns true if the challenge was accepted and user info retrieved.
	Success bool `json:"success"`
}

func (o *V3ChallengeResponse) GetIndividual() Individual {
	if o == nil {
		return Individual{}
	}
	return o.Individual
}

func (o *V3ChallengeResponse) GetNext() map[string]string {
	if o == nil {
		return map[string]string{}
	}
	return o.Next
}

func (o *V3ChallengeResponse) GetSuccess() bool {
	if o == nil {
		return false
	}
	return o.Success
}
