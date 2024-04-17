# PurchaseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantAccountId** | **string** | Merchant account id | 
**Description** | **string** | Description | 
**Quantity** | **int32** | Quantity | 
**UnitAmount** | **int32** | Unit amount | 
**Type** | **int32** | Type&lt;br/&gt;1 &#x3D; Physical.&lt;br/&gt;2 &#x3D; Digital.&lt;br/&gt; | 
**TaxAmount** | Pointer to **int32** | Tax amount | [optional] 

## Methods

### NewPurchaseItem

`func NewPurchaseItem(merchantAccountId string, description string, quantity int32, unitAmount int32, type_ int32, ) *PurchaseItem`

NewPurchaseItem instantiates a new PurchaseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseItemWithDefaults

`func NewPurchaseItemWithDefaults() *PurchaseItem`

NewPurchaseItemWithDefaults instantiates a new PurchaseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantAccountId

`func (o *PurchaseItem) GetMerchantAccountId() string`

GetMerchantAccountId returns the MerchantAccountId field if non-nil, zero value otherwise.

### GetMerchantAccountIdOk

`func (o *PurchaseItem) GetMerchantAccountIdOk() (*string, bool)`

GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountId

`func (o *PurchaseItem) SetMerchantAccountId(v string)`

SetMerchantAccountId sets MerchantAccountId field to given value.


### GetDescription

`func (o *PurchaseItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PurchaseItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PurchaseItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetQuantity

`func (o *PurchaseItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PurchaseItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PurchaseItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetUnitAmount

`func (o *PurchaseItem) GetUnitAmount() int32`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *PurchaseItem) GetUnitAmountOk() (*int32, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *PurchaseItem) SetUnitAmount(v int32)`

SetUnitAmount sets UnitAmount field to given value.


### GetType

`func (o *PurchaseItem) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PurchaseItem) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PurchaseItem) SetType(v int32)`

SetType sets Type field to given value.


### GetTaxAmount

`func (o *PurchaseItem) GetTaxAmount() int32`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *PurchaseItem) GetTaxAmountOk() (*int32, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *PurchaseItem) SetTaxAmount(v int32)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *PurchaseItem) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


