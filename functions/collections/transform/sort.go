package transform

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameSort is the template name for the Sort modifier.
const ModifierNameSort functions.ModifierName = "sort"

// SortAsc sorts a slice ascending, the default direction when none is given. It
// sorts a copy, so the caller's slice is never mutated (coerce may hand this
// clause the caller's own backing array by reference).
func SortAsc(slice []any) ([]any, error) {
	out := append([]any{}, slice...)
	sortSlice(out, true)
	return out, nil
}

// SortDir sorts a copy of a slice in the named direction, 'asc' or 'desc'; any
// other direction is an error. An empty slice is returned untouched without
// validating the direction, matching the original short-circuit.
func SortDir(slice []any, direction string) ([]any, error) {
	if len(slice) == 0 {
		return slice, nil
	}
	out := append([]any{}, slice...)
	switch direction {
	case "asc":
		sortSlice(out, true)
	case "desc":
		sortSlice(out, false)
	default:
		return nil, fmt.Errorf("sort expected direction 'asc' or 'desc', got %q", direction)
	}
	return out, nil
}

// sortNil is the sort clause that passes a nil value straight through as nil, and
// declines any other value so Overload falls through to the typed clauses.
func sortNil(value any, _ []any) (any, error) {
	if value == nil {
		return nil, nil //nolint:nilnil // deliberate, nil input passes through as nil, not an error
	}
	return nil, functions.ErrInvalidValueType
}

// sortSlice orders slice in place. It compares strings alphabetically, numbers
// numerically, and booleans with false before true. A mixed-type slice groups
// elements by type name first so the result stays deterministic, and nils sort
// first regardless of direction.
func sortSlice(slice []any, ascending bool) {
	sort.Slice(slice, func(i, j int) bool {
		vi := reflect.ValueOf(slice[i])
		vj := reflect.ValueOf(slice[j])

		if !vi.IsValid() && !vj.IsValid() {
			return false
		}
		if !vi.IsValid() {
			return true
		}
		if !vj.IsValid() {
			return false
		}

		ti := vi.Type()
		tj := vj.Type()
		if ti != tj {
			result := ti.String() < tj.String()
			if ascending {
				return result
			}
			return !result
		}

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
			less = !vi.Bool() && vj.Bool()
		default:
			less = fmt.Sprintf("%v", vi.Interface()) < fmt.Sprintf("%v", vj.Interface())
		}

		if ascending {
			return less
		}
		return !less
	})
}
