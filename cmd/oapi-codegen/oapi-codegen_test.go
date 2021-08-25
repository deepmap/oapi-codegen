package main

import (
	"testing"

	"github.com/deepmap/oapi-codegen/pkg/util"
)

func TestLoader(t *testing.T) {

	paths := []string{
		"../../examples/petstore-expanded/petstore-expanded.yaml",
		"https://petstore3.swagger.io/api/v3/openapi.json",
	}

	for _, v := range paths {

		spec, err := util.LoadSpec(v)
		if err != nil {
			t.Error(err)
		}
		if spec == nil || spec.Info == nil || spec.Info.Version == "" {
			t.Error("missing data")
		}
	}
}
