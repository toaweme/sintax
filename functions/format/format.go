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

// Format formats a time.Time value using a date format string.
// Strings are passed through unchanged.
//
// value: string, time.Time
// param:0?: string
// returns: string
//
// example: render a date in ISO form
// in:  created_at = 2024-03-14T09:30:00Z
// tpl: {{ created_at | format:'Y-m-d' }}
// out: 2024-03-14
//
// example: render a date and time for display
// in:  published = 2024-03-14T09:30:00Z
// tpl: {{ published | format:'d/m/Y H:i' }}
// out: 14/03/2024 09:30
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
