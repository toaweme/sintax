package edit

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameReplace is the template name for the Replace modifier.
const ModifierNameReplace functions.ModifierName = "replace"

// ModifierNameReplacePattern is the template name for the ReplacePattern modifier.
const ModifierNameReplacePattern functions.ModifierName = "replace_pattern"

// ModifierNameReverse is the template name for the Reverse modifier.
const ModifierNameReverse functions.ModifierName = "reverse"

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
		return nil, errors.New("replace function requires at least two parameters")
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

// ReplacePattern replaces every regular-expression match in the string value.
// param:0 is a Go (RE2) regular expression and param:1 is the replacement. The
// replacement may reference capture groups with $1, $2 and so on, so a pattern
// with parentheses can reorder or reuse the matched text. Pass an empty string
// as the replacement to delete every match. An invalid regex is an error.
//
// value: string
// param:0: string
// param:1: string
// returns: string
//
// example: collapse runs of whitespace into a single space
// in:  text = "hello    world"
// tpl: {{ text | replace_pattern:'\s+',' ' }}
// out: hello world
//
// example: delete everything that is not slug-safe (empty replacement)
// in:  slug = "hello-world!@#"
// tpl: {{ slug | replace_pattern:'[^a-z0-9\-]',"" }}
// out: hello-world
//
// example: reorder captured groups with $1 and $2
// in:  name = "Doe, Jane"
// tpl: {{ name | replace_pattern:'(\w+), (\w+)','$2 $1' }}
// out: Jane Doe
func ReplacePattern(value any, params []any) (any, error) {
	if len(params) < 2 {
		return nil, errors.New("replacePattern function requires at least two parameters")
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

// Reverse reverses the characters in a string. It reverses by rune, not by
// byte, so multi-byte characters (accented letters, CJK glyphs, and similar)
// stay intact rather than being split. Combining sequences and emoji built
// from several code points are reversed code point by code point, which can
// reorder such a cluster.
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
//
// example: multi-byte characters stay intact
// in:  word = "café"
// tpl: {{ word | reverse }}
// out: éfac
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
