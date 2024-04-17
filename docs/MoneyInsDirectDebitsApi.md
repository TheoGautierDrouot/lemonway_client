# \MoneyInsDirectDebitsApi

All URIs are relative to *https://sandbox-api.lemonway.fr/mb/drouot/dev/directkitrest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MoneyInsCancelPut**](MoneyInsDirectDebitsApi.md#MoneyInsCancelPut) | **Put** /v2/moneyins/{transactionid}/cancel | Cancel a Money-In
[**MoneyInsMandateGet**](MoneyInsDirectDebitsApi.md#MoneyInsMandateGet) | **Get** /v2/moneyins/{accountid}/mandate | Get Mandate Associated to a Payment Account
[**MoneyInsMandateGetDocument**](MoneyInsDirectDebitsApi.md#MoneyInsMandateGetDocument) | **Get** /v2/moneyins/{accountid}/mandate/{mandateid}/document | Get Mandate Document
[**MoneyInsSddGet**](MoneyInsDirectDebitsApi.md#MoneyInsSddGet) | **Get** /v2/moneyins/sdd | List of Money-In by SEPA Direct Debit (SDD)
[**MoneyInsSddInitPost**](MoneyInsDirectDebitsApi.md#MoneyInsSddInitPost) | **Post** /v2/moneyins/sdd/init | Request a SEPA Direct Debit (SDD)
[**MoneyInsSddMandatePost**](MoneyInsDirectDebitsApi.md#MoneyInsSddMandatePost) | **Post** /v2/moneyins/sdd/mandate | Register a SDD Mandate
[**MoneyInsSddMandateUnregisterPut**](MoneyInsDirectDebitsApi.md#MoneyInsSddMandateUnregisterPut) | **Put** /v2/moneyins/sdd/mandate/{mandatid}/unregister | Deactivate a Mandate



## MoneyInsCancelPut

> CancelMoneyInOutput MoneyInsCancelPut(ctx, transactionid).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Cancel a Money-In



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
    parameters := *openapiclient.NewCancelMoneyInInput("Account_example") // CancelMoneyInInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyInsDirectDebitsApi.MoneyInsCancelPut(context.Background(), transactionid).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyInsDirectDebitsApi.MoneyInsCancelPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyInsCancelPut`: CancelMoneyInOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyInsDirectDebitsApi.MoneyInsCancelPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionid** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoneyInsCancelPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**CancelMoneyInInput**](CancelMoneyInInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**CancelMoneyInOutput**](CancelMoneyInOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoneyInsMandateGet

> AccountMandatsOutput MoneyInsMandateGet(ctx, accountid).Authorization(authorization).PSUIPAddress(pSUIPAddress).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Get Mandate Associated to a Payment Account

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
    resp, r, err := apiClient.MoneyInsDirectDebitsApi.MoneyInsMandateGet(context.Background(), accountid).Authorization(authorization).PSUIPAddress(pSUIPAddress).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyInsDirectDebitsApi.MoneyInsMandateGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyInsMandateGet`: AccountMandatsOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyInsDirectDebitsApi.MoneyInsMandateGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoneyInsMandateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**AccountMandatsOutput**](AccountMandatsOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoneyInsMandateGetDocument

> MandateGetDocumentOutput MoneyInsMandateGetDocument(ctx, accountid, mandateid).Authorization(authorization).PSUIPAddress(pSUIPAddress).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Get Mandate Document



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
    accountid := "accountid_example" // string | Account Id
    mandateid := int64(789) // int64 | Id of the registered SDD Mandate
    authorization := "authorization_example" // string | Authorization bearer (OAuth 2)
    pSUIPAddress := "pSUIPAddress_example" // string | IP address of the final client (PSU).
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyInsDirectDebitsApi.MoneyInsMandateGetDocument(context.Background(), accountid, mandateid).Authorization(authorization).PSUIPAddress(pSUIPAddress).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyInsDirectDebitsApi.MoneyInsMandateGetDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyInsMandateGetDocument`: MandateGetDocumentOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyInsDirectDebitsApi.MoneyInsMandateGetDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account Id | 
**mandateid** | **int64** | Id of the registered SDD Mandate | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoneyInsMandateGetDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**MandateGetDocumentOutput**](MandateGetDocumentOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoneyInsSddGet

> GetMoneyInSddOutput MoneyInsSddGet(ctx).Authorization(authorization).PSUIPAddress(pSUIPAddress).Reference(reference).UpdateDate(updateDate).UpdateEndDate(updateEndDate).Page(page).Limit(limit).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

List of Money-In by SEPA Direct Debit (SDD)



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
    reference := "reference_example" // string | Unique ID Generated by your Server (optional)
    updateDate := "updateDate_example" // string | UTC Unix Timestamp (optional)
    updateEndDate := "updateEndDate_example" // string | End Date UTC Unix Timestamp (optional)
    page := int32(56) // int32 | Index start from 1, let null to get all (optional)
    limit := int32(56) // int32 | Default 20, let it null if no pagination (optional)
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyInsDirectDebitsApi.MoneyInsSddGet(context.Background()).Authorization(authorization).PSUIPAddress(pSUIPAddress).Reference(reference).UpdateDate(updateDate).UpdateEndDate(updateEndDate).Page(page).Limit(limit).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyInsDirectDebitsApi.MoneyInsSddGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyInsSddGet`: GetMoneyInSddOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyInsDirectDebitsApi.MoneyInsSddGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoneyInsSddGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **reference** | **string** | Unique ID Generated by your Server | 
 **updateDate** | **string** | UTC Unix Timestamp | 
 **updateEndDate** | **string** | End Date UTC Unix Timestamp | 
 **page** | **int32** | Index start from 1, let null to get all | 
 **limit** | **int32** | Default 20, let it null if no pagination | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**GetMoneyInSddOutput**](GetMoneyInSddOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoneyInsSddInitPost

> MoneyInSddInitOutput MoneyInsSddInitPost(ctx).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Request a SEPA Direct Debit (SDD)



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
    parameters := *openapiclient.NewMoneyInSddInitInput(int64(123), "AccountId_example") // MoneyInSddInitInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyInsDirectDebitsApi.MoneyInsSddInitPost(context.Background()).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyInsDirectDebitsApi.MoneyInsSddInitPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyInsSddInitPost`: MoneyInSddInitOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyInsDirectDebitsApi.MoneyInsSddInitPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoneyInsSddInitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**MoneyInSddInitInput**](MoneyInSddInitInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**MoneyInSddInitOutput**](MoneyInSddInitOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoneyInsSddMandatePost

> RegisterSddMandateOutput MoneyInsSddMandatePost(ctx).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Register a SDD Mandate



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
    parameters := *openapiclient.NewRegisterSddMandateInput("Wallet_example", "Holder_example", "Bic_example", "Iban_example", false) // RegisterSddMandateInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyInsDirectDebitsApi.MoneyInsSddMandatePost(context.Background()).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyInsDirectDebitsApi.MoneyInsSddMandatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyInsSddMandatePost`: RegisterSddMandateOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyInsDirectDebitsApi.MoneyInsSddMandatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoneyInsSddMandatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**RegisterSddMandateInput**](RegisterSddMandateInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**RegisterSddMandateOutput**](RegisterSddMandateOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoneyInsSddMandateUnregisterPut

> UnregisterSddMandateOutput MoneyInsSddMandateUnregisterPut(ctx, mandatid).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Deactivate a Mandate



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
    mandatid := int64(789) // int64 | 
    authorization := "authorization_example" // string | Authorization bearer (OAuth 2)
    pSUIPAddress := "pSUIPAddress_example" // string | IP address of the final client (PSU).
    parameters := *openapiclient.NewUnregisterSddMandateInput("Account_example") // UnregisterSddMandateInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyInsDirectDebitsApi.MoneyInsSddMandateUnregisterPut(context.Background(), mandatid).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyInsDirectDebitsApi.MoneyInsSddMandateUnregisterPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyInsSddMandateUnregisterPut`: UnregisterSddMandateOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyInsDirectDebitsApi.MoneyInsSddMandateUnregisterPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mandatid** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoneyInsSddMandateUnregisterPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**UnregisterSddMandateInput**](UnregisterSddMandateInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**UnregisterSddMandateOutput**](UnregisterSddMandateOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

