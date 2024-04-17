# CustomerAccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerAccountId** | Pointer to **string** | Additional information about the account optionally provided by the 3-D Secure Requestor | [optional] 
**NumberOfPurchase180Days** | Pointer to **string** | The number of customer transactions made over the last six months (last 180 days) | [optional] 
**NumberOfTransactionYear** | Pointer to **string** | Number of accepted or abandoned transactions in the last year on the customer account | [optional] 
**CustomerAccountCreationDate** | Pointer to **string** | The customer account creation date. Format: YYYY/MM/DD | [optional] 
**NumberOfAttemptsAddCard24Hours** | Pointer to **string** | The number of add card attempts in 24 hours | [optional] 
**SuspiciousActivityIndicator** | Pointer to **string** | Specifies whether a suspicious activity is detected on customer account. Possible values: true or false. | [optional] 
**NumberOfTransaction24Hours** | Pointer to **string** | Number of abandoned or successful transactions in the last 24 hours on the customer account | [optional] 
**CustomerAccountChangeDate** | Pointer to **string** | Last date the customer account was changed. Format: YYYY/MM/DD | [optional] 
**PasswordChangeDate** | Pointer to **string** | Date of last change of password of the customer account. Format: YYYY/MM/DD | [optional] 
**AddPaymentMeanDate** | Pointer to **string** | Date of last added form of payment made to account. For example a new card registered to the account. Format: YYYY/MM/DD | [optional] 

## Methods

### NewCustomerAccountInfo

`func NewCustomerAccountInfo() *CustomerAccountInfo`

NewCustomerAccountInfo instantiates a new CustomerAccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAccountInfoWithDefaults

`func NewCustomerAccountInfoWithDefaults() *CustomerAccountInfo`

NewCustomerAccountInfoWithDefaults instantiates a new CustomerAccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerAccountId

`func (o *CustomerAccountInfo) GetCustomerAccountId() string`

GetCustomerAccountId returns the CustomerAccountId field if non-nil, zero value otherwise.

### GetCustomerAccountIdOk

`func (o *CustomerAccountInfo) GetCustomerAccountIdOk() (*string, bool)`

GetCustomerAccountIdOk returns a tuple with the CustomerAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAccountId

`func (o *CustomerAccountInfo) SetCustomerAccountId(v string)`

SetCustomerAccountId sets CustomerAccountId field to given value.

### HasCustomerAccountId

`func (o *CustomerAccountInfo) HasCustomerAccountId() bool`

HasCustomerAccountId returns a boolean if a field has been set.

### GetNumberOfPurchase180Days

`func (o *CustomerAccountInfo) GetNumberOfPurchase180Days() string`

GetNumberOfPurchase180Days returns the NumberOfPurchase180Days field if non-nil, zero value otherwise.

### GetNumberOfPurchase180DaysOk

`func (o *CustomerAccountInfo) GetNumberOfPurchase180DaysOk() (*string, bool)`

GetNumberOfPurchase180DaysOk returns a tuple with the NumberOfPurchase180Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPurchase180Days

`func (o *CustomerAccountInfo) SetNumberOfPurchase180Days(v string)`

SetNumberOfPurchase180Days sets NumberOfPurchase180Days field to given value.

### HasNumberOfPurchase180Days

`func (o *CustomerAccountInfo) HasNumberOfPurchase180Days() bool`

HasNumberOfPurchase180Days returns a boolean if a field has been set.

### GetNumberOfTransactionYear

`func (o *CustomerAccountInfo) GetNumberOfTransactionYear() string`

GetNumberOfTransactionYear returns the NumberOfTransactionYear field if non-nil, zero value otherwise.

### GetNumberOfTransactionYearOk

`func (o *CustomerAccountInfo) GetNumberOfTransactionYearOk() (*string, bool)`

GetNumberOfTransactionYearOk returns a tuple with the NumberOfTransactionYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfTransactionYear

`func (o *CustomerAccountInfo) SetNumberOfTransactionYear(v string)`

SetNumberOfTransactionYear sets NumberOfTransactionYear field to given value.

### HasNumberOfTransactionYear

`func (o *CustomerAccountInfo) HasNumberOfTransactionYear() bool`

HasNumberOfTransactionYear returns a boolean if a field has been set.

### GetCustomerAccountCreationDate

`func (o *CustomerAccountInfo) GetCustomerAccountCreationDate() string`

GetCustomerAccountCreationDate returns the CustomerAccountCreationDate field if non-nil, zero value otherwise.

### GetCustomerAccountCreationDateOk

`func (o *CustomerAccountInfo) GetCustomerAccountCreationDateOk() (*string, bool)`

GetCustomerAccountCreationDateOk returns a tuple with the CustomerAccountCreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAccountCreationDate

`func (o *CustomerAccountInfo) SetCustomerAccountCreationDate(v string)`

SetCustomerAccountCreationDate sets CustomerAccountCreationDate field to given value.

### HasCustomerAccountCreationDate

`func (o *CustomerAccountInfo) HasCustomerAccountCreationDate() bool`

HasCustomerAccountCreationDate returns a boolean if a field has been set.

### GetNumberOfAttemptsAddCard24Hours

`func (o *CustomerAccountInfo) GetNumberOfAttemptsAddCard24Hours() string`

