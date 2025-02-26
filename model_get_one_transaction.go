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

// checks if the GetOneTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOneTransaction{}

// GetOneTransaction struct for GetOneTransaction
type GetOneTransaction struct {
	OrderId string `json:"order_id"`
	MethodOfPayment []Wallet `json:"method_of_payment"`
	Amount int32 `json:"amount"`
	AmountToPay int32 `json:"amount_to_pay"`
	Currency string `json:"currency"`
	CreatedAt string `json:"created_at"`
	TransactionStatus TransactionStatusEnum `json:"transaction_status"`
	Products []ProductModel `json:"products"`
	IsDone bool `json:"is_done"`
	IsEscrow bool `json:"is_escrow"`
	IsMerchant bool `json:"is_merchant"`
	CheckoutUrl string `json:"checkout_url"`
}

type _GetOneTransaction GetOneTransaction

// NewGetOneTransaction instantiates a new GetOneTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOneTransaction(orderId string, methodOfPayment []Wallet, amount int32, amountToPay int32, currency string, createdAt string, transactionStatus TransactionStatusEnum, products []ProductModel, isDone bool, isEscrow bool, isMerchant bool, checkoutUrl string) *GetOneTransaction {
	this := GetOneTransaction{}
	this.OrderId = orderId
	this.MethodOfPayment = methodOfPayment
	this.Amount = amount
	this.AmountToPay = amountToPay
	this.Currency = currency
	this.CreatedAt = createdAt
	this.TransactionStatus = transactionStatus
	this.Products = products
	this.IsDone = isDone
	this.IsEscrow = isEscrow
	this.IsMerchant = isMerchant
	this.CheckoutUrl = checkoutUrl
	return &this
}

// NewGetOneTransactionWithDefaults instantiates a new GetOneTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOneTransactionWithDefaults() *GetOneTransaction {
	this := GetOneTransaction{}
	return &this
}

// GetOrderId returns the OrderId field value
func (o *GetOneTransaction) GetOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *GetOneTransaction) GetOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *GetOneTransaction) SetOrderId(v string) {
	o.OrderId = v
}

// GetMethodOfPayment returns the MethodOfPayment field value
func (o *GetOneTransaction) GetMethodOfPayment() []Wallet {
	if o == nil {
		var ret []Wallet
		return ret
	}

	return o.MethodOfPayment
}

// GetMethodOfPaymentOk returns a tuple with the MethodOfPayment field value
// and a boolean to check if the value has been set.
func (o *GetOneTransaction) GetMethodOfPaymentOk() ([]Wallet, bool) {
	if o == nil {
		return nil, false
	}
	return o.MethodOfPayment, true
}

// SetMethodOfPayment sets field value
func (o *GetOneTransaction) SetMethodOfPayment(v []Wallet) {
	o.MethodOfPayment = v
}

// GetAmount returns the Amount field value
func (o *GetOneTransaction) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetOneTransaction) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetOneTransaction) SetAmount(v int32) {
	o.Amount = v
}

// GetAmountToPay returns the AmountToPay field value
func (o *GetOneTransaction) GetAmountToPay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountToPay
}

// GetAmountToPayOk returns a tuple with the AmountToPay field value
// and a boolean to check if the value has been set.
func (o *GetOneTransaction) GetAmountToPayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountToPay, true
}

// SetAmountToPay sets field value
func (o *GetOneTransaction) SetAmountToPay(v int32) {
	o.AmountToPay = v
}

// GetCurrency returns the Currency field value
func (o *GetOneTransaction) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *GetOneTransaction) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *GetOneTransaction) SetCurrency(v string) {
	o.Currency = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *GetOneTransaction) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GetOneTransaction) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GetOneTransaction) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetTransactionStatus returns the TransactionStatus field value
func (o *GetOneTransaction) GetTransactionStatus() TransactionStatusEnum {
	if o == nil {
		var ret TransactionStatusEnum
		return ret
	}

	return o.TransactionStatus
}

// GetTransactionStatusOk returns a tuple with the TransactionStatus field value
// and a boolean to check if the value has been set.
func (o *GetOneTransaction) GetTransactionStatusOk() (*TransactionStatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionStatus, true
}

