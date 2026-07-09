// Package serialize provides modifiers that render a value to a data format.
package serialize

import (
	"encoding/json"
	"fmt"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameJSON is the template name for the JSON modifier.
const ModifierNameJSON functions.ModifierName = "json"

// JSON serializes the value to a JSON string.
// Pass 'pretty' as a parameter to produce indented output.
//
// value: any
// param:0?: string
// returns: string
//
// example: serialize a map to compact JSON
// in:  user = {"name": "Alice", "role": "admin"}
// tpl: {{ user | json }}
// out: {"name":"Alice","role":"admin"}
//
// example: produce indented JSON for human reading
// in:  config = {"region": "eu-west-1", "debug": false}
// tpl: {{ config | json:'pretty' }}
// out: {
// out:   "region": "eu-west-1",
// out:   "debug": false
// out: }
func JSON(value any, params []any) (any, error) {
	if len(params) > 0 && functions.IsParam(params, 0, "pretty") {
		jsonBytes, err := json.MarshalIndent(value, "", "  ")
		if err != nil {
			return "", fmt.Errorf("failed to apply json filter: %w", err)
		}

		return string(jsonBytes), nil
	}
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return "", fmt.Errorf("failed to apply json filter: %w", err)
	}

	return string(jsonBytes), nil
}
