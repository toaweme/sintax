package collections

import (
	"fmt"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameWrap is the template name for the Wrap modifier.
const ModifierNameWrap functions.ModifierName = "wrap"

// Wrap wraps the value in a map under the given key.
//
// value: any
// param:0: string
// returns: map
//
// example: wrap a value for a JSON envelope
// in:  name = "Alice"
// tpl: {{ name | wrap:'user' }}
// out: {"user": "Alice"}
//
// example: wrap a list under a top-level key
// in:  items = ["mug", "pen", "pad"]
// tpl: {{ items | wrap:'data' }}
// out: {"data": ["mug", "pen", "pad"]}
var Wrap = func(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, fmt.Errorf("wrap requires at least one parameter")
	}

	key := ""
	switch v := params[0].(type) {
	case string:
		key = v
	}
	return map[string]any{key: value}, nil
}
