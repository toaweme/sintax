// Package datetime provides template modifiers that do arithmetic on a date,
// shifting one forward or back and measuring the distance between two. The
// format package writes a date out and reads one back in; this one moves it.
//
// Every modifier here takes a real time.Time and refuses a string, the same
// stance the date modifier takes, so a field that never became a date shows up
// where it goes wrong rather than being shifted by nothing. A printed date
// becomes a real one through from_date first.
package datetime

import (
	"time"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameAddDays is the template name for the AddDays modifier.
const ModifierNameAddDays functions.ModifierName = "add_days"

// AddDays shifts a date by a whole number of days, forward on a positive
// amount and back on a negative one, so `{{ now | add_days:14 | date:'Y-m-d' }}`
// is the date a fortnight from today and add_days:-7 is the date a week ago.
//
// The amount is signed rather than split across an add and a subtract pair,
// because the interesting amount is usually a field rather than a literal. A
// notice period read off a row is a number whose sign the template author does
// not know when they write the pipe, and with a pair they would have to.
//
// The clock time and the zone come through untouched, so a shifted timestamp
// still names the same moment of the day.
func AddDays(t time.Time, days int) (time.Time, error) {
	return t.AddDate(0, 0, days), nil
}
