// Package control provides template modifiers that steer value resolution
// rather than transform data, such as falling back when a value is missing.
package control

import (
	"errors"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameDefault is the template name for the Default modifier.
const ModifierNameDefault functions.ModifierName = "default"

// Default supplies a fallback so a template never renders a missing value. It
// swaps in the fallback in exactly two situations, and passes the real value
// through untouched in every other case.
//
// It applies the fallback when the piped value is nil (an absent or null
// variable) or an empty string. That is the whole rule for the value itself.
// An empty list, an empty object, the number zero, and the boolean false are
// all real values, so they are kept as-is and the fallback is NOT used. This
// matters for a non-technical author. `default` guards against "nothing there",
// not against "a value that happens to be empty or zero".
//
// The second situation is a soft failure earlier in the same pipe. When a
// preceding modifier gives up without a hard error, for example `find` not
// matching any row, the engine turns that into nil and lets `default` catch it.
// So `find:'id',99 | default:{}` renders the fallback when nothing matched
// rather than failing the whole template.
//
// The fallback is any literal, including an empty-collection literal. Write `[]`
// for an empty list or `{}` for an empty object when you want a safe, iterable
// stand-in for a missing collection.
//
// value: any
// param:0: any (the fallback, required)
// returns: any
//
// example: fall back to a placeholder name for an empty string
// in:  name = ""
// tpl: {{ name | default:'anonymous' }}
// out: anonymous
//
// example: fall back to a count when the variable is absent or null
// in:  count = null
// tpl: {{ count | default:0 }}
// out: 0
//
// example: a present value is kept, the fallback is ignored
// in:  name = "Ada"
// tpl: {{ name | default:'anonymous' }}
// out: Ada
//
// example: zero is a real value, so it is kept
// in:  count = 0
// tpl: {{ count | default:5 }}
// out: 0
//
// example: fall back to an empty list when a lookup finds nothing
// in:  items = [{"id": 1, "name": "Mug"}]
// tpl: {{ items | find:'id',99 | default:[] }}
// out: []
//
// example: fall back to an empty object when a lookup finds nothing
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
