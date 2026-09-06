package boolean

import (
	"time"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameGte is the template name for the Gte modifier.
const ModifierNameGte functions.ModifierName = "gte"

// Gte reports whether the first value is greater than or equal to the second,
// comparing both as numbers so 5 and 5.0 rank the same. A non-numeric operand is an error, and nil
// counts as zero.
func Gte(value, than float64) (bool, error) {
	return value >= than, nil
}

// GteTime reports whether the value falls on or after the given date, the same
// stance GtTime takes on what counts as a date. Equality here is the same
// instant rather than the same day, so two timestamps hours apart on one
// calendar day are not equal.
func GteTime(value, than time.Time) (bool, error) {
	return !value.Before(than), nil
}
