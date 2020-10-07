# ApplicationAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **string** | ID of the resource | [optional] [readonly] 
**AuthenticationId** | Pointer to **string** | ID of the resource | [optional] [readonly] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] [readonly] 
**Id** | Pointer to **string** | ID of the resource | [optional] [readonly] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] [readonly] 

## Methods

### NewApplicationAuthentication

`func NewApplicationAuthentication() *ApplicationAuthentication`

NewApplicationAuthentication instantiates a new ApplicationAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationAuthenticationWithDefaults

`func NewApplicationAuthenticationWithDefaults() *ApplicationAuthentication`

NewApplicationAuthenticationWithDefaults instantiates a new ApplicationAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *ApplicationAuthentication) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationAuthentication) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationAuthentication) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ApplicationAuthentication) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetAuthenticationId

`func (o *ApplicationAuthentication) GetAuthenticationId() string`

GetAuthenticationId returns the AuthenticationId field if non-nil, zero value otherwise.

### GetAuthenticationIdOk

`func (o *ApplicationAuthentication) GetAuthenticationIdOk() (*string, bool)`

GetAuthenticationIdOk returns a tuple with the AuthenticationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationId

`func (o *ApplicationAuthentication) SetAuthenticationId(v string)`

SetAuthenticationId sets AuthenticationId field to given value.

### HasAuthenticationId

`func (o *ApplicationAuthentication) HasAuthenticationId() bool`

HasAuthenticationId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApplicationAuthentication) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicationAuthentication) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicationAuthentication) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApplicationAuthentication) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *ApplicationAuthentication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationAuthentication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationAuthentication) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationAuthentication) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ApplicationAuthentication) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ApplicationAuthentication) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ApplicationAuthentication) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ApplicationAuthentication) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


