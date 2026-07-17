// Package edit provides modifiers that rewrite string content.
package edit

import (
	"fmt"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameConcat is the template name for the Concat modifier.
const ModifierNameConcat functions.ModifierName = "concat"

// Concat joins the value and every part into one string, in order, with no
// separator. Concatenation is inherently textual, so the value and each part may
// be a scalar (a number or bool) and are taken as their string form; a composite
// value or part (a slice or map) is an error.
func Concat(value any, parts ...any) (string, error) {
	head, ok := functions.Stringish(value)
	if !ok {
		return "", fmt.Errorf("concat expected a scalar value, got %T: %w", value, functions.ErrInvalidValueType)
	}
	var b strings.Builder
	b.WriteString(head)
	for i, p := range parts {
		s, ok := functions.Stringish(p)
		if !ok {
			return "", fmt.Errorf("concat expected a scalar at param %d, got %T: %w", i, p, functions.ErrInvalidParamType)
		}
		b.WriteString(s)
	}
	return b.String(), nil
}
