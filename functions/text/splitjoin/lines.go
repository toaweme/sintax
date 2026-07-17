package splitjoin

import (
	"bytes"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameLines is the template name for the Lines modifier.
const ModifierNameLines functions.ModifierName = "lines"

// LinesString splits a string on "\n" only. It does not strip carriage returns,
// so "\r\n" endings leave a trailing "\r" on each line, and a trailing newline
// produces a final empty element.
func LinesString(s string) ([]string, error) {
	return strings.Split(s, "\n"), nil
}

// LinesBytes splits a byte slice on "\n", with the same rules as LinesString.
func LinesBytes(b []byte) ([][]byte, error) {
	return bytes.Split(b, []byte("\n")), nil
}

// linesNil is the lines clause that passes a nil value straight through as nil
// rather than erroring, and declines any other value so Overload falls through
// to the string and []byte clauses.
func linesNil(value any, _ []any) (any, error) {
	if value == nil {
		return nil, nil //nolint:nilnil // deliberate, nil input passes through as nil, not an error
	}
	return nil, functions.ErrInvalidValueType
}
