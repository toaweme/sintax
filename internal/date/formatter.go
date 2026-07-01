// Package date renders time.Time values using PHP-style date format strings,
// mapping each recognized character to its Go reference time layout token.
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
	var goLayout strings.Builder
	for i := range len(format) {
		char := string(format[i])
		layout, ok := f.mapping[char]
		if !ok {
			goLayout.WriteString(char)
			continue
		}
		goLayout.WriteString(layout)
	}
	return t.Format(goLayout.String())
}
