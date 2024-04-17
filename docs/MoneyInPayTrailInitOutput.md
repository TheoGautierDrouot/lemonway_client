# MoneyInPayTrailInitOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Transaction ID. You will this this value to confirm the transaction | [optional] 
**ActionUrl** | Pointer to **string** | Redirected URL for the client on the iDeal page payment | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewMoneyInPayTrailInitOutput

`func NewMoneyInPayTrailInitOutput() *MoneyInPayTrailInitOutput`

NewMoneyInPayTrailInitOutput instantiates a new MoneyInPayTrailInitOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoneyInPayTrailInitOutputWithDefaults

`func NewMoneyInPayTrailInitOutputWithDefaults() *MoneyInPayTrailInitOutput`

NewMoneyInPayTrailInitOutputWithDefaults instantiates a new MoneyInPayTrailInitOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MoneyInPayTrailInitOutput) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MoneyInPayTrailInitOutput) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MoneyInPayTrailInitOutput) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MoneyInPayTrailInitOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActionUrl

`func (o *MoneyInPayTrailInitOutput) GetActionUrl() string`

GetActionUrl returns the ActionUrl field if non-nil, zero value otherwise.

### GetActionUrlOk

`func (o *MoneyInPayTrailInitOutput) GetActionUrlOk() (*string, bool)`

GetActionUrlOk returns a tuple with the ActionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionUrl

`func (o *MoneyInPayTrailInitOutput) SetActionUrl(v string)`

SetActionUrl sets ActionUrl field to given value.

### HasActionUrl

`func (o *MoneyInPayTrailInitOutput) HasActionUrl() bool`

HasActionUrl returns a boolean if a field has been set.

### GetError

`func (o *MoneyInPayTrailInitOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MoneyInPayTrailInitOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MoneyInPayTrailInitOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *MoneyInPayTrailInitOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


