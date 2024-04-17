# PaymentForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletIp** | Pointer to **string** |  | [optional] 
**WalletUa** | Pointer to **string** |  | [optional] 
**AmountTotRange** | Pointer to **string** |  | [optional] 
**AmountCom** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Style** | Pointer to **string** |  | [optional] 
**AtosStyle** | Pointer to **string** | Mercanet v1 only: Link to a custom Mercanet CSS Stylesheet.  The stylesheet should be publicly accessible via HTTPS | [optional] 
**NotifUrl** | Pointer to **string** |  | [optional] 
**Options** | Pointer to **string** | Reserved for future version | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**RaisonSociale** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Payment Form ID | [optional] 
**OptId** | Pointer to **string** | Optional identity of the payment form. You should not use special character here. | [optional] 
**AccountPayer** | Pointer to **string** | Payer Account  If this field is filled then the money reaches this wallet before arriving at the beneficiary wallet (via a transfer wallet to wallet)  note: Don&#39;t put the &lt;b&gt;SC Wallet&lt;/b&gt; here, it won&#39;t work. You cannot credit the &lt;b&gt;SC Wallet&lt;/b&gt; with a credit card. | [optional] 
**AccountReceiver** | Pointer to **string** | Beneficiary Account  if this field is not filled then the end-user must to fill its value on the payment form. We recommend you to always fill it instead of your end-user.  note: Don&#39;t put the &lt;b&gt;SC Wallet&lt;/b&gt; here, it won&#39;t work. You cannot credit the &lt;b&gt;SC Wallet&lt;/b&gt; with a credit card. | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**ReturnUrl** | Pointer to **string** | URL redirection after the payment procedure is successfully finished | [optional] 
**CancelUrl** | Pointer to **string** | URL redirection after the payment procedure is cancelled | [optional] 
**ErrorUrl** | Pointer to **string** | URL redirection after the payment procedure is failed | [optional] 
**FirstNamePayer** | Pointer to **string** | Payer&#39;s First Name  If this field is not filled then the end-user have to fill it in the payment form. | [optional] 
**LastNamePayer** | Pointer to **string** | Payer&#39;s Last Name  If this field is not filled then the end-user have to fill it in the payment form. | [optional] 
**EmailPayer** | Pointer to **string** | Payer&#39;s Email  If this field is not filled then the end-user have to fill it in the payment form. | [optional] 

## Methods

### NewPaymentForm

`func NewPaymentForm() *PaymentForm`

NewPaymentForm instantiates a new PaymentForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentFormWithDefaults

`func NewPaymentFormWithDefaults() *PaymentForm`

NewPaymentFormWithDefaults instantiates a new PaymentForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletIp

`func (o *PaymentForm) GetWalletIp() string`

GetWalletIp returns the WalletIp field if non-nil, zero value otherwise.

### GetWalletIpOk

`func (o *PaymentForm) GetWalletIpOk() (*string, bool)`

GetWalletIpOk returns a tuple with the WalletIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletIp

`func (o *PaymentForm) SetWalletIp(v string)`

SetWalletIp sets WalletIp field to given value.

### HasWalletIp

`func (o *PaymentForm) HasWalletIp() bool`

HasWalletIp returns a boolean if a field has been set.

### GetWalletUa

`func (o *PaymentForm) GetWalletUa() string`

GetWalletUa returns the WalletUa field if non-nil, zero value otherwise.

### GetWalletUaOk

`func (o *PaymentForm) GetWalletUaOk() (*string, bool)`

GetWalletUaOk returns a tuple with the WalletUa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletUa

`func (o *PaymentForm) SetWalletUa(v string)`

SetWalletUa sets WalletUa field to given value.

### HasWalletUa

`func (o *PaymentForm) HasWalletUa() bool`

HasWalletUa returns a boolean if a field has been set.

### GetAmountTotRange

`func (o *PaymentForm) GetAmountTotRange() string`

GetAmountTotRange returns the AmountTotRange field if non-nil, zero value otherwise.

### GetAmountTotRangeOk

`func (o *PaymentForm) GetAmountTotRangeOk() (*string, bool)`

