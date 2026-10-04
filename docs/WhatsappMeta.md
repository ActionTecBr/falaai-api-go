# WhatsappMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | **string** | Uploaded file name | 
**ChatTxt** | **string** | chat.txt entry name inside the export | 
**Format** | **string** | Detected format: Android | iOS | 
**DateFormat** | **string** | Date order used | 
**Timezone** | **string** | Timezone informed | 
**Start** | **string** | Window start (ISO) | 
**End** | **string** | Window end (ISO) | 
**GapMinutes** | **float32** | Gap used to split conversations | 
**MinMessages** | **int32** | Minimum messages per conversation | 
**CharsPerMinute** | **float32** | Chars per minute used to estimate duration | 
**Turns** | **int32** | Total parsed turns | 
**SystemLines** | **int32** | System lines ignored | 
**ConversationsTotal** | **int32** | Conversations before window filter | 
**ConversationsInWindow** | **int32** | Conversations overlapping the window | 
**MonologuesDropped** | **int32** | Single-speaker conversations dropped | 
**ConversationsSelected** | **int32** | Final conversations returned | 

## Methods

### NewWhatsappMeta

`func NewWhatsappMeta(file string, chatTxt string, format string, dateFormat string, timezone string, start string, end string, gapMinutes float32, minMessages int32, charsPerMinute float32, turns int32, systemLines int32, conversationsTotal int32, conversationsInWindow int32, monologuesDropped int32, conversationsSelected int32, ) *WhatsappMeta`

NewWhatsappMeta instantiates a new WhatsappMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMetaWithDefaults

`func NewWhatsappMetaWithDefaults() *WhatsappMeta`

NewWhatsappMetaWithDefaults instantiates a new WhatsappMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *WhatsappMeta) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *WhatsappMeta) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *WhatsappMeta) SetFile(v string)`

SetFile sets File field to given value.


### GetChatTxt

`func (o *WhatsappMeta) GetChatTxt() string`

GetChatTxt returns the ChatTxt field if non-nil, zero value otherwise.

### GetChatTxtOk

`func (o *WhatsappMeta) GetChatTxtOk() (*string, bool)`

GetChatTxtOk returns a tuple with the ChatTxt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatTxt

`func (o *WhatsappMeta) SetChatTxt(v string)`

SetChatTxt sets ChatTxt field to given value.


### GetFormat

`func (o *WhatsappMeta) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *WhatsappMeta) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *WhatsappMeta) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetDateFormat

`func (o *WhatsappMeta) GetDateFormat() string`

GetDateFormat returns the DateFormat field if non-nil, zero value otherwise.

### GetDateFormatOk

`func (o *WhatsappMeta) GetDateFormatOk() (*string, bool)`

GetDateFormatOk returns a tuple with the DateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormat

`func (o *WhatsappMeta) SetDateFormat(v string)`

SetDateFormat sets DateFormat field to given value.


### GetTimezone

