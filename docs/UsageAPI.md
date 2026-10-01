# \UsageAPI

All URIs are relative to *https://api01-falaai.action.tec.br*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsageByKeyV1UsageByKeyGet**](UsageAPI.md#GetUsageByKeyV1UsageByKeyGet) | **Get** /v1/usage/by-key | Get Usage By Key
[**GetUsageLogV1UsageLogGet**](UsageAPI.md#GetUsageLogV1UsageLogGet) | **Get** /v1/usage/log | Get Usage Log



## GetUsageByKeyV1UsageByKeyGet

> []UsageByKeyItem GetUsageByKeyV1UsageByKeyGet(ctx).KeyId(keyId).Execute()

Get Usage By Key

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
	keyId := "keyId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.GetUsageByKeyV1UsageByKeyGet(context.Background()).KeyId(keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.GetUsageByKeyV1UsageByKeyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsageByKeyV1UsageByKeyGet`: []UsageByKeyItem
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.GetUsageByKeyV1UsageByKeyGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageByKeyV1UsageByKeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyId** | **string** |  | 

### Return type

[**[]UsageByKeyItem**](UsageByKeyItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageLogV1UsageLogGet

> UsageLogResponse GetUsageLogV1UsageLogGet(ctx).Page(page).Limit(limit).ApiKeyId(apiKeyId).Execute()

Get Usage Log

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
	apiKeyId := "apiKeyId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.GetUsageLogV1UsageLogGet(context.Background()).Page(page).Limit(limit).ApiKeyId(apiKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.GetUsageLogV1UsageLogGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsageLogV1UsageLogGet`: UsageLogResponse
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.GetUsageLogV1UsageLogGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageLogV1UsageLogGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 20]
 **apiKeyId** | **string** |  | 

### Return type

[**UsageLogResponse**](UsageLogResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

