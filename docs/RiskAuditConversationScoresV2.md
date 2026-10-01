# RiskAuditConversationScoresV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsolidatedScore** | Pointer to **float32** | Consolidated score | [optional] 
**ViolationDensityPerMin** | Pointer to **float32** | Violation density/min | [optional] 
**SentimentTrend** | Pointer to **interface{}** | Sentiment trend | [optional] 
**PctTurnsWithViolation** | Pointer to **float32** | % turns with violation | [optional] 
**MostCriticalTurn** | Pointer to **interface{}** | Most critical turn | [optional] 
**PositiveNegativeRatio** | Pointer to **interface{}** | Positive:negative ratio | [optional] 
**GlobalRiskSeverity** | Pointer to **string** | Global risk severity code | [optional] 
**GlobalRiskSeverityLabel** | Pointer to **string** | Global risk severity label | [optional] 
**GlobalRiskSeverityColor** | Pointer to **string** | Global risk severity color | [optional] 
**RiskLikelihoodAvg** | Pointer to **float32** | Risk likelihood avg | [optional] 
**RiskImpactAvg** | Pointer to **float32** | Risk impact avg | [optional] 

## Methods

### NewRiskAuditConversationScoresV2

`func NewRiskAuditConversationScoresV2() *RiskAuditConversationScoresV2`

NewRiskAuditConversationScoresV2 instantiates a new RiskAuditConversationScoresV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAuditConversationScoresV2WithDefaults

`func NewRiskAuditConversationScoresV2WithDefaults() *RiskAuditConversationScoresV2`

NewRiskAuditConversationScoresV2WithDefaults instantiates a new RiskAuditConversationScoresV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsolidatedScore

`func (o *RiskAuditConversationScoresV2) GetConsolidatedScore() float32`

GetConsolidatedScore returns the ConsolidatedScore field if non-nil, zero value otherwise.

### GetConsolidatedScoreOk

`func (o *RiskAuditConversationScoresV2) GetConsolidatedScoreOk() (*float32, bool)`

GetConsolidatedScoreOk returns a tuple with the ConsolidatedScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolidatedScore

`func (o *RiskAuditConversationScoresV2) SetConsolidatedScore(v float32)`

SetConsolidatedScore sets ConsolidatedScore field to given value.

### HasConsolidatedScore

`func (o *RiskAuditConversationScoresV2) HasConsolidatedScore() bool`

HasConsolidatedScore returns a boolean if a field has been set.

### GetViolationDensityPerMin

`func (o *RiskAuditConversationScoresV2) GetViolationDensityPerMin() float32`

GetViolationDensityPerMin returns the ViolationDensityPerMin field if non-nil, zero value otherwise.

### GetViolationDensityPerMinOk

`func (o *RiskAuditConversationScoresV2) GetViolationDensityPerMinOk() (*float32, bool)`

GetViolationDensityPerMinOk returns a tuple with the ViolationDensityPerMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationDensityPerMin

`func (o *RiskAuditConversationScoresV2) SetViolationDensityPerMin(v float32)`

SetViolationDensityPerMin sets ViolationDensityPerMin field to given value.

### HasViolationDensityPerMin

`func (o *RiskAuditConversationScoresV2) HasViolationDensityPerMin() bool`

HasViolationDensityPerMin returns a boolean if a field has been set.

### GetSentimentTrend

`func (o *RiskAuditConversationScoresV2) GetSentimentTrend() interface{}`

GetSentimentTrend returns the SentimentTrend field if non-nil, zero value otherwise.

### GetSentimentTrendOk

`func (o *RiskAuditConversationScoresV2) GetSentimentTrendOk() (*interface{}, bool)`

GetSentimentTrendOk returns a tuple with the SentimentTrend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentimentTrend

`func (o *RiskAuditConversationScoresV2) SetSentimentTrend(v interface{})`

SetSentimentTrend sets SentimentTrend field to given value.

### HasSentimentTrend

`func (o *RiskAuditConversationScoresV2) HasSentimentTrend() bool`

HasSentimentTrend returns a boolean if a field has been set.

### SetSentimentTrendNil

`func (o *RiskAuditConversationScoresV2) SetSentimentTrendNil(b bool)`

 SetSentimentTrendNil sets the value for SentimentTrend to be an explicit nil

