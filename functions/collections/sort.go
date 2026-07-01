package collections

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameSort is the template name for the Sort modifier.
const ModifierNameSort functions.ModifierName = "sort"

// Sort sorts a slice in ascending or descending order.
// Handles strings, numbers, and booleans. Elements of mixed types are grouped by type name.
//
// value: array
// param:0?: string
// returns: array
//
// example: sort names alphabetically
// in:  names = ["Charlie", "Alice", "Bob"]
// tpl: {{ names | sort }}
// out: ["Alice", "Bob", "Charlie"]
//
// example: sort scores from highest to lowest
// in:  scores = [72, 95, 88]
// tpl: {{ scores | sort:'desc' }}
// out: [95, 88, 72]
//
// example: sort prices ascending
// in:  prices = [9.99, 4.50, 14.00]
// tpl: {{ prices | sort:'asc' }}
// out: [4.50, 9.99, 14.00]
func Sort(value any, params []any) (any, error) {
	if value == nil {
		return nil, nil //nolint:nilnil // deliberate: nil input passes through as nil, not an error
	}

	slice, err := functions.ValueSlice(value)
	if err != nil {
		return nil, err
	}

	if len(slice) == 0 {
		return slice, nil
	}

	direction, _ := functions.ParamString(params, 0)
	if direction == "" {
		direction = "asc"
	}
	if direction != "asc" && direction != "desc" {
		return nil, fmt.Errorf("sort: invalid direction %s, expected 'asc' or 'desc'", direction)
	}

	ascending := direction == "asc"

	sort.Slice(slice, func(i, j int) bool {
		vi := reflect.ValueOf(slice[i])
		vj := reflect.ValueOf(slice[j])

		// handle nil values - nils always come first (regardless of direction)
		if !vi.IsValid() && !vj.IsValid() {
			return false
		}
		if !vi.IsValid() {
			return true
		}
		if !vj.IsValid() {
			return false
		}

		// get the actual types
		ti := vi.Type()
		tj := vj.Type()

		// if types don't match, compare by type string for consistent ordering
		if ti != tj {
			result := ti.String() < tj.String()
			if ascending {
				return result
			}
			return !result
		}

		// handle different types
		var less bool
		switch vi.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			less = vi.Int() < vj.Int()

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			less = vi.Uint() < vj.Uint()

		case reflect.Float32, reflect.Float64:
			less = vi.Float() < vj.Float()

		case reflect.String:
			less = vi.String() < vj.String()

		case reflect.Bool:
			// false < true
			less = !vi.Bool() && vj.Bool()

		default:
			// for other types, convert to string and compare
			less = fmt.Sprintf("%v", vi.Interface()) < fmt.Sprintf("%v", vj.Interface())
		}

		if ascending {
			return less
		}
		return !less
	})

	return slice, nil
}
