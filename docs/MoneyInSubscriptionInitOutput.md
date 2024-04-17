# MoneyInSubscriptionInitOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | Pointer to [**TransactionsTransactionIn**](TransactionsTransactionIn.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewMoneyInSubscriptionInitOutput

`func NewMoneyInSubscriptionInitOutput() *MoneyInSubscriptionInitOutput`

NewMoneyInSubscriptionInitOutput instantiates a new MoneyInSubscriptionInitOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoneyInSubscriptionInitOutputWithDefaults

`func NewMoneyInSubscriptionInitOutputWithDefaults() *MoneyInSubscriptionInitOutput`

NewMoneyInSubscriptionInitOutputWithDefaults instantiates a new MoneyInSubscriptionInitOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *MoneyInSubscriptionInitOutput) GetTransactions() TransactionsTransactionIn`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *MoneyInSubscriptionInitOutput) GetTransactionsOk() (*TransactionsTransactionIn, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *MoneyInSubscriptionInitOutput) SetTransactions(v TransactionsTransactionIn)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *MoneyInSubscriptionInitOutput) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetError

`func (o *MoneyInSubscriptionInitOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MoneyInSubscriptionInitOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MoneyInSubscriptionInitOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *MoneyInSubscriptionInitOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


