package format

import (
	"fmt"
	"reflect"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameLength is the template name for the Length modifier.
const ModifierNameLength functions.ModifierName = "length"

// Length returns the size of a value: the number of bytes in a string or byte
// slice, or the number of elements in a slice, array, or map. For strings the
// count is UTF-8 bytes, not runes, so a multi-byte character such as "é" counts
// as more than one. A nil pointer or nil interface counts as zero, but a bare
// nil (no type) has no length and returns an error.
//
// value: string, bytes, array, map
// returns: int
//
// example: count bytes in an ASCII name
// in:  name = "Alice"
// tpl: {{ name | length }}
// out: 5
//
// example: count items in a cart
// in:  items = ["mug", "pen", "pad"]
// tpl: {{ items | length }}
// out: 3
//
// example: a multi-byte character counts as its bytes, not one rune
// in:  word = "café"
// tpl: {{ word | length }}
// out: 5
var Length = func(value any, _ []any) (any, error) {
	switch v := value.(type) {
	case string:
		return len(v), nil
	case []byte:
		return len(v), nil
	}

	rv := reflect.ValueOf(value)
	for rv.Kind() == reflect.Pointer || rv.Kind() == reflect.Interface {
		if rv.IsNil() {
			return 0, nil
		}
		rv = rv.Elem()
	}
	switch rv.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		return rv.Len(), nil
	default:
	}

	return nil, fmt.Errorf("length function expected string, bytes, slice, array, or map, got %T", value)
}
