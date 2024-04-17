# \AccountsUpdateApi

All URIs are relative to *https://sandbox-api.lemonway.fr/mb/drouot/dev/directkitrest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountsIndividualPut**](AccountsUpdateApi.md#AccountsIndividualPut) | **Put** /v2/accounts/individual/{accountid} | Update Individual Payment Account Data
[**AccountsKycstatusPut**](AccountsUpdateApi.md#AccountsKycstatusPut) | **Put** /v2/accounts/kycstatus/{accountid} | Update Payment Account Status
[**AccountsLegalSinglePut**](AccountsUpdateApi.md#AccountsLegalSinglePut) | **Put** /v2/accounts/legal/{accountid} | Update Legal Payment Account Data



## AccountsIndividualPut

> UpdateIndividualAccountDetailsOutput AccountsIndividualPut(ctx, accountid).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Update Individual Payment Account Data



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
    parameters := *openapiclient.NewUpdateIndividualAccountDetailsInput() // UpdateIndividualAccountDetailsInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsUpdateApi.AccountsIndividualPut(context.Background(), accountid).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsUpdateApi.AccountsIndividualPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsIndividualPut`: UpdateIndividualAccountDetailsOutput
    fmt.Fprintf(os.Stdout, "Response from `AccountsUpdateApi.AccountsIndividualPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsIndividualPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**UpdateIndividualAccountDetailsInput**](UpdateIndividualAccountDetailsInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**UpdateIndividualAccountDetailsOutput**](UpdateIndividualAccountDetailsOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsKycstatusPut

> UpdateAccountStatusOutput AccountsKycstatusPut(ctx, accountid).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Update Payment Account Status



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
    parameters := *openapiclient.NewUpdateAccountStatusInput(int32(123)) // UpdateAccountStatusInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsUpdateApi.AccountsKycstatusPut(context.Background(), accountid).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsUpdateApi.AccountsKycstatusPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsKycstatusPut`: UpdateAccountStatusOutput
    fmt.Fprintf(os.Stdout, "Response from `AccountsUpdateApi.AccountsKycstatusPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsKycstatusPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**UpdateAccountStatusInput**](UpdateAccountStatusInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**UpdateAccountStatusOutput**](UpdateAccountStatusOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsLegalSinglePut

> UpdateLegalAccountDetailsOutput AccountsLegalSinglePut(ctx, accountid).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Update Legal Payment Account Data



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
    parameters := *openapiclient.NewUpdateLegalAccountDetailsInput(*openapiclient.NewCompany("Name_example", "Description_example")) // UpdateLegalAccountDetailsInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsUpdateApi.AccountsLegalSinglePut(context.Background(), accountid).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsUpdateApi.AccountsLegalSinglePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsLegalSinglePut`: UpdateLegalAccountDetailsOutput
    fmt.Fprintf(os.Stdout, "Response from `AccountsUpdateApi.AccountsLegalSinglePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsLegalSinglePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**UpdateLegalAccountDetailsInput**](UpdateLegalAccountDetailsInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**UpdateLegalAccountDetailsOutput**](UpdateLegalAccountDetailsOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

