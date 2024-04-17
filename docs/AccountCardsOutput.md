# AccountCardsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cards** | Pointer to [**[]Card**](Card.md) | Dispalays a list of documents that have changed since original entry date. | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewAccountCardsOutput

`func NewAccountCardsOutput() *AccountCardsOutput`

NewAccountCardsOutput instantiates a new AccountCardsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCardsOutputWithDefaults

`func NewAccountCardsOutputWithDefaults() *AccountCardsOutput`

NewAccountCardsOutputWithDefaults instantiates a new AccountCardsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCards

`func (o *AccountCardsOutput) GetCards() []Card`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *AccountCardsOutput) GetCardsOk() (*[]Card, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *AccountCardsOutput) SetCards(v []Card)`

SetCards sets Cards field to given value.

### HasCards

`func (o *AccountCardsOutput) HasCards() bool`

HasCards returns a boolean if a field has been set.

### GetError

`func (o *AccountCardsOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AccountCardsOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AccountCardsOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *AccountCardsOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


