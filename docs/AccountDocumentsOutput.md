# AccountDocumentsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Documents** | Pointer to [**[]Document**](Document.md) | List of documents that changed since the entry date. | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewAccountDocumentsOutput

`func NewAccountDocumentsOutput() *AccountDocumentsOutput`

NewAccountDocumentsOutput instantiates a new AccountDocumentsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDocumentsOutputWithDefaults

`func NewAccountDocumentsOutputWithDefaults() *AccountDocumentsOutput`

NewAccountDocumentsOutputWithDefaults instantiates a new AccountDocumentsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocuments

`func (o *AccountDocumentsOutput) GetDocuments() []Document`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *AccountDocumentsOutput) GetDocumentsOk() (*[]Document, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *AccountDocumentsOutput) SetDocuments(v []Document)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *AccountDocumentsOutput) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetError

`func (o *AccountDocumentsOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AccountDocumentsOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AccountDocumentsOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *AccountDocumentsOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


