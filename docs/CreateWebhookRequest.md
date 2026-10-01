# CreateWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Nome identificador do webhook | 
**Url** | **string** | URL HTTPS que recebera POST com HMAC FalaAI-Signature | 
**Events** | [**[]WebhookEvent**](WebhookEvent.md) | Eventos subscritos (10 alertas) | 
**RetryEnabled** | Pointer to **bool** | Retry exponencial 5 tentativas quando true (false&#x3D;1 tentativa) | [optional] [default to false]

## Methods

### NewCreateWebhookRequest

`func NewCreateWebhookRequest(name string, url string, events []WebhookEvent, ) *CreateWebhookRequest`

NewCreateWebhookRequest instantiates a new CreateWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebhookRequestWithDefaults

`func NewCreateWebhookRequestWithDefaults() *CreateWebhookRequest`

NewCreateWebhookRequestWithDefaults instantiates a new CreateWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateWebhookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWebhookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWebhookRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *CreateWebhookRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateWebhookRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateWebhookRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEvents

`func (o *CreateWebhookRequest) GetEvents() []WebhookEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CreateWebhookRequest) GetEventsOk() (*[]WebhookEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CreateWebhookRequest) SetEvents(v []WebhookEvent)`

SetEvents sets Events field to given value.


### GetRetryEnabled

`func (o *CreateWebhookRequest) GetRetryEnabled() bool`

GetRetryEnabled returns the RetryEnabled field if non-nil, zero value otherwise.

### GetRetryEnabledOk

`func (o *CreateWebhookRequest) GetRetryEnabledOk() (*bool, bool)`

GetRetryEnabledOk returns a tuple with the RetryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryEnabled

`func (o *CreateWebhookRequest) SetRetryEnabled(v bool)`

SetRetryEnabled sets RetryEnabled field to given value.

### HasRetryEnabled

`func (o *CreateWebhookRequest) HasRetryEnabled() bool`

HasRetryEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


