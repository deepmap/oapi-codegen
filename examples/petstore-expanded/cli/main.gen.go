// Package main provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"text/template"
)

// Error defines model for Error.
type Error struct {

	// Error code
	Code int32 `json:"code"`

	// Error message
	Message string `json:"message"`
}

// NewPet defines model for NewPet.
type NewPet struct {

	// Name of the pet
	Name string `json:"name"`

	// Type of the pet
	Tag *string `json:"tag,omitempty"`
}

// Pet defines model for Pet.
type Pet struct {
	// Embedded struct due to allOf(#/components/schemas/NewPet)
	NewPet
	// Embedded fields due to inline allOf schema

	// Unique id of the pet
	Id int64 `json:"id"`
}

// FindPetsParams defines parameters for FindPets.
type FindPetsParams struct {

	// tags to filter by
	Tags *[]string `json:"tags,omitempty"`

	// maximum number of results to return
	Limit *int32 `json:"limit,omitempty"`
}

// AddPetJSONBody defines parameters for AddPet.
type AddPetJSONBody NewPet

// AddPetRequestBody defines body for AddPet for application/json ContentType.
type AddPetJSONRequestBody AddPetJSONBody

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
	// FindPets request
	FindPets(ctx context.Context, params *FindPetsParams) (*http.Response, error)

	// AddPet request  with any body
	AddPetWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	AddPet(ctx context.Context, body AddPetJSONRequestBody) (*http.Response, error)

	// DeletePet request
	DeletePet(ctx context.Context, id int64) (*http.Response, error)

	// FindPetById request
	FindPetById(ctx context.Context, id int64) (*http.Response, error)
}

