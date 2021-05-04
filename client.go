/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0-SDK.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/oauth2"
)

var (
	jsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	xmlCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
)

const DepthParam = "depth"
const DefaultDepth = "10"

const (
   RequestStatusQueued  = "QUEUED"
   RequestStatusRunning = "RUNNING"
   RequestStatusFailed  = "FAILED"
   RequestStatusDone    = "DONE"
)

// APIClient manages communication with the CLOUD API API v5.0-SDK.1
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	DefaultApi *DefaultApiService

	BackupUnitApi *BackupUnitApiService

	ContractApi *ContractApiService

	DataCenterApi *DataCenterApiService

	IPBlocksApi *IPBlocksApiService

	ImageApi *ImageApiService

	KubernetesApi *KubernetesApiService

	LabelApi *LabelApiService

	LanApi *LanApiService

	LoadBalancerApi *LoadBalancerApiService

	LocationApi *LocationApiService

	NicApi *NicApiService

	PrivateCrossConnectApi *PrivateCrossConnectApiService

	RequestApi *RequestApiService

	ServerApi *ServerApiService

	SnapshotApi *SnapshotApiService

	UserManagementApi *UserManagementApiService

	VolumeApi *VolumeApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.DefaultApi = (*DefaultApiService)(&c.common)
	c.BackupUnitApi = (*BackupUnitApiService)(&c.common)
	c.ContractApi = (*ContractApiService)(&c.common)
	c.DataCenterApi = (*DataCenterApiService)(&c.common)
	c.IPBlocksApi = (*IPBlocksApiService)(&c.common)
	c.ImageApi = (*ImageApiService)(&c.common)
	c.KubernetesApi = (*KubernetesApiService)(&c.common)
	c.LabelApi = (*LabelApiService)(&c.common)
	c.LanApi = (*LanApiService)(&c.common)
	c.LoadBalancerApi = (*LoadBalancerApiService)(&c.common)
	c.LocationApi = (*LocationApiService)(&c.common)
	c.NicApi = (*NicApiService)(&c.common)
	c.PrivateCrossConnectApi = (*PrivateCrossConnectApiService)(&c.common)
	c.RequestApi = (*RequestApiService)(&c.common)
	c.ServerApi = (*ServerApiService)(&c.common)
	c.SnapshotApi = (*SnapshotApiService)(&c.common)
	c.UserManagementApi = (*UserManagementApiService)(&c.common)
	c.VolumeApi = (*VolumeApiService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insenstive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	} else if t, ok := obj.(time.Time); ok {
		return t.Format(time.RFC3339)
	}

	return fmt.Sprintf("%v", obj)
}

// helper for converting interface{} parameters to json strings
func parameterToJson(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}


// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	retryCount := 0

	var resp *http.Response
	var err error

	for {

		retryCount ++

		/* we need to clone the request with every retry time because Body closes after the request */
		var clonedRequest *http.Request = request.Clone(request.Context())
		if request.Body != nil {
			clonedRequest.Body, err = request.GetBody()
			if err != nil {
				return nil, err
			}
		}

		if c.cfg.Debug {
			dump, err := httputil.DumpRequestOut(clonedRequest, true)
			if err != nil {
				return nil, err
			}
			log.Printf("\ntry no: %d\n", retryCount)
			log.Printf("%s\n", string(dump))
		}

		resp, err = c.cfg.HTTPClient.Do(clonedRequest)
		if err != nil {
			return resp, err
		}

		if c.cfg.Debug {
			dump, err := httputil.DumpResponse(resp, true)
			if err != nil {
				return resp, err
			}
			log.Printf("\n%s\n", string(dump))
		}

		var backoffTime time.Duration

		switch resp.StatusCode {
		case http.StatusServiceUnavailable,
			http.StatusGatewayTimeout,
			http.StatusBadGateway:
			backoffTime = c.GetConfig().WaitTime

		case http.StatusTooManyRequests:
			if retryAfterSeconds := resp.Header.Get("Retry-After"); retryAfterSeconds != "" {
				waitTime, err := time.ParseDuration(retryAfterSeconds + "s")
				if err != nil {
					return resp, err
				}
				backoffTime = waitTime
			} else {
				backoffTime = c.GetConfig().WaitTime
			}
		default:
			return resp, err

		}

		if retryCount >= c.GetConfig().MaxRetries {
			if c.cfg.Debug {
				fmt.Printf("number of maximum retries exceeded (%d retries)\n", c.cfg.MaxRetries)
			}
			break
		} else {
			c.backOff(backoffTime)
		}
	}

	return resp, err
}

