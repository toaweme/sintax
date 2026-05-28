package collections

import (
	"fmt"
	"reflect"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameFind is the template name for the Find modifier.
const ModifierNameFind functions.ModifierName = "find"

// Find returns the first element in a slice or map where a field equals the given value.
// Returns a non-fatal error when not found, allowing the default modifier to handle it.
//
// value: array, map
// param:0: string
// param:1: any
// returns: map
//
// example: look up a user by id
// in:  users = [{"id": 7, "name": "Bob"}, {"id": 42, "name": "Alice"}]
// tpl: {{ users | find:'id',42 }}
// out: {"id": 42, "name": "Alice"}
//
// example: fall back when not found
// in:  items = [{"slug": "hat", "price": 25}]
// tpl: {{ items | find:'slug','scarf' | default:{} }}
// out: {}
func Find(value any, params []any) (any, error) {
	if len(params) < 2 {
		return nil, fmt.Errorf("find requires at least two parameters: key and value")
	}

	key, err := functions.ParamString(params, 0)
	if err != nil {
		return nil, err
	}

	keyValue, err := functions.ParamAny(params, 1)
	if err != nil {
		return nil, err
	}

	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			elem := v.Index(i)
			if elem.Kind() == reflect.Interface {
				elem = elem.Elem()
			}

			if elem.Kind() == reflect.Map {
				mapValue := elem.Interface().(map[string]any)
				if val, ok := mapValue[key]; ok && reflect.DeepEqual(val, keyValue) {
					return mapValue, nil
				}
			}
		}
		return nil, fmt.Errorf("%w: key %q with value %q not found in slice", functions.ErrAllowsDefaultFunc, key, keyValue)

	case reflect.Map:
		mapValue := value.(map[string]any)
		if val, ok := mapValue[key]; ok && reflect.DeepEqual(val, keyValue) {
			return mapValue, nil
		}
		return nil, fmt.Errorf("%w: key %q with value %q not found in map", functions.ErrAllowsDefaultFunc, key, keyValue)
	}

	return nil, fmt.Errorf("expected slice of maps or map, got %T", value)
}
