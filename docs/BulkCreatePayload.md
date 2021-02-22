# BulkCreatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sources** | Pointer to [**[]BulkCreatePayloadSources**](BulkCreatePayloadSources.md) | Array of Source objects to create. Only supported fields are name + type, source_type_id will automatically be set based on the &#x60;source_type_name&#x60;.  | [optional] 
**Endpoints** | Pointer to [**[]BulkCreatePayloadEndpoints**](BulkCreatePayloadEndpoints.md) | Array of Endpoint objects to create. The operation looks up the parent source by the &#x60;source_name&#x60; attribute so the &#x60;source_name&#x60; must match one of the &#x60;source&#x60;&#39;s names in the payload.  | [optional] 
**Applications** | Pointer to [**[]BulkCreatePayloadApplications**](BulkCreatePayloadApplications.md) | Array of Application objects to create. The operation looks up the parent Source by the &#x60;source_name&#x60; attribute so the &#x60;source_name&#x60; must match one of the &#x60;source&#x60;&#39;s names in the payload.  application_type_id will be automatically looked up and set by the &#x60;application_type_name&#x60; attribute via regex.  | [optional] 
**Authentications** | Pointer to [**[]BulkCreatePayloadAuthentications**](BulkCreatePayloadAuthentications.md) | Array of Authentications to create. This one is a bit more tricky. &#x60;resource_type&#x60; tells the action where to look for the parent, must be either application/endpoint/source  if the parent is a source, it looks up by name. if the parent is an endpoint, it looks up via host so the hostname must match. if the parent is an application, it looks up via application type so the value must match the application type which matches  | [optional] 

## Methods

### NewBulkCreatePayload

`func NewBulkCreatePayload() *BulkCreatePayload`

NewBulkCreatePayload instantiates a new BulkCreatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCreatePayloadWithDefaults

`func NewBulkCreatePayloadWithDefaults() *BulkCreatePayload`

NewBulkCreatePayloadWithDefaults instantiates a new BulkCreatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSources

`func (o *BulkCreatePayload) GetSources() []BulkCreatePayloadSources`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *BulkCreatePayload) GetSourcesOk() (*[]BulkCreatePayloadSources, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *BulkCreatePayload) SetSources(v []BulkCreatePayloadSources)`

SetSources sets Sources field to given value.

### HasSources

`func (o *BulkCreatePayload) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetEndpoints

`func (o *BulkCreatePayload) GetEndpoints() []BulkCreatePayloadEndpoints`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *BulkCreatePayload) GetEndpointsOk() (*[]BulkCreatePayloadEndpoints, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *BulkCreatePayload) SetEndpoints(v []BulkCreatePayloadEndpoints)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *BulkCreatePayload) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetApplications

`func (o *BulkCreatePayload) GetApplications() []BulkCreatePayloadApplications`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *BulkCreatePayload) GetApplicationsOk() (*[]BulkCreatePayloadApplications, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *BulkCreatePayload) SetApplications(v []BulkCreatePayloadApplications)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *BulkCreatePayload) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetAuthentications

`func (o *BulkCreatePayload) GetAuthentications() []BulkCreatePayloadAuthentications`

GetAuthentications returns the Authentications field if non-nil, zero value otherwise.

### GetAuthenticationsOk

`func (o *BulkCreatePayload) GetAuthenticationsOk() (*[]BulkCreatePayloadAuthentications, bool)`

GetAuthenticationsOk returns a tuple with the Authentications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentications

`func (o *BulkCreatePayload) SetAuthentications(v []BulkCreatePayloadAuthentications)`

SetAuthentications sets Authentications field to given value.

### HasAuthentications

`func (o *BulkCreatePayload) HasAuthentications() bool`

HasAuthentications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


