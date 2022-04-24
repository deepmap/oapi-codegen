// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version (devel) DO NOT EDIT.
package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// Badrequest defines model for badrequest.
type Badrequest string

// Example defines model for example.
type Example struct {
	Value *string `json:"value,omitempty"`
}

// MultipleRequestAndResponseTypesTextBody defines parameters for MultipleRequestAndResponseTypes.
type MultipleRequestAndResponseTypesTextBody string

// TextExampleTextBody defines parameters for TextExample.
type TextExampleTextBody string

// HeadersExampleParams defines parameters for HeadersExample.
type HeadersExampleParams struct {
	Header1 string `json:"header1"`
	Header2 *int   `json:"header2,omitempty"`
}

// JSONExampleJSONRequestBody defines body for JSONExample for application/json ContentType.
type JSONExampleJSONRequestBody Example

// MultipartExampleMultipartRequestBody defines body for MultipartExample for multipart/form-data ContentType.
type MultipartExampleMultipartRequestBody Example

// MultipleRequestAndResponseTypesJSONRequestBody defines body for MultipleRequestAndResponseTypes for application/json ContentType.
type MultipleRequestAndResponseTypesJSONRequestBody Example

// MultipleRequestAndResponseTypesFormdataRequestBody defines body for MultipleRequestAndResponseTypes for application/x-www-form-urlencoded ContentType.
type MultipleRequestAndResponseTypesFormdataRequestBody Example

// MultipleRequestAndResponseTypesMultipartRequestBody defines body for MultipleRequestAndResponseTypes for multipart/form-data ContentType.
type MultipleRequestAndResponseTypesMultipartRequestBody Example

// MultipleRequestAndResponseTypesTextRequestBody defines body for MultipleRequestAndResponseTypes for text/plain ContentType.
type MultipleRequestAndResponseTypesTextRequestBody MultipleRequestAndResponseTypesTextBody

// TextExampleTextRequestBody defines body for TextExample for text/plain ContentType.
type TextExampleTextRequestBody TextExampleTextBody

// URLEncodedExampleFormdataRequestBody defines body for URLEncodedExample for application/x-www-form-urlencoded ContentType.
type URLEncodedExampleFormdataRequestBody Example

