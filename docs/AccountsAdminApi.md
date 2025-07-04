# \AccountsAdminApi

All URIs are relative to *https://sandbox-api.lemonway.fr/mb/drouot/dev/directkitrest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountsAccountSingleGet**](AccountsAdminApi.md#AccountsAccountSingleGet) | **Get** /v2/accounts/{accountid} | Get Detailed Payment Account Data
[**AccountsBalancesGet**](AccountsAdminApi.md#AccountsBalancesGet) | **Get** /v2/accounts/balances | Get all Payment Account Balances
[**AccountsBalancesHistoryGet**](AccountsAdminApi.md#AccountsBalancesHistoryGet) | **Get** /v2/accounts/{accountId}/balances/history | Get Payment Account Balance History
[**AccountsBlockedPut**](AccountsAdminApi.md#AccountsBlockedPut) | **Put** /v2/accounts/{accountid}/blocked | Block or Unblock an Account
[**AccountsDocumentGet**](AccountsAdminApi.md#AccountsDocumentGet) | **Get** /v2/accounts/{accountid}/documents | Get Documents Associated with a Payment Account
[**AccountsDocumentsSignInitPost**](AccountsAdminApi.md#AccountsDocumentsSignInitPost) | **Post** /v2/accounts/{accountid}/documents/{documentid}/signinit | Generate an Electronic Signature of a Document
[**AccountsEnrolmentInit**](AccountsAdminApi.md#AccountsEnrolmentInit) | **Post** /v2/accounts/{accountid}/enrolment/init | (Deprecated) Initialize a Deutsche Post POSTIDENT Identification
[**AccountsRetrievePost**](AccountsAdminApi.md#AccountsRetrievePost) | **Post** /v2/accounts/retrieve | Get Detailed Payments Accounts Data
[**AccountsTransactionsGet**](AccountsAdminApi.md#AccountsTransactionsGet) | **Get** /v2/accounts/{accountId}/transactions | Get list of all Payment Account transactions



## AccountsAccountSingleGet

> AccountDetailsOutput AccountsAccountSingleGet(ctx, accountid).Authorization(authorization).PSUIPAddress(pSUIPAddress).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Get Detailed Payment Account Data



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
    resp, r, err := apiClient.AccountsAdminApi.AccountsAccountSingleGet(context.Background(), accountid).Authorization(authorization).PSUIPAddress(pSUIPAddress).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAdminApi.AccountsAccountSingleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsAccountSingleGet`: AccountDetailsOutput
    fmt.Fprintf(os.Stdout, "Response from `AccountsAdminApi.AccountsAccountSingleGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsAccountSingleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**AccountDetailsOutput**](AccountDetailsOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsBalancesGet

> AccountBalanceOutput AccountsBalancesGet(ctx).Authorization(authorization).PSUIPAddress(pSUIPAddress).UpdateDate(updateDate).InternalAccountIdStart(internalAccountIdStart).InternalAccountIdEnd(internalAccountIdEnd).Page(page).Limit(limit).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Get all Payment Account Balances



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
    updateDate := "updateDate_example" // string | Date in UTC Seconds  Leave empty to use payment account IDs. (optional)
    internalAccountIdStart := int64(789) // int64 | First payment account internal ID, starting from 12. (optional)
    internalAccountIdEnd := int64(789) // int64 | Last payment account internal ID, starting from 12. (optional)
    page := int32(56) // int32 | Index start from 1, let null to get all (optional)
    limit := int32(56) // int32 | Default 20, let it null if no pagination (optional)
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsAdminApi.AccountsBalancesGet(context.Background()).Authorization(authorization).PSUIPAddress(pSUIPAddress).UpdateDate(updateDate).InternalAccountIdStart(internalAccountIdStart).InternalAccountIdEnd(internalAccountIdEnd).Page(page).Limit(limit).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAdminApi.AccountsBalancesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsBalancesGet`: AccountBalanceOutput
    fmt.Fprintf(os.Stdout, "Response from `AccountsAdminApi.AccountsBalancesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountsBalancesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **updateDate** | **string** | Date in UTC Seconds  Leave empty to use payment account IDs. | 
 **internalAccountIdStart** | **int64** | First payment account internal ID, starting from 12. | 
 **internalAccountIdEnd** | **int64** | Last payment account internal ID, starting from 12. | 
 **page** | **int32** | Index start from 1, let null to get all | 
 **limit** | **int32** | Default 20, let it null if no pagination | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**AccountBalanceOutput**](AccountBalanceOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsBalancesHistoryGet

> BalanceHistoryOutput AccountsBalancesHistoryGet(ctx, accountId).Authorization(authorization).PSUIPAddress(pSUIPAddress).AtDate(atDate).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Get Payment Account Balance History



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
    accountId := "accountId_example" // string | Account ID
    authorization := "authorization_example" // string | Authorization bearer (OAuth 2)
    pSUIPAddress := "pSUIPAddress_example" // string | IP address of the final client (PSU).
    atDate := "atDate_example" // string | Request balance at given time in UTC Unix timestamp. (optional)
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsAdminApi.AccountsBalancesHistoryGet(context.Background(), accountId).Authorization(authorization).PSUIPAddress(pSUIPAddress).AtDate(atDate).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAdminApi.AccountsBalancesHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsBalancesHistoryGet`: BalanceHistoryOutput
    fmt.Fprintf(os.Stdout, "Response from `AccountsAdminApi.AccountsBalancesHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsBalancesHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **atDate** | **string** | Request balance at given time in UTC Unix timestamp. | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**BalanceHistoryOutput**](BalanceHistoryOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsBlockedPut

> AccountBlockedOutput AccountsBlockedPut(ctx, accountid).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Block or Unblock an Account

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
    parameters := *openapiclient.NewAccountBlockedInput(false) // AccountBlockedInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsAdminApi.AccountsBlockedPut(context.Background(), accountid).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAdminApi.AccountsBlockedPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsBlockedPut`: AccountBlockedOutput
    fmt.Fprintf(os.Stdout, "Response from `AccountsAdminApi.AccountsBlockedPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsBlockedPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**AccountBlockedInput**](AccountBlockedInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**AccountBlockedOutput**](AccountBlockedOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsDocumentGet

> AccountDocumentsOutput AccountsDocumentGet(ctx, accountid).Authorization(authorization).PSUIPAddress(pSUIPAddress).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Get Documents Associated with a Payment Account



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
    resp, r, err := apiClient.AccountsAdminApi.AccountsDocumentGet(context.Background(), accountid).Authorization(authorization).PSUIPAddress(pSUIPAddress).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAdminApi.AccountsDocumentGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsDocumentGet`: AccountDocumentsOutput
    fmt.Fprintf(os.Stdout, "Response from `AccountsAdminApi.AccountsDocumentGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsDocumentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**AccountDocumentsOutput**](AccountDocumentsOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsDocumentsSignInitPost

> SignDocumentInitOutput AccountsDocumentsSignInitPost(ctx, accountid, documentid).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Generate an Electronic Signature of a Document



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
    documentid := int64(789) // int64 | 
    authorization := "authorization_example" // string | Authorization bearer (OAuth 2)
    pSUIPAddress := "pSUIPAddress_example" // string | IP address of the final client (PSU).
    parameters := *openapiclient.NewSignDocumentInitInput(int32(123), "ReturnUrl_example", "ErrorUrl_example") // SignDocumentInitInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsAdminApi.AccountsDocumentsSignInitPost(context.Background(), accountid, documentid).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAdminApi.AccountsDocumentsSignInitPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsDocumentsSignInitPost`: SignDocumentInitOutput
    fmt.Fprintf(os.Stdout, "Response from `AccountsAdminApi.AccountsDocumentsSignInitPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account Id | 
**documentid** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsDocumentsSignInitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**SignDocumentInitInput**](SignDocumentInitInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**SignDocumentInitOutput**](SignDocumentInitOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsEnrolmentInit

> EnrolmentInitOutput AccountsEnrolmentInit(ctx, accountid).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

(Deprecated) Initialize a Deutsche Post POSTIDENT Identification

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
    accountid := "accountid_example" // string | 
    authorization := "authorization_example" // string | Authorization bearer (OAuth 2)
    pSUIPAddress := "pSUIPAddress_example" // string | IP address of the final client (PSU).
    parameters := *openapiclient.NewEnrolmentInitInput(int32(123)) // EnrolmentInitInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsAdminApi.AccountsEnrolmentInit(context.Background(), accountid).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAdminApi.AccountsEnrolmentInit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsEnrolmentInit`: EnrolmentInitOutput
    fmt.Fprintf(os.Stdout, "Response from `AccountsAdminApi.AccountsEnrolmentInit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsEnrolmentInitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**EnrolmentInitInput**](EnrolmentInitInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**EnrolmentInitOutput**](EnrolmentInitOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsRetrievePost

> AccountDetailsBatchOutput AccountsRetrievePost(ctx).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).Page(page).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Get Detailed Payments Accounts Data



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
    parameters := *openapiclient.NewAccountDetailsBatchInput([]openapiclient.WalletDetailsInput{*openapiclient.NewWalletDetailsInput()}) // AccountDetailsBatchInput | 
    page := int32(56) // int32 |  (optional)
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsAdminApi.AccountsRetrievePost(context.Background()).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).Page(page).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAdminApi.AccountsRetrievePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsRetrievePost`: AccountDetailsBatchOutput
    fmt.Fprintf(os.Stdout, "Response from `AccountsAdminApi.AccountsRetrievePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountsRetrievePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**AccountDetailsBatchInput**](AccountDetailsBatchInput.md) |  | 
 **page** | **int32** |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**AccountDetailsBatchOutput**](AccountDetailsBatchOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsTransactionsGet

> AccountKycStatusOutput AccountsTransactionsGet(ctx, accountId).Authorization(authorization).PSUIPAddress(pSUIPAddress).StartDate(startDate).EndDate(endDate).ExecutionDateStart(executionDateStart).ExecutionDateEnd(executionDateEnd).Page(page).Limit(limit).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Get list of all Payment Account transactions



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
    accountId := "accountId_example" // string | Account ID
    authorization := "authorization_example" // string | Authorization bearer (OAuth 2)
    pSUIPAddress := "pSUIPAddress_example" // string | IP address of the final client (PSU).
    startDate := "startDate_example" // string | UTC Unix timestamp  In order to return transactions initialized after `startDate`.  If the payment account is a **SC Wallet** then this value is mandatory. (optional)
    endDate := "endDate_example" // string | UTC Unix timestamp  In order to return transactions initialized before `endDate`.  If the payment account is a **SC Wallet** then this value is mandatory and the time span can not exceed 1 week. (optional)
    executionDateStart := "executionDateStart_example" // string | UTC Unix timestamp  In order to return transactions Executed after `endDate`.  If the payment account is a **SC Wallet** then this value is mandatory and the time span can not exceed 1 week. (optional)
    executionDateEnd := "executionDateEnd_example" // string | UTC Unix timestamp  In order to return transactions Executed before `endDate`  If the payment account is a **SC Wallet** then this value is mandatory and the time span can not exceed 1 week. (optional)
    page := int32(56) // int32 | Index start from 1, let null to get all (optional)
    limit := int32(56) // int32 | Default 20, let it null if no pagination (optional)
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsAdminApi.AccountsTransactionsGet(context.Background(), accountId).Authorization(authorization).PSUIPAddress(pSUIPAddress).StartDate(startDate).EndDate(endDate).ExecutionDateStart(executionDateStart).ExecutionDateEnd(executionDateEnd).Page(page).Limit(limit).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAdminApi.AccountsTransactionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsTransactionsGet`: AccountKycStatusOutput
    fmt.Fprintf(os.Stdout, "Response from `AccountsAdminApi.AccountsTransactionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **startDate** | **string** | UTC Unix timestamp  In order to return transactions initialized after &#x60;startDate&#x60;.  If the payment account is a **SC Wallet** then this value is mandatory. | 
 **endDate** | **string** | UTC Unix timestamp  In order to return transactions initialized before &#x60;endDate&#x60;.  If the payment account is a **SC Wallet** then this value is mandatory and the time span can not exceed 1 week. | 
 **executionDateStart** | **string** | UTC Unix timestamp  In order to return transactions Executed after &#x60;endDate&#x60;.  If the payment account is a **SC Wallet** then this value is mandatory and the time span can not exceed 1 week. | 
 **executionDateEnd** | **string** | UTC Unix timestamp  In order to return transactions Executed before &#x60;endDate&#x60;  If the payment account is a **SC Wallet** then this value is mandatory and the time span can not exceed 1 week. | 
 **page** | **int32** | Index start from 1, let null to get all | 
 **limit** | **int32** | Default 20, let it null if no pagination | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**AccountKycStatusOutput**](AccountKycStatusOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

