// Package xgotypename provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package xgotypename

// Client defines model for Client.
type Client struct {
	Id   *float32 `json:"id,omitempty"`
	Name string   `json:"name"`
}

// ClientWithExtension defines model for ClientWithExtension.
type ClientWithExtension = ClientRenamedByExtension

// ClientRenamedByExtension defines model for .
type ClientRenamedByExtension struct {
	Id   *float32 `json:"id,omitempty"`
	Name string   `json:"name"`
}
