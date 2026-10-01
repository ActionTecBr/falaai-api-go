# RiskAuditParticipantV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interlocutor** | Pointer to **string** | Speaker label | [optional] 
**Name** | Pointer to **string** | Participant name | [optional] 
**Role** | Pointer to **string** | Role (agent/client/bot/unknown) | [optional] 
**Confidence** | Pointer to **string** | Role inference confidence (high/medium/low) | [optional] 
**Source** | Pointer to **string** | Role source (input/inferred) | [optional] 
**Evidence** | Pointer to **string** | Role inference evidence | [optional] 

## Methods

### NewRiskAuditParticipantV2

`func NewRiskAuditParticipantV2() *RiskAuditParticipantV2`

NewRiskAuditParticipantV2 instantiates a new RiskAuditParticipantV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAuditParticipantV2WithDefaults

`func NewRiskAuditParticipantV2WithDefaults() *RiskAuditParticipantV2`

NewRiskAuditParticipantV2WithDefaults instantiates a new RiskAuditParticipantV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterlocutor

`func (o *RiskAuditParticipantV2) GetInterlocutor() string`

GetInterlocutor returns the Interlocutor field if non-nil, zero value otherwise.

### GetInterlocutorOk

`func (o *RiskAuditParticipantV2) GetInterlocutorOk() (*string, bool)`

GetInterlocutorOk returns a tuple with the Interlocutor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterlocutor

`func (o *RiskAuditParticipantV2) SetInterlocutor(v string)`

SetInterlocutor sets Interlocutor field to given value.

### HasInterlocutor

`func (o *RiskAuditParticipantV2) HasInterlocutor() bool`

HasInterlocutor returns a boolean if a field has been set.

### GetName

`func (o *RiskAuditParticipantV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskAuditParticipantV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskAuditParticipantV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RiskAuditParticipantV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRole

`func (o *RiskAuditParticipantV2) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RiskAuditParticipantV2) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RiskAuditParticipantV2) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *RiskAuditParticipantV2) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetConfidence

`func (o *RiskAuditParticipantV2) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *RiskAuditParticipantV2) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *RiskAuditParticipantV2) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *RiskAuditParticipantV2) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetSource

`func (o *RiskAuditParticipantV2) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RiskAuditParticipantV2) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RiskAuditParticipantV2) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RiskAuditParticipantV2) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetEvidence

`func (o *RiskAuditParticipantV2) GetEvidence() string`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *RiskAuditParticipantV2) GetEvidenceOk() (*string, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *RiskAuditParticipantV2) SetEvidence(v string)`

SetEvidence sets Evidence field to given value.

### HasEvidence

`func (o *RiskAuditParticipantV2) HasEvidence() bool`

HasEvidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


