# Company

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Company name (at least one alphabethic character is required) | 
**Description** | **string** | Company description | 
**WebsiteUrl** | Pointer to **string** | Website URL | [optional] 
**IdentificationNumber** | Pointer to **string** | Company identification number | [optional] 

## Methods

### NewCompany

`func NewCompany(name string, description string, ) *Company`

NewCompany instantiates a new Company object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyWithDefaults

`func NewCompanyWithDefaults() *Company`

NewCompanyWithDefaults instantiates a new Company object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Company) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Company) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Company) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Company) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Company) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Company) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetWebsiteUrl

`func (o *Company) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *Company) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *Company) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *Company) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetIdentificationNumber

`func (o *Company) GetIdentificationNumber() string`

GetIdentificationNumber returns the IdentificationNumber field if non-nil, zero value otherwise.

### GetIdentificationNumberOk

`func (o *Company) GetIdentificationNumberOk() (*string, bool)`

GetIdentificationNumberOk returns a tuple with the IdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationNumber

`func (o *Company) SetIdentificationNumber(v string)`

SetIdentificationNumber sets IdentificationNumber field to given value.

### HasIdentificationNumber

`func (o *Company) HasIdentificationNumber() bool`

HasIdentificationNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


