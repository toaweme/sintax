// Package access provides modifiers that read elements and fields from collections.
package access

import (
	"errors"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameFirst is the template name for the First modifier.
const ModifierNameFirst functions.ModifierName = "first"

// ModifierNameLast is the template name for the Last modifier.
const ModifierNameLast functions.ModifierName = "last"

// FirstString returns the leading byte of a string as a one-byte string, so it
// is a letter only for ASCII text; a multi-byte character yields a broken
// fragment. An empty string is an error the default modifier can catch.
func FirstString(s string) (string, error) {
	if len(s) == 0 {
		return "", errors.New("first expected a non-empty string")
	}
	return string(s[0]), nil
}

// FirstSlice returns the first element of a slice as-is, keeping its type. An
// empty slice is an error.
func FirstSlice(v []any) (any, error) {
	if len(v) == 0 {
		return nil, errors.New("first expected a non-empty slice")
	}
	return v[0], nil
}

// LastString returns the trailing byte of a string as a one-byte string, with
// the same ASCII caveat as FirstString. An empty string is an error.
func LastString(s string) (string, error) {
	if len(s) == 0 {
		return "", errors.New("last expected a non-empty string")
	}
	return string(s[len(s)-1]), nil
}

// LastSlice returns the last element of a slice as-is. An empty slice is an
// error.
func LastSlice(v []any) (any, error) {
	if len(v) == 0 {
		return nil, errors.New("last expected a non-empty slice")
	}
	return v[len(v)-1], nil
}
