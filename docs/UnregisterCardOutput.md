# UnregisterCardOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardId** | Pointer to **int64** | ID Card | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewUnregisterCardOutput

`func NewUnregisterCardOutput() *UnregisterCardOutput`

NewUnregisterCardOutput instantiates a new UnregisterCardOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnregisterCardOutputWithDefaults

`func NewUnregisterCardOutputWithDefaults() *UnregisterCardOutput`

NewUnregisterCardOutputWithDefaults instantiates a new UnregisterCardOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardId

`func (o *UnregisterCardOutput) GetCardId() int64`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *UnregisterCardOutput) GetCardIdOk() (*int64, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *UnregisterCardOutput) SetCardId(v int64)`

SetCardId sets CardId field to given value.

### HasCardId

`func (o *UnregisterCardOutput) HasCardId() bool`

HasCardId returns a boolean if a field has been set.

### GetError

`func (o *UnregisterCardOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UnregisterCardOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UnregisterCardOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *UnregisterCardOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


