// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version (devel) DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
)

// Error defines model for Error.
type Error struct {
	// Code Error code
	Code int32 `json:"code"`

	// Message Error message
	Message string `json:"message"`
}

// NewPet defines model for NewPet.
type NewPet struct {
	// Name Name of the pet
	Name string `json:"name"`

	// Tag Type of the pet
	Tag *string `json:"tag,omitempty"`
}

// Pet defines model for Pet.
type Pet struct {
	// Id Unique id of the pet
	Id int64 `json:"id"`

	// Name Name of the pet
	Name string `json:"name"`

	// Tag Type of the pet
	Tag *string `json:"tag,omitempty"`
}

// FindPetsParams defines parameters for FindPets.
type FindPetsParams struct {
	// Tags tags to filter by
	Tags *[]string `form:"tags,omitempty" json:"tags,omitempty"`

	// Limit maximum number of results to return
	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

// AddPetJSONRequestBody defines body for AddPet for application/json ContentType.
type AddPetJSONRequestBody = NewPet

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns all pets
	// (GET /pets)
	FindPets(w http.ResponseWriter, r *http.Request, params FindPetsParams)
	// Creates a new pet
	// (POST /pets)
	AddPet(w http.ResponseWriter, r *http.Request)
	// Deletes a pet by ID
	// (DELETE /pets/{id})
	DeletePet(w http.ResponseWriter, r *http.Request, id int64)
	// Returns a pet by ID
	// (GET /pets/{id})
	FindPetByID(w http.ResponseWriter, r *http.Request, id int64)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// FindPets operation middleware
func (siw *ServerInterfaceWrapper) FindPets(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params FindPetsParams

	// ------------- Optional query parameter "tags" -------------

	err = runtime.BindQueryParameter("form", true, false, "tags", r.URL.Query(), &params.Tags)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tags", Err: err})
		return
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.FindPets(w, r, params)
	}

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler(w, r.WithContext(ctx))
}

// AddPet operation middleware
func (siw *ServerInterfaceWrapper) AddPet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddPet(w, r)
	}

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeletePet operation middleware
func (siw *ServerInterfaceWrapper) DeletePet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameter("simple", false, "id", mux.Vars(r)["id"], &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeletePet(w, r, id)
	}

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler(w, r.WithContext(ctx))
}

// FindPetByID operation middleware
func (siw *ServerInterfaceWrapper) FindPetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameter("simple", false, "id", mux.Vars(r)["id"], &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.FindPetByID(w, r, id)
	}

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshallingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshallingParamError) Error() string {
	return fmt.Sprintf("Error unmarshalling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshallingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{})
}

type GorillaServerOptions struct {
	BaseURL          string
	BaseRouter       *mux.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r *mux.Router) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r *mux.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options GorillaServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = mux.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.HandleFunc(options.BaseURL+"/pets", wrapper.FindPets).Methods("GET")

	r.HandleFunc(options.BaseURL+"/pets", wrapper.AddPet).Methods("POST")

	r.HandleFunc(options.BaseURL+"/pets/{id}", wrapper.DeletePet).Methods("DELETE")

	r.HandleFunc(options.BaseURL+"/pets/{id}", wrapper.FindPetByID).Methods("GET")

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AA==",
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
