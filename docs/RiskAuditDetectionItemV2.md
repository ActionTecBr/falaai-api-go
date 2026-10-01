# RiskAuditDetectionItemV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Turn** | Pointer to **int32** | Turn number | [optional] 
**Interlocutor** | Pointer to **string** | Speaker | [optional] 
**Role** | Pointer to **string** | Role | [optional] 
**TimestampStartS** | Pointer to **float32** | Start (s) | [optional] 
**TimestampEndS** | Pointer to **float32** | End (s) | [optional] 
**TimestampFormatted** | Pointer to **string** | Formatted timestamp | [optional] 
**TermText** | Pointer to **string** | Detected term | [optional] 
**SuggestedTermForBank** | Pointer to **interface{}** | Suggested term for bank | [optional] 
**Category** | Pointer to **string** | Category code | [optional] 
**CategoryLabel** | **string** | Category label (i18n) | 
**CategoryColor** | Pointer to **string** | Category color | [optional] 
**CategoryIcon** | Pointer to **string** | Category icon | [optional] 
**Criticality** | Pointer to **string** | Criticality | [optional] 
**CategoryThreshold** | Pointer to **float32** | Category threshold | [optional] 
**CategoryType** | Pointer to **string** | Category type | [optional] 
**CategoryGroup** | **string** | Category group label (i18n) | 
**Nature** | Pointer to **string** | Nature | [optional] 
**LlmConfidence** | Pointer to **float32** | LLM confidence | [optional] 
**Reason** | Pointer to **string** | Reason | [optional] 
**IsValidContext** | Pointer to **bool** | Valid context | [optional] 
**RiskProbability** | Pointer to **float32** | Risk probability | [optional] 
**RiskImpact** | Pointer to **float32** | Risk impact | [optional] 
**CategoryWeight** | Pointer to **float32** | Category weight | [optional] 
**TurnSentiment** | Pointer to **string** | Turn sentiment | [optional] 
**Intensity** | Pointer to **interface{}** | Intensity | [optional] 
**ModApplied** | Pointer to **float32** | Total modifier applied | [optional] 
**MacApplied** | Pointer to **float32** | Audio modifier applied | [optional] 
**MvadApplied** | Pointer to **float32** | Intensity modifier applied | [optional] 
**ModFormula** | Pointer to **string** | Modifier formula | [optional] 
**MacDetails** | Pointer to **[]map[string]interface{}** | MAC details | [optional] [default to {}]
**CalibrationReason** | Pointer to **string** | Calibration reason | [optional] 
**FinalScore** | Pointer to **float32** | Final score | [optional] 
**FinalScoreFormula** | Pointer to **string** | Final score formula | [optional] 
**ConversationLimit** | Pointer to **interface{}** | Conversation limit | [optional] 
**ApplySaturation** | Pointer to **bool** | Apply saturation | [optional] 
**BlockRepetition** | Pointer to **bool** | Block repetition | [optional] 
**Status** | Pointer to **string** | Status | [optional] 
**EffectiveImpact** | Pointer to **float32** | Effective impact | [optional] 
**SaturationFactor** | Pointer to **float32** | Saturation factor | [optional] 
**SaturationFormula** | Pointer to **string** | Saturation formula | [optional] 
**ThresholdFormula** | Pointer to **string** | Threshold formula | [optional] 
**BlockedFormula** | Pointer to **string** | Blocked formula | [optional] 
**ReconciliationNote** | Pointer to **string** | Reconciliation note | [optional] 
**ViolatedFrameworks** | Pointer to **[]interface{}** | Violated frameworks | [optional] [default to {}]
**CitationFidelity** | Pointer to **bool** | Citation fidelity | [optional] [default to true]
**Subcategory** | Pointer to **string** | Subcategory code | [optional] 
**SubcategoryLabel** | Pointer to **string** | Subcategory label (i18n) | [optional] 

## Methods

### NewRiskAuditDetectionItemV2

`func NewRiskAuditDetectionItemV2(categoryLabel string, categoryGroup string, ) *RiskAuditDetectionItemV2`

NewRiskAuditDetectionItemV2 instantiates a new RiskAuditDetectionItemV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAuditDetectionItemV2WithDefaults

`func NewRiskAuditDetectionItemV2WithDefaults() *RiskAuditDetectionItemV2`

NewRiskAuditDetectionItemV2WithDefaults instantiates a new RiskAuditDetectionItemV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTurn

`func (o *RiskAuditDetectionItemV2) GetTurn() int32`

GetTurn returns the Turn field if non-nil, zero value otherwise.

### GetTurnOk

