# AudioEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** | Type of identified audio event. Ex: [riso], [suspiro], [pausa], [tosse] | 
**StartS** | **float32** | Start time of audio event in seconds | 
**EndS** | **float32** | End time of audio event in seconds | 
**DurationS** | **float32** | Event duration in seconds | 
**FormattedTimestamp** | **string** | Formatted timestamp HH:MM:SS.mmm of event start | 

## Methods

### NewAudioEvent

`func NewAudioEvent(event string, startS float32, endS float32, durationS float32, formattedTimestamp string, ) *AudioEvent`

NewAudioEvent instantiates a new AudioEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudioEventWithDefaults

`func NewAudioEventWithDefaults() *AudioEvent`

NewAudioEventWithDefaults instantiates a new AudioEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *AudioEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AudioEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AudioEvent) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetStartS

`func (o *AudioEvent) GetStartS() float32`

GetStartS returns the StartS field if non-nil, zero value otherwise.

### GetStartSOk

`func (o *AudioEvent) GetStartSOk() (*float32, bool)`

GetStartSOk returns a tuple with the StartS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartS

`func (o *AudioEvent) SetStartS(v float32)`

SetStartS sets StartS field to given value.


### GetEndS

`func (o *AudioEvent) GetEndS() float32`

GetEndS returns the EndS field if non-nil, zero value otherwise.

### GetEndSOk

`func (o *AudioEvent) GetEndSOk() (*float32, bool)`

GetEndSOk returns a tuple with the EndS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndS

`func (o *AudioEvent) SetEndS(v float32)`

SetEndS sets EndS field to given value.


### GetDurationS

`func (o *AudioEvent) GetDurationS() float32`

GetDurationS returns the DurationS field if non-nil, zero value otherwise.

### GetDurationSOk

`func (o *AudioEvent) GetDurationSOk() (*float32, bool)`

GetDurationSOk returns a tuple with the DurationS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationS

`func (o *AudioEvent) SetDurationS(v float32)`

SetDurationS sets DurationS field to given value.


### GetFormattedTimestamp

`func (o *AudioEvent) GetFormattedTimestamp() string`

GetFormattedTimestamp returns the FormattedTimestamp field if non-nil, zero value otherwise.

### GetFormattedTimestampOk

`func (o *AudioEvent) GetFormattedTimestampOk() (*string, bool)`

GetFormattedTimestampOk returns a tuple with the FormattedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTimestamp

`func (o *AudioEvent) SetFormattedTimestamp(v string)`

SetFormattedTimestamp sets FormattedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


