package splitjoin

import (
	"fmt"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameJoin is the template name for the Join modifier.
const ModifierNameJoin functions.ModifierName = "join"

// Join combines an array of strings into a single string with a separator.
// The separator is optional and defaults to a newline when it is omitted (or
// when a non-string value is passed where the separator is expected). Every
// element of the array must be a string, a non-string element is an error. An
// empty array joins to an empty string and a single-element array returns that
// element unchanged.
//
// value: array of strings
// param:0?: string, the separator placed between elements (defaults to "\n")
// returns: string
//
// example: join tags into a CSV string
// in:  tags = ["coffee", "sale", "new"]
// tpl: {{ tags | join:',' }}
// out: coffee,sale,new
//
// example: rebuild lines into a paragraph, the default separator is a newline
// in:  lines = ["Dear Alice,", "Welcome aboard.", "Best,"]
// tpl: {{ lines | join }}
// out: Dear Alice,
// out: Welcome aboard.
// out: Best,
//
// example: join with a custom separator
// in:  steps = ["sign in", "verify email", "start tour"]
// tpl: {{ steps | join:' | ' }}
// out: sign in | verify email | start tour
func Join(value any, params []any) (any, error) {
	separator, _ := functions.ParamString(params, 0)
	if separator == "" {
		separator = "\n"
	}

	switch v := value.(type) {
	case []string:
		return strings.Join(v, separator), nil
	case []any:
		// Convert []any to []string
		strs := make([]string, len(v))
		for i, item := range v {
			str, ok := item.(string)
			if !ok {
				return nil, fmt.Errorf("join function expected array of strings, got %T at index %d", item, i)
			}
			strs[i] = str
		}
		return strings.Join(strs, separator), nil
	}

	return nil, fmt.Errorf("join function expected array of strings, got %T", value)
}
