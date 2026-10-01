# RiskAuditSummaryV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalTurns** | Pointer to **int32** | Total turns | [optional] 
**TotalCalibrated** | Pointer to **int32** | Total calibrated detections | [optional] 
**Active** | Pointer to **int32** | Active detections | [optional] 
**Tolerated** | Pointer to **int32** | Tolerated detections | [optional] 
**Blocked** | Pointer to **int32** | Blocked detections | [optional] 
**AudioEventsUsed** | Pointer to **int32** | Audio events used | [optional] 
**AudioEventsAggravated** | Pointer to **int32** | Audio events aggravated | [optional] 
**MacAudioApplied** | Pointer to **interface{}** | MAC audio applied | [optional] 
**MvadApplied** | Pointer to **interface{}** | MVAD applied | [optional] 
**TotalParticipants** | Pointer to **int32** | Total participants | [optional] 
**TotalAgents** | Pointer to **int32** | Total agents | [optional] 
**TotalClients** | Pointer to **int32** | Total clients | [optional] 
**TotalBots** | Pointer to **int32** | Total bots | [optional] 
**TotalUnknown** | Pointer to **int32** | Total unknown | [optional] 
**ClientRiskAlertsCount** | Pointer to **int32** | Client risk alerts count | [optional] 
**ClientBehaviorAlertsCount** | Pointer to **int32** | Client behavior alerts count | [optional] 

## Methods

### NewRiskAuditSummaryV2

`func NewRiskAuditSummaryV2() *RiskAuditSummaryV2`

NewRiskAuditSummaryV2 instantiates a new RiskAuditSummaryV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAuditSummaryV2WithDefaults

`func NewRiskAuditSummaryV2WithDefaults() *RiskAuditSummaryV2`

NewRiskAuditSummaryV2WithDefaults instantiates a new RiskAuditSummaryV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalTurns

`func (o *RiskAuditSummaryV2) GetTotalTurns() int32`

GetTotalTurns returns the TotalTurns field if non-nil, zero value otherwise.

### GetTotalTurnsOk

`func (o *RiskAuditSummaryV2) GetTotalTurnsOk() (*int32, bool)`

GetTotalTurnsOk returns a tuple with the TotalTurns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTurns

`func (o *RiskAuditSummaryV2) SetTotalTurns(v int32)`

SetTotalTurns sets TotalTurns field to given value.

### HasTotalTurns

`func (o *RiskAuditSummaryV2) HasTotalTurns() bool`

HasTotalTurns returns a boolean if a field has been set.

### GetTotalCalibrated

`func (o *RiskAuditSummaryV2) GetTotalCalibrated() int32`

GetTotalCalibrated returns the TotalCalibrated field if non-nil, zero value otherwise.

### GetTotalCalibratedOk

`func (o *RiskAuditSummaryV2) GetTotalCalibratedOk() (*int32, bool)`

GetTotalCalibratedOk returns a tuple with the TotalCalibrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCalibrated

`func (o *RiskAuditSummaryV2) SetTotalCalibrated(v int32)`

SetTotalCalibrated sets TotalCalibrated field to given value.

### HasTotalCalibrated

`func (o *RiskAuditSummaryV2) HasTotalCalibrated() bool`

HasTotalCalibrated returns a boolean if a field has been set.

### GetActive

`func (o *RiskAuditSummaryV2) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RiskAuditSummaryV2) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RiskAuditSummaryV2) SetActive(v int32)`

SetActive sets Active field to given value.

### HasActive