func (c *APIClient) backOff(t time.Duration) {
	if t > c.GetConfig().MaxWaitTime {
		t = c.GetConfig().MaxWaitTime
	}
	if c.cfg.Debug {
		fmt.Printf("sleeping %s before retrying request\n", t.String())
	}
	time.Sleep(t)
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *Configuration {
	return c.cfg
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFileName string,
	fileName string,
	fileBytes []byte) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(fileBytes) > 0 && fileName != "") {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		if len(fileBytes) > 0 && fileName != "" {
			w.Boundary()
			//_, fileNm := filepath.Split(fileName)
			part, err := w.CreateFormFile(formFileName, filepath.Base(fileName))
			if err != nil {
				return nil, err
			}
			_, err = part.Write(fileBytes)
			if err != nil {
				return nil, err
			}
		}

		// Set the Boundary in the Content-Type
		headerParams["Content-Type"] = w.FormDataContentType()

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		url.Host = c.cfg.Host
	}

	// Override request scheme, if applicable
	if c.cfg.Scheme != "" {
		url.Scheme = c.cfg.Scheme
	}

	// Adding Query Param
	query := url.Query()
	/* adding default query params */
    for k, v := range c.cfg.DefaultQueryParams {
        if _, ok := queryParams[k]; !ok {
            queryParams[k] = v
        }
    }
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

    // Adding default depth if needed
    if query.Get(DepthParam) == "" {
        query.Add(DepthParam, DefaultDepth)
    }
 

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

    if c.cfg.Token != "" {
        localVarRequest.Header.Add("Authorization", "Bearer " + c.cfg.Token)
    } else {
        if c.cfg.Username != "" {
            localVarRequest.SetBasicAuth(c.cfg.Username, c.cfg.Password)
        }
    }

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// OAuth2 authentication
		if tok, ok := ctx.Value(ContextOAuth2).(oauth2.TokenSource); ok {
			// We were able to grab an oauth2 token from the context
			var latestToken *oauth2.Token
			if latestToken, err = tok.Token(); err != nil {
				return nil, err
			}

			latestToken.SetAuthHeader(localVarRequest)
		}

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(ContextBasicAuth).(BasicAuth); ok {
			localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
		}

		// AccessToken Authentication
		if auth, ok := ctx.Value(ContextAccessToken).(string); ok {
			localVarRequest.Header.Add("Authorization", "Bearer "+auth)
		}

	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}
	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if xmlCheck.MatchString(contentType) {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	if jsonCheck.MatchString(contentType) {
		if actualObj, ok := v.(interface{GetActualInstance() interface{}}); ok { // oneOf, anyOf schemas
			if unmarshalObj, ok := actualObj.(interface{UnmarshalJSON([]byte) error}); ok { // make sure it has UnmarshalJSON defined
				if err = unmarshalObj.UnmarshalJSON(b); err!= nil {
					return err
				}
			} else {
				errors.New("Unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
			}
		} else if err = json.Unmarshal(b, v); err != nil { // simple model
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}


func (c *APIClient) WaitForRequest(ctx context.Context, path string) (*APIResponse, error) {

	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	r, err := c.prepareRequest(ctx, path, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		resp, err := c.callAPI(r)

		var localVarBody = make([]byte, 0)
		if resp != nil {
			var errRead error
			localVarBody, errRead = ioutil.ReadAll(resp.Body)
			_ = resp.Body.Close()
			if errRead != nil {
				return nil, errRead
			}

		}

		localVarAPIResponse := &APIResponse {
			Response: resp,
			Method: localVarHTTPMethod,
			RequestURL: path,
			Operation: "WaitForRequest",
		}

		localVarAPIResponse.Payload = localVarBody

		if err != nil {
			return localVarAPIResponse, err
		}

		status := RequestStatus{}
		err = c.decode(&status, localVarBody, resp.Header.Get("Content-Type"))
		if err != nil {
			return localVarAPIResponse, err
		}

		if resp.StatusCode != http.StatusOK {
			return localVarAPIResponse, fmt.Errorf("WaitForRequest failed; received status code %d from API", resp.StatusCode)
		}
		if status.Metadata != nil && status.Metadata.Status != nil {
			switch *status.Metadata.Status {
			case RequestStatusDone:
				return localVarAPIResponse, nil
			case RequestStatusFailed:
				var id = "<none>"
				var message = "<none>"
				if status.Id != nil {
					id = *status.Id
				}
				if status.Metadata.Message != nil {
					message = *status.Metadata.Message
				}
				return localVarAPIResponse, errors.New(
					fmt.Sprintf("Request %s failed: %s", id, message),
				)
			}
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			continue
		}
	}
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		err = xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		} else {
			expires = now.Add(lifetime)
		}
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() interface{} {
	return e.model
}
