package format

import (
	"testing"
	"time"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
	"github.com/toaweme/sintax/internal/date"
)

func Test_FromDate(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		params   []any
		expected any
	}{
		{"long month name", "June 26, 2026", []any{"F j, Y"}, time.Date(2026, time.June, 26, 0, 0, 0, 0, time.UTC)},
		{"short month name", "Mar 14, 2024", []any{"M j, Y"}, time.Date(2024, time.March, 14, 0, 0, 0, 0, time.UTC)},
		{"iso date", "2024-03-14", []any{"Y-m-d"}, time.Date(2024, time.March, 14, 0, 0, 0, 0, time.UTC)},
		{"date and time", "14/03/2024 09:30", []any{"d/m/Y H:i"}, time.Date(2024, time.March, 14, 9, 30, 0, 0, time.UTC)},
		{"weekday in the value", "Thursday, March 14, 2024", []any{"l, F j, Y"}, time.Date(2024, time.March, 14, 0, 0, 0, 0, time.UTC)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := fromDateModifier(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

// Test_FromDate_Guesses asserts the no-layout clause reads the forms a printed
// date usually arrives in, and that a numeric value only one reading fits is
// read that way round.
func Test_FromDate_Guesses(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		expected any
	}{
		{"rfc 3339", "2026-06-26T14:30:00Z", time.Date(2026, time.June, 26, 14, 30, 0, 0, time.UTC)},
		{"iso date", "2026-06-26", time.Date(2026, time.June, 26, 0, 0, 0, 0, time.UTC)},
		{"default output layout", "2024-03-14 09:30:05", time.Date(2024, time.March, 14, 9, 30, 5, 0, time.UTC)},
		{"long month first", "June 26, 2026", time.Date(2026, time.June, 26, 0, 0, 0, 0, time.UTC)},
		{"long day first", "26 June 2026", time.Date(2026, time.June, 26, 0, 0, 0, 0, time.UTC)},
		{"short month name", "Jun 26, 2026", time.Date(2026, time.June, 26, 0, 0, 0, 0, time.UTC)},
		{"day beyond twelve reads day first", "25/12/2026", time.Date(2026, time.December, 25, 0, 0, 0, 0, time.UTC)},
		{"month first when the day cannot be", "12/25/2026", time.Date(2026, time.December, 25, 0, 0, 0, 0, time.UTC)},
		{"dotted day first", "25.12.2026", time.Date(2026, time.December, 25, 0, 0, 0, 0, time.UTC)},
		{"same either way round", "05/05/2026", time.Date(2026, time.May, 5, 0, 0, 0, 0, time.UTC)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := fromDateModifier(tt.value, []any{})
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

// Test_FromDate_GuessRefusesAmbiguous asserts a numeric value that reads as a
// real date both ways round is an error naming the value, not a coin flip that
// writes a wrong date with nothing to show for it.
func Test_FromDate_GuessRefusesAmbiguous(t *testing.T) {
	tests := []struct {
		name  string
		value any
	}{
		{"slashes", "03/04/2026"},
		{"dots", "03.04.2026"},
		{"hyphens", "03-04-2026"},
		{"unpadded", "3/4/2026"},
		{"with a time", "03/04/2026 09:30"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := fromDateModifier(tt.value, []any{})
			assert.ErrorIs(t, err, date.ErrAmbiguousDate)
			if functions.IsParamError(err) {
				t.Fatalf("IsParamError() = true, want false for an ambiguous value")
			}
		})
	}
}

// Test_FromDate_GuessUnknown asserts a value no known layout reads is an error
// that names the value, so the message is the thing that tells somebody the
// layout argument exists.
func Test_FromDate_GuessUnknown(t *testing.T) {
	tests := []struct {
		name  string
		value any
	}{
		{"prose", "next week"},
		{"empty value", ""},
		{"trailing text", "2024-03-14 leftovers"},
		{"impossible day", "2024-02-31"},
		{"two digit year", "14/03/24"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := fromDateModifier(tt.value, []any{})
			assert.ErrorIs(t, err, date.ErrUnknownLayout)
			if functions.IsParamError(err) {
				t.Fatalf("IsParamError() = true, want false for an unreadable value")
			}
		})
	}
}

// Test_FromDate_LayoutWinsOverGuess asserts a layout given in the template is
// honored exactly, including where the guesser would have read the same value
// the other way round.
func Test_FromDate_LayoutWinsOverGuess(t *testing.T) {
	guessed, err := fromDateModifier("12/25/2026", []any{})
	assert.NoError(t, err)
	assert.Equal(t, time.Date(2026, time.December, 25, 0, 0, 0, 0, time.UTC), guessed)

	told, err := fromDateModifier("03/04/2026", []any{"m/d/Y"})
	assert.NoError(t, err)
	assert.Equal(t, time.Date(2026, time.March, 4, 0, 0, 0, 0, time.UTC), told)

	other, err := fromDateModifier("03/04/2026", []any{"d/m/Y"})
	assert.NoError(t, err)
	assert.Equal(t, time.Date(2026, time.April, 3, 0, 0, 0, 0, time.UTC), other)
}

// Test_FromDate_BadValue asserts a value that does not match the layout it was
// given is an error the engine stops on, rather than a zero date traveling
// downstream as if the store had said January 1, year 1.
func Test_FromDate_BadValue(t *testing.T) {
	tests := []struct {
		name   string
		value  any
		params []any
	}{
		{"month name against a numeric layout", "June 26, 2026", []any{"Y-m-d"}},
		{"empty value", "", []any{"Y-m-d"}},
		{"trailing text", "2024-03-14 leftovers", []any{"Y-m-d"}},
		{"impossible day", "2024-02-31", []any{"Y-m-d"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := fromDateModifier(tt.value, tt.params)
			assert.Error(t, err)
			if functions.IsParamError(err) {
				t.Fatalf("IsParamError() = true, want false for a value that does not match its layout")
			}
		})
	}
}

// Test_FromDate_BadParams asserts a mistyped or repeated layout is reported as
// a param error, so a broken template cannot fall through to the guesser.
func Test_FromDate_BadParams(t *testing.T) {
	tests := []struct {
		name   string
		params []any
		target error
	}{
		{"numeric layout", []any{123}, functions.ErrInvalidParamType},
		{"two layouts", []any{"Y-m-d", "Y-m-d"}, functions.ErrInvalidParamType},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := fromDateModifier("2024-03-14", tt.params)
			assert.ErrorIs(t, err, tt.target)
			assert.True(t, functions.IsParamError(err))
		})
	}
}

// Test_FromDate_BadValueType asserts a value that is not a string is rejected,
// with and without a layout. A call carrying a layout reports the param
// rejection, since the guess clause takes no param and declines on that first.
func Test_FromDate_BadValueType(t *testing.T) {
	_, err := fromDateModifier(map[string]any{}, []any{"Y-m-d"})
	assert.ErrorIs(t, err, functions.ErrInvalidParamType)

	_, err = fromDateModifier(map[string]any{}, []any{})
	assert.ErrorIs(t, err, functions.ErrInvalidValueType)
}

// Test_FromDate_RoundTripsWithDate asserts the pair reads the same layout
// language, and that the two no-argument defaults meet, so a value written out
// by date comes back through from_date with nothing given either way.
func Test_FromDate_RoundTripsWithDate(t *testing.T) {
	layout := "F j, Y"
	rendered, err := dateModifier(moment, []any{layout})
	assert.NoError(t, err)

	parsed, err := fromDateModifier(rendered, []any{layout})
	assert.NoError(t, err)
	assert.Equal(t, time.Date(2024, time.March, 14, 0, 0, 0, 0, time.UTC), parsed)

	written, err := dateModifier(moment, []any{})
	assert.NoError(t, err)

	read, err := fromDateModifier(written, []any{})
	assert.NoError(t, err)
	assert.Equal(t, moment, read)
}
