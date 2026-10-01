# RiskAuditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** | Analysis model. Always &#39;falaai-risk-audit-1&#39; | [optional] [default to "falaai-risk-audit-1"]
**Text** | Pointer to **string** | Plain transcript (fallback if dialog is empty). At least one of &#39;dialog&#39; or &#39;text&#39; required. Max 300,000 characters | [optional] [default to ""]
**Dialog** | Pointer to **string** | Diarized transcript with speaker turns. PRIMARY source. Speaker labels accepted (any case): &#39;Speaker N&#39;, &#39;Interlocutor N&#39;, &#39;Hablante N&#39;, &#39;Locutor N&#39;, &#39;Orador N&#39; (space or underscore). Normalized internally to &#39;Speaker N&#39; in the response. Max 300,000 characters | [optional] [default to ""]
**AudioEvents** | Pointer to [**[]DiagnosticAudioEvent**](DiagnosticAudioEvent.md) | Audio events with timestamps (correlated with turns when diarization is present) | [optional] [default to {}]
**DurationSeconds** | **float32** | Total audio duration in seconds. Required. Max 3h (10800s). | 
**Language** | **string** | Language of the transcript being analyzed. Must match the dialog/text language. Accepted: pt-BR, en-US, es-ES. | 
**ResponseLanguage** | **string** | Language for analysis results (labels, categories, levels, actions, HTML report). Can differ from &#39;language&#39;. Accepted: pt-BR, en-US, es-ES. | 
**CallDirection** | Pointer to **string** | Who originated the call. inbound&#x3D;client called, outbound&#x3D;company called. If omitted, LLM infers from context. | [optional] 
**Participants** | Pointer to [**[]Participant**](Participant.md) | Explicit participant roles. If omitted, LLM infers from dialog (Lei 17). When provided, used as ground truth — no inference. | [optional] 
**ResponseFormat** | Pointer to **string** | Response format version. v1&#x3D;legacy flat PT-BR, v2&#x3D;structured EN-US blocks. | [optional] [default to "v2"]
**ClientReferenceId** | Pointer to **string** | Optional client-supplied ID echoed verbatim in the response. Use to correlate/sync with your system. Accepted charset: [A-Za-z0-9._:-], max 128 chars. Not idempotency. | [optional] 

## Methods

### NewRiskAuditRequest

`func NewRiskAuditRequest(durationSeconds float32, language string, responseLanguage string, ) *RiskAuditRequest`

NewRiskAuditRequest instantiates a new RiskAuditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAuditRequestWithDefaults

`func NewRiskAuditRequestWithDefaults() *RiskAuditRequest`

NewRiskAuditRequestWithDefaults instantiates a new RiskAuditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *RiskAuditRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *RiskAuditRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *RiskAuditRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *RiskAuditRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetText

`func (o *RiskAuditRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *RiskAuditRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *RiskAuditRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *RiskAuditRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetDialog

`func (o *RiskAuditRequest) GetDialog() string`

GetDialog returns the Dialog field if non-nil, zero value otherwise.

### GetDialogOk

`func (o *RiskAuditRequest) GetDialogOk() (*string, bool)`

GetDialogOk returns a tuple with the Dialog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialog

`func (o *RiskAuditRequest) SetDialog(v string)`

SetDialog sets Dialog field to given value.

### HasDialog

`func (o *RiskAuditRequest) HasDialog() bool`

HasDialog returns a boolean if a field has been set.

### GetAudioEvents

`func (o *RiskAuditRequest) GetAudioEvents() []DiagnosticAudioEvent`

GetAudioEvents returns the AudioEvents field if non-nil, zero value otherwise.

### GetAudioEventsOk

`func (o *RiskAuditRequest) GetAudioEventsOk() (*[]DiagnosticAudioEvent, bool)`

GetAudioEventsOk returns a tuple with the AudioEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioEvents

`func (o *RiskAuditRequest) SetAudioEvents(v []DiagnosticAudioEvent)`

SetAudioEvents sets AudioEvents field to given value.

### HasAudioEvents

`func (o *RiskAuditRequest) HasAudioEvents() bool`

HasAudioEvents returns a boolean if a field has been set.

### GetDurationSeconds

`func (o *RiskAuditRequest) GetDurationSeconds() float32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *RiskAuditRequest) GetDurationSecondsOk() (*float32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *RiskAuditRequest) SetDurationSeconds(v float32)`

SetDurationSeconds sets DurationSeconds field to given value.


### GetLanguage

`func (o *RiskAuditRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *RiskAuditRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *RiskAuditRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetResponseLanguage

`func (o *RiskAuditRequest) GetResponseLanguage() string`

GetResponseLanguage returns the ResponseLanguage field if non-nil, zero value otherwise.

### GetResponseLanguageOk

`func (o *RiskAuditRequest) GetResponseLanguageOk() (*string, bool)`

GetResponseLanguageOk returns a tuple with the ResponseLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLanguage

`func (o *RiskAuditRequest) SetResponseLanguage(v string)`

SetResponseLanguage sets ResponseLanguage field to given value.


### GetCallDirection

`func (o *RiskAuditRequest) GetCallDirection() string`

GetCallDirection returns the CallDirection field if non-nil, zero value otherwise.

### GetCallDirectionOk

`func (o *RiskAuditRequest) GetCallDirectionOk() (*string, bool)`

GetCallDirectionOk returns a tuple with the CallDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallDirection

`func (o *RiskAuditRequest) SetCallDirection(v string)`

SetCallDirection sets CallDirection field to given value.

### HasCallDirection

`func (o *RiskAuditRequest) HasCallDirection() bool`

HasCallDirection returns a boolean if a field has been set.

### GetParticipants

`func (o *RiskAuditRequest) GetParticipants() []Participant`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *RiskAuditRequest) GetParticipantsOk() (*[]Participant, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *RiskAuditRequest) SetParticipants(v []Participant)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *RiskAuditRequest) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### GetResponseFormat

`func (o *RiskAuditRequest) GetResponseFormat() string`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *RiskAuditRequest) GetResponseFormatOk() (*string, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *RiskAuditRequest) SetResponseFormat(v string)`

SetResponseFormat sets ResponseFormat field to given value.

### HasResponseFormat

`func (o *RiskAuditRequest) HasResponseFormat() bool`

HasResponseFormat returns a boolean if a field has been set.

### GetClientReferenceId

`func (o *RiskAuditRequest) GetClientReferenceId() string`

GetClientReferenceId returns the ClientReferenceId field if non-nil, zero value otherwise.

### GetClientReferenceIdOk

`func (o *RiskAuditRequest) GetClientReferenceIdOk() (*string, bool)`

GetClientReferenceIdOk returns a tuple with the ClientReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientReferenceId

`func (o *RiskAuditRequest) SetClientReferenceId(v string)`

SetClientReferenceId sets ClientReferenceId field to given value.

### HasClientReferenceId

`func (o *RiskAuditRequest) HasClientReferenceId() bool`

HasClientReferenceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


