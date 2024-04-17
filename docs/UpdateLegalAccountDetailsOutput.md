# UpdateLegalAccountDetailsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**LegalAccount**](LegalAccount.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewUpdateLegalAccountDetailsOutput

`func NewUpdateLegalAccountDetailsOutput() *UpdateLegalAccountDetailsOutput`

NewUpdateLegalAccountDetailsOutput instantiates a new UpdateLegalAccountDetailsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLegalAccountDetailsOutputWithDefaults

`func NewUpdateLegalAccountDetailsOutputWithDefaults() *UpdateLegalAccountDetailsOutput`

NewUpdateLegalAccountDetailsOutputWithDefaults instantiates a new UpdateLegalAccountDetailsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *UpdateLegalAccountDetailsOutput) GetAccount() LegalAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateLegalAccountDetailsOutput) GetAccountOk() (*LegalAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateLegalAccountDetailsOutput) SetAccount(v LegalAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateLegalAccountDetailsOutput) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetError

`func (o *UpdateLegalAccountDetailsOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UpdateLegalAccountDetailsOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UpdateLegalAccountDetailsOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *UpdateLegalAccountDetailsOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


