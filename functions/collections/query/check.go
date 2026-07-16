// Package query provides modifiers that test and filter collections.
package query

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameHas is the template name for the Has modifier.
const ModifierNameHas functions.ModifierName = "has"

// ModifierNameIs is the template name for the Is modifier.
const ModifierNameIs functions.ModifierName = "is"

// Has reports whether a collection contains something, and what "contains"
// means depends on the shape of the value.
//
// For a plain slice, a single parameter is the element to look for, and Has is
// true when any element equals it. For a slice of maps, the first parameter is a
// field key and the rest are candidate values, so Has is true when any item's
// field equals any of those values. For a map with one parameter, Has tests only
// whether the key exists (the stored value is ignored, so a key mapped to false
// still counts as present). For a map with a key plus one or more values, Has is
// true when the key exists and its value equals any of the given values.
//
// Matching is exact on type, so an integer element is not found by a string
// parameter of the same digits.
func Has(value any, params []any) (any, error) {
	if len(params) == 0 {
		return false, errors.New("`has` requires at least one parameter")
	}

	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		return hasInSlice(v, params)
	case reflect.Map:
		return hasInMap(v, params)
	default:
		return false, fmt.Errorf("expected slice, array, or map, got %T", value)
	}
}

// Is reports whether the value equals any one of the given candidates, a compact
// way to write an "is this one of these" test in a template. Comparison is exact
// on type, so the number 5 does not match the string "5".
func Is(value any, candidates ...any) (bool, error) {
	if len(candidates) == 0 {
		return false, errors.New("is requires at least one candidate")
	}

	for _, candidate := range candidates {
		if reflect.DeepEqual(value, candidate) {
			return true, nil
		}
	}

	return false, nil
}

func hasInSlice(v reflect.Value, params []any) (bool, error) {
	if v.Len() > 0 {
		firstElem := v.Index(0)
		if firstElem.Kind() == reflect.Interface {
			firstElem = firstElem.Elem()
		}

		if firstElem.Kind() == reflect.Map {
			if len(params) < 2 {
				return false, errors.New("has requires key and value parameters for slice of maps")
			}

			key, err := functions.ParamString(params, 0)
			if err != nil {
				return false, fmt.Errorf("first parameter must be a string key: %w", err)
			}

			for i := 1; i < len(params); i++ {
				searchValue, err := functions.ParamAny(params, i)
				if err != nil {
					return false, fmt.Errorf("invalid parameter at index %d: %w", i, err)
				}

				for j := range v.Len() {
					elem := v.Index(j)
					if elem.Kind() == reflect.Interface {
						elem = elem.Elem()
					}

					if elem.Kind() == reflect.Map {
						mapKey := reflect.ValueOf(key)
						if elem.Type().Key().Kind() == reflect.String {
							if val := elem.MapIndex(mapKey); val.IsValid() {
								if reflect.DeepEqual(val.Interface(), searchValue) {
									return true, nil
								}
							}
						}
					}
				}
			}
			return false, nil
		}
	}

	for _, param := range params {
		for i := range v.Len() {
			elem := v.Index(i)
			if elem.Kind() == reflect.Interface {
				elem = elem.Elem()
			}
			if reflect.DeepEqual(elem.Interface(), param) {
				return true, nil
			}
		}
	}

	return false, nil
}

func hasInMap(v reflect.Value, params []any) (bool, error) {
	if len(params) == 1 {
		key := reflect.ValueOf(params[0])
		if v.MapIndex(key).IsValid() {
			return true, nil
		}
		return false, nil
	}

	if len(params) < 2 {
		return false, errors.New("has requires at least one parameter for maps")
	}

	key := reflect.ValueOf(params[0])
	mapValue := v.MapIndex(key)

	if !mapValue.IsValid() {
		return false, nil
	}

	for i := 1; i < len(params); i++ {
		if reflect.DeepEqual(mapValue.Interface(), params[i]) {
			return true, nil
		}
	}

	return false, nil
}
