# GetMoneyInBanksOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Banks** | Pointer to [**map[string][]Bank**](array.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewGetMoneyInBanksOutput

`func NewGetMoneyInBanksOutput() *GetMoneyInBanksOutput`

NewGetMoneyInBanksOutput instantiates a new GetMoneyInBanksOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMoneyInBanksOutputWithDefaults

`func NewGetMoneyInBanksOutputWithDefaults() *GetMoneyInBanksOutput`

NewGetMoneyInBanksOutputWithDefaults instantiates a new GetMoneyInBanksOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBanks

`func (o *GetMoneyInBanksOutput) GetBanks() map[string][]Bank`

GetBanks returns the Banks field if non-nil, zero value otherwise.

### GetBanksOk

`func (o *GetMoneyInBanksOutput) GetBanksOk() (*map[string][]Bank, bool)`

GetBanksOk returns a tuple with the Banks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanks

`func (o *GetMoneyInBanksOutput) SetBanks(v map[string][]Bank)`

SetBanks sets Banks field to given value.

### HasBanks

`func (o *GetMoneyInBanksOutput) HasBanks() bool`

HasBanks returns a boolean if a field has been set.

### GetError

`func (o *GetMoneyInBanksOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetMoneyInBanksOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetMoneyInBanksOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *GetMoneyInBanksOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


