# Endpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityStatus** | Pointer to **string** |  | [optional] 
**AvailabilityStatusError** | Pointer to **string** |  | [optional] 
**CertificateAuthority** | Pointer to **string** | Optional X.509 Certificate Authority | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Default** | Pointer to **bool** |  | [optional] 
**Host** | Pointer to **string** | URI host component | [optional] 
**Id** | Pointer to **string** | ID of the resource | [optional] [readonly] 
**LastAvailableAt** | Pointer to **time.Time** |  | [optional] 
**LastCheckedAt** | Pointer to **time.Time** |  | [optional] 
**Path** | Pointer to **string** | URI path component | [optional] 
**Port** | Pointer to **int32** | URI port component | [optional] 
**ReceptorNode** | Pointer to **string** | Identifier of a receptor node | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Scheme** | Pointer to **string** | URI scheme component | [optional] 
**SourceId** | Pointer to **string** | ID of the resource | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**VerifySsl** | Pointer to **bool** | Should SSL be verified | [optional] 

## Methods

### NewEndpoint

`func NewEndpoint() *Endpoint`

NewEndpoint instantiates a new Endpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointWithDefaults

`func NewEndpointWithDefaults() *Endpoint`

NewEndpointWithDefaults instantiates a new Endpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityStatus

`func (o *Endpoint) GetAvailabilityStatus() string`

GetAvailabilityStatus returns the AvailabilityStatus field if non-nil, zero value otherwise.

### GetAvailabilityStatusOk

`func (o *Endpoint) GetAvailabilityStatusOk() (*string, bool)`

GetAvailabilityStatusOk returns a tuple with the AvailabilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityStatus

`func (o *Endpoint) SetAvailabilityStatus(v string)`

SetAvailabilityStatus sets AvailabilityStatus field to given value.

### HasAvailabilityStatus

`func (o *Endpoint) HasAvailabilityStatus() bool`

HasAvailabilityStatus returns a boolean if a field has been set.

### GetAvailabilityStatusError

`func (o *Endpoint) GetAvailabilityStatusError() string`

GetAvailabilityStatusError returns the AvailabilityStatusError field if non-nil, zero value otherwise.

### GetAvailabilityStatusErrorOk

`func (o *Endpoint) GetAvailabilityStatusErrorOk() (*string, bool)`

GetAvailabilityStatusErrorOk returns a tuple with the AvailabilityStatusError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityStatusError

`func (o *Endpoint) SetAvailabilityStatusError(v string)`

SetAvailabilityStatusError sets AvailabilityStatusError field to given value.

### HasAvailabilityStatusError

`func (o *Endpoint) HasAvailabilityStatusError() bool`

HasAvailabilityStatusError returns a boolean if a field has been set.

### GetCertificateAuthority

`func (o *Endpoint) GetCertificateAuthority() string`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *Endpoint) GetCertificateAuthorityOk() (*string, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *Endpoint) SetCertificateAuthority(v string)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *Endpoint) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Endpoint) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Endpoint) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Endpoint) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Endpoint) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefault

`func (o *Endpoint) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Endpoint) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Endpoint) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *Endpoint) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetHost

`func (o *Endpoint) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Endpoint) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Endpoint) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *Endpoint) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *Endpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Endpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Endpoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Endpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastAvailableAt

`func (o *Endpoint) GetLastAvailableAt() time.Time`

GetLastAvailableAt returns the LastAvailableAt field if non-nil, zero value otherwise.

### GetLastAvailableAtOk

`func (o *Endpoint) GetLastAvailableAtOk() (*time.Time, bool)`

GetLastAvailableAtOk returns a tuple with the LastAvailableAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAvailableAt

`func (o *Endpoint) SetLastAvailableAt(v time.Time)`

SetLastAvailableAt sets LastAvailableAt field to given value.

### HasLastAvailableAt

`func (o *Endpoint) HasLastAvailableAt() bool`

HasLastAvailableAt returns a boolean if a field has been set.

### GetLastCheckedAt

`func (o *Endpoint) GetLastCheckedAt() time.Time`

GetLastCheckedAt returns the LastCheckedAt field if non-nil, zero value otherwise.

### GetLastCheckedAtOk

`func (o *Endpoint) GetLastCheckedAtOk() (*time.Time, bool)`

GetLastCheckedAtOk returns a tuple with the LastCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckedAt

`func (o *Endpoint) SetLastCheckedAt(v time.Time)`

SetLastCheckedAt sets LastCheckedAt field to given value.

### HasLastCheckedAt

`func (o *Endpoint) HasLastCheckedAt() bool`

HasLastCheckedAt returns a boolean if a field has been set.

### GetPath

`func (o *Endpoint) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Endpoint) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Endpoint) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Endpoint) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPort

`func (o *Endpoint) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Endpoint) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Endpoint) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Endpoint) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetReceptorNode

`func (o *Endpoint) GetReceptorNode() string`

GetReceptorNode returns the ReceptorNode field if non-nil, zero value otherwise.

### GetReceptorNodeOk

`func (o *Endpoint) GetReceptorNodeOk() (*string, bool)`

GetReceptorNodeOk returns a tuple with the ReceptorNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceptorNode

`func (o *Endpoint) SetReceptorNode(v string)`

SetReceptorNode sets ReceptorNode field to given value.

### HasReceptorNode

`func (o *Endpoint) HasReceptorNode() bool`

HasReceptorNode returns a boolean if a field has been set.

### GetRole

`func (o *Endpoint) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Endpoint) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Endpoint) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *Endpoint) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetScheme

`func (o *Endpoint) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *Endpoint) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *Endpoint) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *Endpoint) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetSourceId

`func (o *Endpoint) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Endpoint) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Endpoint) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *Endpoint) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Endpoint) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Endpoint) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Endpoint) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Endpoint) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVerifySsl

`func (o *Endpoint) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *Endpoint) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *Endpoint) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *Endpoint) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


