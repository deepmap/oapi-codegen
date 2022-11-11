package codegen

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/require"
)

type testSchema struct {
	schema openapi3.Schema
	path   []string
	want   Schema
	err    string
}

func Test_oapiSchemaToGoType(t *testing.T) {
	tests := make(map[string]*testSchema)
	addTest := func(schemaType, schemaFormat, want, err string) {
		tests[schemaType+"-"+schemaFormat] = &testSchema{
			schema: openapi3.Schema{
				Type:   schemaType,
				Format: schemaFormat,
			},
			want: Schema{
				GoType:         want,
				DefineViaAlias: true,
			},
			err: err,
		}
	}

	// Integers.
	for _, format := range []string{
		"int64",
		"int32", "int16", "int8", "int",
		"uint64", "uint32", "uint16", "uint8", "uint",
		"unknown",
	} {
		want := format
		switch format {
		case "unknown", "":
			want = "int"
		}
		addTest("integer", format, want, "")
	}

	// Numbers.
	for _, tt := range [][3]string{
		{"double", "float64"},
		{"float", "float32"},
		{"", "float32"},
		{"unknown", "", "invalid number format: unknown"},
	} {
		addTest("number", tt[0], tt[1], tt[2])
	}

	// Booleans.
	addTest("boolean", "", "bool", "")
	addTest("boolean", "unknown", "", "invalid format (unknown) for boolean")

	// Strings.
	for _, tt := range [][3]string{
		{"byte", "[]byte"},
		{"email", "openapi_types.Email"},
		{"date", "openapi_types.Date"},
		{"date-time", "time.Time"},
		{"json", "json.RawMessage"},
		{"uuid", "openapi_types.UUID"},
		{"binary", "openapi_types.File"},
		{"", "string"},
		{"other", "string"},
	} {
		addTest("string", tt[0], tt[1], tt[2])
	}
	tests["string-json"].want.SkipOptionalPointer = true

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			outSchema := &Schema{}
			err := oapiSchemaToGoType(&tt.schema, tt.path, outSchema)
			if tt.err != "" {
				require.EqualError(t, err, tt.err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, *outSchema)
		})
	}

	// Overrides
	old := globalState.options.OutputOptions.TypeMappings
	globalState.options.OutputOptions.TypeMappings = map[string]TypeMapping{
		"string-uuid": {
			GoType:              "mypkg.UUID",
			SkipOptionalPointer: true,
		},
	}
	defer func() { globalState.options.OutputOptions.TypeMappings = old }()

	tests = make(map[string]*testSchema)
	addTest("string", "uuid", "mypkg.UUID", "")
	tests["string-uuid"].want.SkipOptionalPointer = true
	tests["string-uuid"].want.DefineViaAlias = false

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			outSchema := &Schema{}
			err := oapiSchemaToGoType(&tt.schema, tt.path, outSchema)
			if tt.err != "" {
				require.EqualError(t, err, tt.err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, *outSchema)
		})
	}
}
