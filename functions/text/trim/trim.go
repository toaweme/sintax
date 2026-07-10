// Package trim provides whitespace and affix trimming modifiers.
package trim

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

// Trim removes leading and trailing whitespace from the value. With a parameter
// it instead removes any of the characters in that parameter from both ends. The
// parameter is a character set (a cutset), not a fixed prefix or suffix, so
// trim:'/' strips every leading and trailing slash. When a cutset is given the
// default whitespace stripping no longer applies.
//
// value: string, bytes
// param:0?: string (a set of characters to trim from both ends; defaults to whitespace)
// returns: string, bytes
//
// example: clean whitespace from a name
// in:  name = "  Alice  "
// tpl: {{ name | trim }}
// out: Alice
//
// example: strip wrapping characters (every leading and trailing comma)
// in:  list = ",apples,bananas,"
// tpl: {{ list | trim:',' }}
// out: apples,bananas
//
// example: the cutset removes any of its characters, not a fixed string
// in:  tags = "xy-hello-yx"
// tpl: {{ tags | trim:'xy' }}
// out: -hello-
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

// TrimPrefix removes the given prefix from the start of the value. The parameter
// is matched as a whole string, once, and if the value does not start with it the
// value is returned unchanged. Without a parameter it trims leading whitespace
// (spaces, tabs, and newlines) instead.
//
// value: string, bytes
// param:0?: string (the exact prefix to remove; defaults to trimming leading whitespace)
// returns: string, bytes
//
// example: drop a leading slash from a path
// in:  path = "/api/v1/users"
// tpl: {{ path | trim_prefix:'/' }}
// out: api/v1/users
//
// example: no-op when the prefix is absent (the value is returned unchanged)
// in:  path = "api/v1/users"
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

// TrimSuffix removes the given suffix from the end of the value. The parameter is
// matched as a whole string, once, and if the value does not end with it the value
// is returned unchanged. Without a parameter it trims trailing whitespace (spaces,
// tabs, and newlines) instead.
//
// value: string, bytes
// param:0?: string (the exact suffix to remove; defaults to trimming trailing whitespace)
// returns: string, bytes
//
// example: drop a trailing slash from a URL
// in:  url = "https://example.com/users/"
// tpl: {{ url | trim_suffix:'/' }}
// out: https://example.com/users
//
// example: drop a file extension
// in:  file = "report.txt"
// tpl: {{ file | trim_suffix:'.txt' }}
// out: report
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
