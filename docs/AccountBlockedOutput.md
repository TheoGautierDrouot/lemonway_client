# AccountBlockedOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**AccountBlock**](AccountBlock.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewAccountBlockedOutput

`func NewAccountBlockedOutput() *AccountBlockedOutput`

NewAccountBlockedOutput instantiates a new AccountBlockedOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountBlockedOutputWithDefaults

`func NewAccountBlockedOutputWithDefaults() *AccountBlockedOutput`

NewAccountBlockedOutputWithDefaults instantiates a new AccountBlockedOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AccountBlockedOutput) GetAccount() AccountBlock`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AccountBlockedOutput) GetAccountOk() (*AccountBlock, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AccountBlockedOutput) SetAccount(v AccountBlock)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AccountBlockedOutput) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetError

`func (o *AccountBlockedOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AccountBlockedOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AccountBlockedOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *AccountBlockedOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


