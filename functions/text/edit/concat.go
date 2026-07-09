// Package edit provides modifiers that rewrite string content.
package edit

import (
	"fmt"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameConcat is the template name for the Concat modifier.
const ModifierNameConcat functions.ModifierName = "concat"

// Concat appends one or more strings to the value.
//
// value: string
// param:...: string
// returns: string
//
// example: append a punctuation mark
// in:  greeting = "Hello"
// tpl: {{ greeting | concat:'!' }}
// out: Hello!
//
// example: build a kebab-cased identifier
// in:  prefix = "user"
// in:  suffix = "profile"
// tpl: {{ prefix | concat:'-' | concat:suffix }}
// out: user-profile
func Concat(value any, params []any) (any, error) {
	v, ok := value.(string)
	if !ok {
		return nil, fmt.Errorf("concat function expected string, got %T", value)
	}

	joined := make([]string, 0)
	joined = append(joined, v)
	for _, p := range params {
		switch pv := p.(type) {
		case string:
			joined = append(joined, pv)
		default:
			return nil, fmt.Errorf("concat function expected string params, got %T", p)
		}
	}
	return strings.Join(joined, ""), nil
}
