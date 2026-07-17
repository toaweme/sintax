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
// column of a list of records. A field absent from a record is a miss, so
// `| sum:'amount' | default:0` falls back rather than failing, while a
// non-numeric value in the column is a terminal error, since silently treating
// "abc" as zero would understate a total that someone is going to act on.
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

// numberFromField reads one numeric column out of a record. A nil record
// contributes zero rather than reporting a miss, unlike pluck, which reports one
// for the same input. The operations differ in what absence means. Addition has
// an identity to fall back on, so a nil record can be counted as contributing
// nothing without inventing anything, while pluck would have to fabricate an
// element to keep its result aligned with its input.
func numberFromField(elem any, field string) (float64, error) {
	rv := reflect.ValueOf(elem)
	for rv.Kind() == reflect.Pointer || rv.Kind() == reflect.Interface {
		if rv.IsNil() {
			return 0, nil
		}
		rv = rv.Elem()
	}
	if elem == nil {
		return 0, nil
	}
	if rv.Kind() != reflect.Map {
		return 0, fmt.Errorf("sum expected a map to read field %q from, got %T: %w", field, elem, functions.ErrInvalidValueType)
	}
	for _, k := range rv.MapKeys() {
		if fmt.Sprint(k.Interface()) == field {
			return functions.ParseNumber(rv.MapIndex(k).Interface())
		}
	}
	return 0, functions.Miss("sum found no field %q to total", field)
}
