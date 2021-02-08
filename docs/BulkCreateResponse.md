# BulkCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Superkey** | Pointer to **bool** |  | [optional] [default to false]
**Sources** | Pointer to [**[]Source**](Source.md) |  | [optional] 
**Endpoints** | Pointer to [**[]Endpoint**](Endpoint.md) |  | [optional] 
**Applications** | Pointer to [**[]Application**](Application.md) |  | [optional] 
**Authentications** | Pointer to [**[]Authentication**](Authentication.md) |  | [optional] 

## Methods

### NewBulkCreateResponse

`func NewBulkCreateResponse() *BulkCreateResponse`

NewBulkCreateResponse instantiates a new BulkCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCreateResponseWithDefaults

`func NewBulkCreateResponseWithDefaults() *BulkCreateResponse`

NewBulkCreateResponseWithDefaults instantiates a new BulkCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuperkey

`func (o *BulkCreateResponse) GetSuperkey() bool`

GetSuperkey returns the Superkey field if non-nil, zero value otherwise.

### GetSuperkeyOk

`func (o *BulkCreateResponse) GetSuperkeyOk() (*bool, bool)`

GetSuperkeyOk returns a tuple with the Superkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperkey

`func (o *BulkCreateResponse) SetSuperkey(v bool)`

SetSuperkey sets Superkey field to given value.

### HasSuperkey

`func (o *BulkCreateResponse) HasSuperkey() bool`

HasSuperkey returns a boolean if a field has been set.

### GetSources

`func (o *BulkCreateResponse) GetSources() []Source`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *BulkCreateResponse) GetSourcesOk() (*[]Source, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *BulkCreateResponse) SetSources(v []Source)`

SetSources sets Sources field to given value.

### HasSources

`func (o *BulkCreateResponse) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetEndpoints

`func (o *BulkCreateResponse) GetEndpoints() []Endpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *BulkCreateResponse) GetEndpointsOk() (*[]Endpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *BulkCreateResponse) SetEndpoints(v []Endpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *BulkCreateResponse) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetApplications

`func (o *BulkCreateResponse) GetApplications() []Application`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *BulkCreateResponse) GetApplicationsOk() (*[]Application, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *BulkCreateResponse) SetApplications(v []Application)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *BulkCreateResponse) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetAuthentications

`func (o *BulkCreateResponse) GetAuthentications() []Authentication`

GetAuthentications returns the Authentications field if non-nil, zero value otherwise.

### GetAuthenticationsOk

`func (o *BulkCreateResponse) GetAuthenticationsOk() (*[]Authentication, bool)`

GetAuthenticationsOk returns a tuple with the Authentications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentications

`func (o *BulkCreateResponse) SetAuthentications(v []Authentication)`

SetAuthentications sets Authentications field to given value.

### HasAuthentications

`func (o *BulkCreateResponse) HasAuthentications() bool`

HasAuthentications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


