// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
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

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/kataras/iris/v12"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns all pets
	// (GET /pets)
	FindPets(ctx iris.Context, params FindPetsParams)
	// Creates a new pet
	// (POST /pets)
	AddPet(ctx iris.Context)
	// Deletes a pet by ID
	// (DELETE /pets/{id})
	DeletePet(ctx iris.Context, id int64)
	// Returns a pet by ID
	// (GET /pets/{id})
	FindPetByID(ctx iris.Context, id int64)
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

type MiddlewareFunc iris.Handler

// FindPets converts iris context to params.
func (w *ServerInterfaceWrapper) FindPets(ctx iris.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params FindPetsParams
	// ------------- Optional query parameter "tags" -------------

	err = runtime.BindQueryParameter("form", true, false, "tags", ctx.Request().URL.Query(), &params.Tags)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.Writef("Invalid format for parameter tags: %s", err)
		return
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.Request().URL.Query(), &params.Limit)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.Writef("Invalid format for parameter limit: %s", err)
		return
	}

	// Invoke the callback with all the unmarshaled arguments
	w.Handler.FindPets(ctx, params)
}

// AddPet converts iris context to params.
func (w *ServerInterfaceWrapper) AddPet(ctx iris.Context) {

	// Invoke the callback with all the unmarshaled arguments
	w.Handler.AddPet(ctx)
}

// DeletePet converts iris context to params.
func (w *ServerInterfaceWrapper) DeletePet(ctx iris.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Params().Get("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.Writef("Invalid format for parameter id: %s", err)
		return
	}

	// Invoke the callback with all the unmarshaled arguments
	w.Handler.DeletePet(ctx, id)
}

// FindPetByID converts iris context to params.
func (w *ServerInterfaceWrapper) FindPetByID(ctx iris.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Params().Get("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.Writef("Invalid format for parameter id: %s", err)
		return
	}

	// Invoke the callback with all the unmarshaled arguments
	w.Handler.FindPetByID(ctx, id)
}

// IrisServerOption is the option for iris server
type IrisServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *iris.Application, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, IrisServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *iris.Application, si ServerInterface, options IrisServerOptions) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.Get(options.BaseURL+"/pets", wrapper.FindPets)
	router.Post(options.BaseURL+"/pets", wrapper.AddPet)
	router.Delete(options.BaseURL+"/pets/:id", wrapper.DeletePet)
	router.Get(options.BaseURL+"/pets/:id", wrapper.FindPetByID)

	router.Build()
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
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
