// Package access provides modifiers that read elements and fields from collections.
package access

import "github.com/toaweme/sintax/functions"

// ModifierNameFirst is the template name for the First modifier.
const ModifierNameFirst functions.ModifierName = "first"

// ModifierNameLast is the template name for the Last modifier.
const ModifierNameLast functions.ModifierName = "last"

// An empty collection is a "nothing here" condition rather than a broken
// template, so every clause below reports it with functions.Miss. That makes
// `{{ items | filter:'id',9 | first | default:'none' }}` fall back instead of
// failing, matching find, which already reports a miss the same way. Without a
// default in the chain the error still surfaces, so an empty collection never
// passes silently as nil.

// FirstString returns the leading byte of a string as a one-byte string, so it
// is a letter only for ASCII text; a multi-byte character yields a broken
// fragment. An empty string is a miss the default modifier can catch.
func FirstString(s string) (string, error) {
	if len(s) == 0 {
		return "", functions.Miss("first expected a non-empty string")
	}
	return string(s[0]), nil
}

// FirstBytes returns the leading byte of a []byte as a one-byte string, so a
// text buffer (a file read, an HTTP body) reads the same as the string clause
// rather than falling through to the slice clause and yielding a raw byte
// number. The same ASCII caveat as FirstString applies. An empty buffer is a
// miss the default modifier can catch.
func FirstBytes(b []byte) (string, error) {
	if len(b) == 0 {
		return "", functions.Miss("first expected a non-empty []byte")
	}
	return string(b[0]), nil
}

// FirstSlice returns the first element of a slice as-is, keeping its type. An
// empty slice is a miss the default modifier can catch.
func FirstSlice(v []any) (any, error) {
	if len(v) == 0 {
		return nil, functions.Miss("first expected a non-empty slice")
	}
	return v[0], nil
}

// LastString returns the trailing byte of a string as a one-byte string, with
// the same ASCII caveat as the first modifier. An empty string is a miss the
// default modifier can catch.
func LastString(s string) (string, error) {
	if len(s) == 0 {
		return "", functions.Miss("last expected a non-empty string")
	}
	return string(s[len(s)-1]), nil
}

// LastBytes returns the trailing byte of a []byte as a one-byte string, with the
// same ASCII caveat as LastString, so a text buffer reads the same as the string
// clause instead of yielding a raw byte number via the slice clause. An empty
// buffer is a miss the default modifier can catch.
func LastBytes(b []byte) (string, error) {
	if len(b) == 0 {
		return "", functions.Miss("last expected a non-empty []byte")
	}
	return string(b[len(b)-1]), nil
}

// LastSlice returns the last element of a slice as-is. An empty slice is a miss
// the default modifier can catch.
func LastSlice(v []any) (any, error) {
	if len(v) == 0 {
		return nil, functions.Miss("last expected a non-empty slice")
	}
	return v[len(v)-1], nil
}
