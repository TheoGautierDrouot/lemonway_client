# ThreeDS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChallengeMode** | Pointer to **string** | If you would like to enforce 3D Secure for your client then use the challenge mode. Possible values are:   - No_preference (default value). The card holder&#39;s bank will decide whether or not to challange  - Force_threeDS | [optional] 

## Methods

### NewThreeDS

`func NewThreeDS() *ThreeDS`

NewThreeDS instantiates a new ThreeDS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSWithDefaults

`func NewThreeDSWithDefaults() *ThreeDS`

NewThreeDSWithDefaults instantiates a new ThreeDS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChallengeMode

`func (o *ThreeDS) GetChallengeMode() string`

GetChallengeMode returns the ChallengeMode field if non-nil, zero value otherwise.

### GetChallengeModeOk

`func (o *ThreeDS) GetChallengeModeOk() (*string, bool)`

GetChallengeModeOk returns a tuple with the ChallengeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeMode

`func (o *ThreeDS) SetChallengeMode(v string)`

SetChallengeMode sets ChallengeMode field to given value.

### HasChallengeMode

`func (o *ThreeDS) HasChallengeMode() bool`

HasChallengeMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


