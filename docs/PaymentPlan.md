# PaymentPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Payment plan ID. | [optional] 
**Provider** | Pointer to **string** | PSP name. | [optional] 
**Type** | Pointer to **string** | Payment Plan Type.  ex: Installments or deferrerd. | [optional] 
**DeferredDays** | Pointer to **int32** | Number of deferred days. | [optional] 
**Installments** | Pointer to **int32** | Number of installments. | [optional] 
**Description** | Pointer to **string** | Payment plan description.  ex: The marketplace bears the fees. | [optional] 

## Methods

### NewPaymentPlan

`func NewPaymentPlan() *PaymentPlan`

NewPaymentPlan instantiates a new PaymentPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentPlanWithDefaults

`func NewPaymentPlanWithDefaults() *PaymentPlan`

NewPaymentPlanWithDefaults instantiates a new PaymentPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentPlan) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentPlan) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentPlan) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProvider

`func (o *PaymentPlan) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PaymentPlan) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PaymentPlan) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *PaymentPlan) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetType

`func (o *PaymentPlan) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentPlan) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentPlan) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentPlan) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDeferredDays

`func (o *PaymentPlan) GetDeferredDays() int32`

GetDeferredDays returns the DeferredDays field if non-nil, zero value otherwise.

### GetDeferredDaysOk

`func (o *PaymentPlan) GetDeferredDaysOk() (*int32, bool)`

GetDeferredDaysOk returns a tuple with the DeferredDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferredDays

`func (o *PaymentPlan) SetDeferredDays(v int32)`

SetDeferredDays sets DeferredDays field to given value.

### HasDeferredDays

`func (o *PaymentPlan) HasDeferredDays() bool`

HasDeferredDays returns a boolean if a field has been set.

### GetInstallments

`func (o *PaymentPlan) GetInstallments() int32`

GetInstallments returns the Installments field if non-nil, zero value otherwise.

### GetInstallmentsOk

`func (o *PaymentPlan) GetInstallmentsOk() (*int32, bool)`

GetInstallmentsOk returns a tuple with the Installments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallments

`func (o *PaymentPlan) SetInstallments(v int32)`

SetInstallments sets Installments field to given value.

### HasInstallments

`func (o *PaymentPlan) HasInstallments() bool`

HasInstallments returns a boolean if a field has been set.

### GetDescription

`func (o *PaymentPlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentPlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentPlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentPlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


