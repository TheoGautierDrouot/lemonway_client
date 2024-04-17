# GetMoneyOutTransDetailsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | Pointer to [**TransactionsTransactionOut**](TransactionsTransactionOut.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewGetMoneyOutTransDetailsOutput

`func NewGetMoneyOutTransDetailsOutput() *GetMoneyOutTransDetailsOutput`

NewGetMoneyOutTransDetailsOutput instantiates a new GetMoneyOutTransDetailsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMoneyOutTransDetailsOutputWithDefaults

`func NewGetMoneyOutTransDetailsOutputWithDefaults() *GetMoneyOutTransDetailsOutput`

NewGetMoneyOutTransDetailsOutputWithDefaults instantiates a new GetMoneyOutTransDetailsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *GetMoneyOutTransDetailsOutput) GetTransactions() TransactionsTransactionOut`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *GetMoneyOutTransDetailsOutput) GetTransactionsOk() (*TransactionsTransactionOut, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *GetMoneyOutTransDetailsOutput) SetTransactions(v TransactionsTransactionOut)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *GetMoneyOutTransDetailsOutput) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetLinks

`func (o *GetMoneyOutTransDetailsOutput) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetMoneyOutTransDetailsOutput) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetMoneyOutTransDetailsOutput) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetMoneyOutTransDetailsOutput) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetError

`func (o *GetMoneyOutTransDetailsOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetMoneyOutTransDetailsOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetMoneyOutTransDetailsOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *GetMoneyOutTransDetailsOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


