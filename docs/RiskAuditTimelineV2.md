# RiskAuditTimelineV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TurnsSentiment** | Pointer to **[]map[string]interface{}** | Per-turn sentiment | [optional] [default to {}]
**AudioEvents** | Pointer to **[]map[string]interface{}** | Audio events (i18n) | [optional] [default to {}]
**AudioGroupsFound** | Pointer to **[]map[string]interface{}** | Audio groups found | [optional] [default to {}]

## Methods

### NewRiskAuditTimelineV2

`func NewRiskAuditTimelineV2() *RiskAuditTimelineV2`

NewRiskAuditTimelineV2 instantiates a new RiskAuditTimelineV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAuditTimelineV2WithDefaults

`func NewRiskAuditTimelineV2WithDefaults() *RiskAuditTimelineV2`

NewRiskAuditTimelineV2WithDefaults instantiates a new RiskAuditTimelineV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTurnsSentiment

`func (o *RiskAuditTimelineV2) GetTurnsSentiment() []map[string]interface{}`

GetTurnsSentiment returns the TurnsSentiment field if non-nil, zero value otherwise.

### GetTurnsSentimentOk

`func (o *RiskAuditTimelineV2) GetTurnsSentimentOk() (*[]map[string]interface{}, bool)`

GetTurnsSentimentOk returns a tuple with the TurnsSentiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurnsSentiment

`func (o *RiskAuditTimelineV2) SetTurnsSentiment(v []map[string]interface{})`

SetTurnsSentiment sets TurnsSentiment field to given value.

### HasTurnsSentiment

`func (o *RiskAuditTimelineV2) HasTurnsSentiment() bool`

HasTurnsSentiment returns a boolean if a field has been set.

### GetAudioEvents

`func (o *RiskAuditTimelineV2) GetAudioEvents() []map[string]interface{}`

GetAudioEvents returns the AudioEvents field if non-nil, zero value otherwise.

### GetAudioEventsOk

`func (o *RiskAuditTimelineV2) GetAudioEventsOk() (*[]map[string]interface{}, bool)`

GetAudioEventsOk returns a tuple with the AudioEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioEvents

`func (o *RiskAuditTimelineV2) SetAudioEvents(v []map[string]interface{})`

SetAudioEvents sets AudioEvents field to given value.

### HasAudioEvents

`func (o *RiskAuditTimelineV2) HasAudioEvents() bool`

HasAudioEvents returns a boolean if a field has been set.

### GetAudioGroupsFound

`func (o *RiskAuditTimelineV2) GetAudioGroupsFound() []map[string]interface{}`

GetAudioGroupsFound returns the AudioGroupsFound field if non-nil, zero value otherwise.

### GetAudioGroupsFoundOk

`func (o *RiskAuditTimelineV2) GetAudioGroupsFoundOk() (*[]map[string]interface{}, bool)`

GetAudioGroupsFoundOk returns a tuple with the AudioGroupsFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioGroupsFound

`func (o *RiskAuditTimelineV2) SetAudioGroupsFound(v []map[string]interface{})`

SetAudioGroupsFound sets AudioGroupsFound field to given value.

### HasAudioGroupsFound

`func (o *RiskAuditTimelineV2) HasAudioGroupsFound() bool`

HasAudioGroupsFound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


