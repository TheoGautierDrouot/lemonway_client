# UploadDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Document ID | [optional] 
**Status** | Pointer to **int32** | Document Status:&lt;br/&gt;0 &#x3D; Document put on hold, waiting for another document.&lt;br/&gt;1 &#x3D; Received, need manual validation.&lt;br/&gt;2 &#x3D; Accepted.&lt;br/&gt;3 &#x3D; Rejected.&lt;br/&gt;4 &#x3D; Rejected. Unreadable by human (Cropped, blur, glare…).&lt;br/&gt;5 &#x3D; Rejected. Expired (Expiration Date is passed).&lt;br/&gt;6 &#x3D; Rejected. Wrong Type (Document not accepted).&lt;br/&gt;7 &#x3D; Rejected. Wrong Name (Name not matching user information).&lt;br/&gt;8 &#x3D; Rejected. Duplicated Document.&lt;br/&gt; | [optional] 
**Substatus** | Pointer to **int32** | Document sub-status. Only when Status &#x3D; 1:&lt;br/&gt;1 &#x3D; Document type not processable automatically.&lt;br/&gt;2 &#x3D; Unreadable by computer (Cropped, blur, glare…).&lt;br/&gt;3 &#x3D; Wrong Type.&lt;br/&gt;4 &#x3D; Wrong Name (Name not matching user information).&lt;br/&gt;5 &#x3D; File too big (File should not be over 4MB to be processed automatically).&lt;br/&gt;6 &#x3D; Automatic Controls Failed.&lt;br/&gt;7 &#x3D; Need 2 sides of the document in one file.&lt;br/&gt;8 &#x3D; Unknown.&lt;br/&gt; | [optional] 
**Accountstatus** | Pointer to **int32** | Payment account status&lt;br/&gt;2 &#x3D; Registered, KYC incomplete.&lt;br/&gt;3 &#x3D; Registered, rejected KYC.&lt;br/&gt;5 &#x3D; Registered, KYC 1 (status given at registration).&lt;br/&gt;6 &#x3D; Registered, KYC 2.&lt;br/&gt;7 &#x3D; Registered, KYC 3.&lt;br/&gt;8 &#x3D; Registered, expired KYC.&lt;br/&gt;10 &#x3D; Blocked.&lt;br/&gt;12 &#x3D; Closed.&lt;br/&gt;13 &#x3D; Registered, status is being updated from KYC 2 to KYC 3.&lt;br/&gt;14 &#x3D; One-time customer.&lt;br/&gt;15 &#x3D; Special account for crowdlending.&lt;br/&gt;16 &#x3D; Technical account.&lt;br/&gt; | [optional] 

## Methods

### NewUploadDocument

`func NewUploadDocument() *UploadDocument`

NewUploadDocument instantiates a new UploadDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadDocumentWithDefaults

`func NewUploadDocumentWithDefaults() *UploadDocument`

NewUploadDocumentWithDefaults instantiates a new UploadDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UploadDocument) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UploadDocument) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UploadDocument) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UploadDocument) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *UploadDocument) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UploadDocument) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UploadDocument) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UploadDocument) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubstatus

`func (o *UploadDocument) GetSubstatus() int32`

GetSubstatus returns the Substatus field if non-nil, zero value otherwise.

### GetSubstatusOk

`func (o *UploadDocument) GetSubstatusOk() (*int32, bool)`

GetSubstatusOk returns a tuple with the Substatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstatus

`func (o *UploadDocument) SetSubstatus(v int32)`

SetSubstatus sets Substatus field to given value.

### HasSubstatus

`func (o *UploadDocument) HasSubstatus() bool`

HasSubstatus returns a boolean if a field has been set.

### GetAccountstatus

`func (o *UploadDocument) GetAccountstatus() int32`

GetAccountstatus returns the Accountstatus field if non-nil, zero value otherwise.

### GetAccountstatusOk

`func (o *UploadDocument) GetAccountstatusOk() (*int32, bool)`

GetAccountstatusOk returns a tuple with the Accountstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountstatus

`func (o *UploadDocument) SetAccountstatus(v int32)`

SetAccountstatus sets Accountstatus field to given value.

### HasAccountstatus

`func (o *UploadDocument) HasAccountstatus() bool`

HasAccountstatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


