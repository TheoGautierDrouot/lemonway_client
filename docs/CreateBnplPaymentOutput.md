# CreateBnplPaymentOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Transaction ID. | [optional] 
**RedirectUrl** | Pointer to **string** | Payment page URL. | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewCreateBnplPaymentOutput

`func NewCreateBnplPaymentOutput() *CreateBnplPaymentOutput`

NewCreateBnplPaymentOutput instantiates a new CreateBnplPaymentOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBnplPaymentOutputWithDefaults

`func NewCreateBnplPaymentOutputWithDefaults() *CreateBnplPaymentOutput`

NewCreateBnplPaymentOutputWithDefaults instantiates a new CreateBnplPaymentOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateBnplPaymentOutput) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateBnplPaymentOutput) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateBnplPaymentOutput) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CreateBnplPaymentOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *CreateBnplPaymentOutput) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *CreateBnplPaymentOutput) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *CreateBnplPaymentOutput) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *CreateBnplPaymentOutput) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetError

`func (o *CreateBnplPaymentOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CreateBnplPaymentOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CreateBnplPaymentOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *CreateBnplPaymentOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


