# RiskAuditDetectionsV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Violations** | Pointer to [**[]RiskAuditDetectionItemV2**](RiskAuditDetectionItemV2.md) | Active violations | [optional] [default to {}]
**Positives** | Pointer to [**[]RiskAuditDetectionItemV2**](RiskAuditDetectionItemV2.md) | Active positives | [optional] [default to {}]
**ClientRiskAlerts** | Pointer to **[]map[string]interface{}** | Client risk alerts | [optional] [default to {}]
**ClientBehaviorAlerts** | Pointer to **[]map[string]interface{}** | Client behavior alerts | [optional] [default to {}]
**ClientNegatives** | Pointer to [**[]RiskAuditDetectionItemV2**](RiskAuditDetectionItemV2.md) | Client negatives | [optional] [default to {}]

## Methods

### NewRiskAuditDetectionsV2

`func NewRiskAuditDetectionsV2() *RiskAuditDetectionsV2`

NewRiskAuditDetectionsV2 instantiates a new RiskAuditDetectionsV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAuditDetectionsV2WithDefaults

`func NewRiskAuditDetectionsV2WithDefaults() *RiskAuditDetectionsV2`

NewRiskAuditDetectionsV2WithDefaults instantiates a new RiskAuditDetectionsV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViolations

`func (o *RiskAuditDetectionsV2) GetViolations() []RiskAuditDetectionItemV2`

GetViolations returns the Violations field if non-nil, zero value otherwise.

### GetViolationsOk

`func (o *RiskAuditDetectionsV2) GetViolationsOk() (*[]RiskAuditDetectionItemV2, bool)`

GetViolationsOk returns a tuple with the Violations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolations

`func (o *RiskAuditDetectionsV2) SetViolations(v []RiskAuditDetectionItemV2)`

SetViolations sets Violations field to given value.

### HasViolations

`func (o *RiskAuditDetectionsV2) HasViolations() bool`

HasViolations returns a boolean if a field has been set.

### GetPositives

`func (o *RiskAuditDetectionsV2) GetPositives() []RiskAuditDetectionItemV2`

GetPositives returns the Positives field if non-nil, zero value otherwise.

### GetPositivesOk

`func (o *RiskAuditDetectionsV2) GetPositivesOk() (*[]RiskAuditDetectionItemV2, bool)`

GetPositivesOk returns a tuple with the Positives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositives

`func (o *RiskAuditDetectionsV2) SetPositives(v []RiskAuditDetectionItemV2)`

SetPositives sets Positives field to given value.

### HasPositives

`func (o *RiskAuditDetectionsV2) HasPositives() bool`

HasPositives returns a boolean if a field has been set.

### GetClientRiskAlerts

`func (o *RiskAuditDetectionsV2) GetClientRiskAlerts() []map[string]interface{}`

GetClientRiskAlerts returns the ClientRiskAlerts field if non-nil, zero value otherwise.

### GetClientRiskAlertsOk

`func (o *RiskAuditDetectionsV2) GetClientRiskAlertsOk() (*[]map[string]interface{}, bool)`

GetClientRiskAlertsOk returns a tuple with the ClientRiskAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRiskAlerts

`func (o *RiskAuditDetectionsV2) SetClientRiskAlerts(v []map[string]interface{})`

SetClientRiskAlerts sets ClientRiskAlerts field to given value.

### HasClientRiskAlerts

`func (o *RiskAuditDetectionsV2) HasClientRiskAlerts() bool`

HasClientRiskAlerts returns a boolean if a field has been set.

### GetClientBehaviorAlerts

`func (o *RiskAuditDetectionsV2) GetClientBehaviorAlerts() []map[string]interface{}`

GetClientBehaviorAlerts returns the ClientBehaviorAlerts field if non-nil, zero value otherwise.

### GetClientBehaviorAlertsOk

`func (o *RiskAuditDetectionsV2) GetClientBehaviorAlertsOk() (*[]map[string]interface{}, bool)`

GetClientBehaviorAlertsOk returns a tuple with the ClientBehaviorAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientBehaviorAlerts

`func (o *RiskAuditDetectionsV2) SetClientBehaviorAlerts(v []map[string]interface{})`

SetClientBehaviorAlerts sets ClientBehaviorAlerts field to given value.

### HasClientBehaviorAlerts

`func (o *RiskAuditDetectionsV2) HasClientBehaviorAlerts() bool`

HasClientBehaviorAlerts returns a boolean if a field has been set.

### GetClientNegatives

`func (o *RiskAuditDetectionsV2) GetClientNegatives() []RiskAuditDetectionItemV2`

GetClientNegatives returns the ClientNegatives field if non-nil, zero value otherwise.

### GetClientNegativesOk

`func (o *RiskAuditDetectionsV2) GetClientNegativesOk() (*[]RiskAuditDetectionItemV2, bool)`

GetClientNegativesOk returns a tuple with the ClientNegatives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNegatives

`func (o *RiskAuditDetectionsV2) SetClientNegatives(v []RiskAuditDetectionItemV2)`

SetClientNegatives sets ClientNegatives field to given value.

### HasClientNegatives

`func (o *RiskAuditDetectionsV2) HasClientNegatives() bool`

HasClientNegatives returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


