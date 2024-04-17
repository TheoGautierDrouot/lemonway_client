# RegisterIndividualAccountOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**IndividualAccount**](IndividualAccount.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewRegisterIndividualAccountOutput

`func NewRegisterIndividualAccountOutput() *RegisterIndividualAccountOutput`

NewRegisterIndividualAccountOutput instantiates a new RegisterIndividualAccountOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterIndividualAccountOutputWithDefaults

`func NewRegisterIndividualAccountOutputWithDefaults() *RegisterIndividualAccountOutput`

NewRegisterIndividualAccountOutputWithDefaults instantiates a new RegisterIndividualAccountOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *RegisterIndividualAccountOutput) GetAccount() IndividualAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *RegisterIndividualAccountOutput) GetAccountOk() (*IndividualAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *RegisterIndividualAccountOutput) SetAccount(v IndividualAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *RegisterIndividualAccountOutput) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetError

`func (o *RegisterIndividualAccountOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RegisterIndividualAccountOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RegisterIndividualAccountOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *RegisterIndividualAccountOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


