# \EmailAlertsAPI

All URIs are relative to *https://api01-falaai.action.tec.br*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmailAlertV1EmailAlertsPost**](EmailAlertsAPI.md#CreateEmailAlertV1EmailAlertsPost) | **Post** /v1/email-alerts | Create email alert
[**DeleteEmailAlertV1EmailAlertsAlertIdDelete**](EmailAlertsAPI.md#DeleteEmailAlertV1EmailAlertsAlertIdDelete) | **Delete** /v1/email-alerts/{alert_id} | Delete email alert
[**ListEmailAlertsV1EmailAlertsGet**](EmailAlertsAPI.md#ListEmailAlertsV1EmailAlertsGet) | **Get** /v1/email-alerts | List email alerts
[**UpdateEmailAlertV1EmailAlertsAlertIdPut**](EmailAlertsAPI.md#UpdateEmailAlertV1EmailAlertsAlertIdPut) | **Put** /v1/email-alerts/{alert_id} | Update email alert



## CreateEmailAlertV1EmailAlertsPost

> EmailAlertItem CreateEmailAlertV1EmailAlertsPost(ctx).CreateEmailAlertRequest(createEmailAlertRequest).Execute()

Create email alert

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createEmailAlertRequest := *openapiclient.NewCreateEmailAlertRequest("Name_example", "Email_example", []openapiclient.EmailEvent{openapiclient.EmailEvent("subscription.created")}) // CreateEmailAlertRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAlertsAPI.CreateEmailAlertV1EmailAlertsPost(context.Background()).CreateEmailAlertRequest(createEmailAlertRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAlertsAPI.CreateEmailAlertV1EmailAlertsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEmailAlertV1EmailAlertsPost`: EmailAlertItem
	fmt.Fprintf(os.Stdout, "Response from `EmailAlertsAPI.CreateEmailAlertV1EmailAlertsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmailAlertV1EmailAlertsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEmailAlertRequest** | [**CreateEmailAlertRequest**](CreateEmailAlertRequest.md) |  | 

### Return type

[**EmailAlertItem**](EmailAlertItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEmailAlertV1EmailAlertsAlertIdDelete

> EmailAlertMessageResponse DeleteEmailAlertV1EmailAlertsAlertIdDelete(ctx, alertId).Execute()

Delete email alert

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	alertId := "alertId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAlertsAPI.DeleteEmailAlertV1EmailAlertsAlertIdDelete(context.Background(), alertId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAlertsAPI.DeleteEmailAlertV1EmailAlertsAlertIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEmailAlertV1EmailAlertsAlertIdDelete`: EmailAlertMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `EmailAlertsAPI.DeleteEmailAlertV1EmailAlertsAlertIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmailAlertV1EmailAlertsAlertIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmailAlertMessageResponse**](EmailAlertMessageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEmailAlertsV1EmailAlertsGet

> EmailAlertListResponse ListEmailAlertsV1EmailAlertsGet(ctx).Page(page).Limit(limit).Execute()

List email alerts

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAlertsAPI.ListEmailAlertsV1EmailAlertsGet(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAlertsAPI.ListEmailAlertsV1EmailAlertsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEmailAlertsV1EmailAlertsGet`: EmailAlertListResponse
	fmt.Fprintf(os.Stdout, "Response from `EmailAlertsAPI.ListEmailAlertsV1EmailAlertsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEmailAlertsV1EmailAlertsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 20]

### Return type

[**EmailAlertListResponse**](EmailAlertListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEmailAlertV1EmailAlertsAlertIdPut

> EmailAlertMessageResponse UpdateEmailAlertV1EmailAlertsAlertIdPut(ctx, alertId).UpdateEmailAlertRequest(updateEmailAlertRequest).Execute()

Update email alert

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	alertId := "alertId_example" // string | 
	updateEmailAlertRequest := *openapiclient.NewUpdateEmailAlertRequest() // UpdateEmailAlertRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAlertsAPI.UpdateEmailAlertV1EmailAlertsAlertIdPut(context.Background(), alertId).UpdateEmailAlertRequest(updateEmailAlertRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAlertsAPI.UpdateEmailAlertV1EmailAlertsAlertIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEmailAlertV1EmailAlertsAlertIdPut`: EmailAlertMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `EmailAlertsAPI.UpdateEmailAlertV1EmailAlertsAlertIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmailAlertV1EmailAlertsAlertIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEmailAlertRequest** | [**UpdateEmailAlertRequest**](UpdateEmailAlertRequest.md) |  | 

### Return type

[**EmailAlertMessageResponse**](EmailAlertMessageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

