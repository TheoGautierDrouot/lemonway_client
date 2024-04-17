# MoneyInIDealInitOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Transaction ID. You will this this value to confirm the transaction | [optional] 
**ActionUrl** | Pointer to **string** | Redirected URL for the client on the iDeal page payment | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewMoneyInIDealInitOutput

`func NewMoneyInIDealInitOutput() *MoneyInIDealInitOutput`

NewMoneyInIDealInitOutput instantiates a new MoneyInIDealInitOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoneyInIDealInitOutputWithDefaults

`func NewMoneyInIDealInitOutputWithDefaults() *MoneyInIDealInitOutput`

NewMoneyInIDealInitOutputWithDefaults instantiates a new MoneyInIDealInitOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MoneyInIDealInitOutput) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MoneyInIDealInitOutput) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MoneyInIDealInitOutput) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MoneyInIDealInitOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActionUrl

`func (o *MoneyInIDealInitOutput) GetActionUrl() string`

GetActionUrl returns the ActionUrl field if non-nil, zero value otherwise.

### GetActionUrlOk

`func (o *MoneyInIDealInitOutput) GetActionUrlOk() (*string, bool)`

GetActionUrlOk returns a tuple with the ActionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionUrl

`func (o *MoneyInIDealInitOutput) SetActionUrl(v string)`

SetActionUrl sets ActionUrl field to given value.

### HasActionUrl

`func (o *MoneyInIDealInitOutput) HasActionUrl() bool`

HasActionUrl returns a boolean if a field has been set.

### GetError

`func (o *MoneyInIDealInitOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MoneyInIDealInitOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MoneyInIDealInitOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *MoneyInIDealInitOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


