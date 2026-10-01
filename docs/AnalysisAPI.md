# \AnalysisAPI

All URIs are relative to *https://api01-falaai.action.tec.br*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDiagnosticV1AnalyzeDiagnosticPost**](AnalysisAPI.md#CreateDiagnosticV1AnalyzeDiagnosticPost) | **Post** /v1/analyze/diagnostic | Analyze a call transcript — 5 parallel analyses
[**CreateRiskAuditV1AnalyzeRiskAuditPost**](AnalysisAPI.md#CreateRiskAuditV1AnalyzeRiskAuditPost) | **Post** /v1/analyze/riskAudit | Compliance Risk Audit — conversation compliance analysis



## CreateDiagnosticV1AnalyzeDiagnosticPost

> DiagnosticResponse CreateDiagnosticV1AnalyzeDiagnosticPost(ctx).DiagnosticRequest(diagnosticRequest).Execute()

Analyze a call transcript — 5 parallel analyses

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
	diagnosticRequest := *openapiclient.NewDiagnosticRequest("Language_example", float32(123)) // DiagnosticRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysisAPI.CreateDiagnosticV1AnalyzeDiagnosticPost(context.Background()).DiagnosticRequest(diagnosticRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysisAPI.CreateDiagnosticV1AnalyzeDiagnosticPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDiagnosticV1AnalyzeDiagnosticPost`: DiagnosticResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalysisAPI.CreateDiagnosticV1AnalyzeDiagnosticPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDiagnosticV1AnalyzeDiagnosticPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **diagnosticRequest** | [**DiagnosticRequest**](DiagnosticRequest.md) |  | 

### Return type

[**DiagnosticResponse**](DiagnosticResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRiskAuditV1AnalyzeRiskAuditPost

> RiskAuditV2Response CreateRiskAuditV1AnalyzeRiskAuditPost(ctx).RiskAuditRequest(riskAuditRequest).Execute()

Compliance Risk Audit — conversation compliance analysis

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
	riskAuditRequest := *openapiclient.NewRiskAuditRequest(float32(123), "Language_example", "ResponseLanguage_example") // RiskAuditRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysisAPI.CreateRiskAuditV1AnalyzeRiskAuditPost(context.Background()).RiskAuditRequest(riskAuditRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysisAPI.CreateRiskAuditV1AnalyzeRiskAuditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRiskAuditV1AnalyzeRiskAuditPost`: RiskAuditV2Response
	fmt.Fprintf(os.Stdout, "Response from `AnalysisAPI.CreateRiskAuditV1AnalyzeRiskAuditPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRiskAuditV1AnalyzeRiskAuditPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **riskAuditRequest** | [**RiskAuditRequest**](RiskAuditRequest.md) |  | 

### Return type

[**RiskAuditV2Response**](RiskAuditV2Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

