package text

import (
	"errors"
	"fmt"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameSplit is the template name for the Split modifier.
const ModifierNameSplit functions.ModifierName = "split"

// Split splits a string into an array using a separator.
//
// value: string
// param:0: string
// returns: array
//
// example: split a CSV line into fields
// in:  csv_line = "Alice,42,admin"
// tpl: {{ csv_line | split:',' }}
// out: ["Alice", "42", "admin"]
//
// example: break a path into segments
// in:  path = "/var/log/app/server.log"
// tpl: {{ path | split:'/' }}
// out: ["", "var", "log", "app", "server.log"]
//
// example: split a sentence into words
// in:  tags = "coffee tea espresso"
// tpl: {{ tags | split:' ' }}
// out: ["coffee", "tea", "espresso"]
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
