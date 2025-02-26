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

// checks if the GetAccountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountResponse{}

// GetAccountResponse struct for GetAccountResponse
type GetAccountResponse struct {
	AccountNumber string `json:"account_number"`
	Balance float32 `json:"balance"`
	AccountIsActivate bool `json:"account_is_activate"`
	MethodOfPayment Wallet `json:"method_of_payment"`
	LoyaltyCredit int32 `json:"loyalty_credit"`
}

type _GetAccountResponse GetAccountResponse

// NewGetAccountResponse instantiates a new GetAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountResponse(accountNumber string, balance float32, accountIsActivate bool, methodOfPayment Wallet, loyaltyCredit int32) *GetAccountResponse {
	this := GetAccountResponse{}
	this.AccountNumber = accountNumber
	this.Balance = balance
	this.AccountIsActivate = accountIsActivate
	this.MethodOfPayment = methodOfPayment
	this.LoyaltyCredit = loyaltyCredit
	return &this
}

// NewGetAccountResponseWithDefaults instantiates a new GetAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountResponseWithDefaults() *GetAccountResponse {
	this := GetAccountResponse{}
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *GetAccountResponse) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *GetAccountResponse) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetBalance returns the Balance field value
func (o *GetAccountResponse) GetBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetBalanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *GetAccountResponse) SetBalance(v float32) {
	o.Balance = v
}

// GetAccountIsActivate returns the AccountIsActivate field value
func (o *GetAccountResponse) GetAccountIsActivate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AccountIsActivate
}

// GetAccountIsActivateOk returns a tuple with the AccountIsActivate field value
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetAccountIsActivateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountIsActivate, true
}

// SetAccountIsActivate sets field value
func (o *GetAccountResponse) SetAccountIsActivate(v bool) {
	o.AccountIsActivate = v
}

// GetMethodOfPayment returns the MethodOfPayment field value
func (o *GetAccountResponse) GetMethodOfPayment() Wallet {
	if o == nil {
		var ret Wallet
		return ret
	}

	return o.MethodOfPayment
}

// GetMethodOfPaymentOk returns a tuple with the MethodOfPayment field value
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetMethodOfPaymentOk() (*Wallet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MethodOfPayment, true
}

// SetMethodOfPayment sets field value
func (o *GetAccountResponse) SetMethodOfPayment(v Wallet) {
	o.MethodOfPayment = v
}

// GetLoyaltyCredit returns the LoyaltyCredit field value
func (o *GetAccountResponse) GetLoyaltyCredit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LoyaltyCredit
}

// GetLoyaltyCreditOk returns a tuple with the LoyaltyCredit field value
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetLoyaltyCreditOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoyaltyCredit, true
}

// SetLoyaltyCredit sets field value
func (o *GetAccountResponse) SetLoyaltyCredit(v int32) {
	o.LoyaltyCredit = v
}

func (o GetAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_number"] = o.AccountNumber
	toSerialize["balance"] = o.Balance
	toSerialize["account_is_activate"] = o.AccountIsActivate
	toSerialize["method_of_payment"] = o.MethodOfPayment
	toSerialize["loyalty_credit"] = o.LoyaltyCredit
	return toSerialize, nil
}

func (o *GetAccountResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_number",
		"balance",
		"account_is_activate",
		"method_of_payment",
		"loyalty_credit",
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

	varGetAccountResponse := _GetAccountResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetAccountResponse)

	if err != nil {
		return err
	}

	*o = GetAccountResponse(varGetAccountResponse)

	return err
}

type NullableGetAccountResponse struct {
	value *GetAccountResponse
	isSet bool
}

func (v NullableGetAccountResponse) Get() *GetAccountResponse {
	return v.value
}

func (v *NullableGetAccountResponse) Set(val *GetAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountResponse(val *GetAccountResponse) *NullableGetAccountResponse {
	return &NullableGetAccountResponse{value: val, isSet: true}
}

func (v NullableGetAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


