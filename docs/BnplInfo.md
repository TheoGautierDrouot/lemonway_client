# BnplInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentPlanId** | Pointer to **int64** | Payment Plan ID | [optional] 
**Type** | Pointer to **string** | Payment plan type. | [optional] 
**SecondInstallmentPaymentDate** | Pointer to **string** | Second installment payment date.  Format: yyyy-MM-dd | [optional] 
**DeferredPaymentDate** | Pointer to **string** | Deferred payment date.  Format: yyyy-MM-dd | [optional] 

## Methods

### NewBnplInfo

`func NewBnplInfo() *BnplInfo`

NewBnplInfo instantiates a new BnplInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnplInfoWithDefaults

`func NewBnplInfoWithDefaults() *BnplInfo`

NewBnplInfoWithDefaults instantiates a new BnplInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentPlanId

`func (o *BnplInfo) GetPaymentPlanId() int64`

GetPaymentPlanId returns the PaymentPlanId field if non-nil, zero value otherwise.

### GetPaymentPlanIdOk

`func (o *BnplInfo) GetPaymentPlanIdOk() (*int64, bool)`

GetPaymentPlanIdOk returns a tuple with the PaymentPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentPlanId

`func (o *BnplInfo) SetPaymentPlanId(v int64)`

SetPaymentPlanId sets PaymentPlanId field to given value.

### HasPaymentPlanId

`func (o *BnplInfo) HasPaymentPlanId() bool`

HasPaymentPlanId returns a boolean if a field has been set.

### GetType

`func (o *BnplInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BnplInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BnplInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BnplInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSecondInstallmentPaymentDate

`func (o *BnplInfo) GetSecondInstallmentPaymentDate() string`

GetSecondInstallmentPaymentDate returns the SecondInstallmentPaymentDate field if non-nil, zero value otherwise.

### GetSecondInstallmentPaymentDateOk

`func (o *BnplInfo) GetSecondInstallmentPaymentDateOk() (*string, bool)`

GetSecondInstallmentPaymentDateOk returns a tuple with the SecondInstallmentPaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondInstallmentPaymentDate

`func (o *BnplInfo) SetSecondInstallmentPaymentDate(v string)`

SetSecondInstallmentPaymentDate sets SecondInstallmentPaymentDate field to given value.

### HasSecondInstallmentPaymentDate

`func (o *BnplInfo) HasSecondInstallmentPaymentDate() bool`

HasSecondInstallmentPaymentDate returns a boolean if a field has been set.

### GetDeferredPaymentDate

`func (o *BnplInfo) GetDeferredPaymentDate() string`

GetDeferredPaymentDate returns the DeferredPaymentDate field if non-nil, zero value otherwise.

### GetDeferredPaymentDateOk

`func (o *BnplInfo) GetDeferredPaymentDateOk() (*string, bool)`

GetDeferredPaymentDateOk returns a tuple with the DeferredPaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferredPaymentDate

`func (o *BnplInfo) SetDeferredPaymentDate(v string)`

SetDeferredPaymentDate sets DeferredPaymentDate field to given value.

### HasDeferredPaymentDate

`func (o *BnplInfo) HasDeferredPaymentDate() bool`

HasDeferredPaymentDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


