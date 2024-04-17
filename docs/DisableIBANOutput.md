# DisableIBANOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | IBAN ID | [optional] 
**Status** | Pointer to **int32** | IBAN Status | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewDisableIBANOutput

`func NewDisableIBANOutput() *DisableIBANOutput`

NewDisableIBANOutput instantiates a new DisableIBANOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisableIBANOutputWithDefaults

`func NewDisableIBANOutputWithDefaults() *DisableIBANOutput`

NewDisableIBANOutputWithDefaults instantiates a new DisableIBANOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DisableIBANOutput) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DisableIBANOutput) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DisableIBANOutput) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DisableIBANOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *DisableIBANOutput) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DisableIBANOutput) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DisableIBANOutput) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DisableIBANOutput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetError

`func (o *DisableIBANOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DisableIBANOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DisableIBANOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *DisableIBANOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


