package edit

import (
	"errors"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameWrap is the template name for the Wrap modifier.
const ModifierNameWrap functions.ModifierName = "wrap"

// Wrap nests the value inside a new single-entry map under the given key. The
// result is a map value, not a JSON string, so it is meant to be piped into an
// encoder (or returned as structured data) rather than rendered inline. When a
// template does stringify it directly, it comes out in Go's native map form
// (for example map[user:Alice]), not as JSON.
//
// The value can be of any type. The key is taken from the first parameter; a
// non-string key coerces to the empty string.
//
// value: any
// param:0: string
// returns: map
//
// example: nest a value under a key, shown in Go's inline map form
// in:  name = "Alice"
// tpl: {{ name | wrap:'user' }}
// out: map[user:Alice]
//
// example: nest a list under a top-level key
// in:  items = ["mug", "pen", "pad"]
// tpl: {{ items | wrap:'data' }}
// out: map[data:[mug pen pad]]
var Wrap = func(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, errors.New("wrap requires at least one parameter")
	}

	key, _ := params[0].(string)
	return map[string]any{key: value}, nil
}
