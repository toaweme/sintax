package collections

import (
	"fmt"
	"reflect"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameFirst is the template name for the First modifier.
const ModifierNameFirst functions.ModifierName = "first"

// ModifierNameLast is the template name for the Last modifier.
const ModifierNameLast functions.ModifierName = "last"

// First returns the first character of a string or the first element of a slice.
//
// value: string, array
// returns: string, any
//
// example: pick the first item from a list
// in:  items = ["espresso", "latte", "macchiato"]
// tpl: {{ items | first }}
// out: espresso
//
// example: take the first letter of a name
// in:  name = "Alice"
// tpl: {{ name | first }}
// out: A
func First(value any, params []any) (any, error) {
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

// Last returns the last character of a string or the last element of a slice.
//
// value: string, array
// returns: string, any
//
// example: pick the last item from a list
// in:  items = ["espresso", "latte", "macchiato"]
// tpl: {{ items | last }}
// out: macchiato
//
// example: take the last letter of a word
// in:  word = "Hello"
// tpl: {{ word | last }}
// out: o
func Last(value any, params []any) (any, error) {
	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.String:
		if v.Len() > 0 {
			return string(v.String()[v.Len()-1]), nil
		}
	case reflect.Slice, reflect.Array:
		if v.Len() > 0 {
			return v.Index(v.Len() - 1).Interface(), nil
		}
	}

	return nil, fmt.Errorf("last function expected a non-empty string, slice, or array, got %T", value)
}
