# CreateIBANInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wallet** | **string** | Payment Account ID. | 
**Country** | **string** | ISO2 code of the country from which the IBAN must be generated  We currently serve 4 countries:  &lt;ul&gt;&lt;li&gt;BE: Belgium (BNP Fortis)&lt;/li&gt;&lt;li&gt;ES: BNP Paribas Sucursal en España&lt;/li&gt;&lt;li&gt;FR: France (BNP Paribas)&lt;/li&gt;&lt;li&gt;IT: Italy (BNL)&lt;/li&gt;&lt;/ul&gt; | 
**GeneratePDFAndQrCode** | Pointer to **bool** | If activated, this function will return a PDF document describing the virtual IBAN, along with a &lt;a href&#x3D;\&quot;https://www.europeanpaymentscouncil.eu/document-library/guidance-documents/quick-response-code-guidelines-enable-data-capture-initiation\&quot; target&#x3D;\&quot;_blank\&quot;&gt;standard EPC QR-Code&lt;/a&gt; image. Please store these documents in your datawarehouse, as they are not accessible through Lemonway&#39;s API after the virtual IBAN has been generated. | [optional] 
**PdfLanguage** | Pointer to **string** | ISO 639-1 language code for the PDF document. The supported languages are:  &lt;ul&gt;&lt;li&gt;en: English&lt;/li&gt;&lt;li&gt;es: Spanish&lt;/li&gt;&lt;li&gt;fr: French&lt;/li&gt;&lt;li&gt;it: Italian&lt;/li&gt;&lt;/ul&gt;  This field is ignored if generatePDFAndQrCode is false. | [optional] 

## Methods

### NewCreateIBANInput

`func NewCreateIBANInput(wallet string, country string, ) *CreateIBANInput`

NewCreateIBANInput instantiates a new CreateIBANInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIBANInputWithDefaults

`func NewCreateIBANInputWithDefaults() *CreateIBANInput`

NewCreateIBANInputWithDefaults instantiates a new CreateIBANInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWallet

`func (o *CreateIBANInput) GetWallet() string`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *CreateIBANInput) GetWalletOk() (*string, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *CreateIBANInput) SetWallet(v string)`

SetWallet sets Wallet field to given value.


### GetCountry

`func (o *CreateIBANInput) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CreateIBANInput) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CreateIBANInput) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetGeneratePDFAndQrCode

`func (o *CreateIBANInput) GetGeneratePDFAndQrCode() bool`

GetGeneratePDFAndQrCode returns the GeneratePDFAndQrCode field if non-nil, zero value otherwise.

### GetGeneratePDFAndQrCodeOk

`func (o *CreateIBANInput) GetGeneratePDFAndQrCodeOk() (*bool, bool)`

GetGeneratePDFAndQrCodeOk returns a tuple with the GeneratePDFAndQrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratePDFAndQrCode

`func (o *CreateIBANInput) SetGeneratePDFAndQrCode(v bool)`

SetGeneratePDFAndQrCode sets GeneratePDFAndQrCode field to given value.

### HasGeneratePDFAndQrCode

`func (o *CreateIBANInput) HasGeneratePDFAndQrCode() bool`

HasGeneratePDFAndQrCode returns a boolean if a field has been set.

### GetPdfLanguage

`func (o *CreateIBANInput) GetPdfLanguage() string`

GetPdfLanguage returns the PdfLanguage field if non-nil, zero value otherwise.

### GetPdfLanguageOk

`func (o *CreateIBANInput) GetPdfLanguageOk() (*string, bool)`

GetPdfLanguageOk returns a tuple with the PdfLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfLanguage

`func (o *CreateIBANInput) SetPdfLanguage(v string)`

SetPdfLanguage sets PdfLanguage field to given value.

### HasPdfLanguage

`func (o *CreateIBANInput) HasPdfLanguage() bool`

HasPdfLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


