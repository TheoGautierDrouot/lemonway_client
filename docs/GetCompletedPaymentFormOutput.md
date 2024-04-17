# GetCompletedPaymentFormOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Form** | Pointer to [**PaymentFormDetails**](PaymentFormDetails.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewGetCompletedPaymentFormOutput

`func NewGetCompletedPaymentFormOutput() *GetCompletedPaymentFormOutput`

NewGetCompletedPaymentFormOutput instantiates a new GetCompletedPaymentFormOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCompletedPaymentFormOutputWithDefaults

`func NewGetCompletedPaymentFormOutputWithDefaults() *GetCompletedPaymentFormOutput`

NewGetCompletedPaymentFormOutputWithDefaults instantiates a new GetCompletedPaymentFormOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForm

`func (o *GetCompletedPaymentFormOutput) GetForm() PaymentFormDetails`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *GetCompletedPaymentFormOutput) GetFormOk() (*PaymentFormDetails, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *GetCompletedPaymentFormOutput) SetForm(v PaymentFormDetails)`

SetForm sets Form field to given value.

### HasForm

`func (o *GetCompletedPaymentFormOutput) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetError

`func (o *GetCompletedPaymentFormOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetCompletedPaymentFormOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetCompletedPaymentFormOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *GetCompletedPaymentFormOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


