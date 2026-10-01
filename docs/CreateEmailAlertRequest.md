# CreateEmailAlertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Nome identificador | 
**Email** | **string** | Email destino | 
**Events** | [**[]EmailEvent**](EmailEvent.md) | Eventos subscritos | 

## Methods

### NewCreateEmailAlertRequest

`func NewCreateEmailAlertRequest(name string, email string, events []EmailEvent, ) *CreateEmailAlertRequest`

NewCreateEmailAlertRequest instantiates a new CreateEmailAlertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEmailAlertRequestWithDefaults

`func NewCreateEmailAlertRequestWithDefaults() *CreateEmailAlertRequest`

NewCreateEmailAlertRequestWithDefaults instantiates a new CreateEmailAlertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateEmailAlertRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEmailAlertRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEmailAlertRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *CreateEmailAlertRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateEmailAlertRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateEmailAlertRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEvents

`func (o *CreateEmailAlertRequest) GetEvents() []EmailEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CreateEmailAlertRequest) GetEventsOk() (*[]EmailEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CreateEmailAlertRequest) SetEvents(v []EmailEvent)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