`func (o *RiskAuditDetectionItemV2) GetTurnOk() (*int32, bool)`

GetTurnOk returns a tuple with the Turn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurn

`func (o *RiskAuditDetectionItemV2) SetTurn(v int32)`

SetTurn sets Turn field to given value.

### HasTurn

`func (o *RiskAuditDetectionItemV2) HasTurn() bool`

HasTurn returns a boolean if a field has been set.

### GetInterlocutor

`func (o *RiskAuditDetectionItemV2) GetInterlocutor() string`

GetInterlocutor returns the Interlocutor field if non-nil, zero value otherwise.

### GetInterlocutorOk

`func (o *RiskAuditDetectionItemV2) GetInterlocutorOk() (*string, bool)`

GetInterlocutorOk returns a tuple with the Interlocutor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterlocutor

`func (o *RiskAuditDetectionItemV2) SetInterlocutor(v string)`

SetInterlocutor sets Interlocutor field to given value.

### HasInterlocutor

`func (o *RiskAuditDetectionItemV2) HasInterlocutor() bool`

HasInterlocutor returns a boolean if a field has been set.

### GetRole

`func (o *RiskAuditDetectionItemV2) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RiskAuditDetectionItemV2) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RiskAuditDetectionItemV2) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *RiskAuditDetectionItemV2) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTimestampStartS

`func (o *RiskAuditDetectionItemV2) GetTimestampStartS() float32`

GetTimestampStartS returns the TimestampStartS field if non-nil, zero value otherwise.

### GetTimestampStartSOk

`func (o *RiskAuditDetectionItemV2) GetTimestampStartSOk() (*float32, bool)`

GetTimestampStartSOk returns a tuple with the TimestampStartS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampStartS

`func (o *RiskAuditDetectionItemV2) SetTimestampStartS(v float32)`

SetTimestampStartS sets TimestampStartS field to given value.

### HasTimestampStartS

`func (o *RiskAuditDetectionItemV2) HasTimestampStartS() bool`

HasTimestampStartS returns a boolean if a field has been set.

### GetTimestampEndS

`func (o *RiskAuditDetectionItemV2) GetTimestampEndS() float32`

GetTimestampEndS returns the TimestampEndS field if non-nil, zero value otherwise.

### GetTimestampEndSOk

`func (o *RiskAuditDetectionItemV2) GetTimestampEndSOk() (*float32, bool)`

GetTimestampEndSOk returns a tuple with the TimestampEndS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampEndS

`func (o *RiskAuditDetectionItemV2) SetTimestampEndS(v float32)`

SetTimestampEndS sets TimestampEndS field to given value.

### HasTimestampEndS

`func (o *RiskAuditDetectionItemV2) HasTimestampEndS() bool`

HasTimestampEndS returns a boolean if a field has been set.

### GetTimestampFormatted

`func (o *RiskAuditDetectionItemV2) GetTimestampFormatted() string`

GetTimestampFormatted returns the TimestampFormatted field if non-nil, zero value otherwise.

### GetTimestampFormattedOk

`func (o *RiskAuditDetectionItemV2) GetTimestampFormattedOk() (*string, bool)`

GetTimestampFormattedOk returns a tuple with the TimestampFormatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampFormatted

`func (o *RiskAuditDetectionItemV2) SetTimestampFormatted(v string)`

SetTimestampFormatted sets TimestampFormatted field to given value.

### HasTimestampFormatted

`func (o *RiskAuditDetectionItemV2) HasTimestampFormatted() bool`

HasTimestampFormatted returns a boolean if a field has been set.

### GetTermText

`func (o *RiskAuditDetectionItemV2) GetTermText() string`

GetTermText returns the TermText field if non-nil, zero value otherwise.

### GetTermTextOk

`func (o *RiskAuditDetectionItemV2) GetTermTextOk() (*string, bool)`

GetTermTextOk returns a tuple with the TermText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermText

`func (o *RiskAuditDetectionItemV2) SetTermText(v string)`

SetTermText sets TermText field to given value.

### HasTermText

`func (o *RiskAuditDetectionItemV2) HasTermText() bool`

HasTermText returns a boolean if a field has been set.

### GetSuggestedTermForBank

`func (o *RiskAuditDetectionItemV2) GetSuggestedTermForBank() interface{}`

GetSuggestedTermForBank returns the SuggestedTermForBank field if non-nil, zero value otherwise.

### GetSuggestedTermForBankOk

`func (o *RiskAuditDetectionItemV2) GetSuggestedTermForBankOk() (*interface{}, bool)`

GetSuggestedTermForBankOk returns a tuple with the SuggestedTermForBank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedTermForBank

`func (o *RiskAuditDetectionItemV2) SetSuggestedTermForBank(v interface{})`

SetSuggestedTermForBank sets SuggestedTermForBank field to given value.

### HasSuggestedTermForBank

`func (o *RiskAuditDetectionItemV2) HasSuggestedTermForBank() bool`

HasSuggestedTermForBank returns a boolean if a field has been set.

### SetSuggestedTermForBankNil

`func (o *RiskAuditDetectionItemV2) SetSuggestedTermForBankNil(b bool)`

 SetSuggestedTermForBankNil sets the value for SuggestedTermForBank to be an explicit nil

### UnsetSuggestedTermForBank
`func (o *RiskAuditDetectionItemV2) UnsetSuggestedTermForBank()`

UnsetSuggestedTermForBank ensures that no value is present for SuggestedTermForBank, not even an explicit nil
### GetCategory

`func (o *RiskAuditDetectionItemV2) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *RiskAuditDetectionItemV2) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *RiskAuditDetectionItemV2) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *RiskAuditDetectionItemV2) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCategoryLabel

