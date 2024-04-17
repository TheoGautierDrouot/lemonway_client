# BankBranchAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Street** | **string** | Number and Street | 
**ZipCode** | **string** | Postal or Zip Code | 
**City** | **string** | City | 

## Methods

### NewBankBranchAddress

`func NewBankBranchAddress(street string, zipCode string, city string, ) *BankBranchAddress`

NewBankBranchAddress instantiates a new BankBranchAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankBranchAddressWithDefaults

`func NewBankBranchAddressWithDefaults() *BankBranchAddress`

NewBankBranchAddressWithDefaults instantiates a new BankBranchAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreet

`func (o *BankBranchAddress) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *BankBranchAddress) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *BankBranchAddress) SetStreet(v string)`

SetStreet sets Street field to given value.


### GetZipCode

`func (o *BankBranchAddress) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *BankBranchAddress) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *BankBranchAddress) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.


### GetCity

`func (o *BankBranchAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *BankBranchAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *BankBranchAddress) SetCity(v string)`

SetCity sets City field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


