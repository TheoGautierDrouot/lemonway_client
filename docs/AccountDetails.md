# AccountDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Payment Account ID | [optional] 
**InternalId** | Pointer to **int64** | Internal ID given by Lemonway (ID displayed on the Dashboard) | [optional] 
**ClientTitle** | Pointer to **string** | Client Title | [optional] 
**Firstname** | Pointer to **string** | Holder&#39;s First Name | [optional] 
**Lastname** | Pointer to **string** | Holder&#39;s Last Name | [optional] 
**Balance** | Pointer to **int32** | Payment Account Balance Amount | [optional] 
**Email** | Pointer to **string** | Email | [optional] 
**Status** | Pointer to **int32** | | Status Code                                                     | Payment Account Status                                         |  |-----------------------------------------------------------------|----------------------------------------------------------------|  | -1                                                              | Wallet SC                                                      |  | 2                                                               | Registered, not verified: missing document(s)                  |  | 3                                                               | Registered, not verified: document(s) rejected                 |  | 5                                                               | Registered, not verified: KYC 1 (status given at registration) |  | 6                                                               | Registered and verified: KYC 2                                 |  | 7*                                                              | Registered and verified by previous PSP: KYC 3                 |  | 8                                                               | Registered, not verified: expired document(s)                  |  | 10                                                              | Blocked                                                        |  | 12                                                              | Closed                                                         |  | 13*                                                             | Registered, status is being updated from KYC 2 to KYC 3        |  | 14*                                                             | One-time customer                                              |  | 15*                                                             | Special wallet for crowdlending                                |  | 16*                                                             | Technical wallet                                               |  | *these status may or may not be used depending on your business |                                                                | | [optional] 
**Isblocked** | Pointer to **bool** | Indicates if Account is blocked or not:  0: Not blocked  1: Blocked | [optional] 
**AccountType** | Pointer to **int32** | Indicates Account type:  0: Individual  1: Legal | [optional] 
**Company** | Pointer to [**Company**](Company.md) |  | [optional] 
**Adresse** | Pointer to [**Address**](Address.md) |  | [optional] 
**Birth** | Pointer to [**Birth**](Birth.md) |  | [optional] 
**Nationality** | Pointer to **string** | Nationality | [optional] 
**PhoneNumber** | Pointer to **string** | Phone number with MSISDN format | [optional] 
**MobileNumber** | Pointer to **string** | Mobile phone number with MSISDN format | [optional] 
**IsDebtor** | Pointer to **bool** | For crowdfunding/loan platforms, indicates if the Account is created for a debtor   0 or empty: no  1: yes | [optional] 
**PayerOrBeneficiary** | Pointer to **int32** | Indicates if the payment Account is created for a payer or a beneficiary:  * Empty: unknown status (default)  * 1 &#x3D; Payer  * 2 &#x3D; Beneficiary | [optional] 

## Methods

### NewAccountDetails

`func NewAccountDetails() *AccountDetails`

NewAccountDetails instantiates a new AccountDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDetailsWithDefaults

`func NewAccountDetailsWithDefaults() *AccountDetails`

NewAccountDetailsWithDefaults instantiates a new AccountDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalId

`func (o *AccountDetails) GetInternalId() int64`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *AccountDetails) GetInternalIdOk() (*int64, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *AccountDetails) SetInternalId(v int64)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *AccountDetails) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetClientTitle

`func (o *AccountDetails) GetClientTitle() string`

GetClientTitle returns the ClientTitle field if non-nil, zero value otherwise.

### GetClientTitleOk

`func (o *AccountDetails) GetClientTitleOk() (*string, bool)`

GetClientTitleOk returns a tuple with the ClientTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTitle

`func (o *AccountDetails) SetClientTitle(v string)`

SetClientTitle sets ClientTitle field to given value.

### HasClientTitle

`func (o *AccountDetails) HasClientTitle() bool`

HasClientTitle returns a boolean if a field has been set.

### GetFirstname

`func (o *AccountDetails) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *AccountDetails) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *AccountDetails) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *AccountDetails) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *AccountDetails) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *AccountDetails) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *AccountDetails) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *AccountDetails) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetBalance

`func (o *AccountDetails) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *AccountDetails) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *AccountDetails) SetBalance(v int32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *AccountDetails) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetEmail

