# LemonWayCommission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Idp2p** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **int32** | Amounts are represented as integer in cents (Euro)  Represented as an integer in cents (Euro) | [optional] 

## Methods

### NewLemonWayCommission

`func NewLemonWayCommission() *LemonWayCommission`

NewLemonWayCommission instantiates a new LemonWayCommission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLemonWayCommissionWithDefaults

`func NewLemonWayCommissionWithDefaults() *LemonWayCommission`

NewLemonWayCommissionWithDefaults instantiates a new LemonWayCommission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdp2p

`func (o *LemonWayCommission) GetIdp2p() string`

GetIdp2p returns the Idp2p field if non-nil, zero value otherwise.

### GetIdp2pOk

`func (o *LemonWayCommission) GetIdp2pOk() (*string, bool)`

GetIdp2pOk returns a tuple with the Idp2p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp2p

`func (o *LemonWayCommission) SetIdp2p(v string)`

SetIdp2p sets Idp2p field to given value.

### HasIdp2p

`func (o *LemonWayCommission) HasIdp2p() bool`

HasIdp2p returns a boolean if a field has been set.

### GetAmount

`func (o *LemonWayCommission) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *LemonWayCommission) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *LemonWayCommission) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *LemonWayCommission) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


