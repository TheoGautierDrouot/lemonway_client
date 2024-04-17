# Bank

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** | Country code associated with the payment bank. ISO Alpha-2 codes accepted: (FR) France, (ES), Spain (DE), (IT) Italy, Germany, and (PT) Portugal | [optional] 
**ParentName** | Pointer to **string** | The bank&#39;s parent company name. Example: Banque Populaire | [optional] 
**LogoURL** | Pointer to **string** | A weblink to the bank&#39;s logo | [optional] 
**BankName** | Pointer to **string** | This is the bank name. Example: Banque Populaire Alsace Lorraine | [optional] 
**BankId** | Pointer to **string** | This is the unique &#x60;bankId&#x60; associated with the bank. | [optional] 
**Segments** | Pointer to **[]string** | Personal or Business Acoount | [optional] 

## Methods

### NewBank

`func NewBank() *Bank`

NewBank instantiates a new Bank object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankWithDefaults

`func NewBankWithDefaults() *Bank`

NewBankWithDefaults instantiates a new Bank object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *Bank) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Bank) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Bank) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Bank) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetParentName

`func (o *Bank) GetParentName() string`

GetParentName returns the ParentName field if non-nil, zero value otherwise.

### GetParentNameOk

`func (o *Bank) GetParentNameOk() (*string, bool)`

GetParentNameOk returns a tuple with the ParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentName

`func (o *Bank) SetParentName(v string)`

SetParentName sets ParentName field to given value.

### HasParentName

`func (o *Bank) HasParentName() bool`

HasParentName returns a boolean if a field has been set.

### GetLogoURL

`func (o *Bank) GetLogoURL() string`

GetLogoURL returns the LogoURL field if non-nil, zero value otherwise.

### GetLogoURLOk

`func (o *Bank) GetLogoURLOk() (*string, bool)`

GetLogoURLOk returns a tuple with the LogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoURL

`func (o *Bank) SetLogoURL(v string)`

SetLogoURL sets LogoURL field to given value.

### HasLogoURL

`func (o *Bank) HasLogoURL() bool`

HasLogoURL returns a boolean if a field has been set.

### GetBankName

`func (o *Bank) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *Bank) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *Bank) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *Bank) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetBankId

`func (o *Bank) GetBankId() string`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *Bank) GetBankIdOk() (*string, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *Bank) SetBankId(v string)`

SetBankId sets BankId field to given value.

### HasBankId

`func (o *Bank) HasBankId() bool`

HasBankId returns a boolean if a field has been set.

### GetSegments

`func (o *Bank) GetSegments() []string`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *Bank) GetSegmentsOk() (*[]string, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *Bank) SetSegments(v []string)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *Bank) HasSegments() bool`

HasSegments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


