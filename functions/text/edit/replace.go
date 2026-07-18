package edit

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameReplace is the template name for the Replace modifier.
const ModifierNameReplace functions.ModifierName = "replace"

// ModifierNameReplacePattern is the template name for the ReplacePattern modifier.
const ModifierNameReplacePattern functions.ModifierName = "replace_pattern"

// ModifierNameReverse is the template name for the Reverse modifier.
const ModifierNameReverse functions.ModifierName = "reverse"

// Replace replaces every occurrence of the old substring with the replacement in
// the input.
func Replace(s, old, replacement string) (string, error) {
	return strings.ReplaceAll(s, old, replacement), nil
}

// ReplacePattern replaces every match of the RE2 pattern in the input with the
// replacement, which may reference capture groups with $1, $2 and so on. An empty
// replacement deletes each match, and an invalid pattern is an error.
func ReplacePattern(s, pattern, replacement string) (string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", fmt.Errorf("failed to compile replace_pattern regex %q: %w", pattern, err)
	}
	return re.ReplaceAllString(s, replacement), nil
}

// Reverse reverses the input by rune, so multi-byte characters stay intact. A
// cluster built from several code points (a combining sequence, or an emoji) is
// reversed code point by code point, which can reorder it.
func Reverse(s string) (string, error) {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes), nil
}
