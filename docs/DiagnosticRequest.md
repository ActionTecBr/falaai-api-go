# DiagnosticRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** | Analysis model. Always &#39;falaai-diagnostic-1&#39; | [optional] [default to "falaai-diagnostic-1"]
**Text** | Pointer to **string** | Plain transcript (fallback if dialog is empty). At least one of &#39;dialog&#39; or &#39;text&#39; required. Max 300,000 characters | [optional] [default to ""]
**Dialog** | Pointer to **string** | Diarized transcript with speaker turns. PRIMARY source. At least one of &#39;dialog&#39; or &#39;text&#39; required. Speaker labels accepted (any case): &#39;Speaker N&#39;, &#39;Interlocutor N&#39;, &#39;Hablante N&#39;, &#39;Locutor N&#39;, &#39;Orador N&#39; (space or underscore). Normalized internally to &#39;Speaker N&#39; in the response. Max 300,000 characters | [optional] [default to ""]
**AudioEvents** | Pointer to [**[]DiagnosticAudioEvent**](DiagnosticAudioEvent.md) | Detected audio events with timestamps (required when using dialog) | [optional] [default to {}]
**Language** | **string** | Transcript language. Required. Accepted: en-US, pt-BR, es-ES, es-MX, fr-FR, de-DE, it-IT, pt-PT, zh-CN, ja-JP, ko-KR, ar-SA, hi-IN, ru-RU, id-ID, tr-TR, nl-NL, pl-PL, vi-VN, th-TH, en-GB | 
**ResponseLanguage** | Pointer to **string** | Language for the analysis results (labels, categories, levels). Can differ from &#39;language&#39; (input). If omitted, uses &#39;language&#39;. Accepted: en-US, es-ES, es-MX, pt-BR, fr-FR, de-DE, it-IT, pt-PT, zh-CN, ja-JP, ko-KR, ar-SA, hi-IN, ru-RU, id-ID, tr-TR, nl-NL, pl-PL, vi-VN, th-TH, en-GB. | [optional] 
**DurationSeconds** | **float32** | Total audio duration in seconds. Required. Max 3h (10800s). | 
**ClientReferenceId** | Pointer to **string** | Optional client-supplied ID echoed verbatim in the response. Use to correlate/sync with your system. Accepted charset: [A-Za-z0-9._:-], max 128 chars. Not idempotency. | [optional] 
**CallDirection** | Pointer to **string** | Who originated the call. inbound&#x3D;client called, outbound&#x3D;company called. If omitted, the LLM infers from context. | [optional] 
**Participants** | Pointer to [**[]DiagnosticParticipant**](DiagnosticParticipant.md) | Explicit participant roles. If omitted, the LLM infers from the dialog. When provided, used as ground truth (no inference) and echoed in analysis.participants_identified. | [optional] 

## Methods

### NewDiagnosticRequest

`func NewDiagnosticRequest(language string, durationSeconds float32, ) *DiagnosticRequest`

NewDiagnosticRequest instantiates a new DiagnosticRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosticRequestWithDefaults

`func NewDiagnosticRequestWithDefaults() *DiagnosticRequest`

NewDiagnosticRequestWithDefaults instantiates a new DiagnosticRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *DiagnosticRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DiagnosticRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DiagnosticRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DiagnosticRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetText

`func (o *DiagnosticRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *DiagnosticRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *DiagnosticRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *DiagnosticRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetDialog

`func (o *DiagnosticRequest) GetDialog() string`

GetDialog returns the Dialog field if non-nil, zero value otherwise.

### GetDialogOk

`func (o *DiagnosticRequest) GetDialogOk() (*string, bool)`

GetDialogOk returns a tuple with the Dialog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialog

`func (o *DiagnosticRequest) SetDialog(v string)`

SetDialog sets Dialog field to given value.

### HasDialog

`func (o *DiagnosticRequest) HasDialog() bool`

HasDialog returns a boolean if a field has been set.

### GetAudioEvents

`func (o *DiagnosticRequest) GetAudioEvents() []DiagnosticAudioEvent`

GetAudioEvents returns the AudioEvents field if non-nil, zero value otherwise.

### GetAudioEventsOk

`func (o *DiagnosticRequest) GetAudioEventsOk() (*[]DiagnosticAudioEvent, bool)`

GetAudioEventsOk returns a tuple with the AudioEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioEvents

`func (o *DiagnosticRequest) SetAudioEvents(v []DiagnosticAudioEvent)`

SetAudioEvents sets AudioEvents field to given value.

### HasAudioEvents

`func (o *DiagnosticRequest) HasAudioEvents() bool`

HasAudioEvents returns a boolean if a field has been set.

### GetLanguage

`func (o *DiagnosticRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *DiagnosticRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *DiagnosticRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetResponseLanguage

`func (o *DiagnosticRequest) GetResponseLanguage() string`

GetResponseLanguage returns the ResponseLanguage field if non-nil, zero value otherwise.

### GetResponseLanguageOk

`func (o *DiagnosticRequest) GetResponseLanguageOk() (*string, bool)`

GetResponseLanguageOk returns a tuple with the ResponseLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLanguage

`func (o *DiagnosticRequest) SetResponseLanguage(v string)`

SetResponseLanguage sets ResponseLanguage field to given value.

### HasResponseLanguage

`func (o *DiagnosticRequest) HasResponseLanguage() bool`

HasResponseLanguage returns a boolean if a field has been set.

### GetDurationSeconds

`func (o *DiagnosticRequest) GetDurationSeconds() float32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *DiagnosticRequest) GetDurationSecondsOk() (*float32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *DiagnosticRequest) SetDurationSeconds(v float32)`

SetDurationSeconds sets DurationSeconds field to given value.


### GetClientReferenceId

`func (o *DiagnosticRequest) GetClientReferenceId() string`

GetClientReferenceId returns the ClientReferenceId field if non-nil, zero value otherwise.

### GetClientReferenceIdOk

`func (o *DiagnosticRequest) GetClientReferenceIdOk() (*string, bool)`

GetClientReferenceIdOk returns a tuple with the ClientReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientReferenceId

`func (o *DiagnosticRequest) SetClientReferenceId(v string)`

SetClientReferenceId sets ClientReferenceId field to given value.

### HasClientReferenceId

`func (o *DiagnosticRequest) HasClientReferenceId() bool`

HasClientReferenceId returns a boolean if a field has been set.

### GetCallDirection

`func (o *DiagnosticRequest) GetCallDirection() string`

GetCallDirection returns the CallDirection field if non-nil, zero value otherwise.

### GetCallDirectionOk

`func (o *DiagnosticRequest) GetCallDirectionOk() (*string, bool)`

GetCallDirectionOk returns a tuple with the CallDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallDirection

`func (o *DiagnosticRequest) SetCallDirection(v string)`

SetCallDirection sets CallDirection field to given value.

### HasCallDirection

`func (o *DiagnosticRequest) HasCallDirection() bool`

HasCallDirection returns a boolean if a field has been set.

### GetParticipants

`func (o *DiagnosticRequest) GetParticipants() []DiagnosticParticipant`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *DiagnosticRequest) GetParticipantsOk() (*[]DiagnosticParticipant, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *DiagnosticRequest) SetParticipants(v []DiagnosticParticipant)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *DiagnosticRequest) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


