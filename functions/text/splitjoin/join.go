package splitjoin

import (
	"fmt"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameJoin is the template name for the Join modifier.
const ModifierNameJoin functions.ModifierName = "join"

// JoinAny joins the elements of a slice into a single string. Every element must
// be a string; a non-string element is an error. The separator defaults to a
// newline when it is omitted or empty.
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
