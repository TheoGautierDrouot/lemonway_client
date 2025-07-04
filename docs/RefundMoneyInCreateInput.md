# RefundMoneyInCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkedP2pIds** | **[]int64** | P2P transactions linked to Money-In refund. | 
**TransactionDateTime** | Pointer to **string** | Transaction DateTime (UTC Unix timestamp) | [optional] 
**AmountToRefund** | Pointer to **int32** | Refund Amount. If empty, the total amount of the payment will be refunded. | [optional] 
**Comment** | Pointer to **string** | Comment on the refund.   **Explain the reason for the refunded amount**   **Note:** In the API Response displayed before the comment will appear the refund transaction id. Example: comment\&quot;: \&quot;Refund 2763789 Items not wanted\&quot; | [optional] 

## Methods

### NewRefundMoneyInCreateInput

`func NewRefundMoneyInCreateInput(linkedP2pIds []int64, ) *RefundMoneyInCreateInput`

NewRefundMoneyInCreateInput instantiates a new RefundMoneyInCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundMoneyInCreateInputWithDefaults

`func NewRefundMoneyInCreateInputWithDefaults() *RefundMoneyInCreateInput`

NewRefundMoneyInCreateInputWithDefaults instantiates a new RefundMoneyInCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkedP2pIds

`func (o *RefundMoneyInCreateInput) GetLinkedP2pIds() []int64`

GetLinkedP2pIds returns the LinkedP2pIds field if non-nil, zero value otherwise.

### GetLinkedP2pIdsOk

`func (o *RefundMoneyInCreateInput) GetLinkedP2pIdsOk() (*[]int64, bool)`

GetLinkedP2pIdsOk returns a tuple with the LinkedP2pIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedP2pIds

`func (o *RefundMoneyInCreateInput) SetLinkedP2pIds(v []int64)`

SetLinkedP2pIds sets LinkedP2pIds field to given value.


### GetTransactionDateTime

`func (o *RefundMoneyInCreateInput) GetTransactionDateTime() string`

GetTransactionDateTime returns the TransactionDateTime field if non-nil, zero value otherwise.

### GetTransactionDateTimeOk

`func (o *RefundMoneyInCreateInput) GetTransactionDateTimeOk() (*string, bool)`

GetTransactionDateTimeOk returns a tuple with the TransactionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDateTime

`func (o *RefundMoneyInCreateInput) SetTransactionDateTime(v string)`

SetTransactionDateTime sets TransactionDateTime field to given value.

### HasTransactionDateTime

`func (o *RefundMoneyInCreateInput) HasTransactionDateTime() bool`

HasTransactionDateTime returns a boolean if a field has been set.

### GetAmountToRefund

`func (o *RefundMoneyInCreateInput) GetAmountToRefund() int32`

GetAmountToRefund returns the AmountToRefund field if non-nil, zero value otherwise.

### GetAmountToRefundOk

`func (o *RefundMoneyInCreateInput) GetAmountToRefundOk() (*int32, bool)`

GetAmountToRefundOk returns a tuple with the AmountToRefund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountToRefund

`func (o *RefundMoneyInCreateInput) SetAmountToRefund(v int32)`

SetAmountToRefund sets AmountToRefund field to given value.

### HasAmountToRefund

`func (o *RefundMoneyInCreateInput) HasAmountToRefund() bool`

HasAmountToRefund returns a boolean if a field has been set.

### GetComment

`func (o *RefundMoneyInCreateInput) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RefundMoneyInCreateInput) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RefundMoneyInCreateInput) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RefundMoneyInCreateInput) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


