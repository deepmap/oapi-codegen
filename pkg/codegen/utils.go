// Copyright 2019 DeepMap, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package codegen

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/getkin/kin-openapi/openapi3"
)

var (
	pathParamRE    *regexp.Regexp
	predeclaredSet map[string]struct{}
	separatorSet   map[rune]struct{}
)

func init() {
	pathParamRE = regexp.MustCompile("{[.;?]?([^{}*]+)\\*?}")

	predeclaredIdentifiers := []string{
		// Types
		"bool",
		"byte",
		"complex64",
		"complex128",
		"error",
		"float32",
		"float64",
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"rune",
		"string",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"uintptr",
		// Constants
		"true",
		"false",
		"iota",
		// Zero value
		"nil",
		// Functions
		"append",
		"cap",
		"close",
		"complex",
		"copy",
		"delete",
		"imag",
		"len",
		"make",
		"new",
		"panic",
		"print",
		"println",
		"real",
		"recover",
	}
	predeclaredSet = map[string]struct{}{}
	for _, id := range predeclaredIdentifiers {
		predeclaredSet[id] = struct{}{}
	}

	separators := "-#@!$&=.+:;_~ (){}[]"
	separatorSet = map[rune]struct{}{}
	for _, r := range separators {
		separatorSet[r] = struct{}{}
	}
}

