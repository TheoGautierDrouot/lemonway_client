# MoneyIn3DConfirmInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPreAuth** | Pointer to **bool** | **Mercanet only**  Indicates if the request is delayed and will require validation.  1. If true, payment will only be pre-authorized, you will have to call the &#x60;MoneyInValidate&#x60; method within 6 days or [delayedDays] in order to request the debit card.  2. If empty, the default behavior is no delay and no validation is necessary (unless stated otherwise in your contract).  3. If false, the card will be debited without delay and without validation necessary. | [optional] 
**DelayedDays** | Pointer to **int32** | **Mercanet only**  If &#x60;isPreAuth&#x60; is not true, this will be ignored.   Please use with caution, if delayedDays &amp;gt; 6, it is possible that the payment will be denied when you try to validate it, because a new authorization will be made. | [optional] 
**Card** | Pointer to [**CardInfo**](CardInfo.md) |  | [optional] 
**Md** | Pointer to **string** | MD Data Returned by 3-D Secure Site | [optional] 
**Pares** | Pointer to **string** | Pares Data Returned by 3-D Secure Authentication Website | [optional] 
**SpecialConfig** | Pointer to **string** | Leave Empty | [optional] 

## Methods

### NewMoneyIn3DConfirmInput

`func NewMoneyIn3DConfirmInput() *MoneyIn3DConfirmInput`

NewMoneyIn3DConfirmInput instantiates a new MoneyIn3DConfirmInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoneyIn3DConfirmInputWithDefaults

`func NewMoneyIn3DConfirmInputWithDefaults() *MoneyIn3DConfirmInput`

NewMoneyIn3DConfirmInputWithDefaults instantiates a new MoneyIn3DConfirmInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPreAuth

`func (o *MoneyIn3DConfirmInput) GetIsPreAuth() bool`

GetIsPreAuth returns the IsPreAuth field if non-nil, zero value otherwise.

### GetIsPreAuthOk

`func (o *MoneyIn3DConfirmInput) GetIsPreAuthOk() (*bool, bool)`

GetIsPreAuthOk returns a tuple with the IsPreAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreAuth

`func (o *MoneyIn3DConfirmInput) SetIsPreAuth(v bool)`

SetIsPreAuth sets IsPreAuth field to given value.

### HasIsPreAuth

`func (o *MoneyIn3DConfirmInput) HasIsPreAuth() bool`

HasIsPreAuth returns a boolean if a field has been set.

### GetDelayedDays

`func (o *MoneyIn3DConfirmInput) GetDelayedDays() int32`

GetDelayedDays returns the DelayedDays field if non-nil, zero value otherwise.

### GetDelayedDaysOk

`func (o *MoneyIn3DConfirmInput) GetDelayedDaysOk() (*int32, bool)`

GetDelayedDaysOk returns a tuple with the DelayedDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayedDays

`func (o *MoneyIn3DConfirmInput) SetDelayedDays(v int32)`

SetDelayedDays sets DelayedDays field to given value.

### HasDelayedDays

`func (o *MoneyIn3DConfirmInput) HasDelayedDays() bool`

HasDelayedDays returns a boolean if a field has been set.

### GetCard

`func (o *MoneyIn3DConfirmInput) GetCard() CardInfo`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *MoneyIn3DConfirmInput) GetCardOk() (*CardInfo, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *MoneyIn3DConfirmInput) SetCard(v CardInfo)`

SetCard sets Card field to given value.

### HasCard

`func (o *MoneyIn3DConfirmInput) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetMd

`func (o *MoneyIn3DConfirmInput) GetMd() string`

GetMd returns the Md field if non-nil, zero value otherwise.

### GetMdOk

`func (o *MoneyIn3DConfirmInput) GetMdOk() (*string, bool)`

GetMdOk returns a tuple with the Md field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd

`func (o *MoneyIn3DConfirmInput) SetMd(v string)`

SetMd sets Md field to given value.

### HasMd

`func (o *MoneyIn3DConfirmInput) HasMd() bool`

HasMd returns a boolean if a field has been set.

### GetPares

`func (o *MoneyIn3DConfirmInput) GetPares() string`

GetPares returns the Pares field if non-nil, zero value otherwise.

### GetParesOk

`func (o *MoneyIn3DConfirmInput) GetParesOk() (*string, bool)`

GetParesOk returns a tuple with the Pares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPares

`func (o *MoneyIn3DConfirmInput) SetPares(v string)`

SetPares sets Pares field to given value.

### HasPares

`func (o *MoneyIn3DConfirmInput) HasPares() bool`

HasPares returns a boolean if a field has been set.

### GetSpecialConfig

`func (o *MoneyIn3DConfirmInput) GetSpecialConfig() string`

GetSpecialConfig returns the SpecialConfig field if non-nil, zero value otherwise.

### GetSpecialConfigOk

`func (o *MoneyIn3DConfirmInput) GetSpecialConfigOk() (*string, bool)`

GetSpecialConfigOk returns a tuple with the SpecialConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialConfig

`func (o *MoneyIn3DConfirmInput) SetSpecialConfig(v string)`

SetSpecialConfig sets SpecialConfig field to given value.

### HasSpecialConfig

`func (o *MoneyIn3DConfirmInput) HasSpecialConfig() bool`

HasSpecialConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


