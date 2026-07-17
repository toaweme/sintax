// Package casing provides case-transformation and slug modifiers.
package casing

import (
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameToLower is the template name for the ToLower modifier.
const ModifierNameToLower functions.ModifierName = "lower"

// ModifierNameToUpper is the template name for the ToUpper modifier.
const ModifierNameToUpper functions.ModifierName = "upper"

// ToLower converts a string to lowercase.
func ToLower(s string) (string, error) {
	return strings.ToLower(s), nil
}

// ToUpper converts a string to uppercase.
func ToUpper(s string) (string, error) {
	return strings.ToUpper(s), nil
}
