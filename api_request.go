/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionossdk

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// RequestApiService RequestApi service
type RequestApiService service

type ApiRequestsFindByIdRequest struct {
	ctx _context.Context
	ApiService *RequestApiService
	requestId string
	pretty *bool
	depth *int32
	xContractNumber *int32
}

func (r ApiRequestsFindByIdRequest) Pretty(pretty bool) ApiRequestsFindByIdRequest {
	r.pretty = &pretty
	return r
}
func (r ApiRequestsFindByIdRequest) Depth(depth int32) ApiRequestsFindByIdRequest {
	r.depth = &depth
	return r
}
func (r ApiRequestsFindByIdRequest) XContractNumber(xContractNumber int32) ApiRequestsFindByIdRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiRequestsFindByIdRequest) Execute() (Request, *APIResponse, error) {
	return r.ApiService.RequestsFindByIdExecute(r)
}

/*
 * RequestsFindById Retrieve a Request
 * Retrieves the attributes of a given request.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param requestId
 * @return ApiRequestsFindByIdRequest
 */
func (a *RequestApiService) RequestsFindById(ctx _context.Context, requestId string) ApiRequestsFindByIdRequest {
	return ApiRequestsFindByIdRequest{
		ApiService: a,
		ctx: ctx,
		requestId: requestId,
	}
}

/*
 * Execute executes the request
 * @return Request
 */
func (a *RequestApiService) RequestsFindByIdExecute(r ApiRequestsFindByIdRequest) (Request, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Request
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestApiService.RequestsFindById")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/requests/{requestId}"
	localVarPath = strings.Replace(localVarPath, "{"+"requestId"+"}", _neturl.PathEscape(parameterToString(r.requestId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse {
		Response: localVarHTTPResponse,
		Method: localVarHTTPMethod,
		RequestURL: localVarPath,
		Operation: "RequestsFindById",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiRequestsGetRequest struct {
	ctx _context.Context
	ApiService *RequestApiService
	pretty *bool
	depth *int32
	xContractNumber *int32
	filterStatus *string
	filterCreatedAfter *string
	filterCreatedBefore *string
	filterUrl *string
	filterCreatedDate *string
	filterMethod *string
	filterBody *string
}

func (r ApiRequestsGetRequest) Pretty(pretty bool) ApiRequestsGetRequest {
	r.pretty = &pretty
	return r
}
func (r ApiRequestsGetRequest) Depth(depth int32) ApiRequestsGetRequest {
	r.depth = &depth
	return r
}
func (r ApiRequestsGetRequest) XContractNumber(xContractNumber int32) ApiRequestsGetRequest {
	r.xContractNumber = &xContractNumber
	return r
}
func (r ApiRequestsGetRequest) FilterStatus(filterStatus string) ApiRequestsGetRequest {
	r.filterStatus = &filterStatus
	return r
}
func (r ApiRequestsGetRequest) FilterCreatedAfter(filterCreatedAfter string) ApiRequestsGetRequest {
	r.filterCreatedAfter = &filterCreatedAfter
	return r
}
func (r ApiRequestsGetRequest) FilterCreatedBefore(filterCreatedBefore string) ApiRequestsGetRequest {
	r.filterCreatedBefore = &filterCreatedBefore
	return r
}
func (r ApiRequestsGetRequest) FilterUrl(filterUrl string) ApiRequestsGetRequest {
	r.filterUrl = &filterUrl
	return r
}
func (r ApiRequestsGetRequest) FilterCreatedDate(filterCreatedDate string) ApiRequestsGetRequest {
	r.filterCreatedDate = &filterCreatedDate
	return r
}
func (r ApiRequestsGetRequest) FilterMethod(filterMethod string) ApiRequestsGetRequest {
	r.filterMethod = &filterMethod
	return r
}
func (r ApiRequestsGetRequest) FilterBody(filterBody string) ApiRequestsGetRequest {
	r.filterBody = &filterBody
	return r
}

func (r ApiRequestsGetRequest) Execute() (Requests, *APIResponse, error) {
	return r.ApiService.RequestsGetExecute(r)
}

/*
 * RequestsGet List Requests
 * Retrieve a list of API requests.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiRequestsGetRequest
 */
func (a *RequestApiService) RequestsGet(ctx _context.Context) ApiRequestsGetRequest {
	return ApiRequestsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return Requests
 */
func (a *RequestApiService) RequestsGetExecute(r ApiRequestsGetRequest) (Requests, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Requests
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestApiService.RequestsGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/requests"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	}
	if r.filterStatus != nil {
		localVarQueryParams.Add("filter.status", parameterToString(*r.filterStatus, ""))
	}
	if r.filterCreatedAfter != nil {
		localVarQueryParams.Add("filter.createdAfter", parameterToString(*r.filterCreatedAfter, ""))
	}
	if r.filterCreatedBefore != nil {
		localVarQueryParams.Add("filter.createdBefore", parameterToString(*r.filterCreatedBefore, ""))
	}
	if r.filterUrl != nil {
		localVarQueryParams.Add("filter.url", parameterToString(*r.filterUrl, ""))
	}
	if r.filterCreatedDate != nil {
		localVarQueryParams.Add("filter.createdDate", parameterToString(*r.filterCreatedDate, ""))
	}
	if r.filterMethod != nil {
		localVarQueryParams.Add("filter.method", parameterToString(*r.filterMethod, ""))
	}
	if r.filterBody != nil {
		localVarQueryParams.Add("filter.body", parameterToString(*r.filterBody, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse {
		Response: localVarHTTPResponse,
		Method: localVarHTTPMethod,
		RequestURL: localVarPath,
		Operation: "RequestsGet",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiRequestsStatusGetRequest struct {
	ctx _context.Context
	ApiService *RequestApiService
	requestId string
	pretty *bool
	depth *int32
	xContractNumber *int32
}

func (r ApiRequestsStatusGetRequest) Pretty(pretty bool) ApiRequestsStatusGetRequest {
	r.pretty = &pretty
	return r
}
func (r ApiRequestsStatusGetRequest) Depth(depth int32) ApiRequestsStatusGetRequest {
	r.depth = &depth
	return r
}
func (r ApiRequestsStatusGetRequest) XContractNumber(xContractNumber int32) ApiRequestsStatusGetRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiRequestsStatusGetRequest) Execute() (RequestStatus, *APIResponse, error) {
	return r.ApiService.RequestsStatusGetExecute(r)
}

/*
 * RequestsStatusGet Retrieve Request Status
 * Retrieves the status of a given request.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param requestId
 * @return ApiRequestsStatusGetRequest
 */
func (a *RequestApiService) RequestsStatusGet(ctx _context.Context, requestId string) ApiRequestsStatusGetRequest {
	return ApiRequestsStatusGetRequest{
		ApiService: a,
		ctx: ctx,
		requestId: requestId,
	}
}

/*
 * Execute executes the request
 * @return RequestStatus
 */
func (a *RequestApiService) RequestsStatusGetExecute(r ApiRequestsStatusGetRequest) (RequestStatus, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RequestStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestApiService.RequestsStatusGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/requests/{requestId}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"requestId"+"}", _neturl.PathEscape(parameterToString(r.requestId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse {
		Response: localVarHTTPResponse,
		Method: localVarHTTPMethod,
		RequestURL: localVarPath,
		Operation: "RequestsStatusGet",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}
