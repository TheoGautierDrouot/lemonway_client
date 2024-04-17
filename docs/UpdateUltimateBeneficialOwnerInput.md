# UpdateUltimateBeneficialOwnerInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | Client First Name | 
**LastName** | **string** | Client Last Name | 
**Nationality** | Pointer to **string** | Client Nationality (use ISO 3166-1 alpha-3 format) | [optional] 
**DateOfBirth** | **string** | Client Birthdate | 
**CityOfBirth** | **string** | Client City of Birth | 
**CountryOfBirth** | **string** | Client Country of Birth (use ISO 3166-1 alpha-3 format) | 
**CountryOfResidence** | Pointer to **string** | Country of Residence (use ISO 3166-1 alpha-3 format) | [optional] 
**IsActive** | Pointer to **bool** | Indicate if the UBO is active or not | [optional] 

## Methods

### NewUpdateUltimateBeneficialOwnerInput

`func NewUpdateUltimateBeneficialOwnerInput(firstName string, lastName string, dateOfBirth string, cityOfBirth string, countryOfBirth string, ) *UpdateUltimateBeneficialOwnerInput`

NewUpdateUltimateBeneficialOwnerInput instantiates a new UpdateUltimateBeneficialOwnerInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUltimateBeneficialOwnerInputWithDefaults

`func NewUpdateUltimateBeneficialOwnerInputWithDefaults() *UpdateUltimateBeneficialOwnerInput`

NewUpdateUltimateBeneficialOwnerInputWithDefaults instantiates a new UpdateUltimateBeneficialOwnerInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *UpdateUltimateBeneficialOwnerInput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UpdateUltimateBeneficialOwnerInput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UpdateUltimateBeneficialOwnerInput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *UpdateUltimateBeneficialOwnerInput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UpdateUltimateBeneficialOwnerInput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UpdateUltimateBeneficialOwnerInput) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetNationality

`func (o *UpdateUltimateBeneficialOwnerInput) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *UpdateUltimateBeneficialOwnerInput) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *UpdateUltimateBeneficialOwnerInput) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *UpdateUltimateBeneficialOwnerInput) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *UpdateUltimateBeneficialOwnerInput) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *UpdateUltimateBeneficialOwnerInput) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *UpdateUltimateBeneficialOwnerInput) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.


### GetCityOfBirth

`func (o *UpdateUltimateBeneficialOwnerInput) GetCityOfBirth() string`

GetCityOfBirth returns the CityOfBirth field if non-nil, zero value otherwise.

### GetCityOfBirthOk

`func (o *UpdateUltimateBeneficialOwnerInput) GetCityOfBirthOk() (*string, bool)`

GetCityOfBirthOk returns a tuple with the CityOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityOfBirth

`func (o *UpdateUltimateBeneficialOwnerInput) SetCityOfBirth(v string)`

SetCityOfBirth sets CityOfBirth field to given value.


### GetCountryOfBirth

`func (o *UpdateUltimateBeneficialOwnerInput) GetCountryOfBirth() string`

GetCountryOfBirth returns the CountryOfBirth field if non-nil, zero value otherwise.

### GetCountryOfBirthOk

`func (o *UpdateUltimateBeneficialOwnerInput) GetCountryOfBirthOk() (*string, bool)`

GetCountryOfBirthOk returns a tuple with the CountryOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfBirth

`func (o *UpdateUltimateBeneficialOwnerInput) SetCountryOfBirth(v string)`

SetCountryOfBirth sets CountryOfBirth field to given value.


### GetCountryOfResidence

`func (o *UpdateUltimateBeneficialOwnerInput) GetCountryOfResidence() string`

GetCountryOfResidence returns the CountryOfResidence field if non-nil, zero value otherwise.

### GetCountryOfResidenceOk

`func (o *UpdateUltimateBeneficialOwnerInput) GetCountryOfResidenceOk() (*string, bool)`

GetCountryOfResidenceOk returns a tuple with the CountryOfResidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfResidence

`func (o *UpdateUltimateBeneficialOwnerInput) SetCountryOfResidence(v string)`

SetCountryOfResidence sets CountryOfResidence field to given value.

### HasCountryOfResidence

`func (o *UpdateUltimateBeneficialOwnerInput) HasCountryOfResidence() bool`

HasCountryOfResidence returns a boolean if a field has been set.

### GetIsActive

`func (o *UpdateUltimateBeneficialOwnerInput) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateUltimateBeneficialOwnerInput) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateUltimateBeneficialOwnerInput) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateUltimateBeneficialOwnerInput) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


