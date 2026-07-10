// Package serialize provides modifiers that render a value to a data format.
package serialize

import (
	"encoding/json"
	"fmt"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameJSON is the template name for the JSON modifier.
const ModifierNameJSON functions.ModifierName = "json"

// JSON serializes the value to a JSON string via Go's encoding/json.
// By default the output is compact (no spaces or newlines). Pass the literal
// string 'pretty' as the first parameter to produce indented output using two
// spaces per level. Any other parameter value, or no parameter, keeps it
// compact. Object keys always come out sorted alphabetically, so the output is
// stable and comparable regardless of the input map's insertion order. Scalars
// serialize to their JSON forms (a string is quoted, a bool becomes true/false,
// nil becomes null) and slices become JSON arrays.
//
// value: any
// param:0: string (optional; 'pretty' for indented output, otherwise compact)
// returns: string
//
// example: serialize a map to compact JSON, keys sorted alphabetically
// in:  user = {"name": "Alice", "role": "admin"}
// tpl: {{ user | json }}
// out: {"name":"Alice","role":"admin"}
//
// example: produce indented JSON for human reading (keys come out sorted)
// in:  config = {"region": "eu-west-1", "debug": false}
// tpl: {{ config | json:'pretty' }}
// out: {
// out:   "debug": false,
// out:   "region": "eu-west-1"
// out: }
//
// example: a slice becomes a JSON array
// in:  ids = [1, 2, 3]
// tpl: {{ ids | json }}
// out: [1,2,3]
//
// example: nil serializes to the JSON null literal
// in:  missing = nil
// tpl: {{ missing | json }}
// out: null
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
