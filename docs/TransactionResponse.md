# TransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** |  | 
**MethodOfPayment** | **[]string** |  | 
**Amount** | Pointer to **int32** |  | [optional] [default to 0]
**AmountToPay** | Pointer to **int32** |  | [optional] [default to 0]
**Currency** | **string** |  | 
**CreatedAt** | **string** |  | 
**TransactionStatus** | Pointer to **string** |  | [optional] [default to "pending"]
**IsEscrow** | Pointer to **bool** |  | [optional] [default to false]
**IsMerchant** | Pointer to **bool** |  | [optional] [default to false]
**CheckoutUrl** | **string** |  | 

## Methods

### NewTransactionResponse

`func NewTransactionResponse(orderId string, methodOfPayment []string, currency string, createdAt string, checkoutUrl string, ) *TransactionResponse`

NewTransactionResponse instantiates a new TransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionResponseWithDefaults

`func NewTransactionResponseWithDefaults() *TransactionResponse`

NewTransactionResponseWithDefaults instantiates a new TransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *TransactionResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *TransactionResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *TransactionResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetMethodOfPayment

`func (o *TransactionResponse) GetMethodOfPayment() []string`

GetMethodOfPayment returns the MethodOfPayment field if non-nil, zero value otherwise.

### GetMethodOfPaymentOk

`func (o *TransactionResponse) GetMethodOfPaymentOk() (*[]string, bool)`

GetMethodOfPaymentOk returns a tuple with the MethodOfPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodOfPayment

`func (o *TransactionResponse) SetMethodOfPayment(v []string)`

SetMethodOfPayment sets MethodOfPayment field to given value.


### GetAmount

`func (o *TransactionResponse) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionResponse) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionResponse) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAmountToPay

`func (o *TransactionResponse) GetAmountToPay() int32`

GetAmountToPay returns the AmountToPay field if non-nil, zero value otherwise.

### GetAmountToPayOk

`func (o *TransactionResponse) GetAmountToPayOk() (*int32, bool)`

GetAmountToPayOk returns a tuple with the AmountToPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountToPay

`func (o *TransactionResponse) SetAmountToPay(v int32)`

SetAmountToPay sets AmountToPay field to given value.

### HasAmountToPay

`func (o *TransactionResponse) HasAmountToPay() bool`

HasAmountToPay returns a boolean if a field has been set.

### GetCurrency

`func (o *TransactionResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransactionResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransactionResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCreatedAt

`func (o *TransactionResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransactionResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransactionResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetTransactionStatus

`func (o *TransactionResponse) GetTransactionStatus() string`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *TransactionResponse) GetTransactionStatusOk() (*string, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *TransactionResponse) SetTransactionStatus(v string)`

SetTransactionStatus sets TransactionStatus field to given value.

### HasTransactionStatus

`func (o *TransactionResponse) HasTransactionStatus() bool`

HasTransactionStatus returns a boolean if a field has been set.

### GetIsEscrow

`func (o *TransactionResponse) GetIsEscrow() bool`

GetIsEscrow returns the IsEscrow field if non-nil, zero value otherwise.

### GetIsEscrowOk

`func (o *TransactionResponse) GetIsEscrowOk() (*bool, bool)`

GetIsEscrowOk returns a tuple with the IsEscrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEscrow

`func (o *TransactionResponse) SetIsEscrow(v bool)`

SetIsEscrow sets IsEscrow field to given value.

### HasIsEscrow

`func (o *TransactionResponse) HasIsEscrow() bool`

HasIsEscrow returns a boolean if a field has been set.

### GetIsMerchant

`func (o *TransactionResponse) GetIsMerchant() bool`

GetIsMerchant returns the IsMerchant field if non-nil, zero value otherwise.

### GetIsMerchantOk

`func (o *TransactionResponse) GetIsMerchantOk() (*bool, bool)`

GetIsMerchantOk returns a tuple with the IsMerchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMerchant

`func (o *TransactionResponse) SetIsMerchant(v bool)`

SetIsMerchant sets IsMerchant field to given value.

### HasIsMerchant

`func (o *TransactionResponse) HasIsMerchant() bool`

HasIsMerchant returns a boolean if a field has been set.

### GetCheckoutUrl

`func (o *TransactionResponse) GetCheckoutUrl() string`

GetCheckoutUrl returns the CheckoutUrl field if non-nil, zero value otherwise.

### GetCheckoutUrlOk

`func (o *TransactionResponse) GetCheckoutUrlOk() (*string, bool)`

GetCheckoutUrlOk returns a tuple with the CheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUrl

`func (o *TransactionResponse) SetCheckoutUrl(v string)`

SetCheckoutUrl sets CheckoutUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


