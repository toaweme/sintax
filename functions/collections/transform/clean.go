package transform

import (
	"reflect"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameClean is the template name for the Clean modifier.
const ModifierNameClean functions.ModifierName = "clean"

// Clean drops the elements that hold nothing, so a list can be built by
// appending unconditionally and tidied once at the end instead of being guarded
// at every append. Order is kept and the remaining elements are untouched.
//
// Nothing means nil, an empty string, and an empty list or object. Zero and
// false are real values and stay, which is the same line default draws between
// absent data and a value that happens to be empty.
func Clean(value []any) ([]any, error) {
	out := make([]any, 0, len(value))
	for _, elem := range value {
		if isNothing(elem) {
			continue
		}
		out = append(out, elem)
	}
	return out, nil
}

// isNothing reports whether an element holds nothing, unwrapping interfaces and
// pointers first so a nil behind either one is recognized as the absence it is.
func isNothing(elem any) bool {
	if elem == nil {
		return true
	}
	rv := reflect.ValueOf(elem)
	for rv.Kind() == reflect.Pointer || rv.Kind() == reflect.Interface {
		if rv.IsNil() {
			return true
		}
		rv = rv.Elem()
	}
	switch rv.Kind() {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map:
		return rv.Len() == 0
	default:
		return false
	}
}

// cleanNil answers a value that is not there at all with a miss, so
// `{{ maybe | clean | default:[] }}` falls back rather than reporting that nil
// is the wrong type to clean. It declines any other value so Overload falls
// through to the typed clause.
func cleanNil(value any, _ []any) (any, error) {
	if value == nil {
		return nil, functions.Miss("clean found no list to clean")
	}
	return nil, functions.ErrInvalidValueType
}
