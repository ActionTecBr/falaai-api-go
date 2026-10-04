# WhatsappUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conversations** | **int32** | Number of conversations returned | 
**Characters** | **int32** | Total characters across conversations | 
**CreditsConsumed** | **int32** | Credits consumed (1 per conversation) | 
**ProcessingMs** | **int32** | Total processing time in milliseconds | 

## Methods

### NewWhatsappUsage

`func NewWhatsappUsage(conversations int32, characters int32, creditsConsumed int32, processingMs int32, ) *WhatsappUsage`

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


