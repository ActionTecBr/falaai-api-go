# DiagnosticResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique analysis identifier. Prefix &#39;di-&#39; + UUID | 
**ResponseLanguage** | **string** | Language used in the response. E.g.: &#39;pt-BR&#39;, &#39;en-US&#39;, &#39;es-ES&#39; | 
**Object** | **string** | Object type. Always &#39;analysis&#39; | 
**Analysis** | [**DiagnosticAnalysisMap**](DiagnosticAnalysisMap.md) | The 6 conversation analyses (5 + participants) | 
**Usage** | [**DiagnosticUsage**](DiagnosticUsage.md) | Usage and processing information | 
**ClientReferenceId** | Pointer to **string** | Client-supplied ID echoed verbatim (if provided in request) | [optional] 

## Methods

### NewDiagnosticResponse

`func NewDiagnosticResponse(id string, responseLanguage string, object string, analysis DiagnosticAnalysisMap, usage DiagnosticUsage, ) *DiagnosticResponse`

NewDiagnosticResponse instantiates a new DiagnosticResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosticResponseWithDefaults

`func NewDiagnosticResponseWithDefaults() *DiagnosticResponse`

NewDiagnosticResponseWithDefaults instantiates a new DiagnosticResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DiagnosticResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiagnosticResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiagnosticResponse) SetId(v string)`

SetId sets Id field to given value.


### GetResponseLanguage

`func (o *DiagnosticResponse) GetResponseLanguage() string`

GetResponseLanguage returns the ResponseLanguage field if non-nil, zero value otherwise.

### GetResponseLanguageOk

`func (o *DiagnosticResponse) GetResponseLanguageOk() (*string, bool)`

GetResponseLanguageOk returns a tuple with the ResponseLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLanguage

`func (o *DiagnosticResponse) SetResponseLanguage(v string)`

SetResponseLanguage sets ResponseLanguage field to given value.


### GetObject

`func (o *DiagnosticResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *DiagnosticResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *DiagnosticResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetAnalysis

`func (o *DiagnosticResponse) GetAnalysis() DiagnosticAnalysisMap`

GetAnalysis returns the Analysis field if non-nil, zero value otherwise.

### GetAnalysisOk

`func (o *DiagnosticResponse) GetAnalysisOk() (*DiagnosticAnalysisMap, bool)`

GetAnalysisOk returns a tuple with the Analysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysis

`func (o *DiagnosticResponse) SetAnalysis(v DiagnosticAnalysisMap)`

SetAnalysis sets Analysis field to given value.


### GetUsage

`func (o *DiagnosticResponse) GetUsage() DiagnosticUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *DiagnosticResponse) GetUsageOk() (*DiagnosticUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *DiagnosticResponse) SetUsage(v DiagnosticUsage)`

SetUsage sets Usage field to given value.


### GetClientReferenceId

`func (o *DiagnosticResponse) GetClientReferenceId() string`

GetClientReferenceId returns the ClientReferenceId field if non-nil, zero value otherwise.

### GetClientReferenceIdOk

`func (o *DiagnosticResponse) GetClientReferenceIdOk() (*string, bool)`

GetClientReferenceIdOk returns a tuple with the ClientReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientReferenceId

`func (o *DiagnosticResponse) SetClientReferenceId(v string)`

SetClientReferenceId sets ClientReferenceId field to given value.

### HasClientReferenceId

`func (o *DiagnosticResponse) HasClientReferenceId() bool`

HasClientReferenceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


