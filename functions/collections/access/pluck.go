package access

import (
	"fmt"
	"reflect"

	"github.com/toaweme/sintax/functions"
)

// ModifierNamePluck is the template name for the Pluck modifier.
const ModifierNamePluck functions.ModifierName = "pluck"

// Pluck reads one named field from every element of a slice of maps and returns
// the collected values as a slice, in order. The result length always matches
// the input length, so a field that is absent from an element is never skipped
// or padded over. An empty slice yields an empty slice.
//
// A field missing from an element, or an element holding nothing to read the
// field from, is a miss, so `| pluck:'key' | default:[]` falls back to an empty
// slice rather than failing. An element that is not a map at all is a terminal
// error, since plucking a field from a number is a template that cannot mean
// anything.
func Pluck(value []any, field string) ([]any, error) {
	out := make([]any, 0, len(value))
	for i, elem := range value {
		ev := reflect.ValueOf(elem)
		for ev.Kind() == reflect.Pointer || ev.Kind() == reflect.Interface {
			if ev.IsNil() {
				return nil, functions.Miss("pluck found nothing at element %d to read %q from", i, field)
			}
			ev = ev.Elem()
		}
		if elem == nil {
			return nil, functions.Miss("pluck found nothing at element %d to read %q from", i, field)
		}
		if ev.Kind() != reflect.Map {
			return nil, fmt.Errorf("pluck expected a map at element %d, got %T: %w", i, elem, functions.ErrInvalidValueType)
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
			return nil, functions.Miss("pluck found no field %q in element %d", field, i)
		}
	}
	return out, nil
}
