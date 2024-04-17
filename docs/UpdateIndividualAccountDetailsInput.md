# UpdateIndividualAccountDetailsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewUpdateIndividualAccountDetailsInput

`func NewUpdateIndividualAccountDetailsInput() *UpdateIndividualAccountDetailsInput`

NewUpdateIndividualAccountDetailsInput instantiates a new UpdateIndividualAccountDetailsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIndividualAccountDetailsInputWithDefaults

`func NewUpdateIndividualAccountDetailsInputWithDefaults() *UpdateIndividualAccountDetailsInput`

NewUpdateIndividualAccountDetailsInputWithDefaults instantiates a new UpdateIndividualAccountDetailsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UpdateIndividualAccountDetailsInput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateIndividualAccountDetailsInput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateIndividualAccountDetailsInput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateIndividualAccountDetailsInput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetTitle

`func (o *UpdateIndividualAccountDetailsInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateIndividualAccountDetailsInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateIndividualAccountDetailsInput) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateIndividualAccountDetailsInput) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFirstName

`func (o *UpdateIndividualAccountDetailsInput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UpdateIndividualAccountDetailsInput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UpdateIndividualAccountDetailsInput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UpdateIndividualAccountDetailsInput) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UpdateIndividualAccountDetailsInput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UpdateIndividualAccountDetailsInput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UpdateIndividualAccountDetailsInput) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UpdateIndividualAccountDetailsInput) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetAdresse

`func (o *UpdateIndividualAccountDetailsInput) GetAdresse() Address`

GetAdresse returns the Adresse field if non-nil, zero value otherwise.

### GetAdresseOk

`func (o *UpdateIndividualAccountDetailsInput) GetAdresseOk() (*Address, bool)`

GetAdresseOk returns a tuple with the Adresse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresse

`func (o *UpdateIndividualAccountDetailsInput) SetAdresse(v Address)`

SetAdresse sets Adresse field to given value.

### HasAdresse

`func (o *UpdateIndividualAccountDetailsInput) HasAdresse() bool`

HasAdresse returns a boolean if a field has been set.

### GetBirth

`func (o *UpdateIndividualAccountDetailsInput) GetBirth() Birth`

GetBirth returns the Birth field if non-nil, zero value otherwise.

### GetBirthOk

`func (o *UpdateIndividualAccountDetailsInput) GetBirthOk() (*Birth, bool)`

GetBirthOk returns a tuple with the Birth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirth

`func (o *UpdateIndividualAccountDetailsInput) SetBirth(v Birth)`

SetBirth sets Birth field to given value.

### HasBirth

`func (o *UpdateIndividualAccountDetailsInput) HasBirth() bool`

HasBirth returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *UpdateIndividualAccountDetailsInput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UpdateIndividualAccountDetailsInput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UpdateIndividualAccountDetailsInput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *UpdateIndividualAccountDetailsInput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetMobileNumber

`func (o *UpdateIndividualAccountDetailsInput) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *UpdateIndividualAccountDetailsInput) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *UpdateIndividualAccountDetailsInput) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.

### HasMobileNumber

`func (o *UpdateIndividualAccountDetailsInput) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.

### GetIsDebtor

`func (o *UpdateIndividualAccountDetailsInput) GetIsDebtor() bool`

GetIsDebtor returns the IsDebtor field if non-nil, zero value otherwise.

### GetIsDebtorOk

`func (o *UpdateIndividualAccountDetailsInput) GetIsDebtorOk() (*bool, bool)`

GetIsDebtorOk returns a tuple with the IsDebtor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDebtor

`func (o *UpdateIndividualAccountDetailsInput) SetIsDebtor(v bool)`

SetIsDebtor sets IsDebtor field to given value.

### HasIsDebtor

`func (o *UpdateIndividualAccountDetailsInput) HasIsDebtor() bool`

HasIsDebtor returns a boolean if a field has been set.

### GetNationality

`func (o *UpdateIndividualAccountDetailsInput) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *UpdateIndividualAccountDetailsInput) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *UpdateIndividualAccountDetailsInput) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *UpdateIndividualAccountDetailsInput) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetPayerOrBeneficiary

`func (o *UpdateIndividualAccountDetailsInput) GetPayerOrBeneficiary() int32`

GetPayerOrBeneficiary returns the PayerOrBeneficiary field if non-nil, zero value otherwise.

### GetPayerOrBeneficiaryOk

`func (o *UpdateIndividualAccountDetailsInput) GetPayerOrBeneficiaryOk() (*int32, bool)`

GetPayerOrBeneficiaryOk returns a tuple with the PayerOrBeneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerOrBeneficiary

`func (o *UpdateIndividualAccountDetailsInput) SetPayerOrBeneficiary(v int32)`

SetPayerOrBeneficiary sets PayerOrBeneficiary field to given value.

### HasPayerOrBeneficiary

`func (o *UpdateIndividualAccountDetailsInput) HasPayerOrBeneficiary() bool`

HasPayerOrBeneficiary returns a boolean if a field has been set.

### GetIsCompany

`func (o *UpdateIndividualAccountDetailsInput) GetIsCompany() bool`

GetIsCompany returns the IsCompany field if non-nil, zero value otherwise.

### GetIsCompanyOk

`func (o *UpdateIndividualAccountDetailsInput) GetIsCompanyOk() (*bool, bool)`

GetIsCompanyOk returns a tuple with the IsCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompany

`func (o *UpdateIndividualAccountDetailsInput) SetIsCompany(v bool)`

SetIsCompany sets IsCompany field to given value.

### HasIsCompany

`func (o *UpdateIndividualAccountDetailsInput) HasIsCompany() bool`

HasIsCompany returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


