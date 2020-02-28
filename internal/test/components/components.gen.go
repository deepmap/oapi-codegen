// Package components provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package components

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// AdditionalPropertiesObject1 defines model for AdditionalPropertiesObject1.
type AdditionalPropertiesObject1 struct {
	Id                   int            `json:"id"`
	Name                 string         `json:"name"`
	Optional             *string        `json:"optional,omitempty"`
	AdditionalProperties map[string]int `json:"-"`
}

// AdditionalPropertiesObject2 defines model for AdditionalPropertiesObject2.
type AdditionalPropertiesObject2 struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// AdditionalPropertiesObject3 defines model for AdditionalPropertiesObject3.
type AdditionalPropertiesObject3 struct {
	Name                 string                 `json:"name"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// AdditionalPropertiesObject4 defines model for AdditionalPropertiesObject4.
type AdditionalPropertiesObject4 struct {
	Inner                AdditionalPropertiesObject4_Inner `json:"inner"`
	Name                 string                            `json:"name"`
	AdditionalProperties map[string]interface{}            `json:"-"`
}

// AdditionalPropertiesObject4_Inner defines model for AdditionalPropertiesObject4.Inner.
type AdditionalPropertiesObject4_Inner struct {
	Name                 string                 `json:"name"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// AdditionalPropertiesObject5 defines model for AdditionalPropertiesObject5.
type AdditionalPropertiesObject5 struct {
	AdditionalProperties map[string]SchemaObject `json:"-"`
}

// ObjectWithJsonField defines model for ObjectWithJsonField.
type ObjectWithJsonField struct {
	Name   string          `json:"name"`
	Value1 json.RawMessage `json:"value1"`
	Value2 json.RawMessage `json:"value2,omitempty"`
}

// SchemaObject defines model for SchemaObject.
type SchemaObject struct {
	FirstName string `json:"firstName"`
	Role      string `json:"role"`
}

// ParameterObject defines model for ParameterObject.
type ParameterObject string

// ResponseObject defines model for ResponseObject.
type ResponseObject struct {
	Field SchemaObject `json:"Field"`
}

// RequestBody defines model for RequestBody.
type RequestBody struct {
	Field SchemaObject `json:"Field"`
}

// ParamsWithAddPropsParams_P1 defines parameters for ParamsWithAddProps.
type ParamsWithAddPropsParams_P1 struct {
	AdditionalProperties map[string]interface{} `json:"-"`
}

// ParamsWithAddPropsParams defines parameters for ParamsWithAddProps.
type ParamsWithAddPropsParams struct {

	// This parameter has additional properties
	P1 ParamsWithAddPropsParams_P1 `json:"p1"`

	// This parameter has an anonymous inner property which needs to be
	// turned into a proper type for additionalProperties to work
	P2 struct {
		Inner ParamsWithAddPropsParams_P2_Inner `json:"inner"`
	} `json:"p2"`
}

// ParamsWithAddPropsParams_P2_Inner defines parameters for ParamsWithAddProps.
type ParamsWithAddPropsParams_P2_Inner struct {
	AdditionalProperties map[string]string `json:"-"`
}

// BodyWithAddPropsJSONBody defines parameters for BodyWithAddProps.
type BodyWithAddPropsJSONBody struct {
	Inner                BodyWithAddPropsJSONBody_Inner `json:"inner"`
	Name                 string                         `json:"name"`
	AdditionalProperties map[string]interface{}         `json:"-"`
}

// BodyWithAddPropsJSONBody_Inner defines parameters for BodyWithAddProps.
type BodyWithAddPropsJSONBody_Inner struct {
	AdditionalProperties map[string]int `json:"-"`
}

// BodyWithAddPropsRequestBody defines body for BodyWithAddProps for application/json ContentType.
type BodyWithAddPropsJSONRequestBody BodyWithAddPropsJSONBody

// Getter for additional properties for ParamsWithAddPropsParams_P1. Returns the specified
// element and whether it was found
func (a ParamsWithAddPropsParams_P1) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ParamsWithAddPropsParams_P1
func (a *ParamsWithAddPropsParams_P1) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ParamsWithAddPropsParams_P1 to handle AdditionalProperties
func (a *ParamsWithAddPropsParams_P1) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ParamsWithAddPropsParams_P1 to handle AdditionalProperties
func (a ParamsWithAddPropsParams_P1) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for ParamsWithAddPropsParams_P2_Inner. Returns the specified
// element and whether it was found
func (a ParamsWithAddPropsParams_P2_Inner) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ParamsWithAddPropsParams_P2_Inner
func (a *ParamsWithAddPropsParams_P2_Inner) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ParamsWithAddPropsParams_P2_Inner to handle AdditionalProperties
func (a *ParamsWithAddPropsParams_P2_Inner) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ParamsWithAddPropsParams_P2_Inner to handle AdditionalProperties
func (a ParamsWithAddPropsParams_P2_Inner) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for BodyWithAddPropsJSONBody. Returns the specified
// element and whether it was found
func (a BodyWithAddPropsJSONBody) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for BodyWithAddPropsJSONBody
func (a *BodyWithAddPropsJSONBody) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for BodyWithAddPropsJSONBody to handle AdditionalProperties
func (a *BodyWithAddPropsJSONBody) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["inner"]; found {
		err = json.Unmarshal(raw, &a.Inner)
		if err != nil {
			return errors.Wrap(err, "error reading 'inner'")
		}
		delete(object, "inner")
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return errors.Wrap(err, "error reading 'name'")
		}
		delete(object, "name")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for BodyWithAddPropsJSONBody to handle AdditionalProperties
func (a BodyWithAddPropsJSONBody) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["inner"], err = json.Marshal(a.Inner)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'inner'"))
	}

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'name'"))
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for BodyWithAddPropsJSONBody_Inner. Returns the specified
// element and whether it was found
func (a BodyWithAddPropsJSONBody_Inner) Get(fieldName string) (value int, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for BodyWithAddPropsJSONBody_Inner
func (a *BodyWithAddPropsJSONBody_Inner) Set(fieldName string, value int) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]int)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for BodyWithAddPropsJSONBody_Inner to handle AdditionalProperties
func (a *BodyWithAddPropsJSONBody_Inner) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]int)
		for fieldName, fieldBuf := range object {
			var fieldVal int
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for BodyWithAddPropsJSONBody_Inner to handle AdditionalProperties
func (a BodyWithAddPropsJSONBody_Inner) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for AdditionalPropertiesObject1. Returns the specified
// element and whether it was found
func (a AdditionalPropertiesObject1) Get(fieldName string) (value int, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AdditionalPropertiesObject1
func (a *AdditionalPropertiesObject1) Set(fieldName string, value int) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]int)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AdditionalPropertiesObject1 to handle AdditionalProperties
func (a *AdditionalPropertiesObject1) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["id"]; found {
		err = json.Unmarshal(raw, &a.Id)
		if err != nil {
			return errors.Wrap(err, "error reading 'id'")
		}
		delete(object, "id")
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return errors.Wrap(err, "error reading 'name'")
		}
		delete(object, "name")
	}

	if raw, found := object["optional"]; found {
		err = json.Unmarshal(raw, &a.Optional)
		if err != nil {
			return errors.Wrap(err, "error reading 'optional'")
		}
		delete(object, "optional")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]int)
		for fieldName, fieldBuf := range object {
			var fieldVal int
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for AdditionalPropertiesObject1 to handle AdditionalProperties
func (a AdditionalPropertiesObject1) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["id"], err = json.Marshal(a.Id)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'id'"))
	}

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'name'"))
	}

	if a.Optional != nil {
		object["optional"], err = json.Marshal(a.Optional)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'optional'"))
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for AdditionalPropertiesObject3. Returns the specified
// element and whether it was found
func (a AdditionalPropertiesObject3) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AdditionalPropertiesObject3
func (a *AdditionalPropertiesObject3) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AdditionalPropertiesObject3 to handle AdditionalProperties
func (a *AdditionalPropertiesObject3) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return errors.Wrap(err, "error reading 'name'")
		}
		delete(object, "name")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for AdditionalPropertiesObject3 to handle AdditionalProperties
