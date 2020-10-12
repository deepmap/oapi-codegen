// Package validator provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package validator

// Color defines model for Color.
type Color string

// List of Color
const (
	Color_black Color = "black"
	Color_white Color = "white"
)

// StructA defines model for StructA.
type StructA struct {
	RangeInt       *int64 `json:"rangeInt,omitempty" validate:"gte=3,lte=42"`
	RequiredString string `json:"requiredString" validate:"required"`
}

// StructB defines model for StructB.
type StructB struct {
	ListItem *[]string `json:"listItem,omitempty" validate:"gte=1"`
}

// StructC defines model for StructC.
type StructC struct {
	Color Color `json:"color" validate:"required,oneof=black white"`
}