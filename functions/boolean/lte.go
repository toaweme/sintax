package boolean

import (
	"time"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameLte is the template name for the Lte modifier.
const ModifierNameLte functions.ModifierName = "lte"

// Lte reports whether the first value is at or below the second, comparing
// both as numbers so 5 and 5.0 rank the same. A non-numeric operand is an
// error, and nil counts as zero.
//
// It is the form a warning threshold is written in, since "expiring within a
// fortnight" is a count of days measured against a limit,
// `{{ now | days_between:expires_at | lte:14 }}`.
func Lte(value, than float64) (bool, error) {
	return value <= than, nil
}

// LteTime reports whether the value falls on or before the given date, the
// same stance LtTime takes on what counts as a date. Equality here is the same
// instant rather than the same day.
func LteTime(value, than time.Time) (bool, error) {
	return !value.After(than), nil
}
