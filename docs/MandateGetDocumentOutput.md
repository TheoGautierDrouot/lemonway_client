# MandateGetDocumentOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Document** | Pointer to **string** | Document in base64 | [optional] 
**Name** | Pointer to **string** | Document name | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewMandateGetDocumentOutput

`func NewMandateGetDocumentOutput() *MandateGetDocumentOutput`

NewMandateGetDocumentOutput instantiates a new MandateGetDocumentOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMandateGetDocumentOutputWithDefaults

`func NewMandateGetDocumentOutputWithDefaults() *MandateGetDocumentOutput`

NewMandateGetDocumentOutputWithDefaults instantiates a new MandateGetDocumentOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocument

`func (o *MandateGetDocumentOutput) GetDocument() string`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *MandateGetDocumentOutput) GetDocumentOk() (*string, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *MandateGetDocumentOutput) SetDocument(v string)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *MandateGetDocumentOutput) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetName

`func (o *MandateGetDocumentOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MandateGetDocumentOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MandateGetDocumentOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MandateGetDocumentOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetError

`func (o *MandateGetDocumentOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MandateGetDocumentOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MandateGetDocumentOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *MandateGetDocumentOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


