# ApplicationTypesCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**CollectionMetadata**](CollectionMetadata.md) |  | [optional] 
**Links** | Pointer to [**CollectionLinks**](CollectionLinks.md) |  | [optional] 
**Data** | Pointer to [**[]ApplicationType**](ApplicationType.md) |  | [optional] 

## Methods

### NewApplicationTypesCollection

`func NewApplicationTypesCollection() *ApplicationTypesCollection`

NewApplicationTypesCollection instantiates a new ApplicationTypesCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationTypesCollectionWithDefaults

`func NewApplicationTypesCollectionWithDefaults() *ApplicationTypesCollection`

NewApplicationTypesCollectionWithDefaults instantiates a new ApplicationTypesCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ApplicationTypesCollection) GetMeta() CollectionMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ApplicationTypesCollection) GetMetaOk() (*CollectionMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ApplicationTypesCollection) SetMeta(v CollectionMetadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ApplicationTypesCollection) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *ApplicationTypesCollection) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApplicationTypesCollection) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApplicationTypesCollection) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApplicationTypesCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetData

`func (o *ApplicationTypesCollection) GetData() []ApplicationType`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApplicationTypesCollection) GetDataOk() (*[]ApplicationType, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApplicationTypesCollection) SetData(v []ApplicationType)`

SetData sets Data field to given value.

### HasData

`func (o *ApplicationTypesCollection) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