### UnsetSentimentTrend
`func (o *RiskAuditConversationScoresV2) UnsetSentimentTrend()`

UnsetSentimentTrend ensures that no value is present for SentimentTrend, not even an explicit nil
### GetPctTurnsWithViolation

`func (o *RiskAuditConversationScoresV2) GetPctTurnsWithViolation() float32`

GetPctTurnsWithViolation returns the PctTurnsWithViolation field if non-nil, zero value otherwise.

### GetPctTurnsWithViolationOk

`func (o *RiskAuditConversationScoresV2) GetPctTurnsWithViolationOk() (*float32, bool)`

GetPctTurnsWithViolationOk returns a tuple with the PctTurnsWithViolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPctTurnsWithViolation

`func (o *RiskAuditConversationScoresV2) SetPctTurnsWithViolation(v float32)`

SetPctTurnsWithViolation sets PctTurnsWithViolation field to given value.

### HasPctTurnsWithViolation

`func (o *RiskAuditConversationScoresV2) HasPctTurnsWithViolation() bool`

HasPctTurnsWithViolation returns a boolean if a field has been set.

### GetMostCriticalTurn

`func (o *RiskAuditConversationScoresV2) GetMostCriticalTurn() interface{}`

GetMostCriticalTurn returns the MostCriticalTurn field if non-nil, zero value otherwise.

### GetMostCriticalTurnOk

`func (o *RiskAuditConversationScoresV2) GetMostCriticalTurnOk() (*interface{}, bool)`

GetMostCriticalTurnOk returns a tuple with the MostCriticalTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostCriticalTurn

`func (o *RiskAuditConversationScoresV2) SetMostCriticalTurn(v interface{})`

SetMostCriticalTurn sets MostCriticalTurn field to given value.

### HasMostCriticalTurn

`func (o *RiskAuditConversationScoresV2) HasMostCriticalTurn() bool`

HasMostCriticalTurn returns a boolean if a field has been set.

### SetMostCriticalTurnNil

`func (o *RiskAuditConversationScoresV2) SetMostCriticalTurnNil(b bool)`

 SetMostCriticalTurnNil sets the value for MostCriticalTurn to be an explicit nil

### UnsetMostCriticalTurn
`func (o *RiskAuditConversationScoresV2) UnsetMostCriticalTurn()`

UnsetMostCriticalTurn ensures that no value is present for MostCriticalTurn, not even an explicit nil
### GetPositiveNegativeRatio

`func (o *RiskAuditConversationScoresV2) GetPositiveNegativeRatio() interface{}`

GetPositiveNegativeRatio returns the PositiveNegativeRatio field if non-nil, zero value otherwise.

### GetPositiveNegativeRatioOk

`func (o *RiskAuditConversationScoresV2) GetPositiveNegativeRatioOk() (*interface{}, bool)`

GetPositiveNegativeRatioOk returns a tuple with the PositiveNegativeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositiveNegativeRatio

`func (o *RiskAuditConversationScoresV2) SetPositiveNegativeRatio(v interface{})`

SetPositiveNegativeRatio sets PositiveNegativeRatio field to given value.

### HasPositiveNegativeRatio

`func (o *RiskAuditConversationScoresV2) HasPositiveNegativeRatio() bool`

HasPositiveNegativeRatio returns a boolean if a field has been set.

### SetPositiveNegativeRatioNil

`func (o *RiskAuditConversationScoresV2) SetPositiveNegativeRatioNil(b bool)`

 SetPositiveNegativeRatioNil sets the value for PositiveNegativeRatio to be an explicit nil

### UnsetPositiveNegativeRatio
`func (o *RiskAuditConversationScoresV2) UnsetPositiveNegativeRatio()`

UnsetPositiveNegativeRatio ensures that no value is present for PositiveNegativeRatio, not even an explicit nil
### GetGlobalRiskSeverity

`func (o *RiskAuditConversationScoresV2) GetGlobalRiskSeverity() string`

GetGlobalRiskSeverity returns the GlobalRiskSeverity field if non-nil, zero value otherwise.

### GetGlobalRiskSeverityOk

