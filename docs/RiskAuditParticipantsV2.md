# RiskAuditParticipantsV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identified** | Pointer to [**[]RiskAuditParticipantV2**](RiskAuditParticipantV2.md) | Identified participants | [optional] [default to {}]
**CallDirection** | Pointer to **string** | inbound/outbound | [optional] 
**RoleInferenceReliable** | Pointer to **bool** | Role inference reliability | [optional] [default to true]
**IdentificationStatus** | Pointer to **string** | Identification status | [optional] [default to "none"]
**UnidentifiedItemsCount** | Pointer to **int32** | Unidentified items count | [optional] [default to 0]

## Methods

### NewRiskAuditParticipantsV2

`func NewRiskAuditParticipantsV2() *RiskAuditParticipantsV2`

NewRiskAuditParticipantsV2 instantiates a new RiskAuditParticipantsV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAuditParticipantsV2WithDefaults

`func NewRiskAuditParticipantsV2WithDefaults() *RiskAuditParticipantsV2`

NewRiskAuditParticipantsV2WithDefaults instantiates a new RiskAuditParticipantsV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentified

`func (o *RiskAuditParticipantsV2) GetIdentified() []RiskAuditParticipantV2`

GetIdentified returns the Identified field if non-nil, zero value otherwise.

### GetIdentifiedOk

`func (o *RiskAuditParticipantsV2) GetIdentifiedOk() (*[]RiskAuditParticipantV2, bool)`

GetIdentifiedOk returns a tuple with the Identified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentified

`func (o *RiskAuditParticipantsV2) SetIdentified(v []RiskAuditParticipantV2)`

SetIdentified sets Identified field to given value.

### HasIdentified

`func (o *RiskAuditParticipantsV2) HasIdentified() bool`

HasIdentified returns a boolean if a field has been set.

### GetCallDirection

`func (o *RiskAuditParticipantsV2) GetCallDirection() string`

GetCallDirection returns the CallDirection field if non-nil, zero value otherwise.

### GetCallDirectionOk

`func (o *RiskAuditParticipantsV2) GetCallDirectionOk() (*string, bool)`

GetCallDirectionOk returns a tuple with the CallDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallDirection

`func (o *RiskAuditParticipantsV2) SetCallDirection(v string)`

SetCallDirection sets CallDirection field to given value.

### HasCallDirection

`func (o *RiskAuditParticipantsV2) HasCallDirection() bool`

HasCallDirection returns a boolean if a field has been set.

### GetRoleInferenceReliable

`func (o *RiskAuditParticipantsV2) GetRoleInferenceReliable() bool`

GetRoleInferenceReliable returns the RoleInferenceReliable field if non-nil, zero value otherwise.

### GetRoleInferenceReliableOk

`func (o *RiskAuditParticipantsV2) GetRoleInferenceReliableOk() (*bool, bool)`

GetRoleInferenceReliableOk returns a tuple with the RoleInferenceReliable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleInferenceReliable

`func (o *RiskAuditParticipantsV2) SetRoleInferenceReliable(v bool)`

SetRoleInferenceReliable sets RoleInferenceReliable field to given value.

### HasRoleInferenceReliable

`func (o *RiskAuditParticipantsV2) HasRoleInferenceReliable() bool`

HasRoleInferenceReliable returns a boolean if a field has been set.

### GetIdentificationStatus

`func (o *RiskAuditParticipantsV2) GetIdentificationStatus() string`

GetIdentificationStatus returns the IdentificationStatus field if non-nil, zero value otherwise.

### GetIdentificationStatusOk

`func (o *RiskAuditParticipantsV2) GetIdentificationStatusOk() (*string, bool)`

GetIdentificationStatusOk returns a tuple with the IdentificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationStatus

`func (o *RiskAuditParticipantsV2) SetIdentificationStatus(v string)`

SetIdentificationStatus sets IdentificationStatus field to given value.

### HasIdentificationStatus

`func (o *RiskAuditParticipantsV2) HasIdentificationStatus() bool`

HasIdentificationStatus returns a boolean if a field has been set.

### GetUnidentifiedItemsCount

`func (o *RiskAuditParticipantsV2) GetUnidentifiedItemsCount() int32`

GetUnidentifiedItemsCount returns the UnidentifiedItemsCount field if non-nil, zero value otherwise.

### GetUnidentifiedItemsCountOk

`func (o *RiskAuditParticipantsV2) GetUnidentifiedItemsCountOk() (*int32, bool)`

GetUnidentifiedItemsCountOk returns a tuple with the UnidentifiedItemsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnidentifiedItemsCount

`func (o *RiskAuditParticipantsV2) SetUnidentifiedItemsCount(v int32)`

SetUnidentifiedItemsCount sets UnidentifiedItemsCount field to given value.

### HasUnidentifiedItemsCount

`func (o *RiskAuditParticipantsV2) HasUnidentifiedItemsCount() bool`

HasUnidentifiedItemsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


