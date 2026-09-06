package datetime

import (
	"testing"
	"time"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

func Test_AddMonths(t *testing.T) {
	tests := []struct {
		name     string
		value    time.Time
		months   int
		expected time.Time
	}{
		{"forward one month", moment, 1, time.Date(2024, 4, 14, 9, 30, 5, 0, time.UTC)},
		{"back one month", moment, -1, time.Date(2024, 2, 14, 9, 30, 5, 0, time.UTC)},
		{"zero is the same day", moment, 0, moment},
		{"forward over a year boundary", time.Date(2024, 11, 15, 0, 0, 0, 0, time.UTC), 3, time.Date(2025, 2, 15, 0, 0, 0, 0, time.UTC)},
		{"back over a year boundary", time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC), -1, time.Date(2023, 12, 15, 0, 0, 0, 0, time.UTC)},
		{"back a whole year", time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC), -12, time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC)},
		{"the 31st clamps to a short month", time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC), 1, time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)},
		{"the 31st clamps to February in a common year", time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC), 1, time.Date(2023, 2, 28, 0, 0, 0, 0, time.UTC)},
		{"the 31st clamps to a thirty day month", time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC), 1, time.Date(2024, 4, 30, 0, 0, 0, 0, time.UTC)},
		{"a day the target month has is untouched", time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC), 1, time.Date(2024, 2, 28, 0, 0, 0, 0, time.UTC)},
		{"back onto a short month clamps too", time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC), -1, time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)},
		{"a leap day back a year", time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC), -12, time.Date(2023, 2, 28, 0, 0, 0, 0, time.UTC)},
		{"far back over several years", time.Date(2024, 5, 15, 0, 0, 0, 0, time.UTC), -30, time.Date(2021, 11, 15, 0, 0, 0, 0, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := AddMonths(tt.value, tt.months)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_AddYears(t *testing.T) {
	tests := []struct {
		name     string
		value    time.Time
		years    int
		expected time.Time
	}{
		{"forward a warranty term", moment, 2, time.Date(2026, 3, 14, 9, 30, 5, 0, time.UTC)},
		{"back one year", moment, -1, time.Date(2023, 3, 14, 9, 30, 5, 0, time.UTC)},
		{"zero is the same day", moment, 0, moment},
		{"a leap day lands on the 28th", time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC), 1, time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC)},
		{"a leap day survives four years", time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC), 4, time.Date(2028, 2, 29, 0, 0, 0, 0, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := AddYears(tt.value, tt.years)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_AddMonths_KeepsZone asserts the clock time and the zone survive a shift.
func Test_AddMonths_KeepsZone(t *testing.T) {
	zone := time.FixedZone("CET", 2*60*60)
	out, err := AddMonths(time.Date(2024, 1, 31, 9, 30, 5, 7, zone), 1)
	assert.NoError(t, err)
	assert.Equal(t, time.Date(2024, 2, 29, 9, 30, 5, 7, zone), out)
}

// Test_AddMonths_Modifier asserts the registered modifiers refuse a value that
// is not a date and an amount that is not a whole number.
func Test_AddMonths_Modifier(t *testing.T) {
	tests := []struct {
		name     string
		modifier functions.GlobalModifier
		value    any
		params   []any
		target   error
	}{
		{"months on a string", addMonthsModifier, "2024-03-14", []any{1}, functions.ErrInvalidValueType},
		{"months with a written amount", addMonthsModifier, moment, []any{"1"}, functions.ErrInvalidParamType},
		{"months with no amount", addMonthsModifier, moment, []any{}, functions.ErrMissingParam},
		{"years on a string", addYearsModifier, "2024-03-14", []any{1}, functions.ErrInvalidValueType},
		{"years with a fractional amount", addYearsModifier, moment, []any{0.5}, functions.ErrInvalidParamType},
		{"years with no amount", addYearsModifier, moment, []any{}, functions.ErrMissingParam},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := tt.modifier(tt.value, tt.params)
			assert.ErrorIs(t, err, tt.target)
			assert.Equal(t, nil, out)
		})
	}
}
