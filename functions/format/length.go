package format

import (
	"fmt"
	"reflect"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameLength is the template name for the Length modifier.
const ModifierNameLength functions.ModifierName = "length"

// LengthString returns the number of UTF-8 bytes in a string, so a multi-byte
// character such as "é" counts as more than one.
func LengthString(s string) (int, error) {
	return len(s), nil
}

// LengthBytes returns the number of bytes in a byte slice.
func LengthBytes(b []byte) (int, error) {
	return len(b), nil
}

// LengthReflect counts the elements of a slice, array, or map via reflection,
// dereferencing a pointer or interface first (a nil pointer counts as zero). It
// is the fallback clause; a value with no meaningful length, including a bare
// nil, is an error.
func LengthReflect(value any) (int, error) {
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
		return 0, fmt.Errorf("length expected a string, bytes, slice, array, or map, got %T", value)
	}
}
