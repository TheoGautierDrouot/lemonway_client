# AccountBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Updated payment account ID | [optional] 
**IsBlocked** | Pointer to **bool** | Status of the payment account  1: is bocked true  0: is blocked false | [optional] 

## Methods

### NewAccountBlock

`func NewAccountBlock() *AccountBlock`

NewAccountBlock instantiates a new AccountBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountBlockWithDefaults

`func NewAccountBlockWithDefaults() *AccountBlock`

NewAccountBlockWithDefaults instantiates a new AccountBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsBlocked

`func (o *AccountBlock) GetIsBlocked() bool`

GetIsBlocked returns the IsBlocked field if non-nil, zero value otherwise.

### GetIsBlockedOk

`func (o *AccountBlock) GetIsBlockedOk() (*bool, bool)`

GetIsBlockedOk returns a tuple with the IsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlocked

`func (o *AccountBlock) SetIsBlocked(v bool)`

SetIsBlocked sets IsBlocked field to given value.

### HasIsBlocked

`func (o *AccountBlock) HasIsBlocked() bool`

HasIsBlocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