GetAmountTotRangeOk returns a tuple with the AmountTotRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountTotRange

`func (o *PaymentForm) SetAmountTotRange(v string)`

SetAmountTotRange sets AmountTotRange field to given value.

### HasAmountTotRange

`func (o *PaymentForm) HasAmountTotRange() bool`

HasAmountTotRange returns a boolean if a field has been set.

### GetAmountCom

`func (o *PaymentForm) GetAmountCom() string`

GetAmountCom returns the AmountCom field if non-nil, zero value otherwise.

### GetAmountComOk

`func (o *PaymentForm) GetAmountComOk() (*string, bool)`

GetAmountComOk returns a tuple with the AmountCom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCom

`func (o *PaymentForm) SetAmountCom(v string)`

SetAmountCom sets AmountCom field to given value.

### HasAmountCom

`func (o *PaymentForm) HasAmountCom() bool`

HasAmountCom returns a boolean if a field has been set.

### GetLanguage

`func (o *PaymentForm) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PaymentForm) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PaymentForm) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *PaymentForm) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetVersion

`func (o *PaymentForm) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PaymentForm) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PaymentForm) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PaymentForm) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetStyle

`func (o *PaymentForm) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *PaymentForm) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *PaymentForm) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *PaymentForm) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### GetAtosStyle

`func (o *PaymentForm) GetAtosStyle() string`

GetAtosStyle returns the AtosStyle field if non-nil, zero value otherwise.

### GetAtosStyleOk

`func (o *PaymentForm) GetAtosStyleOk() (*string, bool)`

GetAtosStyleOk returns a tuple with the AtosStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtosStyle

`func (o *PaymentForm) SetAtosStyle(v string)`

SetAtosStyle sets AtosStyle field to given value.

### HasAtosStyle

`func (o *PaymentForm) HasAtosStyle() bool`

HasAtosStyle returns a boolean if a field has been set.

### GetNotifUrl

`func (o *PaymentForm) GetNotifUrl() string`

GetNotifUrl returns the NotifUrl field if non-nil, zero value otherwise.

### GetNotifUrlOk

`func (o *PaymentForm) GetNotifUrlOk() (*string, bool)`

GetNotifUrlOk returns a tuple with the NotifUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUrl

`func (o *PaymentForm) SetNotifUrl(v string)`

SetNotifUrl sets NotifUrl field to given value.

### HasNotifUrl

`func (o *PaymentForm) HasNotifUrl() bool`

HasNotifUrl returns a boolean if a field has been set.

### GetOptions

`func (o *PaymentForm) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PaymentForm) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PaymentForm) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PaymentForm) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetActive

`func (o *PaymentForm) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PaymentForm) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PaymentForm) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PaymentForm) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetRaisonSociale

`func (o *PaymentForm) GetRaisonSociale() string`

GetRaisonSociale returns the RaisonSociale field if non-nil, zero value otherwise.

### GetRaisonSocialeOk

`func (o *PaymentForm) GetRaisonSocialeOk() (*string, bool)`

GetRaisonSocialeOk returns a tuple with the RaisonSociale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaisonSociale

`func (o *PaymentForm) SetRaisonSociale(v string)`

SetRaisonSociale sets RaisonSociale field to given value.

### HasRaisonSociale

`func (o *PaymentForm) HasRaisonSociale() bool`

HasRaisonSociale returns a boolean if a field has been set.

### GetId

`func (o *PaymentForm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentForm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentForm) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentForm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOptId

`func (o *PaymentForm) GetOptId() string`

GetOptId returns the OptId field if non-nil, zero value otherwise.

### GetOptIdOk

`func (o *PaymentForm) GetOptIdOk() (*string, bool)`

GetOptIdOk returns a tuple with the OptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptId

`func (o *PaymentForm) SetOptId(v string)`

SetOptId sets OptId field to given value.

### HasOptId

`func (o *PaymentForm) HasOptId() bool`

HasOptId returns a boolean if a field has been set.

### GetAccountPayer

`func (o *PaymentForm) GetAccountPayer() string`

