# IndividualAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Payment Account ID | [optional] 
**InternalId** | Pointer to **int64** | Internal ID given by Lemonway (ID displayed on the Lemonway Dashboard) | [optional] 
**Limits** | Pointer to [**Limits**](Limits.md) |  | [optional] 

## Methods

### NewIndividualAccount

`func NewIndividualAccount() *IndividualAccount`

NewIndividualAccount instantiates a new IndividualAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndividualAccountWithDefaults

`func NewIndividualAccountWithDefaults() *IndividualAccount`

NewIndividualAccountWithDefaults instantiates a new IndividualAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IndividualAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IndividualAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IndividualAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IndividualAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalId

`func (o *IndividualAccount) GetInternalId() int64`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *IndividualAccount) GetInternalIdOk() (*int64, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *IndividualAccount) SetInternalId(v int64)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *IndividualAccount) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetLimits

`func (o *IndividualAccount) GetLimits() Limits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *IndividualAccount) GetLimitsOk() (*Limits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *IndividualAccount) SetLimits(v Limits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *IndividualAccount) HasLimits() bool`

HasLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