`func (o *RiskAuditDetectionItemV2) GetCategoryLabel() string`

GetCategoryLabel returns the CategoryLabel field if non-nil, zero value otherwise.

### GetCategoryLabelOk

`func (o *RiskAuditDetectionItemV2) GetCategoryLabelOk() (*string, bool)`

GetCategoryLabelOk returns a tuple with the CategoryLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryLabel

`func (o *RiskAuditDetectionItemV2) SetCategoryLabel(v string)`

SetCategoryLabel sets CategoryLabel field to given value.


### GetCategoryColor

`func (o *RiskAuditDetectionItemV2) GetCategoryColor() string`

GetCategoryColor returns the CategoryColor field if non-nil, zero value otherwise.

### GetCategoryColorOk

`func (o *RiskAuditDetectionItemV2) GetCategoryColorOk() (*string, bool)`

GetCategoryColorOk returns a tuple with the CategoryColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryColor

`func (o *RiskAuditDetectionItemV2) SetCategoryColor(v string)`

SetCategoryColor sets CategoryColor field to given value.

### HasCategoryColor

`func (o *RiskAuditDetectionItemV2) HasCategoryColor() bool`

HasCategoryColor returns a boolean if a field has been set.

### GetCategoryIcon

`func (o *RiskAuditDetectionItemV2) GetCategoryIcon() string`

GetCategoryIcon returns the CategoryIcon field if non-nil, zero value otherwise.

### GetCategoryIconOk

`func (o *RiskAuditDetectionItemV2) GetCategoryIconOk() (*string, bool)`

GetCategoryIconOk returns a tuple with the CategoryIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryIcon

`func (o *RiskAuditDetectionItemV2) SetCategoryIcon(v string)`

SetCategoryIcon sets CategoryIcon field to given value.

### HasCategoryIcon

`func (o *RiskAuditDetectionItemV2) HasCategoryIcon() bool`

HasCategoryIcon returns a boolean if a field has been set.

### GetCriticality

`func (o *RiskAuditDetectionItemV2) GetCriticality() string`

GetCriticality returns the Criticality field if non-nil, zero value otherwise.

### GetCriticalityOk

`func (o *RiskAuditDetectionItemV2) GetCriticalityOk() (*string, bool)`

GetCriticalityOk returns a tuple with the Criticality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticality

`func (o *RiskAuditDetectionItemV2) SetCriticality(v string)`

SetCriticality sets Criticality field to given value.

### HasCriticality

`func (o *RiskAuditDetectionItemV2) HasCriticality() bool`

HasCriticality returns a boolean if a field has been set.

### GetCategoryThreshold

`func (o *RiskAuditDetectionItemV2) GetCategoryThreshold() float32`

GetCategoryThreshold returns the CategoryThreshold field if non-nil, zero value otherwise.

### GetCategoryThresholdOk

`func (o *RiskAuditDetectionItemV2) GetCategoryThresholdOk() (*float32, bool)`

GetCategoryThresholdOk returns a tuple with the CategoryThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryThreshold

`func (o *RiskAuditDetectionItemV2) SetCategoryThreshold(v float32)`

SetCategoryThreshold sets CategoryThreshold field to given value.

### HasCategoryThreshold

`func (o *RiskAuditDetectionItemV2) HasCategoryThreshold() bool`

HasCategoryThreshold returns a boolean if a field has been set.

### GetCategoryType

`func (o *RiskAuditDetectionItemV2) GetCategoryType() string`

GetCategoryType returns the CategoryType field if non-nil, zero value otherwise.

