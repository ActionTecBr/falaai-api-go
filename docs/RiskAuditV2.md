# RiskAuditV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**RiskAuditMetaV2**](RiskAuditMetaV2.md) | Identification + usage | 
**Participants** | [**RiskAuditParticipantsV2**](RiskAuditParticipantsV2.md) | Participants/roles/direction | 
**Verdict** | [**RiskAuditVerdictV2**](RiskAuditVerdictV2.md) | Verdict + level + applied actions | 
**Scores** | [**RiskAuditScoresV2**](RiskAuditScoresV2.md) | Consolidated + per-participant scores | 
**Detections** | [**RiskAuditDetectionsV2**](RiskAuditDetectionsV2.md) | violations/positives/client alerts | 
**Analysis** | [**RiskAuditAnalysisV2**](RiskAuditAnalysisV2.md) | global_metrics + final_analysis + frameworks | 
**Timeline** | [**RiskAuditTimelineV2**](RiskAuditTimelineV2.md) | turns_sentiment + audio_events + groups | 
**AudioEventModel** | [**RiskAuditAudioEventModelV2**](RiskAuditAudioEventModelV2.md) | MAC audio event semantics | 
**CategoriesSummary** | **map[string]interface{}** | Per-category summary (keyed by category) | 
**Indexer** | [**RiskAuditIndexerV2**](RiskAuditIndexerV2.md) | Suggested terms for bank | 
**Summary** | [**RiskAuditSummaryV2**](RiskAuditSummaryV2.md) | Executive summary counts | 
**ActionsI18n** | **map[string]interface{}** | Used actions i18n catalog (keyed by action) | 
**AuditDecisions** | [**RiskAuditAuditDecisionsV2**](RiskAuditAuditDecisionsV2.md) | Risk origin + validator changes | 
**ScoringExplanation** | [**RiskAuditScoringExplanationV2**](RiskAuditScoringExplanationV2.md) | Score composition explanation | 
**HtmlReport** | **string** | HTML report (base64 gzip) | 

## Methods

### NewRiskAuditV2

`func NewRiskAuditV2(meta RiskAuditMetaV2, participants RiskAuditParticipantsV2, verdict RiskAuditVerdictV2, scores RiskAuditScoresV2, detections RiskAuditDetectionsV2, analysis RiskAuditAnalysisV2, timeline RiskAuditTimelineV2, audioEventModel RiskAuditAudioEventModelV2, categoriesSummary map[string]interface{}, indexer RiskAuditIndexerV2, summary RiskAuditSummaryV2, actionsI18n map[string]interface{}, auditDecisions RiskAuditAuditDecisionsV2, scoringExplanation RiskAuditScoringExplanationV2, htmlReport string, ) *RiskAuditV2`

NewRiskAuditV2 instantiates a new RiskAuditV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAuditV2WithDefaults

`func NewRiskAuditV2WithDefaults() *RiskAuditV2`

NewRiskAuditV2WithDefaults instantiates a new RiskAuditV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *RiskAuditV2) GetMeta() RiskAuditMetaV2`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RiskAuditV2) GetMetaOk() (*RiskAuditMetaV2, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RiskAuditV2) SetMeta(v RiskAuditMetaV2)`

SetMeta sets Meta field to given value.


### GetParticipants

`func (o *RiskAuditV2) GetParticipants() RiskAuditParticipantsV2`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *RiskAuditV2) GetParticipantsOk() (*RiskAuditParticipantsV2, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *RiskAuditV2) SetParticipants(v RiskAuditParticipantsV2)`

SetParticipants sets Participants field to given value.


### GetVerdict

`func (o *RiskAuditV2) GetVerdict() RiskAuditVerdictV2`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *RiskAuditV2) GetVerdictOk() (*RiskAuditVerdictV2, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *RiskAuditV2) SetVerdict(v RiskAuditVerdictV2)`

SetVerdict sets Verdict field to given value.


### GetScores

`func (o *RiskAuditV2) GetScores() RiskAuditScoresV2`

GetScores returns the Scores field if non-nil, zero value otherwise.

### GetScoresOk

`func (o *RiskAuditV2) GetScoresOk() (*RiskAuditScoresV2, bool)`

GetScoresOk returns a tuple with the Scores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScores

`func (o *RiskAuditV2) SetScores(v RiskAuditScoresV2)`

SetScores sets Scores field to given value.


### GetDetections

`func (o *RiskAuditV2) GetDetections() RiskAuditDetectionsV2`

GetDetections returns the Detections field if non-nil, zero value otherwise.

### GetDetectionsOk

`func (o *RiskAuditV2) GetDetectionsOk() (*RiskAuditDetectionsV2, bool)`

GetDetectionsOk returns a tuple with the Detections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetections

`func (o *RiskAuditV2) SetDetections(v RiskAuditDetectionsV2)`

SetDetections sets Detections field to given value.


### GetAnalysis

`func (o *RiskAuditV2) GetAnalysis() RiskAuditAnalysisV2`

GetAnalysis returns the Analysis field if non-nil, zero value otherwise.

