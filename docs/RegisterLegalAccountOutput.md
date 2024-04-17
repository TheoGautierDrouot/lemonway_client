# RegisterLegalAccountOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LegalAccount** | Pointer to [**LegalAccount**](LegalAccount.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewRegisterLegalAccountOutput

`func NewRegisterLegalAccountOutput() *RegisterLegalAccountOutput`

NewRegisterLegalAccountOutput instantiates a new RegisterLegalAccountOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterLegalAccountOutputWithDefaults

`func NewRegisterLegalAccountOutputWithDefaults() *RegisterLegalAccountOutput`

NewRegisterLegalAccountOutputWithDefaults instantiates a new RegisterLegalAccountOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLegalAccount

`func (o *RegisterLegalAccountOutput) GetLegalAccount() LegalAccount`

GetLegalAccount returns the LegalAccount field if non-nil, zero value otherwise.

### GetLegalAccountOk

`func (o *RegisterLegalAccountOutput) GetLegalAccountOk() (*LegalAccount, bool)`

GetLegalAccountOk returns a tuple with the LegalAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalAccount

`func (o *RegisterLegalAccountOutput) SetLegalAccount(v LegalAccount)`

SetLegalAccount sets LegalAccount field to given value.

### HasLegalAccount

`func (o *RegisterLegalAccountOutput) HasLegalAccount() bool`

HasLegalAccount returns a boolean if a field has been set.

### GetError

`func (o *RegisterLegalAccountOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RegisterLegalAccountOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RegisterLegalAccountOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *RegisterLegalAccountOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


