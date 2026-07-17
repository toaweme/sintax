package transform

import (
	"reflect"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameFlatten is the template name for the Flatten modifier.
const ModifierNameFlatten functions.ModifierName = "flatten"

// Flatten flattens a slice by exactly one level: any element that is itself a
// slice or array is spread into the result, and every other element is copied
// through unchanged. A nested slice two levels deep still comes out as a slice,
// and an untyped nil element is dropped.
func Flatten(v []any) ([]any, error) {
	out := make([]any, 0, len(v))
	for _, elem := range v {
		ev := reflect.ValueOf(elem)
		for ev.Kind() == reflect.Pointer || ev.Kind() == reflect.Interface {
			if ev.IsNil() {
				out = append(out, nil)
				ev = reflect.Value{}
				break
			}
			ev = ev.Elem()
		}
		if !ev.IsValid() {
			continue
		}
		if ev.Kind() == reflect.Slice || ev.Kind() == reflect.Array {
			for j := range ev.Len() {
				out = append(out, ev.Index(j).Interface())
			}
			continue
		}
		out = append(out, elem)
	}
	return out, nil
}
