package datetime

import (
	"time"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameDaysBetween is the template name for the DaysBetween modifier.
const ModifierNameDaysBetween functions.ModifierName = "days_between"

// DaysBetween counts the whole days from the value to the argument, positive
// when the argument is the later of the two and negative when it is earlier, so
// `{{ now | days_between:expires_at }}` is how many days are left.
//
// It answers with a number, which is what makes the question a template
// actually asks reachable. An argument is a name or a literal and never a pipe,
// so a date shifted by a modifier cannot itself be compared against, and
// "expiring within a fortnight" has to be written as a count instead of as two
// dates: `{{ now | days_between:expires_at | lte:14 }}`. The same number is what
// a notice reads out, "expires in 12 days".
//
// Both sides are read as calendar days in UTC, so a date and a timestamp on the
// same day are zero days apart however far into the day the timestamp sits, and
// midnight tonight is one day away rather than a fraction of one.
func DaysBetween(t time.Time, other time.Time) (int, error) {
	const day = 24 * time.Hour
	return int(startOfDayUTC(other).Sub(startOfDayUTC(t)) / day), nil
}

// startOfDayUTC drops a moment to midnight of the UTC day it falls in, so a
// difference of two of them is always a whole number of days.
func startOfDayUTC(t time.Time) time.Time {
	year, month, day := t.UTC().Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
