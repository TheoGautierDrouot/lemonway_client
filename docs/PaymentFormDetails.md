# PaymentFormDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentId** | Pointer to **string** |  | [optional] 
**ParentComment** | Pointer to **string** |  | [optional] 
**AmountTot** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UtcDate** | Pointer to **int32** | UTC Unix Timestamp | [optional] 
**Created** | Pointer to **int32** | UTC Unix Timestamp | [optional] 
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

### NewPaymentFormDetails

`func NewPaymentFormDetails() *PaymentFormDetails`

NewPaymentFormDetails instantiates a new PaymentFormDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentFormDetailsWithDefaults

`func NewPaymentFormDetailsWithDefaults() *PaymentFormDetails`

NewPaymentFormDetailsWithDefaults instantiates a new PaymentFormDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentId

`func (o *PaymentFormDetails) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *PaymentFormDetails) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *PaymentFormDetails) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *PaymentFormDetails) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetParentComment

`func (o *PaymentFormDetails) GetParentComment() string`

GetParentComment returns the ParentComment field if non-nil, zero value otherwise.

### GetParentCommentOk

`func (o *PaymentFormDetails) GetParentCommentOk() (*string, bool)`

GetParentCommentOk returns a tuple with the ParentComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentComment

`func (o *PaymentFormDetails) SetParentComment(v string)`

SetParentComment sets ParentComment field to given value.

### HasParentComment

`func (o *PaymentFormDetails) HasParentComment() bool`

HasParentComment returns a boolean if a field has been set.

### GetAmountTot

`func (o *PaymentFormDetails) GetAmountTot() string`

GetAmountTot returns the AmountTot field if non-nil, zero value otherwise.

### GetAmountTotOk

`func (o *PaymentFormDetails) GetAmountTotOk() (*string, bool)`

GetAmountTotOk returns a tuple with the AmountTot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountTot

`func (o *PaymentFormDetails) SetAmountTot(v string)`

SetAmountTot sets AmountTot field to given value.

### HasAmountTot

`func (o *PaymentFormDetails) HasAmountTot() bool`

HasAmountTot returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentFormDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentFormDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentFormDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentFormDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUtcDate

`func (o *PaymentFormDetails) GetUtcDate() int32`

GetUtcDate returns the UtcDate field if non-nil, zero value otherwise.

### GetUtcDateOk

`func (o *PaymentFormDetails) GetUtcDateOk() (*int32, bool)`

GetUtcDateOk returns a tuple with the UtcDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcDate

`func (o *PaymentFormDetails) SetUtcDate(v int32)`

SetUtcDate sets UtcDate field to given value.

### HasUtcDate

`func (o *PaymentFormDetails) HasUtcDate() bool`

HasUtcDate returns a boolean if a field has been set.

### GetCreated

`func (o *PaymentFormDetails) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PaymentFormDetails) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PaymentFormDetails) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PaymentFormDetails) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *PaymentFormDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentFormDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentFormDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentFormDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOptId

`func (o *PaymentFormDetails) GetOptId() string`

GetOptId returns the OptId field if non-nil, zero value otherwise.

### GetOptIdOk

`func (o *PaymentFormDetails) GetOptIdOk() (*string, bool)`

GetOptIdOk returns a tuple with the OptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptId

`func (o *PaymentFormDetails) SetOptId(v string)`

SetOptId sets OptId field to given value.

### HasOptId

`func (o *PaymentFormDetails) HasOptId() bool`

HasOptId returns a boolean if a field has been set.

### GetAccountPayer

`func (o *PaymentFormDetails) GetAccountPayer() string`

GetAccountPayer returns the AccountPayer field if non-nil, zero value otherwise.

### GetAccountPayerOk

`func (o *PaymentFormDetails) GetAccountPayerOk() (*string, bool)`

GetAccountPayerOk returns a tuple with the AccountPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPayer

`func (o *PaymentFormDetails) SetAccountPayer(v string)`

SetAccountPayer sets AccountPayer field to given value.

