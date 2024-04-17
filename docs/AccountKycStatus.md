# AccountKycStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Payment account ID | [optional] 
**Status** | Pointer to **int32** | Payment account status&lt;br/&gt;2 &#x3D; Registered, KYC incomplete.&lt;br/&gt;3 &#x3D; Registered, rejected KYC.&lt;br/&gt;5 &#x3D; Registered, KYC 1 (status given at registration).&lt;br/&gt;6 &#x3D; Registered, KYC 2.&lt;br/&gt;7 &#x3D; Registered, KYC 3.&lt;br/&gt;8 &#x3D; Registered, expired KYC.&lt;br/&gt;10 &#x3D; Blocked.&lt;br/&gt;12 &#x3D; Closed.&lt;br/&gt;13 &#x3D; Registered, status is being updated from KYC 2 to KYC 3.&lt;br/&gt;14 &#x3D; One-time customer.&lt;br/&gt;15 &#x3D; Special account for crowdlending.&lt;br/&gt;16 &#x3D; Technical account.&lt;br/&gt; | [optional] 
**Date** | Pointer to **string** | Modification date of the payment account status in Second UTC | [optional] 
**Documents** | Pointer to [**[]Document**](Document.md) |  | [optional] 
**Ibans** | Pointer to [**[]Iban**](Iban.md) |  | [optional] 
**SddMandates** | Pointer to [**[]SddMandate**](SddMandate.md) |  | [optional] 

## Methods

### NewAccountKycStatus

`func NewAccountKycStatus() *AccountKycStatus`

NewAccountKycStatus instantiates a new AccountKycStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountKycStatusWithDefaults

`func NewAccountKycStatusWithDefaults() *AccountKycStatus`

NewAccountKycStatusWithDefaults instantiates a new AccountKycStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountKycStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountKycStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountKycStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountKycStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *AccountKycStatus) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountKycStatus) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountKycStatus) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountKycStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDate

`func (o *AccountKycStatus) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AccountKycStatus) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AccountKycStatus) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *AccountKycStatus) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDocuments

`func (o *AccountKycStatus) GetDocuments() []Document`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *AccountKycStatus) GetDocumentsOk() (*[]Document, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *AccountKycStatus) SetDocuments(v []Document)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *AccountKycStatus) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetIbans

`func (o *AccountKycStatus) GetIbans() []Iban`

GetIbans returns the Ibans field if non-nil, zero value otherwise.

### GetIbansOk

`func (o *AccountKycStatus) GetIbansOk() (*[]Iban, bool)`

GetIbansOk returns a tuple with the Ibans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbans

`func (o *AccountKycStatus) SetIbans(v []Iban)`

SetIbans sets Ibans field to given value.

### HasIbans

`func (o *AccountKycStatus) HasIbans() bool`

HasIbans returns a boolean if a field has been set.

### GetSddMandates

`func (o *AccountKycStatus) GetSddMandates() []SddMandate`

GetSddMandates returns the SddMandates field if non-nil, zero value otherwise.

### GetSddMandatesOk

`func (o *AccountKycStatus) GetSddMandatesOk() (*[]SddMandate, bool)`

GetSddMandatesOk returns a tuple with the SddMandates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSddMandates

`func (o *AccountKycStatus) SetSddMandates(v []SddMandate)`

SetSddMandates sets SddMandates field to given value.

### HasSddMandates

`func (o *AccountKycStatus) HasSddMandates() bool`

HasSddMandates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


