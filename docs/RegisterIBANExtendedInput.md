# RegisterIBANExtendedInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wallet** | **string** | Payment Account ID | 
**AccountType** | Pointer to **int32** | Account Type, by default Other&lt;br/&gt;0 &#x3D; Other.&lt;br/&gt;1 &#x3D; IBAN.&lt;br/&gt;2 &#x3D; BBAN/RIB.&lt;br/&gt; | [optional] 
**HolderName** | **string** | The registered bank account owner: First and Last name, or Company Name | 
**AccountNumber** | **string** | Account Number. The format depends on the account type. | 
**HolderCountry** | Pointer to **string** | Country of the beneficiary. Two-letter country code (ISO alpha-2) for example, France&#x3D;&lt;b&gt;FR&lt;/b&gt; | [optional] 
**BicCode** | Pointer to **string** | BIC/SWIFT codes are arranged like this: AAAABBCCDDD  AAAA: 4 character for bank code  BB: 2 char for country code  CC: 2 char for location code  DDD: 3 char for branch code | [optional] 
**BankName** | Pointer to **string** | Bank Name. This field is mandatory in the following circumstances:  &lt;ul&gt;&lt;li&gt;If the currency is USD&lt;/li&gt;&lt;li&gt;If the selected bank country code requires this field&lt;/li&gt;&lt;/ul&gt; | [optional] 
**BankCountry** | **string** | Country of the Bank. Two-letter country code (ISO alpha-2) for example, France&#x3D;&lt;b&gt;FR&lt;/b&gt; | 
**BankBranchCode** | Pointer to **string** | Bank Branch Code (Sort Code in the United Kingdom). It is mandatory for the examples below:  &lt;ul&gt;&lt;li&gt;If the selected bank country code requires this field.&lt;/li&gt;&lt;/ul&gt; | [optional] 
**BankBranchAddress** | Pointer to [**BankBranchAddress**](BankBranchAddress.md) |  | [optional] 
**IntermediaryBicCode** | Pointer to **string** | BIC/SWIFT Code of the Intermediary Bank. | [optional] 
**IntermediaryBankName** | Pointer to **string** | Intermediary Bank Name | [optional] 
**IntermediaryBankCountry** | Pointer to **string** | Bank Country Code of the Intermediary Bank. Two-letter country code (ISO alpha-2) for example, France&#x3D;&lt;b&gt;FR&lt;/b&gt; | [optional] 
**Comment** | Pointer to **string** | Reason for New Bank Account Details if another one is already linked to the Payment Account. | [optional] 

## Methods

### NewRegisterIBANExtendedInput

`func NewRegisterIBANExtendedInput(wallet string, holderName string, accountNumber string, bankCountry string, ) *RegisterIBANExtendedInput`

NewRegisterIBANExtendedInput instantiates a new RegisterIBANExtendedInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterIBANExtendedInputWithDefaults

`func NewRegisterIBANExtendedInputWithDefaults() *RegisterIBANExtendedInput`

NewRegisterIBANExtendedInputWithDefaults instantiates a new RegisterIBANExtendedInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWallet

`func (o *RegisterIBANExtendedInput) GetWallet() string`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *RegisterIBANExtendedInput) GetWalletOk() (*string, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *RegisterIBANExtendedInput) SetWallet(v string)`

SetWallet sets Wallet field to given value.


### GetAccountType

`func (o *RegisterIBANExtendedInput) GetAccountType() int32`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *RegisterIBANExtendedInput) GetAccountTypeOk() (*int32, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *RegisterIBANExtendedInput) SetAccountType(v int32)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *RegisterIBANExtendedInput) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetHolderName

`func (o *RegisterIBANExtendedInput) GetHolderName() string`

GetHolderName returns the HolderName field if non-nil, zero value otherwise.

### GetHolderNameOk

`func (o *RegisterIBANExtendedInput) GetHolderNameOk() (*string, bool)`

GetHolderNameOk returns a tuple with the HolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderName

`func (o *RegisterIBANExtendedInput) SetHolderName(v string)`

SetHolderName sets HolderName field to given value.


### GetAccountNumber

`func (o *RegisterIBANExtendedInput) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *RegisterIBANExtendedInput) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *RegisterIBANExtendedInput) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetHolderCountry

`func (o *RegisterIBANExtendedInput) GetHolderCountry() string`

GetHolderCountry returns the HolderCountry field if non-nil, zero value otherwise.

### GetHolderCountryOk

`func (o *RegisterIBANExtendedInput) GetHolderCountryOk() (*string, bool)`

GetHolderCountryOk returns a tuple with the HolderCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderCountry

`func (o *RegisterIBANExtendedInput) SetHolderCountry(v string)`

SetHolderCountry sets HolderCountry field to given value.

### HasHolderCountry

`func (o *RegisterIBANExtendedInput) HasHolderCountry() bool`

HasHolderCountry returns a boolean if a field has been set.

### GetBicCode

`func (o *RegisterIBANExtendedInput) GetBicCode() string`

GetBicCode returns the BicCode field if non-nil, zero value otherwise.

### GetBicCodeOk

`func (o *RegisterIBANExtendedInput) GetBicCodeOk() (*string, bool)`

