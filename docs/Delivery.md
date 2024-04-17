# Delivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Receiver** | [**Receiver**](Receiver.md) |  | 
**Address** | [**PaypalDeliveryAddress**](PaypalDeliveryAddress.md) |  | 

## Methods

### NewDelivery

`func NewDelivery(receiver Receiver, address PaypalDeliveryAddress, ) *Delivery`

NewDelivery instantiates a new Delivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryWithDefaults

`func NewDeliveryWithDefaults() *Delivery`

NewDeliveryWithDefaults instantiates a new Delivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceiver

`func (o *Delivery) GetReceiver() Receiver`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *Delivery) GetReceiverOk() (*Receiver, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *Delivery) SetReceiver(v Receiver)`

SetReceiver sets Receiver field to given value.


### GetAddress

`func (o *Delivery) GetAddress() PaypalDeliveryAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Delivery) GetAddressOk() (*PaypalDeliveryAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Delivery) SetAddress(v PaypalDeliveryAddress)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


