# CreatePaymentFormInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptId** | Pointer to **string** | To cross-reference a buyer with a payment, use &#x60;optId&#x60; to help associate them with one another.     **Note:** Do not use special characters in this field, for example: #, @, and !. | [optional] 
**WalletPayer** | Pointer to **string** | Associated with the **Payer Wallet**.    If this field is filled then the money reaches this wallet before arriving at the beneficiary wallet (via a transfer wallet to wallet).  **Important:** Don&#39;t put the &lt;b&gt;SC Wallet&lt;/b&gt; here, it won&#39;t work. You cannot debit the &lt;b&gt;SC Wallet&lt;/b&gt; with a credit card. | [optional] 
**WalletReceiver** | Pointer to **string** | Beneficiary Wallet. If this field is not complete, then the end-user must to fill its value on the payment form.   &lt;b&gt;We recommend you complete the form instead of your end-user&lt;/b&gt;  Don&#39;t put the &lt;b&gt;SC Wallet&lt;/b&gt; here, it won&#39;t work. You cannot credit the &lt;b&gt;SC Wallet&lt;/b&gt; with a credit card. | [optional] 
**TotalAmount** | Pointer to **string** | Amount or a range of the amount to be debited.  1. If this field is configured with an interval (eg, 15.30-500.26) then the final customer will have to enter an appropriate amount in the form.  2. If this field is not filled then the end-user can enter any amount to the form.  3. If this field is filled with a precise value (eg 15.60), then the end-user has no choice in the amount field of the form.  Amounts are given as integer numbers in cents.  Note: If you enter a fixed amount | [optional] 
**CommissionAmount** | Pointer to **int32** | Your Fee  Amounts are given as integer numbers in cents | [optional] 
**Comment** | Pointer to **string** | Comment Regarding the Transaction | [optional] 
**ReturnUrl** | Pointer to **string** | URL redirection after the payment procedure is successfully finished | [optional] 
**CancelUrl** | Pointer to **string** | URL redirection after the payment procedure is cancelled | [optional] 
**ErrorUrl** | Pointer to **string** | URL edirection after the payment procedure is failed | [optional] 
**FirstNamePayer** | Pointer to **string** | Payer&#39;s First Name  If this field is not filled then the end-user have to fill it in the payment form | [optional] 
**LastNamePayer** | Pointer to **string** | Payer&#39;s Last Name  If this field is not filled then the end-user have to fill it in the payment form | [optional] 
**EmailPayer** | Pointer to **string** | Payer&#39;s Email  If this field is not filled then the end-user have to fill it in the payment form | [optional] 
**Style** | Pointer to **string** | Link to a custom CSS Stylesheet  The stylesheet should be publicly accessible via **HTTPS* | [optional] 
**AtosStyle** | Pointer to **string** |  | [optional] 
**NotifUrl** | Pointer to **string** | At the end of the payment procedure, an HTTP POST message containing the payment status (PAID, ERROR, CANCEL) is sent to this address.   It is possible that the same notification might be sent several times. | [optional] 
**Options** | Pointer to **string** | Reserved for a future version | [optional] 

## Methods

### NewCreatePaymentFormInput

`func NewCreatePaymentFormInput() *CreatePaymentFormInput`

NewCreatePaymentFormInput instantiates a new CreatePaymentFormInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaymentFormInputWithDefaults

`func NewCreatePaymentFormInputWithDefaults() *CreatePaymentFormInput`

NewCreatePaymentFormInputWithDefaults instantiates a new CreatePaymentFormInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptId

`func (o *CreatePaymentFormInput) GetOptId() string`

GetOptId returns the OptId field if non-nil, zero value otherwise.

### GetOptIdOk

`func (o *CreatePaymentFormInput) GetOptIdOk() (*string, bool)`

GetOptIdOk returns a tuple with the OptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptId

`func (o *CreatePaymentFormInput) SetOptId(v string)`

SetOptId sets OptId field to given value.

### HasOptId

`func (o *CreatePaymentFormInput) HasOptId() bool`

HasOptId returns a boolean if a field has been set.

### GetWalletPayer

`func (o *CreatePaymentFormInput) GetWalletPayer() string`

GetWalletPayer returns the WalletPayer field if non-nil, zero value otherwise.

### GetWalletPayerOk

`func (o *CreatePaymentFormInput) GetWalletPayerOk() (*string, bool)`

GetWalletPayerOk returns a tuple with the WalletPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletPayer

`func (o *CreatePaymentFormInput) SetWalletPayer(v string)`

SetWalletPayer sets WalletPayer field to given value.

### HasWalletPayer

`func (o *CreatePaymentFormInput) HasWalletPayer() bool`

HasWalletPayer returns a boolean if a field has been set.

### GetWalletReceiver

`func (o *CreatePaymentFormInput) GetWalletReceiver() string`

GetWalletReceiver returns the WalletReceiver field if non-nil, zero value otherwise.

### GetWalletReceiverOk

`func (o *CreatePaymentFormInput) GetWalletReceiverOk() (*string, bool)`

GetWalletReceiverOk returns a tuple with the WalletReceiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletReceiver

`func (o *CreatePaymentFormInput) SetWalletReceiver(v string)`

SetWalletReceiver sets WalletReceiver field to given value.

### HasWalletReceiver

`func (o *CreatePaymentFormInput) HasWalletReceiver() bool`

HasWalletReceiver returns a boolean if a field has been set.

### GetTotalAmount

