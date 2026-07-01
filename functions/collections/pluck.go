package collections

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/toaweme/sintax/functions"
)

// ModifierNamePluck is the template name for the Pluck modifier.
const ModifierNamePluck functions.ModifierName = "pluck"

// Pluck extracts a single field from each element of a slice of maps and
// returns a slice of values.
//
// value: array
// param:0: string
// returns: array
//
// example: collect every user id
// in:  users = [{"id": 1, "name": "Alice"}, {"id": 2, "name": "Bob"}]
// tpl: {{ users | pluck:'id' }}
// out: [1, 2]
//
// example: gather product names
// in:  products = [{"name": "Mug", "price": 12}, {"name": "Pen", "price": 3}]
// tpl: {{ products | pluck:'name' }}
// out: ["Mug", "Pen"]
func Pluck(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, errors.New("pluck: missing field name")
	}
	field, ok := params[0].(string)
	if !ok {
		return nil, fmt.Errorf("pluck: field name must be a string, got %T", params[0])
	}

	rv := reflect.ValueOf(value)
	for rv.Kind() == reflect.Pointer || rv.Kind() == reflect.Interface {
		if rv.IsNil() {
			return []any{}, nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return nil, fmt.Errorf("pluck: expected slice or array, got %T", value)
	}

	out := make([]any, 0, rv.Len())
	for i := range rv.Len() {
		elem := rv.Index(i).Interface()
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
