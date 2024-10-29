# \CashoutService

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CashOutOrangeMoneyCashoutOrangeMoneyPost**](CashoutService.md#CashOutOrangeMoneyCashoutOrangeMoneyPost) | **Post** /cashout/orange-money | Cash Out Orange Money
[**CashOutWaveCashoutWavePost**](CashoutService.md#CashOutWaveCashoutWavePost) | **Post** /cashout/wave | Cash Out Wave



## CashOutOrangeMoneyCashoutOrangeMoneyPost

> CashOutResponse CashOutOrangeMoneyCashoutOrangeMoneyPost(ctx).CashOutOrangeRequest(cashOutOrangeRequest).Execute()

Cash Out Orange Money



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
	cashOutOrangeRequest := *openapiclient.NewCashOutOrangeRequest("FullName_example", int32(123), "PhoneNumber_example") // CashOutOrangeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CashoutService.CashOutOrangeMoneyCashoutOrangeMoneyPost(context.Background()).CashOutOrangeRequest(cashOutOrangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CashoutService.CashOutOrangeMoneyCashoutOrangeMoneyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CashOutOrangeMoneyCashoutOrangeMoneyPost`: CashOutResponse
	fmt.Fprintf(os.Stdout, "Response from `CashoutService.CashOutOrangeMoneyCashoutOrangeMoneyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCashOutOrangeMoneyCashoutOrangeMoneyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cashOutOrangeRequest** | [**CashOutOrangeRequest**](CashOutOrangeRequest.md) |  | 

### Return type

[**CashOutResponse**](CashOutResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CashOutWaveCashoutWavePost

> CashOutResponse CashOutWaveCashoutWavePost(ctx).CashOutWaveRequest(cashOutWaveRequest).Execute()

Cash Out Wave



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
	cashOutWaveRequest := *openapiclient.NewCashOutWaveRequest("FullName_example", int32(123), "PhoneNumber_example") // CashOutWaveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CashoutService.CashOutWaveCashoutWavePost(context.Background()).CashOutWaveRequest(cashOutWaveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CashoutService.CashOutWaveCashoutWavePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CashOutWaveCashoutWavePost`: CashOutResponse
	fmt.Fprintf(os.Stdout, "Response from `CashoutService.CashOutWaveCashoutWavePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCashOutWaveCashoutWavePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cashOutWaveRequest** | [**CashOutWaveRequest**](CashOutWaveRequest.md) |  | 

### Return type

[**CashOutResponse**](CashOutResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