`func (o *WhatsappMeta) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *WhatsappMeta) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *WhatsappMeta) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetStart

`func (o *WhatsappMeta) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *WhatsappMeta) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *WhatsappMeta) SetStart(v string)`

SetStart sets Start field to given value.


### GetEnd

`func (o *WhatsappMeta) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *WhatsappMeta) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *WhatsappMeta) SetEnd(v string)`

SetEnd sets End field to given value.


### GetGapMinutes

`func (o *WhatsappMeta) GetGapMinutes() float32`

GetGapMinutes returns the GapMinutes field if non-nil, zero value otherwise.

### GetGapMinutesOk

`func (o *WhatsappMeta) GetGapMinutesOk() (*float32, bool)`

GetGapMinutesOk returns a tuple with the GapMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGapMinutes

`func (o *WhatsappMeta) SetGapMinutes(v float32)`

SetGapMinutes sets GapMinutes field to given value.


### GetMinMessages

`func (o *WhatsappMeta) GetMinMessages() int32`

GetMinMessages returns the MinMessages field if non-nil, zero value otherwise.

### GetMinMessagesOk

`func (o *WhatsappMeta) GetMinMessagesOk() (*int32, bool)`

GetMinMessagesOk returns a tuple with the MinMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMessages

`func (o *WhatsappMeta) SetMinMessages(v int32)`

SetMinMessages sets MinMessages field to given value.


### GetCharsPerMinute

`func (o *WhatsappMeta) GetCharsPerMinute() float32`

GetCharsPerMinute returns the CharsPerMinute field if non-nil, zero value otherwise.

### GetCharsPerMinuteOk

`func (o *WhatsappMeta) GetCharsPerMinuteOk() (*float32, bool)`

GetCharsPerMinuteOk returns a tuple with the CharsPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharsPerMinute

`func (o *WhatsappMeta) SetCharsPerMinute(v float32)`

SetCharsPerMinute sets CharsPerMinute field to given value.


### GetTurns

`func (o *WhatsappMeta) GetTurns() int32`

GetTurns returns the Turns field if non-nil, zero value otherwise.

### GetTurnsOk

`func (o *WhatsappMeta) GetTurnsOk() (*int32, bool)`

GetTurnsOk returns a tuple with the Turns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurns

`func (o *WhatsappMeta) SetTurns(v int32)`

SetTurns sets Turns field to given value.


### GetSystemLines

`func (o *WhatsappMeta) GetSystemLines() int32`

GetSystemLines returns the SystemLines field if non-nil, zero value otherwise.

### GetSystemLinesOk

`func (o *WhatsappMeta) GetSystemLinesOk() (*int32, bool)`

GetSystemLinesOk returns a tuple with the SystemLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemLines

`func (o *WhatsappMeta) SetSystemLines(v int32)`

SetSystemLines sets SystemLines field to given value.


### GetConversationsTotal

`func (o *WhatsappMeta) GetConversationsTotal() int32`

GetConversationsTotal returns the ConversationsTotal field if non-nil, zero value otherwise.

### GetConversationsTotalOk

`func (o *WhatsappMeta) GetConversationsTotalOk() (*int32, bool)`

GetConversationsTotalOk returns a tuple with the ConversationsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationsTotal

`func (o *WhatsappMeta) SetConversationsTotal(v int32)`

SetConversationsTotal sets ConversationsTotal field to given value.


### GetConversationsInWindow

`func (o *WhatsappMeta) GetConversationsInWindow() int32`

GetConversationsInWindow returns the ConversationsInWindow field if non-nil, zero value otherwise.

### GetConversationsInWindowOk

`func (o *WhatsappMeta) GetConversationsInWindowOk() (*int32, bool)`

GetConversationsInWindowOk returns a tuple with the ConversationsInWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationsInWindow

`func (o *WhatsappMeta) SetConversationsInWindow(v int32)`

SetConversationsInWindow sets ConversationsInWindow field to given value.


### GetMonologuesDropped

`func (o *WhatsappMeta) GetMonologuesDropped() int32`

GetMonologuesDropped returns the MonologuesDropped field if non-nil, zero value otherwise.

### GetMonologuesDroppedOk

`func (o *WhatsappMeta) GetMonologuesDroppedOk() (*int32, bool)`

GetMonologuesDroppedOk returns a tuple with the MonologuesDropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonologuesDropped

`func (o *WhatsappMeta) SetMonologuesDropped(v int32)`

SetMonologuesDropped sets MonologuesDropped field to given value.


### GetConversationsSelected

`func (o *WhatsappMeta) GetConversationsSelected() int32`

GetConversationsSelected returns the ConversationsSelected field if non-nil, zero value otherwise.

### GetConversationsSelectedOk

`func (o *WhatsappMeta) GetConversationsSelectedOk() (*int32, bool)`

GetConversationsSelectedOk returns a tuple with the ConversationsSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationsSelected

`func (o *WhatsappMeta) SetConversationsSelected(v int32)`

SetConversationsSelected sets ConversationsSelected field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


