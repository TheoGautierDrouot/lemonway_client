# Document

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Document ID | [optional] 
**Status** | Pointer to **int32** | Document Status&lt;br/&gt;0 &#x3D; Document put on hold, waiting for another document.&lt;br/&gt;1 &#x3D; Received, need manual validation.&lt;br/&gt;2 &#x3D; Accepted.&lt;br/&gt;3 &#x3D; Rejected.&lt;br/&gt;4 &#x3D; Rejected. Unreadable by human (Cropped, blur, glare…).&lt;br/&gt;5 &#x3D; Rejected. Expired (Expiration Date is passed).&lt;br/&gt;6 &#x3D; Rejected. Wrong Type (Document not accepted).&lt;br/&gt;7 &#x3D; Rejected. Wrong Name (Name not matching user information).&lt;br/&gt;8 &#x3D; Rejected. Duplicated Document.&lt;br/&gt; | [optional] 
**Type** | Pointer to **int32** | Document Type&lt;br/&gt;0 &#x3D; ID card (both sides in one file).&lt;br/&gt;1 &#x3D; Proof of address.&lt;br/&gt;2 &#x3D; Scan of a proof of IBAN.&lt;br/&gt;3 &#x3D; Passport (European Union).&lt;br/&gt;4 &#x3D; Passport (outside the European Union).&lt;br/&gt;5 &#x3D; Residence permit (both sides in one file).&lt;br/&gt;6 &#x3D; Other document type.&lt;br/&gt;7 &#x3D; Official company registration document (Kbis extract or equivalent).&lt;br/&gt;11 &#x3D; Driver licence (both sides in one file).&lt;br/&gt;12 &#x3D; Status.&lt;br/&gt;13 &#x3D; Selfie.&lt;br/&gt;14 &#x3D; Other document type.&lt;br/&gt;15 &#x3D; Other document type.&lt;br/&gt;16 &#x3D; Other document type.&lt;br/&gt;17 &#x3D; Other document type.&lt;br/&gt;18 &#x3D; Other document type.&lt;br/&gt;19 &#x3D; Other document type.&lt;br/&gt;20 &#x3D; Other document type.&lt;br/&gt;21 &#x3D; SDD mandate.&lt;br/&gt; | [optional] 
**ValidityDate** | Pointer to **string** | Document validity date  Format: dd/mm/yyyy  Empty if unknown | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewDocument

`func NewDocument() *Document`

NewDocument instantiates a new Document object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentWithDefaults

`func NewDocumentWithDefaults() *Document`

NewDocumentWithDefaults instantiates a new Document object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Document) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Document) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Document) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Document) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *Document) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Document) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Document) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Document) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *Document) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Document) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Document) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *Document) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValidityDate

`func (o *Document) GetValidityDate() string`

GetValidityDate returns the ValidityDate field if non-nil, zero value otherwise.

### GetValidityDateOk

`func (o *Document) GetValidityDateOk() (*string, bool)`

GetValidityDateOk returns a tuple with the ValidityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityDate

`func (o *Document) SetValidityDate(v string)`

SetValidityDate sets ValidityDate field to given value.

### HasValidityDate

`func (o *Document) HasValidityDate() bool`

HasValidityDate returns a boolean if a field has been set.

### GetComment

`func (o *Document) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Document) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Document) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Document) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


