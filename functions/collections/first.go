package collections

import (
	"fmt"
	"reflect"

	"github.com/toaweme/sintax/functions"
)

var FirstDefinition = functions.ModifierDefinition{
	Description: `Returns the first element of a string, bytes, or array.`,
	Func:        First,
	AcceptedValue: []functions.Type{
		functions.TypeString,
		functions.TypeBytes,
		functions.TypeArray,
	},
	AcceptedParams: nil,
}

func First(value any, params []any) (any, error) {
	// err := FirstDefinition.Validate(value, params)
	// if err != nil {
	// 	return nil, err
	// }
	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.String:
		if v.Len() > 0 {
			return string(v.String()[0]), nil
		}
	case reflect.Slice, reflect.Array:
		if v.Len() > 0 {
			return v.Index(0).Interface(), nil
		}
	}

	return nil, fmt.Errorf("first function expected a non-empty string, slice, or array, got %T", value)
}
