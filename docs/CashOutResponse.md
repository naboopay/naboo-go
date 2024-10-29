# CashOutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | **string** |  | 
**Amount** | **int32** |  | 
**FullName** | **string** |  | 
**Status** | [**TransactionStatusEnum**](TransactionStatusEnum.md) |  | 

## Methods

### NewCashOutResponse

`func NewCashOutResponse(phoneNumber string, amount int32, fullName string, status TransactionStatusEnum, ) *CashOutResponse`

NewCashOutResponse instantiates a new CashOutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashOutResponseWithDefaults

`func NewCashOutResponseWithDefaults() *CashOutResponse`

NewCashOutResponseWithDefaults instantiates a new CashOutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *CashOutResponse) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CashOutResponse) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CashOutResponse) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetAmount

`func (o *CashOutResponse) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CashOutResponse) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CashOutResponse) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetFullName

`func (o *CashOutResponse) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *CashOutResponse) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *CashOutResponse) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetStatus

`func (o *CashOutResponse) GetStatus() TransactionStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CashOutResponse) GetStatusOk() (*TransactionStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CashOutResponse) SetStatus(v TransactionStatusEnum)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