GetAccountPayer returns the AccountPayer field if non-nil, zero value otherwise.

### GetAccountPayerOk

`func (o *PaymentForm) GetAccountPayerOk() (*string, bool)`

GetAccountPayerOk returns a tuple with the AccountPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPayer

`func (o *PaymentForm) SetAccountPayer(v string)`

SetAccountPayer sets AccountPayer field to given value.

### HasAccountPayer

`func (o *PaymentForm) HasAccountPayer() bool`

HasAccountPayer returns a boolean if a field has been set.

### GetAccountReceiver

`func (o *PaymentForm) GetAccountReceiver() string`

GetAccountReceiver returns the AccountReceiver field if non-nil, zero value otherwise.

### GetAccountReceiverOk

`func (o *PaymentForm) GetAccountReceiverOk() (*string, bool)`

GetAccountReceiverOk returns a tuple with the AccountReceiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountReceiver

`func (o *PaymentForm) SetAccountReceiver(v string)`

SetAccountReceiver sets AccountReceiver field to given value.

### HasAccountReceiver

`func (o *PaymentForm) HasAccountReceiver() bool`

HasAccountReceiver returns a boolean if a field has been set.

### GetComment

`func (o *PaymentForm) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PaymentForm) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PaymentForm) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PaymentForm) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetReturnUrl

`func (o *PaymentForm) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *PaymentForm) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *PaymentForm) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *PaymentForm) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetCancelUrl

`func (o *PaymentForm) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *PaymentForm) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *PaymentForm) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *PaymentForm) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### GetErrorUrl

`func (o *PaymentForm) GetErrorUrl() string`

GetErrorUrl returns the ErrorUrl field if non-nil, zero value otherwise.

### GetErrorUrlOk

`func (o *PaymentForm) GetErrorUrlOk() (*string, bool)`

GetErrorUrlOk returns a tuple with the ErrorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorUrl

`func (o *PaymentForm) SetErrorUrl(v string)`

SetErrorUrl sets ErrorUrl field to given value.

### HasErrorUrl

`func (o *PaymentForm) HasErrorUrl() bool`

HasErrorUrl returns a boolean if a field has been set.

### GetFirstNamePayer

`func (o *PaymentForm) GetFirstNamePayer() string`

GetFirstNamePayer returns the FirstNamePayer field if non-nil, zero value otherwise.

### GetFirstNamePayerOk

`func (o *PaymentForm) GetFirstNamePayerOk() (*string, bool)`

GetFirstNamePayerOk returns a tuple with the FirstNamePayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNamePayer

`func (o *PaymentForm) SetFirstNamePayer(v string)`

SetFirstNamePayer sets FirstNamePayer field to given value.

### HasFirstNamePayer

`func (o *PaymentForm) HasFirstNamePayer() bool`

HasFirstNamePayer returns a boolean if a field has been set.

### GetLastNamePayer

`func (o *PaymentForm) GetLastNamePayer() string`

GetLastNamePayer returns the LastNamePayer field if non-nil, zero value otherwise.

### GetLastNamePayerOk

`func (o *PaymentForm) GetLastNamePayerOk() (*string, bool)`

GetLastNamePayerOk returns a tuple with the LastNamePayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNamePayer

`func (o *PaymentForm) SetLastNamePayer(v string)`

SetLastNamePayer sets LastNamePayer field to given value.

### HasLastNamePayer

`func (o *PaymentForm) HasLastNamePayer() bool`

HasLastNamePayer returns a boolean if a field has been set.

### GetEmailPayer

`func (o *PaymentForm) GetEmailPayer() string`

GetEmailPayer returns the EmailPayer field if non-nil, zero value otherwise.

### GetEmailPayerOk

`func (o *PaymentForm) GetEmailPayerOk() (*string, bool)`

GetEmailPayerOk returns a tuple with the EmailPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailPayer

`func (o *PaymentForm) SetEmailPayer(v string)`

SetEmailPayer sets EmailPayer field to given value.

### HasEmailPayer

`func (o *PaymentForm) HasEmailPayer() bool`

HasEmailPayer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


