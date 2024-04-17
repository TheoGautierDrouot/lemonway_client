# MoneyInWebInitOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebKitToken** | Pointer to **string** | Payment Token to Pass to WebKit URL using GET | [optional] 
**Id** | Pointer to **int32** | Transaction ID | [optional] 
**Id** | Pointer to **int32** | OBSOLETE Transaction ID | [optional] 
**CardId** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewMoneyInWebInitOutput

`func NewMoneyInWebInitOutput() *MoneyInWebInitOutput`

NewMoneyInWebInitOutput instantiates a new MoneyInWebInitOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoneyInWebInitOutputWithDefaults

`func NewMoneyInWebInitOutputWithDefaults() *MoneyInWebInitOutput`

NewMoneyInWebInitOutputWithDefaults instantiates a new MoneyInWebInitOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebKitToken

`func (o *MoneyInWebInitOutput) GetWebKitToken() string`

GetWebKitToken returns the WebKitToken field if non-nil, zero value otherwise.

### GetWebKitTokenOk

`func (o *MoneyInWebInitOutput) GetWebKitTokenOk() (*string, bool)`

GetWebKitTokenOk returns a tuple with the WebKitToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebKitToken

`func (o *MoneyInWebInitOutput) SetWebKitToken(v string)`

SetWebKitToken sets WebKitToken field to given value.

### HasWebKitToken

`func (o *MoneyInWebInitOutput) HasWebKitToken() bool`

HasWebKitToken returns a boolean if a field has been set.

### GetId

`func (o *MoneyInWebInitOutput) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MoneyInWebInitOutput) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MoneyInWebInitOutput) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MoneyInWebInitOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetId

`func (o *MoneyInWebInitOutput) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MoneyInWebInitOutput) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MoneyInWebInitOutput) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MoneyInWebInitOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCardId

`func (o *MoneyInWebInitOutput) GetCardId() int32`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *MoneyInWebInitOutput) GetCardIdOk() (*int32, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *MoneyInWebInitOutput) SetCardId(v int32)`

SetCardId sets CardId field to given value.

### HasCardId

`func (o *MoneyInWebInitOutput) HasCardId() bool`

HasCardId returns a boolean if a field has been set.

### GetError

`func (o *MoneyInWebInitOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MoneyInWebInitOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MoneyInWebInitOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *MoneyInWebInitOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