`func (o *AccountDetails) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountDetails) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountDetails) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AccountDetails) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetStatus

`func (o *AccountDetails) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountDetails) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountDetails) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsblocked

`func (o *AccountDetails) GetIsblocked() bool`

GetIsblocked returns the Isblocked field if non-nil, zero value otherwise.

### GetIsblockedOk

`func (o *AccountDetails) GetIsblockedOk() (*bool, bool)`

GetIsblockedOk returns a tuple with the Isblocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsblocked

`func (o *AccountDetails) SetIsblocked(v bool)`

SetIsblocked sets Isblocked field to given value.

### HasIsblocked

`func (o *AccountDetails) HasIsblocked() bool`

HasIsblocked returns a boolean if a field has been set.

### GetAccountType

`func (o *AccountDetails) GetAccountType() int32`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountDetails) GetAccountTypeOk() (*int32, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountDetails) SetAccountType(v int32)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *AccountDetails) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetCompany

`func (o *AccountDetails) GetCompany() Company`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *AccountDetails) GetCompanyOk() (*Company, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *AccountDetails) SetCompany(v Company)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *AccountDetails) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetAdresse

`func (o *AccountDetails) GetAdresse() Address`

GetAdresse returns the Adresse field if non-nil, zero value otherwise.

### GetAdresseOk

`func (o *AccountDetails) GetAdresseOk() (*Address, bool)`

GetAdresseOk returns a tuple with the Adresse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresse

`func (o *AccountDetails) SetAdresse(v Address)`

SetAdresse sets Adresse field to given value.

### HasAdresse

`func (o *AccountDetails) HasAdresse() bool`

HasAdresse returns a boolean if a field has been set.

### GetBirth

`func (o *AccountDetails) GetBirth() Birth`

GetBirth returns the Birth field if non-nil, zero value otherwise.

### GetBirthOk

`func (o *AccountDetails) GetBirthOk() (*Birth, bool)`

GetBirthOk returns a tuple with the Birth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirth

`func (o *AccountDetails) SetBirth(v Birth)`

SetBirth sets Birth field to given value.

### HasBirth

`func (o *AccountDetails) HasBirth() bool`

HasBirth returns a boolean if a field has been set.

### GetNationality

`func (o *AccountDetails) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *AccountDetails) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *AccountDetails) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *AccountDetails) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *AccountDetails) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *AccountDetails) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *AccountDetails) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *AccountDetails) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetMobileNumber

`func (o *AccountDetails) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *AccountDetails) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *AccountDetails) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.

### HasMobileNumber

`func (o *AccountDetails) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.

### GetIsDebtor

`func (o *AccountDetails) GetIsDebtor() bool`

GetIsDebtor returns the IsDebtor field if non-nil, zero value otherwise.

### GetIsDebtorOk

`func (o *AccountDetails) GetIsDebtorOk() (*bool, bool)`

GetIsDebtorOk returns a tuple with the IsDebtor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDebtor

`func (o *AccountDetails) SetIsDebtor(v bool)`

SetIsDebtor sets IsDebtor field to given value.

### HasIsDebtor

`func (o *AccountDetails) HasIsDebtor() bool`

HasIsDebtor returns a boolean if a field has been set.

### GetPayerOrBeneficiary

`func (o *AccountDetails) GetPayerOrBeneficiary() int32`

GetPayerOrBeneficiary returns the PayerOrBeneficiary field if non-nil, zero value otherwise.

### GetPayerOrBeneficiaryOk

`func (o *AccountDetails) GetPayerOrBeneficiaryOk() (*int32, bool)`

GetPayerOrBeneficiaryOk returns a tuple with the PayerOrBeneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerOrBeneficiary

`func (o *AccountDetails) SetPayerOrBeneficiary(v int32)`

SetPayerOrBeneficiary sets PayerOrBeneficiary field to given value.

### HasPayerOrBeneficiary

`func (o *AccountDetails) HasPayerOrBeneficiary() bool`

HasPayerOrBeneficiary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


