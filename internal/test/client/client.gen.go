// Package client provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/ula/oapi-codegen version v0.0.0-00010101000000-000000000000 DO NOT EDIT.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	OpenIdScopes = "OpenId.Scopes"
)

// SchemaObject defines model for SchemaObject.
type SchemaObject struct {
	FirstName string `json:"firstName"`
	Role      string `json:"role"`
}

// PostVendorJsonApplicationVndAPIPlusJSONBody defines parameters for PostVendorJson.
type PostVendorJsonApplicationVndAPIPlusJSONBody = map[string]interface{}

// PostBothJSONRequestBody defines body for PostBoth for application/json ContentType.
type PostBothJSONRequestBody = SchemaObject

// PostJsonJSONRequestBody defines body for PostJson for application/json ContentType.
type PostJsonJSONRequestBody = SchemaObject

// PostVendorJsonApplicationVndAPIPlusJSONRequestBody defines body for PostVendorJson for application/vnd.api+json ContentType.
type PostVendorJsonApplicationVndAPIPlusJSONRequestBody = PostVendorJsonApplicationVndAPIPlusJSONBody

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
	// PostBothWithBody request with any body
	PostBothWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostBoth(ctx context.Context, body PostBothJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetBoth request
	GetBoth(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostJsonWithBody request with any body
	PostJsonWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostJson(ctx context.Context, body PostJsonJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetJson request
	GetJson(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostOtherWithBody request with any body
	PostOtherWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetOther request
	GetOther(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetJsonWithTrailingSlash request
	GetJsonWithTrailingSlash(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostVendorJsonWithBody request with any body
	PostVendorJsonWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostVendorJsonWithApplicationVndAPIPlusJSONBody(ctx context.Context, body PostVendorJsonApplicationVndAPIPlusJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) PostBothWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostBothRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostBoth(ctx context.Context, body PostBothJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostBothRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetBoth(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetBothRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostJsonWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostJsonRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostJson(ctx context.Context, body PostJsonJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostJsonRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetJson(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetJsonRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostOtherWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostOtherRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetOther(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetOtherRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetJsonWithTrailingSlash(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetJsonWithTrailingSlashRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostVendorJsonWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostVendorJsonRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostVendorJsonWithApplicationVndAPIPlusJSONBody(ctx context.Context, body PostVendorJsonApplicationVndAPIPlusJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostVendorJsonRequestWithApplicationVndAPIPlusJSONBody(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewPostBothRequest calls the generic PostBoth builder with application/json body
func NewPostBothRequest(server string, body PostBothJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostBothRequestWithBody(server, "application/json", bodyReader)
}

// NewPostBothRequestWithBody generates requests for PostBoth with any type of body
func NewPostBothRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/with_both_bodies")
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

// NewGetBothRequest generates requests for GetBoth
func NewGetBothRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/with_both_responses")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostJsonRequest calls the generic PostJson builder with application/json body
func NewPostJsonRequest(server string, body PostJsonJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostJsonRequestWithBody(server, "application/json", bodyReader)
}

// NewPostJsonRequestWithBody generates requests for PostJson with any type of body
func NewPostJsonRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/with_json_body")
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

// NewGetJsonRequest generates requests for GetJson
func NewGetJsonRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/with_json_response")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostOtherRequestWithBody generates requests for PostOther with any type of body
func NewPostOtherRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/with_other_body")
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

// NewGetOtherRequest generates requests for GetOther
func NewGetOtherRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/with_other_response")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetJsonWithTrailingSlashRequest generates requests for GetJsonWithTrailingSlash
func NewGetJsonWithTrailingSlashRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/with_trailing_slash/")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostVendorJsonRequestWithApplicationVndAPIPlusJSONBody calls the generic PostVendorJson builder with application/vnd.api+json body
func NewPostVendorJsonRequestWithApplicationVndAPIPlusJSONBody(server string, body PostVendorJsonApplicationVndAPIPlusJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostVendorJsonRequestWithBody(server, "application/vnd.api+json", bodyReader)
}

// NewPostVendorJsonRequestWithBody generates requests for PostVendorJson with any type of body
func NewPostVendorJsonRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/with_vendor_json")
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
	// PostBothWithBodyWithResponse request with any body
	PostBothWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostBothResponse, error)

	PostBothWithResponse(ctx context.Context, body PostBothJSONRequestBody, reqEditors ...RequestEditorFn) (*PostBothResponse, error)

	// GetBothWithResponse request
	GetBothWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetBothResponse, error)

	// PostJsonWithBodyWithResponse request with any body
	PostJsonWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostJsonResponse, error)

	PostJsonWithResponse(ctx context.Context, body PostJsonJSONRequestBody, reqEditors ...RequestEditorFn) (*PostJsonResponse, error)

	// GetJsonWithResponse request
	GetJsonWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetJsonResponse, error)

	// PostOtherWithBodyWithResponse request with any body
	PostOtherWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostOtherResponse, error)

	// GetOtherWithResponse request
	GetOtherWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetOtherResponse, error)

	// GetJsonWithTrailingSlashWithResponse request
	GetJsonWithTrailingSlashWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetJsonWithTrailingSlashResponse, error)

	// PostVendorJsonWithBodyWithResponse request with any body
	PostVendorJsonWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostVendorJsonResponse, error)

	PostVendorJsonWithApplicationVndAPIPlusJSONBodyWithResponse(ctx context.Context, body PostVendorJsonApplicationVndAPIPlusJSONRequestBody, reqEditors ...RequestEditorFn) (*PostVendorJsonResponse, error)
}

type PostBothResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r PostBothResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostBothResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetBothResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetBothResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetBothResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostJsonResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r PostJsonResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostJsonResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetJsonResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetJsonResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetJsonResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostOtherResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r PostOtherResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostOtherResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetOtherResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetOtherResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetOtherResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetJsonWithTrailingSlashResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetJsonWithTrailingSlashResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetJsonWithTrailingSlashResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostVendorJsonResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r PostVendorJsonResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostVendorJsonResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// PostBothWithBodyWithResponse request with arbitrary body returning *PostBothResponse
func (c *ClientWithResponses) PostBothWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostBothResponse, error) {
	rsp, err := c.PostBothWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostBothResponse(rsp)
}

func (c *ClientWithResponses) PostBothWithResponse(ctx context.Context, body PostBothJSONRequestBody, reqEditors ...RequestEditorFn) (*PostBothResponse, error) {
	rsp, err := c.PostBoth(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostBothResponse(rsp)
}

// GetBothWithResponse request returning *GetBothResponse
func (c *ClientWithResponses) GetBothWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetBothResponse, error) {
	rsp, err := c.GetBoth(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetBothResponse(rsp)
}

// PostJsonWithBodyWithResponse request with arbitrary body returning *PostJsonResponse
func (c *ClientWithResponses) PostJsonWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostJsonResponse, error) {
	rsp, err := c.PostJsonWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostJsonResponse(rsp)
}

func (c *ClientWithResponses) PostJsonWithResponse(ctx context.Context, body PostJsonJSONRequestBody, reqEditors ...RequestEditorFn) (*PostJsonResponse, error) {
	rsp, err := c.PostJson(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostJsonResponse(rsp)
}

// GetJsonWithResponse request returning *GetJsonResponse
func (c *ClientWithResponses) GetJsonWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetJsonResponse, error) {
	rsp, err := c.GetJson(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetJsonResponse(rsp)
}

// PostOtherWithBodyWithResponse request with arbitrary body returning *PostOtherResponse
func (c *ClientWithResponses) PostOtherWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostOtherResponse, error) {
	rsp, err := c.PostOtherWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostOtherResponse(rsp)
}

// GetOtherWithResponse request returning *GetOtherResponse
func (c *ClientWithResponses) GetOtherWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetOtherResponse, error) {
	rsp, err := c.GetOther(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetOtherResponse(rsp)
}

// GetJsonWithTrailingSlashWithResponse request returning *GetJsonWithTrailingSlashResponse
func (c *ClientWithResponses) GetJsonWithTrailingSlashWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetJsonWithTrailingSlashResponse, error) {
	rsp, err := c.GetJsonWithTrailingSlash(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetJsonWithTrailingSlashResponse(rsp)
}

// PostVendorJsonWithBodyWithResponse request with arbitrary body returning *PostVendorJsonResponse
func (c *ClientWithResponses) PostVendorJsonWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostVendorJsonResponse, error) {
	rsp, err := c.PostVendorJsonWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostVendorJsonResponse(rsp)
}

func (c *ClientWithResponses) PostVendorJsonWithApplicationVndAPIPlusJSONBodyWithResponse(ctx context.Context, body PostVendorJsonApplicationVndAPIPlusJSONRequestBody, reqEditors ...RequestEditorFn) (*PostVendorJsonResponse, error) {
	rsp, err := c.PostVendorJsonWithApplicationVndAPIPlusJSONBody(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostVendorJsonResponse(rsp)
}

// ParsePostBothResponse parses an HTTP response from a PostBothWithResponse call
func ParsePostBothResponse(rsp *http.Response) (*PostBothResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostBothResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetBothResponse parses an HTTP response from a GetBothWithResponse call
func ParseGetBothResponse(rsp *http.Response) (*GetBothResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetBothResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParsePostJsonResponse parses an HTTP response from a PostJsonWithResponse call
func ParsePostJsonResponse(rsp *http.Response) (*PostJsonResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostJsonResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetJsonResponse parses an HTTP response from a GetJsonWithResponse call
func ParseGetJsonResponse(rsp *http.Response) (*GetJsonResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetJsonResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParsePostOtherResponse parses an HTTP response from a PostOtherWithResponse call
func ParsePostOtherResponse(rsp *http.Response) (*PostOtherResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostOtherResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetOtherResponse parses an HTTP response from a GetOtherWithResponse call
func ParseGetOtherResponse(rsp *http.Response) (*GetOtherResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetOtherResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetJsonWithTrailingSlashResponse parses an HTTP response from a GetJsonWithTrailingSlashWithResponse call
func ParseGetJsonWithTrailingSlashResponse(rsp *http.Response) (*GetJsonWithTrailingSlashResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetJsonWithTrailingSlashResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParsePostVendorJsonResponse parses an HTTP response from a PostVendorJsonWithResponse call
func ParsePostVendorJsonResponse(rsp *http.Response) (*PostVendorJsonResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostVendorJsonResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
