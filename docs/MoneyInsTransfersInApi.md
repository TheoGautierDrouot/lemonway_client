# \MoneyInsTransfersInApi

All URIs are relative to *https://sandbox-api.lemonway.fr/mb/drouot/dev/directkitrest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MoneyInsBankwireGet**](MoneyInsTransfersInApi.md#MoneyInsBankwireGet) | **Get** /v2/moneyins/bankwire | Search for a Money-In by Fund Transfer



## MoneyInsBankwireGet

> GetMoneyInIBANDetailsOutput MoneyInsBankwireGet(ctx).UpdateDate(updateDate).UpdateEndDate(updateEndDate).Authorization(authorization).PSUIPAddress(pSUIPAddress).Page(page).Limit(limit).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Search for a Money-In by Fund Transfer



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
    updateDate := "updateDate_example" // string | UTC Unix Timestamp
    updateEndDate := "updateEndDate_example" // string | End Date tUTC Unix Timestamp
    authorization := "authorization_example" // string | Authorization bearer (OAuth 2)
    pSUIPAddress := "pSUIPAddress_example" // string | IP address of the final client (PSU).
    page := int32(56) // int32 | Index start from 1, let null to get all (optional)
    limit := int32(56) // int32 | Default 20, let it null if no pagination (optional)
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyInsTransfersInApi.MoneyInsBankwireGet(context.Background()).UpdateDate(updateDate).UpdateEndDate(updateEndDate).Authorization(authorization).PSUIPAddress(pSUIPAddress).Page(page).Limit(limit).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyInsTransfersInApi.MoneyInsBankwireGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyInsBankwireGet`: GetMoneyInIBANDetailsOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyInsTransfersInApi.MoneyInsBankwireGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoneyInsBankwireGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDate** | **string** | UTC Unix Timestamp | 
 **updateEndDate** | **string** | End Date tUTC Unix Timestamp | 
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **page** | **int32** | Index start from 1, let null to get all | 
 **limit** | **int32** | Default 20, let it null if no pagination | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**GetMoneyInIBANDetailsOutput**](GetMoneyInIBANDetailsOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

