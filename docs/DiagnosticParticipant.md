# DiagnosticParticipant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interlocutor** | **string** | Exact speaker label from the dialog (e.g. &#39;Speaker 1&#39;) | 
**Name** | Pointer to **string** | Participant name (optional) | [optional] 
**Role** | **string** | agent | client | bot | 

## Methods

### NewDiagnosticParticipant

`func NewDiagnosticParticipant(interlocutor string, role string, ) *DiagnosticParticipant`

NewDiagnosticParticipant instantiates a new DiagnosticParticipant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosticParticipantWithDefaults

`func NewDiagnosticParticipantWithDefaults() *DiagnosticParticipant`

NewDiagnosticParticipantWithDefaults instantiates a new DiagnosticParticipant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterlocutor

`func (o *DiagnosticParticipant) GetInterlocutor() string`

GetInterlocutor returns the Interlocutor field if non-nil, zero value otherwise.

### GetInterlocutorOk

`func (o *DiagnosticParticipant) GetInterlocutorOk() (*string, bool)`

GetInterlocutorOk returns a tuple with the Interlocutor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterlocutor

`func (o *DiagnosticParticipant) SetInterlocutor(v string)`

SetInterlocutor sets Interlocutor field to given value.


### GetName

`func (o *DiagnosticParticipant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiagnosticParticipant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiagnosticParticipant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DiagnosticParticipant) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRole

`func (o *DiagnosticParticipant) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *DiagnosticParticipant) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *DiagnosticParticipant) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


