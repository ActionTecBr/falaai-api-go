# RiskAuditScoresV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conversation** | [**RiskAuditConversationScoresV2**](RiskAuditConversationScoresV2.md) | Conversation scores | 
**PerParticipant** | Pointer to **map[string]interface{}** | Per-participant KPIs | [optional] [default to {}]

## Methods

### NewRiskAuditScoresV2

`func NewRiskAuditScoresV2(conversation RiskAuditConversationScoresV2, ) *RiskAuditScoresV2`

NewRiskAuditScoresV2 instantiates a new RiskAuditScoresV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAuditScoresV2WithDefaults

`func NewRiskAuditScoresV2WithDefaults() *RiskAuditScoresV2`

NewRiskAuditScoresV2WithDefaults instantiates a new RiskAuditScoresV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversation

`func (o *RiskAuditScoresV2) GetConversation() RiskAuditConversationScoresV2`

GetConversation returns the Conversation field if non-nil, zero value otherwise.

### GetConversationOk

`func (o *RiskAuditScoresV2) GetConversationOk() (*RiskAuditConversationScoresV2, bool)`

GetConversationOk returns a tuple with the Conversation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversation

`func (o *RiskAuditScoresV2) SetConversation(v RiskAuditConversationScoresV2)`

SetConversation sets Conversation field to given value.


### GetPerParticipant

`func (o *RiskAuditScoresV2) GetPerParticipant() map[string]interface{}`

GetPerParticipant returns the PerParticipant field if non-nil, zero value otherwise.

### GetPerParticipantOk

`func (o *RiskAuditScoresV2) GetPerParticipantOk() (*map[string]interface{}, bool)`

GetPerParticipantOk returns a tuple with the PerParticipant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerParticipant

`func (o *RiskAuditScoresV2) SetPerParticipant(v map[string]interface{})`

SetPerParticipant sets PerParticipant field to given value.

### HasPerParticipant

`func (o *RiskAuditScoresV2) HasPerParticipant() bool`

HasPerParticipant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


