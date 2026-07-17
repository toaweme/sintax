// Package splitjoin provides modifiers that split and join strings.
package splitjoin

import (
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameSplit is the template name for the Split modifier.
const ModifierNameSplit functions.ModifierName = "split"

// Split splits s on every occurrence of sep. A separator at the start or end
// yields an empty leading or trailing element, a separator that never appears
// returns s as a single element, and an empty separator returns s one UTF-8 rune
// per element.
func Split(s, sep string) ([]string, error) {
	return strings.Split(s, sep), nil
}
