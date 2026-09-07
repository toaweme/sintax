package datetime

import (
	"testing"
	"time"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

func Test_DaysBetween(t *testing.T) {
	tests := []struct {
		name     string
		value    time.Time
		other    time.Time
		expected int
	}{
		{"a fortnight ahead", moment, time.Date(2024, 3, 28, 9, 30, 5, 0, time.UTC), 14},
		{"a week behind", moment, time.Date(2024, 3, 7, 9, 30, 5, 0, time.UTC), -7},
		{"the same day", moment, moment, 0},
		{"later the same day is still zero", moment, time.Date(2024, 3, 14, 23, 59, 59, 0, time.UTC), 0},
		{"midnight tonight is one day", moment, time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC), 1},
		{"across a leap day", time.Date(2024, 2, 27, 0, 0, 0, 0, time.UTC), time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC), 3},
		{"across February in a common year", time.Date(2023, 2, 27, 0, 0, 0, 0, time.UTC), time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC), 2},
		{"across a year boundary", time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC), time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), 7},
		{"a whole leap year", time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), 366},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := DaysBetween(tt.value, tt.other)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_DaysBetween_AcrossZones asserts both sides are read as calendar days in
// UTC, so an offset that puts the two on the same UTC day counts as zero.
func Test_DaysBetween_AcrossZones(t *testing.T) {
	zone := time.FixedZone("CET", 2*60*60)
	out, err := DaysBetween(
		time.Date(2024, 3, 15, 1, 0, 0, 0, zone),
		time.Date(2024, 3, 14, 23, 0, 0, 0, time.UTC),
	)
	assert.NoError(t, err)
	assert.Equal(t, 0, out)
}

// Test_DaysBetween_Modifier asserts both sides have to be real dates.
func Test_DaysBetween_Modifier(t *testing.T) {
	tests := []struct {
		name   string
		value  any
		params []any
		target error
	}{
		{"a string value", "2024-03-14", []any{moment}, functions.ErrInvalidValueType},
		{"a string argument", moment, []any{"2024-03-28"}, functions.ErrInvalidParamType},
		{"a number argument", moment, []any{14}, functions.ErrInvalidParamType},
		{"no argument", moment, []any{}, functions.ErrMissingParam},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := daysBetweenModifier(tt.value, tt.params)
			assert.ErrorIs(t, err, tt.target)
			assert.Equal(t, nil, out)
		})
	}

	t.Run("counts through the registered modifier", func(t *testing.T) {
		out, err := daysBetweenModifier(moment, []any{time.Date(2024, 3, 26, 0, 0, 0, 0, time.UTC)})
		assert.NoError(t, err)
		assert.Equal(t, 12, out)
	})
}
