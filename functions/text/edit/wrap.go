package edit

import "github.com/toaweme/sintax/functions"

// ModifierNameWrap is the template name for the Wrap modifier.
const ModifierNameWrap functions.ModifierName = "wrap"

// Wrap nests the value inside a new single-entry map under the given key.
// The result is a map value, not a JSON string, so it is meant to be piped into an
// encoder such as json or yaml, or returned as structured data. A non-string key
// becomes the empty-string key rather than an error.
func Wrap(value any, key any) (map[string]any, error) {
	k, _ := key.(string)
	return map[string]any{k: value}, nil
}
