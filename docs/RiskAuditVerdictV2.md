# RiskAuditVerdictV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Human-readable verdict | [optional] 
**LevelCode** | Pointer to **string** | Classification level code | [optional] 
**Color** | Pointer to **string** | Level color | [optional] 
**Icon** | Pointer to **string** | Level icon | [optional] 
**RiskMatrix** | Pointer to **map[string]interface{}** | Risk matrix | [optional] [default to {}]
**AppliedActions** | Pointer to [**[]RiskAuditAppliedActionV2**](RiskAuditAppliedActionV2.md) | Applied actions | [optional] [default to {}]
**DecisionDetails** | Pointer to **interface{}** | Decision details | [optional] 

## Methods

### NewRiskAuditVerdictV2

`func NewRiskAuditVerdictV2() *RiskAuditVerdictV2`

NewRiskAuditVerdictV2 instantiates a new RiskAuditVerdictV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAuditVerdictV2WithDefaults

`func NewRiskAuditVerdictV2WithDefaults() *RiskAuditVerdictV2`

NewRiskAuditVerdictV2WithDefaults instantiates a new RiskAuditVerdictV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *RiskAuditVerdictV2) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RiskAuditVerdictV2) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RiskAuditVerdictV2) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *RiskAuditVerdictV2) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLevelCode

`func (o *RiskAuditVerdictV2) GetLevelCode() string`

GetLevelCode returns the LevelCode field if non-nil, zero value otherwise.

### GetLevelCodeOk

`func (o *RiskAuditVerdictV2) GetLevelCodeOk() (*string, bool)`

GetLevelCodeOk returns a tuple with the LevelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelCode

`func (o *RiskAuditVerdictV2) SetLevelCode(v string)`

SetLevelCode sets LevelCode field to given value.

### HasLevelCode

`func (o *RiskAuditVerdictV2) HasLevelCode() bool`

HasLevelCode returns a boolean if a field has been set.

### GetColor

`func (o *RiskAuditVerdictV2) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *RiskAuditVerdictV2) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *RiskAuditVerdictV2) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *RiskAuditVerdictV2) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetIcon

`func (o *RiskAuditVerdictV2) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *RiskAuditVerdictV2) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *RiskAuditVerdictV2) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *RiskAuditVerdictV2) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetRiskMatrix

`func (o *RiskAuditVerdictV2) GetRiskMatrix() map[string]interface{}`

GetRiskMatrix returns the RiskMatrix field if non-nil, zero value otherwise.

### GetRiskMatrixOk

`func (o *RiskAuditVerdictV2) GetRiskMatrixOk() (*map[string]interface{}, bool)`

GetRiskMatrixOk returns a tuple with the RiskMatrix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskMatrix

`func (o *RiskAuditVerdictV2) SetRiskMatrix(v map[string]interface{})`

SetRiskMatrix sets RiskMatrix field to given value.

### HasRiskMatrix

`func (o *RiskAuditVerdictV2) HasRiskMatrix() bool`

HasRiskMatrix returns a boolean if a field has been set.

### GetAppliedActions

`func (o *RiskAuditVerdictV2) GetAppliedActions() []RiskAuditAppliedActionV2`

GetAppliedActions returns the AppliedActions field if non-nil, zero value otherwise.

### GetAppliedActionsOk

`func (o *RiskAuditVerdictV2) GetAppliedActionsOk() (*[]RiskAuditAppliedActionV2, bool)`

GetAppliedActionsOk returns a tuple with the AppliedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedActions

`func (o *RiskAuditVerdictV2) SetAppliedActions(v []RiskAuditAppliedActionV2)`

SetAppliedActions sets AppliedActions field to given value.

### HasAppliedActions

`func (o *RiskAuditVerdictV2) HasAppliedActions() bool`

HasAppliedActions returns a boolean if a field has been set.

### GetDecisionDetails

`func (o *RiskAuditVerdictV2) GetDecisionDetails() interface{}`

GetDecisionDetails returns the DecisionDetails field if non-nil, zero value otherwise.

### GetDecisionDetailsOk

`func (o *RiskAuditVerdictV2) GetDecisionDetailsOk() (*interface{}, bool)`

GetDecisionDetailsOk returns a tuple with the DecisionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDetails

`func (o *RiskAuditVerdictV2) SetDecisionDetails(v interface{})`

SetDecisionDetails sets DecisionDetails field to given value.

### HasDecisionDetails

`func (o *RiskAuditVerdictV2) HasDecisionDetails() bool`

HasDecisionDetails returns a boolean if a field has been set.

### SetDecisionDetailsNil

`func (o *RiskAuditVerdictV2) SetDecisionDetailsNil(b bool)`

 SetDecisionDetailsNil sets the value for DecisionDetails to be an explicit nil

### UnsetDecisionDetails
`func (o *RiskAuditVerdictV2) UnsetDecisionDetails()`

UnsetDecisionDetails ensures that no value is present for DecisionDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


