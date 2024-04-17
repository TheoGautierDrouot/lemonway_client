# InitPayPalTransactionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Redirections** | [**Redirections**](Redirections.md) |  | 
**Transaction** | [**Transaction**](Transaction.md) |  | 
**AmountBreakdown** | Pointer to [**AmountBreakdown**](AmountBreakdown.md) |  | [optional] 
**Delivery** | Pointer to [**Delivery**](Delivery.md) |  | [optional] 
**Items** | [**[]PurchaseItem**](PurchaseItem.md) | Items | 
**RiskData** | Pointer to **map[string]string** | Risk data | [optional] 

## Methods

### NewInitPayPalTransactionInput

`func NewInitPayPalTransactionInput(redirections Redirections, transaction Transaction, items []PurchaseItem, ) *InitPayPalTransactionInput`

NewInitPayPalTransactionInput instantiates a new InitPayPalTransactionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitPayPalTransactionInputWithDefaults

`func NewInitPayPalTransactionInputWithDefaults() *InitPayPalTransactionInput`

NewInitPayPalTransactionInputWithDefaults instantiates a new InitPayPalTransactionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirections

`func (o *InitPayPalTransactionInput) GetRedirections() Redirections`

GetRedirections returns the Redirections field if non-nil, zero value otherwise.

### GetRedirectionsOk

`func (o *InitPayPalTransactionInput) GetRedirectionsOk() (*Redirections, bool)`

GetRedirectionsOk returns a tuple with the Redirections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirections

`func (o *InitPayPalTransactionInput) SetRedirections(v Redirections)`

SetRedirections sets Redirections field to given value.


### GetTransaction

`func (o *InitPayPalTransactionInput) GetTransaction() Transaction`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *InitPayPalTransactionInput) GetTransactionOk() (*Transaction, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *InitPayPalTransactionInput) SetTransaction(v Transaction)`

SetTransaction sets Transaction field to given value.


### GetAmountBreakdown

`func (o *InitPayPalTransactionInput) GetAmountBreakdown() AmountBreakdown`

GetAmountBreakdown returns the AmountBreakdown field if non-nil, zero value otherwise.

### GetAmountBreakdownOk

`func (o *InitPayPalTransactionInput) GetAmountBreakdownOk() (*AmountBreakdown, bool)`

GetAmountBreakdownOk returns a tuple with the AmountBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountBreakdown

`func (o *InitPayPalTransactionInput) SetAmountBreakdown(v AmountBreakdown)`

SetAmountBreakdown sets AmountBreakdown field to given value.

### HasAmountBreakdown

`func (o *InitPayPalTransactionInput) HasAmountBreakdown() bool`

HasAmountBreakdown returns a boolean if a field has been set.

### GetDelivery

`func (o *InitPayPalTransactionInput) GetDelivery() Delivery`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *InitPayPalTransactionInput) GetDeliveryOk() (*Delivery, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *InitPayPalTransactionInput) SetDelivery(v Delivery)`

SetDelivery sets Delivery field to given value.

### HasDelivery

`func (o *InitPayPalTransactionInput) HasDelivery() bool`

HasDelivery returns a boolean if a field has been set.

### GetItems

`func (o *InitPayPalTransactionInput) GetItems() []PurchaseItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InitPayPalTransactionInput) GetItemsOk() (*[]PurchaseItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InitPayPalTransactionInput) SetItems(v []PurchaseItem)`

SetItems sets Items field to given value.


### GetRiskData

`func (o *InitPayPalTransactionInput) GetRiskData() map[string]string`

GetRiskData returns the RiskData field if non-nil, zero value otherwise.

### GetRiskDataOk

`func (o *InitPayPalTransactionInput) GetRiskDataOk() (*map[string]string, bool)`

GetRiskDataOk returns a tuple with the RiskData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskData

`func (o *InitPayPalTransactionInput) SetRiskData(v map[string]string)`

SetRiskData sets RiskData field to given value.

### HasRiskData

`func (o *InitPayPalTransactionInput) HasRiskData() bool`

HasRiskData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


