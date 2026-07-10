// Package edit provides modifiers that rewrite string content.
package edit

import (
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameConcat is the template name for the Concat modifier.
const ModifierNameConcat functions.ModifierName = "concat"

// Concat appends parts to s in order, joining with no separator. Both the value
// and every part must be strings, so coerce numbers to strings before piping
// them in.
func Concat(s string, parts ...string) (string, error) {
	return s + strings.Join(parts, ""), nil
}
