# RegisterIBANInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Payment Account ID | 
**Holder** | **string** | Registered Bank Account Owner: First and Last name, or Company Name | 
**Bic** | Pointer to **string** | BIC/SWIFT Codes are arranged like this : AAAABBCCDDD  AAAA : 4 char for bank code  BB : 2 char for country code  CC : 2 char for location code  DDD : 3 char for branch code | [optional] 
**Iban** | **string** | IBAN | 
**Domiciliation1** | Pointer to **string** | Bank Address Line 1 | [optional] 
**Domiciliation2** | Pointer to **string** | Bank Address Line 2 | [optional] 
**Comment** | Pointer to **string** | Reason for new IBAN if another IBAN is already linked to the Payment Account | [optional] 

## Methods

### NewRegisterIBANInput

`func NewRegisterIBANInput(accountId string, holder string, iban string, ) *RegisterIBANInput`

NewRegisterIBANInput instantiates a new RegisterIBANInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterIBANInputWithDefaults

`func NewRegisterIBANInputWithDefaults() *RegisterIBANInput`

NewRegisterIBANInputWithDefaults instantiates a new RegisterIBANInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *RegisterIBANInput) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RegisterIBANInput) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RegisterIBANInput) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetHolder

`func (o *RegisterIBANInput) GetHolder() string`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *RegisterIBANInput) GetHolderOk() (*string, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *RegisterIBANInput) SetHolder(v string)`

SetHolder sets Holder field to given value.


### GetBic

`func (o *RegisterIBANInput) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *RegisterIBANInput) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *RegisterIBANInput) SetBic(v string)`

SetBic sets Bic field to given value.

### HasBic

`func (o *RegisterIBANInput) HasBic() bool`

HasBic returns a boolean if a field has been set.

### GetIban

`func (o *RegisterIBANInput) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *RegisterIBANInput) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *RegisterIBANInput) SetIban(v string)`

SetIban sets Iban field to given value.


### GetDomiciliation1

`func (o *RegisterIBANInput) GetDomiciliation1() string`

GetDomiciliation1 returns the Domiciliation1 field if non-nil, zero value otherwise.

### GetDomiciliation1Ok

`func (o *RegisterIBANInput) GetDomiciliation1Ok() (*string, bool)`

GetDomiciliation1Ok returns a tuple with the Domiciliation1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomiciliation1

`func (o *RegisterIBANInput) SetDomiciliation1(v string)`

SetDomiciliation1 sets Domiciliation1 field to given value.

### HasDomiciliation1

`func (o *RegisterIBANInput) HasDomiciliation1() bool`

HasDomiciliation1 returns a boolean if a field has been set.

### GetDomiciliation2

`func (o *RegisterIBANInput) GetDomiciliation2() string`

GetDomiciliation2 returns the Domiciliation2 field if non-nil, zero value otherwise.

### GetDomiciliation2Ok

`func (o *RegisterIBANInput) GetDomiciliation2Ok() (*string, bool)`

GetDomiciliation2Ok returns a tuple with the Domiciliation2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomiciliation2

`func (o *RegisterIBANInput) SetDomiciliation2(v string)`

SetDomiciliation2 sets Domiciliation2 field to given value.

### HasDomiciliation2

`func (o *RegisterIBANInput) HasDomiciliation2() bool`

HasDomiciliation2 returns a boolean if a field has been set.

### GetComment

`func (o *RegisterIBANInput) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RegisterIBANInput) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RegisterIBANInput) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RegisterIBANInput) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


