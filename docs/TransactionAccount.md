# TransactionAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionIn** | Pointer to [**TransactionInDetails**](TransactionInDetails.md) |  | [optional] 
**TransactionOut** | Pointer to [**TransactionOut**](TransactionOut.md) |  | [optional] 
**TransactionP2P** | Pointer to [**TransactionP2P**](TransactionP2P.md) |  | [optional] 

## Methods

### NewTransactionAccount

`func NewTransactionAccount() *TransactionAccount`

NewTransactionAccount instantiates a new TransactionAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionAccountWithDefaults

`func NewTransactionAccountWithDefaults() *TransactionAccount`

NewTransactionAccountWithDefaults instantiates a new TransactionAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionIn

`func (o *TransactionAccount) GetTransactionIn() TransactionInDetails`

GetTransactionIn returns the TransactionIn field if non-nil, zero value otherwise.

### GetTransactionInOk

`func (o *TransactionAccount) GetTransactionInOk() (*TransactionInDetails, bool)`

GetTransactionInOk returns a tuple with the TransactionIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIn

`func (o *TransactionAccount) SetTransactionIn(v TransactionInDetails)`

SetTransactionIn sets TransactionIn field to given value.

### HasTransactionIn

`func (o *TransactionAccount) HasTransactionIn() bool`

HasTransactionIn returns a boolean if a field has been set.

### GetTransactionOut

`func (o *TransactionAccount) GetTransactionOut() TransactionOut`

GetTransactionOut returns the TransactionOut field if non-nil, zero value otherwise.

### GetTransactionOutOk

`func (o *TransactionAccount) GetTransactionOutOk() (*TransactionOut, bool)`

GetTransactionOutOk returns a tuple with the TransactionOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionOut

`func (o *TransactionAccount) SetTransactionOut(v TransactionOut)`

SetTransactionOut sets TransactionOut field to given value.

### HasTransactionOut

`func (o *TransactionAccount) HasTransactionOut() bool`

HasTransactionOut returns a boolean if a field has been set.

### GetTransactionP2P

`func (o *TransactionAccount) GetTransactionP2P() TransactionP2P`

GetTransactionP2P returns the TransactionP2P field if non-nil, zero value otherwise.

### GetTransactionP2POk

`func (o *TransactionAccount) GetTransactionP2POk() (*TransactionP2P, bool)`

GetTransactionP2POk returns a tuple with the TransactionP2P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionP2P

`func (o *TransactionAccount) SetTransactionP2P(v TransactionP2P)`

SetTransactionP2P sets TransactionP2P field to given value.

### HasTransactionP2P

`func (o *TransactionAccount) HasTransactionP2P() bool`

HasTransactionP2P returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


