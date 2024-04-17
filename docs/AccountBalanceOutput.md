# AccountBalanceOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]AccountBalance**](AccountBalance.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewAccountBalanceOutput

`func NewAccountBalanceOutput() *AccountBalanceOutput`

NewAccountBalanceOutput instantiates a new AccountBalanceOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountBalanceOutputWithDefaults

`func NewAccountBalanceOutputWithDefaults() *AccountBalanceOutput`

NewAccountBalanceOutputWithDefaults instantiates a new AccountBalanceOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *AccountBalanceOutput) GetAccounts() []AccountBalance`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AccountBalanceOutput) GetAccountsOk() (*[]AccountBalance, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AccountBalanceOutput) SetAccounts(v []AccountBalance)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *AccountBalanceOutput) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetLinks

`func (o *AccountBalanceOutput) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AccountBalanceOutput) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AccountBalanceOutput) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AccountBalanceOutput) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetError

`func (o *AccountBalanceOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AccountBalanceOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AccountBalanceOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *AccountBalanceOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


