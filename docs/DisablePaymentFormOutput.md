# DisablePaymentFormOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewDisablePaymentFormOutput

`func NewDisablePaymentFormOutput() *DisablePaymentFormOutput`

NewDisablePaymentFormOutput instantiates a new DisablePaymentFormOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisablePaymentFormOutputWithDefaults

`func NewDisablePaymentFormOutputWithDefaults() *DisablePaymentFormOutput`

NewDisablePaymentFormOutputWithDefaults instantiates a new DisablePaymentFormOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DisablePaymentFormOutput) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DisablePaymentFormOutput) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DisablePaymentFormOutput) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DisablePaymentFormOutput) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetError

`func (o *DisablePaymentFormOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DisablePaymentFormOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DisablePaymentFormOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *DisablePaymentFormOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


