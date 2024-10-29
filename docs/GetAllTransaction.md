# GetAllTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | [**[]GetOneTransaction**](GetOneTransaction.md) |  | 

## Methods

### NewGetAllTransaction

`func NewGetAllTransaction(transactions []GetOneTransaction, ) *GetAllTransaction`

NewGetAllTransaction instantiates a new GetAllTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllTransactionWithDefaults

`func NewGetAllTransactionWithDefaults() *GetAllTransaction`

NewGetAllTransactionWithDefaults instantiates a new GetAllTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *GetAllTransaction) GetTransactions() []GetOneTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *GetAllTransaction) GetTransactionsOk() (*[]GetOneTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *GetAllTransaction) SetTransactions(v []GetOneTransaction)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


