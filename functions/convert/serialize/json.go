// Package serialize provides modifiers that render a value to a data format.
package serialize

import (
	"encoding/json"
	"fmt"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameJSON is the template name for the JSON modifier.
const ModifierNameJSON functions.ModifierName = "json"

// jsonModePretty selects indented output in JSONMode.
const jsonModePretty = "pretty"

// JSON serializes value to a compact JSON string via Go's encoding/json. Object
// keys always come out sorted alphabetically, so the output is stable regardless
// of the input map's insertion order. It is the no-param json clause.
func JSON(value any) (string, error) {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return "", fmt.Errorf("failed to apply json filter: %w", err)
	}
	return string(jsonBytes), nil
}

// JSONMode is the one-param json clause. The literal 'pretty' selects indented
// output using two spaces per level. Any other mode falls back to the compact
// form, matching the plain JSON clause.
func JSONMode(value any, mode string) (string, error) {
	if mode != jsonModePretty {
		return JSON(value)
	}
	jsonBytes, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to apply json filter: %w", err)
	}
	return string(jsonBytes), nil
}
