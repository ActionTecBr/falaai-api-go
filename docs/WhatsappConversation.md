# WhatsappConversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversationId** | **string** | Conversation identifier in the batch | 
**FirstAt** | **string** | Real start (wall-clock, ISO) | 
**LastAt** | **string** | Real end (wall-clock, ISO) | 
**DurationSeconds** | **float32** | (last - first) + last turn duration | 
**Speakers** | [**[]WhatsappSpeaker**](WhatsappSpeaker.md) | Speakers of THIS conversation (dynamic) | 
**Dialog** | **string** | Lines &#39;Speaker N: [HH:MM:SS.mmm - HH:MM:SS.mmm] text&#39; (real offset) | 
**MessageCount** | **int32** | Number of messages | 
**Characters** | **int32** | Total characters of the conversation | 

## Methods

### NewWhatsappConversation

`func NewWhatsappConversation(conversationId string, firstAt string, lastAt string, durationSeconds float32, speakers []WhatsappSpeaker, dialog string, messageCount int32, characters int32, ) *WhatsappConversation`

NewWhatsappConversation instantiates a new WhatsappConversation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappConversationWithDefaults

`func NewWhatsappConversationWithDefaults() *WhatsappConversation`

NewWhatsappConversationWithDefaults instantiates a new WhatsappConversation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversationId

`func (o *WhatsappConversation) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *WhatsappConversation) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *WhatsappConversation) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.


### GetFirstAt

`func (o *WhatsappConversation) GetFirstAt() string`

GetFirstAt returns the FirstAt field if non-nil, zero value otherwise.

### GetFirstAtOk

`func (o *WhatsappConversation) GetFirstAtOk() (*string, bool)`

GetFirstAtOk returns a tuple with the FirstAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAt

`func (o *WhatsappConversation) SetFirstAt(v string)`

SetFirstAt sets FirstAt field to given value.


### GetLastAt

`func (o *WhatsappConversation) GetLastAt() string`

GetLastAt returns the LastAt field if non-nil, zero value otherwise.

### GetLastAtOk

`func (o *WhatsappConversation) GetLastAtOk() (*string, bool)`

GetLastAtOk returns a tuple with the LastAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAt

`func (o *WhatsappConversation) SetLastAt(v string)`

SetLastAt sets LastAt field to given value.


### GetDurationSeconds

`func (o *WhatsappConversation) GetDurationSeconds() float32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *WhatsappConversation) GetDurationSecondsOk() (*float32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *WhatsappConversation) SetDurationSeconds(v float32)`

SetDurationSeconds sets DurationSeconds field to given value.


### GetSpeakers

`func (o *WhatsappConversation) GetSpeakers() []WhatsappSpeaker`

GetSpeakers returns the Speakers field if non-nil, zero value otherwise.

### GetSpeakersOk

`func (o *WhatsappConversation) GetSpeakersOk() (*[]WhatsappSpeaker, bool)`

GetSpeakersOk returns a tuple with the Speakers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeakers

`func (o *WhatsappConversation) SetSpeakers(v []WhatsappSpeaker)`

SetSpeakers sets Speakers field to given value.


### GetDialog

`func (o *WhatsappConversation) GetDialog() string`

GetDialog returns the Dialog field if non-nil, zero value otherwise.

### GetDialogOk

`func (o *WhatsappConversation) GetDialogOk() (*string, bool)`

GetDialogOk returns a tuple with the Dialog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialog

`func (o *WhatsappConversation) SetDialog(v string)`

SetDialog sets Dialog field to given value.


### GetMessageCount

`func (o *WhatsappConversation) GetMessageCount() int32`

GetMessageCount returns the MessageCount field if non-nil, zero value otherwise.

### GetMessageCountOk

`func (o *WhatsappConversation) GetMessageCountOk() (*int32, bool)`

GetMessageCountOk returns a tuple with the MessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCount

`func (o *WhatsappConversation) SetMessageCount(v int32)`

SetMessageCount sets MessageCount field to given value.


### GetCharacters

`func (o *WhatsappConversation) GetCharacters() int32`

GetCharacters returns the Characters field if non-nil, zero value otherwise.

### GetCharactersOk

`func (o *WhatsappConversation) GetCharactersOk() (*int32, bool)`

GetCharactersOk returns a tuple with the Characters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacters

`func (o *WhatsappConversation) SetCharacters(v int32)`

SetCharacters sets Characters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


