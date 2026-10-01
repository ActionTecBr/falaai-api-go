# \SpeechAPI

All URIs are relative to *https://api01-falaai.action.tec.br*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTranscriptionV1AudioTranscriptionsPost**](SpeechAPI.md#CreateTranscriptionV1AudioTranscriptionsPost) | **Post** /v1/audio/transcriptions | Transcribe audio to text



## CreateTranscriptionV1AudioTranscriptionsPost

> TranscriptionResponse CreateTranscriptionV1AudioTranscriptionsPost(ctx).File(file).Model(model).Language(language).ClientReferenceId(clientReferenceId).Execute()

Transcribe audio to text

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
	model := "model_example" // string |  (optional) (default to "falaai-transcribe-1")
	language := "language_example" // string |  (optional) (default to "pt")
	clientReferenceId := "clientReferenceId_example" // string | Optional client-supplied ID echoed verbatim in the response. Use to correlate/sync with your system. Accepted charset: [A-Za-z0-9._:-]. Not idempotency. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpeechAPI.CreateTranscriptionV1AudioTranscriptionsPost(context.Background()).File(file).Model(model).Language(language).ClientReferenceId(clientReferenceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpeechAPI.CreateTranscriptionV1AudioTranscriptionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTranscriptionV1AudioTranscriptionsPost`: TranscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SpeechAPI.CreateTranscriptionV1AudioTranscriptionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTranscriptionV1AudioTranscriptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 
 **model** | **string** |  | [default to &quot;falaai-transcribe-1&quot;]
 **language** | **string** |  | [default to &quot;pt&quot;]
 **clientReferenceId** | **string** | Optional client-supplied ID echoed verbatim in the response. Use to correlate/sync with your system. Accepted charset: [A-Za-z0-9._:-]. Not idempotency. | 

### Return type

[**TranscriptionResponse**](TranscriptionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

