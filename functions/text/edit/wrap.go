package edit

import "github.com/toaweme/sintax/functions"

// ModifierNameWrap is the template name for the Wrap modifier.
const ModifierNameWrap functions.ModifierName = "wrap"

// Wrap nests value inside a new single-entry map under key. The result is a map
// value, not a JSON string, so it is meant to be piped into an encoder or
// returned as structured data. key is typed as any because a non-string key
// coerces to the empty string rather than erroring, preserving the original
// behaviour.
func Wrap(value any, key any) (map[string]any, error) {
	k, _ := key.(string)
	return map[string]any{k: value}, nil
}
