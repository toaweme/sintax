package access

import (
	"fmt"
	"reflect"

	"github.com/toaweme/sintax/functions"
)

// ModifierNamePluck is the template name for the Pluck modifier.
const ModifierNamePluck functions.ModifierName = "pluck"

// Pluck reads one named field from every element of a slice of maps and returns
// the collected values as a slice, in order. Every element must be a map that
// has the field: a missing field, a non-map element, or a nil element is an
// error rather than a skipped or padded entry, so the result length always
// matches the input length. An empty slice yields an empty slice.
func Pluck(value []any, field string) ([]any, error) {
	out := make([]any, 0, len(value))
	for i, elem := range value {
		ev := reflect.ValueOf(elem)
		for ev.Kind() == reflect.Pointer || ev.Kind() == reflect.Interface {
			if ev.IsNil() {
				return nil, fmt.Errorf("pluck: element %d is nil", i)
			}
			ev = ev.Elem()
		}
		if ev.Kind() != reflect.Map {
			return nil, fmt.Errorf("pluck: element %d is not a map (%T)", i, elem)
		}
		var found bool
		for _, k := range ev.MapKeys() {
			if fmt.Sprint(k.Interface()) == field {
				out = append(out, ev.MapIndex(k).Interface())
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("pluck: field %q not found in element %d", field, i)
		}
	}
	return out, nil
}