// Uppercase the first character in a string. This assumes UTF-8, so we have
// to be careful with unicode, don't treat it as a byte array.
func UppercaseFirstCharacter(str string) string {
	if str == "" {
		return ""
	}
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// Same as above, except lower case
func LowercaseFirstCharacter(str string) string {
	if str == "" {
		return ""
	}
	runes := []rune(str)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

// This function will convert query-arg style strings to CamelCase. We will
// use `., -, +, :, ;, _, ~, ' ', (, ), {, }, [, ]` as valid delimiters for words.
// So, "word.word-word+word:word;word_word~word word(word)word{word}[word]"
// would be converted to WordWordWordWordWordWordWordWordWordWordWordWordWord
func ToCamelCase(str string) string {
	s := strings.Trim(str, " ")

	n := ""
	capNext := true
	for _, v := range s {
		if unicode.IsUpper(v) {
			n += string(v)
		}
		if unicode.IsDigit(v) {
			n += string(v)
		}
		if unicode.IsLower(v) {
			if capNext {
				n += strings.ToUpper(string(v))
			} else {
				n += string(v)
			}
		}
		_, capNext = separatorSet[v]
	}
	return n
}

// This function returns the keys of the given SchemaRef dictionary in sorted
// order, since Golang scrambles dictionary keys
func SortedSchemaKeys(dict map[string]*openapi3.SchemaRef) []string {
	keys := make([]string, len(dict))
	i := 0
	for key := range dict {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	return keys
}

// This function is the same as above, except it sorts the keys for a Paths
// dictionary.
func SortedPathsKeys(dict openapi3.Paths) []string {
	keys := make([]string, len(dict))
	i := 0
	for key := range dict {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	return keys
}

// This function returns Operation dictionary keys in sorted order
func SortedOperationsKeys(dict map[string]*openapi3.Operation) []string {
	keys := make([]string, len(dict))
	i := 0
	for key := range dict {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	return keys
}

// This function returns Responses dictionary keys in sorted order
func SortedResponsesKeys(dict openapi3.Responses) []string {
	keys := make([]string, len(dict))
	i := 0
	for key := range dict {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	return keys
}

func SortedHeadersKeys(dict openapi3.Headers) []string {
	keys := make([]string, len(dict))
	i := 0
	for key := range dict {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	return keys
}

// This returns Content dictionary keys in sorted order
func SortedContentKeys(dict openapi3.Content) []string {
	keys := make([]string, len(dict))
	i := 0
	for key := range dict {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	return keys
}

// This returns string map keys in sorted order
func SortedStringKeys(dict map[string]string) []string {
	keys := make([]string, len(dict))
	i := 0
	for key := range dict {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	return keys
}

// This returns sorted keys for a ParameterRef dict
func SortedParameterKeys(dict map[string]*openapi3.ParameterRef) []string {
	keys := make([]string, len(dict))
	i := 0
	for key := range dict {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	return keys
}

func SortedRequestBodyKeys(dict map[string]*openapi3.RequestBodyRef) []string {
	keys := make([]string, len(dict))
	i := 0
	for key := range dict {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	return keys
}

func SortedSecurityRequirementKeys(sr openapi3.SecurityRequirement) []string {
	keys := make([]string, len(sr))
	i := 0
	for key := range sr {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	return keys
}

// This function checks whether the specified string is present in an array
// of strings
func StringInArray(str string, array []string) bool {
	for _, elt := range array {
		if elt == str {
			return true
		}
	}
	return false
}

// RefPathToGoType takes a $ref value and converts it to a Go typename.
// #/components/schemas/Foo -> Foo
// #/components/parameters/Bar -> Bar
// #/components/responses/Baz -> Baz
// Remote components (document.json#/Foo) are supported if they present in --import-mapping
// URL components (http://deepmap.com/schemas/document.json#/Foo) are supported if they present in --import-mapping
// Remote and URL also support standard local paths even though the spec doesn't mention them.
func RefPathToGoType(refPath string) (string, error) {
	return refPathToGoType(refPath, true)
}

// refPathToGoType returns the Go typename for refPath given its
func refPathToGoType(refPath string, local bool) (string, error) {
	if refPath[0] == '#' {
		pathParts := strings.Split(refPath, "/")
		depth := len(pathParts)
		if local {
			if depth != 4 {
				return "", fmt.Errorf("unexpected reference depth: %d for ref: %s local: %t", depth, refPath, local)
			}
		} else if depth != 4 && depth != 2 {
			return "", fmt.Errorf("unexpected reference depth: %d for ref: %s local: %t", depth, refPath, local)
		}

		// Schemas may have been renamed locally, so look up the actual name in
		// the spec.
		name, err := findSchemaNameByRefPath(refPath, globalState.spec)
		if err != nil {
			return "", fmt.Errorf("error finding ref: %s in spec: %v", refPath, err)
		}
		if name != "" {
			return name, nil
		}
		// lastPart now stores the final element of the type path. This is what
		// we use as the base for a type name.
		lastPart := pathParts[len(pathParts)-1]
		return SchemaNameToTypeName(lastPart), nil
	}
	pathParts := strings.Split(refPath, "#")
	if len(pathParts) != 2 {
		return "", fmt.Errorf("unsupported reference: %s", refPath)
	}
	remoteComponent, flatComponent := pathParts[0], pathParts[1]
	if goImport, ok := importMapping[remoteComponent]; !ok {
		return "", fmt.Errorf("unrecognized external reference '%s'; please provide the known import for this reference using option --import-mapping", remoteComponent)
	} else {
		goType, err := refPathToGoType("#"+flatComponent, false)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s.%s", goImport.Name, goType), nil
	}
}

// IsGoTypeReference takes a $ref value and checks if it has link to go type.
// #/components/schemas/Foo                     -> true
// ./local/file.yml#/components/parameters/Bar  -> true
// ./local/file.yml                             -> false
// The function can be used to check whether RefPathToGoType($ref) is possible.
//
func IsGoTypeReference(ref string) bool {
	return ref != "" && !IsWholeDocumentReference(ref)
}

// IsWholeDocumentReference takes a $ref value and checks if it is whole document reference.
// #/components/schemas/Foo                             -> false
// ./local/file.yml#/components/parameters/Bar          -> false
// ./local/file.yml                                     -> true
// http://deepmap.com/schemas/document.json             -> true
// http://deepmap.com/schemas/document.json#/Foo        -> false
//
func IsWholeDocumentReference(ref string) bool {
	return ref != "" && !strings.ContainsAny(ref, "#")
}

// SwaggerUriToEchoUri converts a OpenAPI style path URI with parameters to a
// Echo compatible path URI. We need to replace all of OpenAPI parameters with
// ":param". Valid input parameters are:
//   {param}
//   {param*}
//   {.param}
//   {.param*}
//   {;param}
//   {;param*}
//   {?param}
//   {?param*}
func SwaggerUriToEchoUri(uri string) string {
	return pathParamRE.ReplaceAllString(uri, ":$1")
}

// SwaggerUriToChiUri converts a swagger style path URI with parameters to a
// Chi compatible path URI. We need to replace all of Swagger parameters with
// "{param}". Valid input parameters are:
//   {param}
//   {param*}
//   {.param}
//   {.param*}
//   {;param}
//   {;param*}
//   {?param}
//   {?param*}
func SwaggerUriToChiUri(uri string) string {
	return pathParamRE.ReplaceAllString(uri, "{$1}")
}

// This function converts a swagger style path URI with parameters to a
// Gin compatible path URI. We need to replace all of Swagger parameters with
// ":param". Valid input parameters are:
//   {param}
//   {param*}
//   {.param}
//   {.param*}
//   {;param}
//   {;param*}
//   {?param}
//   {?param*}
func SwaggerUriToGinUri(uri string) string {
	return pathParamRE.ReplaceAllString(uri, ":$1")
}

// This function converts a swagger style path URI with parameters to a
// Gorilla compatible path URI. We need to replace all of Swagger parameters with
// ":param". Valid input parameters are:
//   {param}
//   {param*}
//   {.param}
//   {.param*}
//   {;param}
//   {;param*}
//   {?param}
//   {?param*}
func SwaggerUriToGorillaUri(uri string) string {
	return pathParamRE.ReplaceAllString(uri, "{$1}")
}

// Returns the argument names, in order, in a given URI string, so for
// /path/{param1}/{.param2*}/{?param3}, it would return param1, param2, param3
func OrderedParamsFromUri(uri string) []string {
	matches := pathParamRE.FindAllStringSubmatch(uri, -1)
	result := make([]string, len(matches))
	for i, m := range matches {
		result[i] = m[1]
	}
	return result
}

// Replaces path parameters of the form {param} with %s
func ReplacePathParamsWithStr(uri string) string {
	return pathParamRE.ReplaceAllString(uri, "%s")
}

// Reorders the given parameter definitions to match those in the path URI.
func SortParamsByPath(path string, in []ParameterDefinition) ([]ParameterDefinition, error) {
	pathParams := OrderedParamsFromUri(path)
	n := len(in)
	if len(pathParams) != n {
		return nil, fmt.Errorf("path '%s' has %d positional parameters, but spec has %d declared",
			path, len(pathParams), n)
	}
	out := make([]ParameterDefinition, len(in))
	for i, name := range pathParams {
		p := ParameterDefinitions(in).FindByName(name)
		if p == nil {
			return nil, fmt.Errorf("path '%s' refers to parameter '%s', which doesn't exist in specification",
				path, name)
		}
		out[i] = *p
	}
	return out, nil
}

// Returns whether the given string is a go keyword
func IsGoKeyword(str string) bool {
	keywords := []string{
		"break",
		"case",
		"chan",
		"const",
		"continue",
		"default",
		"defer",
		"else",
		"fallthrough",
		"for",
		"func",
		"go",
		"goto",
		"if",
		"import",
		"interface",
		"map",
		"package",
		"range",
		"return",
		"select",
		"struct",
		"switch",
		"type",
		"var",
	}

	for _, k := range keywords {
		if k == str {
			return true
		}
	}
	return false
}

// IsPredeclaredGoIdentifier returns whether the given string
// is a predefined go indentifier.
//
// See https://golang.org/ref/spec#Predeclared_identifiers
func IsPredeclaredGoIdentifier(str string) bool {
	_, exists := predeclaredSet[str]
	return exists
}

// IsGoIdentity checks if the given string can be used as an identity
// in the generated code like a type name or constant name.
//
// See https://golang.org/ref/spec#Identifiers
func IsGoIdentity(str string) bool {
	for i, c := range str {
		if !isValidRuneForGoID(i, c) {
			return false
		}
	}

	return IsGoKeyword(str)
}

func isValidRuneForGoID(index int, char rune) bool {
	if index == 0 && unicode.IsNumber(char) {
		return false
	}

	return unicode.IsLetter(char) || char == '_' || unicode.IsNumber(char)
}

// IsValidGoIdentity checks if the given string can be used as a
// name of variable, constant, or type.
func IsValidGoIdentity(str string) bool {
	if IsGoIdentity(str) {
		return false
	}

	return !IsPredeclaredGoIdentifier(str)
}

// SanitizeGoIdentity deletes and replaces the illegal runes in the given
// string to use the string as a valid identity.
func SanitizeGoIdentity(str string) string {
	sanitized := []rune(str)

	for i, c := range sanitized {
		if !isValidRuneForGoID(i, c) {
			sanitized[i] = '_'
		} else {
			sanitized[i] = c
		}
	}

	str = string(sanitized)

	if IsGoKeyword(str) || IsPredeclaredGoIdentifier(str) {
		str = "_" + str
	}

	if !IsValidGoIdentity(str) {
		panic("here is a bug")
	}

	return str
}

// SanitizeEnumNames fixes illegal chars in the enum names
// and removes duplicates
func SanitizeEnumNames(enumNames []string) map[string]string {
	dupCheck := make(map[string]int, len(enumNames))
	deDup := make([]string, 0, len(enumNames))

	for _, n := range enumNames {
		if _, dup := dupCheck[n]; !dup {
			deDup = append(deDup, n)
		}
		dupCheck[n] = 0
	}

	dupCheck = make(map[string]int, len(deDup))
	sanitizedDeDup := make(map[string]string, len(deDup))

	for _, n := range deDup {
		sanitized := SanitizeGoIdentity(SchemaNameToTypeName(n))

		if _, dup := dupCheck[sanitized]; !dup {
			sanitizedDeDup[sanitized] = n
		} else {
			sanitizedDeDup[sanitized+strconv.Itoa(dupCheck[sanitized])] = n
		}
		dupCheck[sanitized]++
	}

	return sanitizedDeDup
}

func typeNamePrefix(name string) (prefix string) {
	if len(name) == 0 {
		return "Empty"
	}
	for _, r := range name {
		switch r {
		case '$':
			if len(name) == 1 {
				return "DollarSign"
			}
		case '-':
			prefix += "Minus"
		case '+':
			prefix += "Plus"
		case '&':
			prefix += "And"
		case '|':
			prefix += "Or"
		case '~':
			prefix += "Tilde"
		case '=':
			prefix += "Equal"
		case '#':
			prefix += "Hash"
		case '.':
			prefix += "Dot"
		case '*':
			prefix += "Asterisk"
		case '^':
			prefix += "Caret"
		case '%':
			prefix += "Percent"
		default:
			// Prepend "N" to schemas starting with a number
			if prefix == "" && unicode.IsDigit(r) {
				return "N"
			}

			// break the loop, done parsing prefix
			return
		}
	}

	return
}

// SchemaNameToTypeName converts a Schema name to a valid Go type name. It converts to camel case, and makes sure the name is
// valid in Go
func SchemaNameToTypeName(name string) string {
	return typeNamePrefix(name) + ToCamelCase(name)
}

// According to the spec, additionalProperties may be true, false, or a
// schema. If not present, true is implied. If it's a schema, true is implied.
// If it's false, no additional properties are allowed. We're going to act a little
// differently, in that if you want additionalProperties code to be generated,
// you must specify an additionalProperties type
// If additionalProperties it true/false, this field will be non-nil.
func SchemaHasAdditionalProperties(schema *openapi3.Schema) bool {
	if schema.AdditionalPropertiesAllowed != nil && *schema.AdditionalPropertiesAllowed {
		return true
	}

	if schema.AdditionalProperties != nil {
		return true
	}
	return false
}

// This converts a path, like Object/field1/nestedField into a go
// type name.
func PathToTypeName(path []string) string {
	for i, p := range path {
		path[i] = ToCamelCase(p)
	}
	return strings.Join(path, "_")
}

// StringToGoComment renders a possible multi-line string as a valid Go-Comment.
// Each line is prefixed as a comment.
func StringToGoComment(in string) string {
	return stringToGoCommentWithPrefix(in, "")
}

// StringWithTypeNameToGoComment renders a possible multi-line string as a
// valid Go-Comment, including the name of the type being referenced. Each line
// is prefixed as a comment.
func StringWithTypeNameToGoComment(in, typeName string) string {
	return stringToGoCommentWithPrefix(in, typeName)
}

func stringToGoCommentWithPrefix(in, prefix string) string {
	if len(in) == 0 || len(strings.TrimSpace(in)) == 0 { // ignore empty comment
		return ""
	}

	// Normalize newlines from Windows/Mac to Linux
	in = strings.Replace(in, "\r\n", "\n", -1)
	in = strings.Replace(in, "\r", "\n", -1)

	// Add comment to each line
	var lines []string
	for i, line := range strings.Split(in, "\n") {
		s := "//"
		if i == 0 && len(prefix) > 0 {
			s += " " + prefix
		}
		lines = append(lines, fmt.Sprintf("%s %s", s, line))
	}
	in = strings.Join(lines, "\n")

	// in case we have a multiline string which ends with \n, we would generate
	// empty-line-comments, like `// `. Therefore remove this line comment.
	in = strings.TrimSuffix(in, "\n// ")
	return in
}

// This function breaks apart a path, and looks at each element. If it's
// not a path parameter, eg, {param}, it will URL-escape the element.
func EscapePathElements(path string) string {
	elems := strings.Split(path, "/")
	for i, e := range elems {
		if strings.HasPrefix(e, "{") && strings.HasSuffix(e, "}") {
			// This is a path parameter, we don't want to mess with its value
			continue
		}
		elems[i] = url.QueryEscape(e)
	}
	return strings.Join(elems, "/")
}

// renameSchema takes as input the name of a schema as provided in the spec,
// and the definition of the schema. If the schema overrides the name via
// x-go-name, the new name is returned, otherwise, the original name is
// returned.
func renameSchema(schemaName string, schemaRef *openapi3.SchemaRef) (string, error) {
	// References will not change type names.
	if schemaRef.Ref != "" {
		return SchemaNameToTypeName(schemaName), nil
	}
	schema := schemaRef.Value

	if extension, ok := schema.Extensions[extGoName]; ok {
		typeName, err := extTypeName(extension)
		if err != nil {
			return "", fmt.Errorf("invalid value for %q: %w", extPropGoType, err)
		}
		return typeName, nil
	}
	return SchemaNameToTypeName(schemaName), nil
}

// renameParameter generates the name for a parameter, taking x-go-name into
// account
func renameParameter(parameterName string, parameterRef *openapi3.ParameterRef) (string, error) {
	if parameterRef.Ref != "" {
		return SchemaNameToTypeName(parameterName), nil
	}
	parameter := parameterRef.Value

	if extension, ok := parameter.Extensions[extGoName]; ok {
		typeName, err := extTypeName(extension)
		if err != nil {
			return "", fmt.Errorf("invalid value for %q: %w", extPropGoType, err)
		}
		return typeName, nil
	}
	return SchemaNameToTypeName(parameterName), nil
}

// renameResponse generates the name for a parameter, taking x-go-name into
// account
func renameResponse(responseName string, responseRef *openapi3.ResponseRef) (string, error) {
	if responseRef.Ref != "" {
		return SchemaNameToTypeName(responseName), nil
	}
	response := responseRef.Value

	if extension, ok := response.Extensions[extGoName]; ok {
		typeName, err := extTypeName(extension)
		if err != nil {
			return "", fmt.Errorf("invalid value for %q: %w", extPropGoType, err)
		}
		return typeName, nil
	}
	return SchemaNameToTypeName(responseName), nil
}

// renameRequestBody generates the name for a parameter, taking x-go-name into
// account
func renameRequestBody(requestBodyName string, requestBodyRef *openapi3.RequestBodyRef) (string, error) {
	if requestBodyRef.Ref != "" {
		return SchemaNameToTypeName(requestBodyName), nil
	}
	requestBody := requestBodyRef.Value

	if extension, ok := requestBody.Extensions[extGoName]; ok {
		typeName, err := extTypeName(extension)
		if err != nil {
			return "", fmt.Errorf("invalid value for %q: %w", extPropGoType, err)
		}
		return typeName, nil
	}
	return SchemaNameToTypeName(requestBodyName), nil
}

// findSchemaByRefPath turns a $ref path into a schema. This will return ""
// if the schema wasn't found, and it'll only work successfully for schemas
// defined within the spec that we parsed.
func findSchemaNameByRefPath(refPath string, spec *openapi3.T) (string, error) {
	pathElements := strings.Split(refPath, "/")
	// All local references will have 4 path elements.
	if len(pathElements) != 4 {
		return "", nil
	}

	// We only support local references
	if pathElements[0] != "#" {
		return "", nil
	}
	// Only components are supported
	if pathElements[1] != "components" {
		return "", nil
	}
	propertyName := pathElements[3]
	switch pathElements[2] {
	case "schemas":
		if schema, found := spec.Components.Schemas[propertyName]; found {
			return renameSchema(propertyName, schema)
		}
	case "parameters":
		if parameter, found := spec.Components.Parameters[propertyName]; found {
			return renameParameter(propertyName, parameter)
		}
	case "responses":
		if response, found := spec.Components.Responses[propertyName]; found {
			return renameResponse(propertyName, response)
		}
	case "requestBodies":
		if requestBody, found := spec.Components.RequestBodies[propertyName]; found {
			return renameRequestBody(propertyName, requestBody)
		}
	}
	return "", nil
}

func GetImports(dict map[string]*openapi3.SchemaRef) (map[string]goImport, error) {
	res := map[string]goImport{}
	for _, v := range dict {
		if v == nil || v.Value == nil {
			continue
		}

		if v.Value.Extensions["x-go-type-import"] == nil || v.Value.Extensions["x-go-type"] == nil {
			continue
		}
		goTypeImportExt := v.Value.Extensions["x-go-type-import"]

		if raw, ok := goTypeImportExt.(json.RawMessage); ok {
			gi := goImport{}
			if err := json.Unmarshal(raw, &gi); err != nil {
				return nil, err
			}
			res[gi.String()] = gi
		} else {
			continue
		}
	}
	return res, nil
}
