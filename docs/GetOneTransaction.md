# GetOneTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** |  | 
**MethodOfPayment** | [**[]Wallet**](Wallet.md) |  | 
**Amount** | **int32** |  | 
**AmountToPay** | **int32** |  | 
**Currency** | **string** |  | 
**CreatedAt** | **string** |  | 
**TransactionStatus** | [**TransactionStatusEnum**](TransactionStatusEnum.md) |  | 
**Products** | [**[]ProductModel**](ProductModel.md) |  | 
**IsDone** | **bool** |  | 
**IsEscrow** | **bool** |  | 
**IsMerchant** | **bool** |  | 
**CheckoutUrl** | **string** |  | 

## Methods

### NewGetOneTransaction

`func NewGetOneTransaction(orderId string, methodOfPayment []Wallet, amount int32, amountToPay int32, currency string, createdAt string, transactionStatus TransactionStatusEnum, products []ProductModel, isDone bool, isEscrow bool, isMerchant bool, checkoutUrl string, ) *GetOneTransaction`

NewGetOneTransaction instantiates a new GetOneTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOneTransactionWithDefaults

`func NewGetOneTransactionWithDefaults() *GetOneTransaction`

NewGetOneTransactionWithDefaults instantiates a new GetOneTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *GetOneTransaction) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetOneTransaction) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetOneTransaction) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetMethodOfPayment

`func (o *GetOneTransaction) GetMethodOfPayment() []Wallet`

GetMethodOfPayment returns the MethodOfPayment field if non-nil, zero value otherwise.

### GetMethodOfPaymentOk

`func (o *GetOneTransaction) GetMethodOfPaymentOk() (*[]Wallet, bool)`

GetMethodOfPaymentOk returns a tuple with the MethodOfPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodOfPayment

`func (o *GetOneTransaction) SetMethodOfPayment(v []Wallet)`

SetMethodOfPayment sets MethodOfPayment field to given value.


### GetAmount

`func (o *GetOneTransaction) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetOneTransaction) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetOneTransaction) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetAmountToPay

`func (o *GetOneTransaction) GetAmountToPay() int32`

GetAmountToPay returns the AmountToPay field if non-nil, zero value otherwise.

### GetAmountToPayOk

`func (o *GetOneTransaction) GetAmountToPayOk() (*int32, bool)`

GetAmountToPayOk returns a tuple with the AmountToPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountToPay

`func (o *GetOneTransaction) SetAmountToPay(v int32)`

SetAmountToPay sets AmountToPay field to given value.


### GetCurrency

`func (o *GetOneTransaction) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetOneTransaction) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetOneTransaction) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCreatedAt

`func (o *GetOneTransaction) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetOneTransaction) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetOneTransaction) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetTransactionStatus

`func (o *GetOneTransaction) GetTransactionStatus() TransactionStatusEnum`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *GetOneTransaction) GetTransactionStatusOk() (*TransactionStatusEnum, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *GetOneTransaction) SetTransactionStatus(v TransactionStatusEnum)`

SetTransactionStatus sets TransactionStatus field to given value.


### GetProducts

`func (o *GetOneTransaction) GetProducts() []ProductModel`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *GetOneTransaction) GetProductsOk() (*[]ProductModel, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *GetOneTransaction) SetProducts(v []ProductModel)`

SetProducts sets Products field to given value.


### SetProductsNil

`func (o *GetOneTransaction) SetProductsNil(b bool)`

 SetProductsNil sets the value for Products to be an explicit nil

### UnsetProducts
`func (o *GetOneTransaction) UnsetProducts()`

UnsetProducts ensures that no value is present for Products, not even an explicit nil
### GetIsDone

`func (o *GetOneTransaction) GetIsDone() bool`

GetIsDone returns the IsDone field if non-nil, zero value otherwise.

### GetIsDoneOk

`func (o *GetOneTransaction) GetIsDoneOk() (*bool, bool)`

GetIsDoneOk returns a tuple with the IsDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDone

`func (o *GetOneTransaction) SetIsDone(v bool)`

SetIsDone sets IsDone field to given value.


### GetIsEscrow

`func (o *GetOneTransaction) GetIsEscrow() bool`

GetIsEscrow returns the IsEscrow field if non-nil, zero value otherwise.

### GetIsEscrowOk

`func (o *GetOneTransaction) GetIsEscrowOk() (*bool, bool)`

GetIsEscrowOk returns a tuple with the IsEscrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEscrow

`func (o *GetOneTransaction) SetIsEscrow(v bool)`

SetIsEscrow sets IsEscrow field to given value.


### GetIsMerchant

`func (o *GetOneTransaction) GetIsMerchant() bool`

GetIsMerchant returns the IsMerchant field if non-nil, zero value otherwise.

### GetIsMerchantOk

`func (o *GetOneTransaction) GetIsMerchantOk() (*bool, bool)`

GetIsMerchantOk returns a tuple with the IsMerchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMerchant

`func (o *GetOneTransaction) SetIsMerchant(v bool)`

SetIsMerchant sets IsMerchant field to given value.


### GetCheckoutUrl

`func (o *GetOneTransaction) GetCheckoutUrl() string`

GetCheckoutUrl returns the CheckoutUrl field if non-nil, zero value otherwise.

### GetCheckoutUrlOk

`func (o *GetOneTransaction) GetCheckoutUrlOk() (*string, bool)`

GetCheckoutUrlOk returns a tuple with the CheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUrl

`func (o *GetOneTransaction) SetCheckoutUrl(v string)`

SetCheckoutUrl sets CheckoutUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