func (c *Client) FindPets(ctx context.Context, params *FindPetsParams) (*http.Response, error) {
	req, err := NewFindPetsRequest(c.Server, params)
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

func (c *Client) AddPetWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewAddPetRequestWithBody(c.Server, contentType, body)
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

func (c *Client) AddPet(ctx context.Context, body AddPetJSONRequestBody) (*http.Response, error) {
	req, err := NewAddPetRequest(c.Server, body)
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

func (c *Client) DeletePet(ctx context.Context, id int64) (*http.Response, error) {
	req, err := NewDeletePetRequest(c.Server, id)
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

func (c *Client) FindPetById(ctx context.Context, id int64) (*http.Response, error) {
	req, err := NewFindPetByIdRequest(c.Server, id)
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

// NewFindPetsRequest generates requests for FindPets
func NewFindPetsRequest(server string, params *FindPetsParams) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/pets")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if params.Tags != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "tags", *params.Tags); err != nil {
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

	}

	if params.Limit != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "limit", *params.Limit); err != nil {
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

	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewAddPetRequest calls the generic AddPet builder with application/json body
func NewAddPetRequest(server string, body AddPetJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAddPetRequestWithBody(server, "application/json", bodyReader)
}

// NewAddPetRequestWithBody generates requests for AddPet with any type of body
func NewAddPetRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/pets")
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

// NewDeletePetRequest generates requests for DeletePet
func NewDeletePetRequest(server string, id int64) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "id", id)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/pets/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewFindPetByIdRequest generates requests for FindPetById
func NewFindPetByIdRequest(server string, id int64) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "id", id)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/pets/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

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

type findPetsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Pet
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r findPetsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r findPetsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type addPetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Pet
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r addPetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r addPetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type deletePetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r deletePetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r deletePetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type findPetByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Pet
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r findPetByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r findPetByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// FindPetsWithResponse request returning *FindPetsResponse
func (c *ClientWithResponses) FindPetsWithResponse(ctx context.Context, params *FindPetsParams) (*findPetsResponse, error) {
	rsp, err := c.FindPets(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseFindPetsResponse(rsp)
}

// AddPetWithBodyWithResponse request with arbitrary body returning *AddPetResponse
func (c *ClientWithResponses) AddPetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*addPetResponse, error) {
	rsp, err := c.AddPetWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseAddPetResponse(rsp)
}

func (c *ClientWithResponses) AddPetWithResponse(ctx context.Context, body AddPetJSONRequestBody) (*addPetResponse, error) {
	rsp, err := c.AddPet(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseAddPetResponse(rsp)
}

// DeletePetWithResponse request returning *DeletePetResponse
func (c *ClientWithResponses) DeletePetWithResponse(ctx context.Context, id int64) (*deletePetResponse, error) {
	rsp, err := c.DeletePet(ctx, id)
	if err != nil {
		return nil, err
	}
	return ParseDeletePetResponse(rsp)
}

// FindPetByIdWithResponse request returning *FindPetByIdResponse
func (c *ClientWithResponses) FindPetByIdWithResponse(ctx context.Context, id int64) (*findPetByIdResponse, error) {
	rsp, err := c.FindPetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return ParseFindPetByIdResponse(rsp)
}

// ParseFindPetsResponse parses an HTTP response from a FindPetsWithResponse call
func ParseFindPetsResponse(rsp *http.Response) (*findPetsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &findPetsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Pet
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json"):
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseAddPetResponse parses an HTTP response from a AddPetWithResponse call
func ParseAddPetResponse(rsp *http.Response) (*addPetResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &addPetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Pet
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json"):
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseDeletePetResponse parses an HTTP response from a DeletePetWithResponse call
func ParseDeletePetResponse(rsp *http.Response) (*deletePetResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &deletePetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json"):
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseFindPetByIdResponse parses an HTTP response from a FindPetByIdWithResponse call
func ParseFindPetByIdResponse(rsp *http.Response) (*findPetByIdResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &findPetByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Pet
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json"):
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

var (
	clientOptions = make([]ClientOption, 0)
	client        *ClientWithResponses

	appBeforeFuncs = make([]cli.BeforeFunc, 0)
	app            = cli.NewApp()
)

func init() {
	app.Flags = append(app.Flags,
		&cli.StringFlag{
			Name:    "server",
			Usage:   "server to send requests to",
			Aliases: []string{"S"},
		},
		&cli.StringFlag{
			Name:    "output",
			Usage:   "output format (text|table|json)",
			Aliases: []string{"O"},
			Value:   "text",
		},
	)

	app.Before = func(ctx *cli.Context) error {
		var err error

		for _, beforeFunc := range appBeforeFuncs {
			if err = beforeFunc(ctx); err != nil {
				return err
			}
		}

		if client, err = NewClientWithResponses(ctx.String("server"), clientOptions...); err != nil {
			return fmt.Errorf("unable to initialize client: %s\n", err)
		}

		return nil
	}

	app.Commands = make([]*cli.Command, 0)

	FindPetsCmd := &cli.Command{
		Name:  "find-pets",
		Usage: "Returns all pets",
		Flags: []cli.Flag{

			&cli.StringSliceFlag{Name: "tags", Usage: "tags to filter by"},

			&cli.IntFlag{Name: "limit", Usage: "maximum number of results to return"},
		},
		Action: func(ctx *cli.Context) error {
			var (
				params FindPetsParams
			)
			tags := ctx.StringSlice("tags")
			limit := int32(ctx.Int("limit"))

			params.Tags = &tags
			params.Limit = &limit

			resp, err := client.FindPetsWithResponse(ctx.Context, &params)
			if err != nil {
				return err
			}

			return outputResp(resp, ctx.String("output"))
		},
	}
	app.Commands = append(app.Commands, FindPetsCmd)

	AddPetCmd := &cli.Command{
		Name:  "add-pet",
		Usage: "Creates a new pet",
		Flags: []cli.Flag{

			&cli.StringFlag{Name: "body", Usage: "request body", Required: true},
		},
		Action: func(ctx *cli.Context) error {
			var (
				body AddPetJSONRequestBody
			)

			if err := json.Unmarshal([]byte(ctx.String("body")), &body); err != nil {
				return fmt.Errorf("unable to parse request body: %s", err)
			}

			resp, err := client.AddPetWithResponse(ctx.Context, body)
			if err != nil {
				return err
			}

			return outputResp(resp, ctx.String("output"))
		},
	}
	app.Commands = append(app.Commands, AddPetCmd)

	DeletePetCmd := &cli.Command{
		Name:  "delete-pet",
		Usage: "Deletes a pet by ID",
		Flags: []cli.Flag{

			&cli.IntFlag{Name: "id", Usage: "ID of pet to delete", Required: true},
		},
		Action: func(ctx *cli.Context) error {

			id := ctx.Int64("id")

			resp, err := client.DeletePetWithResponse(ctx.Context, id)
			if err != nil {
				return err
			}

			return outputResp(resp, ctx.String("output"))
		},
	}
	app.Commands = append(app.Commands, DeletePetCmd)

	FindPetByIdCmd := &cli.Command{
		Name:  "find-pet-by-id",
		Usage: "Returns a pet by ID",
		Flags: []cli.Flag{

			&cli.IntFlag{Name: "id", Usage: "ID of pet to fetch", Required: true},
		},
		Action: func(ctx *cli.Context) error {

			id := ctx.Int64("id")

			resp, err := client.FindPetByIdWithResponse(ctx.Context, id)
			if err != nil {
				return err
			}

			return outputResp(resp, ctx.String("output"))
		},
	}
	app.Commands = append(app.Commands, FindPetByIdCmd)

}

func getResponseObject(resp Response) interface{} {
	status := resp.StatusCode()

	v := reflect.ValueOf(resp)
	v = reflect.Indirect(v)
	t := v.Type()

	// If the interface is iterable (slice only), use the element type
	if v.Kind() == reflect.Slice {
		t = v.Type().Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Name == fmt.Sprintf("JSON%d", status) {
			return v.Field(i).Interface()
		}
	}

	return nil
}

func getResponseObjectFields(o interface{}) []string {
	var fields = make([]string, 0)

	v := reflect.ValueOf(o)
	v = reflect.Indirect(v)
	t := v.Type()

	// If the interface is iterable (slice only), use the element type
	if v.Kind() == reflect.Slice {
		t = v.Type().Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		fields = append(fields, t.Field(i).Name)
	}

	return fields
}

func outputJSON(o interface{}) error {
	j, err := json.Marshal(o)
	if err != nil {
		return fmt.Errorf("output JSON encoding failed: %s", err)
	}

	fmt.Println(string(j))
	return nil
}

func outputText(o interface{}) error {
	tplFields := getResponseObjectFields(o)
	for i := range tplFields {
		tplFields[i] = "{%if ." + tplFields[i] + "%}{%." + tplFields[i] + "%}{%else%}n/a{%end%}"
	}
	tpl := strings.Join(tplFields, "\t")

	t, err := template.New("out").Delims("{%", "%}").Parse(tpl)
	if err != nil {
		return fmt.Errorf("output template execution failed: %s", err)
	}

	// If the interface is iterable (slice only), we loop over the
	// items and perform the templating directly
	if v := reflect.ValueOf(o); reflect.Indirect(v).Kind() == reflect.Slice {
		for i := 0; i < reflect.Indirect(v).Len(); i++ {
			if err := t.Execute(os.Stdout, reflect.Indirect(v).Index(i).Interface()); err != nil {
				return fmt.Errorf("output template execution failed: %s", err)
			}
			fmt.Println()
		}
		return nil
	}

	if err := t.Execute(os.Stdout, o); err != nil {
		return fmt.Errorf("output template execution failed: %s", err)
	}

	fmt.Println()
	return nil
}

func outputTable(o interface{}) error {
	tab := tablewriter.NewWriter(os.Stdout)

	v := reflect.ValueOf(o)
	v = reflect.Indirect(v)
	t := v.Type()

	// If the interface is iterable (slice only), use the element type
	if v.Kind() == reflect.Slice {
		t = v.Type().Elem()
	}

	// Set up the table header
	headers := make([]string, 0)
	for i := 0; i < t.NumField(); i++ {
		headers = append(headers, t.Field(i).Name)
	}

	// If the interface is iterable (slice only), we loop over the items and display
	// each one in a table row
	if v := reflect.ValueOf(o); reflect.Indirect(v).Kind() == reflect.Slice {
		tab.SetHeader(headers)

		for i := 0; i < reflect.Indirect(v).Len(); i++ {
			item := reflect.Indirect(v).Index(i)
			row := make([]string, 0)

			for j := 0; j < item.NumField(); j++ {
				field := item.Field(j)
				switch field.Kind() {
				case reflect.Slice:
					// If the field value is a slice and is empty,
					// print "n/a" instead of an empty slice
					if field.Len() == 0 {
						row = append(row, "n/a")
					} else {
						row = append(row, fmt.Sprint(field.Interface()))
					}

				case reflect.Ptr:
					// If the field value is a nil pointer, print "n/a" instead of "<nil>"
					if field.IsNil() {
						row = append(row, "n/a")
					} else {
						row = append(row, fmt.Sprint(reflect.Indirect(field).Interface()))
					}

				default:
					row = append(row, fmt.Sprint(field.Interface()))
				}
			}

			tab.Append(row)
		}

		tab.Render()
		return nil
	}

	// Single item, loop over the type fields and display each item in a key/value-type table

	for i := 0; i < t.NumField(); i++ {
		label := t.Field(i).Name
		switch v.Field(i).Kind() {
		case reflect.Slice:
			// If the field value is a slice and is empty, print "n/a" instead of 0
			if n := v.Field(i).Len(); n == 0 {
				tab.Append([]string{label, "n/a"})
			} else {
				items := v.Field(i).Interface().([]string)
				tab.Append([]string{label, strings.Join(items, "\n")})
			}

		case reflect.Ptr:
			// If the field value is a nil pointer, print "n/a" instead of "<nil>"
			if v.Field(i).IsNil() {
				tab.Append([]string{label, "n/a"})
			} else {
				tab.Append([]string{label, fmt.Sprint(reflect.Indirect(v.Field(i)).Interface())})
			}

		default:
			tab.Append([]string{label, fmt.Sprint(v.Field(i).Interface())})
		}
	}

	tab.Render()
	return nil
}

func outputResp(resp Response, format string) error {
	o := getResponseObject(resp)
	if o == nil {
		return fmt.Errorf("%s", resp.Status())
	}

	switch format {
	case "text":
		return outputText(o)

	case "table":
		return outputTable(o)

	case "json":
		return outputJSON(o)

	default:
		return fmt.Errorf("unsupported output format %q", format)
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
