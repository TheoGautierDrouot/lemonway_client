# \MoneyInsPayPalApi

All URIs are relative to *https://sandbox-api.lemonway.fr/mb/drouot/dev/directkitrest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MoneyInsMoneyInPayPalInit**](MoneyInsPayPalApi.md#MoneyInsMoneyInPayPalInit) | **Post** /v2/moneyins/paypal/init | Initate Pay by PayPal
[**MoneyInsPayPalTransactionResume**](MoneyInsPayPalApi.md#MoneyInsPayPalTransactionResume) | **Post** /v2/moneyins/paypal/{transactionId}/resume | PayPal Resume



## MoneyInsMoneyInPayPalInit

> InitPayPalTransactionOutput MoneyInsMoneyInPayPalInit(ctx).Authorization(authorization).PSUIPAddress(pSUIPAddress).Input(input).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Initate Pay by PayPal



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "authorization_example" // string | Authorization bearer (OAuth 2)
    pSUIPAddress := "pSUIPAddress_example" // string | IP address of the final client (PSU).
    input := *openapiclient.NewInitPayPalTransactionInput(*openapiclient.NewRedirections("ReturnUrl_example", "ErrorUrl_example", "CancelUrl_example"), *openapiclient.NewTransaction("Reference_example", "AccountId_example", int32(123), int32(123)), []openapiclient.PurchaseItem{*openapiclient.NewPurchaseItem("MerchantAccountId_example", "Description_example", int32(123), int32(123), int32(123))}) // InitPayPalTransactionInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyInsPayPalApi.MoneyInsMoneyInPayPalInit(context.Background()).Authorization(authorization).PSUIPAddress(pSUIPAddress).Input(input).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyInsPayPalApi.MoneyInsMoneyInPayPalInit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyInsMoneyInPayPalInit`: InitPayPalTransactionOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyInsPayPalApi.MoneyInsMoneyInPayPalInit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoneyInsMoneyInPayPalInitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **input** | [**InitPayPalTransactionInput**](InitPayPalTransactionInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**InitPayPalTransactionOutput**](InitPayPalTransactionOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoneyInsPayPalTransactionResume

> MoneyInOutput MoneyInsPayPalTransactionResume(ctx, transactionId).Authorization(authorization).PSUIPAddress(pSUIPAddress).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

PayPal Resume



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transactionId := int64(789) // int64 | 
    authorization := "authorization_example" // string | Authorization bearer (OAuth 2)
    pSUIPAddress := "pSUIPAddress_example" // string | IP address of the final client (PSU).
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyInsPayPalApi.MoneyInsPayPalTransactionResume(context.Background(), transactionId).Authorization(authorization).PSUIPAddress(pSUIPAddress).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyInsPayPalApi.MoneyInsPayPalTransactionResume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyInsPayPalTransactionResume`: MoneyInOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyInsPayPalApi.MoneyInsPayPalTransactionResume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoneyInsPayPalTransactionResumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**MoneyInOutput**](MoneyInOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