`func (o *CreatePaymentFormInput) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *CreatePaymentFormInput) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *CreatePaymentFormInput) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *CreatePaymentFormInput) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetCommissionAmount

`func (o *CreatePaymentFormInput) GetCommissionAmount() int32`

GetCommissionAmount returns the CommissionAmount field if non-nil, zero value otherwise.

### GetCommissionAmountOk

`func (o *CreatePaymentFormInput) GetCommissionAmountOk() (*int32, bool)`

GetCommissionAmountOk returns a tuple with the CommissionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionAmount

`func (o *CreatePaymentFormInput) SetCommissionAmount(v int32)`

SetCommissionAmount sets CommissionAmount field to given value.

### HasCommissionAmount

`func (o *CreatePaymentFormInput) HasCommissionAmount() bool`

HasCommissionAmount returns a boolean if a field has been set.

### GetComment

`func (o *CreatePaymentFormInput) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreatePaymentFormInput) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreatePaymentFormInput) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreatePaymentFormInput) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetReturnUrl

`func (o *CreatePaymentFormInput) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *CreatePaymentFormInput) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *CreatePaymentFormInput) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *CreatePaymentFormInput) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetCancelUrl

`func (o *CreatePaymentFormInput) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *CreatePaymentFormInput) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *CreatePaymentFormInput) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *CreatePaymentFormInput) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### GetErrorUrl

`func (o *CreatePaymentFormInput) GetErrorUrl() string`

GetErrorUrl returns the ErrorUrl field if non-nil, zero value otherwise.

### GetErrorUrlOk

`func (o *CreatePaymentFormInput) GetErrorUrlOk() (*string, bool)`

GetErrorUrlOk returns a tuple with the ErrorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorUrl

`func (o *CreatePaymentFormInput) SetErrorUrl(v string)`

SetErrorUrl sets ErrorUrl field to given value.

### HasErrorUrl

`func (o *CreatePaymentFormInput) HasErrorUrl() bool`

HasErrorUrl returns a boolean if a field has been set.

### GetFirstNamePayer

`func (o *CreatePaymentFormInput) GetFirstNamePayer() string`

GetFirstNamePayer returns the FirstNamePayer field if non-nil, zero value otherwise.

### GetFirstNamePayerOk

`func (o *CreatePaymentFormInput) GetFirstNamePayerOk() (*string, bool)`

GetFirstNamePayerOk returns a tuple with the FirstNamePayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNamePayer

`func (o *CreatePaymentFormInput) SetFirstNamePayer(v string)`

SetFirstNamePayer sets FirstNamePayer field to given value.

### HasFirstNamePayer

`func (o *CreatePaymentFormInput) HasFirstNamePayer() bool`

HasFirstNamePayer returns a boolean if a field has been set.

### GetLastNamePayer

`func (o *CreatePaymentFormInput) GetLastNamePayer() string`

GetLastNamePayer returns the LastNamePayer field if non-nil, zero value otherwise.

### GetLastNamePayerOk

`func (o *CreatePaymentFormInput) GetLastNamePayerOk() (*string, bool)`

GetLastNamePayerOk returns a tuple with the LastNamePayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNamePayer

`func (o *CreatePaymentFormInput) SetLastNamePayer(v string)`

SetLastNamePayer sets LastNamePayer field to given value.

### HasLastNamePayer

`func (o *CreatePaymentFormInput) HasLastNamePayer() bool`

HasLastNamePayer returns a boolean if a field has been set.

### GetEmailPayer

`func (o *CreatePaymentFormInput) GetEmailPayer() string`

GetEmailPayer returns the EmailPayer field if non-nil, zero value otherwise.

### GetEmailPayerOk

`func (o *CreatePaymentFormInput) GetEmailPayerOk() (*string, bool)`

GetEmailPayerOk returns a tuple with the EmailPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailPayer

`func (o *CreatePaymentFormInput) SetEmailPayer(v string)`

SetEmailPayer sets EmailPayer field to given value.

### HasEmailPayer

`func (o *CreatePaymentFormInput) HasEmailPayer() bool`

HasEmailPayer returns a boolean if a field has been set.

### GetStyle

`func (o *CreatePaymentFormInput) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *CreatePaymentFormInput) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *CreatePaymentFormInput) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *CreatePaymentFormInput) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### GetAtosStyle

`func (o *CreatePaymentFormInput) GetAtosStyle() string`

GetAtosStyle returns the AtosStyle field if non-nil, zero value otherwise.

### GetAtosStyleOk

`func (o *CreatePaymentFormInput) GetAtosStyleOk() (*string, bool)`

GetAtosStyleOk returns a tuple with the AtosStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtosStyle

`func (o *CreatePaymentFormInput) SetAtosStyle(v string)`

SetAtosStyle sets AtosStyle field to given value.

### HasAtosStyle

`func (o *CreatePaymentFormInput) HasAtosStyle() bool`

HasAtosStyle returns a boolean if a field has been set.

### GetNotifUrl

`func (o *CreatePaymentFormInput) GetNotifUrl() string`

GetNotifUrl returns the NotifUrl field if non-nil, zero value otherwise.

### GetNotifUrlOk

`func (o *CreatePaymentFormInput) GetNotifUrlOk() (*string, bool)`

GetNotifUrlOk returns a tuple with the NotifUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUrl

`func (o *CreatePaymentFormInput) SetNotifUrl(v string)`

SetNotifUrl sets NotifUrl field to given value.

### HasNotifUrl

`func (o *CreatePaymentFormInput) HasNotifUrl() bool`

HasNotifUrl returns a boolean if a field has been set.

### GetOptions

`func (o *CreatePaymentFormInput) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CreatePaymentFormInput) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CreatePaymentFormInput) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CreatePaymentFormInput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


