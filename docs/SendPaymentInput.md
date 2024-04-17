# SendPaymentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DebitAccountId** | **string** | Wallet ID to be Debited | 
**CreditAccountId** | **string** | Wallet ID to Credited | 
**Amount** | Pointer to **int32** | Payment Amount | [optional] 
**Comment** | Pointer to **string** | Payment Comment | [optional] 
**ScheduledDate** | Pointer to **string** | If scheduledDate is set, the following rules will apply:  1. ScheduledDate cannot be equal to the current date, Paris time (CET).  2. The payment will be inserted but not executed, the status will be pending and necessary checks like user balance or user status will not be checked.  3. At 1am, Paris time (CET), on [scheduledDate], the checks will be performed and the payment will be finalized.  **This feature is only available for some partners, contact commercial services for more information.** | [optional] 
**PrivateData** | Pointer to [**PrivateData**](PrivateData.md) |  | [optional] 
**OriginTransactionId** | Pointer to **int32** | Origin Transaction Identification | [optional] 
**Reference** | Pointer to **string** |  | [optional] 

## Methods

### NewSendPaymentInput

`func NewSendPaymentInput(debitAccountId string, creditAccountId string, ) *SendPaymentInput`

NewSendPaymentInput instantiates a new SendPaymentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendPaymentInputWithDefaults

`func NewSendPaymentInputWithDefaults() *SendPaymentInput`

NewSendPaymentInputWithDefaults instantiates a new SendPaymentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDebitAccountId

`func (o *SendPaymentInput) GetDebitAccountId() string`

GetDebitAccountId returns the DebitAccountId field if non-nil, zero value otherwise.

### GetDebitAccountIdOk

`func (o *SendPaymentInput) GetDebitAccountIdOk() (*string, bool)`

GetDebitAccountIdOk returns a tuple with the DebitAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAccountId

`func (o *SendPaymentInput) SetDebitAccountId(v string)`

SetDebitAccountId sets DebitAccountId field to given value.


### GetCreditAccountId

`func (o *SendPaymentInput) GetCreditAccountId() string`

GetCreditAccountId returns the CreditAccountId field if non-nil, zero value otherwise.

### GetCreditAccountIdOk

`func (o *SendPaymentInput) GetCreditAccountIdOk() (*string, bool)`

GetCreditAccountIdOk returns a tuple with the CreditAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccountId

`func (o *SendPaymentInput) SetCreditAccountId(v string)`

SetCreditAccountId sets CreditAccountId field to given value.


### GetAmount

`func (o *SendPaymentInput) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SendPaymentInput) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SendPaymentInput) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *SendPaymentInput) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetComment

`func (o *SendPaymentInput) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SendPaymentInput) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SendPaymentInput) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SendPaymentInput) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetScheduledDate

`func (o *SendPaymentInput) GetScheduledDate() string`

GetScheduledDate returns the ScheduledDate field if non-nil, zero value otherwise.

### GetScheduledDateOk

`func (o *SendPaymentInput) GetScheduledDateOk() (*string, bool)`

GetScheduledDateOk returns a tuple with the ScheduledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDate

`func (o *SendPaymentInput) SetScheduledDate(v string)`

SetScheduledDate sets ScheduledDate field to given value.

### HasScheduledDate

`func (o *SendPaymentInput) HasScheduledDate() bool`

HasScheduledDate returns a boolean if a field has been set.

### GetPrivateData

`func (o *SendPaymentInput) GetPrivateData() PrivateData`

GetPrivateData returns the PrivateData field if non-nil, zero value otherwise.

### GetPrivateDataOk

`func (o *SendPaymentInput) GetPrivateDataOk() (*PrivateData, bool)`

GetPrivateDataOk returns a tuple with the PrivateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateData

`func (o *SendPaymentInput) SetPrivateData(v PrivateData)`

SetPrivateData sets PrivateData field to given value.

### HasPrivateData

`func (o *SendPaymentInput) HasPrivateData() bool`

HasPrivateData returns a boolean if a field has been set.

### GetOriginTransactionId

`func (o *SendPaymentInput) GetOriginTransactionId() int32`

GetOriginTransactionId returns the OriginTransactionId field if non-nil, zero value otherwise.

### GetOriginTransactionIdOk

`func (o *SendPaymentInput) GetOriginTransactionIdOk() (*int32, bool)`

GetOriginTransactionIdOk returns a tuple with the OriginTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginTransactionId

`func (o *SendPaymentInput) SetOriginTransactionId(v int32)`

SetOriginTransactionId sets OriginTransactionId field to given value.

### HasOriginTransactionId

`func (o *SendPaymentInput) HasOriginTransactionId() bool`

HasOriginTransactionId returns a boolean if a field has been set.

### GetReference

`func (o *SendPaymentInput) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *SendPaymentInput) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *SendPaymentInput) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *SendPaymentInput) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


