# \AccountService

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountInformationAccountGet**](AccountService.md#GetAccountInformationAccountGet) | **Get** /account/ | Get Account Information



## GetAccountInformationAccountGet

> GetAccountResponse GetAccountInformationAccountGet(ctx).Execute()

Get Account Information

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
	resp, r, err := apiClient.AccountService.GetAccountInformationAccountGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.GetAccountInformationAccountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountInformationAccountGet`: GetAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountService.GetAccountInformationAccountGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountInformationAccountGetRequest struct via the builder pattern


### Return type

[**GetAccountResponse**](GetAccountResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

