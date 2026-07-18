package format

import (
	"fmt"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameLineNumbers is the template name for the LineNumbers modifier.
const ModifierNameLineNumbers functions.ModifierName = "line_numbers"

// lineNumbersNilEmpty passes a nil or empty-string value through as nil, and
// declines any other value so Overload falls through to the string clause.
func lineNumbersNilEmpty(value any, _ []any) (any, error) {
	if value == nil || value == "" {
		return nil, nil //nolint:nilnil // deliberate, nil/empty input passes through as nil, not an error
	}
	return nil, functions.ErrInvalidValueType
}

// LineNumbers prepends each line of s with its one-based line number.
func LineNumbers(s string) (string, error) {
	return numberLines(s, 1), nil
}

// LineNumbersFrom prepends each line of s with a line number counting up from
// start, so `line_numbers:6` renders a block as if it began at line six.
func LineNumbersFrom(s string, start int) (string, error) {
	return numberLines(s, start), nil
}

// numberLines prefixes every line of s with a running number beginning at start.
func numberLines(s string, start int) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = fmt.Sprintf("%d. %s", start+i, line)
	}
	return strings.Join(lines, "\n")
}
