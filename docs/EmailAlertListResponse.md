# EmailAlertListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]EmailAlertItem**](EmailAlertItem.md) |  | 
**Page** | **int32** |  | 
**Limit** | **int32** |  | 

## Methods

### NewEmailAlertListResponse

`func NewEmailAlertListResponse(data []EmailAlertItem, page int32, limit int32, ) *EmailAlertListResponse`

NewEmailAlertListResponse instantiates a new EmailAlertListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailAlertListResponseWithDefaults

`func NewEmailAlertListResponseWithDefaults() *EmailAlertListResponse`

NewEmailAlertListResponseWithDefaults instantiates a new EmailAlertListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EmailAlertListResponse) GetData() []EmailAlertItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EmailAlertListResponse) GetDataOk() (*[]EmailAlertItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EmailAlertListResponse) SetData(v []EmailAlertItem)`

SetData sets Data field to given value.


### GetPage

`func (o *EmailAlertListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *EmailAlertListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *EmailAlertListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetLimit

`func (o *EmailAlertListResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *EmailAlertListResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *EmailAlertListResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


