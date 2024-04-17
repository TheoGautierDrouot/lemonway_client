# KycStatusOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]AccountKycStatus**](AccountKycStatus.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewKycStatusOutput

`func NewKycStatusOutput() *KycStatusOutput`

NewKycStatusOutput instantiates a new KycStatusOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKycStatusOutputWithDefaults

`func NewKycStatusOutputWithDefaults() *KycStatusOutput`

NewKycStatusOutputWithDefaults instantiates a new KycStatusOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *KycStatusOutput) GetAccounts() []AccountKycStatus`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *KycStatusOutput) GetAccountsOk() (*[]AccountKycStatus, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *KycStatusOutput) SetAccounts(v []AccountKycStatus)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *KycStatusOutput) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetLinks

`func (o *KycStatusOutput) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *KycStatusOutput) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *KycStatusOutput) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *KycStatusOutput) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetError

`func (o *KycStatusOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *KycStatusOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *KycStatusOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *KycStatusOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


