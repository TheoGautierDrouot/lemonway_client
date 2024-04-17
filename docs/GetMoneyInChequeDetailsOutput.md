# GetMoneyInChequeDetailsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | Pointer to [**TransactionsTransactionIn**](TransactionsTransactionIn.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewGetMoneyInChequeDetailsOutput

`func NewGetMoneyInChequeDetailsOutput() *GetMoneyInChequeDetailsOutput`

NewGetMoneyInChequeDetailsOutput instantiates a new GetMoneyInChequeDetailsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMoneyInChequeDetailsOutputWithDefaults

`func NewGetMoneyInChequeDetailsOutputWithDefaults() *GetMoneyInChequeDetailsOutput`

NewGetMoneyInChequeDetailsOutputWithDefaults instantiates a new GetMoneyInChequeDetailsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *GetMoneyInChequeDetailsOutput) GetTransactions() TransactionsTransactionIn`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *GetMoneyInChequeDetailsOutput) GetTransactionsOk() (*TransactionsTransactionIn, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *GetMoneyInChequeDetailsOutput) SetTransactions(v TransactionsTransactionIn)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *GetMoneyInChequeDetailsOutput) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetLinks

`func (o *GetMoneyInChequeDetailsOutput) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetMoneyInChequeDetailsOutput) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetMoneyInChequeDetailsOutput) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetMoneyInChequeDetailsOutput) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetError

`func (o *GetMoneyInChequeDetailsOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetMoneyInChequeDetailsOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetMoneyInChequeDetailsOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *GetMoneyInChequeDetailsOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