func (a AdditionalPropertiesObject3) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'name'"))
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for AdditionalPropertiesObject4. Returns the specified
// element and whether it was found
func (a AdditionalPropertiesObject4) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AdditionalPropertiesObject4
func (a *AdditionalPropertiesObject4) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AdditionalPropertiesObject4 to handle AdditionalProperties
func (a *AdditionalPropertiesObject4) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["inner"]; found {
		err = json.Unmarshal(raw, &a.Inner)
		if err != nil {
			return errors.Wrap(err, "error reading 'inner'")
		}
		delete(object, "inner")
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return errors.Wrap(err, "error reading 'name'")
		}
		delete(object, "name")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for AdditionalPropertiesObject4 to handle AdditionalProperties
func (a AdditionalPropertiesObject4) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["inner"], err = json.Marshal(a.Inner)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'inner'"))
	}

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'name'"))
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for AdditionalPropertiesObject4_Inner. Returns the specified
// element and whether it was found
func (a AdditionalPropertiesObject4_Inner) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AdditionalPropertiesObject4_Inner
func (a *AdditionalPropertiesObject4_Inner) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AdditionalPropertiesObject4_Inner to handle AdditionalProperties
func (a *AdditionalPropertiesObject4_Inner) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return errors.Wrap(err, "error reading 'name'")
		}
		delete(object, "name")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for AdditionalPropertiesObject4_Inner to handle AdditionalProperties
