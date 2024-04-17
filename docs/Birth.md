# Birth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | Client Date of Birth | 
**City** | Pointer to **string** | Client City of Birth | [optional] 
**Country** | Pointer to **string** | Client Country of Birth, using ISO 3166-1 alpha-3. Three-letter country code, for example: FRA (France), GBR (United Kingdom of Great Britain and Northern Ireland) | [optional] 

## Methods

### NewBirth

`func NewBirth(date string, ) *Birth`

NewBirth instantiates a new Birth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBirthWithDefaults

`func NewBirthWithDefaults() *Birth`

NewBirthWithDefaults instantiates a new Birth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *Birth) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Birth) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Birth) SetDate(v string)`

SetDate sets Date field to given value.


### GetCity

`func (o *Birth) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Birth) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Birth) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Birth) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *Birth) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Birth) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Birth) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Birth) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


