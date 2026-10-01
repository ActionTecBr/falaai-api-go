# RiskAuditAudioEventModelV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** | MAC model text (i18n) | [optional] [default to ""]
**Description** | Pointer to **string** | MAC description (i18n) | [optional] [default to ""]
**WindowsS** | Pointer to **map[string]interface{}** | Temporal windows (s) | [optional] [default to {}]

## Methods

### NewRiskAuditAudioEventModelV2

`func NewRiskAuditAudioEventModelV2() *RiskAuditAudioEventModelV2`

NewRiskAuditAudioEventModelV2 instantiates a new RiskAuditAudioEventModelV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAuditAudioEventModelV2WithDefaults

`func NewRiskAuditAudioEventModelV2WithDefaults() *RiskAuditAudioEventModelV2`

NewRiskAuditAudioEventModelV2WithDefaults instantiates a new RiskAuditAudioEventModelV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *RiskAuditAudioEventModelV2) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *RiskAuditAudioEventModelV2) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *RiskAuditAudioEventModelV2) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *RiskAuditAudioEventModelV2) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetDescription

`func (o *RiskAuditAudioEventModelV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RiskAuditAudioEventModelV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RiskAuditAudioEventModelV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RiskAuditAudioEventModelV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWindowsS

`func (o *RiskAuditAudioEventModelV2) GetWindowsS() map[string]interface{}`

GetWindowsS returns the WindowsS field if non-nil, zero value otherwise.

### GetWindowsSOk

`func (o *RiskAuditAudioEventModelV2) GetWindowsSOk() (*map[string]interface{}, bool)`

GetWindowsSOk returns a tuple with the WindowsS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsS

`func (o *RiskAuditAudioEventModelV2) SetWindowsS(v map[string]interface{})`

SetWindowsS sets WindowsS field to given value.

### HasWindowsS

`func (o *RiskAuditAudioEventModelV2) HasWindowsS() bool`

HasWindowsS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


