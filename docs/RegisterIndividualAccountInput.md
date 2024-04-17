# RegisterIndividualAccountInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**IsUltimateBeneficialOwner** | Pointer to **bool** |  | [optional] 

## Methods

### NewRegisterIndividualAccountInput

`func NewRegisterIndividualAccountInput(accountId string, email string, firstName string, lastName string, nationality string, payerOrBeneficiary int32, ) *RegisterIndividualAccountInput`

NewRegisterIndividualAccountInput instantiates a new RegisterIndividualAccountInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterIndividualAccountInputWithDefaults

`func NewRegisterIndividualAccountInputWithDefaults() *RegisterIndividualAccountInput`

NewRegisterIndividualAccountInputWithDefaults instantiates a new RegisterIndividualAccountInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *RegisterIndividualAccountInput) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RegisterIndividualAccountInput) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RegisterIndividualAccountInput) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetEmail

`func (o *RegisterIndividualAccountInput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RegisterIndividualAccountInput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RegisterIndividualAccountInput) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetTitle

`func (o *RegisterIndividualAccountInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RegisterIndividualAccountInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RegisterIndividualAccountInput) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RegisterIndividualAccountInput) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFirstName

`func (o *RegisterIndividualAccountInput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *RegisterIndividualAccountInput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *RegisterIndividualAccountInput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *RegisterIndividualAccountInput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *RegisterIndividualAccountInput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *RegisterIndividualAccountInput) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetAdresse

`func (o *RegisterIndividualAccountInput) GetAdresse() Address`

GetAdresse returns the Adresse field if non-nil, zero value otherwise.

### GetAdresseOk

`func (o *RegisterIndividualAccountInput) GetAdresseOk() (*Address, bool)`

GetAdresseOk returns a tuple with the Adresse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresse

`func (o *RegisterIndividualAccountInput) SetAdresse(v Address)`

SetAdresse sets Adresse field to given value.

### HasAdresse

`func (o *RegisterIndividualAccountInput) HasAdresse() bool`

HasAdresse returns a boolean if a field has been set.

### GetBirth

`func (o *RegisterIndividualAccountInput) GetBirth() Birth`

GetBirth returns the Birth field if non-nil, zero value otherwise.

### GetBirthOk

`func (o *RegisterIndividualAccountInput) GetBirthOk() (*Birth, bool)`

GetBirthOk returns a tuple with the Birth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirth

`func (o *RegisterIndividualAccountInput) SetBirth(v Birth)`

SetBirth sets Birth field to given value.

### HasBirth

`func (o *RegisterIndividualAccountInput) HasBirth() bool`

HasBirth returns a boolean if a field has been set.

### GetNationality

`func (o *RegisterIndividualAccountInput) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *RegisterIndividualAccountInput) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *RegisterIndividualAccountInput) SetNationality(v string)`

SetNationality sets Nationality field to given value.


### GetPhoneNumber

`func (o *RegisterIndividualAccountInput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *RegisterIndividualAccountInput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *RegisterIndividualAccountInput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *RegisterIndividualAccountInput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetMobileNumber

`func (o *RegisterIndividualAccountInput) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *RegisterIndividualAccountInput) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *RegisterIndividualAccountInput) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.

### HasMobileNumber

`func (o *RegisterIndividualAccountInput) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.

### GetIsDebtor

`func (o *RegisterIndividualAccountInput) GetIsDebtor() bool`

GetIsDebtor returns the IsDebtor field if non-nil, zero value otherwise.

### GetIsDebtorOk

`func (o *RegisterIndividualAccountInput) GetIsDebtorOk() (*bool, bool)`

GetIsDebtorOk returns a tuple with the IsDebtor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDebtor

`func (o *RegisterIndividualAccountInput) SetIsDebtor(v bool)`

SetIsDebtor sets IsDebtor field to given value.

### HasIsDebtor

`func (o *RegisterIndividualAccountInput) HasIsDebtor() bool`

HasIsDebtor returns a boolean if a field has been set.

### GetPayerOrBeneficiary

`func (o *RegisterIndividualAccountInput) GetPayerOrBeneficiary() int32`

GetPayerOrBeneficiary returns the PayerOrBeneficiary field if non-nil, zero value otherwise.

### GetPayerOrBeneficiaryOk

`func (o *RegisterIndividualAccountInput) GetPayerOrBeneficiaryOk() (*int32, bool)`

GetPayerOrBeneficiaryOk returns a tuple with the PayerOrBeneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerOrBeneficiary

`func (o *RegisterIndividualAccountInput) SetPayerOrBeneficiary(v int32)`

SetPayerOrBeneficiary sets PayerOrBeneficiary field to given value.


### GetIsOneTimeCustomerAccount

`func (o *RegisterIndividualAccountInput) GetIsOneTimeCustomerAccount() bool`

GetIsOneTimeCustomerAccount returns the IsOneTimeCustomerAccount field if non-nil, zero value otherwise.

### GetIsOneTimeCustomerAccountOk

`func (o *RegisterIndividualAccountInput) GetIsOneTimeCustomerAccountOk() (*bool, bool)`

GetIsOneTimeCustomerAccountOk returns a tuple with the IsOneTimeCustomerAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOneTimeCustomerAccount

`func (o *RegisterIndividualAccountInput) SetIsOneTimeCustomerAccount(v bool)`

SetIsOneTimeCustomerAccount sets IsOneTimeCustomerAccount field to given value.

### HasIsOneTimeCustomerAccount

`func (o *RegisterIndividualAccountInput) HasIsOneTimeCustomerAccount() bool`

HasIsOneTimeCustomerAccount returns a boolean if a field has been set.

### GetIsTechnicalAccount

`func (o *RegisterIndividualAccountInput) GetIsTechnicalAccount() bool`

GetIsTechnicalAccount returns the IsTechnicalAccount field if non-nil, zero value otherwise.

### GetIsTechnicalAccountOk

`func (o *RegisterIndividualAccountInput) GetIsTechnicalAccountOk() (*bool, bool)`

GetIsTechnicalAccountOk returns a tuple with the IsTechnicalAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTechnicalAccount

`func (o *RegisterIndividualAccountInput) SetIsTechnicalAccount(v bool)`

SetIsTechnicalAccount sets IsTechnicalAccount field to given value.

### HasIsTechnicalAccount

`func (o *RegisterIndividualAccountInput) HasIsTechnicalAccount() bool`

HasIsTechnicalAccount returns a boolean if a field has been set.

### GetIsUltimateBeneficialOwner

`func (o *RegisterIndividualAccountInput) GetIsUltimateBeneficialOwner() bool`

GetIsUltimateBeneficialOwner returns the IsUltimateBeneficialOwner field if non-nil, zero value otherwise.

### GetIsUltimateBeneficialOwnerOk

`func (o *RegisterIndividualAccountInput) GetIsUltimateBeneficialOwnerOk() (*bool, bool)`

GetIsUltimateBeneficialOwnerOk returns a tuple with the IsUltimateBeneficialOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUltimateBeneficialOwner

`func (o *RegisterIndividualAccountInput) SetIsUltimateBeneficialOwner(v bool)`

SetIsUltimateBeneficialOwner sets IsUltimateBeneficialOwner field to given value.

### HasIsUltimateBeneficialOwner

`func (o *RegisterIndividualAccountInput) HasIsUltimateBeneficialOwner() bool`

HasIsUltimateBeneficialOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


