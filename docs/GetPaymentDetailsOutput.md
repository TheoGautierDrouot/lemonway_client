# GetPaymentDetailsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | Pointer to [**TransactionsTransactionP2P**](TransactionsTransactionP2P.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewGetPaymentDetailsOutput

`func NewGetPaymentDetailsOutput() *GetPaymentDetailsOutput`

NewGetPaymentDetailsOutput instantiates a new GetPaymentDetailsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPaymentDetailsOutputWithDefaults

`func NewGetPaymentDetailsOutputWithDefaults() *GetPaymentDetailsOutput`

NewGetPaymentDetailsOutputWithDefaults instantiates a new GetPaymentDetailsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *GetPaymentDetailsOutput) GetTransactions() TransactionsTransactionP2P`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *GetPaymentDetailsOutput) GetTransactionsOk() (*TransactionsTransactionP2P, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *GetPaymentDetailsOutput) SetTransactions(v TransactionsTransactionP2P)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *GetPaymentDetailsOutput) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetLinks

`func (o *GetPaymentDetailsOutput) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetPaymentDetailsOutput) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetPaymentDetailsOutput) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetPaymentDetailsOutput) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetError

`func (o *GetPaymentDetailsOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetPaymentDetailsOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetPaymentDetailsOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *GetPaymentDetailsOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


