# \MoneyInsPayByBankApi

All URIs are relative to *https://sandbox-api.lemonway.fr/mb/drouot/dev/directkitrest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MoneyInsGetMoneyInBanks**](MoneyInsPayByBankApi.md#MoneyInsGetMoneyInBanks) | **Get** /v2/moneyins/paybybank/transfer/banks | Get Pay by Bank List
[**MoneyInsMoneyInTransferInit**](MoneyInsPayByBankApi.md#MoneyInsMoneyInTransferInit) | **Post** /v2/moneyins/paybybank/transfer/init | Initiate Pay by Bank



## MoneyInsGetMoneyInBanks

> GetMoneyInBanksOutput MoneyInsGetMoneyInBanks(ctx).Authorization(authorization).PSUIPAddress(pSUIPAddress).CountryCodes(countryCodes).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Get Pay by Bank List



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
    countryCodes := "countryCodes_example" // string | Comma separated string of ISO Alpha-2 country codes.  Available country codes include:    - France (FR)     - Spain (ES)     - Italy (IT)     - Germany (DE)     - Portugal (PT) (optional)
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyInsPayByBankApi.MoneyInsGetMoneyInBanks(context.Background()).Authorization(authorization).PSUIPAddress(pSUIPAddress).CountryCodes(countryCodes).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyInsPayByBankApi.MoneyInsGetMoneyInBanks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyInsGetMoneyInBanks`: GetMoneyInBanksOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyInsPayByBankApi.MoneyInsGetMoneyInBanks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoneyInsGetMoneyInBanksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **countryCodes** | **string** | Comma separated string of ISO Alpha-2 country codes.  Available country codes include:    - France (FR)     - Spain (ES)     - Italy (IT)     - Germany (DE)     - Portugal (PT) | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**GetMoneyInBanksOutput**](GetMoneyInBanksOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoneyInsMoneyInTransferInit

> MoneyInTransferInitOutput MoneyInsMoneyInTransferInit(ctx).Authorization(authorization).PSUIPAddress(pSUIPAddress).RawInput(rawInput).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Initiate Pay by Bank



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
    rawInput := *openapiclient.NewMoneyInTransferInitInput("ReturnUrl_example", "CancelUrl_example", "ErrorUrl_example", "CountryCode_example", "AccountId_example") // MoneyInTransferInitInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyInsPayByBankApi.MoneyInsMoneyInTransferInit(context.Background()).Authorization(authorization).PSUIPAddress(pSUIPAddress).RawInput(rawInput).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyInsPayByBankApi.MoneyInsMoneyInTransferInit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyInsMoneyInTransferInit`: MoneyInTransferInitOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyInsPayByBankApi.MoneyInsMoneyInTransferInit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoneyInsMoneyInTransferInitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **rawInput** | [**MoneyInTransferInitInput**](MoneyInTransferInitInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**MoneyInTransferInitOutput**](MoneyInTransferInitOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

