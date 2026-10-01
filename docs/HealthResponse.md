# HealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Overall API status | 
**Version** | **string** | Current API version | 
**UptimeSeconds** | **int32** | Uptime in seconds | 
**Database** | **bool** | Database connection status | 
**Phase** | **string** | Development phase | 
**LaunchDate** | **string** | Expected public launch date | 

## Methods

### NewHealthResponse

`func NewHealthResponse(status string, version string, uptimeSeconds int32, database bool, phase string, launchDate string, ) *HealthResponse`

NewHealthResponse instantiates a new HealthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthResponseWithDefaults

`func NewHealthResponseWithDefaults() *HealthResponse`

NewHealthResponseWithDefaults instantiates a new HealthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HealthResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVersion

`func (o *HealthResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HealthResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HealthResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetUptimeSeconds

`func (o *HealthResponse) GetUptimeSeconds() int32`

GetUptimeSeconds returns the UptimeSeconds field if non-nil, zero value otherwise.

### GetUptimeSecondsOk

`func (o *HealthResponse) GetUptimeSecondsOk() (*int32, bool)`

GetUptimeSecondsOk returns a tuple with the UptimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimeSeconds

`func (o *HealthResponse) SetUptimeSeconds(v int32)`

SetUptimeSeconds sets UptimeSeconds field to given value.


### GetDatabase

`func (o *HealthResponse) GetDatabase() bool`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *HealthResponse) GetDatabaseOk() (*bool, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *HealthResponse) SetDatabase(v bool)`

SetDatabase sets Database field to given value.


### GetPhase

`func (o *HealthResponse) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *HealthResponse) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *HealthResponse) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetLaunchDate

`func (o *HealthResponse) GetLaunchDate() string`

GetLaunchDate returns the LaunchDate field if non-nil, zero value otherwise.

### GetLaunchDateOk

`func (o *HealthResponse) GetLaunchDateOk() (*string, bool)`

GetLaunchDateOk returns a tuple with the LaunchDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchDate

`func (o *HealthResponse) SetLaunchDate(v string)`

SetLaunchDate sets LaunchDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


