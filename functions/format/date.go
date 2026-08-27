// Package format provides template modifiers that move a value between its
// stored form and a human-readable string, covering dates, decimals, currency,
// line numbering, and size. Dates travel both ways, since date writes one out
// through a PHP-style layout and from_date reads one back in.
package format

import (
	"time"

	"github.com/toaweme/sintax/functions"
	"github.com/toaweme/sintax/internal/date"
)

// ModifierNameDate is the template name for the Date modifier.
const ModifierNameDate functions.ModifierName = "date"

// Date renders a date/time value using a PHP-style date layout (Y is the
// 4-digit year, m the zero-padded month, d the day, H:i:s the time, and so on).
// Any character with no mapping is emitted literally.
//
// The value has to be a date. A string is rejected rather than passed through,
// because a modifier named for dates that quietly returns whatever it was
// handed hides the field that never became one.
func Date(t time.Time, layout string) (string, error) {
	return date.NewFormatter(date.DefaultMapping).Format(t, layout), nil
}

// DateDefault renders a time.Time with the default "Y-m-d H:i:s" layout, the
// clause reached when no layout is given. from_date reads that same layout
// back without being told it, so a value survives the round trip with no
// argument on either side.
func DateDefault(t time.Time) (string, error) {
	return date.NewFormatter(date.DefaultMapping).Format(t, date.DefaultFormat), nil
}
