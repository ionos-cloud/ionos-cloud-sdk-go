# \RequestsApi

All URIs are relative to *https://api.ionos.com/cloudapi/v6*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**RequestsFindById**](RequestsApi.md#RequestsFindById) | **Get** /requests/{requestId} | Retrieve a Request|
|[**RequestsGet**](RequestsApi.md#RequestsGet) | **Get** /requests | List Requests|
|[**RequestsStatusGet**](RequestsApi.md#RequestsStatusGet) | **Get** /requests/{requestId}/status | Retrieve Request Status|



## RequestsFindById

```go
var result Request = RequestsFindById(ctx, requestId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve a Request



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
    requestId := "requestId_example" // string | 
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RequestsApi.RequestsFindById(context.Background(), requestId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestsApi.RequestsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestsFindById`: Request
    fmt.Fprintf(os.Stdout, "Response from `RequestsApi.RequestsFindById`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**requestId** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**Request**](Request.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## RequestsGet

```go
var result Requests = RequestsGet(ctx)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .FilterStatus(filterStatus)
                      .FilterCreatedAfter(filterCreatedAfter)
                      .FilterCreatedBefore(filterCreatedBefore)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

List Requests



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
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)
    filterStatus := "filterStatus_example" // string | Request status filter to fetch all the request based on a particular status [QUEUED, RUNNING, DONE, FAILED] (optional)
    filterCreatedAfter := "filterCreatedAfter_example" // string | Filter all the requests after the created date (optional)
    filterCreatedBefore := "filterCreatedBefore_example" // string | Filter all the requests before the created date (optional)
    offset := int32(56) // int32 | the first element (of the total list of elements) to include in the response (use together with limit for pagination) (optional) (default to 0)
    limit := int32(56) // int32 | the maximum number of elements to return (use together with offset for pagination) (optional) (default to 1000)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RequestsApi.RequestsGet(context.Background()).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).FilterStatus(filterStatus).FilterCreatedAfter(filterCreatedAfter).FilterCreatedBefore(filterCreatedBefore).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestsApi.RequestsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestsGet`: Requests
    fmt.Fprintf(os.Stdout, "Response from `RequestsApi.RequestsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |
| **filterStatus** | **string** | Request status filter to fetch all the request based on a particular status [QUEUED, RUNNING, DONE, FAILED] | |
| **filterCreatedAfter** | **string** | Filter all the requests after the created date | |
| **filterCreatedBefore** | **string** | Filter all the requests before the created date | |
| **offset** | **int32** | the first element (of the total list of elements) to include in the response (use together with limit for pagination) | [default to 0]|
| **limit** | **int32** | the maximum number of elements to return (use together with offset for pagination) | [default to 1000]|

### Return type

[**Requests**](Requests.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## RequestsStatusGet

```go
var result RequestStatus = RequestsStatusGet(ctx, requestId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve Request Status



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
    requestId := "requestId_example" // string | 
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RequestsApi.RequestsStatusGet(context.Background(), requestId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestsApi.RequestsStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestsStatusGet`: RequestStatus
    fmt.Fprintf(os.Stdout, "Response from `RequestsApi.RequestsStatusGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**requestId** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsStatusGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**RequestStatus**](RequestStatus.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

