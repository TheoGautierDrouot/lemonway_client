# BalanceHistoryInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtDate** | Pointer to **string** | Request balance at given time in UTC Unix timestamp. | [optional] 

## Methods

### NewBalanceHistoryInput

`func NewBalanceHistoryInput() *BalanceHistoryInput`

NewBalanceHistoryInput instantiates a new BalanceHistoryInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceHistoryInputWithDefaults

`func NewBalanceHistoryInputWithDefaults() *BalanceHistoryInput`

NewBalanceHistoryInputWithDefaults instantiates a new BalanceHistoryInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtDate

`func (o *BalanceHistoryInput) GetAtDate() string`

GetAtDate returns the AtDate field if non-nil, zero value otherwise.

### GetAtDateOk

`func (o *BalanceHistoryInput) GetAtDateOk() (*string, bool)`

GetAtDateOk returns a tuple with the AtDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtDate

`func (o *BalanceHistoryInput) SetAtDate(v string)`

SetAtDate sets AtDate field to given value.

### HasAtDate

`func (o *BalanceHistoryInput) HasAtDate() bool`

HasAtDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