GetBicCodeOk returns a tuple with the BicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBicCode

`func (o *RegisterIBANExtendedInput) SetBicCode(v string)`

SetBicCode sets BicCode field to given value.

### HasBicCode

`func (o *RegisterIBANExtendedInput) HasBicCode() bool`

HasBicCode returns a boolean if a field has been set.

### GetBankName

`func (o *RegisterIBANExtendedInput) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *RegisterIBANExtendedInput) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *RegisterIBANExtendedInput) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *RegisterIBANExtendedInput) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetBankCountry

`func (o *RegisterIBANExtendedInput) GetBankCountry() string`

GetBankCountry returns the BankCountry field if non-nil, zero value otherwise.

### GetBankCountryOk

`func (o *RegisterIBANExtendedInput) GetBankCountryOk() (*string, bool)`

GetBankCountryOk returns a tuple with the BankCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCountry

`func (o *RegisterIBANExtendedInput) SetBankCountry(v string)`

SetBankCountry sets BankCountry field to given value.


### GetBankBranchCode

`func (o *RegisterIBANExtendedInput) GetBankBranchCode() string`

GetBankBranchCode returns the BankBranchCode field if non-nil, zero value otherwise.

### GetBankBranchCodeOk

`func (o *RegisterIBANExtendedInput) GetBankBranchCodeOk() (*string, bool)`

GetBankBranchCodeOk returns a tuple with the BankBranchCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankBranchCode

`func (o *RegisterIBANExtendedInput) SetBankBranchCode(v string)`

SetBankBranchCode sets BankBranchCode field to given value.

### HasBankBranchCode

`func (o *RegisterIBANExtendedInput) HasBankBranchCode() bool`

HasBankBranchCode returns a boolean if a field has been set.

### GetBankBranchAddress

`func (o *RegisterIBANExtendedInput) GetBankBranchAddress() BankBranchAddress`

GetBankBranchAddress returns the BankBranchAddress field if non-nil, zero value otherwise.

### GetBankBranchAddressOk

`func (o *RegisterIBANExtendedInput) GetBankBranchAddressOk() (*BankBranchAddress, bool)`

GetBankBranchAddressOk returns a tuple with the BankBranchAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankBranchAddress

`func (o *RegisterIBANExtendedInput) SetBankBranchAddress(v BankBranchAddress)`

SetBankBranchAddress sets BankBranchAddress field to given value.

### HasBankBranchAddress

`func (o *RegisterIBANExtendedInput) HasBankBranchAddress() bool`

HasBankBranchAddress returns a boolean if a field has been set.

### GetIntermediaryBicCode

`func (o *RegisterIBANExtendedInput) GetIntermediaryBicCode() string`

GetIntermediaryBicCode returns the IntermediaryBicCode field if non-nil, zero value otherwise.

### GetIntermediaryBicCodeOk

`func (o *RegisterIBANExtendedInput) GetIntermediaryBicCodeOk() (*string, bool)`

GetIntermediaryBicCodeOk returns a tuple with the IntermediaryBicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediaryBicCode

`func (o *RegisterIBANExtendedInput) SetIntermediaryBicCode(v string)`

SetIntermediaryBicCode sets IntermediaryBicCode field to given value.

### HasIntermediaryBicCode

`func (o *RegisterIBANExtendedInput) HasIntermediaryBicCode() bool`

HasIntermediaryBicCode returns a boolean if a field has been set.

### GetIntermediaryBankName

`func (o *RegisterIBANExtendedInput) GetIntermediaryBankName() string`

GetIntermediaryBankName returns the IntermediaryBankName field if non-nil, zero value otherwise.

### GetIntermediaryBankNameOk

`func (o *RegisterIBANExtendedInput) GetIntermediaryBankNameOk() (*string, bool)`

GetIntermediaryBankNameOk returns a tuple with the IntermediaryBankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediaryBankName

`func (o *RegisterIBANExtendedInput) SetIntermediaryBankName(v string)`

SetIntermediaryBankName sets IntermediaryBankName field to given value.

### HasIntermediaryBankName

`func (o *RegisterIBANExtendedInput) HasIntermediaryBankName() bool`

HasIntermediaryBankName returns a boolean if a field has been set.

### GetIntermediaryBankCountry

`func (o *RegisterIBANExtendedInput) GetIntermediaryBankCountry() string`

GetIntermediaryBankCountry returns the IntermediaryBankCountry field if non-nil, zero value otherwise.

### GetIntermediaryBankCountryOk

`func (o *RegisterIBANExtendedInput) GetIntermediaryBankCountryOk() (*string, bool)`

GetIntermediaryBankCountryOk returns a tuple with the IntermediaryBankCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediaryBankCountry

`func (o *RegisterIBANExtendedInput) SetIntermediaryBankCountry(v string)`

SetIntermediaryBankCountry sets IntermediaryBankCountry field to given value.

### HasIntermediaryBankCountry

`func (o *RegisterIBANExtendedInput) HasIntermediaryBankCountry() bool`

HasIntermediaryBankCountry returns a boolean if a field has been set.

### GetComment

`func (o *RegisterIBANExtendedInput) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RegisterIBANExtendedInput) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RegisterIBANExtendedInput) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RegisterIBANExtendedInput) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


