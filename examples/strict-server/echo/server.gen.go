// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version (devel) DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
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

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /json)
	JSONExample(ctx echo.Context) error

	// (POST /multipart)
	MultipartExample(ctx echo.Context) error

	// (POST /multiple)
	MultipleRequestAndResponseTypes(ctx echo.Context) error

	// (POST /text)
	TextExample(ctx echo.Context) error

	// (POST /unknown)
	UnknownExample(ctx echo.Context) error

	// (POST /urlencoded)
	URLEncodedExample(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// JSONExample converts echo context to params.
func (w *ServerInterfaceWrapper) JSONExample(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.JSONExample(ctx)
	return err
}

// MultipartExample converts echo context to params.
func (w *ServerInterfaceWrapper) MultipartExample(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.MultipartExample(ctx)
	return err
}

// MultipleRequestAndResponseTypes converts echo context to params.
func (w *ServerInterfaceWrapper) MultipleRequestAndResponseTypes(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.MultipleRequestAndResponseTypes(ctx)
	return err
}

// TextExample converts echo context to params.
func (w *ServerInterfaceWrapper) TextExample(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TextExample(ctx)
	return err
}

// UnknownExample converts echo context to params.
func (w *ServerInterfaceWrapper) UnknownExample(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UnknownExample(ctx)
	return err
}

// URLEncodedExample converts echo context to params.
func (w *ServerInterfaceWrapper) URLEncodedExample(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.URLEncodedExample(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/json", wrapper.JSONExample)
	router.POST(baseURL+"/multipart", wrapper.MultipartExample)
	router.POST(baseURL+"/multiple", wrapper.MultipleRequestAndResponseTypes)
	router.POST(baseURL+"/text", wrapper.TextExample)
	router.POST(baseURL+"/unknown", wrapper.UnknownExample)
	router.POST(baseURL+"/urlencoded", wrapper.URLEncodedExample)

}

type JSONExampleRequestObject struct {
	Body *JSONExampleJSONRequestBody
}

type JSONExample200JSONResponse Example

func (t JSONExample200JSONResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal((Example)(t))
}

type JSONExample400TextResponse Badrequest

func (t JSONExample400TextResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal((Badrequest)(t))
}

type JSONExampledefaultResponse struct {
	StatusCode int
}

type MultipartExampleRequestObject struct {
	Body *multipart.Reader
}

type MultipartExample200MultipartformDataResponse struct {
	Body          io.Reader
	ContentLength int64
}

type MultipartExample400TextResponse Badrequest

func (t MultipartExample400TextResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal((Badrequest)(t))
}

type MultipartExampledefaultResponse struct {
	StatusCode int
}

type MultipleRequestAndResponseTypesRequestObject struct {
	JSONBody      *MultipleRequestAndResponseTypesJSONRequestBody
	FormdataBody  *MultipleRequestAndResponseTypesFormdataRequestBody
	Body          io.Reader
	MultipartBody *multipart.Reader
	TextBody      *MultipleRequestAndResponseTypesTextRequestBody
}

type MultipleRequestAndResponseTypes200JSONResponse Example

func (t MultipleRequestAndResponseTypes200JSONResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal((Example)(t))
}

type MultipleRequestAndResponseTypes200FormdataResponse Example

func (t MultipleRequestAndResponseTypes200FormdataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal((Example)(t))
}

type MultipleRequestAndResponseTypes200ImagepngResponse struct {
	Body          io.Reader
	ContentLength int64
}

type MultipleRequestAndResponseTypes200MultipartformDataResponse struct {
	Body          io.Reader
	ContentLength int64
}

type MultipleRequestAndResponseTypes200TextResponse string

func (t MultipleRequestAndResponseTypes200TextResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal((string)(t))
}

type TextExampleRequestObject struct {
	Body *TextExampleTextRequestBody
}

type TextExample200TextResponse string

func (t TextExample200TextResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal((string)(t))
}

type TextExample400TextResponse Badrequest

func (t TextExample400TextResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal((Badrequest)(t))
}

type TextExampledefaultResponse struct {
	StatusCode int
}

type UnknownExampleRequestObject struct {
	Body io.Reader
}

type UnknownExample200Videomp4Response struct {
	Body          io.Reader
	ContentLength int64
}

type UnknownExample400TextResponse Badrequest

func (t UnknownExample400TextResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal((Badrequest)(t))
}

type UnknownExampledefaultResponse struct {
	StatusCode int
}

type URLEncodedExampleRequestObject struct {
	Body *URLEncodedExampleFormdataRequestBody
}

type URLEncodedExample200FormdataResponse Example

func (t URLEncodedExample200FormdataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal((Example)(t))
}

type URLEncodedExample400TextResponse Badrequest

func (t URLEncodedExample400TextResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal((Badrequest)(t))
}

type URLEncodedExampledefaultResponse struct {
	StatusCode int
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (POST /json)
	JSONExample(ctx context.Context, request JSONExampleRequestObject) interface{}

	// (POST /multipart)
	MultipartExample(ctx context.Context, request MultipartExampleRequestObject) interface{}

	// (POST /multiple)
	MultipleRequestAndResponseTypes(ctx context.Context, request MultipleRequestAndResponseTypesRequestObject) interface{}

	// (POST /text)
	TextExample(ctx context.Context, request TextExampleRequestObject) interface{}

	// (POST /unknown)
	UnknownExample(ctx context.Context, request UnknownExampleRequestObject) interface{}

	// (POST /urlencoded)
	URLEncodedExample(ctx context.Context, request URLEncodedExampleRequestObject) interface{}
}

type StrictHandlerFunc func(ctx echo.Context, args interface{}) interface{}

type StrictMiddlewareFunc func(f StrictHandlerFunc, operationID string) StrictHandlerFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// JSONExample operation middleware
func (sh *strictHandler) JSONExample(ctx echo.Context) error {
	var request JSONExampleRequestObject

	var body JSONExampleJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) interface{} {
		return sh.ssi.JSONExample(ctx.Request().Context(), request.(JSONExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "JSONExample")
	}

	response := handler(ctx, request)

	switch v := response.(type) {
	case JSONExample200JSONResponse:
		return ctx.JSON(200, v)
	case JSONExample400TextResponse:
		return ctx.Blob(400, "text/plain", []byte(v))
	case JSONExampledefaultResponse:
		return ctx.NoContent(v.StatusCode)
	case error:
		return v
	case nil:
	default:
		return fmt.Errorf("Unexpected response type: %T", v)
	}
	return nil
}

// MultipartExample operation middleware
func (sh *strictHandler) MultipartExample(ctx echo.Context) error {
	var request MultipartExampleRequestObject

	if reader, err := ctx.Request().MultipartReader(); err != nil {
		return err
	} else {
		request.Body = reader
	}

	handler := func(ctx echo.Context, request interface{}) interface{} {
		return sh.ssi.MultipartExample(ctx.Request().Context(), request.(MultipartExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "MultipartExample")
	}

	response := handler(ctx, request)

	switch v := response.(type) {
	case MultipartExample200MultipartformDataResponse:
		if v.ContentLength != 0 {
			ctx.Response().Header().Set("Content-Length", fmt.Sprint(v.ContentLength))
		}
		if closer, ok := v.Body.(io.ReadCloser); ok {
			defer closer.Close()
		}
		return ctx.Stream(200, "multipart/form-data", v.Body)
	case MultipartExample400TextResponse:
		return ctx.Blob(400, "text/plain", []byte(v))
	case MultipartExampledefaultResponse:
		return ctx.NoContent(v.StatusCode)
	case error:
		return v
	case nil:
	default:
		return fmt.Errorf("Unexpected response type: %T", v)
	}
	return nil
}

// MultipleRequestAndResponseTypes operation middleware
func (sh *strictHandler) MultipleRequestAndResponseTypes(ctx echo.Context) error {
	var request MultipleRequestAndResponseTypesRequestObject

	if strings.HasPrefix(ctx.Request().Header.Get("Content-Type"), "application/json") {
		var body MultipleRequestAndResponseTypesJSONRequestBody
		if err := ctx.Bind(&body); err != nil {
			return err
		}
		request.JSONBody = &body
	}
	if strings.HasPrefix(ctx.Request().Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
		if form, err := ctx.FormParams(); err == nil {
			var body MultipleRequestAndResponseTypesFormdataRequestBody
			if err := runtime.BindForm(&body, form, nil, nil); err != nil {
				return err
			}
			request.FormdataBody = &body
		} else {
			return err
		}
	}
	if strings.HasPrefix(ctx.Request().Header.Get("Content-Type"), "image/png") {
		request.Body = ctx.Request().Body
	}
	if strings.HasPrefix(ctx.Request().Header.Get("Content-Type"), "multipart/form-data") {
		if reader, err := ctx.Request().MultipartReader(); err != nil {
			return err
		} else {
			request.MultipartBody = reader
		}
	}
	if strings.HasPrefix(ctx.Request().Header.Get("Content-Type"), "text/plain") {
		data, err := ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			return err
		}
		body := MultipleRequestAndResponseTypesTextRequestBody(data)
		request.TextBody = &body
	}

	handler := func(ctx echo.Context, request interface{}) interface{} {
		return sh.ssi.MultipleRequestAndResponseTypes(ctx.Request().Context(), request.(MultipleRequestAndResponseTypesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "MultipleRequestAndResponseTypes")
	}

	response := handler(ctx, request)

	switch v := response.(type) {
	case MultipleRequestAndResponseTypes200JSONResponse:
		return ctx.JSON(200, v)
	case MultipleRequestAndResponseTypes200FormdataResponse:
		if form, err := runtime.MarshalForm(v, nil); err != nil {
			return err
		} else {
			return ctx.Blob(200, "application/x-www-form-urlencoded", []byte(form.Encode()))
		}
	case MultipleRequestAndResponseTypes200ImagepngResponse:
		if v.ContentLength != 0 {
			ctx.Response().Header().Set("Content-Length", fmt.Sprint(v.ContentLength))
		}
		if closer, ok := v.Body.(io.ReadCloser); ok {
			defer closer.Close()
		}
		return ctx.Stream(200, "image/png", v.Body)
	case MultipleRequestAndResponseTypes200MultipartformDataResponse:
		if v.ContentLength != 0 {
			ctx.Response().Header().Set("Content-Length", fmt.Sprint(v.ContentLength))
		}
		if closer, ok := v.Body.(io.ReadCloser); ok {
			defer closer.Close()
		}
		return ctx.Stream(200, "multipart/form-data", v.Body)
	case MultipleRequestAndResponseTypes200TextResponse:
		return ctx.Blob(200, "text/plain", []byte(v))
	case error:
		return v
	case nil:
	default:
		return fmt.Errorf("Unexpected response type: %T", v)
	}
	return nil
}

// TextExample operation middleware
func (sh *strictHandler) TextExample(ctx echo.Context) error {
	var request TextExampleRequestObject

	data, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}
	body := TextExampleTextRequestBody(data)
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) interface{} {
		return sh.ssi.TextExample(ctx.Request().Context(), request.(TextExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "TextExample")
	}

	response := handler(ctx, request)

	switch v := response.(type) {
	case TextExample200TextResponse:
		return ctx.Blob(200, "text/plain", []byte(v))
	case TextExample400TextResponse:
		return ctx.Blob(400, "text/plain", []byte(v))
	case TextExampledefaultResponse:
		return ctx.NoContent(v.StatusCode)
	case error:
		return v
	case nil:
	default:
		return fmt.Errorf("Unexpected response type: %T", v)
	}
	return nil
}

// UnknownExample operation middleware
func (sh *strictHandler) UnknownExample(ctx echo.Context) error {
	var request UnknownExampleRequestObject

	request.Body = ctx.Request().Body

	handler := func(ctx echo.Context, request interface{}) interface{} {
		return sh.ssi.UnknownExample(ctx.Request().Context(), request.(UnknownExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UnknownExample")
	}

	response := handler(ctx, request)

	switch v := response.(type) {
	case UnknownExample200Videomp4Response:
		if v.ContentLength != 0 {
			ctx.Response().Header().Set("Content-Length", fmt.Sprint(v.ContentLength))
		}
		if closer, ok := v.Body.(io.ReadCloser); ok {
			defer closer.Close()
		}
		return ctx.Stream(200, "video/mp4", v.Body)
	case UnknownExample400TextResponse:
		return ctx.Blob(400, "text/plain", []byte(v))
	case UnknownExampledefaultResponse:
		return ctx.NoContent(v.StatusCode)
	case error:
		return v
	case nil:
	default:
		return fmt.Errorf("Unexpected response type: %T", v)
	}
	return nil
}

// URLEncodedExample operation middleware
func (sh *strictHandler) URLEncodedExample(ctx echo.Context) error {
	var request URLEncodedExampleRequestObject

	if form, err := ctx.FormParams(); err == nil {
		var body URLEncodedExampleFormdataRequestBody
		if err := runtime.BindForm(&body, form, nil, nil); err != nil {
			return err
		}
		request.Body = &body
	} else {
		return err
	}

	handler := func(ctx echo.Context, request interface{}) interface{} {
		return sh.ssi.URLEncodedExample(ctx.Request().Context(), request.(URLEncodedExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "URLEncodedExample")
	}

	response := handler(ctx, request)

	switch v := response.(type) {
	case URLEncodedExample200FormdataResponse:
		if form, err := runtime.MarshalForm(v, nil); err != nil {
			return err
		} else {
			return ctx.Blob(200, "application/x-www-form-urlencoded", []byte(form.Encode()))
		}
	case URLEncodedExample400TextResponse:
		return ctx.Blob(400, "text/plain", []byte(v))
	case URLEncodedExampledefaultResponse:
		return ctx.NoContent(v.StatusCode)
	case error:
		return v
	case nil:
	default:
		return fmt.Errorf("Unexpected response type: %T", v)
	}
	return nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xWy27bOhD9FWLuXSqW22alXVNk0WeAJF0VXdDi2GJKcVhyZNkI9O8FJdmNXBlwUife",
	"dCc+5szhmYfmHnIqHVm0HCC7B4/BkQ3YLmZSefxZYeC4ysky2vaTccWpM1LbuAp5gaWMX/97nEMG/6W/",
	"QdPuNKQPwJqmSUBhyL12rMlCBhdSXW9Pkx5yhASvHUIGgb22C2gSwJUsncF45jw59Kw78ktpKhwxaZLN",
	"Ds3uMO/ZaDuneHnI6h1ZltoGofR8jh4ti14FETGCCJVz5BmVmK1F9JCzCOiX6CEB1hyJwc3DfdETDpDA",
	"En3oHL2aTCfT+BxyaKXTkMGbdisBJ7loH5TeBWr1dtRpMeT64ebqi9BByIqplKxzacxalNKHQhqDSmjL",
	"FDlWOYcJtK68jMbvVW9+2WuZQK/4Ban1Tuilc0bnrd2W0GEJsIlU0wo+SLTX0+lzuNlNsquPUeLzztkY",
	"xpbUIFsjzFxWZkT0r/aHpdoK9J58/7K0rAxrJz0/DNZQ7c+bK4dIvsVL5+TLMyVZPpPqx/J0UuH7ZjBa",
	"JDcF1UEUVAsmoVAaUWsuxMZwp7q1FVIEbRcGxYZUMhpJg333emvVdf+W24jx7LWUDFBWZ3Vdn7XBq7xB",
	"m5NC9TRYXcoFps4uhuYRWzJkMFtzTNs/u+uRkijZ+5fZdflC7eSf0uOF3dVehNjf725xdVCrO2LI/+pN",
	"L9Csqm5zv2a91SGyPTGDDlBxqRVSWrrzRyKfStRBKe7R9frTZXfnsfPO0Wr+kR3reH5PEZY4zrejb4Ds",
	"2z1U3kAGBbPL0rQbmSehlosF+ommNA6/zffmVwAAAP//pen/+5gMAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
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
	var res = make(map[string]func() ([]byte, error))
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
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
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
