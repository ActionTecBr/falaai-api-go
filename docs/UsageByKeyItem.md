# UsageByKeyItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyId** | **string** | API key id | 
**KeyName** | **string** | API key name | 
**TotalCredits** | **int32** | Total credits consumed by the key | 
**RequestCount** | **int32** | Number of requests | 
**LastUsed** | Pointer to **string** | ISO 8601 of last use (null if never) | [optional] 

## Methods

### NewUsageByKeyItem

`func NewUsageByKeyItem(keyId string, keyName string, totalCredits int32, requestCount int32, ) *UsageByKeyItem`

NewUsageByKeyItem instantiates a new UsageByKeyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageByKeyItemWithDefaults

`func NewUsageByKeyItemWithDefaults() *UsageByKeyItem`

NewUsageByKeyItemWithDefaults instantiates a new UsageByKeyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyId

`func (o *UsageByKeyItem) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *UsageByKeyItem) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *UsageByKeyItem) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetKeyName

`func (o *UsageByKeyItem) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *UsageByKeyItem) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *UsageByKeyItem) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetTotalCredits

`func (o *UsageByKeyItem) GetTotalCredits() int32`

GetTotalCredits returns the TotalCredits field if non-nil, zero value otherwise.

### GetTotalCreditsOk

`func (o *UsageByKeyItem) GetTotalCreditsOk() (*int32, bool)`

GetTotalCreditsOk returns a tuple with the TotalCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCredits

`func (o *UsageByKeyItem) SetTotalCredits(v int32)`

SetTotalCredits sets TotalCredits field to given value.


### GetRequestCount

`func (o *UsageByKeyItem) GetRequestCount() int32`

GetRequestCount returns the RequestCount field if non-nil, zero value otherwise.

### GetRequestCountOk

`func (o *UsageByKeyItem) GetRequestCountOk() (*int32, bool)`

GetRequestCountOk returns a tuple with the RequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCount

`func (o *UsageByKeyItem) SetRequestCount(v int32)`

SetRequestCount sets RequestCount field to given value.


### GetLastUsed

`func (o *UsageByKeyItem) GetLastUsed() string`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *UsageByKeyItem) GetLastUsedOk() (*string, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *UsageByKeyItem) SetLastUsed(v string)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *UsageByKeyItem) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


