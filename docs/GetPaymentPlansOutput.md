# GetPaymentPlansOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentPlans** | Pointer to [**[]PaymentPlan**](PaymentPlan.md) | List of actived BNPL payment plans | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewGetPaymentPlansOutput

`func NewGetPaymentPlansOutput() *GetPaymentPlansOutput`

NewGetPaymentPlansOutput instantiates a new GetPaymentPlansOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPaymentPlansOutputWithDefaults

`func NewGetPaymentPlansOutputWithDefaults() *GetPaymentPlansOutput`

NewGetPaymentPlansOutputWithDefaults instantiates a new GetPaymentPlansOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentPlans

`func (o *GetPaymentPlansOutput) GetPaymentPlans() []PaymentPlan`

GetPaymentPlans returns the PaymentPlans field if non-nil, zero value otherwise.

### GetPaymentPlansOk

`func (o *GetPaymentPlansOutput) GetPaymentPlansOk() (*[]PaymentPlan, bool)`

GetPaymentPlansOk returns a tuple with the PaymentPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentPlans

`func (o *GetPaymentPlansOutput) SetPaymentPlans(v []PaymentPlan)`

SetPaymentPlans sets PaymentPlans field to given value.

### HasPaymentPlans

`func (o *GetPaymentPlansOutput) HasPaymentPlans() bool`

HasPaymentPlans returns a boolean if a field has been set.

### GetError

`func (o *GetPaymentPlansOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetPaymentPlansOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetPaymentPlansOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *GetPaymentPlansOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


