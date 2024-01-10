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
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"text/template"

	"github.com/pb33f/libopenapi"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
)

// GenerateInlinedSpec generates a gzipped, base64 encoded JSON representation of the
// swagger definition, which we embed inside the generated code.
func GenerateInlinedSpec(t *template.Template, importMapping importMap, openapi *libopenapi.DocumentModel[v3.Document]) (string, error) {

	encoded := openapi.Model.RenderJSON("")

	// We are doing all this madness because libopenapi doesn't provide JSON Marshalling without indention
	v := json.RawMessage{}
	err := json.Unmarshal(encoded, &v)
	if err != nil {
		return "", fmt.Errorf("error marshaling openapi: %w", err)
	}
	encoded, err = json.Marshal(v)
	if err != nil {
		return "", fmt.Errorf("error marshaling openapi: %w", err)
	}

	var buf bytes.Buffer
	zw, err := gzip.NewWriterLevel(&buf, gzip.BestCompression)
	if err != nil {
		return "", fmt.Errorf("error creating gzip compressor: %w", err)
	}
	_, err = zw.Write(encoded)
	if err != nil {
		return "", fmt.Errorf("error gzipping swagger file: %w", err)
	}
	err = zw.Close()
	if err != nil {
		return "", fmt.Errorf("error gzipping swagger file: %w", err)
	}
	str := base64.StdEncoding.EncodeToString(buf.Bytes())

	var parts []string
	const width = 80

	// Chop up the string into an array of strings.
	for len(str) > width {
		part := str[0:width]
		parts = append(parts, part)
		str = str[width:]
	}
	if len(str) > 0 {
		parts = append(parts, str)
	}

	return GenerateTemplates(
		[]string{"inline.tmpl"},
		t,
		struct {
			SpecParts     []string
			ImportMapping importMap
		}{
			SpecParts:     parts,
			ImportMapping: importMapping,
		})
}
