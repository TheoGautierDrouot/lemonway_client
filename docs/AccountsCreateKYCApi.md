# \AccountsCreateKYCApi

All URIs are relative to *https://sandbox-api.lemonway.fr/mb/drouot/dev/directkitrest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountsAddUltimateBeneficialOwner**](AccountsCreateKYCApi.md#AccountsAddUltimateBeneficialOwner) | **Post** /v2/accounts/{accountId}/ultimateBeneficialOwner | Create an Ultimate Beneficial Owner
[**AccountsDocumentsUploadPost**](AccountsCreateKYCApi.md#AccountsDocumentsUploadPost) | **Post** /v2/accounts/{accountid}/documents/upload | Upload Documents for KYC (Know Your Customers)
[**AccountsGetUltimateBeneficialOwner**](AccountsCreateKYCApi.md#AccountsGetUltimateBeneficialOwner) | **Get** /v2/accounts/{accountId}ultimateBeneficialOwner | Get all Ultimate Beneficial Owners associated to a payment account.
[**AccountsIndividualPost**](AccountsCreateKYCApi.md#AccountsIndividualPost) | **Post** /v2/accounts/individual | Create a New Individual Account
[**AccountsKycStatusGet**](AccountsCreateKYCApi.md#AccountsKycStatusGet) | **Get** /v2/accounts/kycstatus | Find a user, document or an IBAN that has been modified since an entry date
[**AccountsLegalPost**](AccountsCreateKYCApi.md#AccountsLegalPost) | **Post** /v2/accounts/legal | Create a New Legal Account
[**AccountsUpdateUltimateBeneficialOwner**](AccountsCreateKYCApi.md#AccountsUpdateUltimateBeneficialOwner) | **Put** /v2/accounts/{accountId}/ultimateBeneficialOwner/{UltimateBeneficialOwnerId} | Update Ultimate Beneficial Owner data



## AccountsAddUltimateBeneficialOwner

> AddUltimateBeneficialOwnerOutput AccountsAddUltimateBeneficialOwner(ctx, accountId).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Create an Ultimate Beneficial Owner



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
    parameters := *openapiclient.NewAddUltimateBeneficialOwnerInput("FirstName_example", "LastName_example", "DateOfBirth_example", "CityOfBirth_example", "CountryOfBirth_example") // AddUltimateBeneficialOwnerInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsCreateKYCApi.AccountsAddUltimateBeneficialOwner(context.Background(), accountId).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsCreateKYCApi.AccountsAddUltimateBeneficialOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsAddUltimateBeneficialOwner`: AddUltimateBeneficialOwnerOutput
    fmt.Fprintf(os.Stdout, "Response from `AccountsCreateKYCApi.AccountsAddUltimateBeneficialOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsAddUltimateBeneficialOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**AddUltimateBeneficialOwnerInput**](AddUltimateBeneficialOwnerInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**AddUltimateBeneficialOwnerOutput**](AddUltimateBeneficialOwnerOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsDocumentsUploadPost

> UploadDocumentOutput AccountsDocumentsUploadPost(ctx, accountid).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Upload Documents for KYC (Know Your Customers)



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
    parameters := *openapiclient.NewUploadDocumentInput("Name_example", int32(123), "Buffer_example") // UploadDocumentInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsCreateKYCApi.AccountsDocumentsUploadPost(context.Background(), accountid).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsCreateKYCApi.AccountsDocumentsUploadPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsDocumentsUploadPost`: UploadDocumentOutput
    fmt.Fprintf(os.Stdout, "Response from `AccountsCreateKYCApi.AccountsDocumentsUploadPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsDocumentsUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**UploadDocumentInput**](UploadDocumentInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**UploadDocumentOutput**](UploadDocumentOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsGetUltimateBeneficialOwner

> GetUltimateBeneficialOwnerOutput AccountsGetUltimateBeneficialOwner(ctx, accountId).Authorization(authorization).PSUIPAddress(pSUIPAddress).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Get all Ultimate Beneficial Owners associated to a payment account.



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
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsCreateKYCApi.AccountsGetUltimateBeneficialOwner(context.Background(), accountId).Authorization(authorization).PSUIPAddress(pSUIPAddress).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsCreateKYCApi.AccountsGetUltimateBeneficialOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsGetUltimateBeneficialOwner`: GetUltimateBeneficialOwnerOutput
    fmt.Fprintf(os.Stdout, "Response from `AccountsCreateKYCApi.AccountsGetUltimateBeneficialOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsGetUltimateBeneficialOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**GetUltimateBeneficialOwnerOutput**](GetUltimateBeneficialOwnerOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsIndividualPost

> RegisterIndividualAccountOutput AccountsIndividualPost(ctx).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Create a New Individual Account



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
    parameters := *openapiclient.NewRegisterIndividualAccountInput("AccountId_example", "Email_example", "FirstName_example", "LastName_example", "Nationality_example", int32(123)) // RegisterIndividualAccountInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsCreateKYCApi.AccountsIndividualPost(context.Background()).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsCreateKYCApi.AccountsIndividualPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsIndividualPost`: RegisterIndividualAccountOutput
    fmt.Fprintf(os.Stdout, "Response from `AccountsCreateKYCApi.AccountsIndividualPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountsIndividualPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**RegisterIndividualAccountInput**](RegisterIndividualAccountInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**RegisterIndividualAccountOutput**](RegisterIndividualAccountOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsKycStatusGet

> KycStatusOutput AccountsKycStatusGet(ctx).UpdateDate(updateDate).Authorization(authorization).PSUIPAddress(pSUIPAddress).Page(page).Limit(limit).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Find a user, document or an IBAN that has been modified since an entry date



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
    resp, r, err := apiClient.AccountsCreateKYCApi.AccountsKycStatusGet(context.Background()).UpdateDate(updateDate).Authorization(authorization).PSUIPAddress(pSUIPAddress).Page(page).Limit(limit).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsCreateKYCApi.AccountsKycStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsKycStatusGet`: KycStatusOutput
    fmt.Fprintf(os.Stdout, "Response from `AccountsCreateKYCApi.AccountsKycStatusGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountsKycStatusGetRequest struct via the builder pattern


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

[**KycStatusOutput**](KycStatusOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsLegalPost

> RegisterLegalAccountOutput AccountsLegalPost(ctx).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Create a New Legal Account



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
    parameters := *openapiclient.NewRegisterLegalAccountInput(*openapiclient.NewCompany("Name_example", "Description_example"), "AccountId_example", "Email_example", "FirstName_example", "LastName_example", "Nationality_example", int32(123)) // RegisterLegalAccountInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsCreateKYCApi.AccountsLegalPost(context.Background()).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsCreateKYCApi.AccountsLegalPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsLegalPost`: RegisterLegalAccountOutput
    fmt.Fprintf(os.Stdout, "Response from `AccountsCreateKYCApi.AccountsLegalPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountsLegalPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**RegisterLegalAccountInput**](RegisterLegalAccountInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**RegisterLegalAccountOutput**](RegisterLegalAccountOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsUpdateUltimateBeneficialOwner

> UpdateUltimateBeneficialOwnerOutput AccountsUpdateUltimateBeneficialOwner(ctx, accountId, ultimateBeneficialOwnerId).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()

Update Ultimate Beneficial Owner data



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
    ultimateBeneficialOwnerId := int64(789) // int64 | UltimateBeneficial Owner ID
    authorization := "authorization_example" // string | Authorization bearer (OAuth 2)
    pSUIPAddress := "pSUIPAddress_example" // string | IP address of the final client (PSU).
    parameters := *openapiclient.NewUpdateUltimateBeneficialOwnerInput("FirstName_example", "LastName_example", "DateOfBirth_example", "CityOfBirth_example", "CountryOfBirth_example") // UpdateUltimateBeneficialOwnerInput | 
    pSUAcceptLanguage := "pSUAcceptLanguage_example" // string | Response language accepted by final client (PSU). English by default (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | User-agent of the final client (PSU). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsCreateKYCApi.AccountsUpdateUltimateBeneficialOwner(context.Background(), accountId, ultimateBeneficialOwnerId).Authorization(authorization).PSUIPAddress(pSUIPAddress).Parameters(parameters).PSUAcceptLanguage(pSUAcceptLanguage).PSUUserAgent(pSUUserAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsCreateKYCApi.AccountsUpdateUltimateBeneficialOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsUpdateUltimateBeneficialOwner`: UpdateUltimateBeneficialOwnerOutput
    fmt.Fprintf(os.Stdout, "Response from `AccountsCreateKYCApi.AccountsUpdateUltimateBeneficialOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account ID | 
**ultimateBeneficialOwnerId** | **int64** | UltimateBeneficial Owner ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsUpdateUltimateBeneficialOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Authorization bearer (OAuth 2) | 
 **pSUIPAddress** | **string** | IP address of the final client (PSU). | 
 **parameters** | [**UpdateUltimateBeneficialOwnerInput**](UpdateUltimateBeneficialOwnerInput.md) |  | 
 **pSUAcceptLanguage** | **string** | Response language accepted by final client (PSU). English by default | 
 **pSUUserAgent** | **string** | User-agent of the final client (PSU). | 

### Return type

[**UpdateUltimateBeneficialOwnerOutput**](UpdateUltimateBeneficialOwnerOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

