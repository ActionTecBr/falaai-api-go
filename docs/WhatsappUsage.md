# WhatsappUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conversations** | **int32** | Number of conversations returned | 
**Characters** | **int32** | Total characters across conversations | 
**Expressions** | **int32** | Expression tokens across conversations (sticker/emoji/gif) | 
**VoiceNotes** | **int32** | Voice notes (PTT) resolved to text | 
**ConversationCredits** | **int32** | Credits for conversations (2 each) | 
**ExpressionCredits** | **float32** | Credits for expressions (0.2 each) | 
**AudioCredits** | **float32** | Credits for audio duration (0.2 per second) | 
**CreditsConsumed** | **int32** | Total credits &#x3D; ceil(conversation_credits + expression_credits + audio_credits) | 
**ProcessingMs** | **int32** | Total processing time in milliseconds | 

## Methods

### NewWhatsappUsage

`func NewWhatsappUsage(conversations int32, characters int32, expressions int32, voiceNotes int32, conversationCredits int32, expressionCredits float32, audioCredits float32, creditsConsumed int32, processingMs int32, ) *WhatsappUsage`

NewWhatsappUsage instantiates a new WhatsappUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappUsageWithDefaults

`func NewWhatsappUsageWithDefaults() *WhatsappUsage`

NewWhatsappUsageWithDefaults instantiates a new WhatsappUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversations

`func (o *WhatsappUsage) GetConversations() int32`

GetConversations returns the Conversations field if non-nil, zero value otherwise.

### GetConversationsOk

`func (o *WhatsappUsage) GetConversationsOk() (*int32, bool)`

GetConversationsOk returns a tuple with the Conversations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversations

`func (o *WhatsappUsage) SetConversations(v int32)`

SetConversations sets Conversations field to given value.


### GetCharacters

`func (o *WhatsappUsage) GetCharacters() int32`

GetCharacters returns the Characters field if non-nil, zero value otherwise.

### GetCharactersOk

`func (o *WhatsappUsage) GetCharactersOk() (*int32, bool)`

GetCharactersOk returns a tuple with the Characters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacters

`func (o *WhatsappUsage) SetCharacters(v int32)`

SetCharacters sets Characters field to given value.


### GetExpressions

`func (o *WhatsappUsage) GetExpressions() int32`

GetExpressions returns the Expressions field if non-nil, zero value otherwise.

### GetExpressionsOk

`func (o *WhatsappUsage) GetExpressionsOk() (*int32, bool)`

GetExpressionsOk returns a tuple with the Expressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressions

`func (o *WhatsappUsage) SetExpressions(v int32)`

SetExpressions sets Expressions field to given value.


### GetVoiceNotes

`func (o *WhatsappUsage) GetVoiceNotes() int32`

GetVoiceNotes returns the VoiceNotes field if non-nil, zero value otherwise.

### GetVoiceNotesOk

`func (o *WhatsappUsage) GetVoiceNotesOk() (*int32, bool)`

GetVoiceNotesOk returns a tuple with the VoiceNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceNotes

`func (o *WhatsappUsage) SetVoiceNotes(v int32)`

SetVoiceNotes sets VoiceNotes field to given value.


### GetConversationCredits

`func (o *WhatsappUsage) GetConversationCredits() int32`

GetConversationCredits returns the ConversationCredits field if non-nil, zero value otherwise.

### GetConversationCreditsOk

`func (o *WhatsappUsage) GetConversationCreditsOk() (*int32, bool)`

GetConversationCreditsOk returns a tuple with the ConversationCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationCredits

`func (o *WhatsappUsage) SetConversationCredits(v int32)`

SetConversationCredits sets ConversationCredits field to given value.


### GetExpressionCredits

`func (o *WhatsappUsage) GetExpressionCredits() float32`

GetExpressionCredits returns the ExpressionCredits field if non-nil, zero value otherwise.

### GetExpressionCreditsOk

`func (o *WhatsappUsage) GetExpressionCreditsOk() (*float32, bool)`

GetExpressionCreditsOk returns a tuple with the ExpressionCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressionCredits

`func (o *WhatsappUsage) SetExpressionCredits(v float32)`

SetExpressionCredits sets ExpressionCredits field to given value.


### GetAudioCredits

`func (o *WhatsappUsage) GetAudioCredits() float32`

GetAudioCredits returns the AudioCredits field if non-nil, zero value otherwise.

### GetAudioCreditsOk

`func (o *WhatsappUsage) GetAudioCreditsOk() (*float32, bool)`

GetAudioCreditsOk returns a tuple with the AudioCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCredits

`func (o *WhatsappUsage) SetAudioCredits(v float32)`

SetAudioCredits sets AudioCredits field to given value.


### GetCreditsConsumed

`func (o *WhatsappUsage) GetCreditsConsumed() int32`

GetCreditsConsumed returns the CreditsConsumed field if non-nil, zero value otherwise.

### GetCreditsConsumedOk

`func (o *WhatsappUsage) GetCreditsConsumedOk() (*int32, bool)`

GetCreditsConsumedOk returns a tuple with the CreditsConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsConsumed

`func (o *WhatsappUsage) SetCreditsConsumed(v int32)`

SetCreditsConsumed sets CreditsConsumed field to given value.


### GetProcessingMs

`func (o *WhatsappUsage) GetProcessingMs() int32`

GetProcessingMs returns the ProcessingMs field if non-nil, zero value otherwise.

### GetProcessingMsOk

`func (o *WhatsappUsage) GetProcessingMsOk() (*int32, bool)`

GetProcessingMsOk returns a tuple with the ProcessingMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingMs

`func (o *WhatsappUsage) SetProcessingMs(v int32)`

SetProcessingMs sets ProcessingMs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


