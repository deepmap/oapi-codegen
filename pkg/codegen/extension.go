package codegen

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

const (
	extPropGoType       = "x-go-type"
	extPropOmitEmpty    = "x-omitempty"
	extPropExtraTags    = "x-oapi-codegen-extra-tags"
	extPropGoJsonIgnore = "x-go-json-ignore"
)

func extTypeName(extPropValue interface{}) (string, error) {
	raw, ok := extPropValue.(json.RawMessage)
	if !ok {
		return "", fmt.Errorf("failed to convert type: %T", extPropValue)
	}
	var name string
	if err := json.Unmarshal(raw, &name); err != nil {
		return "", errors.Wrap(err, "failed to unmarshal json")
	}

	return name, nil
}

func extParseOmitEmpty(extPropValue interface{}) (bool, error) {
	raw, ok := extPropValue.(json.RawMessage)
	if !ok {
		return false, fmt.Errorf("failed to convert type: %T", extPropValue)
	}

	var omitEmpty bool
	if err := json.Unmarshal(raw, &omitEmpty); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal json")
	}

	return omitEmpty, nil
}

func extExtraTags(extPropValue interface{}) (map[string]string, error) {
	raw, ok := extPropValue.(json.RawMessage)
	if !ok {
		return nil, fmt.Errorf("failed to convert type: %T", extPropValue)
	}
	var tags map[string]string
	if err := json.Unmarshal(raw, &tags); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal json")
	}
	return tags, nil
}

func extParseGoJsonIgnore(extPropValue interface{}) (bool, error) {

	raw, ok := extPropValue.(json.RawMessage)
	if !ok {
		return false, fmt.Errorf("failed to convert type: %T", extPropValue)
	}

	var goJsonIgnore bool
	if err := json.Unmarshal(raw, &goJsonIgnore); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal json")
	}

	return goJsonIgnore, nil
}