func (a AdditionalPropertiesObject4_Inner) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'name'"))
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for AdditionalPropertiesObject5. Returns the specified
// element and whether it was found
func (a AdditionalPropertiesObject5) Get(fieldName string) (value SchemaObject, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AdditionalPropertiesObject5
func (a *AdditionalPropertiesObject5) Set(fieldName string, value SchemaObject) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]SchemaObject)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AdditionalPropertiesObject5 to handle AdditionalProperties
func (a *AdditionalPropertiesObject5) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]SchemaObject)
		for fieldName, fieldBuf := range object {
			var fieldVal SchemaObject
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for AdditionalPropertiesObject5 to handle AdditionalProperties
func (a AdditionalPropertiesObject5) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

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
	// https://api.deepmap.com for example.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestEditor RequestEditorFn
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
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
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
		c.RequestEditor = fn
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// ParamsWithAddProps request
	ParamsWithAddProps(ctx context.Context, params *ParamsWithAddPropsParams) (*http.Response, error)

	// BodyWithAddProps request  with any body
	BodyWithAddPropsWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	BodyWithAddProps(ctx context.Context, body BodyWithAddPropsJSONRequestBody) (*http.Response, error)
}

func (c *Client) ParamsWithAddProps(ctx context.Context, params *ParamsWithAddPropsParams) (*http.Response, error) {
	req, err := NewParamsWithAddPropsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) BodyWithAddPropsWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewBodyWithAddPropsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) BodyWithAddProps(ctx context.Context, body BodyWithAddPropsJSONRequestBody) (*http.Response, error) {
	req, err := NewBodyWithAddPropsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// NewParamsWithAddPropsRequest generates requests for ParamsWithAddProps
func NewParamsWithAddPropsRequest(server string, params *ParamsWithAddPropsParams) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/params_with_add_props")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if queryFrag, err := runtime.StyleParam("simple", true, "p1", params.P1); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if queryFrag, err := runtime.StyleParam("form", true, "p2", params.P2); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewBodyWithAddPropsRequest calls the generic BodyWithAddProps builder with application/json body
func NewBodyWithAddPropsRequest(server string, body BodyWithAddPropsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewBodyWithAddPropsRequestWithBody(server, "application/json", bodyReader)
}

// NewBodyWithAddPropsRequestWithBody generates requests for BodyWithAddProps with any type of body
func NewBodyWithAddPropsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/params_with_add_props")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
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
		if !strings.HasSuffix(baseURL, "/") {
			baseURL += "/"
		}
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

type Response interface {
	Status() string
	StatusCode() int
}

type paramsWithAddPropsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r paramsWithAddPropsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r paramsWithAddPropsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type bodyWithAddPropsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r bodyWithAddPropsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r bodyWithAddPropsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ParamsWithAddPropsWithResponse request returning *ParamsWithAddPropsResponse
func (c *ClientWithResponses) ParamsWithAddPropsWithResponse(ctx context.Context, params *ParamsWithAddPropsParams) (*paramsWithAddPropsResponse, error) {
	rsp, err := c.ParamsWithAddProps(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseParamsWithAddPropsResponse(rsp)
}

// BodyWithAddPropsWithBodyWithResponse request with arbitrary body returning *BodyWithAddPropsResponse
func (c *ClientWithResponses) BodyWithAddPropsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*bodyWithAddPropsResponse, error) {
	rsp, err := c.BodyWithAddPropsWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseBodyWithAddPropsResponse(rsp)
}

func (c *ClientWithResponses) BodyWithAddPropsWithResponse(ctx context.Context, body BodyWithAddPropsJSONRequestBody) (*bodyWithAddPropsResponse, error) {
	rsp, err := c.BodyWithAddProps(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseBodyWithAddPropsResponse(rsp)
}

// ParseParamsWithAddPropsResponse parses an HTTP response from a ParamsWithAddPropsWithResponse call
func ParseParamsWithAddPropsResponse(rsp *http.Response) (*paramsWithAddPropsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &paramsWithAddPropsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ParseBodyWithAddPropsResponse parses an HTTP response from a BodyWithAddPropsWithResponse call
func ParseBodyWithAddPropsResponse(rsp *http.Response) (*bodyWithAddPropsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &bodyWithAddPropsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /params_with_add_props)
	ParamsWithAddProps(ctx echo.Context, params ParamsWithAddPropsParams) error

	// (POST /params_with_add_props)
	BodyWithAddProps(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ParamsWithAddProps converts echo context to params.
func (w *ServerInterfaceWrapper) ParamsWithAddProps(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ParamsWithAddPropsParams
	// ------------- Required query parameter "p1" -------------

	err = runtime.BindQueryParameter("simple", true, true, "p1", ctx.QueryParams(), &params.P1)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter p1: %s", err))
	}

	// ------------- Required query parameter "p2" -------------

	err = runtime.BindQueryParameter("form", true, true, "p2", ctx.QueryParams(), &params.P2)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter p2: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ParamsWithAddProps(ctx, params)
	return err
}

// BodyWithAddProps converts echo context to params.
func (w *ServerInterfaceWrapper) BodyWithAddProps(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.BodyWithAddProps(ctx)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/params_with_add_props", wrapper.ParamsWithAddProps)
	router.POST("/params_with_add_props", wrapper.BodyWithAddProps)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9xWS2/bOBD+KwR3j4Sdx+5FNy8WRVOgbdAE6CExAkYcRUxlUiHpuEKg/14MKVm2RLly",
	"mkt7sklxHt83zxea6lWpFShnafJCS274ChwYf7psT5/vHyF1eJVq5UD5v7wsC5lyJ7WaP1qt8M6mOay4",
	"12R0CcZJ8JreSSgE/vnbQEYT+te8szsPQnZ+5X8bW3XNqIGntTQgaHLTaFjitYPvbl4WXPZMuqoEmlDr",
	"jFQPtManAmxqZIk+0oRyssVHGUVx+rQGU22NgXX/adH4/GV7Uf12yIMOW2plWzDh8IdEckGsXJUFkBYk",
	"0Z2xxgtUtBBCoggvLrcoglunHnjk8459qRw8gKED8++5JZ0s6RgiOiMoTKRylPWokyKuW/EVRFAzqstg",
	"IEbJPqdeBUMLS9Y+bRlhB1g4G2ch44WFPvD/NViitCO8KPQmzsGv4n4jaOfj0JxZD5AtEJAlXFURVNUA",
	"0xG+H+f2P8e57TNRaVWt9NqSDEuLbHKZ5iQfy9FhfJQC8zOzbwp/mnjwi72GxX8PVff0zjW97jfS5SQo",
	"IZk2RMjUPzKB8IHrwcJX6fIPVqttU53EMqPPvFiD72CZNivuaEJ932YjT88mPI2XXWMpxv4eVQPfM2ms",
	"+zQGwOhiQgL4V2xH1dKPAqkyjcKFTEFZ6JiiHy+uUbuTDtXTa7COXIF59mn0DMaGMJ7OTmYnocGC4qWk",
	"CT2fncxOsTK4y73/c78q2DsM7B0X4g7h+S8P4OH2BxJKhjTolijClSCc3GtRNVXZwItn0S2GBQ9+FF8I",
	"moQVzGKeLIS49C6wvS3tpu/JdS5t58J4G/DGdnagtippiTx0YQjl383mQ01ikCPWVT4SYVrTmk3xVu10",
	"NN8Dtm24IVEBCEucJvdwq9zaKBA4cDXhzcswg7EOY96i5Eabb+MMnB1k4KjuGUn+HkvRrres6+XeCqfW",
	"RVEzWmobyT7fl0izwe6mGy55XCr8KqSB1EUJYZint+og8ZjYMdlIzuLG3MtY88pdevpEmhqGnf3jlWOp",
	"XUjaONX9XKmHgavr+kcAAAD//xxDQTJrDQAA",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
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

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
