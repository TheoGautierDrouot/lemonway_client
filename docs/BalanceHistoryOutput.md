# BalanceHistoryOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | Pointer to **int32** | Payment account balance. | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewBalanceHistoryOutput

`func NewBalanceHistoryOutput() *BalanceHistoryOutput`

NewBalanceHistoryOutput instantiates a new BalanceHistoryOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceHistoryOutputWithDefaults

`func NewBalanceHistoryOutputWithDefaults() *BalanceHistoryOutput`

NewBalanceHistoryOutputWithDefaults instantiates a new BalanceHistoryOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *BalanceHistoryOutput) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *BalanceHistoryOutput) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *BalanceHistoryOutput) SetBalance(v int32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *BalanceHistoryOutput) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetError

`func (o *BalanceHistoryOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BalanceHistoryOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BalanceHistoryOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *BalanceHistoryOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


