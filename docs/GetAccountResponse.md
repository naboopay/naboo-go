# GetAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** |  | 
**Balance** | **float32** |  | 
**AccountIsActivate** | **bool** |  | 
**MethodOfPayment** | [**Wallet**](Wallet.md) |  | 
**LoyaltyCredit** | **int32** |  | 

## Methods

### NewGetAccountResponse

`func NewGetAccountResponse(accountNumber string, balance float32, accountIsActivate bool, methodOfPayment Wallet, loyaltyCredit int32, ) *GetAccountResponse`

NewGetAccountResponse instantiates a new GetAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountResponseWithDefaults

`func NewGetAccountResponseWithDefaults() *GetAccountResponse`

NewGetAccountResponseWithDefaults instantiates a new GetAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *GetAccountResponse) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *GetAccountResponse) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *GetAccountResponse) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetBalance

`func (o *GetAccountResponse) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *GetAccountResponse) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *GetAccountResponse) SetBalance(v float32)`

SetBalance sets Balance field to given value.


### GetAccountIsActivate

`func (o *GetAccountResponse) GetAccountIsActivate() bool`

GetAccountIsActivate returns the AccountIsActivate field if non-nil, zero value otherwise.

### GetAccountIsActivateOk

`func (o *GetAccountResponse) GetAccountIsActivateOk() (*bool, bool)`

GetAccountIsActivateOk returns a tuple with the AccountIsActivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIsActivate

`func (o *GetAccountResponse) SetAccountIsActivate(v bool)`

SetAccountIsActivate sets AccountIsActivate field to given value.


### GetMethodOfPayment

`func (o *GetAccountResponse) GetMethodOfPayment() Wallet`

GetMethodOfPayment returns the MethodOfPayment field if non-nil, zero value otherwise.

### GetMethodOfPaymentOk

`func (o *GetAccountResponse) GetMethodOfPaymentOk() (*Wallet, bool)`

GetMethodOfPaymentOk returns a tuple with the MethodOfPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodOfPayment

`func (o *GetAccountResponse) SetMethodOfPayment(v Wallet)`

SetMethodOfPayment sets MethodOfPayment field to given value.


### GetLoyaltyCredit

`func (o *GetAccountResponse) GetLoyaltyCredit() int32`

GetLoyaltyCredit returns the LoyaltyCredit field if non-nil, zero value otherwise.

### GetLoyaltyCreditOk

`func (o *GetAccountResponse) GetLoyaltyCreditOk() (*int32, bool)`

GetLoyaltyCreditOk returns a tuple with the LoyaltyCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyCredit

`func (o *GetAccountResponse) SetLoyaltyCredit(v int32)`

SetLoyaltyCredit sets LoyaltyCredit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


