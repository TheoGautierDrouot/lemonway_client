# \MoneyOutsApi

All URIs are relative to *https://sandbox-api.lemonway.fr/mb/drouot/dev/directkitrest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MoneyOutsCancelPut**](MoneyOutsApi.md#MoneyOutsCancelPut) | **Put** /v2/moneyouts/{transactionid}/cancel | Money-Out Cancellation
[**MoneyOutsIbanExtendedPost**](MoneyOutsApi.md#MoneyOutsIbanExtendedPost) | **Post** /v2/moneyouts/iban/extended | Add Bank Account to a Payment Account for Money-Outs
[**MoneyOutsIbanGet**](MoneyOutsApi.md#MoneyOutsIbanGet) | **Get** /v2/moneyouts/{accountid}/iban | Get the IBAN Associated with a Payment Account
[**MoneyOutsIbanPost**](MoneyOutsApi.md#MoneyOutsIbanPost) | **Post** /v2/moneyouts/iban | Add an IBAN to a Payment Account for Money-Outs
[**MoneyOutsIbanUnregisterPut**](MoneyOutsApi.md#MoneyOutsIbanUnregisterPut) | **Put** /v2/moneyouts/iban/{IbanId}/unregister | Disable Bank Information (IBAN) from a Payment Account
[**MoneyOutsMoneyOutGet**](MoneyOutsApi.md#MoneyOutsMoneyOutGet) | **Get** /v2/moneyouts | Search for a Money-Out
[**MoneyOutsMoneyOutPost**](MoneyOutsApi.md#MoneyOutsMoneyOutPost) | **Post** /v2/moneyouts | External Fund Transfer from a Payment Account to a Bank Account



## MoneyOutsCancelPut

> CancelMoneyOutOutput MoneyOutsCancelPut(ctx, transactionid).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Money-Out Cancellation



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
    transactionid := int64(789) // int64 | 
    authorization := "authorization_example" // string | Authorization bearer (OAuth 2)
    pSUIPAddress := "pSUIPAddress_example" // string | IP address of the final client (PSU).
    parameters := *openapiclient.NewCancelMoneyOutInput() // CancelMoneyOutInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyOutsApi.MoneyOutsCancelPut(context.Background(), transactionid).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyOutsApi.MoneyOutsCancelPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyOutsCancelPut`: CancelMoneyOutOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyOutsApi.MoneyOutsCancelPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionid** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoneyOutsCancelPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**CancelMoneyOutInput**](CancelMoneyOutInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**CancelMoneyOutOutput**](CancelMoneyOutOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoneyOutsIbanExtendedPost

> RegisterIBANExtendedOutput MoneyOutsIbanExtendedPost(ctx).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Add Bank Account to a Payment Account for Money-Outs



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
    parameters := *openapiclient.NewRegisterIBANExtendedInput("Wallet_example", "HolderName_example", "AccountNumber_example", "BankCountry_example") // RegisterIBANExtendedInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyOutsApi.MoneyOutsIbanExtendedPost(context.Background()).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyOutsApi.MoneyOutsIbanExtendedPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyOutsIbanExtendedPost`: RegisterIBANExtendedOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyOutsApi.MoneyOutsIbanExtendedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoneyOutsIbanExtendedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**RegisterIBANExtendedInput**](RegisterIBANExtendedInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**RegisterIBANExtendedOutput**](RegisterIBANExtendedOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoneyOutsIbanGet

> AccountIbansOutput MoneyOutsIbanGet(ctx, accountid).Authorization(authorization).PSUIPAddress(pSUIPAddress).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Get the IBAN Associated with a Payment Account

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
    accountid := "accountid_example" // string | Account ID
    authorization := "authorization_example" // string | Authorization bearer (OAuth 2)
    pSUIPAddress := "pSUIPAddress_example" // string | IP address of the final client (PSU).
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyOutsApi.MoneyOutsIbanGet(context.Background(), accountid).Authorization(authorization).PSUIPAddress(pSUIPAddress).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyOutsApi.MoneyOutsIbanGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyOutsIbanGet`: AccountIbansOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyOutsApi.MoneyOutsIbanGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoneyOutsIbanGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**AccountIbansOutput**](AccountIbansOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoneyOutsIbanPost

> RegisterIBANOutput MoneyOutsIbanPost(ctx).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Add an IBAN to a Payment Account for Money-Outs



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
    parameters := *openapiclient.NewRegisterIBANInput("AccountId_example", "Holder_example", "Iban_example") // RegisterIBANInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyOutsApi.MoneyOutsIbanPost(context.Background()).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyOutsApi.MoneyOutsIbanPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyOutsIbanPost`: RegisterIBANOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyOutsApi.MoneyOutsIbanPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoneyOutsIbanPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**RegisterIBANInput**](RegisterIBANInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**RegisterIBANOutput**](RegisterIBANOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoneyOutsIbanUnregisterPut

> UnregisterIBANOutput MoneyOutsIbanUnregisterPut(ctx, ibanId).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Disable Bank Information (IBAN) from a Payment Account



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
    ibanId := int64(789) // int64 | 
    authorization := "authorization_example" // string | Authorization bearer (OAuth 2)
    pSUIPAddress := "pSUIPAddress_example" // string | IP address of the final client (PSU).
    parameters := *openapiclient.NewUnregisterIBANInput("Wallet_example") // UnregisterIBANInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyOutsApi.MoneyOutsIbanUnregisterPut(context.Background(), ibanId).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyOutsApi.MoneyOutsIbanUnregisterPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyOutsIbanUnregisterPut`: UnregisterIBANOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyOutsApi.MoneyOutsIbanUnregisterPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ibanId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoneyOutsIbanUnregisterPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**UnregisterIBANInput**](UnregisterIBANInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**UnregisterIBANOutput**](UnregisterIBANOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoneyOutsMoneyOutGet

> GetMoneyOutTransDetailsOutput MoneyOutsMoneyOutGet(ctx).Authorization(authorization).PSUIPAddress(pSUIPAddress).Transactionid(transactionid).TransactionComment(transactionComment).Reference(reference).Page(page).Limit(limit).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Search for a Money-Out



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
    transactionid := int64(789) // int64 | Money-Out ID (optional)
    transactionComment := "transactionComment_example" // string | Money-Out Comment (optional)
    reference := "reference_example" // string | Unique ID Generated by your Server (optional)
    page := int32(56) // int32 | Index start from 1, let null to get all (optional)
    limit := int32(56) // int32 | Default 20, let it null if no pagination (optional)
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyOutsApi.MoneyOutsMoneyOutGet(context.Background()).Authorization(authorization).PSUIPAddress(pSUIPAddress).Transactionid(transactionid).TransactionComment(transactionComment).Reference(reference).Page(page).Limit(limit).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyOutsApi.MoneyOutsMoneyOutGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyOutsMoneyOutGet`: GetMoneyOutTransDetailsOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyOutsApi.MoneyOutsMoneyOutGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoneyOutsMoneyOutGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **transactionid** | **int64** | Money-Out ID | 
 **transactionComment** | **string** | Money-Out Comment | 
 **reference** | **string** | Unique ID Generated by your Server | 
 **page** | **int32** | Index start from 1, let null to get all | 
 **limit** | **int32** | Default 20, let it null if no pagination | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**GetMoneyOutTransDetailsOutput**](GetMoneyOutTransDetailsOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoneyOutsMoneyOutPost

> MoneyOutOutput MoneyOutsMoneyOutPost(ctx).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

External Fund Transfer from a Payment Account to a Bank Account



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
    parameters := *openapiclient.NewMoneyOutInput("AccountId_example", false) // MoneyOutInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyOutsApi.MoneyOutsMoneyOutPost(context.Background()).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyOutsApi.MoneyOutsMoneyOutPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyOutsMoneyOutPost`: MoneyOutOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyOutsApi.MoneyOutsMoneyOutPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoneyOutsMoneyOutPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**MoneyOutInput**](MoneyOutInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**MoneyOutOutput**](MoneyOutOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

