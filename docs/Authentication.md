# Authentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantCustomerAuthentMethod** | Pointer to **string** | Possible Authentication Methods: &lt;br /&gt;  - NOAUTHENT &#x3D; 1 - No authentication of the customer by the merchant                                                       - OWNCREDENTIAL &#x3D; 2 - Customer authentication by the merchant using his own system                                            - FEDERATEDID &#x3D; 3 - Customer authentication by the merchant using an identifier federated(facebook, ...) (e.g.Facebook)    - ISSUERID &#x3D; 4 - Customer authentication by the merchant using information of the issuer&#39;s payment mean                  - THIRDPARTY &#x3D; 5 - Customer authentication by the merchant using a third system                                            - FIDO &#x3D; 6 - Customer authentication by the merchant with FIDO(Fast IDentity Online) system | [optional] 
**MerchantCustomerAuthentDateTime** | Pointer to **string** | ISO8601 date time format | [optional] 

## Methods

### NewAuthentication

`func NewAuthentication() *Authentication`

NewAuthentication instantiates a new Authentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationWithDefaults

`func NewAuthenticationWithDefaults() *Authentication`

NewAuthenticationWithDefaults instantiates a new Authentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantCustomerAuthentMethod

`func (o *Authentication) GetMerchantCustomerAuthentMethod() string`

GetMerchantCustomerAuthentMethod returns the MerchantCustomerAuthentMethod field if non-nil, zero value otherwise.

### GetMerchantCustomerAuthentMethodOk

`func (o *Authentication) GetMerchantCustomerAuthentMethodOk() (*string, bool)`

GetMerchantCustomerAuthentMethodOk returns a tuple with the MerchantCustomerAuthentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCustomerAuthentMethod

`func (o *Authentication) SetMerchantCustomerAuthentMethod(v string)`

SetMerchantCustomerAuthentMethod sets MerchantCustomerAuthentMethod field to given value.

### HasMerchantCustomerAuthentMethod

`func (o *Authentication) HasMerchantCustomerAuthentMethod() bool`

HasMerchantCustomerAuthentMethod returns a boolean if a field has been set.

### GetMerchantCustomerAuthentDateTime

`func (o *Authentication) GetMerchantCustomerAuthentDateTime() string`

GetMerchantCustomerAuthentDateTime returns the MerchantCustomerAuthentDateTime field if non-nil, zero value otherwise.

### GetMerchantCustomerAuthentDateTimeOk

`func (o *Authentication) GetMerchantCustomerAuthentDateTimeOk() (*string, bool)`

GetMerchantCustomerAuthentDateTimeOk returns a tuple with the MerchantCustomerAuthentDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCustomerAuthentDateTime

`func (o *Authentication) SetMerchantCustomerAuthentDateTime(v string)`

SetMerchantCustomerAuthentDateTime sets MerchantCustomerAuthentDateTime field to given value.

### HasMerchantCustomerAuthentDateTime

`func (o *Authentication) HasMerchantCustomerAuthentDateTime() bool`

HasMerchantCustomerAuthentDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


