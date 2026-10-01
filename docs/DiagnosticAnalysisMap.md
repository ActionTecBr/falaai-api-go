# DiagnosticAnalysisMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DialogueSummary** | [**DiagnosticTextAnalysis**](DiagnosticTextAnalysis.md) | Detailed conversation summary | 
**ContactReason** | [**DiagnosticTextAnalysis**](DiagnosticTextAnalysis.md) | Initial contact reason | 
**IdentifiedAction** | [**DiagnosticCategoricalAnalysis**](DiagnosticCategoricalAnalysis.md) | Action taken / resolution | 
**IdentifiedLabel** | [**DiagnosticCategoricalAnalysis**](DiagnosticCategoricalAnalysis.md) | Theme classification | 
**Sentiment** | [**DiagnosticCategoricalAnalysis**](DiagnosticCategoricalAnalysis.md) | Predominant sentiment | 
**ParticipantsIdentified** | Pointer to [**[]ParticipantDiagnostic**](ParticipantDiagnostic.md) | Identified participants and roles (same field names as auditoria) | [optional] [default to {}]

## Methods

### NewDiagnosticAnalysisMap

`func NewDiagnosticAnalysisMap(dialogueSummary DiagnosticTextAnalysis, contactReason DiagnosticTextAnalysis, identifiedAction DiagnosticCategoricalAnalysis, identifiedLabel DiagnosticCategoricalAnalysis, sentiment DiagnosticCategoricalAnalysis, ) *DiagnosticAnalysisMap`

NewDiagnosticAnalysisMap instantiates a new DiagnosticAnalysisMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosticAnalysisMapWithDefaults

`func NewDiagnosticAnalysisMapWithDefaults() *DiagnosticAnalysisMap`

NewDiagnosticAnalysisMapWithDefaults instantiates a new DiagnosticAnalysisMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDialogueSummary

`func (o *DiagnosticAnalysisMap) GetDialogueSummary() DiagnosticTextAnalysis`

GetDialogueSummary returns the DialogueSummary field if non-nil, zero value otherwise.

### GetDialogueSummaryOk

`func (o *DiagnosticAnalysisMap) GetDialogueSummaryOk() (*DiagnosticTextAnalysis, bool)`

GetDialogueSummaryOk returns a tuple with the DialogueSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialogueSummary

`func (o *DiagnosticAnalysisMap) SetDialogueSummary(v DiagnosticTextAnalysis)`

SetDialogueSummary sets DialogueSummary field to given value.


### GetContactReason

`func (o *DiagnosticAnalysisMap) GetContactReason() DiagnosticTextAnalysis`

GetContactReason returns the ContactReason field if non-nil, zero value otherwise.

### GetContactReasonOk

`func (o *DiagnosticAnalysisMap) GetContactReasonOk() (*DiagnosticTextAnalysis, bool)`

GetContactReasonOk returns a tuple with the ContactReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactReason

`func (o *DiagnosticAnalysisMap) SetContactReason(v DiagnosticTextAnalysis)`

SetContactReason sets ContactReason field to given value.


### GetIdentifiedAction

`func (o *DiagnosticAnalysisMap) GetIdentifiedAction() DiagnosticCategoricalAnalysis`

GetIdentifiedAction returns the IdentifiedAction field if non-nil, zero value otherwise.

### GetIdentifiedActionOk

`func (o *DiagnosticAnalysisMap) GetIdentifiedActionOk() (*DiagnosticCategoricalAnalysis, bool)`

GetIdentifiedActionOk returns a tuple with the IdentifiedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiedAction

`func (o *DiagnosticAnalysisMap) SetIdentifiedAction(v DiagnosticCategoricalAnalysis)`

SetIdentifiedAction sets IdentifiedAction field to given value.


### GetIdentifiedLabel

`func (o *DiagnosticAnalysisMap) GetIdentifiedLabel() DiagnosticCategoricalAnalysis`

GetIdentifiedLabel returns the IdentifiedLabel field if non-nil, zero value otherwise.

### GetIdentifiedLabelOk

`func (o *DiagnosticAnalysisMap) GetIdentifiedLabelOk() (*DiagnosticCategoricalAnalysis, bool)`

GetIdentifiedLabelOk returns a tuple with the IdentifiedLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiedLabel

`func (o *DiagnosticAnalysisMap) SetIdentifiedLabel(v DiagnosticCategoricalAnalysis)`

SetIdentifiedLabel sets IdentifiedLabel field to given value.


### GetSentiment

`func (o *DiagnosticAnalysisMap) GetSentiment() DiagnosticCategoricalAnalysis`

GetSentiment returns the Sentiment field if non-nil, zero value otherwise.

### GetSentimentOk

`func (o *DiagnosticAnalysisMap) GetSentimentOk() (*DiagnosticCategoricalAnalysis, bool)`

GetSentimentOk returns a tuple with the Sentiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentiment

`func (o *DiagnosticAnalysisMap) SetSentiment(v DiagnosticCategoricalAnalysis)`

SetSentiment sets Sentiment field to given value.


### GetParticipantsIdentified

`func (o *DiagnosticAnalysisMap) GetParticipantsIdentified() []ParticipantDiagnostic`

GetParticipantsIdentified returns the ParticipantsIdentified field if non-nil, zero value otherwise.

### GetParticipantsIdentifiedOk

`func (o *DiagnosticAnalysisMap) GetParticipantsIdentifiedOk() (*[]ParticipantDiagnostic, bool)`

GetParticipantsIdentifiedOk returns a tuple with the ParticipantsIdentified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantsIdentified

`func (o *DiagnosticAnalysisMap) SetParticipantsIdentified(v []ParticipantDiagnostic)`

SetParticipantsIdentified sets ParticipantsIdentified field to given value.

### HasParticipantsIdentified

`func (o *DiagnosticAnalysisMap) HasParticipantsIdentified() bool`

HasParticipantsIdentified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


