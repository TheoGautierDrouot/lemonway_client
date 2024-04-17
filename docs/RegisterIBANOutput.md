# RegisterIBANOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IbanId** | Pointer to **int64** | IBAN ID | [optional] 
**Status** | Pointer to **int32** | IBAN Status&lt;br/&gt;1 &#x3D; None.&lt;br/&gt;2 &#x3D; Internal.&lt;br/&gt;3 &#x3D; Not used.&lt;br/&gt;4 &#x3D; Waiting to be verified by Lemon Way.&lt;br/&gt;5 &#x3D; Activated.&lt;br/&gt;6 &#x3D; Rejected by the bank.&lt;br/&gt;7 &#x3D; Rejected, no owner.&lt;br/&gt;8 &#x3D; Deactivated.&lt;br/&gt;9 &#x3D; Rejected.&lt;br/&gt; | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewRegisterIBANOutput

`func NewRegisterIBANOutput() *RegisterIBANOutput`

NewRegisterIBANOutput instantiates a new RegisterIBANOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterIBANOutputWithDefaults

`func NewRegisterIBANOutputWithDefaults() *RegisterIBANOutput`

NewRegisterIBANOutputWithDefaults instantiates a new RegisterIBANOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIbanId

`func (o *RegisterIBANOutput) GetIbanId() int64`

GetIbanId returns the IbanId field if non-nil, zero value otherwise.

### GetIbanIdOk

`func (o *RegisterIBANOutput) GetIbanIdOk() (*int64, bool)`

GetIbanIdOk returns a tuple with the IbanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbanId

`func (o *RegisterIBANOutput) SetIbanId(v int64)`

SetIbanId sets IbanId field to given value.

### HasIbanId

`func (o *RegisterIBANOutput) HasIbanId() bool`

HasIbanId returns a boolean if a field has been set.

### GetStatus

`func (o *RegisterIBANOutput) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegisterIBANOutput) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegisterIBANOutput) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RegisterIBANOutput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetError

`func (o *RegisterIBANOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RegisterIBANOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RegisterIBANOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *RegisterIBANOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