### GetCategoryTypeOk

`func (o *RiskAuditDetectionItemV2) GetCategoryTypeOk() (*string, bool)`

GetCategoryTypeOk returns a tuple with the CategoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryType

`func (o *RiskAuditDetectionItemV2) SetCategoryType(v string)`

SetCategoryType sets CategoryType field to given value.

### HasCategoryType

`func (o *RiskAuditDetectionItemV2) HasCategoryType() bool`

HasCategoryType returns a boolean if a field has been set.

### GetCategoryGroup

`func (o *RiskAuditDetectionItemV2) GetCategoryGroup() string`

GetCategoryGroup returns the CategoryGroup field if non-nil, zero value otherwise.

### GetCategoryGroupOk

`func (o *RiskAuditDetectionItemV2) GetCategoryGroupOk() (*string, bool)`

GetCategoryGroupOk returns a tuple with the CategoryGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryGroup

`func (o *RiskAuditDetectionItemV2) SetCategoryGroup(v string)`

SetCategoryGroup sets CategoryGroup field to given value.


### GetNature

`func (o *RiskAuditDetectionItemV2) GetNature() string`

GetNature returns the Nature field if non-nil, zero value otherwise.

### GetNatureOk

`func (o *RiskAuditDetectionItemV2) GetNatureOk() (*string, bool)`

GetNatureOk returns a tuple with the Nature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNature

`func (o *RiskAuditDetectionItemV2) SetNature(v string)`

SetNature sets Nature field to given value.

### HasNature

`func (o *RiskAuditDetectionItemV2) HasNature() bool`

HasNature returns a boolean if a field has been set.

### GetLlmConfidence

`func (o *RiskAuditDetectionItemV2) GetLlmConfidence() float32`

GetLlmConfidence returns the LlmConfidence field if non-nil, zero value otherwise.

### GetLlmConfidenceOk

`func (o *RiskAuditDetectionItemV2) GetLlmConfidenceOk() (*float32, bool)`

GetLlmConfidenceOk returns a tuple with the LlmConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlmConfidence

`func (o *RiskAuditDetectionItemV2) SetLlmConfidence(v float32)`

SetLlmConfidence sets LlmConfidence field to given value.

### HasLlmConfidence

`func (o *RiskAuditDetectionItemV2) HasLlmConfidence() bool`

HasLlmConfidence returns a boolean if a field has been set.

### GetReason

