# UsageLogItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Usage log entry id | 
**Endpoint** | **string** | Endpoint called | 
**CreditsCost** | **int32** | Credits consumed | 
**Status** | **string** | Result status | 
**ErrorsCount** | **int32** | Errors count | 
**CreatedAt** | **string** | ISO 8601 timestamp | 

## Methods

### NewUsageLogItem

`func NewUsageLogItem(id string, endpoint string, creditsCost int32, status string, errorsCount int32, createdAt string, ) *UsageLogItem`

NewUsageLogItem instantiates a new UsageLogItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageLogItemWithDefaults

`func NewUsageLogItemWithDefaults() *UsageLogItem`

NewUsageLogItemWithDefaults instantiates a new UsageLogItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UsageLogItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsageLogItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsageLogItem) SetId(v string)`

SetId sets Id field to given value.


### GetEndpoint

`func (o *UsageLogItem) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *UsageLogItem) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *UsageLogItem) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetCreditsCost

`func (o *UsageLogItem) GetCreditsCost() int32`

GetCreditsCost returns the CreditsCost field if non-nil, zero value otherwise.

### GetCreditsCostOk

`func (o *UsageLogItem) GetCreditsCostOk() (*int32, bool)`

GetCreditsCostOk returns a tuple with the CreditsCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsCost

`func (o *UsageLogItem) SetCreditsCost(v int32)`

SetCreditsCost sets CreditsCost field to given value.


### GetStatus

`func (o *UsageLogItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UsageLogItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UsageLogItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetErrorsCount

`func (o *UsageLogItem) GetErrorsCount() int32`

GetErrorsCount returns the ErrorsCount field if non-nil, zero value otherwise.

### GetErrorsCountOk

`func (o *UsageLogItem) GetErrorsCountOk() (*int32, bool)`

GetErrorsCountOk returns a tuple with the ErrorsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsCount

`func (o *UsageLogItem) SetErrorsCount(v int32)`

SetErrorsCount sets ErrorsCount field to given value.


### GetCreatedAt

`func (o *UsageLogItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UsageLogItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UsageLogItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


