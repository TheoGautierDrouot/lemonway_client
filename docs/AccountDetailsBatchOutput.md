# AccountDetailsBatchOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]AccountDetails**](AccountDetails.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewAccountDetailsBatchOutput

`func NewAccountDetailsBatchOutput() *AccountDetailsBatchOutput`

NewAccountDetailsBatchOutput instantiates a new AccountDetailsBatchOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDetailsBatchOutputWithDefaults

`func NewAccountDetailsBatchOutputWithDefaults() *AccountDetailsBatchOutput`

NewAccountDetailsBatchOutputWithDefaults instantiates a new AccountDetailsBatchOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *AccountDetailsBatchOutput) GetAccounts() []AccountDetails`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AccountDetailsBatchOutput) GetAccountsOk() (*[]AccountDetails, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AccountDetailsBatchOutput) SetAccounts(v []AccountDetails)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *AccountDetailsBatchOutput) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetError

`func (o *AccountDetailsBatchOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AccountDetailsBatchOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AccountDetailsBatchOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *AccountDetailsBatchOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


