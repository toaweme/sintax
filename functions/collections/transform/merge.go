package transform

import (
	"errors"
	"fmt"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameMerge is the template name for the Merge modifier.
const ModifierNameMerge functions.ModifierName = "merge"

// Merge converts a slice of maps into a map keyed by the given field's string value.
// Behaves identically to map; use as an alias when "map" conflicts with other syntax.
//
// value: array
// param:0: string
// returns: map
//
// example: index a list of users by id
// in:  users = [{"id": "u1", "name": "Alice"}, {"id": "u2", "name": "Bob"}]
// tpl: {{ users | merge:'id' }}
// out: {"u1": {"id": "u1", "name": "Alice"}, "u2": {"id": "u2", "name": "Bob"}}
//
// example: index records by name
// in:  records = [{"name": "draft", "value": 1}, {"name": "final", "value": 2}]
// tpl: {{ records | merge:'name' }}
// out: {"draft": {"name": "draft", "value": 1}, "final": {"name": "final", "value": 2}}
func Merge(value any, params []any) (any, error) {
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
