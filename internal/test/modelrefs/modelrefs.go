package modelrefs

// This is an example of a spec that uses $ref to reference the same schema
// multiple times, that would otherwise gennerate anonymous structs.
// See https://github.com/deepmap/oapi-codegen/issues/1362

//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -config cfg.yaml api.yaml
