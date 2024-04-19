// Package xgoname provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package xgoname

// Client defines model for Client.
type Client struct {
	Id   *float32 `json:"id,omitempty"`
	Name string   `json:"name"`
}

// ClientRenamedByExtension defines model for ClientWithExtension.
type ClientRenamedByExtension struct {
	AccountIdentifier *float32 `json:"id,omitempty"`
	Name              string   `json:"name"`
}
