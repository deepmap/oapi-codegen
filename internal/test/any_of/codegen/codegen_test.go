package codegen

import (
	"github.com/deepmap/oapi-codegen/pkg/codegen"
	"github.com/deepmap/oapi-codegen/pkg/util"
	"github.com/golangci/lint-1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go/format"
	"testing"
)

func TestAnyOfInlineSchema(t *testing.T) {
	opts := codegen.Configuration{
		PackageName: "api",
		Generate: codegen.GenerateOptions{
			Models:     true,
			EchoServer: true,
			Client:     true,
		},
	}
	swagger, err := util.LoadSwagger("anyof-inline.yaml")
	assert.NoError(t, err)

	// Generate code
	code, err := codegen.Generate(swagger, opts)

	validateGeneratedCode(t, err, code)
}

func TestAnyOfRefSchema(t *testing.T) {
	opts := codegen.Configuration{
		PackageName: "api",
		Generate: codegen.GenerateOptions{
			Models:     true,
			EchoServer: true,
			Client:     true,
		},
	}
	swagger, err := util.LoadSwagger("anyof-ref-schema.yaml")
	require.NoError(t, err)

	// Generate code
	code, err := codegen.Generate(swagger, opts)
	assert.NoError(t, err)

	validateGeneratedCode(t, err, code)
}

func validateGeneratedCode(t *testing.T, err error, code string) {
	assert.NoError(t, err)
	assert.NotEmpty(t, code)

	// Check that we have valid (formattable) code:
	_, err = format.Source([]byte(code))
	assert.NoError(t, err)

	// Check that we have a package:
	assert.Contains(t, code, "package api")

	// Make sure the generated code is valid:
	checkLint(t, "test.gen.go", []byte(code))
}

func checkLint(t *testing.T, filename string, code []byte) {
	linter := new(lint.Linter)
	problems, err := linter.Lint(filename, code)
	assert.NoError(t, err)
	assert.Len(t, problems, 0)
}
