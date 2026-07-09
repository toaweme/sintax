// Package casing provides case-transformation and slug modifiers.
package casing

import (
	"fmt"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameToLower is the template name for the ToLower modifier.
const ModifierNameToLower functions.ModifierName = "lower"

// ModifierNameToUpper is the template name for the ToUpper modifier.
const ModifierNameToUpper functions.ModifierName = "upper"

// ToLower converts a string to lowercase.
//
// value: string
// returns: string
//
// example: normalize an email address
// in:  email = "Alice@Example.COM"
// tpl: {{ email | lower }}
// out: alice@example.com
//
// example: lowercase a heading
// in:  title = "Hello, World!"
// tpl: {{ title | lower }}
// out: hello, world!
func ToLower(value any, _ []any) (any, error) {
	str, err := functions.ValueString(value)
	if err != nil {
		return nil, fmt.Errorf("tolower function expected a string, got %T", value)
	}
	str = strings.ToLower(str)
	return str, nil
}

// ToUpper converts a string to uppercase.
//
// value: string
// returns: string
//
// example: shout a name
// in:  name = "Alice"
// tpl: {{ name | upper }}
// out: ALICE
//
// example: format a country code
// in:  code = "us"
// tpl: {{ code | upper }}
// out: US
func ToUpper(value any, _ []any) (any, error) {
	str, err := functions.ValueString(value)
	if err != nil {
		return nil, fmt.Errorf("toupper function expected a string, got %T", value)
	}
	str = strings.ToUpper(str)
	return str, nil
}
