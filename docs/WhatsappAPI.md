# \WhatsappAPI

All URIs are relative to *https://api01-falaai.action.tec.br*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExtractConversationsV1WhatsappExtractConversationsPost**](WhatsappAPI.md#ExtractConversationsV1WhatsappExtractConversationsPost) | **Post** /v1/whatsapp/extractConversations | Extract and segment WhatsApp conversations from an export



## ExtractConversationsV1WhatsappExtractConversationsPost

> WhatsappConversationsResponse ExtractConversationsV1WhatsappExtractConversationsPost(ctx).File(file).Start(start).End(end).Timezone(timezone).DateFormat(dateFormat).GapMinutes(gapMinutes).MinMessages(minMessages).CharsPerMinute(charsPerMinute).ClientReferenceId(clientReferenceId).Execute()

Extract and segment WhatsApp conversations from an export

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
	file := os.NewFile(1234, "some_file") // *os.File | 
	start := "start_example" // string | 
	end := "end_example" // string | 
	timezone := "timezone_example" // string | 
	dateFormat := "dateFormat_example" // string | 
	gapMinutes := float32(8.14) // float32 |  (optional) (default to 720)
	minMessages := int32(56) // int32 |  (optional) (default to 2)
	charsPerMinute := float32(8.14) // float32 |  (optional) (default to 800)
	clientReferenceId := "clientReferenceId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsappAPI.ExtractConversationsV1WhatsappExtractConversationsPost(context.Background()).File(file).Start(start).End(end).Timezone(timezone).DateFormat(dateFormat).GapMinutes(gapMinutes).MinMessages(minMessages).CharsPerMinute(charsPerMinute).ClientReferenceId(clientReferenceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsappAPI.ExtractConversationsV1WhatsappExtractConversationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExtractConversationsV1WhatsappExtractConversationsPost`: WhatsappConversationsResponse
	fmt.Fprintf(os.Stdout, "Response from `WhatsappAPI.ExtractConversationsV1WhatsappExtractConversationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExtractConversationsV1WhatsappExtractConversationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 
 **start** | **string** |  | 
 **end** | **string** |  | 
 **timezone** | **string** |  | 
 **dateFormat** | **string** |  | 
 **gapMinutes** | **float32** |  | [default to 720]
 **minMessages** | **int32** |  | [default to 2]
 **charsPerMinute** | **float32** |  | [default to 800]
 **clientReferenceId** | **string** |  | 

### Return type

[**WhatsappConversationsResponse**](WhatsappConversationsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

