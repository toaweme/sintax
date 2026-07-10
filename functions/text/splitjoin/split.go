// Package splitjoin provides modifiers that split and join strings.
package splitjoin

import (
	"errors"
	"fmt"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameSplit is the template name for the Split modifier.
const ModifierNameSplit functions.ModifierName = "split"

// Split splits a string into an array using a separator. The separator is
// required. Every occurrence of the separator is a cut, so a separator at the
// very start or end of the string produces an empty leading or trailing
// element, and a separator that never appears returns the whole string as a
// single-element array. Splitting on an empty separator ("") returns the input
// one UTF-8 rune per element.
//
// value: string
// param:0: string, the separator to cut on (required)
// returns: array
//
// example: split a CSV line into fields
// in:  csv_line = "Alice,42,admin"
// tpl: {{ csv_line | split:',' }}
// out: ["Alice", "42", "admin"]
//
// example: break a path into segments, note the empty leading element from the root slash
// in:  path = "/var/log/app/server.log"
// tpl: {{ path | split:'/' }}
// out: ["", "var", "log", "app", "server.log"]
//
// example: split a sentence into words
// in:  tags = "coffee tea espresso"
// tpl: {{ tags | split:' ' }}
// out: ["coffee", "tea", "espresso"]
//
// example: a separator that is not present returns the whole string
// in:  name = "Alice"
// tpl: {{ name | split:',' }}
// out: ["Alice"]
func Split(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, errors.New("split function requires at least one parameter")
	}

	str, err := functions.ValueString(value)
	if err != nil {
		return nil, fmt.Errorf("split function expected a string, got %T", value)
	}

	sep, err := functions.ParamString(params, 0)
	if err != nil {
		return nil, fmt.Errorf("split function expected first parameter to be a string, got %T", params[0])
	}

	result := strings.Split(str, sep)

	return result, nil
}
