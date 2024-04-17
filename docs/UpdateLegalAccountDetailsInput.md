# UpdateLegalAccountDetailsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | [**Company**](Company.md) |  | 
**Email** | Pointer to **string** | New Email. | [optional] 
**Title** | Pointer to **string** | Client title | [optional] 
**FirstName** | Pointer to **string** | Client first name | [optional] 
**LastName** | Pointer to **string** | Client last name | [optional] 
**Adresse** | Pointer to [**Address**](Address.md) |  | [optional] 
**Birth** | Pointer to [**Birth**](Birth.md) |  | [optional] 
**PhoneNumber** | Pointer to **string** | MSISDN phone number | [optional] 
**MobileNumber** | Pointer to **string** | Mobile phone number with MSISDN format: international number with country code without \&quot;00\&quot; neither \&quot;+\&quot;.&lt;br /&gt;   This will be used by default when eletronically signing documents. | [optional] 
**IsDebtor** | Pointer to **bool** | For crowdfunding/loan platforms, indicate if the wallet is created for a debtor. | [optional] 
**Nationality** | Pointer to **string** | client nationality, using ISO 3166-1 alpha-3 format.  Separate multiple nationalities with a comma. | [optional] 
**PayerOrBeneficiary** | Pointer to **int32** | Indicates if the payment account is created for a payer or a beneficiary:  Empty: unknown status (default)  1: Payer  2: Beneficiary | [optional] 
**IsCompany** | Pointer to **bool** | **Important:** This object is deprecated and no longer permissible in accordance with EU compliance rules. | [optional] 

## Methods

### NewUpdateLegalAccountDetailsInput

`func NewUpdateLegalAccountDetailsInput(company Company, ) *UpdateLegalAccountDetailsInput`

NewUpdateLegalAccountDetailsInput instantiates a new UpdateLegalAccountDetailsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLegalAccountDetailsInputWithDefaults

`func NewUpdateLegalAccountDetailsInputWithDefaults() *UpdateLegalAccountDetailsInput`

NewUpdateLegalAccountDetailsInputWithDefaults instantiates a new UpdateLegalAccountDetailsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *UpdateLegalAccountDetailsInput) GetCompany() Company`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UpdateLegalAccountDetailsInput) GetCompanyOk() (*Company, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UpdateLegalAccountDetailsInput) SetCompany(v Company)`

SetCompany sets Company field to given value.


### GetEmail

`func (o *UpdateLegalAccountDetailsInput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateLegalAccountDetailsInput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateLegalAccountDetailsInput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateLegalAccountDetailsInput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetTitle

`func (o *UpdateLegalAccountDetailsInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateLegalAccountDetailsInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateLegalAccountDetailsInput) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateLegalAccountDetailsInput) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFirstName

`func (o *UpdateLegalAccountDetailsInput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UpdateLegalAccountDetailsInput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UpdateLegalAccountDetailsInput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UpdateLegalAccountDetailsInput) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UpdateLegalAccountDetailsInput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UpdateLegalAccountDetailsInput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UpdateLegalAccountDetailsInput) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UpdateLegalAccountDetailsInput) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetAdresse

`func (o *UpdateLegalAccountDetailsInput) GetAdresse() Address`

GetAdresse returns the Adresse field if non-nil, zero value otherwise.

### GetAdresseOk

`func (o *UpdateLegalAccountDetailsInput) GetAdresseOk() (*Address, bool)`

GetAdresseOk returns a tuple with the Adresse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresse

`func (o *UpdateLegalAccountDetailsInput) SetAdresse(v Address)`

SetAdresse sets Adresse field to given value.

### HasAdresse

`func (o *UpdateLegalAccountDetailsInput) HasAdresse() bool`

HasAdresse returns a boolean if a field has been set.

### GetBirth

`func (o *UpdateLegalAccountDetailsInput) GetBirth() Birth`

GetBirth returns the Birth field if non-nil, zero value otherwise.

### GetBirthOk

`func (o *UpdateLegalAccountDetailsInput) GetBirthOk() (*Birth, bool)`

GetBirthOk returns a tuple with the Birth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirth

`func (o *UpdateLegalAccountDetailsInput) SetBirth(v Birth)`

SetBirth sets Birth field to given value.

### HasBirth

`func (o *UpdateLegalAccountDetailsInput) HasBirth() bool`

HasBirth returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *UpdateLegalAccountDetailsInput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UpdateLegalAccountDetailsInput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UpdateLegalAccountDetailsInput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *UpdateLegalAccountDetailsInput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetMobileNumber

`func (o *UpdateLegalAccountDetailsInput) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *UpdateLegalAccountDetailsInput) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *UpdateLegalAccountDetailsInput) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.

### HasMobileNumber

`func (o *UpdateLegalAccountDetailsInput) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.

### GetIsDebtor

`func (o *UpdateLegalAccountDetailsInput) GetIsDebtor() bool`

GetIsDebtor returns the IsDebtor field if non-nil, zero value otherwise.

### GetIsDebtorOk

`func (o *UpdateLegalAccountDetailsInput) GetIsDebtorOk() (*bool, bool)`

GetIsDebtorOk returns a tuple with the IsDebtor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDebtor

`func (o *UpdateLegalAccountDetailsInput) SetIsDebtor(v bool)`

SetIsDebtor sets IsDebtor field to given value.

### HasIsDebtor

`func (o *UpdateLegalAccountDetailsInput) HasIsDebtor() bool`

HasIsDebtor returns a boolean if a field has been set.

### GetNationality

`func (o *UpdateLegalAccountDetailsInput) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *UpdateLegalAccountDetailsInput) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *UpdateLegalAccountDetailsInput) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *UpdateLegalAccountDetailsInput) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetPayerOrBeneficiary

`func (o *UpdateLegalAccountDetailsInput) GetPayerOrBeneficiary() int32`

GetPayerOrBeneficiary returns the PayerOrBeneficiary field if non-nil, zero value otherwise.

### GetPayerOrBeneficiaryOk

`func (o *UpdateLegalAccountDetailsInput) GetPayerOrBeneficiaryOk() (*int32, bool)`

GetPayerOrBeneficiaryOk returns a tuple with the PayerOrBeneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerOrBeneficiary

`func (o *UpdateLegalAccountDetailsInput) SetPayerOrBeneficiary(v int32)`

SetPayerOrBeneficiary sets PayerOrBeneficiary field to given value.

### HasPayerOrBeneficiary

`func (o *UpdateLegalAccountDetailsInput) HasPayerOrBeneficiary() bool`

HasPayerOrBeneficiary returns a boolean if a field has been set.

### GetIsCompany

`func (o *UpdateLegalAccountDetailsInput) GetIsCompany() bool`

GetIsCompany returns the IsCompany field if non-nil, zero value otherwise.

### GetIsCompanyOk

`func (o *UpdateLegalAccountDetailsInput) GetIsCompanyOk() (*bool, bool)`

GetIsCompanyOk returns a tuple with the IsCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompany

`func (o *UpdateLegalAccountDetailsInput) SetIsCompany(v bool)`

SetIsCompany sets IsCompany field to given value.

### HasIsCompany

`func (o *UpdateLegalAccountDetailsInput) HasIsCompany() bool`

HasIsCompany returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


