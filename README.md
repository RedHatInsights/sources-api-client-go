# Go API client for sourcesapi

Sources

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./sourcesapi"
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://cloud.stage.redhat.com//api/sources/v3.0*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**CheckAvailabilitySource**](docs/DefaultApi.md#checkavailabilitysource) | **Post** /sources/{id}/check_availability | Checks Availability of a Source
*DefaultApi* | [**CreateApplication**](docs/DefaultApi.md#createapplication) | **Post** /applications | Create a new Application
*DefaultApi* | [**CreateApplicationAuthentication**](docs/DefaultApi.md#createapplicationauthentication) | **Post** /application_authentications | Create a new ApplicationAuthentication
*DefaultApi* | [**CreateAuthentication**](docs/DefaultApi.md#createauthentication) | **Post** /authentications | Create a new Authentication
*DefaultApi* | [**CreateEndpoint**](docs/DefaultApi.md#createendpoint) | **Post** /endpoints | Create a new Endpoint
*DefaultApi* | [**CreateSource**](docs/DefaultApi.md#createsource) | **Post** /sources | Create a new Source
*DefaultApi* | [**DeleteApplication**](docs/DefaultApi.md#deleteapplication) | **Delete** /applications/{id} | Delete an existing Application
*DefaultApi* | [**DeleteApplicationAuthentication**](docs/DefaultApi.md#deleteapplicationauthentication) | **Delete** /application_authentications/{id} | Delete an existing ApplicationAuthentication
*DefaultApi* | [**DeleteAuthentication**](docs/DefaultApi.md#deleteauthentication) | **Delete** /authentications/{id} | Delete an existing Authentication
*DefaultApi* | [**DeleteEndpoint**](docs/DefaultApi.md#deleteendpoint) | **Delete** /endpoints/{id} | Delete an existing Endpoint
*DefaultApi* | [**DeleteSource**](docs/DefaultApi.md#deletesource) | **Delete** /sources/{id} | Delete an existing Source
*DefaultApi* | [**GetDocumentation**](docs/DefaultApi.md#getdocumentation) | **Get** /openapi.json | Return this API document in JSON format
*DefaultApi* | [**ListAllApplicationAuthentications**](docs/DefaultApi.md#listallapplicationauthentications) | **Get** /application_authentications | List ApplicationAuthentications
*DefaultApi* | [**ListApplicationAuthentications**](docs/DefaultApi.md#listapplicationauthentications) | **Get** /applications/{id}/authentications | List Authentications for Application
*DefaultApi* | [**ListApplicationTypeSources**](docs/DefaultApi.md#listapplicationtypesources) | **Get** /application_types/{id}/sources | List Sources for ApplicationType
*DefaultApi* | [**ListApplicationTypes**](docs/DefaultApi.md#listapplicationtypes) | **Get** /application_types | List ApplicationTypes
*DefaultApi* | [**ListApplications**](docs/DefaultApi.md#listapplications) | **Get** /applications | List Applications
*DefaultApi* | [**ListAuthentications**](docs/DefaultApi.md#listauthentications) | **Get** /authentications | List Authentications
*DefaultApi* | [**ListEndpointAuthentications**](docs/DefaultApi.md#listendpointauthentications) | **Get** /endpoints/{id}/authentications | List Authentications for Endpoint
*DefaultApi* | [**ListEndpoints**](docs/DefaultApi.md#listendpoints) | **Get** /endpoints | List Endpoints
*DefaultApi* | [**ListSourceApplicationTypes**](docs/DefaultApi.md#listsourceapplicationtypes) | **Get** /sources/{id}/application_types | List ApplicationTypes for Source
*DefaultApi* | [**ListSourceApplications**](docs/DefaultApi.md#listsourceapplications) | **Get** /sources/{id}/applications | List Applications for Source
*DefaultApi* | [**ListSourceEndpoints**](docs/DefaultApi.md#listsourceendpoints) | **Get** /sources/{id}/endpoints | List Endpoints for Source
*DefaultApi* | [**ListSourceTypeSources**](docs/DefaultApi.md#listsourcetypesources) | **Get** /source_types/{id}/sources | List Sources for SourceType
*DefaultApi* | [**ListSourceTypes**](docs/DefaultApi.md#listsourcetypes) | **Get** /source_types | List SourceTypes
*DefaultApi* | [**ListSources**](docs/DefaultApi.md#listsources) | **Get** /sources | List Sources
*DefaultApi* | [**PostGraphQL**](docs/DefaultApi.md#postgraphql) | **Post** /graphql | Perform a GraphQL Query
*DefaultApi* | [**ShowApplication**](docs/DefaultApi.md#showapplication) | **Get** /applications/{id} | Show an existing Application
*DefaultApi* | [**ShowApplicationAuthentication**](docs/DefaultApi.md#showapplicationauthentication) | **Get** /application_authentications/{id} | Show an existing ApplicationAuthentication
*DefaultApi* | [**ShowApplicationType**](docs/DefaultApi.md#showapplicationtype) | **Get** /application_types/{id} | Show an existing ApplicationType
*DefaultApi* | [**ShowAuthentication**](docs/DefaultApi.md#showauthentication) | **Get** /authentications/{id} | Show an existing Authentication
*DefaultApi* | [**ShowEndpoint**](docs/DefaultApi.md#showendpoint) | **Get** /endpoints/{id} | Show an existing Endpoint
*DefaultApi* | [**ShowSource**](docs/DefaultApi.md#showsource) | **Get** /sources/{id} | Show an existing Source
*DefaultApi* | [**ShowSourceType**](docs/DefaultApi.md#showsourcetype) | **Get** /source_types/{id} | Show an existing SourceType
*DefaultApi* | [**UpdateApplication**](docs/DefaultApi.md#updateapplication) | **Patch** /applications/{id} | Update an existing Application
*DefaultApi* | [**UpdateApplicationAuthentication**](docs/DefaultApi.md#updateapplicationauthentication) | **Patch** /application_authentications/{id} | Update an existing ApplicationAuthentication
*DefaultApi* | [**UpdateAuthentication**](docs/DefaultApi.md#updateauthentication) | **Patch** /authentications/{id} | Update an existing Authentication
*DefaultApi* | [**UpdateEndpoint**](docs/DefaultApi.md#updateendpoint) | **Patch** /endpoints/{id} | Update an existing Endpoint
*DefaultApi* | [**UpdateSource**](docs/DefaultApi.md#updatesource) | **Patch** /sources/{id} | Update an existing Source


## Documentation For Models

 - [Application](docs/Application.md)
 - [ApplicationAuthentication](docs/ApplicationAuthentication.md)
 - [ApplicationAuthenticationsCollection](docs/ApplicationAuthenticationsCollection.md)
 - [ApplicationType](docs/ApplicationType.md)
 - [ApplicationTypesCollection](docs/ApplicationTypesCollection.md)
 - [ApplicationsCollection](docs/ApplicationsCollection.md)
 - [Authentication](docs/Authentication.md)
 - [AuthenticationExtra](docs/AuthenticationExtra.md)
 - [AuthenticationExtraAzure](docs/AuthenticationExtraAzure.md)
 - [AuthenticationsCollection](docs/AuthenticationsCollection.md)
 - [CollectionLinks](docs/CollectionLinks.md)
 - [CollectionMetadata](docs/CollectionMetadata.md)
 - [Endpoint](docs/Endpoint.md)
 - [EndpointsCollection](docs/EndpointsCollection.md)
 - [ErrorNotFound](docs/ErrorNotFound.md)
 - [ErrorNotFoundErrors](docs/ErrorNotFoundErrors.md)
 - [GraphQLRequest](docs/GraphQLRequest.md)
 - [GraphQLResponse](docs/GraphQLResponse.md)
 - [Source](docs/Source.md)
 - [SourceType](docs/SourceType.md)
 - [SourceTypesCollection](docs/SourceTypesCollection.md)
 - [SourcesCollection](docs/SourcesCollection.md)


## Documentation For Authorization



### UserSecurity

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@redhat.com

