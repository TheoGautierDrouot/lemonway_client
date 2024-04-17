# UpdateAccountStatusInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int32** | New payment account status&lt;br/&gt;0 &#x3D; Unknown.&lt;br/&gt;1 &#x3D; Not Opened.&lt;br/&gt;2 &#x3D; Opened, need more documents.&lt;br/&gt;3 &#x3D; Opened, document rejected.&lt;br/&gt;5 &#x3D; Opened, KYC1.&lt;br/&gt;6 &#x3D; Opened, KYC2.&lt;br/&gt;7 &#x3D; Opened, KYC3.&lt;br/&gt;8 &#x3D; Opened, document expired.&lt;br/&gt;9 &#x3D; Frozen (by backoffice).&lt;br/&gt;10 &#x3D; Blocked.&lt;br/&gt;11 &#x3D; Locked (by Web Service).&lt;br/&gt;12 &#x3D; Closed.&lt;br/&gt;13 &#x3D; Pending KYC3.&lt;br/&gt;14 &#x3D; One-time customer.&lt;br/&gt;15 &#x3D; CGE.&lt;br/&gt;16 &#x3D; Technical Payment Account.&lt;br/&gt; | 

## Methods

### NewUpdateAccountStatusInput

`func NewUpdateAccountStatusInput(status int32, ) *UpdateAccountStatusInput`

NewUpdateAccountStatusInput instantiates a new UpdateAccountStatusInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountStatusInputWithDefaults

`func NewUpdateAccountStatusInputWithDefaults() *UpdateAccountStatusInput`

NewUpdateAccountStatusInputWithDefaults instantiates a new UpdateAccountStatusInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateAccountStatusInput) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateAccountStatusInput) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateAccountStatusInput) SetStatus(v int32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


