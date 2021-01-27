# Authentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authtype** | Pointer to **string** |  | [optional] [readonly] 
**AvailabilityStatus** | Pointer to **string** |  | [optional] 
**AvailabilityStatusError** | Pointer to **string** |  | [optional] 
**Extra** | Pointer to [**AuthenticationExtra**](Authentication_extra.md) |  | [optional] 
**Id** | Pointer to **string** | ID of the resource | [optional] [readonly] 
**LastAvailableAt** | Pointer to **time.Time** |  | [optional] 
**LastCheckedAt** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**ResourceId** | Pointer to **string** | ID of the resource | [optional] [readonly] 
**ResourceType** | Pointer to **string** |  | [optional] 
**SourceId** | Pointer to **string** | ID of the resource | [optional] [readonly] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDetails** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthentication

`func NewAuthentication() *Authentication`

NewAuthentication instantiates a new Authentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationWithDefaults

`func NewAuthenticationWithDefaults() *Authentication`

NewAuthenticationWithDefaults instantiates a new Authentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthtype

`func (o *Authentication) GetAuthtype() string`

GetAuthtype returns the Authtype field if non-nil, zero value otherwise.

### GetAuthtypeOk

`func (o *Authentication) GetAuthtypeOk() (*string, bool)`

GetAuthtypeOk returns a tuple with the Authtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthtype

`func (o *Authentication) SetAuthtype(v string)`

SetAuthtype sets Authtype field to given value.

### HasAuthtype

`func (o *Authentication) HasAuthtype() bool`

HasAuthtype returns a boolean if a field has been set.

### GetAvailabilityStatus

`func (o *Authentication) GetAvailabilityStatus() string`

GetAvailabilityStatus returns the AvailabilityStatus field if non-nil, zero value otherwise.

### GetAvailabilityStatusOk

`func (o *Authentication) GetAvailabilityStatusOk() (*string, bool)`

GetAvailabilityStatusOk returns a tuple with the AvailabilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityStatus

`func (o *Authentication) SetAvailabilityStatus(v string)`

SetAvailabilityStatus sets AvailabilityStatus field to given value.

### HasAvailabilityStatus

`func (o *Authentication) HasAvailabilityStatus() bool`

HasAvailabilityStatus returns a boolean if a field has been set.

### GetAvailabilityStatusError

`func (o *Authentication) GetAvailabilityStatusError() string`

GetAvailabilityStatusError returns the AvailabilityStatusError field if non-nil, zero value otherwise.

### GetAvailabilityStatusErrorOk

`func (o *Authentication) GetAvailabilityStatusErrorOk() (*string, bool)`

GetAvailabilityStatusErrorOk returns a tuple with the AvailabilityStatusError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityStatusError

`func (o *Authentication) SetAvailabilityStatusError(v string)`

SetAvailabilityStatusError sets AvailabilityStatusError field to given value.

### HasAvailabilityStatusError

`func (o *Authentication) HasAvailabilityStatusError() bool`

HasAvailabilityStatusError returns a boolean if a field has been set.

### GetExtra

`func (o *Authentication) GetExtra() AuthenticationExtra`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *Authentication) GetExtraOk() (*AuthenticationExtra, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *Authentication) SetExtra(v AuthenticationExtra)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *Authentication) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetId

`func (o *Authentication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Authentication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Authentication) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Authentication) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastAvailableAt

`func (o *Authentication) GetLastAvailableAt() time.Time`

GetLastAvailableAt returns the LastAvailableAt field if non-nil, zero value otherwise.

### GetLastAvailableAtOk

`func (o *Authentication) GetLastAvailableAtOk() (*time.Time, bool)`

GetLastAvailableAtOk returns a tuple with the LastAvailableAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAvailableAt

`func (o *Authentication) SetLastAvailableAt(v time.Time)`

SetLastAvailableAt sets LastAvailableAt field to given value.

### HasLastAvailableAt

`func (o *Authentication) HasLastAvailableAt() bool`

HasLastAvailableAt returns a boolean if a field has been set.

### GetLastCheckedAt

`func (o *Authentication) GetLastCheckedAt() time.Time`

GetLastCheckedAt returns the LastCheckedAt field if non-nil, zero value otherwise.

### GetLastCheckedAtOk

`func (o *Authentication) GetLastCheckedAtOk() (*time.Time, bool)`

GetLastCheckedAtOk returns a tuple with the LastCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckedAt

`func (o *Authentication) SetLastCheckedAt(v time.Time)`

SetLastCheckedAt sets LastCheckedAt field to given value.

### HasLastCheckedAt

`func (o *Authentication) HasLastCheckedAt() bool`

HasLastCheckedAt returns a boolean if a field has been set.

### GetName

`func (o *Authentication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Authentication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Authentication) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Authentication) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *Authentication) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Authentication) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Authentication) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Authentication) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetResourceId

`func (o *Authentication) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *Authentication) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *Authentication) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *Authentication) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceType

`func (o *Authentication) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *Authentication) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *Authentication) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *Authentication) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetSourceId

`func (o *Authentication) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Authentication) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Authentication) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *Authentication) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetStatus

`func (o *Authentication) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Authentication) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Authentication) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Authentication) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetails

`func (o *Authentication) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *Authentication) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *Authentication) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *Authentication) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

### GetUsername

`func (o *Authentication) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Authentication) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Authentication) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Authentication) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


