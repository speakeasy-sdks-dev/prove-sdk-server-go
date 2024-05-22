// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/payfone/prove-sdk-server-go/internal/utils"
)

type Individual struct {
	// Addresses that belong to the individual.
	Addresses []AddressEntry `json:"addresses,omitempty"`
	// DOB is the date of birth of the individual.
	Dob *string `default:"2024-05-02 00:00:00 +0000 UTC" json:"dob"`
	// Email addresses that belong to the individual.
	EmailAddresses []string `json:"emailAddresses,omitempty"`
	// First name of the individual.
	FirstName *string `default:"Tod" json:"firstName"`
	// Las name of the individual.
	LastName *string `default:"Weedall" json:"lastName"`
	// SSN is the social security number of the individual.
	Ssn *string `default:"265228370" json:"ssn"`
}

func (i Individual) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *Individual) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Individual) GetAddresses() []AddressEntry {
	if o == nil {
		return nil
	}
	return o.Addresses
}

func (o *Individual) GetDob() *string {
	if o == nil {
		return nil
	}
	return o.Dob
}

func (o *Individual) GetEmailAddresses() []string {
	if o == nil {
		return nil
	}
	return o.EmailAddresses
}

func (o *Individual) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *Individual) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *Individual) GetSsn() *string {
	if o == nil {
		return nil
	}
	return o.Ssn
}
