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

// LineNumbers prepends each line of s with its zero-based line number.
func LineNumbers(s string) (string, error) {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = fmt.Sprintf("%d. %s", i, line)
	}
	return strings.Join(lines, "\n"), nil
}
