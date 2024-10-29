# Go API client for naboo

Here you have the first version of the naboo api to create checkout automatically

- API version: 0.1.0

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import naboo "github.com/naboopay/naboo-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `naboo.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), naboo.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `naboo.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), naboo.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `naboo.ContextOperationServerIndices` and `naboo.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), naboo.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), naboo.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to */api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountService* | [**GetAccountInformationAccountGet**](docs/AccountService.md#getaccountinformationaccountget) | **Get** /account/ | Get Account Information
*CashoutService* | [**CashOutOrangeMoneyCashoutOrangeMoneyPost**](docs/CashoutService.md#cashoutorangemoneycashoutorangemoneypost) | **Post** /cashout/orange-money | Cash Out Orange Money
*CashoutService* | [**CashOutWaveCashoutWavePost**](docs/CashoutService.md#cashoutwavecashoutwavepost) | **Post** /cashout/wave | Cash Out Wave
*TransactionService* | [**CreateTransactionTransactionCreateTransactionPost**](docs/TransactionService.md#createtransactiontransactioncreatetransactionpost) | **Put** /transaction/create-transaction | Create Transaction
*TransactionService* | [**DeleteTransactionTransactionDeleteTransactionDelete**](docs/TransactionService.md#deletetransactiontransactiondeletetransactiondelete) | **Delete** /transaction/delete-transaction | Delete Transaction
*TransactionService* | [**GetOneTransactionTransactionGetOneTransactionGet**](docs/TransactionService.md#getonetransactiontransactiongetonetransactionget) | **Get** /transaction/get-one-transaction | Get One Transaction
*TransactionService* | [**GetTransactionsTransactionGetTransactionsGet**](docs/TransactionService.md#gettransactionstransactiongettransactionsget) | **Get** /transaction/get-transactions | Get Transactions


## Documentation For Models

 - [CashOutOrangeRequest](docs/CashOutOrangeRequest.md)
 - [CashOutResponse](docs/CashOutResponse.md)
 - [CashOutWaveRequest](docs/CashOutWaveRequest.md)
 - [DeleteTransactionRequest](docs/DeleteTransactionRequest.md)
 - [DeleteTransactionResponse](docs/DeleteTransactionResponse.md)
 - [GetAccountResponse](docs/GetAccountResponse.md)
 - [GetAllTransaction](docs/GetAllTransaction.md)
 - [GetOneTransaction](docs/GetOneTransaction.md)
 - [HTTPValidationError](docs/HTTPValidationError.md)
 - [ProductModel](docs/ProductModel.md)
 - [TransactionRequest](docs/TransactionRequest.md)
 - [TransactionResponse](docs/TransactionResponse.md)
 - [TransactionStatusEnum](docs/TransactionStatusEnum.md)
 - [ValidationError](docs/ValidationError.md)
 - [ValidationErrorLocInner](docs/ValidationErrorLocInner.md)
 - [Wallet](docs/Wallet.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### HTTPBearer

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), naboo.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


