# RegisterLegalAccountInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | [**Company**](Company.md) |  | 
**IsUltimateBeneficialOwner** | Pointer to **bool** | Indicates if the legal representative is also an Ultimate Beneficial owner (i.e. shareholder with &amp;gt;25% of capital or voting rights). | [optional] 
**AccountId** | **string** | Payment account ID that you use to identify the customer.Choose your unique number.&lt;br /&gt;&lt;b&gt;Note:&lt;/b&gt; If you plan to credit payments accounts  by fund transfer, please use short alphanumeric payment account identifiers (max 20 char.).   Your customers will have to write their payment account identifier in the transfer label/comment, a label of more that 20 characters could be cut when passing the the banking system. | 
**Email** | **string** | Unique Email | 
**Title** | Pointer to **string** | Client title | [optional] 
**FirstName** | **string** | Client first name | 
**LastName** | **string** | Client last name | 
**Adresse** | Pointer to [**Address**](Address.md) |  | [optional] 
**Birth** | Pointer to [**Birth**](Birth.md) |  | [optional] 
**Nationality** | **string** | Client nationality, using ISO 3166-1 alpha-3 format  Please separate multiple nationalities with a comma | 
**PhoneNumber** | Pointer to **string** | Phone number with MSISDN format: international number with country code without \&quot;00\&quot; neither \&quot;+\&quot;. | [optional] 
**MobileNumber** | Pointer to **string** | Mobile phone number with MSISDN format: international number with country code without \&quot;00\&quot; neither \&quot;+\&quot;.   This will be used by default when eletronically signing documents. | [optional] 
**IsDebtor** | Pointer to **bool** | For crowdfunding/loan platforms, indicates if the wallet is created for a debtor. | [optional] 
**PayerOrBeneficiary** | **int32** | Indicates if the payment account is created for a Payer or a Beneficiary.  1: Payer (default)  2: Beneficiary | 
**IsOneTimeCustomerAccount** | Pointer to **bool** | Indicates if the payment account is for a one-time customer.   If yes, the payment account will be created with status 14, allowing only one payment.   The maximum amount will be defined with Lemonway. | [optional] 
**IsTechnicalAccount** | Pointer to **bool** | **Note:** This option is available depending on your contract  True, in case this option is enabled in your contract.  Otherwise it will be considered a Client Wallet. | [optional] 

## Methods

### NewRegisterLegalAccountInput

`func NewRegisterLegalAccountInput(company Company, accountId string, email string, firstName string, lastName string, nationality string, payerOrBeneficiary int32, ) *RegisterLegalAccountInput`

NewRegisterLegalAccountInput instantiates a new RegisterLegalAccountInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterLegalAccountInputWithDefaults

`func NewRegisterLegalAccountInputWithDefaults() *RegisterLegalAccountInput`

NewRegisterLegalAccountInputWithDefaults instantiates a new RegisterLegalAccountInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *RegisterLegalAccountInput) GetCompany() Company`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *RegisterLegalAccountInput) GetCompanyOk() (*Company, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *RegisterLegalAccountInput) SetCompany(v Company)`

SetCompany sets Company field to given value.


### GetIsUltimateBeneficialOwner

`func (o *RegisterLegalAccountInput) GetIsUltimateBeneficialOwner() bool`

GetIsUltimateBeneficialOwner returns the IsUltimateBeneficialOwner field if non-nil, zero value otherwise.

### GetIsUltimateBeneficialOwnerOk

`func (o *RegisterLegalAccountInput) GetIsUltimateBeneficialOwnerOk() (*bool, bool)`

GetIsUltimateBeneficialOwnerOk returns a tuple with the IsUltimateBeneficialOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUltimateBeneficialOwner

`func (o *RegisterLegalAccountInput) SetIsUltimateBeneficialOwner(v bool)`

SetIsUltimateBeneficialOwner sets IsUltimateBeneficialOwner field to given value.

### HasIsUltimateBeneficialOwner

`func (o *RegisterLegalAccountInput) HasIsUltimateBeneficialOwner() bool`

HasIsUltimateBeneficialOwner returns a boolean if a field has been set.

### GetAccountId

`func (o *RegisterLegalAccountInput) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RegisterLegalAccountInput) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RegisterLegalAccountInput) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetEmail

`func (o *RegisterLegalAccountInput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RegisterLegalAccountInput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RegisterLegalAccountInput) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetTitle

`func (o *RegisterLegalAccountInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RegisterLegalAccountInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RegisterLegalAccountInput) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RegisterLegalAccountInput) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFirstName

`func (o *RegisterLegalAccountInput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *RegisterLegalAccountInput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *RegisterLegalAccountInput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *RegisterLegalAccountInput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *RegisterLegalAccountInput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *RegisterLegalAccountInput) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetAdresse

`func (o *RegisterLegalAccountInput) GetAdresse() Address`

