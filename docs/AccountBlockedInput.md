# AccountBlockedInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsBlocked** | **bool** |  | 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountBlockedInput

`func NewAccountBlockedInput(isBlocked bool, ) *AccountBlockedInput`

NewAccountBlockedInput instantiates a new AccountBlockedInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountBlockedInputWithDefaults

`func NewAccountBlockedInputWithDefaults() *AccountBlockedInput`

NewAccountBlockedInputWithDefaults instantiates a new AccountBlockedInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsBlocked

`func (o *AccountBlockedInput) GetIsBlocked() bool`

GetIsBlocked returns the IsBlocked field if non-nil, zero value otherwise.

### GetIsBlockedOk

`func (o *AccountBlockedInput) GetIsBlockedOk() (*bool, bool)`

GetIsBlockedOk returns a tuple with the IsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlocked

`func (o *AccountBlockedInput) SetIsBlocked(v bool)`

SetIsBlocked sets IsBlocked field to given value.


### GetComment

`func (o *AccountBlockedInput) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AccountBlockedInput) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AccountBlockedInput) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AccountBlockedInput) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


