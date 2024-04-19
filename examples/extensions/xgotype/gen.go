// Package xgotype provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package xgotype

import (
	googleuuid "github.com/google/uuid"
)

// Client defines model for Client.
type Client struct {
	Id   *float32 `json:"id,omitempty"`
	Name string   `json:"name"`
}

// ClientWithExtension defines model for ClientWithExtension.
type ClientWithExtension struct {
	Id   *int64          `json:"id,omitempty"`
	Name googleuuid.UUID `json:"name"`
}
