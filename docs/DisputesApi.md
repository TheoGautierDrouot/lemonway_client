# \DisputesApi

All URIs are relative to *https://sandbox-api.lemonway.fr/mb/drouot/dev/directkitrest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisputesDisputesGet**](DisputesApi.md#DisputesDisputesGet) | **Get** /v2/disputes | Get list of disputes from a given date



## DisputesDisputesGet

> GetChargebacksOutput DisputesDisputesGet(ctx).UpdateDate(updateDate).Authorization(authorization).PSUIPAddress(pSUIPAddress).Page(page).Limit(limit).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Get list of disputes from a given date

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
    updateDate := "updateDate_example" // string | UTC Unix timestamp
    authorization := "authorization_example" // string | Authorization bearer (OAuth 2)
    pSUIPAddress := "pSUIPAddress_example" // string | IP address of the final client (PSU).
    page := int32(56) // int32 | Index start from 1, let null to get all (optional)
    limit := int32(56) // int32 | Default 20, let it null if no pagination (optional)
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DisputesApi.DisputesDisputesGet(context.Background()).UpdateDate(updateDate).Authorization(authorization).PSUIPAddress(pSUIPAddress).Page(page).Limit(limit).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DisputesApi.DisputesDisputesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisputesDisputesGet`: GetChargebacksOutput
    fmt.Fprintf(os.Stdout, "Response from `DisputesApi.DisputesDisputesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisputesDisputesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDate** | **string** | UTC Unix timestamp | 
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **page** | **int32** | Index start from 1, let null to get all | 
 **limit** | **int32** | Default 20, let it null if no pagination | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**GetChargebacksOutput**](GetChargebacksOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

