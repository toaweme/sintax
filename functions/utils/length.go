package utils

import (
	"fmt"
	"reflect"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameLength is the template name for the Length modifier.
const ModifierNameLength functions.ModifierName = "length"

// Length returns the number of characters in a string, bytes in a byte slice,
// or elements in a slice/array/map.
//
// value: string, bytes, array, map
// returns: int
//
// example: count characters in a name
// in:  name = "Alice"
// tpl: {{ name | length }}
// out: 5
//
// example: count items in a cart
// in:  items = ["mug", "pen", "pad"]
// tpl: {{ items | length }}
// out: 3
var Length = func(value any, _ []any) (any, error) {
	switch v := value.(type) {
	case string:
		return len(v), nil
	case []byte:
		return len(v), nil
	}

	rv := reflect.ValueOf(value)
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		if rv.IsNil() {
			return 0, nil
		}
		rv = rv.Elem()
	}
	switch rv.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		return rv.Len(), nil
	}

	return nil, fmt.Errorf("length function expected string, bytes, slice, array, or map, got %T", value)
}
