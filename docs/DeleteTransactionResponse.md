# DeleteTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** |  | 
**Message** | **string** |  | 

## Methods

### NewDeleteTransactionResponse

`func NewDeleteTransactionResponse(orderId string, message string, ) *DeleteTransactionResponse`

NewDeleteTransactionResponse instantiates a new DeleteTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteTransactionResponseWithDefaults

`func NewDeleteTransactionResponseWithDefaults() *DeleteTransactionResponse`

NewDeleteTransactionResponseWithDefaults instantiates a new DeleteTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *DeleteTransactionResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *DeleteTransactionResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *DeleteTransactionResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetMessage

`func (o *DeleteTransactionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeleteTransactionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeleteTransactionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


