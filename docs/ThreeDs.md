# ThreeDs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreeDSModeRequested** | Pointer to **string** | 3DS Exemption Request    Possible value:      NO_PREFERENCE: The issuing Bank can choose to activate or not 3DS v2. | [optional] 
**ThreeDSResult** | Pointer to **string** | 3DS Authentication Result    Possible values:      - CHALLENGE: 3DS v2 was activated and the card owner was asked to be strongly authenticated.    - FRICTIONLESS: 3DS v2 was not activated.    - NONE: We could not get the SCA result. | [optional] 

## Methods

### NewThreeDs

`func NewThreeDs() *ThreeDs`

NewThreeDs instantiates a new ThreeDs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDsWithDefaults

`func NewThreeDsWithDefaults() *ThreeDs`

NewThreeDsWithDefaults instantiates a new ThreeDs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreeDSModeRequested

`func (o *ThreeDs) GetThreeDSModeRequested() string`

GetThreeDSModeRequested returns the ThreeDSModeRequested field if non-nil, zero value otherwise.

### GetThreeDSModeRequestedOk

`func (o *ThreeDs) GetThreeDSModeRequestedOk() (*string, bool)`

GetThreeDSModeRequestedOk returns a tuple with the ThreeDSModeRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSModeRequested

`func (o *ThreeDs) SetThreeDSModeRequested(v string)`

SetThreeDSModeRequested sets ThreeDSModeRequested field to given value.

### HasThreeDSModeRequested

`func (o *ThreeDs) HasThreeDSModeRequested() bool`

HasThreeDSModeRequested returns a boolean if a field has been set.

### GetThreeDSResult

`func (o *ThreeDs) GetThreeDSResult() string`

GetThreeDSResult returns the ThreeDSResult field if non-nil, zero value otherwise.

### GetThreeDSResultOk

`func (o *ThreeDs) GetThreeDSResultOk() (*string, bool)`

GetThreeDSResultOk returns a tuple with the ThreeDSResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSResult

`func (o *ThreeDs) SetThreeDSResult(v string)`

SetThreeDSResult sets ThreeDSResult field to given value.

### HasThreeDSResult

`func (o *ThreeDs) HasThreeDSResult() bool`

HasThreeDSResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


