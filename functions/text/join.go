package text

import (
	"fmt"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameJoin is the template name for the Join modifier.
const ModifierNameJoin functions.ModifierName = "join"

// Join combines an array of strings into a single string with a separator.
// Defaults to newline if no separator is provided.
//
// value: array
// param:0?: string
// returns: string
//
// example: join tags into a CSV string
// in:  tags = ["coffee", "sale", "new"]
// tpl: {{ tags | join:',' }}
// out: coffee,sale,new
//
// example: rebuild lines into a paragraph
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
			if str, ok := item.(string); ok {
				strs[i] = str
			} else {
				return nil, fmt.Errorf("join function expected array of strings, got %T at index %d", item, i)
			}
		}
		return strings.Join(strs, separator), nil
	}

	return nil, fmt.Errorf("join function expected array of strings, got %T", value)
}
