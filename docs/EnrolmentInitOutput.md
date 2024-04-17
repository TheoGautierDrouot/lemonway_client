# EnrolmentInitOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Redirecturl** | Pointer to **string** | Redirect URL to Deutsche Post PostIdent | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewEnrolmentInitOutput

`func NewEnrolmentInitOutput() *EnrolmentInitOutput`

NewEnrolmentInitOutput instantiates a new EnrolmentInitOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrolmentInitOutputWithDefaults

`func NewEnrolmentInitOutputWithDefaults() *EnrolmentInitOutput`

NewEnrolmentInitOutputWithDefaults instantiates a new EnrolmentInitOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirecturl

`func (o *EnrolmentInitOutput) GetRedirecturl() string`

GetRedirecturl returns the Redirecturl field if non-nil, zero value otherwise.

### GetRedirecturlOk

`func (o *EnrolmentInitOutput) GetRedirecturlOk() (*string, bool)`

GetRedirecturlOk returns a tuple with the Redirecturl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirecturl

`func (o *EnrolmentInitOutput) SetRedirecturl(v string)`

SetRedirecturl sets Redirecturl field to given value.

### HasRedirecturl

`func (o *EnrolmentInitOutput) HasRedirecturl() bool`

HasRedirecturl returns a boolean if a field has been set.

### GetError

`func (o *EnrolmentInitOutput) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *EnrolmentInitOutput) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *EnrolmentInitOutput) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *EnrolmentInitOutput) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


