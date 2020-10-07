# AuthenticationExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Azure** | Pointer to [**AuthenticationExtraAzure**](Authentication_extra_azure.md) |  | [optional] 

## Methods

### NewAuthenticationExtra

`func NewAuthenticationExtra() *AuthenticationExtra`

NewAuthenticationExtra instantiates a new AuthenticationExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationExtraWithDefaults

`func NewAuthenticationExtraWithDefaults() *AuthenticationExtra`

NewAuthenticationExtraWithDefaults instantiates a new AuthenticationExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzure

`func (o *AuthenticationExtra) GetAzure() AuthenticationExtraAzure`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *AuthenticationExtra) GetAzureOk() (*AuthenticationExtraAzure, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *AuthenticationExtra) SetAzure(v AuthenticationExtraAzure)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *AuthenticationExtra) HasAzure() bool`

HasAzure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


