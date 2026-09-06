package boolean

import (
	"time"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameLt is the template name for the Lt modifier.
const ModifierNameLt functions.ModifierName = "lt"

// Lt reports whether the first value is below the second, comparing both as
// numbers so 5 and 5.0 rank the same. A non-numeric operand is an error, and
// nil counts as zero.
//
// It is a name of its own rather than gte inverted with not, because the
// argument is what is fixed in a pipeline and the value is what flows, so the
// question "is this under the threshold" cannot be turned round the way it
// could if either side were free to move.
func Lt(value, than float64) (bool, error) {
	return value < than, nil
}

// LtTime reports whether the value falls before the given date, so
// `{{ starts_at | lt:now }}` asks whether a date has passed. Both sides have to
// be real dates, on the same terms GtTime states.
func LtTime(value, than time.Time) (bool, error) {
	return value.Before(than), nil
}