GetAdresse returns the Adresse field if non-nil, zero value otherwise.

### GetAdresseOk

`func (o *RegisterLegalAccountInput) GetAdresseOk() (*Address, bool)`

GetAdresseOk returns a tuple with the Adresse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresse

`func (o *RegisterLegalAccountInput) SetAdresse(v Address)`

SetAdresse sets Adresse field to given value.

### HasAdresse

`func (o *RegisterLegalAccountInput) HasAdresse() bool`

HasAdresse returns a boolean if a field has been set.

### GetBirth

`func (o *RegisterLegalAccountInput) GetBirth() Birth`

GetBirth returns the Birth field if non-nil, zero value otherwise.

### GetBirthOk

`func (o *RegisterLegalAccountInput) GetBirthOk() (*Birth, bool)`

GetBirthOk returns a tuple with the Birth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirth

`func (o *RegisterLegalAccountInput) SetBirth(v Birth)`

SetBirth sets Birth field to given value.

### HasBirth

`func (o *RegisterLegalAccountInput) HasBirth() bool`

HasBirth returns a boolean if a field has been set.

### GetNationality

`func (o *RegisterLegalAccountInput) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *RegisterLegalAccountInput) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *RegisterLegalAccountInput) SetNationality(v string)`

SetNationality sets Nationality field to given value.


### GetPhoneNumber

`func (o *RegisterLegalAccountInput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *RegisterLegalAccountInput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *RegisterLegalAccountInput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *RegisterLegalAccountInput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetMobileNumber

`func (o *RegisterLegalAccountInput) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *RegisterLegalAccountInput) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *RegisterLegalAccountInput) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.

### HasMobileNumber

`func (o *RegisterLegalAccountInput) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.

### GetIsDebtor

`func (o *RegisterLegalAccountInput) GetIsDebtor() bool`

GetIsDebtor returns the IsDebtor field if non-nil, zero value otherwise.

### GetIsDebtorOk

`func (o *RegisterLegalAccountInput) GetIsDebtorOk() (*bool, bool)`

GetIsDebtorOk returns a tuple with the IsDebtor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDebtor

`func (o *RegisterLegalAccountInput) SetIsDebtor(v bool)`

SetIsDebtor sets IsDebtor field to given value.

### HasIsDebtor

`func (o *RegisterLegalAccountInput) HasIsDebtor() bool`

HasIsDebtor returns a boolean if a field has been set.

### GetPayerOrBeneficiary

`func (o *RegisterLegalAccountInput) GetPayerOrBeneficiary() int32`

GetPayerOrBeneficiary returns the PayerOrBeneficiary field if non-nil, zero value otherwise.

### GetPayerOrBeneficiaryOk

`func (o *RegisterLegalAccountInput) GetPayerOrBeneficiaryOk() (*int32, bool)`

GetPayerOrBeneficiaryOk returns a tuple with the PayerOrBeneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerOrBeneficiary

`func (o *RegisterLegalAccountInput) SetPayerOrBeneficiary(v int32)`

SetPayerOrBeneficiary sets PayerOrBeneficiary field to given value.


### GetIsOneTimeCustomerAccount

`func (o *RegisterLegalAccountInput) GetIsOneTimeCustomerAccount() bool`

GetIsOneTimeCustomerAccount returns the IsOneTimeCustomerAccount field if non-nil, zero value otherwise.

### GetIsOneTimeCustomerAccountOk

`func (o *RegisterLegalAccountInput) GetIsOneTimeCustomerAccountOk() (*bool, bool)`

GetIsOneTimeCustomerAccountOk returns a tuple with the IsOneTimeCustomerAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOneTimeCustomerAccount

`func (o *RegisterLegalAccountInput) SetIsOneTimeCustomerAccount(v bool)`

SetIsOneTimeCustomerAccount sets IsOneTimeCustomerAccount field to given value.

### HasIsOneTimeCustomerAccount

`func (o *RegisterLegalAccountInput) HasIsOneTimeCustomerAccount() bool`

HasIsOneTimeCustomerAccount returns a boolean if a field has been set.

### GetIsTechnicalAccount

`func (o *RegisterLegalAccountInput) GetIsTechnicalAccount() bool`

GetIsTechnicalAccount returns the IsTechnicalAccount field if non-nil, zero value otherwise.

### GetIsTechnicalAccountOk

`func (o *RegisterLegalAccountInput) GetIsTechnicalAccountOk() (*bool, bool)`

GetIsTechnicalAccountOk returns a tuple with the IsTechnicalAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTechnicalAccount

`func (o *RegisterLegalAccountInput) SetIsTechnicalAccount(v bool)`

SetIsTechnicalAccount sets IsTechnicalAccount field to given value.

### HasIsTechnicalAccount

`func (o *RegisterLegalAccountInput) HasIsTechnicalAccount() bool`

HasIsTechnicalAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


