// Package format provides template modifiers that render a value as a
// human-readable string, covering dates, decimals, currency, line numbering,
// and size.
package format

import (
	"time"

	"github.com/toaweme/sintax/functions"
	"github.com/toaweme/sintax/internal/date"
)

// ModifierNameFormat is the template name for the Format modifier.
const ModifierNameFormat functions.ModifierName = "format"

// formatStringPassthrough returns a string value unchanged, ignoring any params,
// so format is safe to apply to a field that may already be formatted. It
// declines a non-string value so Overload falls through to the time clauses.
func formatStringPassthrough(value any, _ []any) (any, error) {
	if s, ok := value.(string); ok {
		return s, nil
	}
	return nil, functions.ErrInvalidValueType
}

// FormatTime renders a date/time value using a PHP-style date layout (Y is the
// 4-digit year, m the zero-padded month, d the day, H:i:s the time, and so on).
// Any character with no mapping is emitted literally.
func FormatTime(t time.Time, layout string) (string, error) {
	return date.NewFormatter(date.DefaultMapping).Format(t, layout), nil
}

// FormatTimeDefault renders a time.Time with the default "Y-m-d H:i:s" layout,
// the clause reached when no layout is given.
func FormatTimeDefault(t time.Time) (string, error) {
	return date.NewFormatter(date.DefaultMapping).Format(t, date.DefaultFormat), nil
}