`func (o *RiskAuditDetectionItemV2) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RiskAuditDetectionItemV2) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RiskAuditDetectionItemV2) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RiskAuditDetectionItemV2) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetIsValidContext

`func (o *RiskAuditDetectionItemV2) GetIsValidContext() bool`

GetIsValidContext returns the IsValidContext field if non-nil, zero value otherwise.

### GetIsValidContextOk

`func (o *RiskAuditDetectionItemV2) GetIsValidContextOk() (*bool, bool)`

GetIsValidContextOk returns a tuple with the IsValidContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValidContext

`func (o *RiskAuditDetectionItemV2) SetIsValidContext(v bool)`

SetIsValidContext sets IsValidContext field to given value.

### HasIsValidContext

`func (o *RiskAuditDetectionItemV2) HasIsValidContext() bool`

HasIsValidContext returns a boolean if a field has been set.

### GetRiskProbability

`func (o *RiskAuditDetectionItemV2) GetRiskProbability() float32`

GetRiskProbability returns the RiskProbability field if non-nil, zero value otherwise.

### GetRiskProbabilityOk

`func (o *RiskAuditDetectionItemV2) GetRiskProbabilityOk() (*float32, bool)`

GetRiskProbabilityOk returns a tuple with the RiskProbability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskProbability

`func (o *RiskAuditDetectionItemV2) SetRiskProbability(v float32)`

SetRiskProbability sets RiskProbability field to given value.

### HasRiskProbability

`func (o *RiskAuditDetectionItemV2) HasRiskProbability() bool`

HasRiskProbability returns a boolean if a field has been set.

### GetRiskImpact

`func (o *RiskAuditDetectionItemV2) GetRiskImpact() float32`

GetRiskImpact returns the RiskImpact field if non-nil, zero value otherwise.

### GetRiskImpactOk

`func (o *RiskAuditDetectionItemV2) GetRiskImpactOk() (*float32, bool)`

GetRiskImpactOk returns a tuple with the RiskImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskImpact

`func (o *RiskAuditDetectionItemV2) SetRiskImpact(v float32)`

SetRiskImpact sets RiskImpact field to given value.

### HasRiskImpact

`func (o *RiskAuditDetectionItemV2) HasRiskImpact() bool`

HasRiskImpact returns a boolean if a field has been set.

### GetCategoryWeight

`func (o *RiskAuditDetectionItemV2) GetCategoryWeight() float32`

GetCategoryWeight returns the CategoryWeight field if non-nil, zero value otherwise.

### GetCategoryWeightOk

`func (o *RiskAuditDetectionItemV2) GetCategoryWeightOk() (*float32, bool)`

GetCategoryWeightOk returns a tuple with the CategoryWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryWeight

`func (o *RiskAuditDetectionItemV2) SetCategoryWeight(v float32)`

SetCategoryWeight sets CategoryWeight field to given value.

### HasCategoryWeight

`func (o *RiskAuditDetectionItemV2) HasCategoryWeight() bool`

HasCategoryWeight returns a boolean if a field has been set.

### GetTurnSentiment

`func (o *RiskAuditDetectionItemV2) GetTurnSentiment() string`

GetTurnSentiment returns the TurnSentiment field if non-nil, zero value otherwise.

### GetTurnSentimentOk

`func (o *RiskAuditDetectionItemV2) GetTurnSentimentOk() (*string, bool)`

GetTurnSentimentOk returns a tuple with the TurnSentiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurnSentiment

`func (o *RiskAuditDetectionItemV2) SetTurnSentiment(v string)`

SetTurnSentiment sets TurnSentiment field to given value.

### HasTurnSentiment

`func (o *RiskAuditDetectionItemV2) HasTurnSentiment() bool`

HasTurnSentiment returns a boolean if a field has been set.

### GetIntensity

`func (o *RiskAuditDetectionItemV2) GetIntensity() interface{}`

GetIntensity returns the Intensity field if non-nil, zero value otherwise.

### GetIntensityOk

`func (o *RiskAuditDetectionItemV2) GetIntensityOk() (*interface{}, bool)`

GetIntensityOk returns a tuple with the Intensity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntensity

`func (o *RiskAuditDetectionItemV2) SetIntensity(v interface{})`

SetIntensity sets Intensity field to given value.

### HasIntensity

`func (o *RiskAuditDetectionItemV2) HasIntensity() bool`

HasIntensity returns a boolean if a field has been set.

### SetIntensityNil

`func (o *RiskAuditDetectionItemV2) SetIntensityNil(b bool)`

 SetIntensityNil sets the value for Intensity to be an explicit nil

### UnsetIntensity
`func (o *RiskAuditDetectionItemV2) UnsetIntensity()`

UnsetIntensity ensures that no value is present for Intensity, not even an explicit nil
### GetModApplied

`func (o *RiskAuditDetectionItemV2) GetModApplied() float32`

GetModApplied returns the ModApplied field if non-nil, zero value otherwise.

### GetModAppliedOk

`func (o *RiskAuditDetectionItemV2) GetModAppliedOk() (*float32, bool)`

GetModAppliedOk returns a tuple with the ModApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModApplied

`func (o *RiskAuditDetectionItemV2) SetModApplied(v float32)`

SetModApplied sets ModApplied field to given value.

### HasModApplied

`func (o *RiskAuditDetectionItemV2) HasModApplied() bool`

HasModApplied returns a boolean if a field has been set.

### GetMacApplied

`func (o *RiskAuditDetectionItemV2) GetMacApplied() float32`

GetMacApplied returns the MacApplied field if non-nil, zero value otherwise.

### GetMacAppliedOk

`func (o *RiskAuditDetectionItemV2) GetMacAppliedOk() (*float32, bool)`

GetMacAppliedOk returns a tuple with the MacApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacApplied

`func (o *RiskAuditDetectionItemV2) SetMacApplied(v float32)`

SetMacApplied sets MacApplied field to given value.

### HasMacApplied

`func (o *RiskAuditDetectionItemV2) HasMacApplied() bool`

HasMacApplied returns a boolean if a field has been set.

### GetMvadApplied

`func (o *RiskAuditDetectionItemV2) GetMvadApplied() float32`

GetMvadApplied returns the MvadApplied field if non-nil, zero value otherwise.

### GetMvadAppliedOk

`func (o *RiskAuditDetectionItemV2) GetMvadAppliedOk() (*float32, bool)`

GetMvadAppliedOk returns a tuple with the MvadApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMvadApplied

`func (o *RiskAuditDetectionItemV2) SetMvadApplied(v float32)`

SetMvadApplied sets MvadApplied field to given value.

### HasMvadApplied

`func (o *RiskAuditDetectionItemV2) HasMvadApplied() bool`

HasMvadApplied returns a boolean if a field has been set.

### GetModFormula

`func (o *RiskAuditDetectionItemV2) GetModFormula() string`

GetModFormula returns the ModFormula field if non-nil, zero value otherwise.

### GetModFormulaOk

`func (o *RiskAuditDetectionItemV2) GetModFormulaOk() (*string, bool)`

GetModFormulaOk returns a tuple with the ModFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModFormula

`func (o *RiskAuditDetectionItemV2) SetModFormula(v string)`

SetModFormula sets ModFormula field to given value.

### HasModFormula

`func (o *RiskAuditDetectionItemV2) HasModFormula() bool`

HasModFormula returns a boolean if a field has been set.

### GetMacDetails

`func (o *RiskAuditDetectionItemV2) GetMacDetails() []map[string]interface{}`

GetMacDetails returns the MacDetails field if non-nil, zero value otherwise.

### GetMacDetailsOk

`func (o *RiskAuditDetectionItemV2) GetMacDetailsOk() (*[]map[string]interface{}, bool)`

GetMacDetailsOk returns a tuple with the MacDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacDetails

`func (o *RiskAuditDetectionItemV2) SetMacDetails(v []map[string]interface{})`

SetMacDetails sets MacDetails field to given value.

### HasMacDetails

`func (o *RiskAuditDetectionItemV2) HasMacDetails() bool`

HasMacDetails returns a boolean if a field has been set.

### GetCalibrationReason

`func (o *RiskAuditDetectionItemV2) GetCalibrationReason() string`

GetCalibrationReason returns the CalibrationReason field if non-nil, zero value otherwise.

### GetCalibrationReasonOk

`func (o *RiskAuditDetectionItemV2) GetCalibrationReasonOk() (*string, bool)`

GetCalibrationReasonOk returns a tuple with the CalibrationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalibrationReason

`func (o *RiskAuditDetectionItemV2) SetCalibrationReason(v string)`

SetCalibrationReason sets CalibrationReason field to given value.

### HasCalibrationReason

`func (o *RiskAuditDetectionItemV2) HasCalibrationReason() bool`

HasCalibrationReason returns a boolean if a field has been set.

### GetFinalScore

`func (o *RiskAuditDetectionItemV2) GetFinalScore() float32`

GetFinalScore returns the FinalScore field if non-nil, zero value otherwise.

### GetFinalScoreOk

`func (o *RiskAuditDetectionItemV2) GetFinalScoreOk() (*float32, bool)`

GetFinalScoreOk returns a tuple with the FinalScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalScore

`func (o *RiskAuditDetectionItemV2) SetFinalScore(v float32)`

SetFinalScore sets FinalScore field to given value.

### HasFinalScore

`func (o *RiskAuditDetectionItemV2) HasFinalScore() bool`

HasFinalScore returns a boolean if a field has been set.

### GetFinalScoreFormula

`func (o *RiskAuditDetectionItemV2) GetFinalScoreFormula() string`

GetFinalScoreFormula returns the FinalScoreFormula field if non-nil, zero value otherwise.

### GetFinalScoreFormulaOk

`func (o *RiskAuditDetectionItemV2) GetFinalScoreFormulaOk() (*string, bool)`

GetFinalScoreFormulaOk returns a tuple with the FinalScoreFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalScoreFormula

`func (o *RiskAuditDetectionItemV2) SetFinalScoreFormula(v string)`

SetFinalScoreFormula sets FinalScoreFormula field to given value.

### HasFinalScoreFormula

`func (o *RiskAuditDetectionItemV2) HasFinalScoreFormula() bool`

HasFinalScoreFormula returns a boolean if a field has been set.

### GetConversationLimit

`func (o *RiskAuditDetectionItemV2) GetConversationLimit() interface{}`

GetConversationLimit returns the ConversationLimit field if non-nil, zero value otherwise.

### GetConversationLimitOk

`func (o *RiskAuditDetectionItemV2) GetConversationLimitOk() (*interface{}, bool)`

GetConversationLimitOk returns a tuple with the ConversationLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationLimit

`func (o *RiskAuditDetectionItemV2) SetConversationLimit(v interface{})`

SetConversationLimit sets ConversationLimit field to given value.

### HasConversationLimit

`func (o *RiskAuditDetectionItemV2) HasConversationLimit() bool`

HasConversationLimit returns a boolean if a field has been set.

### SetConversationLimitNil

`func (o *RiskAuditDetectionItemV2) SetConversationLimitNil(b bool)`

 SetConversationLimitNil sets the value for ConversationLimit to be an explicit nil

### UnsetConversationLimit
`func (o *RiskAuditDetectionItemV2) UnsetConversationLimit()`

UnsetConversationLimit ensures that no value is present for ConversationLimit, not even an explicit nil
### GetApplySaturation

`func (o *RiskAuditDetectionItemV2) GetApplySaturation() bool`

GetApplySaturation returns the ApplySaturation field if non-nil, zero value otherwise.

### GetApplySaturationOk

`func (o *RiskAuditDetectionItemV2) GetApplySaturationOk() (*bool, bool)`

GetApplySaturationOk returns a tuple with the ApplySaturation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplySaturation

`func (o *RiskAuditDetectionItemV2) SetApplySaturation(v bool)`

SetApplySaturation sets ApplySaturation field to given value.

### HasApplySaturation

`func (o *RiskAuditDetectionItemV2) HasApplySaturation() bool`

HasApplySaturation returns a boolean if a field has been set.

### GetBlockRepetition

`func (o *RiskAuditDetectionItemV2) GetBlockRepetition() bool`

GetBlockRepetition returns the BlockRepetition field if non-nil, zero value otherwise.

### GetBlockRepetitionOk

`func (o *RiskAuditDetectionItemV2) GetBlockRepetitionOk() (*bool, bool)`

GetBlockRepetitionOk returns a tuple with the BlockRepetition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockRepetition

`func (o *RiskAuditDetectionItemV2) SetBlockRepetition(v bool)`

SetBlockRepetition sets BlockRepetition field to given value.

### HasBlockRepetition

`func (o *RiskAuditDetectionItemV2) HasBlockRepetition() bool`

HasBlockRepetition returns a boolean if a field has been set.

### GetStatus

`func (o *RiskAuditDetectionItemV2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RiskAuditDetectionItemV2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RiskAuditDetectionItemV2) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RiskAuditDetectionItemV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEffectiveImpact

