# AudioInputMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationS** | **float32** | Exact audio duration sent in seconds | 
**OriginalFormat** | **string** | Original file format (wav, mp3, ogg, etc) | 
**Codec** | **string** | Audio codec sent | 
**SampleRate** | **int32** | Audio sample rate in Hz | 
**Channels** | **int32** | Number of channels (1&#x3D;mono, 2&#x3D;stereo) | 

## Methods

### NewAudioInputMeta

`func NewAudioInputMeta(durationS float32, originalFormat string, codec string, sampleRate int32, channels int32, ) *AudioInputMeta`

NewAudioInputMeta instantiates a new AudioInputMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudioInputMetaWithDefaults

`func NewAudioInputMetaWithDefaults() *AudioInputMeta`

NewAudioInputMetaWithDefaults instantiates a new AudioInputMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationS

`func (o *AudioInputMeta) GetDurationS() float32`

GetDurationS returns the DurationS field if non-nil, zero value otherwise.

### GetDurationSOk

`func (o *AudioInputMeta) GetDurationSOk() (*float32, bool)`

GetDurationSOk returns a tuple with the DurationS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationS

`func (o *AudioInputMeta) SetDurationS(v float32)`

SetDurationS sets DurationS field to given value.


### GetOriginalFormat

`func (o *AudioInputMeta) GetOriginalFormat() string`

GetOriginalFormat returns the OriginalFormat field if non-nil, zero value otherwise.

### GetOriginalFormatOk

`func (o *AudioInputMeta) GetOriginalFormatOk() (*string, bool)`

GetOriginalFormatOk returns a tuple with the OriginalFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFormat

`func (o *AudioInputMeta) SetOriginalFormat(v string)`

SetOriginalFormat sets OriginalFormat field to given value.


### GetCodec

`func (o *AudioInputMeta) GetCodec() string`

GetCodec returns the Codec field if non-nil, zero value otherwise.

### GetCodecOk

`func (o *AudioInputMeta) GetCodecOk() (*string, bool)`

GetCodecOk returns a tuple with the Codec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodec

`func (o *AudioInputMeta) SetCodec(v string)`

SetCodec sets Codec field to given value.


### GetSampleRate

`func (o *AudioInputMeta) GetSampleRate() int32`

GetSampleRate returns the SampleRate field if non-nil, zero value otherwise.

### GetSampleRateOk

`func (o *AudioInputMeta) GetSampleRateOk() (*int32, bool)`

GetSampleRateOk returns a tuple with the SampleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleRate

`func (o *AudioInputMeta) SetSampleRate(v int32)`

SetSampleRate sets SampleRate field to given value.


### GetChannels

`func (o *AudioInputMeta) GetChannels() int32`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *AudioInputMeta) GetChannelsOk() (*int32, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *AudioInputMeta) SetChannels(v int32)`

SetChannels sets Channels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


