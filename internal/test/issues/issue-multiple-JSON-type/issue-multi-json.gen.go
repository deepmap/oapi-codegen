// Package multijson provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version (devel) DO NOT EDIT.
package multijson

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

// Bar defines model for bar.
type Bar struct {
	Field2 *string `json:"field2,omitempty"`
}

// Foo defines model for foo.
type Foo struct {
	Field1 *string `json:"field1,omitempty"`
}

// BazApplicationBarPlusJSON defines model for baz.
type BazApplicationBarPlusJSON = Bar

// BazApplicationFooPlusJSON defines model for baz.
type BazApplicationFooPlusJSON = Foo

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
	// Test request
	Test(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) Test(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewTestRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewTestRequest generates requests for Test
func NewTestRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/test")
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
	// TestWithResponse request
	TestWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*TestResponse, error)
}

type TestResponse struct {
	Body                  []byte
	HTTPResponse          *http.Response
	ApplicationbarJSON200 *Bar
	ApplicationfooJSON200 *Foo
	ApplicationbarJSON201 *BazApplicationBarPlusJSON
	ApplicationfooJSON201 *BazApplicationFooPlusJSON
}

// Status returns HTTPResponse.Status
func (r TestResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r TestResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// TestWithResponse request returning *TestResponse
func (c *ClientWithResponses) TestWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*TestResponse, error) {
	rsp, err := c.Test(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseTestResponse(rsp)
}

// ParseTestResponse parses an HTTP response from a TestWithResponse call
func ParseTestResponse(rsp *http.Response) (*TestResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &TestResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case rsp.Header.Get("Content-Type") == "application/bar+json" && rsp.StatusCode == 200:
		var dest Bar
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationbarJSON200 = &dest

	case rsp.Header.Get("Content-Type") == "application/bar+json" && rsp.StatusCode == 201:
		var dest BazApplicationBarPlusJSON
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationbarJSON201 = &dest

	case rsp.Header.Get("Content-Type") == "application/foo+json" && rsp.StatusCode == 200:
		var dest Foo
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationfooJSON200 = &dest

	case rsp.Header.Get("Content-Type") == "application/foo+json" && rsp.StatusCode == 201:
		var dest BazApplicationFooPlusJSON
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationfooJSON201 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /test)
	Test(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// Test operation middleware
func (siw *ServerInterfaceWrapper) Test(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.Test(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/test", wrapper.Test)
}

type BazApplicationBarPlusJSONResponse Bar
type BazApplicationFooPlusJSONResponse Foo

type TestRequestObject struct {
}

type TestResponseObject interface {
	VisitTestResponse(w http.ResponseWriter) error
}

type Test200ApplicationBarPlusJSONResponse Bar

func (response Test200ApplicationBarPlusJSONResponse) VisitTestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/bar+json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type Test200ApplicationFooPlusJSONResponse Foo

func (response Test200ApplicationFooPlusJSONResponse) VisitTestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/foo+json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type Test201ApplicationBarPlusJSONResponse struct {
	BazApplicationBarPlusJSONResponse
}

func (response Test201ApplicationBarPlusJSONResponse) VisitTestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/bar+json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type Test201ApplicationFooPlusJSONResponse struct {
	BazApplicationFooPlusJSONResponse
}

func (response Test201ApplicationFooPlusJSONResponse) VisitTestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/foo+json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /test)
	Test(ctx context.Context, request TestRequestObject) (TestResponseObject, error)
}

type StrictHandlerFunc = runtime.StrictGinHandlerFunc
type StrictMiddlewareFunc = runtime.StrictGinMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// Test operation middleware
func (sh *strictHandler) Test(ctx *gin.Context) {
	var request TestRequestObject

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.Test(ctx, request.(TestRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Test")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(TestResponseObject); ok {
		if err := validResponse.VisitTestResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("Unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8ySQU4DMQxF7/JhR9Skwy43YM8F0qnTBk1tKwkLqObuyGkFqlQW7JiNPXHeT/5Xzpjl",
	"pMLEvSGeUampcKPxs0ufVmbhTtytTapLmVMvwn6X6tNbE7b1Nh/plKx7rJQR8eB/dP1l2ozAuroblSzy",
	"R5UsgtU+dwWud61WtIpS7eViIBda9pN1/UMJEa3XwgcM2HTuE9u7hDGFsyDy+7I4iBInLYh43oRNgIOm",
	"fhwqvlMbeR1oFDth2H3ZI+LVhu426imEfxy1wxS2v23+9uHtvYyg1q8AAAD//25zbQ5XAgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
