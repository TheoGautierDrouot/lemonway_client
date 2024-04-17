# EnrolmentInitInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **int32** | Identification provider&lt;br/&gt;0 &#x3D; Deutsche Post POSTIDENT.&lt;br/&gt; | 
**SuccessUrl** | Pointer to **string** | Forwarding URL to which in case of a successful identification should be forwarded | [optional] 
**DeclinedUrl** | Pointer to **string** | Forwarding URL to which in case of an unsuccessful identification should be forwarded | [optional] 
**CouponRequestedUrl** | Pointer to **string** | Forwarding URL to forward to when the coupon has been downloaded | [optional] 
**ReviewPendingUrl** | Pointer to **string** | Forwarding URL to forward to when the coupon has been downloaded | [optional] 
**IncompleteUrl** | Pointer to **string** | Forwarding URL to which in case of incomplete identification should be forwarded | [optional] 

## Methods

### NewEnrolmentInitInput

`func NewEnrolmentInitInput(provider int32, ) *EnrolmentInitInput`

NewEnrolmentInitInput instantiates a new EnrolmentInitInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrolmentInitInputWithDefaults

`func NewEnrolmentInitInputWithDefaults() *EnrolmentInitInput`

NewEnrolmentInitInputWithDefaults instantiates a new EnrolmentInitInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *EnrolmentInitInput) GetProvider() int32`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EnrolmentInitInput) GetProviderOk() (*int32, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EnrolmentInitInput) SetProvider(v int32)`

SetProvider sets Provider field to given value.


### GetSuccessUrl

`func (o *EnrolmentInitInput) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *EnrolmentInitInput) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *EnrolmentInitInput) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *EnrolmentInitInput) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.

### GetDeclinedUrl

`func (o *EnrolmentInitInput) GetDeclinedUrl() string`

GetDeclinedUrl returns the DeclinedUrl field if non-nil, zero value otherwise.

### GetDeclinedUrlOk

`func (o *EnrolmentInitInput) GetDeclinedUrlOk() (*string, bool)`

GetDeclinedUrlOk returns a tuple with the DeclinedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclinedUrl

`func (o *EnrolmentInitInput) SetDeclinedUrl(v string)`

SetDeclinedUrl sets DeclinedUrl field to given value.

### HasDeclinedUrl

`func (o *EnrolmentInitInput) HasDeclinedUrl() bool`

HasDeclinedUrl returns a boolean if a field has been set.

### GetCouponRequestedUrl

`func (o *EnrolmentInitInput) GetCouponRequestedUrl() string`

GetCouponRequestedUrl returns the CouponRequestedUrl field if non-nil, zero value otherwise.

### GetCouponRequestedUrlOk

`func (o *EnrolmentInitInput) GetCouponRequestedUrlOk() (*string, bool)`

GetCouponRequestedUrlOk returns a tuple with the CouponRequestedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponRequestedUrl

`func (o *EnrolmentInitInput) SetCouponRequestedUrl(v string)`

SetCouponRequestedUrl sets CouponRequestedUrl field to given value.

### HasCouponRequestedUrl

`func (o *EnrolmentInitInput) HasCouponRequestedUrl() bool`

HasCouponRequestedUrl returns a boolean if a field has been set.

### GetReviewPendingUrl

`func (o *EnrolmentInitInput) GetReviewPendingUrl() string`

GetReviewPendingUrl returns the ReviewPendingUrl field if non-nil, zero value otherwise.

### GetReviewPendingUrlOk

`func (o *EnrolmentInitInput) GetReviewPendingUrlOk() (*string, bool)`

GetReviewPendingUrlOk returns a tuple with the ReviewPendingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewPendingUrl

`func (o *EnrolmentInitInput) SetReviewPendingUrl(v string)`

SetReviewPendingUrl sets ReviewPendingUrl field to given value.

### HasReviewPendingUrl

`func (o *EnrolmentInitInput) HasReviewPendingUrl() bool`

HasReviewPendingUrl returns a boolean if a field has been set.

### GetIncompleteUrl

`func (o *EnrolmentInitInput) GetIncompleteUrl() string`

GetIncompleteUrl returns the IncompleteUrl field if non-nil, zero value otherwise.

### GetIncompleteUrlOk

`func (o *EnrolmentInitInput) GetIncompleteUrlOk() (*string, bool)`

GetIncompleteUrlOk returns a tuple with the IncompleteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompleteUrl

`func (o *EnrolmentInitInput) SetIncompleteUrl(v string)`

SetIncompleteUrl sets IncompleteUrl field to given value.

### HasIncompleteUrl

`func (o *EnrolmentInitInput) HasIncompleteUrl() bool`

HasIncompleteUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


