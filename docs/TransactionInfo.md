# TransactionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reference** | **string** | Transaction Reference Sent to the PSP | 
**Order** | **string** | Merchant Order ID | 
**DateTime** | **string** | Transaction DateTime (UTC Unix timestamp) | 
**MerchantId** | **string** | Merchant ID Sent to PSP | 
**AuthorisationId** | **string** | Transaction Authorisation ID Received from PSP | 

## Methods

### NewTransactionInfo

`func NewTransactionInfo(reference string, order string, dateTime string, merchantId string, authorisationId string, ) *TransactionInfo`

NewTransactionInfo instantiates a new TransactionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionInfoWithDefaults

`func NewTransactionInfoWithDefaults() *TransactionInfo`

NewTransactionInfoWithDefaults instantiates a new TransactionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReference

`func (o *TransactionInfo) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *TransactionInfo) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *TransactionInfo) SetReference(v string)`

SetReference sets Reference field to given value.


### GetOrder

`func (o *TransactionInfo) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TransactionInfo) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TransactionInfo) SetOrder(v string)`

SetOrder sets Order field to given value.


### GetDateTime

`func (o *TransactionInfo) GetDateTime() string`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *TransactionInfo) GetDateTimeOk() (*string, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *TransactionInfo) SetDateTime(v string)`

SetDateTime sets DateTime field to given value.


### GetMerchantId

`func (o *TransactionInfo) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *TransactionInfo) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *TransactionInfo) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetAuthorisationId

`func (o *TransactionInfo) GetAuthorisationId() string`

GetAuthorisationId returns the AuthorisationId field if non-nil, zero value otherwise.

### GetAuthorisationIdOk

`func (o *TransactionInfo) GetAuthorisationIdOk() (*string, bool)`

GetAuthorisationIdOk returns a tuple with the AuthorisationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorisationId

`func (o *TransactionInfo) SetAuthorisationId(v string)`

SetAuthorisationId sets AuthorisationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


