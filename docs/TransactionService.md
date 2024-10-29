# \TransactionService

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTransactionTransactionCreateTransactionPost**](TransactionService.md#CreateTransactionTransactionCreateTransactionPost) | **Put** /transaction/create-transaction | Create Transaction
[**DeleteTransactionTransactionDeleteTransactionDelete**](TransactionService.md#DeleteTransactionTransactionDeleteTransactionDelete) | **Delete** /transaction/delete-transaction | Delete Transaction
[**GetOneTransactionTransactionGetOneTransactionGet**](TransactionService.md#GetOneTransactionTransactionGetOneTransactionGet) | **Get** /transaction/get-one-transaction | Get One Transaction
[**GetTransactionsTransactionGetTransactionsGet**](TransactionService.md#GetTransactionsTransactionGetTransactionsGet) | **Get** /transaction/get-transactions | Get Transactions



## CreateTransactionTransactionCreateTransactionPost

> TransactionResponse CreateTransactionTransactionCreateTransactionPost(ctx).TransactionRequest(transactionRequest).Execute()

Create Transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/naboopay/naboo-go"
)

func main() {
	transactionRequest := *openapiclient.NewTransactionRequest([]openapiclient.Wallet{openapiclient.Wallet("WAVE")}, []openapiclient.ProductModel{*openapiclient.NewProductModel("Name_example", "Category_example", int32(123), int32(123), "Description_example")}) // TransactionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionService.CreateTransactionTransactionCreateTransactionPost(context.Background()).TransactionRequest(transactionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionService.CreateTransactionTransactionCreateTransactionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTransactionTransactionCreateTransactionPost`: TransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionService.CreateTransactionTransactionCreateTransactionPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransactionTransactionCreateTransactionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionRequest** | [**TransactionRequest**](TransactionRequest.md) |  | 

### Return type

[**TransactionResponse**](TransactionResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTransactionTransactionDeleteTransactionDelete

> DeleteTransactionResponse DeleteTransactionTransactionDeleteTransactionDelete(ctx).DeleteTransactionRequest(deleteTransactionRequest).Execute()

Delete Transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/naboopay/naboo-go"
)

func main() {
	deleteTransactionRequest := *openapiclient.NewDeleteTransactionRequest("OrderId_example") // DeleteTransactionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionService.DeleteTransactionTransactionDeleteTransactionDelete(context.Background()).DeleteTransactionRequest(deleteTransactionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionService.DeleteTransactionTransactionDeleteTransactionDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTransactionTransactionDeleteTransactionDelete`: DeleteTransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionService.DeleteTransactionTransactionDeleteTransactionDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTransactionTransactionDeleteTransactionDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteTransactionRequest** | [**DeleteTransactionRequest**](DeleteTransactionRequest.md) |  | 

### Return type

[**DeleteTransactionResponse**](DeleteTransactionResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOneTransactionTransactionGetOneTransactionGet

> GetOneTransaction GetOneTransactionTransactionGetOneTransactionGet(ctx).OrderId(orderId).Execute()

Get One Transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/naboopay/naboo-go"
)

func main() {
	orderId := "orderId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionService.GetOneTransactionTransactionGetOneTransactionGet(context.Background()).OrderId(orderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionService.GetOneTransactionTransactionGetOneTransactionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOneTransactionTransactionGetOneTransactionGet`: GetOneTransaction
	fmt.Fprintf(os.Stdout, "Response from `TransactionService.GetOneTransactionTransactionGetOneTransactionGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOneTransactionTransactionGetOneTransactionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **string** |  | 

### Return type

[**GetOneTransaction**](GetOneTransaction.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionsTransactionGetTransactionsGet

> GetAllTransaction GetTransactionsTransactionGetTransactionsGet(ctx).Execute()

Get Transactions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/naboopay/naboo-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionService.GetTransactionsTransactionGetTransactionsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionService.GetTransactionsTransactionGetTransactionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionsTransactionGetTransactionsGet`: GetAllTransaction
	fmt.Fprintf(os.Stdout, "Response from `TransactionService.GetTransactionsTransactionGetTransactionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsTransactionGetTransactionsGetRequest struct via the builder pattern


### Return type

[**GetAllTransaction**](GetAllTransaction.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

