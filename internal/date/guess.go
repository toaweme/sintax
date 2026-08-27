package date

import (
	"errors"
	"time"
)

var (
	// ErrUnknownLayout reports that a value is written in none of the layouts
	// Guess knows, so the caller has to say which layout it is written in.
	ErrUnknownLayout = errors.New("no built-in date layout matches")

	// ErrAmbiguousDate reports that a numeric value reads as a real date with
	// the day and the month either way round, and the two readings disagree.
	ErrAmbiguousDate = errors.New("day and month could be either way round")
)

// guessLayouts are the Go reference layouts Guess tries, in order, and that
// order is the whole contract. It runs from the machine forms nothing else can
// look like, through the written-out forms a month name pins down, to the
// numeric forms only one reading fits.
//
// They are Go layouts rather than the PHP-style layout language the rest of
// this package speaks, because that language has no escape and its T means a
// zone abbreviation, so it cannot write the literal T that separates the date
// and the time in ISO 8601.
var guessLayouts = []string{
	time.RFC3339Nano,
	time.RFC3339,
	"2006-01-02T15:04:05",
	"2006-01-02T15:04",
	"2006-01-02 15:04:05Z07:00",
	"2006-01-02 15:04:05",
	"2006-01-02 15:04",
	"2006-01-02",
	time.RFC1123Z,
	time.RFC1123,
	"Monday, January 2, 2006 15:04",
	"Monday, January 2, 2006",
	"January 2, 2006 15:04:05",
	"January 2, 2006 15:04",
	"January 2, 2006",
	"2 January 2006 15:04",
	"2 January 2006",
	"Jan 2, 2006 15:04",
	"Jan 2, 2006",
	"2 Jan 2006 15:04",
	"2 Jan 2006",
	"2006/01/02 15:04:05",
	"2006/01/02 15:04",
	"2006/01/02",
}

// ambiguousLayouts are the numeric forms whose first two fields could be the
// day and the month either way round. Each entry pairs the day-first reading
// with the month-first one, and both are unpadded so a single-digit field and
// a zero-padded one land on the same entry.
var ambiguousLayouts = [][2]string{
	{"2/1/2006 15:04:05", "1/2/2006 15:04:05"},
	{"2/1/2006 15:04", "1/2/2006 15:04"},
	{"2/1/2006", "1/2/2006"},
	{"2.1.2006 15:04", "1.2.2006 15:04"},
	{"2.1.2006", "1.2.2006"},
	{"2-1-2006 15:04", "1-2-2006 15:04"},
	{"2-1-2006", "1-2-2006"},
}

// Guess reads a date out of a value that does not say which layout it is
// written in, using loc for a value carrying no zone of its own.
//
// The layouts in guessLayouts are tried in their listed order and the first
// that parses the value whole wins. The numeric day-and-month forms come last
// and are handled as pairs, so 25/12/2026 reads as December 25 because only
// the day-first reading gives a real date, while 03/04/2026 reads two ways and
// is refused with ErrAmbiguousDate rather than guessed. A wrong-but-plausible
// date written silently is worse than an error naming the value, and the
// layout argument is right there to settle it.
func Guess(value string, loc *time.Location) (time.Time, error) {
	if loc == nil {
		loc = time.UTC
	}
	for _, layout := range guessLayouts {
		if parsed, err := time.ParseInLocation(layout, value, loc); err == nil {
			return parsed, nil
		}
	}
	for _, pair := range ambiguousLayouts {
		dayFirst, dayErr := time.ParseInLocation(pair[0], value, loc)
		monthFirst, monthErr := time.ParseInLocation(pair[1], value, loc)
		switch {
		case dayErr == nil && monthErr == nil:
			if dayFirst.Equal(monthFirst) {
				return dayFirst, nil
			}
			return time.Time{}, ErrAmbiguousDate
		case dayErr == nil:
			return dayFirst, nil
		case monthErr == nil:
			return monthFirst, nil
		}
	}
	return time.Time{}, ErrUnknownLayout
}
