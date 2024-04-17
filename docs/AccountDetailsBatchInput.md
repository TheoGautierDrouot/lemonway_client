# AccountDetailsBatchInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | [**[]WalletDetailsInput**](WalletDetailsInput.md) |  | 

## Methods

### NewAccountDetailsBatchInput

`func NewAccountDetailsBatchInput(accounts []WalletDetailsInput, ) *AccountDetailsBatchInput`

NewAccountDetailsBatchInput instantiates a new AccountDetailsBatchInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDetailsBatchInputWithDefaults

`func NewAccountDetailsBatchInputWithDefaults() *AccountDetailsBatchInput`

NewAccountDetailsBatchInputWithDefaults instantiates a new AccountDetailsBatchInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *AccountDetailsBatchInput) GetAccounts() []WalletDetailsInput`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AccountDetailsBatchInput) GetAccountsOk() (*[]WalletDetailsInput, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AccountDetailsBatchInput) SetAccounts(v []WalletDetailsInput)`

SetAccounts sets Accounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


