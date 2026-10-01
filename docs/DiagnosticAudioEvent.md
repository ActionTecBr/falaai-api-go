# DiagnosticAudioEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** | Audio event type. E.g.: [laughter], [sigh] | 
**StartS** | Pointer to **float32** | Start time in seconds | [optional] 
**EndS** | Pointer to **float32** | End time in seconds | [optional] 
**DurationS** | Pointer to **float32** | Duration in seconds | [optional] 
**FormattedTimestamp** | Pointer to **string** | Formatted timestamp (HH:MM:SS.ms) | [optional] 

## Methods

### NewDiagnosticAudioEvent

`func NewDiagnosticAudioEvent(event string, ) *DiagnosticAudioEvent`

NewDiagnosticAudioEvent instantiates a new DiagnosticAudioEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosticAudioEventWithDefaults

`func NewDiagnosticAudioEventWithDefaults() *DiagnosticAudioEvent`

NewDiagnosticAudioEventWithDefaults instantiates a new DiagnosticAudioEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *DiagnosticAudioEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *DiagnosticAudioEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *DiagnosticAudioEvent) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetStartS

`func (o *DiagnosticAudioEvent) GetStartS() float32`

GetStartS returns the StartS field if non-nil, zero value otherwise.

### GetStartSOk

`func (o *DiagnosticAudioEvent) GetStartSOk() (*float32, bool)`

GetStartSOk returns a tuple with the StartS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartS

`func (o *DiagnosticAudioEvent) SetStartS(v float32)`

SetStartS sets StartS field to given value.

### HasStartS

`func (o *DiagnosticAudioEvent) HasStartS() bool`

HasStartS returns a boolean if a field has been set.

### GetEndS

`func (o *DiagnosticAudioEvent) GetEndS() float32`

GetEndS returns the EndS field if non-nil, zero value otherwise.

### GetEndSOk

`func (o *DiagnosticAudioEvent) GetEndSOk() (*float32, bool)`

GetEndSOk returns a tuple with the EndS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndS

`func (o *DiagnosticAudioEvent) SetEndS(v float32)`

SetEndS sets EndS field to given value.

### HasEndS

`func (o *DiagnosticAudioEvent) HasEndS() bool`

HasEndS returns a boolean if a field has been set.

### GetDurationS

`func (o *DiagnosticAudioEvent) GetDurationS() float32`

GetDurationS returns the DurationS field if non-nil, zero value otherwise.

### GetDurationSOk

`func (o *DiagnosticAudioEvent) GetDurationSOk() (*float32, bool)`

GetDurationSOk returns a tuple with the DurationS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationS

`func (o *DiagnosticAudioEvent) SetDurationS(v float32)`

SetDurationS sets DurationS field to given value.

### HasDurationS

`func (o *DiagnosticAudioEvent) HasDurationS() bool`

HasDurationS returns a boolean if a field has been set.

### GetFormattedTimestamp

`func (o *DiagnosticAudioEvent) GetFormattedTimestamp() string`

GetFormattedTimestamp returns the FormattedTimestamp field if non-nil, zero value otherwise.

### GetFormattedTimestampOk

`func (o *DiagnosticAudioEvent) GetFormattedTimestampOk() (*string, bool)`

GetFormattedTimestampOk returns a tuple with the FormattedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTimestamp

`func (o *DiagnosticAudioEvent) SetFormattedTimestamp(v string)`

SetFormattedTimestamp sets FormattedTimestamp field to given value.

### HasFormattedTimestamp

`func (o *DiagnosticAudioEvent) HasFormattedTimestamp() bool`

HasFormattedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


