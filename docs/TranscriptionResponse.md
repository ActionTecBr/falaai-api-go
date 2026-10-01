# TranscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique transcription identifier. Prefix &#39;tr-&#39; followed by UUID | 
**Object** | **string** | Returned object type. Always &#39;transcription&#39; | 
**Model** | **string** | Model used for transcription. Ex: &#39;falaai-transcribe-1&#39; | 
**Filename** | **string** | Original audio file name uploaded | 
**ProcessedAt** | **string** | Processing datetime in ISO 8601 UTC format | 
**Usage** | [**TranscriptionUsage**](TranscriptionUsage.md) | Usage and processing information | 
**Language** | **string** | ISO 639-3 language code detected in audio. Ex: &#39;por&#39; (Portuguese), &#39;eng&#39; (English), &#39;spa&#39; (Spanish) | 
**LanguageConfidence** | Pointer to **float32** | Language detection confidence level (0.0 to 1.0). Higher is more reliable | [optional] 
**DurationSeconds** | **float32** | Total audio duration in seconds | 
**Text** | **string** | Full transcription as plain text, including audio events in brackets | 
**Dialog** | **string** | Turn-by-turn formatted transcript with speaker identification and start/end timestamps | 
**AudioEvents** | [**[]AudioEvent**](AudioEvent.md) | List of detected audio events (laughs, sighs, pauses, etc) with timestamps and duration | 
**EventTypes** | **[]string** | Unique audio event types found in transcription, alphabetically sorted | 
**WordCount** | **int32** | Total number of recognized words in transcription | 
**Input** | [**AudioInputMeta**](AudioInputMeta.md) | Metadados do arquivo de audio enviado (duracao, formato, codec, sample rate, canais) | 
**ClientReferenceId** | Pointer to **string** | Client-supplied ID echoed verbatim (if provided in request) | [optional] 

## Methods

### NewTranscriptionResponse

`func NewTranscriptionResponse(id string, object string, model string, filename string, processedAt string, usage TranscriptionUsage, language string, durationSeconds float32, text string, dialog string, audioEvents []AudioEvent, eventTypes []string, wordCount int32, input AudioInputMeta, ) *TranscriptionResponse`

NewTranscriptionResponse instantiates a new TranscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscriptionResponseWithDefaults

`func NewTranscriptionResponseWithDefaults() *TranscriptionResponse`

NewTranscriptionResponseWithDefaults instantiates a new TranscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TranscriptionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TranscriptionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TranscriptionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *TranscriptionResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *TranscriptionResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *TranscriptionResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetModel

`func (o *TranscriptionResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TranscriptionResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TranscriptionResponse) SetModel(v string)`

SetModel sets Model field to given value.


### GetFilename

`func (o *TranscriptionResponse) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *TranscriptionResponse) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *TranscriptionResponse) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetProcessedAt

`func (o *TranscriptionResponse) GetProcessedAt() string`

GetProcessedAt returns the ProcessedAt field if non-nil, zero value otherwise.

### GetProcessedAtOk

`func (o *TranscriptionResponse) GetProcessedAtOk() (*string, bool)`

GetProcessedAtOk returns a tuple with the ProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAt

`func (o *TranscriptionResponse) SetProcessedAt(v string)`

SetProcessedAt sets ProcessedAt field to given value.


### GetUsage

`func (o *TranscriptionResponse) GetUsage() TranscriptionUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *TranscriptionResponse) GetUsageOk() (*TranscriptionUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *TranscriptionResponse) SetUsage(v TranscriptionUsage)`

SetUsage sets Usage field to given value.


### GetLanguage

`func (o *TranscriptionResponse) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *TranscriptionResponse) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *TranscriptionResponse) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetLanguageConfidence

`func (o *TranscriptionResponse) GetLanguageConfidence() float32`

GetLanguageConfidence returns the LanguageConfidence field if non-nil, zero value otherwise.

### GetLanguageConfidenceOk

`func (o *TranscriptionResponse) GetLanguageConfidenceOk() (*float32, bool)`

GetLanguageConfidenceOk returns a tuple with the LanguageConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageConfidence

`func (o *TranscriptionResponse) SetLanguageConfidence(v float32)`

SetLanguageConfidence sets LanguageConfidence field to given value.

### HasLanguageConfidence

`func (o *TranscriptionResponse) HasLanguageConfidence() bool`

HasLanguageConfidence returns a boolean if a field has been set.

### GetDurationSeconds

`func (o *TranscriptionResponse) GetDurationSeconds() float32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *TranscriptionResponse) GetDurationSecondsOk() (*float32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *TranscriptionResponse) SetDurationSeconds(v float32)`

SetDurationSeconds sets DurationSeconds field to given value.


### GetText

`func (o *TranscriptionResponse) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TranscriptionResponse) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TranscriptionResponse) SetText(v string)`

SetText sets Text field to given value.


### GetDialog

`func (o *TranscriptionResponse) GetDialog() string`

GetDialog returns the Dialog field if non-nil, zero value otherwise.

### GetDialogOk

`func (o *TranscriptionResponse) GetDialogOk() (*string, bool)`

GetDialogOk returns a tuple with the Dialog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialog

`func (o *TranscriptionResponse) SetDialog(v string)`

SetDialog sets Dialog field to given value.


### GetAudioEvents

`func (o *TranscriptionResponse) GetAudioEvents() []AudioEvent`

GetAudioEvents returns the AudioEvents field if non-nil, zero value otherwise.

### GetAudioEventsOk

`func (o *TranscriptionResponse) GetAudioEventsOk() (*[]AudioEvent, bool)`

GetAudioEventsOk returns a tuple with the AudioEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioEvents

`func (o *TranscriptionResponse) SetAudioEvents(v []AudioEvent)`

SetAudioEvents sets AudioEvents field to given value.


### GetEventTypes

`func (o *TranscriptionResponse) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *TranscriptionResponse) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *TranscriptionResponse) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.


### GetWordCount

`func (o *TranscriptionResponse) GetWordCount() int32`

GetWordCount returns the WordCount field if non-nil, zero value otherwise.

### GetWordCountOk

`func (o *TranscriptionResponse) GetWordCountOk() (*int32, bool)`

GetWordCountOk returns a tuple with the WordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordCount

`func (o *TranscriptionResponse) SetWordCount(v int32)`

SetWordCount sets WordCount field to given value.


### GetInput

`func (o *TranscriptionResponse) GetInput() AudioInputMeta`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *TranscriptionResponse) GetInputOk() (*AudioInputMeta, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *TranscriptionResponse) SetInput(v AudioInputMeta)`

SetInput sets Input field to given value.


### GetClientReferenceId

`func (o *TranscriptionResponse) GetClientReferenceId() string`

GetClientReferenceId returns the ClientReferenceId field if non-nil, zero value otherwise.

### GetClientReferenceIdOk

`func (o *TranscriptionResponse) GetClientReferenceIdOk() (*string, bool)`

GetClientReferenceIdOk returns a tuple with the ClientReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientReferenceId

`func (o *TranscriptionResponse) SetClientReferenceId(v string)`

SetClientReferenceId sets ClientReferenceId field to given value.

### HasClientReferenceId

`func (o *TranscriptionResponse) HasClientReferenceId() bool`

HasClientReferenceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


