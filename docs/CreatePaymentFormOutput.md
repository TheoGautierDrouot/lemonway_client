# CreatePaymentFormOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Form** | Pointer to [**PaymentForm**](PaymentForm.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewCreatePaymentFormOutput

`func NewCreatePaymentFormOutput() *CreatePaymentFormOutput`

NewCreatePaymentFormOutput instantiates a new CreatePaymentFormOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaymentFormOutputWithDefaults

`func NewCreatePaymentFormOutputWithDefaults() *CreatePaymentFormOutput`

NewCreatePaymentFormOutputWithDefaults instantiates a new CreatePaymentFormOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForm

`func (o *CreatePaymentFormOutput) GetForm() PaymentForm`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *CreatePaymentFormOutput) GetFormOk() (*PaymentForm, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *CreatePaymentFormOutput) SetForm(v PaymentForm)`

SetForm sets Form field to given value.

### HasForm

`func (o *CreatePaymentFormOutput) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetError

`func (o *CreatePaymentFormOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CreatePaymentFormOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CreatePaymentFormOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *CreatePaymentFormOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


