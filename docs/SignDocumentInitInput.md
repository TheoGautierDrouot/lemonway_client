# SignDocumentInitInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MobileNumber** | Pointer to **string** | Required if no mobile number is already assigned to the payment account (using RegisterWallet or UpdateWalletDetails).  Format must be MSISDN : international number with country code, without \&quot;+\&quot; and \&quot;00\&quot;. | [optional] 
**Type** | **int32** | Use **21** for SDD mandate | 
**ReturnUrl** | **string** | Your website will the return URL, called by WEBKIT to terminate the operation. | 
**ErrorUrl** | **string** | You site will return the URL, called by WEBKIT in case of an error. | 

## Methods

### NewSignDocumentInitInput

`func NewSignDocumentInitInput(type_ int32, returnUrl string, errorUrl string, ) *SignDocumentInitInput`

NewSignDocumentInitInput instantiates a new SignDocumentInitInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignDocumentInitInputWithDefaults

`func NewSignDocumentInitInputWithDefaults() *SignDocumentInitInput`

NewSignDocumentInitInputWithDefaults instantiates a new SignDocumentInitInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMobileNumber

`func (o *SignDocumentInitInput) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *SignDocumentInitInput) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *SignDocumentInitInput) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.

### HasMobileNumber

`func (o *SignDocumentInitInput) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.

### GetType

`func (o *SignDocumentInitInput) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SignDocumentInitInput) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SignDocumentInitInput) SetType(v int32)`

SetType sets Type field to given value.


### GetReturnUrl

`func (o *SignDocumentInitInput) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *SignDocumentInitInput) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *SignDocumentInitInput) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.


### GetErrorUrl

`func (o *SignDocumentInitInput) GetErrorUrl() string`

GetErrorUrl returns the ErrorUrl field if non-nil, zero value otherwise.

### GetErrorUrlOk

`func (o *SignDocumentInitInput) GetErrorUrlOk() (*string, bool)`

GetErrorUrlOk returns a tuple with the ErrorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorUrl

`func (o *SignDocumentInitInput) SetErrorUrl(v string)`

SetErrorUrl sets ErrorUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


