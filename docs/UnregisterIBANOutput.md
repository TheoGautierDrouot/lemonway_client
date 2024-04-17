# UnregisterIBANOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewUnregisterIBANOutput

`func NewUnregisterIBANOutput() *UnregisterIBANOutput`

NewUnregisterIBANOutput instantiates a new UnregisterIBANOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnregisterIBANOutputWithDefaults

`func NewUnregisterIBANOutputWithDefaults() *UnregisterIBANOutput`

NewUnregisterIBANOutputWithDefaults instantiates a new UnregisterIBANOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UnregisterIBANOutput) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnregisterIBANOutput) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnregisterIBANOutput) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnregisterIBANOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetError

`func (o *UnregisterIBANOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UnregisterIBANOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UnregisterIBANOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *UnregisterIBANOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


