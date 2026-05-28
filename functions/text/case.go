package text

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameToLower is the template name for the ToLower modifier.
const ModifierNameToLower functions.ModifierName = "lower"

// ModifierNameToUpper is the template name for the ToUpper modifier.
const ModifierNameToUpper functions.ModifierName = "upper"

// ModifierNameReplace is the template name for the Replace modifier.
const ModifierNameReplace functions.ModifierName = "replace"

// ModifierNameReplacePattern is the template name for the ReplacePattern modifier.
const ModifierNameReplacePattern functions.ModifierName = "replace_pattern"

// ModifierNameReverse is the template name for the Reverse modifier.
const ModifierNameReverse functions.ModifierName = "reverse"

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

// Replace replaces all occurrences of a substring within the string value.
//
// value: string
// param:0: string
// param:1: string
// returns: string
//
// example: swap one word for another
// in:  greeting = "Hello world"
// tpl: {{ greeting | replace:'world','everyone' }}
// out: Hello everyone
//
// example: redact a phrase
// in:  text = "The password is hunter2"
// tpl: {{ text | replace:'hunter2','******' }}
// out: The password is ******
func Replace(value any, params []any) (any, error) {
	if len(params) < 2 {
		return nil, fmt.Errorf("replace function requires at least two parameters")
	}

	str, err := functions.ValueString(value)
	if err != nil {
		return nil, fmt.Errorf("replace function expected a string, got %T", value)
	}

	old, err := functions.ParamString(params, 0)
	if err != nil {
		return nil, fmt.Errorf("replace function expected first parameter to be a string, got %T", params[0])
	}

	newStr, err := functions.ParamString(params, 1)
	if err != nil {
		return nil, fmt.Errorf("replace function expected second parameter to be a string, got %T", params[1])
	}

	result := strings.ReplaceAll(str, old, newStr)
	return result, nil
}

// ReplacePattern replaces all regex matches within the string value.
// Supports capture group references (e.g. $1) in the replacement string.
//
// value: string
// param:0: string
// param:1: string
// returns: string
//
// example: collapse runs of whitespace
// in:  text = "hello    world"
// tpl: {{ text | replace_pattern:'\s+',' ' }}
// out: hello world
//
// example: keep only slug-safe characters
// in:  slug = "hello-world!@#"
// tpl: {{ slug | replace_pattern:'[^a-z0-9\-]',” }}
// out: hello-world
func ReplacePattern(value any, params []any) (any, error) {
	if len(params) < 2 {
		return nil, fmt.Errorf("replacePattern function requires at least two parameters")
	}

	str, err := functions.ValueString(value)
	if err != nil {
		return nil, fmt.Errorf("replacePattern function expected a string, got %T", value)
	}

	pattern, err := functions.ParamString(params, 0)
	if err != nil {
		return nil, fmt.Errorf("replacePattern function expected first parameter to be a string, got %T", params[0])
	}

	replacement, err := functions.ParamString(params, 1)
	if err != nil {
		return nil, fmt.Errorf("replacePattern function expected second parameter to be a string, got %T", params[1])
	}

	// compile the regular expression
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("replacePattern function invalid regex pattern: %w", err)
	}

	result := re.ReplaceAllString(str, replacement)
	return result, nil
}

// Reverse reverses the characters in a string.
//
// value: string
// returns: string
//
// example: reverse a name
// in:  name = "Alice"
// tpl: {{ name | reverse }}
// out: ecilA
//
// example: reverse a short code
// in:  code = "ABC123"
// tpl: {{ code | reverse }}
// out: 321CBA
func Reverse(value any, _ []any) (any, error) {
	str, err := functions.ValueString(value)
	if err != nil {
		return nil, fmt.Errorf("reverse function expected a string, got %T", value)
	}

	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes), nil
}
