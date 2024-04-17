# RefundMoneyInInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountToRefund** | Pointer to **int32** | Refund Amount. If empty, the total amount of the payment will be refunded. | [optional] 
**Comment** | Pointer to **string** | Comment on the refund  **Please explain why you refunded** | [optional] 

## Methods

### NewRefundMoneyInInput

`func NewRefundMoneyInInput() *RefundMoneyInInput`

NewRefundMoneyInInput instantiates a new RefundMoneyInInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundMoneyInInputWithDefaults

`func NewRefundMoneyInInputWithDefaults() *RefundMoneyInInput`

NewRefundMoneyInInputWithDefaults instantiates a new RefundMoneyInInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountToRefund

`func (o *RefundMoneyInInput) GetAmountToRefund() int32`

GetAmountToRefund returns the AmountToRefund field if non-nil, zero value otherwise.

### GetAmountToRefundOk

`func (o *RefundMoneyInInput) GetAmountToRefundOk() (*int32, bool)`

GetAmountToRefundOk returns a tuple with the AmountToRefund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountToRefund

`func (o *RefundMoneyInInput) SetAmountToRefund(v int32)`

SetAmountToRefund sets AmountToRefund field to given value.

### HasAmountToRefund

`func (o *RefundMoneyInInput) HasAmountToRefund() bool`

HasAmountToRefund returns a boolean if a field has been set.

### GetComment

`func (o *RefundMoneyInInput) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RefundMoneyInInput) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RefundMoneyInInput) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RefundMoneyInInput) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


