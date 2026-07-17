package transform

import "github.com/toaweme/sintax/functions"

// ModifierNameMerge is the template name for the Merge modifier.
const ModifierNameMerge functions.ModifierName = "merge"

// Merge keys a slice of maps by the named field, producing a lookup map. It is
// an alias for the map modifier, useful when the name "map" reads awkwardly or
// clashes with other syntax.
func Merge(value []map[string]any, key string) (map[string]map[string]any, error) {
	return Map(value, key)
}