// HeadersExampleJSONRequestBody defines body for HeadersExample for application/json ContentType.
type HeadersExampleJSONRequestBody Example

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// JSONExample request with any body
	JSONExampleWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	JSONExample(ctx context.Context, body JSONExampleJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MultipartExample request with any body
	MultipartExampleWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MultipleRequestAndResponseTypes request with any body
	MultipleRequestAndResponseTypesWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MultipleRequestAndResponseTypes(ctx context.Context, body MultipleRequestAndResponseTypesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	MultipleRequestAndResponseTypesWithFormdataBody(ctx context.Context, body MultipleRequestAndResponseTypesFormdataRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	MultipleRequestAndResponseTypesWithTextBody(ctx context.Context, body MultipleRequestAndResponseTypesTextRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TextExample request with any body
	TextExampleWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	TextExampleWithTextBody(ctx context.Context, body TextExampleTextRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UnknownExample request with any body
	UnknownExampleWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// URLEncodedExample request with any body
	URLEncodedExampleWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	URLEncodedExampleWithFormdataBody(ctx context.Context, body URLEncodedExampleFormdataRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// HeadersExample request with any body
	HeadersExampleWithBody(ctx context.Context, params *HeadersExampleParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	HeadersExample(ctx context.Context, params *HeadersExampleParams, body HeadersExampleJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) JSONExampleWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewJSONExampleRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) JSONExample(ctx context.Context, body JSONExampleJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewJSONExampleRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) MultipartExampleWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewMultipartExampleRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) MultipleRequestAndResponseTypesWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewMultipleRequestAndResponseTypesRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) MultipleRequestAndResponseTypes(ctx context.Context, body MultipleRequestAndResponseTypesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewMultipleRequestAndResponseTypesRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) MultipleRequestAndResponseTypesWithFormdataBody(ctx context.Context, body MultipleRequestAndResponseTypesFormdataRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewMultipleRequestAndResponseTypesRequestWithFormdataBody(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) MultipleRequestAndResponseTypesWithTextBody(ctx context.Context, body MultipleRequestAndResponseTypesTextRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewMultipleRequestAndResponseTypesRequestWithTextBody(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) TextExampleWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewTextExampleRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) TextExampleWithTextBody(ctx context.Context, body TextExampleTextRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewTextExampleRequestWithTextBody(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UnknownExampleWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUnknownExampleRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) URLEncodedExampleWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewURLEncodedExampleRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) URLEncodedExampleWithFormdataBody(ctx context.Context, body URLEncodedExampleFormdataRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewURLEncodedExampleRequestWithFormdataBody(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) HeadersExampleWithBody(ctx context.Context, params *HeadersExampleParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewHeadersExampleRequestWithBody(c.Server, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) HeadersExample(ctx context.Context, params *HeadersExampleParams, body HeadersExampleJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewHeadersExampleRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewJSONExampleRequest calls the generic JSONExample builder with application/json body
func NewJSONExampleRequest(server string, body JSONExampleJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewJSONExampleRequestWithBody(server, "application/json", bodyReader)
}

// NewJSONExampleRequestWithBody generates requests for JSONExample with any type of body
func NewJSONExampleRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/json")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewMultipartExampleRequestWithBody generates requests for MultipartExample with any type of body
func NewMultipartExampleRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/multipart")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewMultipleRequestAndResponseTypesRequest calls the generic MultipleRequestAndResponseTypes builder with application/json body
func NewMultipleRequestAndResponseTypesRequest(server string, body MultipleRequestAndResponseTypesJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewMultipleRequestAndResponseTypesRequestWithBody(server, "application/json", bodyReader)
}

// NewMultipleRequestAndResponseTypesRequestWithFormdataBody calls the generic MultipleRequestAndResponseTypes builder with application/x-www-form-urlencoded body
func NewMultipleRequestAndResponseTypesRequestWithFormdataBody(server string, body MultipleRequestAndResponseTypesFormdataRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	bodyStr, err := runtime.MarshalForm(body, nil)
	if err != nil {
		return nil, err
	}
	bodyReader = strings.NewReader(bodyStr.Encode())
	return NewMultipleRequestAndResponseTypesRequestWithBody(server, "application/x-www-form-urlencoded", bodyReader)
}

// NewMultipleRequestAndResponseTypesRequestWithTextBody calls the generic MultipleRequestAndResponseTypes builder with text/plain body
func NewMultipleRequestAndResponseTypesRequestWithTextBody(server string, body MultipleRequestAndResponseTypesTextRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	bodyReader = strings.NewReader(string(body))
	return NewMultipleRequestAndResponseTypesRequestWithBody(server, "text/plain", bodyReader)
}

// NewMultipleRequestAndResponseTypesRequestWithBody generates requests for MultipleRequestAndResponseTypes with any type of body
func NewMultipleRequestAndResponseTypesRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/multiple")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewTextExampleRequestWithTextBody calls the generic TextExample builder with text/plain body
func NewTextExampleRequestWithTextBody(server string, body TextExampleTextRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	bodyReader = strings.NewReader(string(body))
	return NewTextExampleRequestWithBody(server, "text/plain", bodyReader)
}

// NewTextExampleRequestWithBody generates requests for TextExample with any type of body
func NewTextExampleRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/text")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewUnknownExampleRequestWithBody generates requests for UnknownExample with any type of body
func NewUnknownExampleRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/unknown")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewURLEncodedExampleRequestWithFormdataBody calls the generic URLEncodedExample builder with application/x-www-form-urlencoded body
func NewURLEncodedExampleRequestWithFormdataBody(server string, body URLEncodedExampleFormdataRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	bodyStr, err := runtime.MarshalForm(body, nil)
	if err != nil {
		return nil, err
	}
	bodyReader = strings.NewReader(bodyStr.Encode())
	return NewURLEncodedExampleRequestWithBody(server, "application/x-www-form-urlencoded", bodyReader)
}

// NewURLEncodedExampleRequestWithBody generates requests for URLEncodedExample with any type of body
func NewURLEncodedExampleRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/urlencoded")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewHeadersExampleRequest calls the generic HeadersExample builder with application/json body
func NewHeadersExampleRequest(server string, params *HeadersExampleParams, body HeadersExampleJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewHeadersExampleRequestWithBody(server, params, "application/json", bodyReader)
}

// NewHeadersExampleRequestWithBody generates requests for HeadersExample with any type of body
func NewHeadersExampleRequestWithBody(server string, params *HeadersExampleParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/with-headers")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	var headerParam0 string

	headerParam0, err = runtime.StyleParamWithLocation("simple", false, "header1", runtime.ParamLocationHeader, params.Header1)
	if err != nil {
		return nil, err
	}

	req.Header.Set("header1", headerParam0)

	if params.Header2 != nil {
		var headerParam1 string

		headerParam1, err = runtime.StyleParamWithLocation("simple", false, "header2", runtime.ParamLocationHeader, *params.Header2)
		if err != nil {
			return nil, err
		}

		req.Header.Set("header2", headerParam1)
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// JSONExample request with any body
	JSONExampleWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*JSONExampleResponse, error)

	JSONExampleWithResponse(ctx context.Context, body JSONExampleJSONRequestBody, reqEditors ...RequestEditorFn) (*JSONExampleResponse, error)

	// MultipartExample request with any body
	MultipartExampleWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*MultipartExampleResponse, error)

	// MultipleRequestAndResponseTypes request with any body
	MultipleRequestAndResponseTypesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*MultipleRequestAndResponseTypesResponse, error)

	MultipleRequestAndResponseTypesWithResponse(ctx context.Context, body MultipleRequestAndResponseTypesJSONRequestBody, reqEditors ...RequestEditorFn) (*MultipleRequestAndResponseTypesResponse, error)

	MultipleRequestAndResponseTypesWithFormdataBodyWithResponse(ctx context.Context, body MultipleRequestAndResponseTypesFormdataRequestBody, reqEditors ...RequestEditorFn) (*MultipleRequestAndResponseTypesResponse, error)

	MultipleRequestAndResponseTypesWithTextBodyWithResponse(ctx context.Context, body MultipleRequestAndResponseTypesTextRequestBody, reqEditors ...RequestEditorFn) (*MultipleRequestAndResponseTypesResponse, error)

	// TextExample request with any body
	TextExampleWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*TextExampleResponse, error)

	TextExampleWithTextBodyWithResponse(ctx context.Context, body TextExampleTextRequestBody, reqEditors ...RequestEditorFn) (*TextExampleResponse, error)

	// UnknownExample request with any body
	UnknownExampleWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UnknownExampleResponse, error)

	// URLEncodedExample request with any body
	URLEncodedExampleWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*URLEncodedExampleResponse, error)

	URLEncodedExampleWithFormdataBodyWithResponse(ctx context.Context, body URLEncodedExampleFormdataRequestBody, reqEditors ...RequestEditorFn) (*URLEncodedExampleResponse, error)

	// HeadersExample request with any body
	HeadersExampleWithBodyWithResponse(ctx context.Context, params *HeadersExampleParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*HeadersExampleResponse, error)

	HeadersExampleWithResponse(ctx context.Context, params *HeadersExampleParams, body HeadersExampleJSONRequestBody, reqEditors ...RequestEditorFn) (*HeadersExampleResponse, error)
}

type JSONExampleResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Example
}

// Status returns HTTPResponse.Status
func (r JSONExampleResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r JSONExampleResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type MultipartExampleResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r MultipartExampleResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r MultipartExampleResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type MultipleRequestAndResponseTypesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Example
}

// Status returns HTTPResponse.Status
func (r MultipleRequestAndResponseTypesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r MultipleRequestAndResponseTypesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type TextExampleResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r TextExampleResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r TextExampleResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UnknownExampleResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r UnknownExampleResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UnknownExampleResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type URLEncodedExampleResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r URLEncodedExampleResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r URLEncodedExampleResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type HeadersExampleResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Example
}

// Status returns HTTPResponse.Status
func (r HeadersExampleResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r HeadersExampleResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// JSONExampleWithBodyWithResponse request with arbitrary body returning *JSONExampleResponse
func (c *ClientWithResponses) JSONExampleWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*JSONExampleResponse, error) {
	rsp, err := c.JSONExampleWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseJSONExampleResponse(rsp)
}

func (c *ClientWithResponses) JSONExampleWithResponse(ctx context.Context, body JSONExampleJSONRequestBody, reqEditors ...RequestEditorFn) (*JSONExampleResponse, error) {
	rsp, err := c.JSONExample(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseJSONExampleResponse(rsp)
}

// MultipartExampleWithBodyWithResponse request with arbitrary body returning *MultipartExampleResponse
func (c *ClientWithResponses) MultipartExampleWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*MultipartExampleResponse, error) {
	rsp, err := c.MultipartExampleWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseMultipartExampleResponse(rsp)
}

// MultipleRequestAndResponseTypesWithBodyWithResponse request with arbitrary body returning *MultipleRequestAndResponseTypesResponse
func (c *ClientWithResponses) MultipleRequestAndResponseTypesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*MultipleRequestAndResponseTypesResponse, error) {
	rsp, err := c.MultipleRequestAndResponseTypesWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseMultipleRequestAndResponseTypesResponse(rsp)
}

func (c *ClientWithResponses) MultipleRequestAndResponseTypesWithResponse(ctx context.Context, body MultipleRequestAndResponseTypesJSONRequestBody, reqEditors ...RequestEditorFn) (*MultipleRequestAndResponseTypesResponse, error) {
	rsp, err := c.MultipleRequestAndResponseTypes(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseMultipleRequestAndResponseTypesResponse(rsp)
}

func (c *ClientWithResponses) MultipleRequestAndResponseTypesWithFormdataBodyWithResponse(ctx context.Context, body MultipleRequestAndResponseTypesFormdataRequestBody, reqEditors ...RequestEditorFn) (*MultipleRequestAndResponseTypesResponse, error) {
	rsp, err := c.MultipleRequestAndResponseTypesWithFormdataBody(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseMultipleRequestAndResponseTypesResponse(rsp)
}

func (c *ClientWithResponses) MultipleRequestAndResponseTypesWithTextBodyWithResponse(ctx context.Context, body MultipleRequestAndResponseTypesTextRequestBody, reqEditors ...RequestEditorFn) (*MultipleRequestAndResponseTypesResponse, error) {
	rsp, err := c.MultipleRequestAndResponseTypesWithTextBody(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseMultipleRequestAndResponseTypesResponse(rsp)
}

// TextExampleWithBodyWithResponse request with arbitrary body returning *TextExampleResponse
func (c *ClientWithResponses) TextExampleWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*TextExampleResponse, error) {
	rsp, err := c.TextExampleWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseTextExampleResponse(rsp)
}

func (c *ClientWithResponses) TextExampleWithTextBodyWithResponse(ctx context.Context, body TextExampleTextRequestBody, reqEditors ...RequestEditorFn) (*TextExampleResponse, error) {
	rsp, err := c.TextExampleWithTextBody(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseTextExampleResponse(rsp)
}

// UnknownExampleWithBodyWithResponse request with arbitrary body returning *UnknownExampleResponse
func (c *ClientWithResponses) UnknownExampleWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UnknownExampleResponse, error) {
	rsp, err := c.UnknownExampleWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUnknownExampleResponse(rsp)
}

// URLEncodedExampleWithBodyWithResponse request with arbitrary body returning *URLEncodedExampleResponse
func (c *ClientWithResponses) URLEncodedExampleWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*URLEncodedExampleResponse, error) {
	rsp, err := c.URLEncodedExampleWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseURLEncodedExampleResponse(rsp)
}

func (c *ClientWithResponses) URLEncodedExampleWithFormdataBodyWithResponse(ctx context.Context, body URLEncodedExampleFormdataRequestBody, reqEditors ...RequestEditorFn) (*URLEncodedExampleResponse, error) {
	rsp, err := c.URLEncodedExampleWithFormdataBody(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseURLEncodedExampleResponse(rsp)
}

// HeadersExampleWithBodyWithResponse request with arbitrary body returning *HeadersExampleResponse
func (c *ClientWithResponses) HeadersExampleWithBodyWithResponse(ctx context.Context, params *HeadersExampleParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*HeadersExampleResponse, error) {
	rsp, err := c.HeadersExampleWithBody(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseHeadersExampleResponse(rsp)
}

func (c *ClientWithResponses) HeadersExampleWithResponse(ctx context.Context, params *HeadersExampleParams, body HeadersExampleJSONRequestBody, reqEditors ...RequestEditorFn) (*HeadersExampleResponse, error) {
	rsp, err := c.HeadersExample(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseHeadersExampleResponse(rsp)
}

// ParseJSONExampleResponse parses an HTTP response from a JSONExampleWithResponse call
func ParseJSONExampleResponse(rsp *http.Response) (*JSONExampleResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &JSONExampleResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Example
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseMultipartExampleResponse parses an HTTP response from a MultipartExampleWithResponse call
func ParseMultipartExampleResponse(rsp *http.Response) (*MultipartExampleResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &MultipartExampleResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseMultipleRequestAndResponseTypesResponse parses an HTTP response from a MultipleRequestAndResponseTypesWithResponse call
func ParseMultipleRequestAndResponseTypesResponse(rsp *http.Response) (*MultipleRequestAndResponseTypesResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &MultipleRequestAndResponseTypesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Example
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case rsp.StatusCode == 200:
		// Content-type (text/plain) unsupported

	}

	return response, nil
}

// ParseTextExampleResponse parses an HTTP response from a TextExampleWithResponse call
func ParseTextExampleResponse(rsp *http.Response) (*TextExampleResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &TextExampleResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseUnknownExampleResponse parses an HTTP response from a UnknownExampleWithResponse call
func ParseUnknownExampleResponse(rsp *http.Response) (*UnknownExampleResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UnknownExampleResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseURLEncodedExampleResponse parses an HTTP response from a URLEncodedExampleWithResponse call
func ParseURLEncodedExampleResponse(rsp *http.Response) (*URLEncodedExampleResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &URLEncodedExampleResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseHeadersExampleResponse parses an HTTP response from a HeadersExampleWithResponse call
func ParseHeadersExampleResponse(rsp *http.Response) (*HeadersExampleResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &HeadersExampleResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Example
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
