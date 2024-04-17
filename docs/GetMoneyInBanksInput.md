# GetMoneyInBanksInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCodes** | Pointer to **string** | Comma separated string of ISO Alpha-2 country codes.  Available country codes include:    - France (FR)     - Spain (ES)     - Italy (IT)     - Germany (DE)     - Portugal (PT) | [optional] 

## Methods

### NewGetMoneyInBanksInput

`func NewGetMoneyInBanksInput() *GetMoneyInBanksInput`

NewGetMoneyInBanksInput instantiates a new GetMoneyInBanksInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMoneyInBanksInputWithDefaults

`func NewGetMoneyInBanksInputWithDefaults() *GetMoneyInBanksInput`

NewGetMoneyInBanksInputWithDefaults instantiates a new GetMoneyInBanksInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCodes

`func (o *GetMoneyInBanksInput) GetCountryCodes() string`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *GetMoneyInBanksInput) GetCountryCodesOk() (*string, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *GetMoneyInBanksInput) SetCountryCodes(v string)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *GetMoneyInBanksInput) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


