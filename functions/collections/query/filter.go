package query

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/toaweme/sintax/functions"
	"github.com/toaweme/sintax/functions/collections/access"
)

// ModifierNameFilter is the template name for the Filter modifier.
const ModifierNameFilter functions.ModifierName = "filter"

// Filter returns the items of a slice whose named field equals a search value.
// Each item is expected to be a map, and the field is looked up by key. Use dot
// notation to reach into a nested map (for example "meta.published"). Numbers
// compare by value across integer and float types, so 10 matches 10.0. An item
// is dropped when the field is missing or the values differ, and a slice where
// nothing matches comes back empty rather than as an error.
//
// value: array
// param:0: string (the field key, with optional dot notation for nested maps)
// param:1: any (the value the field must equal)
// returns: array
//
// example: keep only active items
// in:  items = [{"name": "Coffee", "status": "active"}, {"name": "Tea", "status": "sold-out"}]
// tpl: {{ items | filter:'status','active' }}
// out: [{"name": "Coffee", "status": "active"}]
//
// example: keep only admin users
// in:  users = [{"name": "Alice", "role": "admin"}, {"name": "Bob", "role": "viewer"}]
// tpl: {{ users | filter:'role','admin' }}
// out: [{"name": "Alice", "role": "admin"}]
//
// example: filter on a nested field
// in:  posts = [{"title": "Hello", "meta": {"published": true}}, {"title": "Draft", "meta": {"published": false}}]
// tpl: {{ posts | filter:'meta.published',true }}
// out: [{"title": "Hello", "meta": {"published": true}}]
//
// example: numbers match by value across int and float
// in:  products = [{"name": "Mug", "price": 10}, {"name": "Cup", "price": 8}]
// tpl: {{ products | filter:'price',10 }}
// out: [{"name": "Mug", "price": 10}]
//
// example: nothing matches, so the result is empty
// in:  users = [{"name": "Alice", "role": "admin"}, {"name": "Bob", "role": "viewer"}]
// tpl: {{ users | filter:'role','owner' }}
// out: []
func Filter(value any, params []any) (any, error) {
	slice, err := functions.ValueSlice(value) // []any
	if err != nil {
		return nil, fmt.Errorf("filter function: %w", err)
	}

	if len(params) < 2 {
		return nil, errors.New("filter function requires at least 2 parameters: key and search value")
	}

	key, err := functions.ParamString(params, 0) // 'some_key' or 'some_parent.child_key'
	if err != nil {
		return nil, fmt.Errorf("filter function: %w", err)
	}

	search, err := functions.ParamAny(params, 1) // 'some_value', 10, true, etc.
	if err != nil {
		return nil, fmt.Errorf("filter function: %w", err)
	}

	var filtered []any
	keyParams := []any{key} // Prepare params for key function

	for _, item := range slice {
		// Use the existing key function to extract the nested value
		extractedValue, err := access.Key(item, keyParams)
		if err != nil {
			// Skip items where key extraction fails (key doesn't exist)
			continue
		}

		// Compare the extracted value with the search value
		if valuesEqual(extractedValue, search) {
			filtered = append(filtered, item)
		}
	}

	return filtered, nil
}

// valuesEqual compares two values for equality, handling different types
func valuesEqual(a, b any) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	// Direct equality check first
	if a == b {
		return true
	}

	// Use reflection for deeper comparison
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)

	// If types are exactly the same, use reflect.DeepEqual
	if va.Type() == vb.Type() {
		return reflect.DeepEqual(a, b)
	}

	// Handle numeric type conversions
	if isNumeric(va) && isNumeric(vb) {
		return compareNumeric(va, vb)
	}

	// Handle string comparisons
	if va.Kind() == reflect.String && vb.Kind() == reflect.String {
		return va.String() == vb.String()
	}

	// Handle boolean comparisons
	if va.Kind() == reflect.Bool && vb.Kind() == reflect.Bool {
		return va.Bool() == vb.Bool()
	}

	return false
}

// isNumeric checks if a reflect.Value represents a numeric type
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

// compareNumeric compares two numeric values by converting them to float64
func compareNumeric(a, b reflect.Value) bool {
	aFloat := convertToFloat64(a)
	bFloat := convertToFloat64(b)
	return aFloat == bFloat
}

// convertToFloat64 converts a numeric reflect.Value to float64
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
