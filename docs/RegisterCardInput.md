# RegisterCardInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Payment Account ID | 
**Card** | [**CardInfo**](CardInfo.md) |  | 
**SpecialConfiguration** | Pointer to **string** | Leave Empty | [optional] 

## Methods

### NewRegisterCardInput

`func NewRegisterCardInput(accountId string, card CardInfo, ) *RegisterCardInput`

NewRegisterCardInput instantiates a new RegisterCardInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterCardInputWithDefaults

`func NewRegisterCardInputWithDefaults() *RegisterCardInput`

NewRegisterCardInputWithDefaults instantiates a new RegisterCardInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *RegisterCardInput) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RegisterCardInput) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RegisterCardInput) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCard

`func (o *RegisterCardInput) GetCard() CardInfo`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *RegisterCardInput) GetCardOk() (*CardInfo, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *RegisterCardInput) SetCard(v CardInfo)`

SetCard sets Card field to given value.


### GetSpecialConfiguration

`func (o *RegisterCardInput) GetSpecialConfiguration() string`

GetSpecialConfiguration returns the SpecialConfiguration field if non-nil, zero value otherwise.

### GetSpecialConfigurationOk

`func (o *RegisterCardInput) GetSpecialConfigurationOk() (*string, bool)`

GetSpecialConfigurationOk returns a tuple with the SpecialConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialConfiguration

`func (o *RegisterCardInput) SetSpecialConfiguration(v string)`

SetSpecialConfiguration sets SpecialConfiguration field to given value.

### HasSpecialConfiguration

`func (o *RegisterCardInput) HasSpecialConfiguration() bool`

HasSpecialConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


