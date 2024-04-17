# EuPagoInit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Transaction ID. You require this ID to confirm the transaction | [optional] 
**Reference** | Pointer to **string** | References in Format: &lt;b&gt;entitades.referencias&lt;/b&gt; | [optional] 

## Methods

### NewEuPagoInit

`func NewEuPagoInit() *EuPagoInit`

NewEuPagoInit instantiates a new EuPagoInit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEuPagoInitWithDefaults

`func NewEuPagoInitWithDefaults() *EuPagoInit`

NewEuPagoInitWithDefaults instantiates a new EuPagoInit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EuPagoInit) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EuPagoInit) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EuPagoInit) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *EuPagoInit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReference

`func (o *EuPagoInit) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *EuPagoInit) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *EuPagoInit) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *EuPagoInit) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