GetNumberOfAttemptsAddCard24Hours returns the NumberOfAttemptsAddCard24Hours field if non-nil, zero value otherwise.

### GetNumberOfAttemptsAddCard24HoursOk

`func (o *CustomerAccountInfo) GetNumberOfAttemptsAddCard24HoursOk() (*string, bool)`

GetNumberOfAttemptsAddCard24HoursOk returns a tuple with the NumberOfAttemptsAddCard24Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAttemptsAddCard24Hours

`func (o *CustomerAccountInfo) SetNumberOfAttemptsAddCard24Hours(v string)`

SetNumberOfAttemptsAddCard24Hours sets NumberOfAttemptsAddCard24Hours field to given value.

### HasNumberOfAttemptsAddCard24Hours

`func (o *CustomerAccountInfo) HasNumberOfAttemptsAddCard24Hours() bool`

HasNumberOfAttemptsAddCard24Hours returns a boolean if a field has been set.

### GetSuspiciousActivityIndicator

`func (o *CustomerAccountInfo) GetSuspiciousActivityIndicator() string`

GetSuspiciousActivityIndicator returns the SuspiciousActivityIndicator field if non-nil, zero value otherwise.

### GetSuspiciousActivityIndicatorOk

`func (o *CustomerAccountInfo) GetSuspiciousActivityIndicatorOk() (*string, bool)`

GetSuspiciousActivityIndicatorOk returns a tuple with the SuspiciousActivityIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspiciousActivityIndicator

`func (o *CustomerAccountInfo) SetSuspiciousActivityIndicator(v string)`

SetSuspiciousActivityIndicator sets SuspiciousActivityIndicator field to given value.

### HasSuspiciousActivityIndicator

`func (o *CustomerAccountInfo) HasSuspiciousActivityIndicator() bool`

HasSuspiciousActivityIndicator returns a boolean if a field has been set.

### GetNumberOfTransaction24Hours

`func (o *CustomerAccountInfo) GetNumberOfTransaction24Hours() string`

GetNumberOfTransaction24Hours returns the NumberOfTransaction24Hours field if non-nil, zero value otherwise.

### GetNumberOfTransaction24HoursOk

`func (o *CustomerAccountInfo) GetNumberOfTransaction24HoursOk() (*string, bool)`

GetNumberOfTransaction24HoursOk returns a tuple with the NumberOfTransaction24Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfTransaction24Hours

`func (o *CustomerAccountInfo) SetNumberOfTransaction24Hours(v string)`

SetNumberOfTransaction24Hours sets NumberOfTransaction24Hours field to given value.

### HasNumberOfTransaction24Hours

`func (o *CustomerAccountInfo) HasNumberOfTransaction24Hours() bool`

HasNumberOfTransaction24Hours returns a boolean if a field has been set.

### GetCustomerAccountChangeDate

`func (o *CustomerAccountInfo) GetCustomerAccountChangeDate() string`

GetCustomerAccountChangeDate returns the CustomerAccountChangeDate field if non-nil, zero value otherwise.

### GetCustomerAccountChangeDateOk

`func (o *CustomerAccountInfo) GetCustomerAccountChangeDateOk() (*string, bool)`

GetCustomerAccountChangeDateOk returns a tuple with the CustomerAccountChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAccountChangeDate

`func (o *CustomerAccountInfo) SetCustomerAccountChangeDate(v string)`

SetCustomerAccountChangeDate sets CustomerAccountChangeDate field to given value.

### HasCustomerAccountChangeDate

`func (o *CustomerAccountInfo) HasCustomerAccountChangeDate() bool`

HasCustomerAccountChangeDate returns a boolean if a field has been set.

### GetPasswordChangeDate

`func (o *CustomerAccountInfo) GetPasswordChangeDate() string`

GetPasswordChangeDate returns the PasswordChangeDate field if non-nil, zero value otherwise.

### GetPasswordChangeDateOk

`func (o *CustomerAccountInfo) GetPasswordChangeDateOk() (*string, bool)`

GetPasswordChangeDateOk returns a tuple with the PasswordChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChangeDate

`func (o *CustomerAccountInfo) SetPasswordChangeDate(v string)`

SetPasswordChangeDate sets PasswordChangeDate field to given value.

### HasPasswordChangeDate

`func (o *CustomerAccountInfo) HasPasswordChangeDate() bool`

HasPasswordChangeDate returns a boolean if a field has been set.

### GetAddPaymentMeanDate

`func (o *CustomerAccountInfo) GetAddPaymentMeanDate() string`

GetAddPaymentMeanDate returns the AddPaymentMeanDate field if non-nil, zero value otherwise.

### GetAddPaymentMeanDateOk

`func (o *CustomerAccountInfo) GetAddPaymentMeanDateOk() (*string, bool)`

GetAddPaymentMeanDateOk returns a tuple with the AddPaymentMeanDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddPaymentMeanDate

`func (o *CustomerAccountInfo) SetAddPaymentMeanDate(v string)`

SetAddPaymentMeanDate sets AddPaymentMeanDate field to given value.

### HasAddPaymentMeanDate

`func (o *CustomerAccountInfo) HasAddPaymentMeanDate() bool`

HasAddPaymentMeanDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


