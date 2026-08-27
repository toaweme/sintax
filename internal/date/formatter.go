// Package date renders time.Time values as strings, and reads them back, using
// PHP-style date format strings that map each recognized character to its Go
// reference time layout token.
package date

import (
	"strings"
	"time"
)

// Formatter renders time.Time values using a PHP-style date format string,
// translating each recognized character to its Go reference time layout token.
type Formatter struct {
	mapping map[string]string
}

// NewFormatter returns a Formatter that renders using the given character-to-Go-layout mapping.
func NewFormatter(mapping map[string]string) *Formatter {
	return &Formatter{mapping: mapping}
}

// Format renders t according to format. Characters without a mapping entry
// are copied through to the output unchanged.
func (f *Formatter) Format(t time.Time, format string) string {
	return t.Format(translate(f.mapping, format))
}

// translate rewrites a PHP-style format into a Go reference time layout,
// copying through any character the mapping does not name. Formatting and
// parsing share it, so the two always read the same layout language.
func translate(mapping map[string]string, format string) string {
	var goLayout strings.Builder
	for i := range len(format) {
		char := string(format[i])
		layout, ok := mapping[char]
		if !ok {
			goLayout.WriteString(char)
			continue
		}
		goLayout.WriteString(layout)
	}
	return goLayout.String()
}
