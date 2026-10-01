# ParticipantDiagnostic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interlocutor** | **string** | Exact speaker label from the dialog (e.g. &#39;Speaker 1&#39;) | 
**Role** | **string** | Role: agent | client | bot | agent_requester | agent_custodian | 
**Name** | Pointer to **string** | Participant name if mentioned in the dialogue | [optional] 
**Confidence** | Pointer to **string** | high | medium | low | [optional] 
**Evidence** | Pointer to **string** | Exact verbatim quote supporting the role (no timestamps) | [optional] 

## Methods

### NewParticipantDiagnostic

`func NewParticipantDiagnostic(interlocutor string, role string, ) *ParticipantDiagnostic`

NewParticipantDiagnostic instantiates a new ParticipantDiagnostic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantDiagnosticWithDefaults

`func NewParticipantDiagnosticWithDefaults() *ParticipantDiagnostic`

NewParticipantDiagnosticWithDefaults instantiates a new ParticipantDiagnostic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterlocutor

`func (o *ParticipantDiagnostic) GetInterlocutor() string`

GetInterlocutor returns the Interlocutor field if non-nil, zero value otherwise.

### GetInterlocutorOk

`func (o *ParticipantDiagnostic) GetInterlocutorOk() (*string, bool)`

GetInterlocutorOk returns a tuple with the Interlocutor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterlocutor

`func (o *ParticipantDiagnostic) SetInterlocutor(v string)`

SetInterlocutor sets Interlocutor field to given value.


### GetRole

`func (o *ParticipantDiagnostic) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ParticipantDiagnostic) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ParticipantDiagnostic) SetRole(v string)`

SetRole sets Role field to given value.


### GetName

`func (o *ParticipantDiagnostic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParticipantDiagnostic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParticipantDiagnostic) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ParticipantDiagnostic) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfidence

`func (o *ParticipantDiagnostic) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *ParticipantDiagnostic) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *ParticipantDiagnostic) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *ParticipantDiagnostic) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetEvidence

`func (o *ParticipantDiagnostic) GetEvidence() string`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *ParticipantDiagnostic) GetEvidenceOk() (*string, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *ParticipantDiagnostic) SetEvidence(v string)`

SetEvidence sets Evidence field to given value.

### HasEvidence

`func (o *ParticipantDiagnostic) HasEvidence() bool`

HasEvidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


