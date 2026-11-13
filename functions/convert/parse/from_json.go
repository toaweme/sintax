package parse

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameFromJSON is the template name for the FromJSON modifier.
const ModifierNameFromJSON functions.ModifierName = "from_json"

// FromJSON parses a JSON string into the value it holds, so a serialized
// payload (an API response body, a config blob) becomes data that later
// template steps can index into. An object becomes a map[string]any, an array
// a []any, and a string, bool or null its Go counterpart.
//
// Numbers decode to native int64 or float64 (a value with a decimal point or
// exponent becomes float64, otherwise int64) rather than json.Number, so
// downstream numeric modifiers see real numbers.
func FromJSON(value string) (any, error) {
	dec := json.NewDecoder(strings.NewReader(value))
	dec.UseNumber()

	var raw any
	if err := dec.Decode(&raw); err != nil {
		return nil, fmt.Errorf("failed to parse JSON value: %w", err)
	}

	return functions.ConvertNumbersJSON(raw), nil
}
