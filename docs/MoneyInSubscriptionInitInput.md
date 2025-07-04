# MoneyInSubscriptionInitInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** | Subscription ID | 
**Count** | **int32** | Number of operations to be performed on this subscription | 
**AccountId** | **string** | Payment Account ID to Credit | 
**TotalAmount** | Pointer to **int32** | Amount to Debit  Amounts are given as integer numbers in cents | [optional] 
**CommissionAmount** | Pointer to **int32** | Your Fee  Amounts are given as integer numbers in cents | [optional] 
**Comment** | Pointer to **string** | Comment Regarding the Transaction | [optional] 
**AutoCommission** | Pointer to **bool** | If true:  1. [amountCom] will be ignored and will be replaced with Lemonway&#39;s fee    2. You will not receive any fee | [optional] 

## Methods

### NewMoneyInSubscriptionInitInput

`func NewMoneyInSubscriptionInitInput(subscriptionId string, count int32, accountId string, ) *MoneyInSubscriptionInitInput`

NewMoneyInSubscriptionInitInput instantiates a new MoneyInSubscriptionInitInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoneyInSubscriptionInitInputWithDefaults

`func NewMoneyInSubscriptionInitInputWithDefaults() *MoneyInSubscriptionInitInput`

NewMoneyInSubscriptionInitInputWithDefaults instantiates a new MoneyInSubscriptionInitInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *MoneyInSubscriptionInitInput) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *MoneyInSubscriptionInitInput) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *MoneyInSubscriptionInitInput) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetCount

`func (o *MoneyInSubscriptionInitInput) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MoneyInSubscriptionInitInput) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MoneyInSubscriptionInitInput) SetCount(v int32)`

SetCount sets Count field to given value.


### GetAccountId

`func (o *MoneyInSubscriptionInitInput) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *MoneyInSubscriptionInitInput) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *MoneyInSubscriptionInitInput) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetTotalAmount

`func (o *MoneyInSubscriptionInitInput) GetTotalAmount() int32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *MoneyInSubscriptionInitInput) GetTotalAmountOk() (*int32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *MoneyInSubscriptionInitInput) SetTotalAmount(v int32)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *MoneyInSubscriptionInitInput) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetCommissionAmount

`func (o *MoneyInSubscriptionInitInput) GetCommissionAmount() int32`

GetCommissionAmount returns the CommissionAmount field if non-nil, zero value otherwise.

### GetCommissionAmountOk

`func (o *MoneyInSubscriptionInitInput) GetCommissionAmountOk() (*int32, bool)`

GetCommissionAmountOk returns a tuple with the CommissionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionAmount

`func (o *MoneyInSubscriptionInitInput) SetCommissionAmount(v int32)`

SetCommissionAmount sets CommissionAmount field to given value.

### HasCommissionAmount

`func (o *MoneyInSubscriptionInitInput) HasCommissionAmount() bool`

HasCommissionAmount returns a boolean if a field has been set.

### GetComment

`func (o *MoneyInSubscriptionInitInput) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MoneyInSubscriptionInitInput) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MoneyInSubscriptionInitInput) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MoneyInSubscriptionInitInput) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetAutoCommission

`func (o *MoneyInSubscriptionInitInput) GetAutoCommission() bool`

GetAutoCommission returns the AutoCommission field if non-nil, zero value otherwise.

### GetAutoCommissionOk

`func (o *MoneyInSubscriptionInitInput) GetAutoCommissionOk() (*bool, bool)`

GetAutoCommissionOk returns a tuple with the AutoCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCommission

`func (o *MoneyInSubscriptionInitInput) SetAutoCommission(v bool)`

SetAutoCommission sets AutoCommission field to given value.

### HasAutoCommission

`func (o *MoneyInSubscriptionInitInput) HasAutoCommission() bool`

HasAutoCommission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


