# DiagnosticCategoricalAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListChoice** | Pointer to **string** | Selected value from classification list (used in action, label, sentiment) | [optional] 
**Justification** | Pointer to **string** | Justification for the choice | [optional] 
**EvidencePhrases** | Pointer to **[]string** | Verbatim transcript excerpts supporting the analysis | [optional] [default to {}]

## Methods

### NewDiagnosticCategoricalAnalysis

`func NewDiagnosticCategoricalAnalysis() *DiagnosticCategoricalAnalysis`

NewDiagnosticCategoricalAnalysis instantiates a new DiagnosticCategoricalAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosticCategoricalAnalysisWithDefaults

`func NewDiagnosticCategoricalAnalysisWithDefaults() *DiagnosticCategoricalAnalysis`

NewDiagnosticCategoricalAnalysisWithDefaults instantiates a new DiagnosticCategoricalAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListChoice

`func (o *DiagnosticCategoricalAnalysis) GetListChoice() string`

GetListChoice returns the ListChoice field if non-nil, zero value otherwise.

### GetListChoiceOk

`func (o *DiagnosticCategoricalAnalysis) GetListChoiceOk() (*string, bool)`

GetListChoiceOk returns a tuple with the ListChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListChoice

`func (o *DiagnosticCategoricalAnalysis) SetListChoice(v string)`

SetListChoice sets ListChoice field to given value.

### HasListChoice

`func (o *DiagnosticCategoricalAnalysis) HasListChoice() bool`

HasListChoice returns a boolean if a field has been set.

### GetJustification

`func (o *DiagnosticCategoricalAnalysis) GetJustification() string`

GetJustification returns the Justification field if non-nil, zero value otherwise.

### GetJustificationOk

`func (o *DiagnosticCategoricalAnalysis) GetJustificationOk() (*string, bool)`

GetJustificationOk returns a tuple with the Justification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJustification

`func (o *DiagnosticCategoricalAnalysis) SetJustification(v string)`

SetJustification sets Justification field to given value.

### HasJustification

`func (o *DiagnosticCategoricalAnalysis) HasJustification() bool`

HasJustification returns a boolean if a field has been set.

### GetEvidencePhrases

`func (o *DiagnosticCategoricalAnalysis) GetEvidencePhrases() []string`

GetEvidencePhrases returns the EvidencePhrases field if non-nil, zero value otherwise.

### GetEvidencePhrasesOk

`func (o *DiagnosticCategoricalAnalysis) GetEvidencePhrasesOk() (*[]string, bool)`

GetEvidencePhrasesOk returns a tuple with the EvidencePhrases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidencePhrases

`func (o *DiagnosticCategoricalAnalysis) SetEvidencePhrases(v []string)`

SetEvidencePhrases sets EvidencePhrases field to given value.

### HasEvidencePhrases

`func (o *DiagnosticCategoricalAnalysis) HasEvidencePhrases() bool`

HasEvidencePhrases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


