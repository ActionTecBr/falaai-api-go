# UpdateEmailAlertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Events** | Pointer to [**[]EmailEvent**](EmailEvent.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateEmailAlertRequest

`func NewUpdateEmailAlertRequest() *UpdateEmailAlertRequest`

NewUpdateEmailAlertRequest instantiates a new UpdateEmailAlertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEmailAlertRequestWithDefaults

`func NewUpdateEmailAlertRequestWithDefaults() *UpdateEmailAlertRequest`

NewUpdateEmailAlertRequestWithDefaults instantiates a new UpdateEmailAlertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateEmailAlertRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateEmailAlertRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateEmailAlertRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateEmailAlertRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateEmailAlertRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateEmailAlertRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateEmailAlertRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateEmailAlertRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEvents

`func (o *UpdateEmailAlertRequest) GetEvents() []EmailEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *UpdateEmailAlertRequest) GetEventsOk() (*[]EmailEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *UpdateEmailAlertRequest) SetEvents(v []EmailEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *UpdateEmailAlertRequest) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetActive

`func (o *UpdateEmailAlertRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateEmailAlertRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateEmailAlertRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateEmailAlertRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


