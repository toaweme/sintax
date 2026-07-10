// Package edit provides modifiers that rewrite string content.
package edit

import (
	"fmt"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameConcat is the template name for the Concat modifier.
const ModifierNameConcat functions.ModifierName = "concat"

// Concat appends one or more strings to the value, joining them with no
// separator. It takes any number of parameters and glues them on in order.
// Both the value and every parameter must be strings; a non-string anywhere
// is an error, so coerce numbers to strings before piping them in.
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
// example: append several parts in one call
// in:  base = "file"
// tpl: {{ base | concat:'-','01','.txt' }}
// out: file-01.txt
//
// example: build a kebab-cased identifier across two calls
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
