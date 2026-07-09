package text

import (
	"fmt"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameTrim is the template name for the Trim modifier.
const ModifierNameTrim functions.ModifierName = "trim"

// ModifierNameTrimPrefix is the template name for the TrimPrefix modifier.
const ModifierNameTrimPrefix functions.ModifierName = "trim_prefix"

// ModifierNameTrimSuffix is the template name for the TrimSuffix modifier.
const ModifierNameTrimSuffix functions.ModifierName = "trim_suffix"

// Trim removes leading and trailing whitespace, or the given character set.
//
// value: string, bytes
// param:0?: string
// returns: string, bytes
//
// example: clean whitespace from a name
// in:  name = "  Alice  "
// tpl: {{ name | trim }}
// out: Alice
//
// example: strip a wrapping character
// in:  list = ",apples,bananas,"
// tpl: {{ list | trim:',' }}
// out: apples,bananas
var Trim = func(value any, params []any) (any, error) {
	switch v := value.(type) {
	case string:
		if len(params) == 0 {
			return strings.TrimSpace(v), nil
		}
		chars, ok := params[0].(string)
		if !ok {
			return nil, fmt.Errorf("expected string as first parameter, got %T", params[0])
		}
		return strings.Trim(v, chars), nil
	case []byte:
		if len(params) == 0 {
			return []byte(strings.TrimSpace(string(v))), nil
		}
		chars, ok := params[0].(string)
		if !ok {
			return nil, fmt.Errorf("expected string as first parameter, got %T", params[0])
		}
		return []byte(strings.Trim(string(v), chars)), nil
	default:
		return nil, fmt.Errorf("expected string or []byte, got %T", value)
	}
}

// TrimPrefix removes a leading prefix string or leading whitespace from the value.
//
// value: string, bytes
// param:0?: string
// returns: string, bytes
//
// example: drop a leading slash from a path
// in:  path = "/api/v1/users"
// tpl: {{ path | trim_prefix:'/' }}
// out: api/v1/users
//
// example: trim leading whitespace from a paragraph
// in:  text = "   Welcome aboard."
// tpl: {{ text | trim_prefix }}
// out: Welcome aboard.
var TrimPrefix = func(value any, params []any) (any, error) {
	cutset := "\n \t"

	switch v := value.(type) {
	case string:
		if len(params) == 0 {
			return strings.TrimLeft(v, cutset), nil
		}
		chars, ok := params[0].(string)
		if !ok {
			return nil, fmt.Errorf("expected string as first parameter, got %T", params[0])
		}
		return strings.TrimPrefix(v, chars), nil
	case []byte:
		if len(params) == 0 {
			return []byte(strings.TrimLeft(string(v), cutset)), nil
		}
		chars, ok := params[0].(string)
		if !ok {
			return nil, fmt.Errorf("expected string as first parameter, got %T", params[0])
		}

		return []byte(strings.TrimPrefix(string(v), chars)), nil
	default:
		return nil, fmt.Errorf("expected string or []byte, got %T", value)
	}
}

// TrimSuffix removes a trailing suffix string or trailing whitespace from the value.
//
// value: string, bytes
// param:0?: string
// returns: string, bytes
//
// example: drop a trailing slash from a URL
// in:  url = "https://example.com/users/"
// tpl: {{ url | trim_suffix:'/' }}
// out: https://example.com/users
//
// example: trim trailing whitespace from a paragraph
// in:  text = "Welcome aboard.   "
// tpl: {{ text | trim_suffix }}
// out: Welcome aboard.
var TrimSuffix = func(value any, params []any) (any, error) {
	cutset := "\n \t"

	switch v := value.(type) {
	case string:
		if len(params) == 0 {
			return strings.TrimRight(v, cutset), nil
		}
		chars, ok := params[0].(string)
		if !ok {
			return nil, fmt.Errorf("expected string as first parameter, got %T", params[0])
		}
		return strings.TrimSuffix(v, chars), nil
	case []byte:
		if len(params) == 0 {
			return []byte(strings.TrimRight(string(v), cutset)), nil
		}
		chars, ok := params[0].(string)
		if !ok {
			return nil, fmt.Errorf("expected string as first parameter, got %T", params[0])
		}

		return []byte(strings.TrimSuffix(string(v), chars)), nil
	default:
		return nil, fmt.Errorf("expected string or []byte, got %T", value)
	}
}
