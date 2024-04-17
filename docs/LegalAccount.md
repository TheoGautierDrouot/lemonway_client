# LegalAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Updated Wallet ID (your choice when you create the Wallet) | [optional] 
**LemonwayId** | Pointer to **int64** | Lemonway Wallet ID (you can not choose this number). This number is also displayed in our Dashboard. | [optional] 
**Limits** | Pointer to [**Limits**](Limits.md) |  | [optional] 

## Methods

### NewLegalAccount

`func NewLegalAccount() *LegalAccount`

NewLegalAccount instantiates a new LegalAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalAccountWithDefaults

`func NewLegalAccountWithDefaults() *LegalAccount`

NewLegalAccountWithDefaults instantiates a new LegalAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LegalAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LegalAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LegalAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LegalAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLemonwayId

`func (o *LegalAccount) GetLemonwayId() int64`

GetLemonwayId returns the LemonwayId field if non-nil, zero value otherwise.

### GetLemonwayIdOk

`func (o *LegalAccount) GetLemonwayIdOk() (*int64, bool)`

GetLemonwayIdOk returns a tuple with the LemonwayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLemonwayId

`func (o *LegalAccount) SetLemonwayId(v int64)`

SetLemonwayId sets LemonwayId field to given value.

### HasLemonwayId

`func (o *LegalAccount) HasLemonwayId() bool`

HasLemonwayId returns a boolean if a field has been set.

### GetLimits

`func (o *LegalAccount) GetLimits() Limits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *LegalAccount) GetLimitsOk() (*Limits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *LegalAccount) SetLimits(v Limits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *LegalAccount) HasLimits() bool`

HasLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


