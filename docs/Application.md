# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationTypeId** | Pointer to **string** | ID of the resource | [optional] [readonly] 
**AvailabilityStatus** | Pointer to **string** |  | [optional] 
**AvailabilityStatusError** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Extra** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | Pointer to **string** | ID of the resource | [optional] [readonly] 
**LastAvailableAt** | Pointer to **time.Time** |  | [optional] 
**LastCheckedAt** | Pointer to **time.Time** |  | [optional] 
**SourceId** | Pointer to **string** | ID of the resource | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewApplication

`func NewApplication() *Application`

NewApplication instantiates a new Application object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWithDefaults

`func NewApplicationWithDefaults() *Application`

NewApplicationWithDefaults instantiates a new Application object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationTypeId

`func (o *Application) GetApplicationTypeId() string`

GetApplicationTypeId returns the ApplicationTypeId field if non-nil, zero value otherwise.

### GetApplicationTypeIdOk

`func (o *Application) GetApplicationTypeIdOk() (*string, bool)`

GetApplicationTypeIdOk returns a tuple with the ApplicationTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationTypeId

`func (o *Application) SetApplicationTypeId(v string)`

SetApplicationTypeId sets ApplicationTypeId field to given value.

### HasApplicationTypeId

`func (o *Application) HasApplicationTypeId() bool`

HasApplicationTypeId returns a boolean if a field has been set.

### GetAvailabilityStatus

`func (o *Application) GetAvailabilityStatus() string`

GetAvailabilityStatus returns the AvailabilityStatus field if non-nil, zero value otherwise.

### GetAvailabilityStatusOk

`func (o *Application) GetAvailabilityStatusOk() (*string, bool)`

GetAvailabilityStatusOk returns a tuple with the AvailabilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityStatus

`func (o *Application) SetAvailabilityStatus(v string)`

SetAvailabilityStatus sets AvailabilityStatus field to given value.

### HasAvailabilityStatus

`func (o *Application) HasAvailabilityStatus() bool`

HasAvailabilityStatus returns a boolean if a field has been set.

### GetAvailabilityStatusError

`func (o *Application) GetAvailabilityStatusError() string`

GetAvailabilityStatusError returns the AvailabilityStatusError field if non-nil, zero value otherwise.

### GetAvailabilityStatusErrorOk

`func (o *Application) GetAvailabilityStatusErrorOk() (*string, bool)`

GetAvailabilityStatusErrorOk returns a tuple with the AvailabilityStatusError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityStatusError

`func (o *Application) SetAvailabilityStatusError(v string)`

SetAvailabilityStatusError sets AvailabilityStatusError field to given value.

### HasAvailabilityStatusError

`func (o *Application) HasAvailabilityStatusError() bool`

HasAvailabilityStatusError returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Application) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Application) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Application) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Application) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExtra

`func (o *Application) GetExtra() map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *Application) GetExtraOk() (*map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *Application) SetExtra(v map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *Application) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetId

`func (o *Application) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Application) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Application) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Application) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastAvailableAt

`func (o *Application) GetLastAvailableAt() time.Time`

GetLastAvailableAt returns the LastAvailableAt field if non-nil, zero value otherwise.

### GetLastAvailableAtOk

`func (o *Application) GetLastAvailableAtOk() (*time.Time, bool)`

GetLastAvailableAtOk returns a tuple with the LastAvailableAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAvailableAt

`func (o *Application) SetLastAvailableAt(v time.Time)`

SetLastAvailableAt sets LastAvailableAt field to given value.

### HasLastAvailableAt

`func (o *Application) HasLastAvailableAt() bool`

HasLastAvailableAt returns a boolean if a field has been set.

### GetLastCheckedAt

`func (o *Application) GetLastCheckedAt() time.Time`

GetLastCheckedAt returns the LastCheckedAt field if non-nil, zero value otherwise.

### GetLastCheckedAtOk

`func (o *Application) GetLastCheckedAtOk() (*time.Time, bool)`

GetLastCheckedAtOk returns a tuple with the LastCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckedAt

`func (o *Application) SetLastCheckedAt(v time.Time)`

SetLastCheckedAt sets LastCheckedAt field to given value.

### HasLastCheckedAt

`func (o *Application) HasLastCheckedAt() bool`

HasLastCheckedAt returns a boolean if a field has been set.

### GetSourceId

`func (o *Application) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Application) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Application) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *Application) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Application) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Application) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Application) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Application) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