`func (o *RiskAuditDetectionItemV2) GetEffectiveImpact() float32`

GetEffectiveImpact returns the EffectiveImpact field if non-nil, zero value otherwise.

### GetEffectiveImpactOk

`func (o *RiskAuditDetectionItemV2) GetEffectiveImpactOk() (*float32, bool)`

GetEffectiveImpactOk returns a tuple with the EffectiveImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveImpact

`func (o *RiskAuditDetectionItemV2) SetEffectiveImpact(v float32)`

SetEffectiveImpact sets EffectiveImpact field to given value.

### HasEffectiveImpact

`func (o *RiskAuditDetectionItemV2) HasEffectiveImpact() bool`

HasEffectiveImpact returns a boolean if a field has been set.

### GetSaturationFactor

`func (o *RiskAuditDetectionItemV2) GetSaturationFactor() float32`

GetSaturationFactor returns the SaturationFactor field if non-nil, zero value otherwise.

### GetSaturationFactorOk

`func (o *RiskAuditDetectionItemV2) GetSaturationFactorOk() (*float32, bool)`

GetSaturationFactorOk returns a tuple with the SaturationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturationFactor

`func (o *RiskAuditDetectionItemV2) SetSaturationFactor(v float32)`

SetSaturationFactor sets SaturationFactor field to given value.

