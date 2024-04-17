# BnplDelivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryMode** | **string** |  | 
**ExpressDelivery** | **bool** |  | 
**DeliveryAddress** | [**BnplAddress**](BnplAddress.md) |  | 
**DeliveryPointName** | Pointer to **string** | Pickup name or recipient name.  Mandatory only if DeliveryMode is not \&quot;DeliveryToCustomerAddress\&quot;. | [optional] 

## Methods

### NewBnplDelivery

`func NewBnplDelivery(deliveryMode string, expressDelivery bool, deliveryAddress BnplAddress, ) *BnplDelivery`

NewBnplDelivery instantiates a new BnplDelivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnplDeliveryWithDefaults

`func NewBnplDeliveryWithDefaults() *BnplDelivery`

NewBnplDeliveryWithDefaults instantiates a new BnplDelivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryMode

`func (o *BnplDelivery) GetDeliveryMode() string`

GetDeliveryMode returns the DeliveryMode field if non-nil, zero value otherwise.

### GetDeliveryModeOk

`func (o *BnplDelivery) GetDeliveryModeOk() (*string, bool)`

GetDeliveryModeOk returns a tuple with the DeliveryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMode

`func (o *BnplDelivery) SetDeliveryMode(v string)`

SetDeliveryMode sets DeliveryMode field to given value.


### GetExpressDelivery

`func (o *BnplDelivery) GetExpressDelivery() bool`

GetExpressDelivery returns the ExpressDelivery field if non-nil, zero value otherwise.

### GetExpressDeliveryOk

`func (o *BnplDelivery) GetExpressDeliveryOk() (*bool, bool)`

GetExpressDeliveryOk returns a tuple with the ExpressDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressDelivery

`func (o *BnplDelivery) SetExpressDelivery(v bool)`

SetExpressDelivery sets ExpressDelivery field to given value.


### GetDeliveryAddress

`func (o *BnplDelivery) GetDeliveryAddress() BnplAddress`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *BnplDelivery) GetDeliveryAddressOk() (*BnplAddress, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *BnplDelivery) SetDeliveryAddress(v BnplAddress)`

SetDeliveryAddress sets DeliveryAddress field to given value.


### GetDeliveryPointName

`func (o *BnplDelivery) GetDeliveryPointName() string`

GetDeliveryPointName returns the DeliveryPointName field if non-nil, zero value otherwise.

### GetDeliveryPointNameOk

`func (o *BnplDelivery) GetDeliveryPointNameOk() (*string, bool)`

GetDeliveryPointNameOk returns a tuple with the DeliveryPointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryPointName

`func (o *BnplDelivery) SetDeliveryPointName(v string)`

SetDeliveryPointName sets DeliveryPointName field to given value.

### HasDeliveryPointName

`func (o *BnplDelivery) HasDeliveryPointName() bool`

HasDeliveryPointName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


