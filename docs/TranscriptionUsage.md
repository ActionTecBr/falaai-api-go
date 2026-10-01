# TranscriptionUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioSeconds** | **float32** | Actual audio duration processed in seconds | 
**CreditsConsumed** | **int32** | Number of credits consumed in this request | 
**ProcessingMs** | **int32** | Total processing time in milliseconds | 

## Methods

### NewTranscriptionUsage

`func NewTranscriptionUsage(audioSeconds float32, creditsConsumed int32, processingMs int32, ) *TranscriptionUsage`

NewTranscriptionUsage instantiates a new TranscriptionUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscriptionUsageWithDefaults

`func NewTranscriptionUsageWithDefaults() *TranscriptionUsage`

NewTranscriptionUsageWithDefaults instantiates a new TranscriptionUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioSeconds

`func (o *TranscriptionUsage) GetAudioSeconds() float32`

GetAudioSeconds returns the AudioSeconds field if non-nil, zero value otherwise.

### GetAudioSecondsOk

`func (o *TranscriptionUsage) GetAudioSecondsOk() (*float32, bool)`

GetAudioSecondsOk returns a tuple with the AudioSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioSeconds

`func (o *TranscriptionUsage) SetAudioSeconds(v float32)`

SetAudioSeconds sets AudioSeconds field to given value.


### GetCreditsConsumed

`func (o *TranscriptionUsage) GetCreditsConsumed() int32`

GetCreditsConsumed returns the CreditsConsumed field if non-nil, zero value otherwise.

### GetCreditsConsumedOk

`func (o *TranscriptionUsage) GetCreditsConsumedOk() (*int32, bool)`

GetCreditsConsumedOk returns a tuple with the CreditsConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsConsumed

`func (o *TranscriptionUsage) SetCreditsConsumed(v int32)`

SetCreditsConsumed sets CreditsConsumed field to given value.


### GetProcessingMs

`func (o *TranscriptionUsage) GetProcessingMs() int32`

GetProcessingMs returns the ProcessingMs field if non-nil, zero value otherwise.

### GetProcessingMsOk

`func (o *TranscriptionUsage) GetProcessingMsOk() (*int32, bool)`

GetProcessingMsOk returns a tuple with the ProcessingMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingMs

`func (o *TranscriptionUsage) SetProcessingMs(v int32)`

SetProcessingMs sets ProcessingMs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


