# GetCardOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Card** | Pointer to [**Card**](Card.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewGetCardOutput

`func NewGetCardOutput() *GetCardOutput`

NewGetCardOutput instantiates a new GetCardOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCardOutputWithDefaults

`func NewGetCardOutputWithDefaults() *GetCardOutput`

NewGetCardOutputWithDefaults instantiates a new GetCardOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCard

`func (o *GetCardOutput) GetCard() Card`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *GetCardOutput) GetCardOk() (*Card, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *GetCardOutput) SetCard(v Card)`

SetCard sets Card field to given value.

### HasCard

`func (o *GetCardOutput) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetError

`func (o *GetCardOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetCardOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetCardOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *GetCardOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


