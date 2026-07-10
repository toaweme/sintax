// Package transform provides modifiers that reshape collections.
package transform

import (
	"fmt"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameMap is the template name for the Map modifier.
const ModifierNameMap functions.ModifierName = "map"

// Map converts a slice of string-keyed maps into a single map keyed by the named
// field's value, turning a list you have to scan into a lookup table. The field
// must hold a string, which becomes the key; elements missing that field are
// skipped, and when two elements share a key value the later one wins.
func Map(value []map[string]any, key string) (map[string]map[string]any, error) {
	out := make(map[string]map[string]any)
	for _, m := range value {
		v, ok := m[key]
		if !ok {
			continue
		}
		s, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("map expected a string value at key %q, got %T", key, v)
		}
		out[s] = m
	}
	return out, nil
}