`func (o *RiskAuditConversationScoresV2) GetGlobalRiskSeverityOk() (*string, bool)`

GetGlobalRiskSeverityOk returns a tuple with the GlobalRiskSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalRiskSeverity

`func (o *RiskAuditConversationScoresV2) SetGlobalRiskSeverity(v string)`

SetGlobalRiskSeverity sets GlobalRiskSeverity field to given value.

### HasGlobalRiskSeverity

`func (o *RiskAuditConversationScoresV2) HasGlobalRiskSeverity() bool`

HasGlobalRiskSeverity returns a boolean if a field has been set.

### GetGlobalRiskSeverityLabel

`func (o *RiskAuditConversationScoresV2) GetGlobalRiskSeverityLabel() string`

GetGlobalRiskSeverityLabel returns the GlobalRiskSeverityLabel field if non-nil, zero value otherwise.

### GetGlobalRiskSeverityLabelOk

`func (o *RiskAuditConversationScoresV2) GetGlobalRiskSeverityLabelOk() (*string, bool)`

GetGlobalRiskSeverityLabelOk returns a tuple with the GlobalRiskSeverityLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalRiskSeverityLabel

`func (o *RiskAuditConversationScoresV2) SetGlobalRiskSeverityLabel(v string)`

SetGlobalRiskSeverityLabel sets GlobalRiskSeverityLabel field to given value.

### HasGlobalRiskSeverityLabel

`func (o *RiskAuditConversationScoresV2) HasGlobalRiskSeverityLabel() bool`

HasGlobalRiskSeverityLabel returns a boolean if a field has been set.

### GetGlobalRiskSeverityColor

`func (o *RiskAuditConversationScoresV2) GetGlobalRiskSeverityColor() string`

GetGlobalRiskSeverityColor returns the GlobalRiskSeverityColor field if non-nil, zero value otherwise.

### GetGlobalRiskSeverityColorOk

`func (o *RiskAuditConversationScoresV2) GetGlobalRiskSeverityColorOk() (*string, bool)`

GetGlobalRiskSeverityColorOk returns a tuple with the GlobalRiskSeverityColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalRiskSeverityColor

`func (o *RiskAuditConversationScoresV2) SetGlobalRiskSeverityColor(v string)`

SetGlobalRiskSeverityColor sets GlobalRiskSeverityColor field to given value.

### HasGlobalRiskSeverityColor

`func (o *RiskAuditConversationScoresV2) HasGlobalRiskSeverityColor() bool`

HasGlobalRiskSeverityColor returns a boolean if a field has been set.

### GetRiskLikelihoodAvg

`func (o *RiskAuditConversationScoresV2) GetRiskLikelihoodAvg() float32`

GetRiskLikelihoodAvg returns the RiskLikelihoodAvg field if non-nil, zero value otherwise.

### GetRiskLikelihoodAvgOk

`func (o *RiskAuditConversationScoresV2) GetRiskLikelihoodAvgOk() (*float32, bool)`

GetRiskLikelihoodAvgOk returns a tuple with the RiskLikelihoodAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLikelihoodAvg

`func (o *RiskAuditConversationScoresV2) SetRiskLikelihoodAvg(v float32)`

SetRiskLikelihoodAvg sets RiskLikelihoodAvg field to given value.

### HasRiskLikelihoodAvg

`func (o *RiskAuditConversationScoresV2) HasRiskLikelihoodAvg() bool`

HasRiskLikelihoodAvg returns a boolean if a field has been set.

### GetRiskImpactAvg

`func (o *RiskAuditConversationScoresV2) GetRiskImpactAvg() float32`

GetRiskImpactAvg returns the RiskImpactAvg field if non-nil, zero value otherwise.

### GetRiskImpactAvgOk

`func (o *RiskAuditConversationScoresV2) GetRiskImpactAvgOk() (*float32, bool)`

GetRiskImpactAvgOk returns a tuple with the RiskImpactAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskImpactAvg

`func (o *RiskAuditConversationScoresV2) SetRiskImpactAvg(v float32)`

SetRiskImpactAvg sets RiskImpactAvg field to given value.

### HasRiskImpactAvg

`func (o *RiskAuditConversationScoresV2) HasRiskImpactAvg() bool`

HasRiskImpactAvg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


