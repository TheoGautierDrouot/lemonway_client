# AccountBalancesInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateDate** | Pointer to **string** | Date in UTC Seconds  Leave empty to use payment account IDs. | [optional] 
**InternalAccountIdStart** | Pointer to **int64** | First payment account internal ID, starting from 12. | [optional] 
**InternalAccountIdEnd** | Pointer to **int64** | Last payment account internal ID, starting from 12. | [optional] 

## Methods

### NewAccountBalancesInput

`func NewAccountBalancesInput() *AccountBalancesInput`

NewAccountBalancesInput instantiates a new AccountBalancesInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountBalancesInputWithDefaults

`func NewAccountBalancesInputWithDefaults() *AccountBalancesInput`

NewAccountBalancesInputWithDefaults instantiates a new AccountBalancesInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateDate

`func (o *AccountBalancesInput) GetUpdateDate() string`

GetUpdateDate returns the UpdateDate field if non-nil, zero value otherwise.

### GetUpdateDateOk

`func (o *AccountBalancesInput) GetUpdateDateOk() (*string, bool)`

GetUpdateDateOk returns a tuple with the UpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDate

`func (o *AccountBalancesInput) SetUpdateDate(v string)`

SetUpdateDate sets UpdateDate field to given value.

### HasUpdateDate

`func (o *AccountBalancesInput) HasUpdateDate() bool`

HasUpdateDate returns a boolean if a field has been set.

### GetInternalAccountIdStart

`func (o *AccountBalancesInput) GetInternalAccountIdStart() int64`

GetInternalAccountIdStart returns the InternalAccountIdStart field if non-nil, zero value otherwise.

### GetInternalAccountIdStartOk

`func (o *AccountBalancesInput) GetInternalAccountIdStartOk() (*int64, bool)`

GetInternalAccountIdStartOk returns a tuple with the InternalAccountIdStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAccountIdStart

`func (o *AccountBalancesInput) SetInternalAccountIdStart(v int64)`

SetInternalAccountIdStart sets InternalAccountIdStart field to given value.

### HasInternalAccountIdStart

`func (o *AccountBalancesInput) HasInternalAccountIdStart() bool`

HasInternalAccountIdStart returns a boolean if a field has been set.

### GetInternalAccountIdEnd

`func (o *AccountBalancesInput) GetInternalAccountIdEnd() int64`

GetInternalAccountIdEnd returns the InternalAccountIdEnd field if non-nil, zero value otherwise.

### GetInternalAccountIdEndOk

`func (o *AccountBalancesInput) GetInternalAccountIdEndOk() (*int64, bool)`

GetInternalAccountIdEndOk returns a tuple with the InternalAccountIdEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAccountIdEnd

`func (o *AccountBalancesInput) SetInternalAccountIdEnd(v int64)`

SetInternalAccountIdEnd sets InternalAccountIdEnd field to given value.

### HasInternalAccountIdEnd

`func (o *AccountBalancesInput) HasInternalAccountIdEnd() bool`

HasInternalAccountIdEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


