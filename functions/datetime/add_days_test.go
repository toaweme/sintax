package datetime

import (
	"testing"
	"time"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

// moment is a fixed timestamp so every case is deterministic and never depends
// on the wall clock.
var moment = time.Date(2024, 3, 14, 9, 30, 5, 0, time.UTC)

func Test_AddDays(t *testing.T) {
	tests := []struct {
		name     string
		value    time.Time
		days     int
		expected time.Time
	}{
		{"forward a fortnight", moment, 14, time.Date(2024, 3, 28, 9, 30, 5, 0, time.UTC)},
		{"back a week", moment, -7, time.Date(2024, 3, 7, 9, 30, 5, 0, time.UTC)},
		{"zero is the same day", moment, 0, moment},
		{"over a month boundary", time.Date(2024, 1, 30, 0, 0, 0, 0, time.UTC), 3, time.Date(2024, 2, 2, 0, 0, 0, 0, time.UTC)},
		{"through a leap day", time.Date(2024, 2, 28, 0, 0, 0, 0, time.UTC), 1, time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)},
		{"a year with no leap day skips it", time.Date(2023, 2, 28, 0, 0, 0, 0, time.UTC), 1, time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC)},
		{"across a year boundary", time.Date(2024, 12, 30, 0, 0, 0, 0, time.UTC), 5, time.Date(2025, 1, 4, 0, 0, 0, 0, time.UTC)},
		{"back across a year boundary", time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC), -5, time.Date(2024, 12, 28, 0, 0, 0, 0, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := AddDays(tt.value, tt.days)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_AddDays_KeepsZone asserts the clock time and the zone survive a shift,
// so a shifted timestamp still names the same moment of the day.
func Test_AddDays_KeepsZone(t *testing.T) {
	zone := time.FixedZone("CET", 2*60*60)
	out, err := AddDays(time.Date(2024, 3, 14, 9, 30, 5, 0, zone), 1)
	assert.NoError(t, err)
	assert.Equal(t, time.Date(2024, 3, 15, 9, 30, 5, 0, zone), out)
}

// Test_AddDays_Modifier exercises the registered modifier, where WrapOne
// coerces the untyped value and param before AddDays runs.
func Test_AddDays_Modifier(t *testing.T) {
	tests := []struct {
		name   string
		value  any
		params []any
		target error
	}{
		{"a string is not a date", "2024-03-14", []any{14}, functions.ErrInvalidValueType},
		{"a number is not a date", 20240314, []any{14}, functions.ErrInvalidValueType},
		{"nil is not a date", nil, []any{14}, functions.ErrInvalidValueType},
		{"a written amount is not a number", moment, []any{"14"}, functions.ErrInvalidParamType},
		{"a fractional amount is not whole", moment, []any{1.5}, functions.ErrInvalidParamType},
		{"no amount at all", moment, []any{}, functions.ErrMissingParam},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := addDaysModifier(tt.value, tt.params)
			assert.ErrorIs(t, err, tt.target)
			assert.Equal(t, nil, out)
		})
	}

	t.Run("a whole float amount coerces", func(t *testing.T) {
		out, err := addDaysModifier(moment, []any{float64(14)})
		assert.NoError(t, err)
		assert.Equal(t, time.Date(2024, 3, 28, 9, 30, 5, 0, time.UTC), out)
	})
}
