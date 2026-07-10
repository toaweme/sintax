package transform

import (
	"fmt"
	"reflect"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameSum is the template name for the Sum modifier.
const ModifierNameSum functions.ModifierName = "sum"

// SumElements adds up the elements of a slice, returning a float. Every element
// must be a number or a numeric string, and a nil element counts as zero.
func SumElements(v []any) (float64, error) {
	var total float64
	for i, elem := range v {
		n, err := functions.ParseNumber(elem)
		if err != nil {
			return 0, fmt.Errorf("failed to sum element %d: %w", i, err)
		}
		total += n
	}
	return total, nil
}

// SumField totals the named field across a slice of maps, the way you sum one
// column of a list of records. A missing field or a non-numeric value is an
// error.
func SumField(v []any, field string) (float64, error) {
	var total float64
	for i, elem := range v {
		n, err := numberFromField(elem, field)
		if err != nil {
			return 0, fmt.Errorf("failed to sum element %d: %w", i, err)
		}
		total += n
	}
	return total, nil
}

func numberFromField(elem any, field string) (float64, error) {
	rv := reflect.ValueOf(elem)
	for rv.Kind() == reflect.Pointer || rv.Kind() == reflect.Interface {
		if rv.IsNil() {
			return 0, nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Map {
		return 0, fmt.Errorf("expected a map for field %q, got %T", field, elem)
	}
	for _, k := range rv.MapKeys() {
		if fmt.Sprint(k.Interface()) == field {
			return functions.ParseNumber(rv.MapIndex(k).Interface())
		}
	}
	return 0, fmt.Errorf("field %q not found", field)
}
