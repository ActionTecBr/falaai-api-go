# RiskAuditAnalysisV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalMetrics** | Pointer to **map[string]interface{}** | Global metrics | [optional] [default to {}]
**FinalAnalysis** | Pointer to **map[string]interface{}** | Final analysis | [optional] [default to {}]
**Frameworks** | Pointer to **map[string]interface{}** | Frameworks (COPC/ISO/Kirkpatrick/CES) | [optional] [default to {}]

## Methods

### NewRiskAuditAnalysisV2

`func NewRiskAuditAnalysisV2() *RiskAuditAnalysisV2`

NewRiskAuditAnalysisV2 instantiates a new RiskAuditAnalysisV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAuditAnalysisV2WithDefaults

`func NewRiskAuditAnalysisV2WithDefaults() *RiskAuditAnalysisV2`

NewRiskAuditAnalysisV2WithDefaults instantiates a new RiskAuditAnalysisV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalMetrics

`func (o *RiskAuditAnalysisV2) GetGlobalMetrics() map[string]interface{}`

GetGlobalMetrics returns the GlobalMetrics field if non-nil, zero value otherwise.

### GetGlobalMetricsOk

`func (o *RiskAuditAnalysisV2) GetGlobalMetricsOk() (*map[string]interface{}, bool)`

GetGlobalMetricsOk returns a tuple with the GlobalMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalMetrics

`func (o *RiskAuditAnalysisV2) SetGlobalMetrics(v map[string]interface{})`

SetGlobalMetrics sets GlobalMetrics field to given value.

### HasGlobalMetrics

`func (o *RiskAuditAnalysisV2) HasGlobalMetrics() bool`

HasGlobalMetrics returns a boolean if a field has been set.

### GetFinalAnalysis

`func (o *RiskAuditAnalysisV2) GetFinalAnalysis() map[string]interface{}`

GetFinalAnalysis returns the FinalAnalysis field if non-nil, zero value otherwise.

### GetFinalAnalysisOk

`func (o *RiskAuditAnalysisV2) GetFinalAnalysisOk() (*map[string]interface{}, bool)`

GetFinalAnalysisOk returns a tuple with the FinalAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalAnalysis

`func (o *RiskAuditAnalysisV2) SetFinalAnalysis(v map[string]interface{})`

SetFinalAnalysis sets FinalAnalysis field to given value.

### HasFinalAnalysis

`func (o *RiskAuditAnalysisV2) HasFinalAnalysis() bool`

HasFinalAnalysis returns a boolean if a field has been set.

### GetFrameworks

`func (o *RiskAuditAnalysisV2) GetFrameworks() map[string]interface{}`

GetFrameworks returns the Frameworks field if non-nil, zero value otherwise.

### GetFrameworksOk

`func (o *RiskAuditAnalysisV2) GetFrameworksOk() (*map[string]interface{}, bool)`

GetFrameworksOk returns a tuple with the Frameworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworks

`func (o *RiskAuditAnalysisV2) SetFrameworks(v map[string]interface{})`

SetFrameworks sets Frameworks field to given value.

### HasFrameworks

`func (o *RiskAuditAnalysisV2) HasFrameworks() bool`

HasFrameworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


