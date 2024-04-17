# \MoneyInsBNPLApi

All URIs are relative to *https://sandbox-api.lemonway.fr/mb/drouot/dev/directkitrest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MoneyInsCreateBnplPayment**](MoneyInsBNPLApi.md#MoneyInsCreateBnplPayment) | **Post** /v2/moneyins/buynowpaylater/init | Create New Pending Payment
[**MoneyInsGetBnplPaymentPlans**](MoneyInsBNPLApi.md#MoneyInsGetBnplPaymentPlans) | **Get** /v2/moneyins/buynowpaylater/plans | Get all Actived Payment Plans



## MoneyInsCreateBnplPayment

> CreateBnplPaymentOutput MoneyInsCreateBnplPayment(ctx).Authorization(authorization).PSUIPAddress(pSUIPAddress).CreateBnplPaymentInput(createBnplPaymentInput).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Create New Pending Payment



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
    createBnplPaymentInput := *openapiclient.NewCreateBnplPaymentInput(int64(123), *openapiclient.NewBnplDelivery("DeliveryMode_example", false, *openapiclient.NewBnplAddress("Street_example", "City_example", "PostCode_example", "Country_example")), []openapiclient.BnplPurchaseItem{*openapiclient.NewBnplPurchaseItem("Type_example", "Description_example", "Reference_example", int32(123), int32(123), "MerchantAccountId_example")}, *openapiclient.NewCustomer("AccountId_example", "Civility_example", "FirstName_example", "LastName_example", "MobileNumber_example", "Email_example", *openapiclient.NewBnplAddress("Street_example", "City_example", "PostCode_example", "Country_example")), "ReturnUrl_example", "ErrorUrl_example", "CancelUrl_example", "AccountId_example") // CreateBnplPaymentInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyInsBNPLApi.MoneyInsCreateBnplPayment(context.Background()).Authorization(authorization).PSUIPAddress(pSUIPAddress).CreateBnplPaymentInput(createBnplPaymentInput).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyInsBNPLApi.MoneyInsCreateBnplPayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyInsCreateBnplPayment`: CreateBnplPaymentOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyInsBNPLApi.MoneyInsCreateBnplPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoneyInsCreateBnplPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **createBnplPaymentInput** | [**CreateBnplPaymentInput**](CreateBnplPaymentInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**CreateBnplPaymentOutput**](CreateBnplPaymentOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoneyInsGetBnplPaymentPlans

> GetPaymentPlansOutput MoneyInsGetBnplPaymentPlans(ctx).Authorization(authorization).PSUIPAddress(pSUIPAddress).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Get all Actived Payment Plans



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
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoneyInsBNPLApi.MoneyInsGetBnplPaymentPlans(context.Background()).Authorization(authorization).PSUIPAddress(pSUIPAddress).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoneyInsBNPLApi.MoneyInsGetBnplPaymentPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoneyInsGetBnplPaymentPlans`: GetPaymentPlansOutput
    fmt.Fprintf(os.Stdout, "Response from `MoneyInsBNPLApi.MoneyInsGetBnplPaymentPlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoneyInsGetBnplPaymentPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**GetPaymentPlansOutput**](GetPaymentPlansOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

