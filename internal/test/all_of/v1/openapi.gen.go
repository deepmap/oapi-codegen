// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package v1

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Person defines model for Person.
type Person struct {
	// Embedded struct due to allOf(#/components/schemas/PersonProperties)
	PersonProperties `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	// Embedded fields due to inline allOf schema
}

// PersonProperties These are fields that specify a person. They are all optional, and
// would be used by an `Edit` style API endpoint, where each is optional.
type PersonProperties struct {
	FirstName          *string `json:"FirstName,omitempty"`
	GovernmentIDNumber *int64  `json:"GovernmentIDNumber,omitempty"`
	LastName           *string `json:"LastName,omitempty"`
}

// PersonWithID defines model for PersonWithID.
type PersonWithID struct {
	// Embedded struct due to allOf(#/components/schemas/Person)
	Person `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	ID int64 `json:"ID"`
	// Embedded fields due to inline allOf schema
}

// PersonWithMorePropertiesOutsideOfAllOf defines model for PersonWithMorePropertiesOutsideOfAllOf.
type PersonWithMorePropertiesOutsideOfAllOf struct {
	// Embedded struct due to allOf(#/components/schemas/Person)
	Person `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	AdditionalProperty string `json:"additionalProperty"`
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xVzY7bNhB+lQHboyBv0KIH3YK6DQS0WQNNmkO8wI7FkcWUGrLkyIZg+N0LUvIfvG26",
	"t/g0gDkf5/vh6KAa13vHxBJVdVCx6ajHXK4oRMepQmsfW1V9PqjvA7WqUt8tLl2LuWUxnV8F5ymIoaiO",
	"xUEF+nswgbSqPqtfTYjyHntShfoN5/Lp+FQoTbEJxotJ96kPnYlgIiD4DFnA3kgHPbJGcWGENgEBsgaL",
	"UYCxpwI2g4DLEGihXq6Zh35DoYQMt3eD1bAhCCRDYNKwGQHh+R3JM0QZLcHbVV3CJ4KewpZAOpqvX7M/",
	"c5omQXbSUYA/MnPYd6bpwLEdwQe3M5oinHhDa8jqWK5ZFUpGT6pSbvOFGlHHQt1JVh3utKBIgIFmIJAO",
	"BaKnxrTjWaFEksZ8DK09y1AkjdZ85j7EmTfD8y/aXDMHYu2dYSlg31EgIGy6ZMIJa2Lgb0a9GFodTuSi",
	"BMPbRO6d21Hgnljq5fvsRTrWutCjqEoZlp9+vIhiWGhLITWes3GPevxXET8Z6erla9OaM3pLagL56pjH",
	"4ibb9fL/JBkCNS5owHjJYRtcDwg/B0Khsw0l1AKNY0HDcc3J1ZTIOQSuBYTV9eNABtTazPEPFN0QGoKP",
	"H+vlf2Yvyfa7C3TJ4OMg0Wh6bN9OOr5W0G9GhNnWEdzEKDckKqBNoEbMjuJLqb5AzKKMLwfx2v4Xep7u",
	"NE9NhluX4YzY9N8HihIhSw1Z1JhxVKF2FOIk4JvyoXxIljlPjN6oSv1QPpRv0uQoXR564S021Dmrp2e2",
	"JblfJn+iNXmFRtgjC6CApbRBHRMkqAKiA0l+9fgXpWVDPXTo/TgJlahhAqu1qtTq6sokR/SO42mJtTjY",
	"PELyjziX6L01TQZYfJm/LVN8UvX1cM1vPAt5y+ya/TH//gkAAP//d44HiNkGAAA=",
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
