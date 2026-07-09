// Package control provides template modifiers that steer value resolution
// rather than transform data, such as falling back when a value is missing.
package control

import (
	"errors"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameDefault is the template name for the Default modifier.
const ModifierNameDefault functions.ModifierName = "default"

// Default returns the fallback value if the input is nil or an empty string.
//
// value: any
// param:0: any
// returns: any
//
// example: fall back to a placeholder name
// in:  name = ""
// tpl: {{ name | default:'anonymous' }}
// out: anonymous
//
// example: use a default count when missing
// in:  count = null
// tpl: {{ count | default:0 }}
// out: 0
//
// example: fall back to an empty object when a lookup misses
// in:  items = [{"id": 1, "name": "Mug"}]
// tpl: {{ items | find:'id',99 | default:{} }}
// out: {}
func Default(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, errors.New("default requires at least one parameter")
	}
	if value == nil {
		return params[0], nil
	}
	if value == "" {
		return params[0], nil
	}

	return value, nil
}
