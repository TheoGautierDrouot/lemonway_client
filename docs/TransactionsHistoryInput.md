# TransactionsHistoryInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **string** | UTC Unix timestamp  In order to return transactions initialized after &#x60;startDate&#x60;.  If the payment account is a **SC Wallet** then this value is mandatory. | [optional] 
**EndDate** | Pointer to **string** | UTC Unix timestamp  In order to return transactions initialized before &#x60;endDate&#x60;.  If the payment account is a **SC Wallet** then this value is mandatory and the time span can not exceed 1 week. | [optional] 
**ExecutionDateStart** | Pointer to **string** | UTC Unix timestamp  In order to return transactions Executed after &#x60;endDate&#x60;.  If the payment account is a **SC Wallet** then this value is mandatory and the time span can not exceed 1 week. | [optional] 
**ExecutionDateEnd** | Pointer to **string** | UTC Unix timestamp  In order to return transactions Executed before &#x60;endDate&#x60;  If the payment account is a **SC Wallet** then this value is mandatory and the time span can not exceed 1 week. | [optional] 

## Methods

### NewTransactionsHistoryInput

`func NewTransactionsHistoryInput() *TransactionsHistoryInput`

NewTransactionsHistoryInput instantiates a new TransactionsHistoryInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionsHistoryInputWithDefaults

`func NewTransactionsHistoryInputWithDefaults() *TransactionsHistoryInput`

NewTransactionsHistoryInputWithDefaults instantiates a new TransactionsHistoryInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *TransactionsHistoryInput) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TransactionsHistoryInput) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TransactionsHistoryInput) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TransactionsHistoryInput) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *TransactionsHistoryInput) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TransactionsHistoryInput) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TransactionsHistoryInput) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *TransactionsHistoryInput) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetExecutionDateStart

`func (o *TransactionsHistoryInput) GetExecutionDateStart() string`

GetExecutionDateStart returns the ExecutionDateStart field if non-nil, zero value otherwise.

### GetExecutionDateStartOk

`func (o *TransactionsHistoryInput) GetExecutionDateStartOk() (*string, bool)`

GetExecutionDateStartOk returns a tuple with the ExecutionDateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDateStart

`func (o *TransactionsHistoryInput) SetExecutionDateStart(v string)`

SetExecutionDateStart sets ExecutionDateStart field to given value.

### HasExecutionDateStart

`func (o *TransactionsHistoryInput) HasExecutionDateStart() bool`

HasExecutionDateStart returns a boolean if a field has been set.

### GetExecutionDateEnd

`func (o *TransactionsHistoryInput) GetExecutionDateEnd() string`

GetExecutionDateEnd returns the ExecutionDateEnd field if non-nil, zero value otherwise.

### GetExecutionDateEndOk

`func (o *TransactionsHistoryInput) GetExecutionDateEndOk() (*string, bool)`

GetExecutionDateEndOk returns a tuple with the ExecutionDateEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDateEnd

`func (o *TransactionsHistoryInput) SetExecutionDateEnd(v string)`

SetExecutionDateEnd sets ExecutionDateEnd field to given value.

### HasExecutionDateEnd

`func (o *TransactionsHistoryInput) HasExecutionDateEnd() bool`

HasExecutionDateEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


