# SignDocumentInitOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignDocument** | Pointer to **string** | Payment token to use as GET parameter when redirecting your user to the WEBKIT.  so even if the user comes back to your error page (if the user cancelled for example),   you can still submit the same token to the WEBKIT. | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewSignDocumentInitOutput

`func NewSignDocumentInitOutput() *SignDocumentInitOutput`

NewSignDocumentInitOutput instantiates a new SignDocumentInitOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignDocumentInitOutputWithDefaults

`func NewSignDocumentInitOutputWithDefaults() *SignDocumentInitOutput`

NewSignDocumentInitOutputWithDefaults instantiates a new SignDocumentInitOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignDocument

`func (o *SignDocumentInitOutput) GetSignDocument() string`

GetSignDocument returns the SignDocument field if non-nil, zero value otherwise.

### GetSignDocumentOk

`func (o *SignDocumentInitOutput) GetSignDocumentOk() (*string, bool)`

GetSignDocumentOk returns a tuple with the SignDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignDocument

`func (o *SignDocumentInitOutput) SetSignDocument(v string)`

SetSignDocument sets SignDocument field to given value.

### HasSignDocument

`func (o *SignDocumentInitOutput) HasSignDocument() bool`

HasSignDocument returns a boolean if a field has been set.

### GetError

`func (o *SignDocumentInitOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SignDocumentInitOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SignDocumentInitOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *SignDocumentInitOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