// SetTransactionStatus sets field value
func (o *GetOneTransaction) SetTransactionStatus(v TransactionStatusEnum) {
	o.TransactionStatus = v
}

// GetProducts returns the Products field value
// If the value is explicit nil, the zero value for []ProductModel will be returned
func (o *GetOneTransaction) GetProducts() []ProductModel {
	if o == nil {
		var ret []ProductModel
		return ret
	}

	return o.Products
}

// GetProductsOk returns a tuple with the Products field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetOneTransaction) GetProductsOk() ([]ProductModel, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// SetProducts sets field value
func (o *GetOneTransaction) SetProducts(v []ProductModel) {
	o.Products = v
}

// GetIsDone returns the IsDone field value
func (o *GetOneTransaction) GetIsDone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDone
}

// GetIsDoneOk returns a tuple with the IsDone field value
// and a boolean to check if the value has been set.
func (o *GetOneTransaction) GetIsDoneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDone, true
}

// SetIsDone sets field value
func (o *GetOneTransaction) SetIsDone(v bool) {
	o.IsDone = v
}

// GetIsEscrow returns the IsEscrow field value
func (o *GetOneTransaction) GetIsEscrow() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEscrow
}

// GetIsEscrowOk returns a tuple with the IsEscrow field value
// and a boolean to check if the value has been set.
func (o *GetOneTransaction) GetIsEscrowOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEscrow, true
}

// SetIsEscrow sets field value
func (o *GetOneTransaction) SetIsEscrow(v bool) {
	o.IsEscrow = v
}

// GetIsMerchant returns the IsMerchant field value
func (o *GetOneTransaction) GetIsMerchant() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMerchant
}

// GetIsMerchantOk returns a tuple with the IsMerchant field value
// and a boolean to check if the value has been set.
func (o *GetOneTransaction) GetIsMerchantOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMerchant, true
}

// SetIsMerchant sets field value
func (o *GetOneTransaction) SetIsMerchant(v bool) {
	o.IsMerchant = v
}

// GetCheckoutUrl returns the CheckoutUrl field value
func (o *GetOneTransaction) GetCheckoutUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CheckoutUrl
}

// GetCheckoutUrlOk returns a tuple with the CheckoutUrl field value
// and a boolean to check if the value has been set.
func (o *GetOneTransaction) GetCheckoutUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckoutUrl, true
}

// SetCheckoutUrl sets field value
func (o *GetOneTransaction) SetCheckoutUrl(v string) {
	o.CheckoutUrl = v
}

func (o GetOneTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOneTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["order_id"] = o.OrderId
	toSerialize["method_of_payment"] = o.MethodOfPayment
	toSerialize["amount"] = o.Amount
	toSerialize["amount_to_pay"] = o.AmountToPay
	toSerialize["currency"] = o.Currency
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["transaction_status"] = o.TransactionStatus
	if o.Products != nil {
		toSerialize["products"] = o.Products
	}
	toSerialize["is_done"] = o.IsDone
	toSerialize["is_escrow"] = o.IsEscrow
	toSerialize["is_merchant"] = o.IsMerchant
	toSerialize["checkout_url"] = o.CheckoutUrl
	return toSerialize, nil
}

func (o *GetOneTransaction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"order_id",
		"method_of_payment",
		"amount",
		"amount_to_pay",
		"currency",
		"created_at",
		"transaction_status",
		"products",
		"is_done",
		"is_escrow",
		"is_merchant",
		"checkout_url",
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

	varGetOneTransaction := _GetOneTransaction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetOneTransaction)

	if err != nil {
		return err
	}

	*o = GetOneTransaction(varGetOneTransaction)

	return err
}

type NullableGetOneTransaction struct {
	value *GetOneTransaction
	isSet bool
}

func (v NullableGetOneTransaction) Get() *GetOneTransaction {
	return v.value
}

func (v *NullableGetOneTransaction) Set(val *GetOneTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOneTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOneTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOneTransaction(val *GetOneTransaction) *NullableGetOneTransaction {
	return &NullableGetOneTransaction{value: val, isSet: true}
}

func (v NullableGetOneTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOneTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


