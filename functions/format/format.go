// Package format provides template modifiers that render a value as a
// human-readable string: dates, decimals, currency, line numbering, and size.
package format

import (
	"fmt"
	"time"

	"github.com/toaweme/sintax/functions"
	"github.com/toaweme/sintax/internal/date"
)

// ModifierNameFormat is the template name for the Format modifier.
const ModifierNameFormat functions.ModifierName = "format"

// Format renders a time.Time value as a string using a PHP-style date format.
// A string value is passed through unchanged, so it is safe to apply to a field
// that may already be formatted. When no format is given, the default is
// "Y-m-d H:i:s". Any character with no mapping (a space, a slash, punctuation)
// is emitted literally, so format strings can freely mix tokens and separators.
//
// Common tokens: Y is the 4-digit year, y the 2-digit year, m the zero-padded
// month, n the month without padding, F the full month name, M the short month
// name, d the zero-padded day, j the day without padding, l the full weekday
// name, D the short weekday name, H the zero-padded 24-hour, h the zero-padded
// 12-hour, i the minutes, s the seconds, A the uppercase meridiem (AM/PM), a the
// lowercase meridiem, T the timezone abbreviation, and P the timezone offset with
// a colon.
//
// value: string, time.Time
// param:0?: string - a PHP-style date format (default "Y-m-d H:i:s")
// returns: string
//
// example: render a date in ISO form
// in:  created_at = time.Date(2024, 3, 14, 9, 30, 0, 0, time.UTC)
// tpl: {{ created_at | format:'Y-m-d' }}
// out: 2024-03-14
//
// example: render a date and time for display
// in:  published = time.Date(2024, 3, 14, 9, 30, 0, 0, time.UTC)
// tpl: {{ published | format:'d/m/Y H:i' }}
// out: 14/03/2024 09:30
//
// example: render a long human-readable date
// in:  published = time.Date(2024, 3, 14, 9, 30, 0, 0, time.UTC)
// tpl: {{ published | format:'l, F j, Y' }}
// out: Thursday, March 14, 2024
//
// example: apply the default format when none is given
// in:  published = time.Date(2024, 3, 14, 9, 30, 5, 0, time.UTC)
// tpl: {{ published | format }}
// out: 2024-03-14 09:30:05
//
// example: pass an already-formatted string straight through
// in:  label = "next week"
// tpl: {{ label | format:'Y-m-d' }}
// out: next week
func Format(value any, params []any) (any, error) {
	switch timeValue := value.(type) {
	case string:
		return timeValue, nil
	case time.Time:
		d := date.NewFormatter(date.DefaultMapping)
		format := date.DefaultFormat
		if len(params) > 0 {
			s, ok := params[0].(string)
			if !ok {
				return nil, fmt.Errorf("format function expected string, got %T", params[0])
			}
			format = s
		}

		return d.Format(timeValue, format), nil
	}

	return nil, fmt.Errorf("format function expected string or time.Time, got %T", value)
}
