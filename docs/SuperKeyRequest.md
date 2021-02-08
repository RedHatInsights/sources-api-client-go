# SuperKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to **string** |  | [optional] 
**Applications** | Pointer to **[]string** |  | [optional] 
**SourceId** | Pointer to **string** | the superkey source to create applications under. only required if not included in a &#x60;bulk_create&#x60; request  | [optional] 

## Methods

### NewSuperKeyRequest

`func NewSuperKeyRequest() *SuperKeyRequest`

NewSuperKeyRequest instantiates a new SuperKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuperKeyRequestWithDefaults

`func NewSuperKeyRequestWithDefaults() *SuperKeyRequest`

NewSuperKeyRequestWithDefaults instantiates a new SuperKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *SuperKeyRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SuperKeyRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SuperKeyRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *SuperKeyRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetApplications

`func (o *SuperKeyRequest) GetApplications() []string`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *SuperKeyRequest) GetApplicationsOk() (*[]string, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *SuperKeyRequest) SetApplications(v []string)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *SuperKeyRequest) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetSourceId

`func (o *SuperKeyRequest) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *SuperKeyRequest) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *SuperKeyRequest) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *SuperKeyRequest) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


