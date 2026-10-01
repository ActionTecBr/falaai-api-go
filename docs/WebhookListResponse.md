# WebhookListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]WebhookItem**](WebhookItem.md) |  | 
**Page** | **int32** |  | 
**Limit** | **int32** |  | 

## Methods

### NewWebhookListResponse

`func NewWebhookListResponse(data []WebhookItem, page int32, limit int32, ) *WebhookListResponse`

NewWebhookListResponse instantiates a new WebhookListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookListResponseWithDefaults

`func NewWebhookListResponseWithDefaults() *WebhookListResponse`

NewWebhookListResponseWithDefaults instantiates a new WebhookListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WebhookListResponse) GetData() []WebhookItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookListResponse) GetDataOk() (*[]WebhookItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookListResponse) SetData(v []WebhookItem)`

SetData sets Data field to given value.


### GetPage

`func (o *WebhookListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *WebhookListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *WebhookListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetLimit

`func (o *WebhookListResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *WebhookListResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *WebhookListResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


