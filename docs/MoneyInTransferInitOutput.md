# MoneyInTransferInitOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedirectUrl** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewMoneyInTransferInitOutput

`func NewMoneyInTransferInitOutput() *MoneyInTransferInitOutput`

NewMoneyInTransferInitOutput instantiates a new MoneyInTransferInitOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoneyInTransferInitOutputWithDefaults

`func NewMoneyInTransferInitOutputWithDefaults() *MoneyInTransferInitOutput`

NewMoneyInTransferInitOutputWithDefaults instantiates a new MoneyInTransferInitOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirectUrl

`func (o *MoneyInTransferInitOutput) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *MoneyInTransferInitOutput) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *MoneyInTransferInitOutput) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *MoneyInTransferInitOutput) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetId

`func (o *MoneyInTransferInitOutput) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MoneyInTransferInitOutput) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MoneyInTransferInitOutput) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MoneyInTransferInitOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetError

`func (o *MoneyInTransferInitOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MoneyInTransferInitOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MoneyInTransferInitOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *MoneyInTransferInitOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