### HasAccountPayer

`func (o *PaymentFormDetails) HasAccountPayer() bool`

HasAccountPayer returns a boolean if a field has been set.

### GetAccountReceiver

`func (o *PaymentFormDetails) GetAccountReceiver() string`

GetAccountReceiver returns the AccountReceiver field if non-nil, zero value otherwise.

### GetAccountReceiverOk

`func (o *PaymentFormDetails) GetAccountReceiverOk() (*string, bool)`

GetAccountReceiverOk returns a tuple with the AccountReceiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountReceiver

`func (o *PaymentFormDetails) SetAccountReceiver(v string)`

SetAccountReceiver sets AccountReceiver field to given value.

### HasAccountReceiver

`func (o *PaymentFormDetails) HasAccountReceiver() bool`

HasAccountReceiver returns a boolean if a field has been set.

### GetComment

`func (o *PaymentFormDetails) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PaymentFormDetails) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PaymentFormDetails) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PaymentFormDetails) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetReturnUrl

`func (o *PaymentFormDetails) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *PaymentFormDetails) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *PaymentFormDetails) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *PaymentFormDetails) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetCancelUrl

`func (o *PaymentFormDetails) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *PaymentFormDetails) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *PaymentFormDetails) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *PaymentFormDetails) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### GetErrorUrl

`func (o *PaymentFormDetails) GetErrorUrl() string`

GetErrorUrl returns the ErrorUrl field if non-nil, zero value otherwise.

### GetErrorUrlOk

`func (o *PaymentFormDetails) GetErrorUrlOk() (*string, bool)`

GetErrorUrlOk returns a tuple with the ErrorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorUrl

`func (o *PaymentFormDetails) SetErrorUrl(v string)`

SetErrorUrl sets ErrorUrl field to given value.

### HasErrorUrl

`func (o *PaymentFormDetails) HasErrorUrl() bool`

HasErrorUrl returns a boolean if a field has been set.

### GetFirstNamePayer

`func (o *PaymentFormDetails) GetFirstNamePayer() string`

GetFirstNamePayer returns the FirstNamePayer field if non-nil, zero value otherwise.

### GetFirstNamePayerOk

`func (o *PaymentFormDetails) GetFirstNamePayerOk() (*string, bool)`

GetFirstNamePayerOk returns a tuple with the FirstNamePayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNamePayer

`func (o *PaymentFormDetails) SetFirstNamePayer(v string)`

SetFirstNamePayer sets FirstNamePayer field to given value.

### HasFirstNamePayer

`func (o *PaymentFormDetails) HasFirstNamePayer() bool`

HasFirstNamePayer returns a boolean if a field has been set.

### GetLastNamePayer

`func (o *PaymentFormDetails) GetLastNamePayer() string`

GetLastNamePayer returns the LastNamePayer field if non-nil, zero value otherwise.

### GetLastNamePayerOk

`func (o *PaymentFormDetails) GetLastNamePayerOk() (*string, bool)`

GetLastNamePayerOk returns a tuple with the LastNamePayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNamePayer

`func (o *PaymentFormDetails) SetLastNamePayer(v string)`

SetLastNamePayer sets LastNamePayer field to given value.

### HasLastNamePayer

`func (o *PaymentFormDetails) HasLastNamePayer() bool`

HasLastNamePayer returns a boolean if a field has been set.

### GetEmailPayer

`func (o *PaymentFormDetails) GetEmailPayer() string`

GetEmailPayer returns the EmailPayer field if non-nil, zero value otherwise.

### GetEmailPayerOk

`func (o *PaymentFormDetails) GetEmailPayerOk() (*string, bool)`

GetEmailPayerOk returns a tuple with the EmailPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailPayer

`func (o *PaymentFormDetails) SetEmailPayer(v string)`

SetEmailPayer sets EmailPayer field to given value.

### HasEmailPayer

`func (o *PaymentFormDetails) HasEmailPayer() bool`

HasEmailPayer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


