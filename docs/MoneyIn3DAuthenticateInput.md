# MoneyIn3DAuthenticateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Card** | Pointer to [**CardInfo**](CardInfo.md) |  | [optional] 
**Md** | Pointer to **string** | MD Data Returned by 3-D Secure Site | [optional] 
**Pares** | Pointer to **string** | Pares Data Returned by 3-D Secure Authentication Website | [optional] 
**SpecialConfig** | Pointer to **string** | Leave Empty | [optional] 

## Methods

### NewMoneyIn3DAuthenticateInput

`func NewMoneyIn3DAuthenticateInput() *MoneyIn3DAuthenticateInput`

NewMoneyIn3DAuthenticateInput instantiates a new MoneyIn3DAuthenticateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoneyIn3DAuthenticateInputWithDefaults

`func NewMoneyIn3DAuthenticateInputWithDefaults() *MoneyIn3DAuthenticateInput`

NewMoneyIn3DAuthenticateInputWithDefaults instantiates a new MoneyIn3DAuthenticateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCard

`func (o *MoneyIn3DAuthenticateInput) GetCard() CardInfo`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *MoneyIn3DAuthenticateInput) GetCardOk() (*CardInfo, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *MoneyIn3DAuthenticateInput) SetCard(v CardInfo)`

SetCard sets Card field to given value.

### HasCard

`func (o *MoneyIn3DAuthenticateInput) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetMd

`func (o *MoneyIn3DAuthenticateInput) GetMd() string`

GetMd returns the Md field if non-nil, zero value otherwise.

### GetMdOk

`func (o *MoneyIn3DAuthenticateInput) GetMdOk() (*string, bool)`

GetMdOk returns a tuple with the Md field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd

`func (o *MoneyIn3DAuthenticateInput) SetMd(v string)`

SetMd sets Md field to given value.

### HasMd

`func (o *MoneyIn3DAuthenticateInput) HasMd() bool`

HasMd returns a boolean if a field has been set.

### GetPares

`func (o *MoneyIn3DAuthenticateInput) GetPares() string`

GetPares returns the Pares field if non-nil, zero value otherwise.

### GetParesOk

`func (o *MoneyIn3DAuthenticateInput) GetParesOk() (*string, bool)`

GetParesOk returns a tuple with the Pares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPares

`func (o *MoneyIn3DAuthenticateInput) SetPares(v string)`

SetPares sets Pares field to given value.

### HasPares

`func (o *MoneyIn3DAuthenticateInput) HasPares() bool`

HasPares returns a boolean if a field has been set.

### GetSpecialConfig

`func (o *MoneyIn3DAuthenticateInput) GetSpecialConfig() string`

GetSpecialConfig returns the SpecialConfig field if non-nil, zero value otherwise.

### GetSpecialConfigOk

`func (o *MoneyIn3DAuthenticateInput) GetSpecialConfigOk() (*string, bool)`

GetSpecialConfigOk returns a tuple with the SpecialConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialConfig

`func (o *MoneyIn3DAuthenticateInput) SetSpecialConfig(v string)`

SetSpecialConfig sets SpecialConfig field to given value.

### HasSpecialConfig

`func (o *MoneyIn3DAuthenticateInput) HasSpecialConfig() bool`

HasSpecialConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


