package access

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameFind is the template name for the Find modifier.
const ModifierNameFind functions.ModifierName = "find"

// Find returns the first map in a slice whose named field equals a given value.
// It scans the slice in order and returns the whole matching map, not just the
// field. The value can also be a single map, in which case find returns that
// map when its field matches. Matching is exact on both value and type, so the
// field's stored type must equal the type you pass: a field holding the integer
// 42 is not matched by the string "42", and vice versa. When nothing matches,
// find returns a non-fatal error so the default modifier can supply a fallback
// instead of failing the render.
//
// value: array, map
// param:0: string (field name)
// param:1: any (value to match, compared by exact type and value)
// returns: map
//
// example: look up a user by id
// in:  users = [{"id": 7, "name": "Bob"}, {"id": 42, "name": "Alice"}]
// tpl: {{ users | find:'id',42 }}
// out: {"id": 42, "name": "Alice"}
//
// example: match on a string field
// in:  products = [{"slug": "hat", "price": 25}, {"slug": "mug", "price": 12}]
// tpl: {{ products | find:'slug','mug' }}
// out: {"slug": "mug", "price": 12}
//
// example: fall back when not found
// in:  items = [{"slug": "hat", "price": 25}]
// tpl: {{ items | find:'slug','scarf' | default:{} }}
// out: {}
func Find(value any, params []any) (any, error) {
	if len(params) < 2 {
		return nil, errors.New("find requires at least two parameters: key and value")
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
		for i := range v.Len() {
			elem := v.Index(i)
			if elem.Kind() == reflect.Interface {
				elem = elem.Elem()
			}

			if elem.Kind() == reflect.Map {
				if mapValue, ok := elem.Interface().(map[string]any); ok {
					if val, ok := mapValue[key]; ok && reflect.DeepEqual(val, keyValue) {
						return mapValue, nil
					}
				}
			}
		}
		return nil, fmt.Errorf("%w: key %q with value %q not found in slice", functions.ErrAllowsDefaultFunc, key, keyValue)

	case reflect.Map:
		mapValue, ok := value.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("expected map[string]any, got %T", value)
		}
		if val, ok := mapValue[key]; ok && reflect.DeepEqual(val, keyValue) {
			return mapValue, nil
		}
		return nil, fmt.Errorf("%w: key %q with value %q not found in map", functions.ErrAllowsDefaultFunc, key, keyValue)

	default:
		return nil, fmt.Errorf("expected slice of maps or map, got %T", value)
	}
}
