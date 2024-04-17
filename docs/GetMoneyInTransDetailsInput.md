# GetMoneyInTransDetailsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | Pointer to **int64** | Money-In ID | [optional] 
**TransactionComment** | Pointer to **string** | Money-In Comment | [optional] 
**TransactionMerchantToken** | Pointer to **string** | Token from wkToken variable | [optional] 
**StartDate** | Pointer to **string** | UTC Unix timestamp, in order to return transactions initialized after it | [optional] 
**EndDate** | Pointer to **string** | UTC Unix timestamp, in order to return transactions initialized before it | [optional] 

## Methods

### NewGetMoneyInTransDetailsInput

`func NewGetMoneyInTransDetailsInput() *GetMoneyInTransDetailsInput`

NewGetMoneyInTransDetailsInput instantiates a new GetMoneyInTransDetailsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMoneyInTransDetailsInputWithDefaults

`func NewGetMoneyInTransDetailsInputWithDefaults() *GetMoneyInTransDetailsInput`

NewGetMoneyInTransDetailsInputWithDefaults instantiates a new GetMoneyInTransDetailsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *GetMoneyInTransDetailsInput) GetTransactionId() int64`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *GetMoneyInTransDetailsInput) GetTransactionIdOk() (*int64, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *GetMoneyInTransDetailsInput) SetTransactionId(v int64)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *GetMoneyInTransDetailsInput) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTransactionComment

`func (o *GetMoneyInTransDetailsInput) GetTransactionComment() string`

GetTransactionComment returns the TransactionComment field if non-nil, zero value otherwise.

### GetTransactionCommentOk

`func (o *GetMoneyInTransDetailsInput) GetTransactionCommentOk() (*string, bool)`

GetTransactionCommentOk returns a tuple with the TransactionComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionComment

`func (o *GetMoneyInTransDetailsInput) SetTransactionComment(v string)`

SetTransactionComment sets TransactionComment field to given value.

### HasTransactionComment

`func (o *GetMoneyInTransDetailsInput) HasTransactionComment() bool`

HasTransactionComment returns a boolean if a field has been set.

### GetTransactionMerchantToken

`func (o *GetMoneyInTransDetailsInput) GetTransactionMerchantToken() string`

GetTransactionMerchantToken returns the TransactionMerchantToken field if non-nil, zero value otherwise.

### GetTransactionMerchantTokenOk

`func (o *GetMoneyInTransDetailsInput) GetTransactionMerchantTokenOk() (*string, bool)`

GetTransactionMerchantTokenOk returns a tuple with the TransactionMerchantToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionMerchantToken

`func (o *GetMoneyInTransDetailsInput) SetTransactionMerchantToken(v string)`

SetTransactionMerchantToken sets TransactionMerchantToken field to given value.

### HasTransactionMerchantToken

`func (o *GetMoneyInTransDetailsInput) HasTransactionMerchantToken() bool`

HasTransactionMerchantToken returns a boolean if a field has been set.

### GetStartDate

`func (o *GetMoneyInTransDetailsInput) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetMoneyInTransDetailsInput) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetMoneyInTransDetailsInput) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetMoneyInTransDetailsInput) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetMoneyInTransDetailsInput) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetMoneyInTransDetailsInput) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetMoneyInTransDetailsInput) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetMoneyInTransDetailsInput) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


