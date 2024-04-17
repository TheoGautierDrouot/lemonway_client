# RefundMoneyInOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transaction** | Pointer to [**TransactionRefund**](TransactionRefund.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewRefundMoneyInOutput

`func NewRefundMoneyInOutput() *RefundMoneyInOutput`

NewRefundMoneyInOutput instantiates a new RefundMoneyInOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundMoneyInOutputWithDefaults

`func NewRefundMoneyInOutputWithDefaults() *RefundMoneyInOutput`

NewRefundMoneyInOutputWithDefaults instantiates a new RefundMoneyInOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransaction

`func (o *RefundMoneyInOutput) GetTransaction() TransactionRefund`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *RefundMoneyInOutput) GetTransactionOk() (*TransactionRefund, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *RefundMoneyInOutput) SetTransaction(v TransactionRefund)`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *RefundMoneyInOutput) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.

### GetError

`func (o *RefundMoneyInOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RefundMoneyInOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RefundMoneyInOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *RefundMoneyInOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


