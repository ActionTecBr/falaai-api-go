# WhatsappConversationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier. Prefix &#39;wc-&#39; + UUID | 
**Object** | **string** | Object type. Always &#39;conversations&#39; | 
**Usage** | [**WhatsappUsage**](WhatsappUsage.md) | Usage and processing information | 
**Conversations** | [**[]WhatsappConversation**](WhatsappConversation.md) | Segmented conversations | 
**ClientReferenceId** | Pointer to **string** | Client-supplied ID echoed verbatim (if provided) | [optional] 
**Meta** | [**WhatsappMeta**](WhatsappMeta.md) | Segmentation parameters and counts | 

## Methods

### NewWhatsappConversationsResponse

`func NewWhatsappConversationsResponse(id string, object string, usage WhatsappUsage, conversations []WhatsappConversation, meta WhatsappMeta, ) *WhatsappConversationsResponse`

NewWhatsappConversationsResponse instantiates a new WhatsappConversationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappConversationsResponseWithDefaults

`func NewWhatsappConversationsResponseWithDefaults() *WhatsappConversationsResponse`

NewWhatsappConversationsResponseWithDefaults instantiates a new WhatsappConversationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WhatsappConversationsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WhatsappConversationsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WhatsappConversationsResponse) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *WhatsappConversationsResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *WhatsappConversationsResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *WhatsappConversationsResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetUsage

`func (o *WhatsappConversationsResponse) GetUsage() WhatsappUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *WhatsappConversationsResponse) GetUsageOk() (*WhatsappUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *WhatsappConversationsResponse) SetUsage(v WhatsappUsage)`

SetUsage sets Usage field to given value.


### GetConversations

`func (o *WhatsappConversationsResponse) GetConversations() []WhatsappConversation`

GetConversations returns the Conversations field if non-nil, zero value otherwise.

### GetConversationsOk

`func (o *WhatsappConversationsResponse) GetConversationsOk() (*[]WhatsappConversation, bool)`

GetConversationsOk returns a tuple with the Conversations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversations

`func (o *WhatsappConversationsResponse) SetConversations(v []WhatsappConversation)`

SetConversations sets Conversations field to given value.


### GetClientReferenceId

`func (o *WhatsappConversationsResponse) GetClientReferenceId() string`

GetClientReferenceId returns the ClientReferenceId field if non-nil, zero value otherwise.

### GetClientReferenceIdOk

`func (o *WhatsappConversationsResponse) GetClientReferenceIdOk() (*string, bool)`

GetClientReferenceIdOk returns a tuple with the ClientReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientReferenceId

`func (o *WhatsappConversationsResponse) SetClientReferenceId(v string)`

SetClientReferenceId sets ClientReferenceId field to given value.

### HasClientReferenceId

`func (o *WhatsappConversationsResponse) HasClientReferenceId() bool`

HasClientReferenceId returns a boolean if a field has been set.

### GetMeta

`func (o *WhatsappConversationsResponse) GetMeta() WhatsappMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WhatsappConversationsResponse) GetMetaOk() (*WhatsappMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WhatsappConversationsResponse) SetMeta(v WhatsappMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


