# UploadDocumentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Document Name | 
**Type** | **int32** | Document Type:    **Note:**     If you have previously uploaded a document in a reserved **slot(0-13)** and need to upload another document of the same type, use the slot **Other document(6, 14-20)**.&lt;br/&gt;0 &#x3D; ID card (both sides in one file).&lt;br/&gt;1 &#x3D; Proof of address.&lt;br/&gt;2 &#x3D; Scan of a proof of IBAN.&lt;br/&gt;3 &#x3D; Passport (European Union).&lt;br/&gt;4 &#x3D; Passport (outside the European Union).&lt;br/&gt;5 &#x3D; Residence permit (both sides in one file).&lt;br/&gt;6 &#x3D; Other document type.&lt;br/&gt;7 &#x3D; Official company registration document (Kbis extract or equivalent).&lt;br/&gt;11 &#x3D; Driver licence (both sides in one file).&lt;br/&gt;12 &#x3D; Status.&lt;br/&gt;13 &#x3D; Selfie.&lt;br/&gt;14 &#x3D; Other document type.&lt;br/&gt;15 &#x3D; Other document type.&lt;br/&gt;16 &#x3D; Other document type.&lt;br/&gt;17 &#x3D; Other document type.&lt;br/&gt;18 &#x3D; Other document type.&lt;br/&gt;19 &#x3D; Other document type.&lt;br/&gt;20 &#x3D; Other document type.&lt;br/&gt;21 &#x3D; SDD mandate.&lt;br/&gt; | 
**Buffer** | **string** | Byte array with the document. Encode in base 64 if necessary. | 
**SddMandateId** | Pointer to **int64** | Lets you upload your signed(with your own signing partner) mandate document, to validate a mandate ID you previously created with &#x60;RegisterSddMandate&#x60;. | [optional] 

## Methods

### NewUploadDocumentInput

`func NewUploadDocumentInput(name string, type_ int32, buffer string, ) *UploadDocumentInput`

NewUploadDocumentInput instantiates a new UploadDocumentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadDocumentInputWithDefaults

`func NewUploadDocumentInputWithDefaults() *UploadDocumentInput`

NewUploadDocumentInputWithDefaults instantiates a new UploadDocumentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UploadDocumentInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UploadDocumentInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UploadDocumentInput) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *UploadDocumentInput) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UploadDocumentInput) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UploadDocumentInput) SetType(v int32)`

SetType sets Type field to given value.


### GetBuffer

`func (o *UploadDocumentInput) GetBuffer() string`

GetBuffer returns the Buffer field if non-nil, zero value otherwise.

### GetBufferOk

`func (o *UploadDocumentInput) GetBufferOk() (*string, bool)`

GetBufferOk returns a tuple with the Buffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuffer

`func (o *UploadDocumentInput) SetBuffer(v string)`

SetBuffer sets Buffer field to given value.


### GetSddMandateId

`func (o *UploadDocumentInput) GetSddMandateId() int64`

GetSddMandateId returns the SddMandateId field if non-nil, zero value otherwise.

### GetSddMandateIdOk

`func (o *UploadDocumentInput) GetSddMandateIdOk() (*int64, bool)`

GetSddMandateIdOk returns a tuple with the SddMandateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSddMandateId

`func (o *UploadDocumentInput) SetSddMandateId(v int64)`

SetSddMandateId sets SddMandateId field to given value.

### HasSddMandateId

`func (o *UploadDocumentInput) HasSddMandateId() bool`

HasSddMandateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


