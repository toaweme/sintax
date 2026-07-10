// Package access provides modifiers that read elements and fields from collections.
package access

import (
	"fmt"
	"reflect"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameFirst is the template name for the First modifier.
const ModifierNameFirst functions.ModifierName = "first"

// ModifierNameLast is the template name for the Last modifier.
const ModifierNameLast functions.ModifierName = "last"

// First returns the first element of a slice or the first byte of a string.
// For a slice it returns the element as-is (a map, number, or nested slice
// keeps its type). For a string it returns the leading byte, so it is a letter
// only for ASCII text. A multi-byte character (an accented letter, an emoji, a
// CJK glyph) is more than one byte, so on such text first yields a broken
// fragment of that character rather than the character itself. An empty string,
// an empty slice, or a non-collection value (a number, a bool, nil) is an
// error, which the default modifier can then supply a fallback for.
//
// value: string, array
// returns: string, any
//
// example: pick the first item from a list
// in:  items = ["espresso", "latte", "macchiato"]
// tpl: {{ items | first }}
// out: espresso
//
// example: take the first letter of an ASCII name
// in:  name = "Alice"
// tpl: {{ name | first }}
// out: A
//
// example: the first element of a list of maps keeps its shape
// in:  users = [{"id": 1, "name": "Alice"}, {"id": 2, "name": "Bob"}]
// tpl: {{ users | first | key:'name' }}
// out: Alice
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
	default:
	}

	return nil, fmt.Errorf("first function expected a non-empty string, slice, or array, got %T", value)
}

// Last returns the last element of a slice or the last byte of a string.
// For a slice it returns the element as-is, keeping its type. For a string it
// returns the trailing byte, so it is a letter only for ASCII text. On text
// ending in a multi-byte character it yields a broken fragment of that
// character rather than the character itself. An empty string, an empty slice,
// or a non-collection value is an error, which the default modifier can catch.
//
// value: string, array
// returns: string, any
//
// example: pick the last item from a list
// in:  items = ["espresso", "latte", "macchiato"]
// tpl: {{ items | last }}
// out: macchiato
//
// example: take the last letter of an ASCII word
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
	default:
	}

	return nil, fmt.Errorf("last function expected a non-empty string, slice, or array, got %T", value)
}
