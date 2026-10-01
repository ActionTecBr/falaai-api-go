# RiskAuditAppliedActionV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | **string** | Action type code | 
**Label** | **string** | Action label | 
**Description** | **string** | Action description | 
**Priority** | **string** | CRITICO/ALTO/MEDIO/BAIXO | 
**Color** | Pointer to **string** | Color | [optional] [default to ""]
**Icon** | Pointer to **string** | Icon | [optional] [default to ""]
**Condition** | Pointer to **string** | Condition | [optional] 
**Reason** | Pointer to **string** | Reason | [optional] [default to ""]

## Methods

### NewRiskAuditAppliedActionV2

`func NewRiskAuditAppliedActionV2(actionType string, label string, description string, priority string, ) *RiskAuditAppliedActionV2`

NewRiskAuditAppliedActionV2 instantiates a new RiskAuditAppliedActionV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAuditAppliedActionV2WithDefaults

`func NewRiskAuditAppliedActionV2WithDefaults() *RiskAuditAppliedActionV2`

NewRiskAuditAppliedActionV2WithDefaults instantiates a new RiskAuditAppliedActionV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *RiskAuditAppliedActionV2) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *RiskAuditAppliedActionV2) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *RiskAuditAppliedActionV2) SetActionType(v string)`

SetActionType sets ActionType field to given value.


### GetLabel

`func (o *RiskAuditAppliedActionV2) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RiskAuditAppliedActionV2) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RiskAuditAppliedActionV2) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *RiskAuditAppliedActionV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RiskAuditAppliedActionV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RiskAuditAppliedActionV2) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPriority

`func (o *RiskAuditAppliedActionV2) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RiskAuditAppliedActionV2) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RiskAuditAppliedActionV2) SetPriority(v string)`

SetPriority sets Priority field to given value.


### GetColor

`func (o *RiskAuditAppliedActionV2) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *RiskAuditAppliedActionV2) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *RiskAuditAppliedActionV2) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *RiskAuditAppliedActionV2) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetIcon

`func (o *RiskAuditAppliedActionV2) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *RiskAuditAppliedActionV2) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *RiskAuditAppliedActionV2) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *RiskAuditAppliedActionV2) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetCondition

`func (o *RiskAuditAppliedActionV2) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *RiskAuditAppliedActionV2) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *RiskAuditAppliedActionV2) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *RiskAuditAppliedActionV2) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetReason

`func (o *RiskAuditAppliedActionV2) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RiskAuditAppliedActionV2) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RiskAuditAppliedActionV2) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RiskAuditAppliedActionV2) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


