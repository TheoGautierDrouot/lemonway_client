# RegisterSddMandateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wallet** | **string** | Payment Account ID | 
**Holder** | **string** | IBAN Holder | 
**Bic** | **string** | BIC Code (SWIFT) | 
**Iban** | **string** | IBAN | 
**IsB2B** | Pointer to **bool** | Indicates if the mandate is a SEPA business to business direct debit | [optional] 
**IsRecurring** | **bool** | 1. If true, mandate is for recurring payments  2. If false, mandate is for one payment only | 
**Street** | Pointer to **string** | Client Number and Street  Mandatory if you plan to electronically sign the mandate | [optional] 
**PostCode** | Pointer to **string** | Client Postal Code/ZIP  Mandatory if you plan to electronically sign the mandate | [optional] 
**City** | Pointer to **string** | Client City  Mandatory if you plan to electronically sign the mandate | [optional] 
**Country** | Pointer to **string** | Client Country  Mandatory if you plan to electronically sign the mandate. | [optional] 
**MandateLanguage** | Pointer to **string** | Language of the automatically generated mandate(for electronic signature).  Currently, the available languages are: fr, en, de, es, it.  If you need another, please contact us.   \&quot;fr\&quot; is used by default if no language is requested. | [optional] 
**Rum** | Pointer to **string** | Unique Mandate Reference Number (UMR) | [optional] 

## Methods

### NewRegisterSddMandateInput

`func NewRegisterSddMandateInput(wallet string, holder string, bic string, iban string, isRecurring bool, ) *RegisterSddMandateInput`

NewRegisterSddMandateInput instantiates a new RegisterSddMandateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterSddMandateInputWithDefaults

`func NewRegisterSddMandateInputWithDefaults() *RegisterSddMandateInput`

NewRegisterSddMandateInputWithDefaults instantiates a new RegisterSddMandateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWallet

`func (o *RegisterSddMandateInput) GetWallet() string`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *RegisterSddMandateInput) GetWalletOk() (*string, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *RegisterSddMandateInput) SetWallet(v string)`

SetWallet sets Wallet field to given value.


### GetHolder

`func (o *RegisterSddMandateInput) GetHolder() string`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *RegisterSddMandateInput) GetHolderOk() (*string, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *RegisterSddMandateInput) SetHolder(v string)`

SetHolder sets Holder field to given value.


### GetBic

`func (o *RegisterSddMandateInput) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *RegisterSddMandateInput) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *RegisterSddMandateInput) SetBic(v string)`

SetBic sets Bic field to given value.


### GetIban

`func (o *RegisterSddMandateInput) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *RegisterSddMandateInput) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *RegisterSddMandateInput) SetIban(v string)`

SetIban sets Iban field to given value.


### GetIsB2B

`func (o *RegisterSddMandateInput) GetIsB2B() bool`

GetIsB2B returns the IsB2B field if non-nil, zero value otherwise.

### GetIsB2BOk

`func (o *RegisterSddMandateInput) GetIsB2BOk() (*bool, bool)`

GetIsB2BOk returns a tuple with the IsB2B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsB2B

`func (o *RegisterSddMandateInput) SetIsB2B(v bool)`

SetIsB2B sets IsB2B field to given value.

### HasIsB2B

`func (o *RegisterSddMandateInput) HasIsB2B() bool`

HasIsB2B returns a boolean if a field has been set.

### GetIsRecurring

`func (o *RegisterSddMandateInput) GetIsRecurring() bool`

GetIsRecurring returns the IsRecurring field if non-nil, zero value otherwise.

### GetIsRecurringOk

`func (o *RegisterSddMandateInput) GetIsRecurringOk() (*bool, bool)`

GetIsRecurringOk returns a tuple with the IsRecurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecurring

`func (o *RegisterSddMandateInput) SetIsRecurring(v bool)`

SetIsRecurring sets IsRecurring field to given value.


### GetStreet

`func (o *RegisterSddMandateInput) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *RegisterSddMandateInput) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *RegisterSddMandateInput) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *RegisterSddMandateInput) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetPostCode

`func (o *RegisterSddMandateInput) GetPostCode() string`

GetPostCode returns the PostCode field if non-nil, zero value otherwise.

### GetPostCodeOk

`func (o *RegisterSddMandateInput) GetPostCodeOk() (*string, bool)`

GetPostCodeOk returns a tuple with the PostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCode

`func (o *RegisterSddMandateInput) SetPostCode(v string)`

SetPostCode sets PostCode field to given value.

### HasPostCode

`func (o *RegisterSddMandateInput) HasPostCode() bool`

HasPostCode returns a boolean if a field has been set.

### GetCity

`func (o *RegisterSddMandateInput) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *RegisterSddMandateInput) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *RegisterSddMandateInput) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *RegisterSddMandateInput) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *RegisterSddMandateInput) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *RegisterSddMandateInput) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *RegisterSddMandateInput) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *RegisterSddMandateInput) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetMandateLanguage

`func (o *RegisterSddMandateInput) GetMandateLanguage() string`

GetMandateLanguage returns the MandateLanguage field if non-nil, zero value otherwise.

### GetMandateLanguageOk

`func (o *RegisterSddMandateInput) GetMandateLanguageOk() (*string, bool)`

GetMandateLanguageOk returns a tuple with the MandateLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateLanguage

`func (o *RegisterSddMandateInput) SetMandateLanguage(v string)`

SetMandateLanguage sets MandateLanguage field to given value.

### HasMandateLanguage

`func (o *RegisterSddMandateInput) HasMandateLanguage() bool`

HasMandateLanguage returns a boolean if a field has been set.

### GetRum

`func (o *RegisterSddMandateInput) GetRum() string`

GetRum returns the Rum field if non-nil, zero value otherwise.

### GetRumOk

`func (o *RegisterSddMandateInput) GetRumOk() (*string, bool)`

GetRumOk returns a tuple with the Rum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRum

`func (o *RegisterSddMandateInput) SetRum(v string)`

SetRum sets Rum field to given value.

### HasRum

`func (o *RegisterSddMandateInput) HasRum() bool`

HasRum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


