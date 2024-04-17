# InitPayPalTransactionOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Money-in Id | 
**RedirectionUrl** | **string** | Redirection Url | 
**PayPalOrderId** | **string** | PayPal Order Id | 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewInitPayPalTransactionOutput

`func NewInitPayPalTransactionOutput(id int64, redirectionUrl string, payPalOrderId string, ) *InitPayPalTransactionOutput`

NewInitPayPalTransactionOutput instantiates a new InitPayPalTransactionOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitPayPalTransactionOutputWithDefaults

`func NewInitPayPalTransactionOutputWithDefaults() *InitPayPalTransactionOutput`

NewInitPayPalTransactionOutputWithDefaults instantiates a new InitPayPalTransactionOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InitPayPalTransactionOutput) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InitPayPalTransactionOutput) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InitPayPalTransactionOutput) SetId(v int64)`

SetId sets Id field to given value.


### GetRedirectionUrl

`func (o *InitPayPalTransactionOutput) GetRedirectionUrl() string`

GetRedirectionUrl returns the RedirectionUrl field if non-nil, zero value otherwise.

### GetRedirectionUrlOk

`func (o *InitPayPalTransactionOutput) GetRedirectionUrlOk() (*string, bool)`

GetRedirectionUrlOk returns a tuple with the RedirectionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectionUrl

`func (o *InitPayPalTransactionOutput) SetRedirectionUrl(v string)`

SetRedirectionUrl sets RedirectionUrl field to given value.


### GetPayPalOrderId

`func (o *InitPayPalTransactionOutput) GetPayPalOrderId() string`

GetPayPalOrderId returns the PayPalOrderId field if non-nil, zero value otherwise.

### GetPayPalOrderIdOk

`func (o *InitPayPalTransactionOutput) GetPayPalOrderIdOk() (*string, bool)`

GetPayPalOrderIdOk returns a tuple with the PayPalOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPalOrderId

`func (o *InitPayPalTransactionOutput) SetPayPalOrderId(v string)`

SetPayPalOrderId sets PayPalOrderId field to given value.


### GetError

`func (o *InitPayPalTransactionOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InitPayPalTransactionOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InitPayPalTransactionOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *InitPayPalTransactionOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


