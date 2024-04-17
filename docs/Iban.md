# Iban

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | IBAN ID | [optional] 
**Status** | Pointer to **int32** | Only status 5 allows the use of the IBAN for a Money-Out.&lt;br/&gt;1 &#x3D; None.&lt;br/&gt;2 &#x3D; Internal.&lt;br/&gt;3 &#x3D; Not used.&lt;br/&gt;4 &#x3D; Waiting to be verified by Lemon Way.&lt;br/&gt;5 &#x3D; Activated.&lt;br/&gt;6 &#x3D; Rejected by the bank.&lt;br/&gt;7 &#x3D; Rejected, no owner.&lt;br/&gt;8 &#x3D; Deactivated.&lt;br/&gt;9 &#x3D; Rejected.&lt;br/&gt; | [optional] 
**Iban** | Pointer to **string** | IBAN saved | [optional] 
**Swift** | Pointer to **string** | SWIFT BIC code associated with the IBAN | [optional] 
**Holder** | Pointer to **string** | IBAN Holder | [optional] 
**Type** | Pointer to **int32** | Indicates if a merchant IBAN or a virtual client IBAN  1: Merchant IBAN  2: IBAN virtual client | [optional] 

## Methods

### NewIban

`func NewIban() *Iban`

NewIban instantiates a new Iban object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIbanWithDefaults

`func NewIbanWithDefaults() *Iban`

NewIbanWithDefaults instantiates a new Iban object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Iban) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Iban) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Iban) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Iban) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *Iban) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Iban) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Iban) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Iban) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIban

`func (o *Iban) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *Iban) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *Iban) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *Iban) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetSwift

`func (o *Iban) GetSwift() string`

GetSwift returns the Swift field if non-nil, zero value otherwise.

### GetSwiftOk

`func (o *Iban) GetSwiftOk() (*string, bool)`

GetSwiftOk returns a tuple with the Swift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwift

`func (o *Iban) SetSwift(v string)`

SetSwift sets Swift field to given value.

### HasSwift

`func (o *Iban) HasSwift() bool`

HasSwift returns a boolean if a field has been set.

### GetHolder

`func (o *Iban) GetHolder() string`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *Iban) GetHolderOk() (*string, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *Iban) SetHolder(v string)`

SetHolder sets Holder field to given value.

### HasHolder

`func (o *Iban) HasHolder() bool`

HasHolder returns a boolean if a field has been set.

### GetType

`func (o *Iban) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Iban) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Iban) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *Iban) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