### GetAnalysisOk

`func (o *RiskAuditV2) GetAnalysisOk() (*RiskAuditAnalysisV2, bool)`

GetAnalysisOk returns a tuple with the Analysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysis

`func (o *RiskAuditV2) SetAnalysis(v RiskAuditAnalysisV2)`

SetAnalysis sets Analysis field to given value.


### GetTimeline

`func (o *RiskAuditV2) GetTimeline() RiskAuditTimelineV2`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *RiskAuditV2) GetTimelineOk() (*RiskAuditTimelineV2, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *RiskAuditV2) SetTimeline(v RiskAuditTimelineV2)`

SetTimeline sets Timeline field to given value.


### GetAudioEventModel

`func (o *RiskAuditV2) GetAudioEventModel() RiskAuditAudioEventModelV2`

GetAudioEventModel returns the AudioEventModel field if non-nil, zero value otherwise.

### GetAudioEventModelOk

`func (o *RiskAuditV2) GetAudioEventModelOk() (*RiskAuditAudioEventModelV2, bool)`

GetAudioEventModelOk returns a tuple with the AudioEventModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioEventModel

`func (o *RiskAuditV2) SetAudioEventModel(v RiskAuditAudioEventModelV2)`

SetAudioEventModel sets AudioEventModel field to given value.


### GetCategoriesSummary

`func (o *RiskAuditV2) GetCategoriesSummary() map[string]interface{}`

GetCategoriesSummary returns the CategoriesSummary field if non-nil, zero value otherwise.

### GetCategoriesSummaryOk

`func (o *RiskAuditV2) GetCategoriesSummaryOk() (*map[string]interface{}, bool)`

GetCategoriesSummaryOk returns a tuple with the CategoriesSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoriesSummary

`func (o *RiskAuditV2) SetCategoriesSummary(v map[string]interface{})`

SetCategoriesSummary sets CategoriesSummary field to given value.


### GetIndexer

`func (o *RiskAuditV2) GetIndexer() RiskAuditIndexerV2`

GetIndexer returns the Indexer field if non-nil, zero value otherwise.

### GetIndexerOk

`func (o *RiskAuditV2) GetIndexerOk() (*RiskAuditIndexerV2, bool)`

GetIndexerOk returns a tuple with the Indexer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexer

`func (o *RiskAuditV2) SetIndexer(v RiskAuditIndexerV2)`

SetIndexer sets Indexer field to given value.


### GetSummary

`func (o *RiskAuditV2) GetSummary() RiskAuditSummaryV2`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *RiskAuditV2) GetSummaryOk() (*RiskAuditSummaryV2, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *RiskAuditV2) SetSummary(v RiskAuditSummaryV2)`

SetSummary sets Summary field to given value.


### GetActionsI18n

`func (o *RiskAuditV2) GetActionsI18n() map[string]interface{}`

GetActionsI18n returns the ActionsI18n field if non-nil, zero value otherwise.

### GetActionsI18nOk

`func (o *RiskAuditV2) GetActionsI18nOk() (*map[string]interface{}, bool)`

GetActionsI18nOk returns a tuple with the ActionsI18n field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionsI18n

`func (o *RiskAuditV2) SetActionsI18n(v map[string]interface{})`

SetActionsI18n sets ActionsI18n field to given value.


### GetAuditDecisions

`func (o *RiskAuditV2) GetAuditDecisions() RiskAuditAuditDecisionsV2`

GetAuditDecisions returns the AuditDecisions field if non-nil, zero value otherwise.

### GetAuditDecisionsOk

`func (o *RiskAuditV2) GetAuditDecisionsOk() (*RiskAuditAuditDecisionsV2, bool)`

GetAuditDecisionsOk returns a tuple with the AuditDecisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditDecisions

`func (o *RiskAuditV2) SetAuditDecisions(v RiskAuditAuditDecisionsV2)`

SetAuditDecisions sets AuditDecisions field to given value.


### GetScoringExplanation

`func (o *RiskAuditV2) GetScoringExplanation() RiskAuditScoringExplanationV2`

GetScoringExplanation returns the ScoringExplanation field if non-nil, zero value otherwise.

### GetScoringExplanationOk

`func (o *RiskAuditV2) GetScoringExplanationOk() (*RiskAuditScoringExplanationV2, bool)`

GetScoringExplanationOk returns a tuple with the ScoringExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoringExplanation

`func (o *RiskAuditV2) SetScoringExplanation(v RiskAuditScoringExplanationV2)`

SetScoringExplanation sets ScoringExplanation field to given value.


### GetHtmlReport

`func (o *RiskAuditV2) GetHtmlReport() string`

GetHtmlReport returns the HtmlReport field if non-nil, zero value otherwise.

### GetHtmlReportOk

`func (o *RiskAuditV2) GetHtmlReportOk() (*string, bool)`

GetHtmlReportOk returns a tuple with the HtmlReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlReport

`func (o *RiskAuditV2) SetHtmlReport(v string)`

SetHtmlReport sets HtmlReport field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


