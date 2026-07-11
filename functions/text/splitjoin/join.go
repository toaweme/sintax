package splitjoin

import (
	"fmt"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameJoin is the template name for the Join modifier.
const ModifierNameJoin functions.ModifierName = "join"

// JoinAny joins the elements of a slice into a single string. Every element must
// be a string; a non-string element is an error. The separator is typed as any
// and defaults to a newline both when it is absent and when it is a non-string
// or the empty string, preserving the original lenient behavior. A []string
// value reaches here as []any via the engine's slice coercion.
func JoinAny(v []any, separator any) (string, error) {
	sep, _ := separator.(string)
	if sep == "" {
		sep = "\n"
	}
	parts := make([]string, len(v))
	for i, item := range v {
		s, ok := item.(string)
		if !ok {
			return "", fmt.Errorf("join expected an array of strings, got %T at index %d", item, i)
		}
		parts[i] = s
	}
	return strings.Join(parts, sep), nil
}

// JoinDefault is the join clause for a call with no separator, joining on a
// newline.
func JoinDefault(v []any) (string, error) {
	return JoinAny(v, "\n")
}
