# AccountBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Payment account ID | [optional] 
**Balance** | Pointer to **int32** | Payment account balance | [optional] 
**Status** | Pointer to **int32** | Payment account status&lt;br/&gt;2 &#x3D; Registered, KYC incomplete.&lt;br/&gt;3 &#x3D; Registered, rejected KYC.&lt;br/&gt;5 &#x3D; Registered, KYC 1 (status given at registration).&lt;br/&gt;6 &#x3D; Registered, KYC 2.&lt;br/&gt;7 &#x3D; Registered, KYC 3.&lt;br/&gt;8 &#x3D; Registered, expired KYC.&lt;br/&gt;10 &#x3D; Blocked.&lt;br/&gt;12 &#x3D; Closed.&lt;br/&gt;13 &#x3D; Registered, status is being updated from KYC 2 to KYC 3.&lt;br/&gt;14 &#x3D; One-time customer.&lt;br/&gt;15 &#x3D; Special account for crowdlending.&lt;br/&gt;16 &#x3D; Technical account.&lt;br/&gt; | [optional] 

## Methods

### NewAccountBalance

`func NewAccountBalance() *AccountBalance`

NewAccountBalance instantiates a new AccountBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountBalanceWithDefaults

`func NewAccountBalanceWithDefaults() *AccountBalance`

NewAccountBalanceWithDefaults instantiates a new AccountBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountBalance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountBalance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountBalance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountBalance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBalance

`func (o *AccountBalance) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *AccountBalance) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *AccountBalance) SetBalance(v int32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *AccountBalance) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetStatus

`func (o *AccountBalance) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountBalance) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountBalance) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountBalance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


