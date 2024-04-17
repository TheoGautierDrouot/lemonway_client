# CardInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardType** | **int32** | Card Type&lt;br/&gt;0 &#x3D; CB.&lt;br/&gt;1 &#x3D; Visa.&lt;br/&gt;2 &#x3D; Mastercard.&lt;br/&gt; | 
**CardNumber** | **string** | Card Number | 
**CardCode** | **string** | CVV Code | 
**CardDate** | **string** | Card Expiration Date | 

## Methods

### NewCardInfo

`func NewCardInfo(cardType int32, cardNumber string, cardCode string, cardDate string, ) *CardInfo`

NewCardInfo instantiates a new CardInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardInfoWithDefaults

`func NewCardInfoWithDefaults() *CardInfo`

NewCardInfoWithDefaults instantiates a new CardInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardType

`func (o *CardInfo) GetCardType() int32`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *CardInfo) GetCardTypeOk() (*int32, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *CardInfo) SetCardType(v int32)`

SetCardType sets CardType field to given value.


### GetCardNumber

`func (o *CardInfo) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *CardInfo) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *CardInfo) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.


### GetCardCode

`func (o *CardInfo) GetCardCode() string`

GetCardCode returns the CardCode field if non-nil, zero value otherwise.

### GetCardCodeOk

`func (o *CardInfo) GetCardCodeOk() (*string, bool)`

GetCardCodeOk returns a tuple with the CardCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCode

`func (o *CardInfo) SetCardCode(v string)`

SetCardCode sets CardCode field to given value.


### GetCardDate

`func (o *CardInfo) GetCardDate() string`

GetCardDate returns the CardDate field if non-nil, zero value otherwise.

### GetCardDateOk

`func (o *CardInfo) GetCardDateOk() (*string, bool)`

GetCardDateOk returns a tuple with the CardDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardDate

`func (o *CardInfo) SetCardDate(v string)`

SetCardDate sets CardDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


