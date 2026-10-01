# DiagnosticUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Characters** | **int32** | Total characters analyzed | 
**CreditsConsumed** | **int32** | Credits consumed: max(ceil(chars/500)*3, 3) * 5 | 
**ProcessingMs** | **int32** | Total processing time in milliseconds | 

## Methods

### NewDiagnosticUsage

`func NewDiagnosticUsage(characters int32, creditsConsumed int32, processingMs int32, ) *DiagnosticUsage`

NewDiagnosticUsage instantiates a new DiagnosticUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosticUsageWithDefaults

`func NewDiagnosticUsageWithDefaults() *DiagnosticUsage`

NewDiagnosticUsageWithDefaults instantiates a new DiagnosticUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharacters

`func (o *DiagnosticUsage) GetCharacters() int32`

GetCharacters returns the Characters field if non-nil, zero value otherwise.

### GetCharactersOk

`func (o *DiagnosticUsage) GetCharactersOk() (*int32, bool)`

GetCharactersOk returns a tuple with the Characters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacters

`func (o *DiagnosticUsage) SetCharacters(v int32)`

SetCharacters sets Characters field to given value.


### GetCreditsConsumed

`func (o *DiagnosticUsage) GetCreditsConsumed() int32`

GetCreditsConsumed returns the CreditsConsumed field if non-nil, zero value otherwise.

### GetCreditsConsumedOk

`func (o *DiagnosticUsage) GetCreditsConsumedOk() (*int32, bool)`

GetCreditsConsumedOk returns a tuple with the CreditsConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsConsumed

`func (o *DiagnosticUsage) SetCreditsConsumed(v int32)`

SetCreditsConsumed sets CreditsConsumed field to given value.


### GetProcessingMs

`func (o *DiagnosticUsage) GetProcessingMs() int32`

GetProcessingMs returns the ProcessingMs field if non-nil, zero value otherwise.

### GetProcessingMsOk

`func (o *DiagnosticUsage) GetProcessingMsOk() (*int32, bool)`

GetProcessingMsOk returns a tuple with the ProcessingMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingMs

`func (o *DiagnosticUsage) SetProcessingMs(v int32)`

SetProcessingMs sets ProcessingMs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


