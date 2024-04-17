# PaypalDeliveryAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | **string** | ISO 3166-1 alpha-2 | 
**City** | **string** | City | 
**Street** | **string** | Number and street | 
**PostCode** | **string** | Postcode/ZIP | 
**State** | Pointer to **string** | Postcode/ZIP | [optional] 

## Methods

### NewPaypalDeliveryAddress

`func NewPaypalDeliveryAddress(country string, city string, street string, postCode string, ) *PaypalDeliveryAddress`

NewPaypalDeliveryAddress instantiates a new PaypalDeliveryAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaypalDeliveryAddressWithDefaults

`func NewPaypalDeliveryAddressWithDefaults() *PaypalDeliveryAddress`

NewPaypalDeliveryAddressWithDefaults instantiates a new PaypalDeliveryAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *PaypalDeliveryAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PaypalDeliveryAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PaypalDeliveryAddress) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCity

`func (o *PaypalDeliveryAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PaypalDeliveryAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PaypalDeliveryAddress) SetCity(v string)`

SetCity sets City field to given value.


### GetStreet

`func (o *PaypalDeliveryAddress) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *PaypalDeliveryAddress) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *PaypalDeliveryAddress) SetStreet(v string)`

SetStreet sets Street field to given value.


### GetPostCode

`func (o *PaypalDeliveryAddress) GetPostCode() string`

GetPostCode returns the PostCode field if non-nil, zero value otherwise.

### GetPostCodeOk

`func (o *PaypalDeliveryAddress) GetPostCodeOk() (*string, bool)`

GetPostCodeOk returns a tuple with the PostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCode

`func (o *PaypalDeliveryAddress) SetPostCode(v string)`

SetPostCode sets PostCode field to given value.


### GetState

`func (o *PaypalDeliveryAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PaypalDeliveryAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PaypalDeliveryAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PaypalDeliveryAddress) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


