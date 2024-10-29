# TransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MethodOfPayment** | [**[]Wallet**](Wallet.md) |  | 
**Products** | [**[]ProductModel**](ProductModel.md) |  | 
**SuccessUrl** | Pointer to **NullableString** |  | [optional] 
**ErrorUrl** | Pointer to **NullableString** |  | [optional] 
**IsEscrow** | Pointer to **bool** |  | [optional] [default to false]
**IsMerchant** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewTransactionRequest

`func NewTransactionRequest(methodOfPayment []Wallet, products []ProductModel, ) *TransactionRequest`

NewTransactionRequest instantiates a new TransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRequestWithDefaults

`func NewTransactionRequestWithDefaults() *TransactionRequest`

NewTransactionRequestWithDefaults instantiates a new TransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethodOfPayment

`func (o *TransactionRequest) GetMethodOfPayment() []Wallet`

GetMethodOfPayment returns the MethodOfPayment field if non-nil, zero value otherwise.

### GetMethodOfPaymentOk

`func (o *TransactionRequest) GetMethodOfPaymentOk() (*[]Wallet, bool)`

GetMethodOfPaymentOk returns a tuple with the MethodOfPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodOfPayment

`func (o *TransactionRequest) SetMethodOfPayment(v []Wallet)`

SetMethodOfPayment sets MethodOfPayment field to given value.


### GetProducts

`func (o *TransactionRequest) GetProducts() []ProductModel`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *TransactionRequest) GetProductsOk() (*[]ProductModel, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *TransactionRequest) SetProducts(v []ProductModel)`

SetProducts sets Products field to given value.


### SetProductsNil

`func (o *TransactionRequest) SetProductsNil(b bool)`

 SetProductsNil sets the value for Products to be an explicit nil

### UnsetProducts
`func (o *TransactionRequest) UnsetProducts()`

UnsetProducts ensures that no value is present for Products, not even an explicit nil
### GetSuccessUrl

`func (o *TransactionRequest) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *TransactionRequest) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *TransactionRequest) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *TransactionRequest) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.

### SetSuccessUrlNil

`func (o *TransactionRequest) SetSuccessUrlNil(b bool)`

 SetSuccessUrlNil sets the value for SuccessUrl to be an explicit nil

### UnsetSuccessUrl
`func (o *TransactionRequest) UnsetSuccessUrl()`

UnsetSuccessUrl ensures that no value is present for SuccessUrl, not even an explicit nil
### GetErrorUrl

`func (o *TransactionRequest) GetErrorUrl() string`

GetErrorUrl returns the ErrorUrl field if non-nil, zero value otherwise.

### GetErrorUrlOk

`func (o *TransactionRequest) GetErrorUrlOk() (*string, bool)`

GetErrorUrlOk returns a tuple with the ErrorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorUrl

`func (o *TransactionRequest) SetErrorUrl(v string)`

SetErrorUrl sets ErrorUrl field to given value.

### HasErrorUrl

`func (o *TransactionRequest) HasErrorUrl() bool`

HasErrorUrl returns a boolean if a field has been set.

### SetErrorUrlNil

`func (o *TransactionRequest) SetErrorUrlNil(b bool)`

 SetErrorUrlNil sets the value for ErrorUrl to be an explicit nil

### UnsetErrorUrl
`func (o *TransactionRequest) UnsetErrorUrl()`

UnsetErrorUrl ensures that no value is present for ErrorUrl, not even an explicit nil
### GetIsEscrow

`func (o *TransactionRequest) GetIsEscrow() bool`

GetIsEscrow returns the IsEscrow field if non-nil, zero value otherwise.

### GetIsEscrowOk

`func (o *TransactionRequest) GetIsEscrowOk() (*bool, bool)`

GetIsEscrowOk returns a tuple with the IsEscrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEscrow

`func (o *TransactionRequest) SetIsEscrow(v bool)`

SetIsEscrow sets IsEscrow field to given value.

### HasIsEscrow

`func (o *TransactionRequest) HasIsEscrow() bool`

HasIsEscrow returns a boolean if a field has been set.

### GetIsMerchant

`func (o *TransactionRequest) GetIsMerchant() bool`

GetIsMerchant returns the IsMerchant field if non-nil, zero value otherwise.

### GetIsMerchantOk

`func (o *TransactionRequest) GetIsMerchantOk() (*bool, bool)`

GetIsMerchantOk returns a tuple with the IsMerchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMerchant

`func (o *TransactionRequest) SetIsMerchant(v bool)`

SetIsMerchant sets IsMerchant field to given value.

### HasIsMerchant

`func (o *TransactionRequest) HasIsMerchant() bool`

HasIsMerchant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


