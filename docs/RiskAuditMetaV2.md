# RiskAuditMetaV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Analysis id | 
**Object** | Pointer to **string** | Object type | [optional] [default to "risk_audit"]
**CallDurationS** | Pointer to **float32** | Call duration (s) | [optional] 
**AnalyzedAt** | Pointer to **string** | ISO 8601 analyzed timestamp | [optional] 
**Usage** | [**RiskAuditUsageV2**](RiskAuditUsageV2.md) | Usage block | 
**ClientReferenceId** | Pointer to **string** | Echoed client reference id | [optional] 

## Methods

### NewRiskAuditMetaV2

`func NewRiskAuditMetaV2(id string, usage RiskAuditUsageV2, ) *RiskAuditMetaV2`

NewRiskAuditMetaV2 instantiates a new RiskAuditMetaV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAuditMetaV2WithDefaults

`func NewRiskAuditMetaV2WithDefaults() *RiskAuditMetaV2`

NewRiskAuditMetaV2WithDefaults instantiates a new RiskAuditMetaV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RiskAuditMetaV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskAuditMetaV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskAuditMetaV2) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *RiskAuditMetaV2) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *RiskAuditMetaV2) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *RiskAuditMetaV2) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *RiskAuditMetaV2) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCallDurationS

`func (o *RiskAuditMetaV2) GetCallDurationS() float32`

GetCallDurationS returns the CallDurationS field if non-nil, zero value otherwise.

### GetCallDurationSOk

`func (o *RiskAuditMetaV2) GetCallDurationSOk() (*float32, bool)`

GetCallDurationSOk returns a tuple with the CallDurationS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallDurationS

`func (o *RiskAuditMetaV2) SetCallDurationS(v float32)`

SetCallDurationS sets CallDurationS field to given value.

### HasCallDurationS

`func (o *RiskAuditMetaV2) HasCallDurationS() bool`

HasCallDurationS returns a boolean if a field has been set.

### GetAnalyzedAt

`func (o *RiskAuditMetaV2) GetAnalyzedAt() string`

GetAnalyzedAt returns the AnalyzedAt field if non-nil, zero value otherwise.

### GetAnalyzedAtOk

`func (o *RiskAuditMetaV2) GetAnalyzedAtOk() (*string, bool)`

GetAnalyzedAtOk returns a tuple with the AnalyzedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzedAt

`func (o *RiskAuditMetaV2) SetAnalyzedAt(v string)`

SetAnalyzedAt sets AnalyzedAt field to given value.

### HasAnalyzedAt

`func (o *RiskAuditMetaV2) HasAnalyzedAt() bool`

HasAnalyzedAt returns a boolean if a field has been set.

### GetUsage

`func (o *RiskAuditMetaV2) GetUsage() RiskAuditUsageV2`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *RiskAuditMetaV2) GetUsageOk() (*RiskAuditUsageV2, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *RiskAuditMetaV2) SetUsage(v RiskAuditUsageV2)`

SetUsage sets Usage field to given value.


### GetClientReferenceId

`func (o *RiskAuditMetaV2) GetClientReferenceId() string`

GetClientReferenceId returns the ClientReferenceId field if non-nil, zero value otherwise.

### GetClientReferenceIdOk

`func (o *RiskAuditMetaV2) GetClientReferenceIdOk() (*string, bool)`

GetClientReferenceIdOk returns a tuple with the ClientReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientReferenceId

`func (o *RiskAuditMetaV2) SetClientReferenceId(v string)`

SetClientReferenceId sets ClientReferenceId field to given value.

### HasClientReferenceId

`func (o *RiskAuditMetaV2) HasClientReferenceId() bool`

HasClientReferenceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


