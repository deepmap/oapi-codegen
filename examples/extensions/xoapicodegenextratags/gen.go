// Package xoapicodegenextratags provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package xoapicodegenextratags

// Client defines model for Client.
type Client struct {
	Id   float32 `json:"id"`
	Name string  `json:"name"`
}

// ClientWithExtension defines model for ClientWithExtension.
type ClientWithExtension struct {
	Id   float32 `gorm:"primarykey" json:"id" safe-to-log:"true" validate:"required,min=1,max=256"`
	Name string  `json:"name"`
}
