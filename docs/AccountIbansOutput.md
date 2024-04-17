# AccountIbansOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ibans** | Pointer to [**[]Iban**](Iban.md) | List of documents that changed since the entry date. | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewAccountIbansOutput

`func NewAccountIbansOutput() *AccountIbansOutput`

NewAccountIbansOutput instantiates a new AccountIbansOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountIbansOutputWithDefaults

`func NewAccountIbansOutputWithDefaults() *AccountIbansOutput`

NewAccountIbansOutputWithDefaults instantiates a new AccountIbansOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIbans

`func (o *AccountIbansOutput) GetIbans() []Iban`

GetIbans returns the Ibans field if non-nil, zero value otherwise.

### GetIbansOk

`func (o *AccountIbansOutput) GetIbansOk() (*[]Iban, bool)`

GetIbansOk returns a tuple with the Ibans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbans

`func (o *AccountIbansOutput) SetIbans(v []Iban)`

SetIbans sets Ibans field to given value.

### HasIbans

`func (o *AccountIbansOutput) HasIbans() bool`

HasIbans returns a boolean if a field has been set.

### GetError

`func (o *AccountIbansOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AccountIbansOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AccountIbansOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *AccountIbansOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


