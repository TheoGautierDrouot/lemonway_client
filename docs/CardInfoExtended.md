# CardInfoExtended

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardHolder** | Pointer to **string** | Card Holder | [optional] 
**CardType** | **int32** | Card Type&lt;br/&gt;0 &#x3D; CB.&lt;br/&gt;1 &#x3D; Visa.&lt;br/&gt;2 &#x3D; Mastercard.&lt;br/&gt; | 
**CardNumber** | **string** | Masked Card Number | 
**CardDate** | **string** | Card Expiration Date | 
**CardCountry** | Pointer to **string** | Card Country | [optional] 

## Methods

### NewCardInfoExtended

`func NewCardInfoExtended(cardType int32, cardNumber string, cardDate string, ) *CardInfoExtended`

NewCardInfoExtended instantiates a new CardInfoExtended object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardInfoExtendedWithDefaults

`func NewCardInfoExtendedWithDefaults() *CardInfoExtended`

NewCardInfoExtendedWithDefaults instantiates a new CardInfoExtended object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardHolder

`func (o *CardInfoExtended) GetCardHolder() string`

GetCardHolder returns the CardHolder field if non-nil, zero value otherwise.

### GetCardHolderOk

`func (o *CardInfoExtended) GetCardHolderOk() (*string, bool)`

GetCardHolderOk returns a tuple with the CardHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardHolder

`func (o *CardInfoExtended) SetCardHolder(v string)`

SetCardHolder sets CardHolder field to given value.

### HasCardHolder

`func (o *CardInfoExtended) HasCardHolder() bool`

HasCardHolder returns a boolean if a field has been set.

### GetCardType

`func (o *CardInfoExtended) GetCardType() int32`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *CardInfoExtended) GetCardTypeOk() (*int32, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *CardInfoExtended) SetCardType(v int32)`

SetCardType sets CardType field to given value.


### GetCardNumber

`func (o *CardInfoExtended) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *CardInfoExtended) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *CardInfoExtended) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.


### GetCardDate

`func (o *CardInfoExtended) GetCardDate() string`

GetCardDate returns the CardDate field if non-nil, zero value otherwise.

### GetCardDateOk

`func (o *CardInfoExtended) GetCardDateOk() (*string, bool)`

GetCardDateOk returns a tuple with the CardDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardDate

`func (o *CardInfoExtended) SetCardDate(v string)`

SetCardDate sets CardDate field to given value.


### GetCardCountry

`func (o *CardInfoExtended) GetCardCountry() string`

GetCardCountry returns the CardCountry field if non-nil, zero value otherwise.

### GetCardCountryOk

`func (o *CardInfoExtended) GetCardCountryOk() (*string, bool)`

GetCardCountryOk returns a tuple with the CardCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCountry

`func (o *CardInfoExtended) SetCardCountry(v string)`

SetCardCountry sets CardCountry field to given value.

### HasCardCountry

`func (o *CardInfoExtended) HasCardCountry() bool`

HasCardCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


