# MoneyIn3DAuthenticateOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation3DCode** | Pointer to **string** | Mercanet return code:  00: Authenticated owner  55: Owner not authenticated  62: Owner By-pass on ACS | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewMoneyIn3DAuthenticateOutput

`func NewMoneyIn3DAuthenticateOutput() *MoneyIn3DAuthenticateOutput`

NewMoneyIn3DAuthenticateOutput instantiates a new MoneyIn3DAuthenticateOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoneyIn3DAuthenticateOutputWithDefaults

`func NewMoneyIn3DAuthenticateOutputWithDefaults() *MoneyIn3DAuthenticateOutput`

NewMoneyIn3DAuthenticateOutputWithDefaults instantiates a new MoneyIn3DAuthenticateOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation3DCode

`func (o *MoneyIn3DAuthenticateOutput) GetOperation3DCode() string`

GetOperation3DCode returns the Operation3DCode field if non-nil, zero value otherwise.

### GetOperation3DCodeOk

`func (o *MoneyIn3DAuthenticateOutput) GetOperation3DCodeOk() (*string, bool)`

GetOperation3DCodeOk returns a tuple with the Operation3DCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation3DCode

`func (o *MoneyIn3DAuthenticateOutput) SetOperation3DCode(v string)`

SetOperation3DCode sets Operation3DCode field to given value.

### HasOperation3DCode

`func (o *MoneyIn3DAuthenticateOutput) HasOperation3DCode() bool`

HasOperation3DCode returns a boolean if a field has been set.

### GetError

`func (o *MoneyIn3DAuthenticateOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MoneyIn3DAuthenticateOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MoneyIn3DAuthenticateOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *MoneyIn3DAuthenticateOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


