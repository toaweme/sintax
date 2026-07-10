package query

import (
	"reflect"

	"github.com/toaweme/sintax/functions"
	"github.com/toaweme/sintax/functions/collections/access"
)

// ModifierNameFilter is the template name for the Filter modifier.
const ModifierNameFilter functions.ModifierName = "filter"

// Filter returns the items of a slice whose named field equals search. Each item
// is looked up by key, with dot notation reaching into a nested map (for example
// "meta.published"). Numbers compare by value across the int and float kinds, so
// 10 matches 10.0. An item is dropped when the field is missing or the values
// differ, and a slice where nothing matches comes back empty rather than as an
// error.
func Filter(value []any, key string, search any) ([]any, error) {
	var filtered []any
	keyParams := []any{key}
	for _, item := range value {
		// access.Key never returns an error - a missing key yields nil - so the
		// missing-field case simply compares nil against search and does not match.
		extracted, _ := access.Key(item, keyParams)
		if valuesEqual(extracted, search) {
			filtered = append(filtered, item)
		}
	}
	return filtered, nil
}

// valuesEqual compares two values for equality, treating the numeric kinds as
// interchangeable so an int field matches a float search of the same value.
func valuesEqual(a, b any) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a == b {
		return true
	}

	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)
	if va.Type() == vb.Type() {
		return reflect.DeepEqual(a, b)
	}
	if isNumeric(va) && isNumeric(vb) {
		return convertToFloat64(va) == convertToFloat64(vb)
	}
	if va.Kind() == reflect.String && vb.Kind() == reflect.String {
		return va.String() == vb.String()
	}
	if va.Kind() == reflect.Bool && vb.Kind() == reflect.Bool {
		return va.Bool() == vb.Bool()
	}
	return false
}

func isNumeric(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

func convertToFloat64(v reflect.Value) float64 {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(v.Uint())
	case reflect.Float32, reflect.Float64:
		return v.Float()
	default:
		return 0
	}
}
