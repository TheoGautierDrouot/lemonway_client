# Limits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalMoneyInAllowed** | Pointer to **int64** | Total Money-In number allowed for this account | [optional] 
**AmountMoneyInAllowed** | Pointer to **float64** | Total Money-In amount allowed for this wallet | [optional] 

## Methods

### NewLimits

`func NewLimits() *Limits`

NewLimits instantiates a new Limits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLimitsWithDefaults

`func NewLimitsWithDefaults() *Limits`

NewLimitsWithDefaults instantiates a new Limits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalMoneyInAllowed

`func (o *Limits) GetTotalMoneyInAllowed() int64`

GetTotalMoneyInAllowed returns the TotalMoneyInAllowed field if non-nil, zero value otherwise.

### GetTotalMoneyInAllowedOk

`func (o *Limits) GetTotalMoneyInAllowedOk() (*int64, bool)`

GetTotalMoneyInAllowedOk returns a tuple with the TotalMoneyInAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMoneyInAllowed

`func (o *Limits) SetTotalMoneyInAllowed(v int64)`

SetTotalMoneyInAllowed sets TotalMoneyInAllowed field to given value.

### HasTotalMoneyInAllowed

`func (o *Limits) HasTotalMoneyInAllowed() bool`

HasTotalMoneyInAllowed returns a boolean if a field has been set.

### GetAmountMoneyInAllowed

`func (o *Limits) GetAmountMoneyInAllowed() float64`

GetAmountMoneyInAllowed returns the AmountMoneyInAllowed field if non-nil, zero value otherwise.

### GetAmountMoneyInAllowedOk

`func (o *Limits) GetAmountMoneyInAllowedOk() (*float64, bool)`

GetAmountMoneyInAllowedOk returns a tuple with the AmountMoneyInAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMoneyInAllowed

`func (o *Limits) SetAmountMoneyInAllowed(v float64)`

SetAmountMoneyInAllowed sets AmountMoneyInAllowed field to given value.

### HasAmountMoneyInAllowed

`func (o *Limits) HasAmountMoneyInAllowed() bool`

HasAmountMoneyInAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


