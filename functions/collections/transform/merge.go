package transform

import "github.com/toaweme/sintax/functions"

// ModifierNameMerge is the template name for the Merge modifier.
const ModifierNameMerge functions.ModifierName = "merge"

// Merge is an alias for Map, useful when the name "map" reads awkwardly or
// clashes with other syntax. It keys a slice of maps by the named field exactly
// as Map does.
func Merge(value []map[string]any, key string) (map[string]map[string]any, error) {
	return Map(value, key)
}
