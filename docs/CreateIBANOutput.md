# CreateIBANOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | IBAN ID | [optional] 
**Iban** | Pointer to **string** | IBAN | [optional] 
**Bic** | Pointer to **string** | BIC/SWIFT Code | [optional] 
**Holder** | Pointer to **string** | IBAN Owner: First Name, Last Name, or Enterprise Name | [optional] 
**Domiciliation** | Pointer to **string** | Domiciliation | [optional] 
**Status** | Pointer to **int32** | IBAN Status | [optional] 
**MaxAvailableIbanPerWallet** | Pointer to **int32** | IBAN left that can be generated for this wallet | [optional] 
**MaxAvailableIban** | Pointer to **int32** | IBAN left that can be generated in total | [optional] 
**Pdf** | Pointer to **string** | PDF in base64 | [optional] 
**QrCode** | Pointer to **string** | QR Code PNG in base64 | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewCreateIBANOutput

`func NewCreateIBANOutput() *CreateIBANOutput`

NewCreateIBANOutput instantiates a new CreateIBANOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIBANOutputWithDefaults

`func NewCreateIBANOutputWithDefaults() *CreateIBANOutput`

NewCreateIBANOutputWithDefaults instantiates a new CreateIBANOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateIBANOutput) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateIBANOutput) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateIBANOutput) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CreateIBANOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIban

`func (o *CreateIBANOutput) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *CreateIBANOutput) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *CreateIBANOutput) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *CreateIBANOutput) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetBic

`func (o *CreateIBANOutput) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *CreateIBANOutput) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *CreateIBANOutput) SetBic(v string)`

SetBic sets Bic field to given value.

### HasBic

`func (o *CreateIBANOutput) HasBic() bool`

HasBic returns a boolean if a field has been set.

### GetHolder

`func (o *CreateIBANOutput) GetHolder() string`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *CreateIBANOutput) GetHolderOk() (*string, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *CreateIBANOutput) SetHolder(v string)`

SetHolder sets Holder field to given value.

### HasHolder

`func (o *CreateIBANOutput) HasHolder() bool`

HasHolder returns a boolean if a field has been set.

### GetDomiciliation

`func (o *CreateIBANOutput) GetDomiciliation() string`

GetDomiciliation returns the Domiciliation field if non-nil, zero value otherwise.

### GetDomiciliationOk

`func (o *CreateIBANOutput) GetDomiciliationOk() (*string, bool)`

GetDomiciliationOk returns a tuple with the Domiciliation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomiciliation

`func (o *CreateIBANOutput) SetDomiciliation(v string)`

SetDomiciliation sets Domiciliation field to given value.

### HasDomiciliation

`func (o *CreateIBANOutput) HasDomiciliation() bool`

HasDomiciliation returns a boolean if a field has been set.

### GetStatus

`func (o *CreateIBANOutput) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateIBANOutput) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateIBANOutput) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateIBANOutput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMaxAvailableIbanPerWallet

`func (o *CreateIBANOutput) GetMaxAvailableIbanPerWallet() int32`

GetMaxAvailableIbanPerWallet returns the MaxAvailableIbanPerWallet field if non-nil, zero value otherwise.

### GetMaxAvailableIbanPerWalletOk

`func (o *CreateIBANOutput) GetMaxAvailableIbanPerWalletOk() (*int32, bool)`

GetMaxAvailableIbanPerWalletOk returns a tuple with the MaxAvailableIbanPerWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAvailableIbanPerWallet

`func (o *CreateIBANOutput) SetMaxAvailableIbanPerWallet(v int32)`

SetMaxAvailableIbanPerWallet sets MaxAvailableIbanPerWallet field to given value.

### HasMaxAvailableIbanPerWallet

`func (o *CreateIBANOutput) HasMaxAvailableIbanPerWallet() bool`

HasMaxAvailableIbanPerWallet returns a boolean if a field has been set.

### GetMaxAvailableIban

`func (o *CreateIBANOutput) GetMaxAvailableIban() int32`

GetMaxAvailableIban returns the MaxAvailableIban field if non-nil, zero value otherwise.

### GetMaxAvailableIbanOk

`func (o *CreateIBANOutput) GetMaxAvailableIbanOk() (*int32, bool)`

GetMaxAvailableIbanOk returns a tuple with the MaxAvailableIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAvailableIban

`func (o *CreateIBANOutput) SetMaxAvailableIban(v int32)`

SetMaxAvailableIban sets MaxAvailableIban field to given value.

### HasMaxAvailableIban

`func (o *CreateIBANOutput) HasMaxAvailableIban() bool`

HasMaxAvailableIban returns a boolean if a field has been set.

### GetPdf

`func (o *CreateIBANOutput) GetPdf() string`

GetPdf returns the Pdf field if non-nil, zero value otherwise.

### GetPdfOk

`func (o *CreateIBANOutput) GetPdfOk() (*string, bool)`

GetPdfOk returns a tuple with the Pdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdf

`func (o *CreateIBANOutput) SetPdf(v string)`

SetPdf sets Pdf field to given value.

### HasPdf

`func (o *CreateIBANOutput) HasPdf() bool`

HasPdf returns a boolean if a field has been set.

### GetQrCode

`func (o *CreateIBANOutput) GetQrCode() string`

GetQrCode returns the QrCode field if non-nil, zero value otherwise.

### GetQrCodeOk

`func (o *CreateIBANOutput) GetQrCodeOk() (*string, bool)`

GetQrCodeOk returns a tuple with the QrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCode

`func (o *CreateIBANOutput) SetQrCode(v string)`

SetQrCode sets QrCode field to given value.

### HasQrCode

`func (o *CreateIBANOutput) HasQrCode() bool`

HasQrCode returns a boolean if a field has been set.

### GetError

`func (o *CreateIBANOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CreateIBANOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CreateIBANOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *CreateIBANOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