### HasSaturationFactor

`func (o *RiskAuditDetectionItemV2) HasSaturationFactor() bool`

HasSaturationFactor returns a boolean if a field has been set.

### GetSaturationFormula

`func (o *RiskAuditDetectionItemV2) GetSaturationFormula() string`

GetSaturationFormula returns the SaturationFormula field if non-nil, zero value otherwise.

### GetSaturationFormulaOk

`func (o *RiskAuditDetectionItemV2) GetSaturationFormulaOk() (*string, bool)`

GetSaturationFormulaOk returns a tuple with the SaturationFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturationFormula

`func (o *RiskAuditDetectionItemV2) SetSaturationFormula(v string)`

SetSaturationFormula sets SaturationFormula field to given value.

### HasSaturationFormula

`func (o *RiskAuditDetectionItemV2) HasSaturationFormula() bool`

HasSaturationFormula returns a boolean if a field has been set.

### GetThresholdFormula

`func (o *RiskAuditDetectionItemV2) GetThresholdFormula() string`

GetThresholdFormula returns the ThresholdFormula field if non-nil, zero value otherwise.

### GetThresholdFormulaOk

`func (o *RiskAuditDetectionItemV2) GetThresholdFormulaOk() (*string, bool)`

GetThresholdFormulaOk returns a tuple with the ThresholdFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdFormula

`func (o *RiskAuditDetectionItemV2) SetThresholdFormula(v string)`

SetThresholdFormula sets ThresholdFormula field to given value.

### HasThresholdFormula

