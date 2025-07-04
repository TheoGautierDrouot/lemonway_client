# Card

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Is3DS** | Pointer to **bool** | Card with 3-DS authentication | [optional] 
**Country** | Pointer to **string** | Card issuing country | [optional] 
**AuthorizationNumber** | Pointer to **string** | Authorization number | [optional] 
**MaskedNumber** | Pointer to **string** | Masked card number | [optional] 
**Expiration** | Pointer to **string** | Expiration date (if available) | [optional] 
**Type** | Pointer to **string** | Card type (example: Visa or Mastercard) | [optional] 
**IsRegistered** | Pointer to **bool** | Card is registered | [optional] 
**HolderName** | Pointer to **string** | Card&#39;s holder name | [optional] 

## Methods

### NewCard

`func NewCard() *Card`

NewCard instantiates a new Card object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardWithDefaults

`func NewCardWithDefaults() *Card`

NewCardWithDefaults instantiates a new Card object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Card) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Card) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Card) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Card) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIs3DS

`func (o *Card) GetIs3DS() bool`

GetIs3DS returns the Is3DS field if non-nil, zero value otherwise.

### GetIs3DSOk

`func (o *Card) GetIs3DSOk() (*bool, bool)`

GetIs3DSOk returns a tuple with the Is3DS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIs3DS

`func (o *Card) SetIs3DS(v bool)`

SetIs3DS sets Is3DS field to given value.

### HasIs3DS

`func (o *Card) HasIs3DS() bool`

HasIs3DS returns a boolean if a field has been set.

### GetCountry

`func (o *Card) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Card) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Card) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Card) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetAuthorizationNumber

`func (o *Card) GetAuthorizationNumber() string`

GetAuthorizationNumber returns the AuthorizationNumber field if non-nil, zero value otherwise.

### GetAuthorizationNumberOk

`func (o *Card) GetAuthorizationNumberOk() (*string, bool)`

GetAuthorizationNumberOk returns a tuple with the AuthorizationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationNumber

`func (o *Card) SetAuthorizationNumber(v string)`

SetAuthorizationNumber sets AuthorizationNumber field to given value.

### HasAuthorizationNumber

`func (o *Card) HasAuthorizationNumber() bool`

HasAuthorizationNumber returns a boolean if a field has been set.

### GetMaskedNumber

`func (o *Card) GetMaskedNumber() string`

GetMaskedNumber returns the MaskedNumber field if non-nil, zero value otherwise.

### GetMaskedNumberOk

`func (o *Card) GetMaskedNumberOk() (*string, bool)`

GetMaskedNumberOk returns a tuple with the MaskedNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedNumber

`func (o *Card) SetMaskedNumber(v string)`

SetMaskedNumber sets MaskedNumber field to given value.

### HasMaskedNumber

`func (o *Card) HasMaskedNumber() bool`

HasMaskedNumber returns a boolean if a field has been set.

### GetExpiration

`func (o *Card) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *Card) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *Card) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *Card) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetType

`func (o *Card) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Card) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Card) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Card) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsRegistered

`func (o *Card) GetIsRegistered() bool`

GetIsRegistered returns the IsRegistered field if non-nil, zero value otherwise.

### GetIsRegisteredOk

`func (o *Card) GetIsRegisteredOk() (*bool, bool)`

GetIsRegisteredOk returns a tuple with the IsRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRegistered

`func (o *Card) SetIsRegistered(v bool)`

SetIsRegistered sets IsRegistered field to given value.

### HasIsRegistered

`func (o *Card) HasIsRegistered() bool`

HasIsRegistered returns a boolean if a field has been set.

### GetHolderName

`func (o *Card) GetHolderName() string`

GetHolderName returns the HolderName field if non-nil, zero value otherwise.

### GetHolderNameOk

`func (o *Card) GetHolderNameOk() (*string, bool)`

GetHolderNameOk returns a tuple with the HolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderName

`func (o *Card) SetHolderName(v string)`

SetHolderName sets HolderName field to given value.

### HasHolderName

`func (o *Card) HasHolderName() bool`

HasHolderName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


