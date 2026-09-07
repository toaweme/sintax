package boolean

import (
	"time"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameGt is the template name for the Gt modifier.
const ModifierNameGt functions.ModifierName = "gt"

// Gt reports whether the first value exceeds the second, comparing both as
// numbers so 5 and 5.0 rank the same. A non-numeric operand is an error, and nil counts as zero.
func Gt(value, than float64) (bool, error) {
	return value > than, nil
}

// GtTime reports whether the value falls after the given date, the gt clause
// reached when both sides are real dates, so `{{ expires_at | gt:now }}` asks
// whether a date is still ahead.
//
// Both sides have to be dates. A printed one becomes a real one through
// from_date, and it is not read here, because a modifier that took a string
// would have to guess how the value is written and would answer confidently
// either way. An argument is a name or a literal and never a pipe, so the
// other side of the comparison is a date the caller already resolved, and a
// date the template computes is compared through days_between instead.
func GtTime(value, than time.Time) (bool, error) {
	return value.After(than), nil
}