`func (o *RiskAuditDetectionItemV2) HasThresholdFormula() bool`

HasThresholdFormula returns a boolean if a field has been set.

### GetBlockedFormula

`func (o *RiskAuditDetectionItemV2) GetBlockedFormula() string`

GetBlockedFormula returns the BlockedFormula field if non-nil, zero value otherwise.

### GetBlockedFormulaOk

`func (o *RiskAuditDetectionItemV2) GetBlockedFormulaOk() (*string, bool)`

GetBlockedFormulaOk returns a tuple with the BlockedFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedFormula

`func (o *RiskAuditDetectionItemV2) SetBlockedFormula(v string)`

SetBlockedFormula sets BlockedFormula field to given value.

### HasBlockedFormula

`func (o *RiskAuditDetectionItemV2) HasBlockedFormula() bool`

HasBlockedFormula returns a boolean if a field has been set.

### GetReconciliationNote

`func (o *RiskAuditDetectionItemV2) GetReconciliationNote() string`

GetReconciliationNote returns the ReconciliationNote field if non-nil, zero value otherwise.

### GetReconciliationNoteOk

`func (o *RiskAuditDetectionItemV2) GetReconciliationNoteOk() (*string, bool)`

GetReconciliationNoteOk returns a tuple with the ReconciliationNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciliationNote

`func (o *RiskAuditDetectionItemV2) SetReconciliationNote(v string)`

SetReconciliationNote sets ReconciliationNote field to given value.

### HasReconciliationNote

`func (o *RiskAuditDetectionItemV2) HasReconciliationNote() bool`

HasReconciliationNote returns a boolean if a field has been set.

### GetViolatedFrameworks

`func (o *RiskAuditDetectionItemV2) GetViolatedFrameworks() []interface{}`

GetViolatedFrameworks returns the ViolatedFrameworks field if non-nil, zero value otherwise.

### GetViolatedFrameworksOk

`func (o *RiskAuditDetectionItemV2) GetViolatedFrameworksOk() (*[]interface{}, bool)`

GetViolatedFrameworksOk returns a tuple with the ViolatedFrameworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolatedFrameworks

`func (o *RiskAuditDetectionItemV2) SetViolatedFrameworks(v []interface{})`

SetViolatedFrameworks sets ViolatedFrameworks field to given value.

### HasViolatedFrameworks

`func (o *RiskAuditDetectionItemV2) HasViolatedFrameworks() bool`

HasViolatedFrameworks returns a boolean if a field has been set.

### GetCitationFidelity

`func (o *RiskAuditDetectionItemV2) GetCitationFidelity() bool`

GetCitationFidelity returns the CitationFidelity field if non-nil, zero value otherwise.

### GetCitationFidelityOk

`func (o *RiskAuditDetectionItemV2) GetCitationFidelityOk() (*bool, bool)`

GetCitationFidelityOk returns a tuple with the CitationFidelity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitationFidelity

`func (o *RiskAuditDetectionItemV2) SetCitationFidelity(v bool)`

SetCitationFidelity sets CitationFidelity field to given value.

### HasCitationFidelity

`func (o *RiskAuditDetectionItemV2) HasCitationFidelity() bool`

HasCitationFidelity returns a boolean if a field has been set.

### GetSubcategory

`func (o *RiskAuditDetectionItemV2) GetSubcategory() string`

GetSubcategory returns the Subcategory field if non-nil, zero value otherwise.

### GetSubcategoryOk

`func (o *RiskAuditDetectionItemV2) GetSubcategoryOk() (*string, bool)`

GetSubcategoryOk returns a tuple with the Subcategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcategory

`func (o *RiskAuditDetectionItemV2) SetSubcategory(v string)`

SetSubcategory sets Subcategory field to given value.

### HasSubcategory

`func (o *RiskAuditDetectionItemV2) HasSubcategory() bool`

HasSubcategory returns a boolean if a field has been set.

### GetSubcategoryLabel

`func (o *RiskAuditDetectionItemV2) GetSubcategoryLabel() string`

GetSubcategoryLabel returns the SubcategoryLabel field if non-nil, zero value otherwise.

### GetSubcategoryLabelOk

`func (o *RiskAuditDetectionItemV2) GetSubcategoryLabelOk() (*string, bool)`

GetSubcategoryLabelOk returns a tuple with the SubcategoryLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcategoryLabel

`func (o *RiskAuditDetectionItemV2) SetSubcategoryLabel(v string)`

SetSubcategoryLabel sets SubcategoryLabel field to given value.

### HasSubcategoryLabel

`func (o *RiskAuditDetectionItemV2) HasSubcategoryLabel() bool`

HasSubcategoryLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


