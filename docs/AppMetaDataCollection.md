# AppMetaDataCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**CollectionMetadata**](CollectionMetadata.md) |  | [optional] 
**Links** | Pointer to [**CollectionLinks**](CollectionLinks.md) |  | [optional] 
**Data** | Pointer to [**[]AppMetaData**](AppMetaData.md) |  | [optional] 

## Methods

### NewAppMetaDataCollection

`func NewAppMetaDataCollection() *AppMetaDataCollection`

NewAppMetaDataCollection instantiates a new AppMetaDataCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppMetaDataCollectionWithDefaults

`func NewAppMetaDataCollectionWithDefaults() *AppMetaDataCollection`

NewAppMetaDataCollectionWithDefaults instantiates a new AppMetaDataCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *AppMetaDataCollection) GetMeta() CollectionMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AppMetaDataCollection) GetMetaOk() (*CollectionMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AppMetaDataCollection) SetMeta(v CollectionMetadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AppMetaDataCollection) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *AppMetaDataCollection) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppMetaDataCollection) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppMetaDataCollection) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AppMetaDataCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetData

`func (o *AppMetaDataCollection) GetData() []AppMetaData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppMetaDataCollection) GetDataOk() (*[]AppMetaData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppMetaDataCollection) SetData(v []AppMetaData)`

SetData sets Data field to given value.

### HasData

`func (o *AppMetaDataCollection) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


