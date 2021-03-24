[![LatestVersion](https://img.shields.io/github/v/tag/ionos-cloud/sdk-go)](https://github.com/ionos-cloud/sdk-go/releases/latest)
[![Gitter](https://img.shields.io/gitter/room/ionos-cloud/sdk-general)](https://gitter.im/ionos-cloud/sdk-general)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=sdk-go&metric=alert_status&token=fb404d0cd27bc4bfd057afc86a3e76f1e4c2e340)](https://sonarcloud.io/dashboard?id=sdk-go)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=sdk-go&metric=bugs&token=fb404d0cd27bc4bfd057afc86a3e76f1e4c2e340)](https://sonarcloud.io/dashboard?id=sdk-go)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=sdk-go&metric=sqale_rating&token=fb404d0cd27bc4bfd057afc86a3e76f1e4c2e340)](https://sonarcloud.io/dashboard?id=sdk-go)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=sdk-go&metric=reliability_rating&token=fb404d0cd27bc4bfd057afc86a3e76f1e4c2e340)](https://sonarcloud.io/dashboard?id=sdk-go)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=sdk-go&metric=security_rating&token=fb404d0cd27bc4bfd057afc86a3e76f1e4c2e340)](https://sonarcloud.io/dashboard?id=sdk-go)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=sdk-go&metric=vulnerabilities&token=fb404d0cd27bc4bfd057afc86a3e76f1e4c2e340)](https://sonarcloud.io/dashboard?id=sdk-go)

# GoLang API Client For IONOS Cloud

An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API. 

The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 5.0
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
import sw "./ionoscloud"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
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

## Our latest, most up to date usage documentation is available [here](https://docs.ionos.com/golang-sdk/)
