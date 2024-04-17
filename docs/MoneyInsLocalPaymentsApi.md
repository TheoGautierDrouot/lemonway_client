# \MoneyInsLocalPaymentsApi

All URIs are relative to *https://sandbox-api.lemonway.fr/mb/drouot/dev/directkitrest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MoneyInsIDealConfirmPut**](MoneyInsLocalPaymentsApi.md#MoneyInsIDealConfirmPut) | **Put** /v2/moneyins/ideal/{transactionid}/confirm | Finalize an iDeal Payment
[**MoneyInsIDealInitPost**](MoneyInsLocalPaymentsApi.md#MoneyInsIDealInitPost) | **Post** /v2/moneyins/ideal/init | Initialize iDeal Payment
[**MoneyInsMbwayInitPost**](MoneyInsLocalPaymentsApi.md#MoneyInsMbwayInitPost) | **Post** /v2/moneyins/mbway/init | Initialize MB WAY Payment
[**MoneyInsMobilePayInitPost**](MoneyInsLocalPaymentsApi.md#MoneyInsMobilePayInitPost) | **Post** /v2/moneyins/mobilePay/init | Initialize MobilePay Payment
[**MoneyInsMultibancoInitPost**](MoneyInsLocalPaymentsApi.md#MoneyInsMultibancoInitPost) | **Post** /v2/moneyins/multibanco/init | Initialize Multibanco Payment
[**MoneyInsPayTrailInitPost**](MoneyInsLocalPaymentsApi.md#MoneyInsPayTrailInitPost) | **Post** /v2/moneyins/paytrail/init | Initialize PayTrail Payment (Deprecated)
[**MoneyInsPayshopInitPost**](MoneyInsLocalPaymentsApi.md#MoneyInsPayshopInitPost) | **Post** /v2/moneyins/payshop/init | Initialize Payshop Payment
[**MoneyInsSofortInitPost**](MoneyInsLocalPaymentsApi.md#MoneyInsSofortInitPost) | **Post** /v2/moneyins/sofort/init | Initialize Sofort Payment
[**MoneyInsTrustlyInitPost**](MoneyInsLocalPaymentsApi.md#MoneyInsTrustlyInitPost) | **Post** /v2/moneyins/trustly/init | Initialize Trustly Payment



## MoneyInsIDealConfirmPut

> MoneyInIDealConfirmOutput MoneyInsIDealConfirmPut(ctx, transactionid).Authorization(authorization).PSUIPAddress(pSUIPAddress).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Finalize an iDeal Payment



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
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyInsLocalPaymentsApi.MoneyInsIDealConfirmPut(context.Background(), transactionid).Authorization(authorization).PSUIPAddress(pSUIPAddress).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyInsLocalPaymentsApi.MoneyInsIDealConfirmPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyInsIDealConfirmPut`: MoneyInIDealConfirmOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyInsLocalPaymentsApi.MoneyInsIDealConfirmPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionid** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoneyInsIDealConfirmPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**MoneyInIDealConfirmOutput**](MoneyInIDealConfirmOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoneyInsIDealInitPost

> MoneyInIDealInitOutput MoneyInsIDealInitPost(ctx).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Initialize iDeal Payment



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
    parameters := *openapiclient.NewMoneyInIDealInitInput("ReturnUrl_example", "IssuerId_example", "AccountId_example") // MoneyInIDealInitInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyInsLocalPaymentsApi.MoneyInsIDealInitPost(context.Background()).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyInsLocalPaymentsApi.MoneyInsIDealInitPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyInsIDealInitPost`: MoneyInIDealInitOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyInsLocalPaymentsApi.MoneyInsIDealInitPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoneyInsIDealInitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**MoneyInIDealInitInput**](MoneyInIDealInitInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**MoneyInIDealInitOutput**](MoneyInIDealInitOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoneyInsMbwayInitPost

> MoneyInMbwayInitOutput MoneyInsMbwayInitPost(ctx).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Initialize MB WAY Payment



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
    parameters := *openapiclient.NewMoneyInMbwayInitInput("Phone_example", "ReturnUrl_example", "AccountId_example") // MoneyInMbwayInitInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyInsLocalPaymentsApi.MoneyInsMbwayInitPost(context.Background()).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyInsLocalPaymentsApi.MoneyInsMbwayInitPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyInsMbwayInitPost`: MoneyInMbwayInitOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyInsLocalPaymentsApi.MoneyInsMbwayInitPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoneyInsMbwayInitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**MoneyInMbwayInitInput**](MoneyInMbwayInitInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**MoneyInMbwayInitOutput**](MoneyInMbwayInitOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoneyInsMobilePayInitPost

> MoneyInMobilePayInitOutput MoneyInsMobilePayInitPost(ctx).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Initialize MobilePay Payment



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
    parameters := *openapiclient.NewMoneyInMobilePayInitInput("ReturnUrl_example", "AccountId_example") // MoneyInMobilePayInitInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyInsLocalPaymentsApi.MoneyInsMobilePayInitPost(context.Background()).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyInsLocalPaymentsApi.MoneyInsMobilePayInitPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyInsMobilePayInitPost`: MoneyInMobilePayInitOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyInsLocalPaymentsApi.MoneyInsMobilePayInitPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoneyInsMobilePayInitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**MoneyInMobilePayInitInput**](MoneyInMobilePayInitInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**MoneyInMobilePayInitOutput**](MoneyInMobilePayInitOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoneyInsMultibancoInitPost

> MoneyInMultibancoInitOutput MoneyInsMultibancoInitPost(ctx).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Initialize Multibanco Payment



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
    parameters := *openapiclient.NewMoneyInMultibancoInitInput("ReturnUrl_example", "AccountId_example") // MoneyInMultibancoInitInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyInsLocalPaymentsApi.MoneyInsMultibancoInitPost(context.Background()).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyInsLocalPaymentsApi.MoneyInsMultibancoInitPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyInsMultibancoInitPost`: MoneyInMultibancoInitOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyInsLocalPaymentsApi.MoneyInsMultibancoInitPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoneyInsMultibancoInitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**MoneyInMultibancoInitInput**](MoneyInMultibancoInitInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**MoneyInMultibancoInitOutput**](MoneyInMultibancoInitOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoneyInsPayTrailInitPost

> MoneyInPayTrailInitOutput MoneyInsPayTrailInitPost(ctx).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Initialize PayTrail Payment (Deprecated)



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
    parameters := *openapiclient.NewMoneyInPayTrailInitInput("ReturnUrl_example", "AccountId_example") // MoneyInPayTrailInitInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyInsLocalPaymentsApi.MoneyInsPayTrailInitPost(context.Background()).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyInsLocalPaymentsApi.MoneyInsPayTrailInitPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyInsPayTrailInitPost`: MoneyInPayTrailInitOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyInsLocalPaymentsApi.MoneyInsPayTrailInitPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoneyInsPayTrailInitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**MoneyInPayTrailInitInput**](MoneyInPayTrailInitInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**MoneyInPayTrailInitOutput**](MoneyInPayTrailInitOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoneyInsPayshopInitPost

> MoneyInPayshopInitOutput MoneyInsPayshopInitPost(ctx).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Initialize Payshop Payment



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
    parameters := *openapiclient.NewMoneyInPayshopInitInput("ReturnUrl_example", "AccountId_example") // MoneyInPayshopInitInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyInsLocalPaymentsApi.MoneyInsPayshopInitPost(context.Background()).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyInsLocalPaymentsApi.MoneyInsPayshopInitPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyInsPayshopInitPost`: MoneyInPayshopInitOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyInsLocalPaymentsApi.MoneyInsPayshopInitPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoneyInsPayshopInitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**MoneyInPayshopInitInput**](MoneyInPayshopInitInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**MoneyInPayshopInitOutput**](MoneyInPayshopInitOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoneyInsSofortInitPost

> MoneyInSofortInitOutput MoneyInsSofortInitPost(ctx).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Initialize Sofort Payment



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
    parameters := *openapiclient.NewMoneyInSofortInitInput("ReturnUrl_example", "AccountId_example") // MoneyInSofortInitInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyInsLocalPaymentsApi.MoneyInsSofortInitPost(context.Background()).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyInsLocalPaymentsApi.MoneyInsSofortInitPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyInsSofortInitPost`: MoneyInSofortInitOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyInsLocalPaymentsApi.MoneyInsSofortInitPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoneyInsSofortInitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**MoneyInSofortInitInput**](MoneyInSofortInitInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**MoneyInSofortInitOutput**](MoneyInSofortInitOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoneyInsTrustlyInitPost

> MoneyInTrustlyInitOutput MoneyInsTrustlyInitPost(ctx).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Initialize Trustly Payment



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
    parameters := *openapiclient.NewMoneyInTrustlyInitInput("FinalCustomerEmail_example", "FinalCustomerFirstName_example", "FinalCustomerLastName_example", "ReturnUrl_example", "AccountId_example") // MoneyInTrustlyInitInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyInsLocalPaymentsApi.MoneyInsTrustlyInitPost(context.Background()).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyInsLocalPaymentsApi.MoneyInsTrustlyInitPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyInsTrustlyInitPost`: MoneyInTrustlyInitOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyInsLocalPaymentsApi.MoneyInsTrustlyInitPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoneyInsTrustlyInitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**MoneyInTrustlyInitInput**](MoneyInTrustlyInitInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**MoneyInTrustlyInitOutput**](MoneyInTrustlyInitOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