`func (o *RiskAuditSummaryV2) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetTolerated

`func (o *RiskAuditSummaryV2) GetTolerated() int32`

GetTolerated returns the Tolerated field if non-nil, zero value otherwise.

### GetToleratedOk

`func (o *RiskAuditSummaryV2) GetToleratedOk() (*int32, bool)`

GetToleratedOk returns a tuple with the Tolerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTolerated

`func (o *RiskAuditSummaryV2) SetTolerated(v int32)`

SetTolerated sets Tolerated field to given value.

### HasTolerated

`func (o *RiskAuditSummaryV2) HasTolerated() bool`

HasTolerated returns a boolean if a field has been set.

### GetBlocked

`func (o *RiskAuditSummaryV2) GetBlocked() int32`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *RiskAuditSummaryV2) GetBlockedOk() (*int32, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *RiskAuditSummaryV2) SetBlocked(v int32)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *RiskAuditSummaryV2) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetAudioEventsUsed

`func (o *RiskAuditSummaryV2) GetAudioEventsUsed() int32`

GetAudioEventsUsed returns the AudioEventsUsed field if non-nil, zero value otherwise.

### GetAudioEventsUsedOk

`func (o *RiskAuditSummaryV2) GetAudioEventsUsedOk() (*int32, bool)`

GetAudioEventsUsedOk returns a tuple with the AudioEventsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioEventsUsed

`func (o *RiskAuditSummaryV2) SetAudioEventsUsed(v int32)`

SetAudioEventsUsed sets AudioEventsUsed field to given value.

### HasAudioEventsUsed

`func (o *RiskAuditSummaryV2) HasAudioEventsUsed() bool`

HasAudioEventsUsed returns a boolean if a field has been set.

### GetAudioEventsAggravated

`func (o *RiskAuditSummaryV2) GetAudioEventsAggravated() int32`

GetAudioEventsAggravated returns the AudioEventsAggravated field if non-nil, zero value otherwise.

### GetAudioEventsAggravatedOk

`func (o *RiskAuditSummaryV2) GetAudioEventsAggravatedOk() (*int32, bool)`

GetAudioEventsAggravatedOk returns a tuple with the AudioEventsAggravated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioEventsAggravated

`func (o *RiskAuditSummaryV2) SetAudioEventsAggravated(v int32)`

SetAudioEventsAggravated sets AudioEventsAggravated field to given value.

### HasAudioEventsAggravated

`func (o *RiskAuditSummaryV2) HasAudioEventsAggravated() bool`

HasAudioEventsAggravated returns a boolean if a field has been set.

### GetMacAudioApplied

`func (o *RiskAuditSummaryV2) GetMacAudioApplied() interface{}`

GetMacAudioApplied returns the MacAudioApplied field if non-nil, zero value otherwise.

### GetMacAudioAppliedOk

`func (o *RiskAuditSummaryV2) GetMacAudioAppliedOk() (*interface{}, bool)`

GetMacAudioAppliedOk returns a tuple with the MacAudioApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAudioApplied

`func (o *RiskAuditSummaryV2) SetMacAudioApplied(v interface{})`

SetMacAudioApplied sets MacAudioApplied field to given value.

### HasMacAudioApplied

`func (o *RiskAuditSummaryV2) HasMacAudioApplied() bool`

HasMacAudioApplied returns a boolean if a field has been set.

### SetMacAudioAppliedNil

`func (o *RiskAuditSummaryV2) SetMacAudioAppliedNil(b bool)`

 SetMacAudioAppliedNil sets the value for MacAudioApplied to be an explicit nil

### UnsetMacAudioApplied
`func (o *RiskAuditSummaryV2) UnsetMacAudioApplied()`

UnsetMacAudioApplied ensures that no value is present for MacAudioApplied, not even an explicit nil
### GetMvadApplied

`func (o *RiskAuditSummaryV2) GetMvadApplied() interface{}`

GetMvadApplied returns the MvadApplied field if non-nil, zero value otherwise.

### GetMvadAppliedOk

`func (o *RiskAuditSummaryV2) GetMvadAppliedOk() (*interface{}, bool)`

GetMvadAppliedOk returns a tuple with the MvadApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMvadApplied

`func (o *RiskAuditSummaryV2) SetMvadApplied(v interface{})`

SetMvadApplied sets MvadApplied field to given value.

### HasMvadApplied

`func (o *RiskAuditSummaryV2) HasMvadApplied() bool`

HasMvadApplied returns a boolean if a field has been set.

### SetMvadAppliedNil

`func (o *RiskAuditSummaryV2) SetMvadAppliedNil(b bool)`

 SetMvadAppliedNil sets the value for MvadApplied to be an explicit nil

### UnsetMvadApplied
`func (o *RiskAuditSummaryV2) UnsetMvadApplied()`

UnsetMvadApplied ensures that no value is present for MvadApplied, not even an explicit nil
### GetTotalParticipants

`func (o *RiskAuditSummaryV2) GetTotalParticipants() int32`

GetTotalParticipants returns the TotalParticipants field if non-nil, zero value otherwise.

### GetTotalParticipantsOk

`func (o *RiskAuditSummaryV2) GetTotalParticipantsOk() (*int32, bool)`

GetTotalParticipantsOk returns a tuple with the TotalParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalParticipants

`func (o *RiskAuditSummaryV2) SetTotalParticipants(v int32)`

SetTotalParticipants sets TotalParticipants field to given value.

### HasTotalParticipants

`func (o *RiskAuditSummaryV2) HasTotalParticipants() bool`

HasTotalParticipants returns a boolean if a field has been set.

### GetTotalAgents

`func (o *RiskAuditSummaryV2) GetTotalAgents() int32`

GetTotalAgents returns the TotalAgents field if non-nil, zero value otherwise.

### GetTotalAgentsOk

`func (o *RiskAuditSummaryV2) GetTotalAgentsOk() (*int32, bool)`

GetTotalAgentsOk returns a tuple with the TotalAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAgents

`func (o *RiskAuditSummaryV2) SetTotalAgents(v int32)`

SetTotalAgents sets TotalAgents field to given value.

### HasTotalAgents

`func (o *RiskAuditSummaryV2) HasTotalAgents() bool`

HasTotalAgents returns a boolean if a field has been set.

### GetTotalClients

`func (o *RiskAuditSummaryV2) GetTotalClients() int32`

GetTotalClients returns the TotalClients field if non-nil, zero value otherwise.

### GetTotalClientsOk

`func (o *RiskAuditSummaryV2) GetTotalClientsOk() (*int32, bool)`

GetTotalClientsOk returns a tuple with the TotalClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalClients

`func (o *RiskAuditSummaryV2) SetTotalClients(v int32)`

SetTotalClients sets TotalClients field to given value.

### HasTotalClients

`func (o *RiskAuditSummaryV2) HasTotalClients() bool`

HasTotalClients returns a boolean if a field has been set.

### GetTotalBots

`func (o *RiskAuditSummaryV2) GetTotalBots() int32`

GetTotalBots returns the TotalBots field if non-nil, zero value otherwise.

### GetTotalBotsOk

`func (o *RiskAuditSummaryV2) GetTotalBotsOk() (*int32, bool)`

GetTotalBotsOk returns a tuple with the TotalBots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBots

`func (o *RiskAuditSummaryV2) SetTotalBots(v int32)`

SetTotalBots sets TotalBots field to given value.

### HasTotalBots

`func (o *RiskAuditSummaryV2) HasTotalBots() bool`

HasTotalBots returns a boolean if a field has been set.

### GetTotalUnknown

`func (o *RiskAuditSummaryV2) GetTotalUnknown() int32`

GetTotalUnknown returns the TotalUnknown field if non-nil, zero value otherwise.

### GetTotalUnknownOk

`func (o *RiskAuditSummaryV2) GetTotalUnknownOk() (*int32, bool)`

GetTotalUnknownOk returns a tuple with the TotalUnknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnknown

`func (o *RiskAuditSummaryV2) SetTotalUnknown(v int32)`

SetTotalUnknown sets TotalUnknown field to given value.

### HasTotalUnknown

`func (o *RiskAuditSummaryV2) HasTotalUnknown() bool`

HasTotalUnknown returns a boolean if a field has been set.

### GetClientRiskAlertsCount

`func (o *RiskAuditSummaryV2) GetClientRiskAlertsCount() int32`

GetClientRiskAlertsCount returns the ClientRiskAlertsCount field if non-nil, zero value otherwise.

### GetClientRiskAlertsCountOk

`func (o *RiskAuditSummaryV2) GetClientRiskAlertsCountOk() (*int32, bool)`

GetClientRiskAlertsCountOk returns a tuple with the ClientRiskAlertsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRiskAlertsCount

`func (o *RiskAuditSummaryV2) SetClientRiskAlertsCount(v int32)`

SetClientRiskAlertsCount sets ClientRiskAlertsCount field to given value.

### HasClientRiskAlertsCount

`func (o *RiskAuditSummaryV2) HasClientRiskAlertsCount() bool`

HasClientRiskAlertsCount returns a boolean if a field has been set.

### GetClientBehaviorAlertsCount

`func (o *RiskAuditSummaryV2) GetClientBehaviorAlertsCount() int32`

GetClientBehaviorAlertsCount returns the ClientBehaviorAlertsCount field if non-nil, zero value otherwise.

### GetClientBehaviorAlertsCountOk

`func (o *RiskAuditSummaryV2) GetClientBehaviorAlertsCountOk() (*int32, bool)`

GetClientBehaviorAlertsCountOk returns a tuple with the ClientBehaviorAlertsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientBehaviorAlertsCount

`func (o *RiskAuditSummaryV2) SetClientBehaviorAlertsCount(v int32)`

SetClientBehaviorAlertsCount sets ClientBehaviorAlertsCount field to given value.

### HasClientBehaviorAlertsCount

`func (o *RiskAuditSummaryV2) HasClientBehaviorAlertsCount() bool`

HasClientBehaviorAlertsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


