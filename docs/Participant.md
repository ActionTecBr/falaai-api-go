# Participant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interlocutor** | **string** | Exact identifier as used in dialog (e.g. &#39;Interlocutor 1&#39;, &#39;Antonio&#39;) | 
**Name** | Pointer to **string** | Participant name (humanizes report, does not affect logic) | [optional] 
**Role** | **string** | Role: agent (human operator), client (customer), bot (IVR/AI) | 

## Methods

### NewParticipant

`func NewParticipant(interlocutor string, role string, ) *Participant`

NewParticipant instantiates a new Participant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantWithDefaults

`func NewParticipantWithDefaults() *Participant`

NewParticipantWithDefaults instantiates a new Participant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterlocutor

`func (o *Participant) GetInterlocutor() string`

GetInterlocutor returns the Interlocutor field if non-nil, zero value otherwise.

### GetInterlocutorOk

`func (o *Participant) GetInterlocutorOk() (*string, bool)`

GetInterlocutorOk returns a tuple with the Interlocutor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterlocutor

`func (o *Participant) SetInterlocutor(v string)`

SetInterlocutor sets Interlocutor field to given value.


### GetName

`func (o *Participant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Participant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Participant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Participant) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRole

`func (o *Participant) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Participant) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Participant) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


