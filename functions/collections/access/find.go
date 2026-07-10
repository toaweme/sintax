package access

import (
	"fmt"
	"reflect"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameFind is the template name for the Find modifier.
const ModifierNameFind functions.ModifierName = "find"

// FindSlice returns the first map in a slice whose key field equals keyValue,
// scanning in order and returning the whole matching map. Matching is exact on
// value and type, so a field holding the integer 42 is not matched by the string
// "42". When nothing matches it returns a non-fatal ErrAllowsDefaultFunc error so
// the default modifier can supply a fallback.
func FindSlice(v []any, key string, keyValue any) (any, error) {
	for _, elem := range v {
		if m, ok := elem.(map[string]any); ok {
			if val, ok := m[key]; ok && reflect.DeepEqual(val, keyValue) {
				return m, nil
			}
		}
	}
	return nil, fmt.Errorf("%w: key %q with value %v not found in slice", functions.ErrAllowsDefaultFunc, key, keyValue)
}

// FindMap returns the map itself when its key field equals keyValue, and
// otherwise a non-fatal ErrAllowsDefaultFunc error. It is the single-map form of
// find.
func FindMap(v map[string]any, key string, keyValue any) (any, error) {
	if val, ok := v[key]; ok && reflect.DeepEqual(val, keyValue) {
		return v, nil
	}
	return nil, fmt.Errorf("%w: key %q with value %v not found in map", functions.ErrAllowsDefaultFunc, key, keyValue)
}
