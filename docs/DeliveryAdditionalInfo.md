# DeliveryAdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**NumberOfItemsBasket** | Pointer to **string** | Total quantity of all products in the basket | [optional] 
**AddressDeliveryBillingMatchIndicator** | Pointer to **string** | Specifies whether the delivery and the billing addresses are the same | [optional] 
**DeliveryAddressCreationDate** | Pointer to **string** | The date on which the last delivery address used by the merchant&#39;s account was reported in the transaction | [optional] 
**EstimatedDeliveryDelay** | Pointer to **string** | Estimated Delivery Delay (in days) by the Merchant | [optional] 
**DeliveryMode** | Pointer to **string** | Delivery Method (Postal Office or Amazon box and so on) | [optional] 

## Methods

### NewDeliveryAdditionalInfo

`func NewDeliveryAdditionalInfo() *DeliveryAdditionalInfo`

NewDeliveryAdditionalInfo instantiates a new DeliveryAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryAdditionalInfoWithDefaults

`func NewDeliveryAdditionalInfoWithDefaults() *DeliveryAdditionalInfo`

NewDeliveryAdditionalInfoWithDefaults instantiates a new DeliveryAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *DeliveryAdditionalInfo) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *DeliveryAdditionalInfo) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *DeliveryAdditionalInfo) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *DeliveryAdditionalInfo) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetNumberOfItemsBasket

`func (o *DeliveryAdditionalInfo) GetNumberOfItemsBasket() string`

GetNumberOfItemsBasket returns the NumberOfItemsBasket field if non-nil, zero value otherwise.

### GetNumberOfItemsBasketOk

`func (o *DeliveryAdditionalInfo) GetNumberOfItemsBasketOk() (*string, bool)`

GetNumberOfItemsBasketOk returns a tuple with the NumberOfItemsBasket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfItemsBasket

`func (o *DeliveryAdditionalInfo) SetNumberOfItemsBasket(v string)`

SetNumberOfItemsBasket sets NumberOfItemsBasket field to given value.

### HasNumberOfItemsBasket

`func (o *DeliveryAdditionalInfo) HasNumberOfItemsBasket() bool`

HasNumberOfItemsBasket returns a boolean if a field has been set.

### GetAddressDeliveryBillingMatchIndicator

`func (o *DeliveryAdditionalInfo) GetAddressDeliveryBillingMatchIndicator() string`

GetAddressDeliveryBillingMatchIndicator returns the AddressDeliveryBillingMatchIndicator field if non-nil, zero value otherwise.

### GetAddressDeliveryBillingMatchIndicatorOk

`func (o *DeliveryAdditionalInfo) GetAddressDeliveryBillingMatchIndicatorOk() (*string, bool)`

GetAddressDeliveryBillingMatchIndicatorOk returns a tuple with the AddressDeliveryBillingMatchIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressDeliveryBillingMatchIndicator

`func (o *DeliveryAdditionalInfo) SetAddressDeliveryBillingMatchIndicator(v string)`

SetAddressDeliveryBillingMatchIndicator sets AddressDeliveryBillingMatchIndicator field to given value.

### HasAddressDeliveryBillingMatchIndicator

`func (o *DeliveryAdditionalInfo) HasAddressDeliveryBillingMatchIndicator() bool`

HasAddressDeliveryBillingMatchIndicator returns a boolean if a field has been set.

### GetDeliveryAddressCreationDate

`func (o *DeliveryAdditionalInfo) GetDeliveryAddressCreationDate() string`

GetDeliveryAddressCreationDate returns the DeliveryAddressCreationDate field if non-nil, zero value otherwise.

### GetDeliveryAddressCreationDateOk

`func (o *DeliveryAdditionalInfo) GetDeliveryAddressCreationDateOk() (*string, bool)`

GetDeliveryAddressCreationDateOk returns a tuple with the DeliveryAddressCreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddressCreationDate

`func (o *DeliveryAdditionalInfo) SetDeliveryAddressCreationDate(v string)`

SetDeliveryAddressCreationDate sets DeliveryAddressCreationDate field to given value.

### HasDeliveryAddressCreationDate

`func (o *DeliveryAdditionalInfo) HasDeliveryAddressCreationDate() bool`

HasDeliveryAddressCreationDate returns a boolean if a field has been set.

### GetEstimatedDeliveryDelay

`func (o *DeliveryAdditionalInfo) GetEstimatedDeliveryDelay() string`

GetEstimatedDeliveryDelay returns the EstimatedDeliveryDelay field if non-nil, zero value otherwise.

### GetEstimatedDeliveryDelayOk

`func (o *DeliveryAdditionalInfo) GetEstimatedDeliveryDelayOk() (*string, bool)`

GetEstimatedDeliveryDelayOk returns a tuple with the EstimatedDeliveryDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDeliveryDelay

`func (o *DeliveryAdditionalInfo) SetEstimatedDeliveryDelay(v string)`

SetEstimatedDeliveryDelay sets EstimatedDeliveryDelay field to given value.

### HasEstimatedDeliveryDelay

`func (o *DeliveryAdditionalInfo) HasEstimatedDeliveryDelay() bool`

HasEstimatedDeliveryDelay returns a boolean if a field has been set.

### GetDeliveryMode

`func (o *DeliveryAdditionalInfo) GetDeliveryMode() string`

GetDeliveryMode returns the DeliveryMode field if non-nil, zero value otherwise.

### GetDeliveryModeOk

`func (o *DeliveryAdditionalInfo) GetDeliveryModeOk() (*string, bool)`

GetDeliveryModeOk returns a tuple with the DeliveryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMode

`func (o *DeliveryAdditionalInfo) SetDeliveryMode(v string)`

SetDeliveryMode sets DeliveryMode field to given value.

### HasDeliveryMode

`func (o *DeliveryAdditionalInfo) HasDeliveryMode() bool`

HasDeliveryMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


