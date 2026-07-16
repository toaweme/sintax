package parse

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameFromJSON is the template name for the FromJSON modifier.
const ModifierNameFromJSON functions.ModifierName = "from_json"

// FromJSON parses a JSON object string into a map, so a serialized payload (an
// API response body, a config blob) becomes data that later template steps can
// index into.
//
// Numbers decode to native int64 or float64 (a value with a decimal point or
// exponent becomes float64, otherwise int64) rather than json.Number, so
// downstream numeric modifiers see real numbers. A top-level JSON array or
// scalar is not an object and returns an error.
func FromJSON(value string) (map[string]any, error) {
	dec := json.NewDecoder(strings.NewReader(value))
	dec.UseNumber()

	var raw map[string]any
	if err := dec.Decode(&raw); err != nil {
		return nil, fmt.Errorf("failed to convert JSON to map: %w", err)
	}

	// ConvertNumbersJSON walks the value it is given, so a map goes in and the
	// same map comes back out with its json.Number leaves swapped for natives.
	converted, ok := functions.ConvertNumbersJSON(raw).(map[string]any)
	if !ok {
		return nil, fmt.Errorf("failed to convert JSON numbers to native types")
	}
	return converted, nil
}
