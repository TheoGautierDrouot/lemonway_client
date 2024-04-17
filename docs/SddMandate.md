# SddMandate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Mandate ID | [optional] 
**Status** | Pointer to **int32** | Mandate Status&lt;br/&gt;0 &#x3D; Pending validation.&lt;br/&gt;5 &#x3D; Validated, First SDD to be done.&lt;br/&gt;6 &#x3D; Enabled.&lt;br/&gt;8 &#x3D; Disabled.&lt;br/&gt;9 &#x3D; Rejected.&lt;br/&gt; | [optional] 
**IBAN** | Pointer to **string** | IBAN associated to the mandate | [optional] 
**BIC** | Pointer to **string** | SWIFT code associated to the mandate | [optional] 
**Holder** | Pointer to **string** | IBAN Holder | [optional] 
**UniqueMandateReference** | Pointer to **string** | Unique mandate reference | [optional] 

## Methods

### NewSddMandate

`func NewSddMandate() *SddMandate`

NewSddMandate instantiates a new SddMandate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSddMandateWithDefaults

`func NewSddMandateWithDefaults() *SddMandate`

NewSddMandateWithDefaults instantiates a new SddMandate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SddMandate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SddMandate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SddMandate) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SddMandate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *SddMandate) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SddMandate) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SddMandate) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SddMandate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIBAN

`func (o *SddMandate) GetIBAN() string`

GetIBAN returns the IBAN field if non-nil, zero value otherwise.

### GetIBANOk

`func (o *SddMandate) GetIBANOk() (*string, bool)`

GetIBANOk returns a tuple with the IBAN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBAN

`func (o *SddMandate) SetIBAN(v string)`

SetIBAN sets IBAN field to given value.

### HasIBAN

`func (o *SddMandate) HasIBAN() bool`

HasIBAN returns a boolean if a field has been set.

### GetBIC

`func (o *SddMandate) GetBIC() string`

GetBIC returns the BIC field if non-nil, zero value otherwise.

### GetBICOk

`func (o *SddMandate) GetBICOk() (*string, bool)`

GetBICOk returns a tuple with the BIC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBIC

`func (o *SddMandate) SetBIC(v string)`

SetBIC sets BIC field to given value.

### HasBIC

`func (o *SddMandate) HasBIC() bool`

HasBIC returns a boolean if a field has been set.

### GetHolder

`func (o *SddMandate) GetHolder() string`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *SddMandate) GetHolderOk() (*string, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *SddMandate) SetHolder(v string)`

SetHolder sets Holder field to given value.

### HasHolder

`func (o *SddMandate) HasHolder() bool`

HasHolder returns a boolean if a field has been set.

### GetUniqueMandateReference

`func (o *SddMandate) GetUniqueMandateReference() string`

GetUniqueMandateReference returns the UniqueMandateReference field if non-nil, zero value otherwise.

### GetUniqueMandateReferenceOk

`func (o *SddMandate) GetUniqueMandateReferenceOk() (*string, bool)`

GetUniqueMandateReferenceOk returns a tuple with the UniqueMandateReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueMandateReference

`func (o *SddMandate) SetUniqueMandateReference(v string)`

SetUniqueMandateReference sets UniqueMandateReference field to given value.

### HasUniqueMandateReference

`func (o *SddMandate) HasUniqueMandateReference() bool`

HasUniqueMandateReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


