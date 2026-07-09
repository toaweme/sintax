// Package transform provides modifiers that reshape collections.
package transform

import (
	"errors"
	"fmt"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameMap is the template name for the Map modifier.
const ModifierNameMap functions.ModifierName = "map"

// Map converts a slice of maps into a map keyed by the given field's string value.
//
// value: array
// param:0: string
// returns: map
//
// example: index a list of users by id
// in:  users = [{"id": "u1", "name": "Alice"}, {"id": "u2", "name": "Bob"}]
// tpl: {{ users | map:'id' }}
// out: {"u1": {"id": "u1", "name": "Alice"}, "u2": {"id": "u2", "name": "Bob"}}
//
// example: index categories by slug
// in:  tags = [{"slug": "coffee", "label": "Coffee"}, {"slug": "tea", "label": "Tea"}]
// tpl: {{ tags | map:'slug' }}
// out: {"coffee": {"slug": "coffee", "label": "Coffee"}, "tea": {"slug": "tea", "label": "Tea"}}
func Map(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, errors.New("in requires at least one parameter")
	}

	key, err := functions.ParamString(params, 0)
	if err != nil {
		return nil, err
	}

	slice, ok := value.([]map[string]any)
	if !ok {
		return nil, fmt.Errorf("expected slice of maps, got %T", value)
	}

	n := make(map[string]map[string]any)
	for _, m := range slice {
		v, ok := m[key]
		if !ok {
			continue
		}
		s, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("expected string value in map at key %q, got %T", key, v)
		}
		n[s] = m
	}
	return n, nil
}
