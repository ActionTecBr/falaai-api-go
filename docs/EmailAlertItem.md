# EmailAlertItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Email alert id | 
**UserId** | **string** | Owner user id | 
**Name** | **string** | Email alert name | 
**Email** | **string** | Destination email | 
**Events** | **[]string** | Subscribed events | 
**Active** | **bool** | Is active | 
**CreatedAt** | **string** | ISO 8601 created | 
**UpdatedAt** | **string** | ISO 8601 updated | 

## Methods

### NewEmailAlertItem

`func NewEmailAlertItem(id string, userId string, name string, email string, events []string, active bool, createdAt string, updatedAt string, ) *EmailAlertItem`

NewEmailAlertItem instantiates a new EmailAlertItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailAlertItemWithDefaults

`func NewEmailAlertItemWithDefaults() *EmailAlertItem`

NewEmailAlertItemWithDefaults instantiates a new EmailAlertItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailAlertItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailAlertItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailAlertItem) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *EmailAlertItem) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EmailAlertItem) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EmailAlertItem) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetName

`func (o *EmailAlertItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailAlertItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailAlertItem) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *EmailAlertItem) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailAlertItem) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailAlertItem) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEvents

`func (o *EmailAlertItem) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EmailAlertItem) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EmailAlertItem) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetActive

`func (o *EmailAlertItem) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EmailAlertItem) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EmailAlertItem) SetActive(v bool)`

SetActive sets Active field to given value.


### GetCreatedAt

`func (o *EmailAlertItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EmailAlertItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EmailAlertItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EmailAlertItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EmailAlertItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EmailAlertItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


