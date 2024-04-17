# GetMoneyInSddOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | Pointer to [**TransactionsTransactionIn**](TransactionsTransactionIn.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewGetMoneyInSddOutput

`func NewGetMoneyInSddOutput() *GetMoneyInSddOutput`

NewGetMoneyInSddOutput instantiates a new GetMoneyInSddOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMoneyInSddOutputWithDefaults

`func NewGetMoneyInSddOutputWithDefaults() *GetMoneyInSddOutput`

NewGetMoneyInSddOutputWithDefaults instantiates a new GetMoneyInSddOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *GetMoneyInSddOutput) GetTransactions() TransactionsTransactionIn`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *GetMoneyInSddOutput) GetTransactionsOk() (*TransactionsTransactionIn, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *GetMoneyInSddOutput) SetTransactions(v TransactionsTransactionIn)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *GetMoneyInSddOutput) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetLinks

`func (o *GetMoneyInSddOutput) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetMoneyInSddOutput) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetMoneyInSddOutput) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetMoneyInSddOutput) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetError

`func (o *GetMoneyInSddOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetMoneyInSddOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetMoneyInSddOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *GetMoneyInSddOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


