# UltimateBeneficialOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Firstname** | **string** | Client First Name | 
**Lastname** | **string** | Client Last Name | 
**Nationality** | **string** | Client Nationality (use ISO 3166-1 alpha-3 format) | 
**DateOfBirth** | **string** | Client Birthdate | 
**CityOfBirth** | Pointer to **string** | Client City of Birth | [optional] 
**CountryOfBirth** | Pointer to **string** | Client Country of Birth (use ISO 3166-1 alpha-3 format) | [optional] 
**CountryOfResidence** | **string** | Country of Residence (use ISO 3166-1 alpha-3 format) | 
**StartDate** | Pointer to **int32** |  | [optional] 
**EndDate** | Pointer to **int32** |  | [optional] 
**UltimateBeneficialOwnerId** | Pointer to **int64** | Ultimate Beneficial Owner ID | [optional] 

## Methods

### NewUltimateBeneficialOwner

`func NewUltimateBeneficialOwner(firstname string, lastname string, nationality string, dateOfBirth string, countryOfResidence string, ) *UltimateBeneficialOwner`

NewUltimateBeneficialOwner instantiates a new UltimateBeneficialOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUltimateBeneficialOwnerWithDefaults

`func NewUltimateBeneficialOwnerWithDefaults() *UltimateBeneficialOwner`

NewUltimateBeneficialOwnerWithDefaults instantiates a new UltimateBeneficialOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstname

`func (o *UltimateBeneficialOwner) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *UltimateBeneficialOwner) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *UltimateBeneficialOwner) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.


### GetLastname

`func (o *UltimateBeneficialOwner) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *UltimateBeneficialOwner) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *UltimateBeneficialOwner) SetLastname(v string)`

SetLastname sets Lastname field to given value.


### GetNationality

`func (o *UltimateBeneficialOwner) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *UltimateBeneficialOwner) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *UltimateBeneficialOwner) SetNationality(v string)`

SetNationality sets Nationality field to given value.


### GetDateOfBirth

`func (o *UltimateBeneficialOwner) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *UltimateBeneficialOwner) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *UltimateBeneficialOwner) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.


### GetCityOfBirth

`func (o *UltimateBeneficialOwner) GetCityOfBirth() string`

GetCityOfBirth returns the CityOfBirth field if non-nil, zero value otherwise.

### GetCityOfBirthOk

`func (o *UltimateBeneficialOwner) GetCityOfBirthOk() (*string, bool)`

GetCityOfBirthOk returns a tuple with the CityOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityOfBirth

`func (o *UltimateBeneficialOwner) SetCityOfBirth(v string)`

SetCityOfBirth sets CityOfBirth field to given value.

### HasCityOfBirth

`func (o *UltimateBeneficialOwner) HasCityOfBirth() bool`

HasCityOfBirth returns a boolean if a field has been set.

### GetCountryOfBirth

`func (o *UltimateBeneficialOwner) GetCountryOfBirth() string`

GetCountryOfBirth returns the CountryOfBirth field if non-nil, zero value otherwise.

### GetCountryOfBirthOk

`func (o *UltimateBeneficialOwner) GetCountryOfBirthOk() (*string, bool)`

GetCountryOfBirthOk returns a tuple with the CountryOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfBirth

`func (o *UltimateBeneficialOwner) SetCountryOfBirth(v string)`

SetCountryOfBirth sets CountryOfBirth field to given value.

### HasCountryOfBirth

`func (o *UltimateBeneficialOwner) HasCountryOfBirth() bool`

HasCountryOfBirth returns a boolean if a field has been set.

### GetCountryOfResidence

`func (o *UltimateBeneficialOwner) GetCountryOfResidence() string`

GetCountryOfResidence returns the CountryOfResidence field if non-nil, zero value otherwise.

### GetCountryOfResidenceOk

`func (o *UltimateBeneficialOwner) GetCountryOfResidenceOk() (*string, bool)`

GetCountryOfResidenceOk returns a tuple with the CountryOfResidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfResidence

`func (o *UltimateBeneficialOwner) SetCountryOfResidence(v string)`

SetCountryOfResidence sets CountryOfResidence field to given value.


### GetStartDate

`func (o *UltimateBeneficialOwner) GetStartDate() int32`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UltimateBeneficialOwner) GetStartDateOk() (*int32, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UltimateBeneficialOwner) SetStartDate(v int32)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UltimateBeneficialOwner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *UltimateBeneficialOwner) GetEndDate() int32`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UltimateBeneficialOwner) GetEndDateOk() (*int32, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UltimateBeneficialOwner) SetEndDate(v int32)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UltimateBeneficialOwner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetUltimateBeneficialOwnerId

`func (o *UltimateBeneficialOwner) GetUltimateBeneficialOwnerId() int64`

GetUltimateBeneficialOwnerId returns the UltimateBeneficialOwnerId field if non-nil, zero value otherwise.

### GetUltimateBeneficialOwnerIdOk

`func (o *UltimateBeneficialOwner) GetUltimateBeneficialOwnerIdOk() (*int64, bool)`

GetUltimateBeneficialOwnerIdOk returns a tuple with the UltimateBeneficialOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUltimateBeneficialOwnerId

`func (o *UltimateBeneficialOwner) SetUltimateBeneficialOwnerId(v int64)`

SetUltimateBeneficialOwnerId sets UltimateBeneficialOwnerId field to given value.

### HasUltimateBeneficialOwnerId

`func (o *UltimateBeneficialOwner) HasUltimateBeneficialOwnerId() bool`

HasUltimateBeneficialOwnerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


