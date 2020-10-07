# Catalog\DefaultApi

All URIs are relative to *https://cloud.redhat.com//api/sources/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckAvailabilitySource**](DefaultApi.md#CheckAvailabilitySource) | **Post** /sources/{id}/check_availability | Checks Availability of a Source
[**CreateApplication**](DefaultApi.md#CreateApplication) | **Post** /applications | Create a new Application
[**CreateApplicationAuthentication**](DefaultApi.md#CreateApplicationAuthentication) | **Post** /application_authentications | Create a new ApplicationAuthentication
[**CreateAuthentication**](DefaultApi.md#CreateAuthentication) | **Post** /authentications | Create a new Authentication
[**CreateEndpoint**](DefaultApi.md#CreateEndpoint) | **Post** /endpoints | Create a new Endpoint
[**CreateSource**](DefaultApi.md#CreateSource) | **Post** /sources | Create a new Source
[**DeleteApplication**](DefaultApi.md#DeleteApplication) | **Delete** /applications/{id} | Delete an existing Application
[**DeleteApplicationAuthentication**](DefaultApi.md#DeleteApplicationAuthentication) | **Delete** /application_authentications/{id} | Delete an existing ApplicationAuthentication
[**DeleteAuthentication**](DefaultApi.md#DeleteAuthentication) | **Delete** /authentications/{id} | Delete an existing Authentication
[**DeleteEndpoint**](DefaultApi.md#DeleteEndpoint) | **Delete** /endpoints/{id} | Delete an existing Endpoint
[**DeleteSource**](DefaultApi.md#DeleteSource) | **Delete** /sources/{id} | Delete an existing Source
[**GetDocumentation**](DefaultApi.md#GetDocumentation) | **Get** /openapi.json | Return this API document in JSON format
[**ListAllApplicationAuthentications**](DefaultApi.md#ListAllApplicationAuthentications) | **Get** /application_authentications | List ApplicationAuthentications
[**ListApplicationAuthentications**](DefaultApi.md#ListApplicationAuthentications) | **Get** /applications/{id}/authentications | List Authentications for Application
[**ListApplicationTypeSources**](DefaultApi.md#ListApplicationTypeSources) | **Get** /application_types/{id}/sources | List Sources for ApplicationType
[**ListApplicationTypes**](DefaultApi.md#ListApplicationTypes) | **Get** /application_types | List ApplicationTypes
[**ListApplications**](DefaultApi.md#ListApplications) | **Get** /applications | List Applications
[**ListAuthentications**](DefaultApi.md#ListAuthentications) | **Get** /authentications | List Authentications
[**ListEndpointAuthentications**](DefaultApi.md#ListEndpointAuthentications) | **Get** /endpoints/{id}/authentications | List Authentications for Endpoint
[**ListEndpoints**](DefaultApi.md#ListEndpoints) | **Get** /endpoints | List Endpoints
[**ListSourceApplicationTypes**](DefaultApi.md#ListSourceApplicationTypes) | **Get** /sources/{id}/application_types | List ApplicationTypes for Source
[**ListSourceApplications**](DefaultApi.md#ListSourceApplications) | **Get** /sources/{id}/applications | List Applications for Source
[**ListSourceEndpoints**](DefaultApi.md#ListSourceEndpoints) | **Get** /sources/{id}/endpoints | List Endpoints for Source
[**ListSourceTypeSources**](DefaultApi.md#ListSourceTypeSources) | **Get** /source_types/{id}/sources | List Sources for SourceType
[**ListSourceTypes**](DefaultApi.md#ListSourceTypes) | **Get** /source_types | List SourceTypes
[**ListSources**](DefaultApi.md#ListSources) | **Get** /sources | List Sources
[**PostGraphQL**](DefaultApi.md#PostGraphQL) | **Post** /graphql | Perform a GraphQL Query
[**ShowApplication**](DefaultApi.md#ShowApplication) | **Get** /applications/{id} | Show an existing Application
[**ShowApplicationAuthentication**](DefaultApi.md#ShowApplicationAuthentication) | **Get** /application_authentications/{id} | Show an existing ApplicationAuthentication
[**ShowApplicationType**](DefaultApi.md#ShowApplicationType) | **Get** /application_types/{id} | Show an existing ApplicationType
[**ShowAuthentication**](DefaultApi.md#ShowAuthentication) | **Get** /authentications/{id} | Show an existing Authentication
[**ShowEndpoint**](DefaultApi.md#ShowEndpoint) | **Get** /endpoints/{id} | Show an existing Endpoint
[**ShowSource**](DefaultApi.md#ShowSource) | **Get** /sources/{id} | Show an existing Source
[**ShowSourceType**](DefaultApi.md#ShowSourceType) | **Get** /source_types/{id} | Show an existing SourceType
[**UpdateApplication**](DefaultApi.md#UpdateApplication) | **Patch** /applications/{id} | Update an existing Application
[**UpdateApplicationAuthentication**](DefaultApi.md#UpdateApplicationAuthentication) | **Patch** /application_authentications/{id} | Update an existing ApplicationAuthentication
[**UpdateAuthentication**](DefaultApi.md#UpdateAuthentication) | **Patch** /authentications/{id} | Update an existing Authentication
[**UpdateEndpoint**](DefaultApi.md#UpdateEndpoint) | **Patch** /endpoints/{id} | Update an existing Endpoint
[**UpdateSource**](DefaultApi.md#UpdateSource) | **Patch** /sources/{id} | Update an existing Source



## CheckAvailabilitySource

> CheckAvailabilitySource(ctx, id).Execute()

Checks Availability of a Source



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CheckAvailabilitySource(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CheckAvailabilitySource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckAvailabilitySourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplication

> Application CreateApplication(ctx).Application(application).Execute()

Create a new Application



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    application := openapiclient.Application{ApplicationTypeId: "ApplicationTypeId_example", AvailabilityStatus: "AvailabilityStatus_example", AvailabilityStatusError: "AvailabilityStatusError_example", CreatedAt: "TODO", Id: "Id_example", LastAvailableAt: "TODO", LastCheckedAt: "TODO", SourceId: "SourceId_example", UpdatedAt: "TODO"} // Application | Application attributes to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateApplication(context.Background()).Application(application).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplication`: Application
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | [**Application**](Application.md) | Application attributes to create | 

### Return type

[**Application**](Application.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicationAuthentication

> ApplicationAuthentication CreateApplicationAuthentication(ctx).ApplicationAuthentication(applicationAuthentication).Execute()

Create a new ApplicationAuthentication



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicationAuthentication := openapiclient.ApplicationAuthentication{ApplicationId: "ApplicationId_example", AuthenticationId: "AuthenticationId_example", CreatedAt: "TODO", Id: "Id_example", UpdatedAt: "TODO"} // ApplicationAuthentication | ApplicationAuthentication attributes to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateApplicationAuthentication(context.Background()).ApplicationAuthentication(applicationAuthentication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateApplicationAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationAuthentication`: ApplicationAuthentication
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateApplicationAuthentication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationAuthentication** | [**ApplicationAuthentication**](ApplicationAuthentication.md) | ApplicationAuthentication attributes to create | 

### Return type

[**ApplicationAuthentication**](ApplicationAuthentication.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthentication

> Authentication CreateAuthentication(ctx).Authentication(authentication).Execute()

Create a new Authentication



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authentication := openapiclient.Authentication{Authtype: "Authtype_example", AvailabilityStatus: "AvailabilityStatus_example", AvailabilityStatusError: "AvailabilityStatusError_example", Extra: openapiclient.Authentication_extra{Azure: openapiclient.Authentication_extra_azure{TenantId: "TenantId_example"}}, Id: "Id_example", LastAvailableAt: "TODO", LastCheckedAt: "TODO", Name: "Name_example", Password: "Password_example", ResourceId: "ResourceId_example", ResourceType: "ResourceType_example", Status: "Status_example", StatusDetails: "StatusDetails_example", Username: "Username_example"} // Authentication | Authentication attributes to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateAuthentication(context.Background()).Authentication(authentication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthentication`: Authentication
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAuthentication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authentication** | [**Authentication**](Authentication.md) | Authentication attributes to create | 

### Return type

[**Authentication**](Authentication.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEndpoint

> Endpoint CreateEndpoint(ctx).Endpoint(endpoint).Execute()

Create a new Endpoint



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    endpoint := openapiclient.Endpoint{AvailabilityStatus: "AvailabilityStatus_example", AvailabilityStatusError: "AvailabilityStatusError_example", CertificateAuthority: "CertificateAuthority_example", CreatedAt: "TODO", Default: false, Host: "Host_example", Id: "Id_example", LastAvailableAt: "TODO", LastCheckedAt: "TODO", Path: "Path_example", Port: 123, ReceptorNode: "ReceptorNode_example", Role: "Role_example", Scheme: "Scheme_example", SourceId: "SourceId_example", UpdatedAt: "TODO", VerifySsl: true} // Endpoint | Endpoint attributes to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateEndpoint(context.Background()).Endpoint(endpoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEndpoint`: Endpoint
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpoint** | [**Endpoint**](Endpoint.md) | Endpoint attributes to create | 

### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSource

> Source CreateSource(ctx).Source(source).Execute()

Create a new Source



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    source := openapiclient.Source{AvailabilityStatus: "AvailabilityStatus_example", CreatedAt: "TODO", Id: "Id_example", Imported: "Imported_example", LastAvailableAt: "TODO", LastCheckedAt: "TODO", Name: "Name_example", SourceRef: "SourceRef_example", SourceTypeId: "SourceTypeId_example", Uid: "Uid_example", UpdatedAt: "TODO", Version: "Version_example"} // Source | Source attributes to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSource(context.Background()).Source(source).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSource`: Source
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **source** | [**Source**](Source.md) | Source attributes to create | 

### Return type

[**Source**](Source.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplication

> DeleteApplication(ctx, id).Execute()

Delete an existing Application



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteApplication(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationAuthentication

> DeleteApplicationAuthentication(ctx, id).Execute()

Delete an existing ApplicationAuthentication



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteApplicationAuthentication(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteApplicationAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthentication

> DeleteAuthentication(ctx, id).Execute()

Delete an existing Authentication



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteAuthentication(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEndpoint

> DeleteEndpoint(ctx, id).Execute()

Delete an existing Endpoint



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteEndpoint(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSource

> DeleteSource(ctx, id).Execute()

Delete an existing Source



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSource(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentation

> map[string]interface{} GetDocumentation(ctx).Execute()

Return this API document in JSON format

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetDocumentation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDocumentation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentation`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDocumentation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentationRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllApplicationAuthentications

> ApplicationAuthenticationsCollection ListAllApplicationAuthentications(ctx).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()

List ApplicationAuthentications



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := 987 // int32 | The numbers of items to return per page. (optional) (default to 100)
    offset := 987 // int32 | The number of items to skip before starting to collect the result set. (optional) (default to 0)
    filter := TODO // map[string]interface{} | Filter for querying collections. (optional)
    sortBy := TODO // map[string]interface{} | The list of attribute and order to sort the result set by. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListAllApplicationAuthentications(context.Background()).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAllApplicationAuthentications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllApplicationAuthentications`: ApplicationAuthenticationsCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAllApplicationAuthentications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllApplicationAuthenticationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The numbers of items to return per page. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | [default to 0]
 **filter** | [**map[string]interface{}**](.md) | Filter for querying collections. | 
 **sortBy** | [**map[string]interface{}**](.md) | The list of attribute and order to sort the result set by. | 

### Return type

[**ApplicationAuthenticationsCollection**](ApplicationAuthenticationsCollection.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationAuthentications

> AuthenticationsCollection ListApplicationAuthentications(ctx, id).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()

List Authentications for Application



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource
    limit := 987 // int32 | The numbers of items to return per page. (optional) (default to 100)
    offset := 987 // int32 | The number of items to skip before starting to collect the result set. (optional) (default to 0)
    filter := TODO // map[string]interface{} | Filter for querying collections. (optional)
    sortBy := TODO // map[string]interface{} | The list of attribute and order to sort the result set by. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListApplicationAuthentications(context.Background(), id).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListApplicationAuthentications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationAuthentications`: AuthenticationsCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListApplicationAuthentications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationAuthenticationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The numbers of items to return per page. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | [default to 0]
 **filter** | [**map[string]interface{}**](.md) | Filter for querying collections. | 
 **sortBy** | [**map[string]interface{}**](.md) | The list of attribute and order to sort the result set by. | 

### Return type

[**AuthenticationsCollection**](AuthenticationsCollection.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationTypeSources

> SourcesCollection ListApplicationTypeSources(ctx, id).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()

List Sources for ApplicationType



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource
    limit := 987 // int32 | The numbers of items to return per page. (optional) (default to 100)
    offset := 987 // int32 | The number of items to skip before starting to collect the result set. (optional) (default to 0)
    filter := TODO // map[string]interface{} | Filter for querying collections. (optional)
    sortBy := TODO // map[string]interface{} | The list of attribute and order to sort the result set by. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListApplicationTypeSources(context.Background(), id).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListApplicationTypeSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationTypeSources`: SourcesCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListApplicationTypeSources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationTypeSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The numbers of items to return per page. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | [default to 0]
 **filter** | [**map[string]interface{}**](.md) | Filter for querying collections. | 
 **sortBy** | [**map[string]interface{}**](.md) | The list of attribute and order to sort the result set by. | 

### Return type

[**SourcesCollection**](SourcesCollection.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationTypes

> ApplicationTypesCollection ListApplicationTypes(ctx).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()

List ApplicationTypes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := 987 // int32 | The numbers of items to return per page. (optional) (default to 100)
    offset := 987 // int32 | The number of items to skip before starting to collect the result set. (optional) (default to 0)
    filter := TODO // map[string]interface{} | Filter for querying collections. (optional)
    sortBy := TODO // map[string]interface{} | The list of attribute and order to sort the result set by. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListApplicationTypes(context.Background()).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListApplicationTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationTypes`: ApplicationTypesCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListApplicationTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The numbers of items to return per page. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | [default to 0]
 **filter** | [**map[string]interface{}**](.md) | Filter for querying collections. | 
 **sortBy** | [**map[string]interface{}**](.md) | The list of attribute and order to sort the result set by. | 

### Return type

[**ApplicationTypesCollection**](ApplicationTypesCollection.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplications

> ApplicationsCollection ListApplications(ctx).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()

List Applications



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := 987 // int32 | The numbers of items to return per page. (optional) (default to 100)
    offset := 987 // int32 | The number of items to skip before starting to collect the result set. (optional) (default to 0)
    filter := TODO // map[string]interface{} | Filter for querying collections. (optional)
    sortBy := TODO // map[string]interface{} | The list of attribute and order to sort the result set by. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListApplications(context.Background()).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplications`: ApplicationsCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The numbers of items to return per page. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | [default to 0]
 **filter** | [**map[string]interface{}**](.md) | Filter for querying collections. | 
 **sortBy** | [**map[string]interface{}**](.md) | The list of attribute and order to sort the result set by. | 

### Return type

[**ApplicationsCollection**](ApplicationsCollection.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthentications

> AuthenticationsCollection ListAuthentications(ctx).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()

List Authentications



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := 987 // int32 | The numbers of items to return per page. (optional) (default to 100)
    offset := 987 // int32 | The number of items to skip before starting to collect the result set. (optional) (default to 0)
    filter := TODO // map[string]interface{} | Filter for querying collections. (optional)
    sortBy := TODO // map[string]interface{} | The list of attribute and order to sort the result set by. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListAuthentications(context.Background()).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAuthentications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuthentications`: AuthenticationsCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAuthentications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAuthenticationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The numbers of items to return per page. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | [default to 0]
 **filter** | [**map[string]interface{}**](.md) | Filter for querying collections. | 
 **sortBy** | [**map[string]interface{}**](.md) | The list of attribute and order to sort the result set by. | 

### Return type

[**AuthenticationsCollection**](AuthenticationsCollection.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEndpointAuthentications

> AuthenticationsCollection ListEndpointAuthentications(ctx, id).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()

List Authentications for Endpoint



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource
    limit := 987 // int32 | The numbers of items to return per page. (optional) (default to 100)
    offset := 987 // int32 | The number of items to skip before starting to collect the result set. (optional) (default to 0)
    filter := TODO // map[string]interface{} | Filter for querying collections. (optional)
    sortBy := TODO // map[string]interface{} | The list of attribute and order to sort the result set by. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListEndpointAuthentications(context.Background(), id).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListEndpointAuthentications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEndpointAuthentications`: AuthenticationsCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListEndpointAuthentications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEndpointAuthenticationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The numbers of items to return per page. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | [default to 0]
 **filter** | [**map[string]interface{}**](.md) | Filter for querying collections. | 
 **sortBy** | [**map[string]interface{}**](.md) | The list of attribute and order to sort the result set by. | 

### Return type

[**AuthenticationsCollection**](AuthenticationsCollection.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEndpoints

> EndpointsCollection ListEndpoints(ctx).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()

List Endpoints



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := 987 // int32 | The numbers of items to return per page. (optional) (default to 100)
    offset := 987 // int32 | The number of items to skip before starting to collect the result set. (optional) (default to 0)
    filter := TODO // map[string]interface{} | Filter for querying collections. (optional)
    sortBy := TODO // map[string]interface{} | The list of attribute and order to sort the result set by. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListEndpoints(context.Background()).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEndpoints`: EndpointsCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListEndpoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The numbers of items to return per page. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | [default to 0]
 **filter** | [**map[string]interface{}**](.md) | Filter for querying collections. | 
 **sortBy** | [**map[string]interface{}**](.md) | The list of attribute and order to sort the result set by. | 

### Return type

[**EndpointsCollection**](EndpointsCollection.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSourceApplicationTypes

> ApplicationTypesCollection ListSourceApplicationTypes(ctx, id).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()

List ApplicationTypes for Source



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource
    limit := 987 // int32 | The numbers of items to return per page. (optional) (default to 100)
    offset := 987 // int32 | The number of items to skip before starting to collect the result set. (optional) (default to 0)
    filter := TODO // map[string]interface{} | Filter for querying collections. (optional)
    sortBy := TODO // map[string]interface{} | The list of attribute and order to sort the result set by. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSourceApplicationTypes(context.Background(), id).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSourceApplicationTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSourceApplicationTypes`: ApplicationTypesCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSourceApplicationTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSourceApplicationTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The numbers of items to return per page. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | [default to 0]
 **filter** | [**map[string]interface{}**](.md) | Filter for querying collections. | 
 **sortBy** | [**map[string]interface{}**](.md) | The list of attribute and order to sort the result set by. | 

### Return type

[**ApplicationTypesCollection**](ApplicationTypesCollection.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSourceApplications

> ApplicationsCollection ListSourceApplications(ctx, id).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()

List Applications for Source



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource
    limit := 987 // int32 | The numbers of items to return per page. (optional) (default to 100)
    offset := 987 // int32 | The number of items to skip before starting to collect the result set. (optional) (default to 0)
    filter := TODO // map[string]interface{} | Filter for querying collections. (optional)
    sortBy := TODO // map[string]interface{} | The list of attribute and order to sort the result set by. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSourceApplications(context.Background(), id).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSourceApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSourceApplications`: ApplicationsCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSourceApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSourceApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The numbers of items to return per page. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | [default to 0]
 **filter** | [**map[string]interface{}**](.md) | Filter for querying collections. | 
 **sortBy** | [**map[string]interface{}**](.md) | The list of attribute and order to sort the result set by. | 

### Return type

[**ApplicationsCollection**](ApplicationsCollection.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSourceEndpoints

> EndpointsCollection ListSourceEndpoints(ctx, id).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()

List Endpoints for Source



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource
    limit := 987 // int32 | The numbers of items to return per page. (optional) (default to 100)
    offset := 987 // int32 | The number of items to skip before starting to collect the result set. (optional) (default to 0)
    filter := TODO // map[string]interface{} | Filter for querying collections. (optional)
    sortBy := TODO // map[string]interface{} | The list of attribute and order to sort the result set by. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSourceEndpoints(context.Background(), id).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSourceEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSourceEndpoints`: EndpointsCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSourceEndpoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSourceEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The numbers of items to return per page. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | [default to 0]
 **filter** | [**map[string]interface{}**](.md) | Filter for querying collections. | 
 **sortBy** | [**map[string]interface{}**](.md) | The list of attribute and order to sort the result set by. | 

### Return type

[**EndpointsCollection**](EndpointsCollection.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSourceTypeSources

> SourcesCollection ListSourceTypeSources(ctx, id).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()

List Sources for SourceType



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource
    limit := 987 // int32 | The numbers of items to return per page. (optional) (default to 100)
    offset := 987 // int32 | The number of items to skip before starting to collect the result set. (optional) (default to 0)
    filter := TODO // map[string]interface{} | Filter for querying collections. (optional)
    sortBy := TODO // map[string]interface{} | The list of attribute and order to sort the result set by. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSourceTypeSources(context.Background(), id).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSourceTypeSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSourceTypeSources`: SourcesCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSourceTypeSources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSourceTypeSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The numbers of items to return per page. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | [default to 0]
 **filter** | [**map[string]interface{}**](.md) | Filter for querying collections. | 
 **sortBy** | [**map[string]interface{}**](.md) | The list of attribute and order to sort the result set by. | 

### Return type

[**SourcesCollection**](SourcesCollection.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSourceTypes

> SourceTypesCollection ListSourceTypes(ctx).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()

List SourceTypes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := 987 // int32 | The numbers of items to return per page. (optional) (default to 100)
    offset := 987 // int32 | The number of items to skip before starting to collect the result set. (optional) (default to 0)
    filter := TODO // map[string]interface{} | Filter for querying collections. (optional)
    sortBy := TODO // map[string]interface{} | The list of attribute and order to sort the result set by. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSourceTypes(context.Background()).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSourceTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSourceTypes`: SourceTypesCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSourceTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSourceTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The numbers of items to return per page. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | [default to 0]
 **filter** | [**map[string]interface{}**](.md) | Filter for querying collections. | 
 **sortBy** | [**map[string]interface{}**](.md) | The list of attribute and order to sort the result set by. | 

### Return type

[**SourceTypesCollection**](SourceTypesCollection.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSources

> SourcesCollection ListSources(ctx).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()

List Sources



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := 987 // int32 | The numbers of items to return per page. (optional) (default to 100)
    offset := 987 // int32 | The number of items to skip before starting to collect the result set. (optional) (default to 0)
    filter := TODO // map[string]interface{} | Filter for querying collections. (optional)
    sortBy := TODO // map[string]interface{} | The list of attribute and order to sort the result set by. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSources(context.Background()).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSources`: SourcesCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The numbers of items to return per page. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | [default to 0]
 **filter** | [**map[string]interface{}**](.md) | Filter for querying collections. | 
 **sortBy** | [**map[string]interface{}**](.md) | The list of attribute and order to sort the result set by. | 

### Return type

[**SourcesCollection**](SourcesCollection.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGraphQL

> GraphQLResponse PostGraphQL(ctx).GraphQLRequest(graphQLRequest).Execute()

Perform a GraphQL Query



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    graphQLRequest := openapiclient.GraphQLRequest{Query: "Query_example", OperationName: "OperationName_example", Variables: 123} // GraphQLRequest | GraphQL Query Request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.PostGraphQL(context.Background()).GraphQLRequest(graphQLRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostGraphQL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostGraphQL`: GraphQLResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostGraphQL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostGraphQLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphQLRequest** | [**GraphQLRequest**](GraphQLRequest.md) | GraphQL Query Request | 

### Return type

[**GraphQLResponse**](GraphQLResponse.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowApplication

> Application ShowApplication(ctx, id).Execute()

Show an existing Application



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ShowApplication(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowApplication`: Application
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Application**](Application.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowApplicationAuthentication

> ApplicationAuthentication ShowApplicationAuthentication(ctx, id).Execute()

Show an existing ApplicationAuthentication



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ShowApplicationAuthentication(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowApplicationAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowApplicationAuthentication`: ApplicationAuthentication
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowApplicationAuthentication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowApplicationAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationAuthentication**](ApplicationAuthentication.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowApplicationType

> ApplicationType ShowApplicationType(ctx, id).Execute()

Show an existing ApplicationType



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ShowApplicationType(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowApplicationType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowApplicationType`: ApplicationType
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowApplicationType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowApplicationTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationType**](ApplicationType.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowAuthentication

> Authentication ShowAuthentication(ctx, id).Execute()

Show an existing Authentication



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ShowAuthentication(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowAuthentication`: Authentication
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowAuthentication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Authentication**](Authentication.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowEndpoint

> Endpoint ShowEndpoint(ctx, id).Execute()

Show an existing Endpoint



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ShowEndpoint(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowEndpoint`: Endpoint
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowSource

> Source ShowSource(ctx, id).Execute()

Show an existing Source



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ShowSource(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowSource`: Source
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Source**](Source.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowSourceType

> SourceType ShowSourceType(ctx, id).Execute()

Show an existing SourceType



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ShowSourceType(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowSourceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowSourceType`: SourceType
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowSourceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowSourceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SourceType**](SourceType.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplication

> UpdateApplication(ctx, id).Application(application).Execute()

Update an existing Application



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource
    application := openapiclient.Application{ApplicationTypeId: "ApplicationTypeId_example", AvailabilityStatus: "AvailabilityStatus_example", AvailabilityStatusError: "AvailabilityStatusError_example", CreatedAt: "TODO", Id: "Id_example", LastAvailableAt: "TODO", LastCheckedAt: "TODO", SourceId: "SourceId_example", UpdatedAt: "TODO"} // Application | Application attributes to update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateApplication(context.Background(), id).Application(application).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **application** | [**Application**](Application.md) | Application attributes to update | 

### Return type

 (empty response body)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationAuthentication

> UpdateApplicationAuthentication(ctx, id).ApplicationAuthentication(applicationAuthentication).Execute()

Update an existing ApplicationAuthentication



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource
    applicationAuthentication := openapiclient.ApplicationAuthentication{ApplicationId: "ApplicationId_example", AuthenticationId: "AuthenticationId_example", CreatedAt: "TODO", Id: "Id_example", UpdatedAt: "TODO"} // ApplicationAuthentication | ApplicationAuthentication attributes to update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateApplicationAuthentication(context.Background(), id).ApplicationAuthentication(applicationAuthentication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateApplicationAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationAuthentication** | [**ApplicationAuthentication**](ApplicationAuthentication.md) | ApplicationAuthentication attributes to update | 

### Return type

 (empty response body)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthentication

> UpdateAuthentication(ctx, id).Authentication(authentication).Execute()

Update an existing Authentication



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource
    authentication := openapiclient.Authentication{Authtype: "Authtype_example", AvailabilityStatus: "AvailabilityStatus_example", AvailabilityStatusError: "AvailabilityStatusError_example", Extra: openapiclient.Authentication_extra{Azure: openapiclient.Authentication_extra_azure{TenantId: "TenantId_example"}}, Id: "Id_example", LastAvailableAt: "TODO", LastCheckedAt: "TODO", Name: "Name_example", Password: "Password_example", ResourceId: "ResourceId_example", ResourceType: "ResourceType_example", Status: "Status_example", StatusDetails: "StatusDetails_example", Username: "Username_example"} // Authentication | Authentication attributes to update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateAuthentication(context.Background(), id).Authentication(authentication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authentication** | [**Authentication**](Authentication.md) | Authentication attributes to update | 

### Return type

 (empty response body)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndpoint

> UpdateEndpoint(ctx, id).Endpoint(endpoint).Execute()

Update an existing Endpoint



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource
    endpoint := openapiclient.Endpoint{AvailabilityStatus: "AvailabilityStatus_example", AvailabilityStatusError: "AvailabilityStatusError_example", CertificateAuthority: "CertificateAuthority_example", CreatedAt: "TODO", Default: false, Host: "Host_example", Id: "Id_example", LastAvailableAt: "TODO", LastCheckedAt: "TODO", Path: "Path_example", Port: 123, ReceptorNode: "ReceptorNode_example", Role: "Role_example", Scheme: "Scheme_example", SourceId: "SourceId_example", UpdatedAt: "TODO", VerifySsl: true} // Endpoint | Endpoint attributes to update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateEndpoint(context.Background(), id).Endpoint(endpoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endpoint** | [**Endpoint**](Endpoint.md) | Endpoint attributes to update | 

### Return type

 (empty response body)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSource

> UpdateSource(ctx, id).Source(source).Execute()

Update an existing Source



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the resource
    source := openapiclient.Source{AvailabilityStatus: "AvailabilityStatus_example", CreatedAt: "TODO", Id: "Id_example", Imported: "Imported_example", LastAvailableAt: "TODO", LastCheckedAt: "TODO", Name: "Name_example", SourceRef: "SourceRef_example", SourceTypeId: "SourceTypeId_example", Uid: "Uid_example", UpdatedAt: "TODO", Version: "Version_example"} // Source | Source attributes to update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSource(context.Background(), id).Source(source).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | [**Source**](Source.md) | Source attributes to update | 

### Return type

 (empty response body)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

