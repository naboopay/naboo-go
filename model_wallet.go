/*
NabooApi V1

Here you have the first version of the naboo api to create checkout automatically

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package naboo

import (
	"encoding/json"
	"fmt"
)

// Wallet the model 'Wallet'
type Wallet string

// List of Wallet
const (
	WAVE Wallet = "WAVE"
	ORANGE_MONEY Wallet = "ORANGE_MONEY"
	FREE_MONEY Wallet = "FREE_MONEY"
)

// All allowed values of Wallet enum
var AllowedWalletEnumValues = []Wallet{
	"WAVE",
	"ORANGE_MONEY",
	"FREE_MONEY",
}

func (v *Wallet) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Wallet(value)
	for _, existing := range AllowedWalletEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Wallet", value)
}

// NewWalletFromValue returns a pointer to a valid Wallet
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWalletFromValue(v string) (*Wallet, error) {
	ev := Wallet(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Wallet: valid values are %v", v, AllowedWalletEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Wallet) IsValid() bool {
	for _, existing := range AllowedWalletEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Wallet value
func (v Wallet) Ptr() *Wallet {
	return &v
}

type NullableWallet struct {
	value *Wallet
	isSet bool
}

func (v NullableWallet) Get() *Wallet {
	return v.value
}

func (v *NullableWallet) Set(val *Wallet) {
	v.value = val
	v.isSet = true
}

func (v NullableWallet) IsSet() bool {
	return v.isSet
}

func (v *NullableWallet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWallet(val *Wallet) *NullableWallet {
	return &NullableWallet{value: val, isSet: true}
}

func (v NullableWallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWallet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

