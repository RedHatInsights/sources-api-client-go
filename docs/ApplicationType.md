# ApplicationType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] [readonly] 
**DependentApplications** | Pointer to **map[string]interface{}** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | ID of the resource | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**SupportedAuthenticationTypes** | Pointer to **map[string]interface{}** |  | [optional] 
**SupportedSourceTypes** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] [readonly] 

## Methods

### NewApplicationType

`func NewApplicationType() *ApplicationType`

NewApplicationType instantiates a new ApplicationType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationTypeWithDefaults

`func NewApplicationTypeWithDefaults() *ApplicationType`

NewApplicationTypeWithDefaults instantiates a new ApplicationType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ApplicationType) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicationType) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicationType) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApplicationType) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDependentApplications

`func (o *ApplicationType) GetDependentApplications() map[string]interface{}`

GetDependentApplications returns the DependentApplications field if non-nil, zero value otherwise.

### GetDependentApplicationsOk

`func (o *ApplicationType) GetDependentApplicationsOk() (*map[string]interface{}, bool)`

GetDependentApplicationsOk returns a tuple with the DependentApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentApplications

`func (o *ApplicationType) SetDependentApplications(v map[string]interface{})`

SetDependentApplications sets DependentApplications field to given value.

### HasDependentApplications

`func (o *ApplicationType) HasDependentApplications() bool`

HasDependentApplications returns a boolean if a field has been set.

### GetDisplayName

`func (o *ApplicationType) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApplicationType) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApplicationType) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApplicationType) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *ApplicationType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApplicationType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSupportedAuthenticationTypes

`func (o *ApplicationType) GetSupportedAuthenticationTypes() map[string]interface{}`

GetSupportedAuthenticationTypes returns the SupportedAuthenticationTypes field if non-nil, zero value otherwise.

### GetSupportedAuthenticationTypesOk

`func (o *ApplicationType) GetSupportedAuthenticationTypesOk() (*map[string]interface{}, bool)`

GetSupportedAuthenticationTypesOk returns a tuple with the SupportedAuthenticationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedAuthenticationTypes

`func (o *ApplicationType) SetSupportedAuthenticationTypes(v map[string]interface{})`

SetSupportedAuthenticationTypes sets SupportedAuthenticationTypes field to given value.

### HasSupportedAuthenticationTypes

`func (o *ApplicationType) HasSupportedAuthenticationTypes() bool`

HasSupportedAuthenticationTypes returns a boolean if a field has been set.

### GetSupportedSourceTypes

`func (o *ApplicationType) GetSupportedSourceTypes() map[string]interface{}`

GetSupportedSourceTypes returns the SupportedSourceTypes field if non-nil, zero value otherwise.

### GetSupportedSourceTypesOk

`func (o *ApplicationType) GetSupportedSourceTypesOk() (*map[string]interface{}, bool)`

GetSupportedSourceTypesOk returns a tuple with the SupportedSourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedSourceTypes

`func (o *ApplicationType) SetSupportedSourceTypes(v map[string]interface{})`

SetSupportedSourceTypes sets SupportedSourceTypes field to given value.

### HasSupportedSourceTypes

`func (o *ApplicationType) HasSupportedSourceTypes() bool`

HasSupportedSourceTypes returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ApplicationType) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ApplicationType) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ApplicationType) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ApplicationType) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


