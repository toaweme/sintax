package transform

import (
	"fmt"
	"reflect"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameFlatten is the template name for the Flatten modifier.
const ModifierNameFlatten functions.ModifierName = "flatten"

// Flatten flattens a slice of slices by exactly one level. Any element that is
// itself a slice or array is spread into the result; every other element is
// copied through unchanged. Only one level is removed, so a slice nested two
// deep still comes out as a slice. An untyped nil element is dropped, and a nil
// value overall is an error rather than an empty result.
//
// value: array
// returns: array
//
// example: combine items from multiple groups
// in:  groups = [{"name": "Drinks", "items": ["coffee", "tea"]}, {"name": "Snacks", "items": ["cookie", "muffin"]}]
// tpl: {{ groups | pluck:'items' | flatten }}
// out: ["coffee", "tea", "cookie", "muffin"]
//
// example: merge a list of lists
// in:  weeks = [["Mon", "Tue"], ["Wed", "Thu"]]
// tpl: {{ weeks | flatten }}
// out: ["Mon", "Tue", "Wed", "Thu"]
func Flatten(value any, params []any) (any, error) {
	rv := reflect.ValueOf(value)
	for rv.Kind() == reflect.Pointer || rv.Kind() == reflect.Interface {
		if rv.IsNil() {
			return []any{}, nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return nil, fmt.Errorf("flatten: expected slice or array, got %T", value)
	}

	out := make([]any, 0, rv.Len())
	for i := range rv.Len() {
		elem := rv.Index(i).Interface()
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
