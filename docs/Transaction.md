# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reference** | **string** | Reference | 
**AccountId** | **string** | Account Id | 
**TotalAmount** | **int32** | Total Amount | 
**CommissionAmount** | **int32** | Commission Amount | 
**Comment** | Pointer to **string** | Comment | [optional] 
**AutoCommission** | Pointer to **bool** | Auto Commission | [optional] 
**IsPreAuth** | Pointer to **bool** | Intention to authorize | [optional] 

## Methods

### NewTransaction

`func NewTransaction(reference string, accountId string, totalAmount int32, commissionAmount int32, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReference

`func (o *Transaction) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Transaction) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Transaction) SetReference(v string)`

SetReference sets Reference field to given value.


### GetAccountId

`func (o *Transaction) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Transaction) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Transaction) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetTotalAmount

`func (o *Transaction) GetTotalAmount() int32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *Transaction) GetTotalAmountOk() (*int32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *Transaction) SetTotalAmount(v int32)`

SetTotalAmount sets TotalAmount field to given value.


### GetCommissionAmount

`func (o *Transaction) GetCommissionAmount() int32`

GetCommissionAmount returns the CommissionAmount field if non-nil, zero value otherwise.

### GetCommissionAmountOk

`func (o *Transaction) GetCommissionAmountOk() (*int32, bool)`

GetCommissionAmountOk returns a tuple with the CommissionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionAmount

`func (o *Transaction) SetCommissionAmount(v int32)`

SetCommissionAmount sets CommissionAmount field to given value.


### GetComment

`func (o *Transaction) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Transaction) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Transaction) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Transaction) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetAutoCommission

`func (o *Transaction) GetAutoCommission() bool`

GetAutoCommission returns the AutoCommission field if non-nil, zero value otherwise.

### GetAutoCommissionOk

`func (o *Transaction) GetAutoCommissionOk() (*bool, bool)`

GetAutoCommissionOk returns a tuple with the AutoCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCommission

`func (o *Transaction) SetAutoCommission(v bool)`

SetAutoCommission sets AutoCommission field to given value.

### HasAutoCommission

`func (o *Transaction) HasAutoCommission() bool`

HasAutoCommission returns a boolean if a field has been set.

### GetIsPreAuth

`func (o *Transaction) GetIsPreAuth() bool`

GetIsPreAuth returns the IsPreAuth field if non-nil, zero value otherwise.

### GetIsPreAuthOk

`func (o *Transaction) GetIsPreAuthOk() (*bool, bool)`

GetIsPreAuthOk returns a tuple with the IsPreAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreAuth

`func (o *Transaction) SetIsPreAuth(v bool)`

SetIsPreAuth sets IsPreAuth field to given value.

### HasIsPreAuth

`func (o *Transaction) HasIsPreAuth() bool`

HasIsPreAuth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


