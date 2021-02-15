# Source

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppCreationWorkflow** | Pointer to **string** |  | [optional] 
**AvailabilityStatus** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Id** | Pointer to **string** | ID of the resource | [optional] [readonly] 
**Imported** | Pointer to **string** |  | [optional] 
**LastAvailableAt** | Pointer to **time.Time** |  | [optional] 
**LastCheckedAt** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SourceRef** | Pointer to **string** |  | [optional] 
**SourceTypeId** | Pointer to **string** | ID of the resource | [optional] [readonly] 
**Uid** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewSource

`func NewSource() *Source`

NewSource instantiates a new Source object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceWithDefaults

`func NewSourceWithDefaults() *Source`

NewSourceWithDefaults instantiates a new Source object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppCreationWorkflow

`func (o *Source) GetAppCreationWorkflow() string`

GetAppCreationWorkflow returns the AppCreationWorkflow field if non-nil, zero value otherwise.

### GetAppCreationWorkflowOk

`func (o *Source) GetAppCreationWorkflowOk() (*string, bool)`

GetAppCreationWorkflowOk returns a tuple with the AppCreationWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCreationWorkflow

`func (o *Source) SetAppCreationWorkflow(v string)`

SetAppCreationWorkflow sets AppCreationWorkflow field to given value.

### HasAppCreationWorkflow

`func (o *Source) HasAppCreationWorkflow() bool`

HasAppCreationWorkflow returns a boolean if a field has been set.

### GetAvailabilityStatus

`func (o *Source) GetAvailabilityStatus() string`

GetAvailabilityStatus returns the AvailabilityStatus field if non-nil, zero value otherwise.

### GetAvailabilityStatusOk

`func (o *Source) GetAvailabilityStatusOk() (*string, bool)`

GetAvailabilityStatusOk returns a tuple with the AvailabilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityStatus

`func (o *Source) SetAvailabilityStatus(v string)`

SetAvailabilityStatus sets AvailabilityStatus field to given value.

### HasAvailabilityStatus

`func (o *Source) HasAvailabilityStatus() bool`

HasAvailabilityStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Source) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Source) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Source) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Source) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *Source) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Source) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Source) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Source) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImported

`func (o *Source) GetImported() string`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *Source) GetImportedOk() (*string, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *Source) SetImported(v string)`

SetImported sets Imported field to given value.

### HasImported

`func (o *Source) HasImported() bool`

HasImported returns a boolean if a field has been set.

### GetLastAvailableAt

`func (o *Source) GetLastAvailableAt() time.Time`

GetLastAvailableAt returns the LastAvailableAt field if non-nil, zero value otherwise.

### GetLastAvailableAtOk

`func (o *Source) GetLastAvailableAtOk() (*time.Time, bool)`

GetLastAvailableAtOk returns a tuple with the LastAvailableAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAvailableAt

`func (o *Source) SetLastAvailableAt(v time.Time)`

SetLastAvailableAt sets LastAvailableAt field to given value.

### HasLastAvailableAt

`func (o *Source) HasLastAvailableAt() bool`

HasLastAvailableAt returns a boolean if a field has been set.

### GetLastCheckedAt

`func (o *Source) GetLastCheckedAt() time.Time`

GetLastCheckedAt returns the LastCheckedAt field if non-nil, zero value otherwise.

### GetLastCheckedAtOk

`func (o *Source) GetLastCheckedAtOk() (*time.Time, bool)`

GetLastCheckedAtOk returns a tuple with the LastCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckedAt

`func (o *Source) SetLastCheckedAt(v time.Time)`

SetLastCheckedAt sets LastCheckedAt field to given value.

### HasLastCheckedAt

`func (o *Source) HasLastCheckedAt() bool`

HasLastCheckedAt returns a boolean if a field has been set.

### GetName

`func (o *Source) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Source) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Source) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Source) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceRef

`func (o *Source) GetSourceRef() string`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *Source) GetSourceRefOk() (*string, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *Source) SetSourceRef(v string)`

SetSourceRef sets SourceRef field to given value.

### HasSourceRef

`func (o *Source) HasSourceRef() bool`

HasSourceRef returns a boolean if a field has been set.

### GetSourceTypeId

`func (o *Source) GetSourceTypeId() string`

GetSourceTypeId returns the SourceTypeId field if non-nil, zero value otherwise.

### GetSourceTypeIdOk

`func (o *Source) GetSourceTypeIdOk() (*string, bool)`

GetSourceTypeIdOk returns a tuple with the SourceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTypeId

`func (o *Source) SetSourceTypeId(v string)`

SetSourceTypeId sets SourceTypeId field to given value.

### HasSourceTypeId

`func (o *Source) HasSourceTypeId() bool`

HasSourceTypeId returns a boolean if a field has been set.

### GetUid

`func (o *Source) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Source) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Source) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *Source) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Source) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Source) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Source) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Source) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *Source) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Source) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Source) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Source) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


