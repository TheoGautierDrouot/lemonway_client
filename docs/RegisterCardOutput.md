# RegisterCardOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Card** | Pointer to [**Card**](Card.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewRegisterCardOutput

`func NewRegisterCardOutput() *RegisterCardOutput`

NewRegisterCardOutput instantiates a new RegisterCardOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterCardOutputWithDefaults

`func NewRegisterCardOutputWithDefaults() *RegisterCardOutput`

NewRegisterCardOutputWithDefaults instantiates a new RegisterCardOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCard

`func (o *RegisterCardOutput) GetCard() Card`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *RegisterCardOutput) GetCardOk() (*Card, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *RegisterCardOutput) SetCard(v Card)`

SetCard sets Card field to given value.

### HasCard

`func (o *RegisterCardOutput) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetError

`func (o *RegisterCardOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RegisterCardOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RegisterCardOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *RegisterCardOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


