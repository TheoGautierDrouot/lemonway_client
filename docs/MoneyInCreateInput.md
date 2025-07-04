# MoneyInCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Card** | [**CardInfoExtended**](CardInfoExtended.md) |  | 
**Transaction** | [**TransactionInfo**](TransactionInfo.md) |  | 
**ServiceProvider** | Pointer to **string** | Service Provider | [optional] 
**AccountId** | **string** | Payment Account ID to Credit | 
**TotalAmount** | Pointer to **int32** | Amount to Debit  Amounts are given as integer numbers in cents | [optional] 
**CommissionAmount** | Pointer to **int32** | Your Fee  Amounts are given as integer numbers in cents | [optional] 
**Comment** | Pointer to **string** | Comment Regarding the Transaction | [optional] 
**AutoCommission** | Pointer to **bool** | If true:  1. [amountCom] will be ignored and will be replaced with Lemonway&#39;s fee    2. You will not receive any fee | [optional] 

## Methods

### NewMoneyInCreateInput

`func NewMoneyInCreateInput(card CardInfoExtended, transaction TransactionInfo, accountId string, ) *MoneyInCreateInput`

NewMoneyInCreateInput instantiates a new MoneyInCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoneyInCreateInputWithDefaults

`func NewMoneyInCreateInputWithDefaults() *MoneyInCreateInput`

NewMoneyInCreateInputWithDefaults instantiates a new MoneyInCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCard

`func (o *MoneyInCreateInput) GetCard() CardInfoExtended`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *MoneyInCreateInput) GetCardOk() (*CardInfoExtended, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *MoneyInCreateInput) SetCard(v CardInfoExtended)`

SetCard sets Card field to given value.


### GetTransaction

`func (o *MoneyInCreateInput) GetTransaction() TransactionInfo`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *MoneyInCreateInput) GetTransactionOk() (*TransactionInfo, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *MoneyInCreateInput) SetTransaction(v TransactionInfo)`

SetTransaction sets Transaction field to given value.


### GetServiceProvider

`func (o *MoneyInCreateInput) GetServiceProvider() string`

GetServiceProvider returns the ServiceProvider field if non-nil, zero value otherwise.

### GetServiceProviderOk

`func (o *MoneyInCreateInput) GetServiceProviderOk() (*string, bool)`

GetServiceProviderOk returns a tuple with the ServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvider

`func (o *MoneyInCreateInput) SetServiceProvider(v string)`

SetServiceProvider sets ServiceProvider field to given value.

### HasServiceProvider

`func (o *MoneyInCreateInput) HasServiceProvider() bool`

HasServiceProvider returns a boolean if a field has been set.

### GetAccountId

`func (o *MoneyInCreateInput) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *MoneyInCreateInput) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *MoneyInCreateInput) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetTotalAmount

`func (o *MoneyInCreateInput) GetTotalAmount() int32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *MoneyInCreateInput) GetTotalAmountOk() (*int32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *MoneyInCreateInput) SetTotalAmount(v int32)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *MoneyInCreateInput) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetCommissionAmount

`func (o *MoneyInCreateInput) GetCommissionAmount() int32`

GetCommissionAmount returns the CommissionAmount field if non-nil, zero value otherwise.

### GetCommissionAmountOk

`func (o *MoneyInCreateInput) GetCommissionAmountOk() (*int32, bool)`

GetCommissionAmountOk returns a tuple with the CommissionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionAmount

`func (o *MoneyInCreateInput) SetCommissionAmount(v int32)`

SetCommissionAmount sets CommissionAmount field to given value.

### HasCommissionAmount

`func (o *MoneyInCreateInput) HasCommissionAmount() bool`

HasCommissionAmount returns a boolean if a field has been set.

### GetComment

`func (o *MoneyInCreateInput) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MoneyInCreateInput) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MoneyInCreateInput) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MoneyInCreateInput) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetAutoCommission

`func (o *MoneyInCreateInput) GetAutoCommission() bool`

GetAutoCommission returns the AutoCommission field if non-nil, zero value otherwise.

### GetAutoCommissionOk

`func (o *MoneyInCreateInput) GetAutoCommissionOk() (*bool, bool)`

GetAutoCommissionOk returns a tuple with the AutoCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCommission

`func (o *MoneyInCreateInput) SetAutoCommission(v bool)`

SetAutoCommission sets AutoCommission field to given value.

### HasAutoCommission

`func (o *MoneyInCreateInput) HasAutoCommission() bool`

HasAutoCommission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


