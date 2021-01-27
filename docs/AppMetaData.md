# AppMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ApplicationTypeId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | ID of the resource | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Payload** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewAppMetaData

`func NewAppMetaData() *AppMetaData`

NewAppMetaData instantiates a new AppMetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppMetaDataWithDefaults

`func NewAppMetaDataWithDefaults() *AppMetaData`

NewAppMetaDataWithDefaults instantiates a new AppMetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AppMetaData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppMetaData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppMetaData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AppMetaData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetApplicationTypeId

`func (o *AppMetaData) GetApplicationTypeId() string`

GetApplicationTypeId returns the ApplicationTypeId field if non-nil, zero value otherwise.

### GetApplicationTypeIdOk

`func (o *AppMetaData) GetApplicationTypeIdOk() (*string, bool)`

GetApplicationTypeIdOk returns a tuple with the ApplicationTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationTypeId

`func (o *AppMetaData) SetApplicationTypeId(v string)`

SetApplicationTypeId sets ApplicationTypeId field to given value.

### HasApplicationTypeId

`func (o *AppMetaData) HasApplicationTypeId() bool`

HasApplicationTypeId returns a boolean if a field has been set.

### GetId

`func (o *AppMetaData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppMetaData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppMetaData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppMetaData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AppMetaData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppMetaData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppMetaData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppMetaData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPayload

`func (o *AppMetaData) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *AppMetaData) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *AppMetaData) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *AppMetaData) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AppMetaData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppMetaData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppMetaData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AppMetaData) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


