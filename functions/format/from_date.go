package format

import (
	"fmt"
	"time"

	"github.com/toaweme/sintax/functions"
	"github.com/toaweme/sintax/internal/date"
)

// ModifierNameFromDate is the template name for the FromDate modifier.
const ModifierNameFromDate functions.ModifierName = "from_date"

// FromDate reads a date out of a string using the same PHP-style layout the
// date modifier writes with, so `June 26, 2026` piped through
// from_date:'F j, Y' becomes a real time.Time that a date column can take.
// Without it a template has to rewrite a month name by hand, which is a chain
// of a dozen replacements that no reader can check.
//
// The value has to match the layout in full, separators included, and any
// character the layout language does not name is matched literally. A value
// carrying its own offset keeps it, and anything else is read as UTC, so a
// date-only value lands on midnight UTC rather than on whatever zone the
// process happens to run in.
//
// A layout given here is honored exactly and never second-guessed. A value
// that does not match it is an error. Returning the zero time instead would
// send January 1, year 1 downstream as if the store had really said so.
//
// The layout is optional. With none given the value is read as it comes,
// through the forms a printed date usually arrives in, ISO 8601 and RFC first,
// then the ones a month name spells out, then the numeric ones. A numeric
// value whose day and month could be either way round, such as 03/04/2026, is
// refused rather than guessed, and a layout settles it.
func FromDate(value string, layout string) (time.Time, error) {
	return date.NewParser(date.DefaultMapping).Parse(value, layout, time.UTC)
}

// FromDateGuess reads a date out of a string with no layout given, which is
// the clause `{{ printed | from_date }}` reaches. It works through the layouts
// a printed date is usually written in, machine forms first, then the forms a
// month name spells out, then the numeric ones.
//
// A numeric value that reads as a real date with the day and the month either
// way round, such as 03/04/2026, is refused rather than guessed, since one of
// the two readings would put a wrong date in a column with nothing to show for
// it. The error says so and names the layout argument that settles it.
func FromDateGuess(value string) (time.Time, error) {
	parsed, err := date.Guess(value, time.UTC)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to read date %q without a layout (pass one such as 'd/m/Y'): %w", value, err)
	}
	return parsed, nil
}
