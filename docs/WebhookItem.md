# WebhookItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Webhook id | 
**UserId** | **string** | Owner user id | 
**Name** | **string** | Webhook name | 
**Url** | **string** | Destination URL | 
**Secret** | **string** | HMAC signing secret | 
**Events** | **[]string** | Subscribed events | 
**Active** | **bool** | Is active | 
**RetryEnabled** | **bool** | Retry enabled | 
**LastDeliveryAt** | Pointer to **string** | ISO 8601 of last delivery | [optional] 
**LastStatus** | Pointer to **int32** | Last HTTP status delivered | [optional] 
**FailureCount** | Pointer to **int32** | Consecutive failures | [optional] [default to 0]
**CreatedAt** | **string** | ISO 8601 created | 
**UpdatedAt** | **string** | ISO 8601 updated | 

## Methods

### NewWebhookItem

`func NewWebhookItem(id string, userId string, name string, url string, secret string, events []string, active bool, retryEnabled bool, createdAt string, updatedAt string, ) *WebhookItem`

NewWebhookItem instantiates a new WebhookItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookItemWithDefaults

`func NewWebhookItemWithDefaults() *WebhookItem`

NewWebhookItemWithDefaults instantiates a new WebhookItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookItem) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *WebhookItem) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WebhookItem) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WebhookItem) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetName

`func (o *WebhookItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookItem) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *WebhookItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookItem) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSecret

`func (o *WebhookItem) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *WebhookItem) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *WebhookItem) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetEvents

`func (o *WebhookItem) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookItem) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookItem) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetActive

`func (o *WebhookItem) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WebhookItem) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WebhookItem) SetActive(v bool)`

SetActive sets Active field to given value.


### GetRetryEnabled

`func (o *WebhookItem) GetRetryEnabled() bool`

GetRetryEnabled returns the RetryEnabled field if non-nil, zero value otherwise.

### GetRetryEnabledOk

`func (o *WebhookItem) GetRetryEnabledOk() (*bool, bool)`

GetRetryEnabledOk returns a tuple with the RetryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryEnabled

`func (o *WebhookItem) SetRetryEnabled(v bool)`

SetRetryEnabled sets RetryEnabled field to given value.


### GetLastDeliveryAt

`func (o *WebhookItem) GetLastDeliveryAt() string`

GetLastDeliveryAt returns the LastDeliveryAt field if non-nil, zero value otherwise.

### GetLastDeliveryAtOk

`func (o *WebhookItem) GetLastDeliveryAtOk() (*string, bool)`

GetLastDeliveryAtOk returns a tuple with the LastDeliveryAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeliveryAt

`func (o *WebhookItem) SetLastDeliveryAt(v string)`

SetLastDeliveryAt sets LastDeliveryAt field to given value.

### HasLastDeliveryAt

`func (o *WebhookItem) HasLastDeliveryAt() bool`

HasLastDeliveryAt returns a boolean if a field has been set.

### GetLastStatus

`func (o *WebhookItem) GetLastStatus() int32`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *WebhookItem) GetLastStatusOk() (*int32, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *WebhookItem) SetLastStatus(v int32)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *WebhookItem) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### GetFailureCount

`func (o *WebhookItem) GetFailureCount() int32`

GetFailureCount returns the FailureCount field if non-nil, zero value otherwise.

### GetFailureCountOk

`func (o *WebhookItem) GetFailureCountOk() (*int32, bool)`

GetFailureCountOk returns a tuple with the FailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCount

`func (o *WebhookItem) SetFailureCount(v int32)`

SetFailureCount sets FailureCount field to given value.

### HasFailureCount

`func (o *WebhookItem) HasFailureCount() bool`

HasFailureCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WebhookItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WebhookItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WebhookItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WebhookItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


