// Package packagea provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package packagea

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"
	"time"

	externalRef0 "github.com/oapi-codegen/oapi-codegen/internal/test/externalref/packageB"
	"github.com/getkin/kin-openapi/openapi3"
)

const (
	Api_keyScopes       = "api_key.Scopes"
	Petstore_authScopes = "petstore_auth.Scopes"
)

// Defines values for OrderStatus.
const (
	Approved  OrderStatus = "approved"
	Delivered OrderStatus = "delivered"
	Placed    OrderStatus = "placed"
)

// Defines values for PetStatus.
const (
	PetStatusAvailable PetStatus = "available"
	PetStatusPending   PetStatus = "pending"
	PetStatusSold      PetStatus = "sold"
)

// Defines values for FindPetsByStatusParamsStatus.
const (
	FindPetsByStatusParamsStatusAvailable FindPetsByStatusParamsStatus = "available"
	FindPetsByStatusParamsStatusPending   FindPetsByStatusParamsStatus = "pending"
	FindPetsByStatusParamsStatusSold      FindPetsByStatusParamsStatus = "sold"
)

// Address defines model for Address.
type Address struct {
	City   *string `json:"city,omitempty"`
	State  *string `json:"state,omitempty"`
	Street *string `json:"street,omitempty"`
	Zip    *string `json:"zip,omitempty"`
}

// ApiResponse defines model for ApiResponse.
type ApiResponse struct {
	Code    *int32  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Type    *string `json:"type,omitempty"`
}

// Category defines model for Category.
type Category struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Customer defines model for Customer.
type Customer struct {
	Address  *[]Address `json:"address,omitempty"`
	Id       *int64     `json:"id,omitempty"`
	Username *string    `json:"username,omitempty"`
}

// Order defines model for Order.
type Order struct {
	Complete *bool      `json:"complete,omitempty"`
	Id       *int64     `json:"id,omitempty"`
	PetId    *int64     `json:"petId,omitempty"`
	Quantity *int32     `json:"quantity,omitempty"`
	ShipDate *time.Time `json:"shipDate,omitempty"`

	// Status Order Status
	Status *OrderStatus `json:"status,omitempty"`
}

// OrderStatus Order Status
type OrderStatus string

// Pet defines model for Pet.
type Pet struct {
	Category  *Category `json:"category,omitempty"`
	Id        *int64    `json:"id,omitempty"`
	Name      string    `json:"name"`
	PhotoUrls []string  `json:"photoUrls"`

	// Status pet status in the store
	Status *PetStatus `json:"status,omitempty"`
	Tags   *[]Tag     `json:"tags,omitempty"`
}

// PetStatus pet status in the store
type PetStatus string

// Tag defines model for Tag.
type Tag struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// User defines model for User.
type User struct {
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	Id        *int64  `json:"id,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Password  *string `json:"password,omitempty"`
	Phone     *string `json:"phone,omitempty"`

	// UserStatus User Status
	UserStatus *int32  `json:"userStatus,omitempty"`
	Username   *string `json:"username,omitempty"`
}

// UserArray defines model for UserArray.
type UserArray = []User

// FindPetsByStatusParams defines parameters for FindPetsByStatus.
type FindPetsByStatusParams struct {
	// Status Status values that need to be considered for filter
	Status *FindPetsByStatusParamsStatus `form:"status,omitempty" json:"status,omitempty"`
}

// FindPetsByStatusParamsStatus defines parameters for FindPetsByStatus.
type FindPetsByStatusParamsStatus string

// FindPetsByTagsParams defines parameters for FindPetsByTags.
type FindPetsByTagsParams struct {
	// Tags Tags to filter by
	Tags *[]string `form:"tags,omitempty" json:"tags,omitempty"`
}

// DeletePetParams defines parameters for DeletePet.
type DeletePetParams struct {
	ApiKey *string `json:"api_key,omitempty"`
}

// UpdatePetWithFormParams defines parameters for UpdatePetWithForm.
type UpdatePetWithFormParams struct {
	// Name Name of pet that needs to be updated
	Name *string `form:"name,omitempty" json:"name,omitempty"`

	// Status Status of pet that needs to be updated
	Status *string `form:"status,omitempty" json:"status,omitempty"`
}

// UploadFileParams defines parameters for UploadFile.
type UploadFileParams struct {
	// AdditionalMetadata Additional Metadata
	AdditionalMetadata *string `form:"additionalMetadata,omitempty" json:"additionalMetadata,omitempty"`
}

// CreateUsersWithListInputJSONBody defines parameters for CreateUsersWithListInput.
type CreateUsersWithListInputJSONBody = []User

// LoginUserParams defines parameters for LoginUser.
type LoginUserParams struct {
	// Username The user name for login
	Username *string `form:"username,omitempty" json:"username,omitempty"`

	// Password The password for login in clear text
	Password *string `form:"password,omitempty" json:"password,omitempty"`
}

// AddPetJSONRequestBody defines body for AddPet for application/json ContentType.
type AddPetJSONRequestBody = Pet

// AddPetFormdataRequestBody defines body for AddPet for application/x-www-form-urlencoded ContentType.
type AddPetFormdataRequestBody = Pet

// UpdatePetJSONRequestBody defines body for UpdatePet for application/json ContentType.
type UpdatePetJSONRequestBody = Pet

// UpdatePetFormdataRequestBody defines body for UpdatePet for application/x-www-form-urlencoded ContentType.
type UpdatePetFormdataRequestBody = Pet

// PlaceOrderJSONRequestBody defines body for PlaceOrder for application/json ContentType.
type PlaceOrderJSONRequestBody = Order

// PlaceOrderFormdataRequestBody defines body for PlaceOrder for application/x-www-form-urlencoded ContentType.
type PlaceOrderFormdataRequestBody = Order

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody = User

// CreateUserFormdataRequestBody defines body for CreateUser for application/x-www-form-urlencoded ContentType.
type CreateUserFormdataRequestBody = User

// CreateUsersWithListInputJSONRequestBody defines body for CreateUsersWithListInput for application/json ContentType.
type CreateUsersWithListInputJSONRequestBody = CreateUsersWithListInputJSONBody

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody = User

// UpdateUserFormdataRequestBody defines body for UpdateUser for application/x-www-form-urlencoded ContentType.
type UpdateUserFormdataRequestBody = User

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"NRcKJSi3rBN77Wcrd4//DgAA///gocX++z0AAA==",
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

	pathPrefix := path.Dir(pathToFile)

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(pathPrefix, "../packageB/spec.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
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
