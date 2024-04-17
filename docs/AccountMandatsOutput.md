# AccountMandatsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mandates** | Pointer to [**[]SddMandate**](SddMandate.md) | List of documents that changed since the entry date. | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewAccountMandatsOutput

`func NewAccountMandatsOutput() *AccountMandatsOutput`

NewAccountMandatsOutput instantiates a new AccountMandatsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountMandatsOutputWithDefaults

`func NewAccountMandatsOutputWithDefaults() *AccountMandatsOutput`

NewAccountMandatsOutputWithDefaults instantiates a new AccountMandatsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMandates

`func (o *AccountMandatsOutput) GetMandates() []SddMandate`

GetMandates returns the Mandates field if non-nil, zero value otherwise.

### GetMandatesOk

`func (o *AccountMandatsOutput) GetMandatesOk() (*[]SddMandate, bool)`

GetMandatesOk returns a tuple with the Mandates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandates

`func (o *AccountMandatsOutput) SetMandates(v []SddMandate)`

SetMandates sets Mandates field to given value.

### HasMandates

`func (o *AccountMandatsOutput) HasMandates() bool`

HasMandates returns a boolean if a field has been set.

### GetError

`func (o *AccountMandatsOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AccountMandatsOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AccountMandatsOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *AccountMandatsOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


