/*
NabooApi V1

Here you have the first version of the naboo api to create checkout automatically

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package naboo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CashOutOrangeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashOutOrangeRequest{}

// CashOutOrangeRequest struct for CashOutOrangeRequest
type CashOutOrangeRequest struct {
	FullName string `json:"full_name"`
	Amount int32 `json:"amount"`
	PhoneNumber string `json:"phone_number"`
}

type _CashOutOrangeRequest CashOutOrangeRequest

// NewCashOutOrangeRequest instantiates a new CashOutOrangeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashOutOrangeRequest(fullName string, amount int32, phoneNumber string) *CashOutOrangeRequest {
	this := CashOutOrangeRequest{}
	this.FullName = fullName
	this.Amount = amount
	this.PhoneNumber = phoneNumber
	return &this
}

// NewCashOutOrangeRequestWithDefaults instantiates a new CashOutOrangeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashOutOrangeRequestWithDefaults() *CashOutOrangeRequest {
	this := CashOutOrangeRequest{}
	return &this
}

// GetFullName returns the FullName field value
func (o *CashOutOrangeRequest) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *CashOutOrangeRequest) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *CashOutOrangeRequest) SetFullName(v string) {
	o.FullName = v
}

// GetAmount returns the Amount field value
func (o *CashOutOrangeRequest) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CashOutOrangeRequest) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CashOutOrangeRequest) SetAmount(v int32) {
	o.Amount = v
}

// GetPhoneNumber returns the PhoneNumber field value
func (o *CashOutOrangeRequest) GetPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
func (o *CashOutOrangeRequest) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhoneNumber, true
}

// SetPhoneNumber sets field value
func (o *CashOutOrangeRequest) SetPhoneNumber(v string) {
	o.PhoneNumber = v
}

func (o CashOutOrangeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashOutOrangeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["full_name"] = o.FullName
	toSerialize["amount"] = o.Amount
	toSerialize["phone_number"] = o.PhoneNumber
	return toSerialize, nil
}

func (o *CashOutOrangeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"full_name",
		"amount",
		"phone_number",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCashOutOrangeRequest := _CashOutOrangeRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCashOutOrangeRequest)

	if err != nil {
		return err
	}

	*o = CashOutOrangeRequest(varCashOutOrangeRequest)

	return err
}

type NullableCashOutOrangeRequest struct {
	value *CashOutOrangeRequest
	isSet bool
}

func (v NullableCashOutOrangeRequest) Get() *CashOutOrangeRequest {
	return v.value
}

func (v *NullableCashOutOrangeRequest) Set(val *CashOutOrangeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCashOutOrangeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCashOutOrangeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashOutOrangeRequest(val *CashOutOrangeRequest) *NullableCashOutOrangeRequest {
	return &NullableCashOutOrangeRequest{value: val, isSet: true}
}

func (v NullableCashOutOrangeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashOutOrangeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


