// Package collections provides template modifiers for slices and maps.
package collections

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

// Has returns true if the slice or map contains the given value.
// For a slice of maps, provide a key and value to match on a specific field.
//
// value: array, map
// param:0: any
// param:1?: any
// returns: bool
//
// example: check a list of tags
// in:  tags = ["featured", "sale", "new"]
// tpl: {{ tags | has:'featured' }}
// out: true
//
// example: find an active item in a list
// in:  items = [{"name": "Coffee", "status": "sold-out"}, {"name": "Tea", "status": "active"}]
// tpl: {{ items | has:'status','active' }}
// out: true
//
// example: check a config map for a key
// in:  config = {"debug_mode": false, "region": "eu-west-1"}
// tpl: {{ config | has:'debug_mode' }}
// out: true
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

// Is returns true if the value equals any of the given parameters.
//
// value: any
// param:...: any
// returns: bool
//
// example: match a single status
// in:  status = "active"
// tpl: {{ status | is:'active' }}
// out: true
//
// example: match any of several roles
// in:  role = "admin"
// tpl: {{ role | is:'admin','superuser' }}
// out: true
func Is(value any, params []any) (any, error) {
	if len(params) == 0 {
		return false, errors.New("`is` requires at least one parameter")
	}

	for _, param := range params {
		if reflect.DeepEqual(value, param) {
			// log.Trace("collections.Is", "matched", true, "value", value, "param", param)
			return true, nil
		}
	}

	// log.Trace("collections.Is", "matched", false, "value", value, "params", params)
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

				for j := 0; j < v.Len(); j++ {
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
		for i := 0; i < v.Len(); i++ {
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
