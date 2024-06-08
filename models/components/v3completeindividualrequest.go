// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type V3CompleteIndividualRequest struct {
	// Addresses that belong to the individual.
	Addresses []V3CompleteAddressEntryRequest `json:"addresses,omitempty"`
	// DOB is the date of birth of the individual.
	Dob *string `json:"dob,omitempty"`
	// Email addresses that belong to the individual.
	EmailAddresses []string `json:"emailAddresses,omitempty"`
	// First name of the individual.
	FirstName *string `json:"firstName,omitempty"`
	// Last4SSN is last 4 digits of SSN.
	Last4SSN *string `json:"last4SSN,omitempty"`
	// Last name of the individual.
	LastName *string `json:"lastName,omitempty"`
	// SSN is the social security number of the individual.
	Ssn *string `json:"ssn,omitempty"`
}

func (o *V3CompleteIndividualRequest) GetAddresses() []V3CompleteAddressEntryRequest {
	if o == nil {
		return nil
	}
	return o.Addresses
}

func (o *V3CompleteIndividualRequest) GetDob() *string {
	if o == nil {
		return nil
	}
	return o.Dob
}

func (o *V3CompleteIndividualRequest) GetEmailAddresses() []string {
	if o == nil {
		return nil
	}
	return o.EmailAddresses
}

func (o *V3CompleteIndividualRequest) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *V3CompleteIndividualRequest) GetLast4SSN() *string {
	if o == nil {
		return nil
	}
	return o.Last4SSN
}

func (o *V3CompleteIndividualRequest) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *V3CompleteIndividualRequest) GetSsn() *string {
	if o == nil {
		return nil
	}
	return o.Ssn
}
