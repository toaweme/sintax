package format

import (
	"errors"
	"testing"
	"time"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

// moment is a fixed timestamp so every case is deterministic and never depends
// on the wall clock.
var moment = time.Date(2024, 3, 14, 9, 30, 5, 0, time.UTC)

func Test_Date(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		params   []any
		expected any
	}{
		{"iso date", moment, []any{"Y-m-d"}, "2024-03-14"},
		{"day slash date and time", moment, []any{"d/m/Y H:i"}, "14/03/2024 09:30"},
		{"long human date", moment, []any{"l, F j, Y"}, "Thursday, March 14, 2024"},
		{"time only", moment, []any{"H:i:s"}, "09:30:05"},
		{"default layout when no param", moment, []any{}, "2024-03-14 09:30:05"},
		{"literal separators pass through", moment, []any{"Y . H"}, "2024 . 09"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := dateModifier(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

// Test_Date_BadParamType asserts a non-string layout parameter is rejected.
func Test_Date_BadParamType(t *testing.T) {
	_, err := dateModifier(moment, []any{123})
	assert.Error(t, err)
}

// Test_Date_BadValueType asserts a value that is not a date is rejected, with
// or without a layout. A string in particular gets the same terminal error as
// any other wrong shape, so a field that never became a date shows up where it
// goes wrong rather than sliding through untouched.
//
// Which sentinel comes out is the arity of the call. Every clause declines a
// value it cannot take, and the layout clause is the only one that accepts a
// param, so a call carrying one reports the param rejection and a bare call
// reports the value rejection.
func Test_Date_BadValueType(t *testing.T) {
	tests := []struct {
		name   string
		value  any
		params []any
		target error
	}{
		{"number with a layout", 42, []any{"Y-m-d"}, functions.ErrInvalidParamType},
		{"number with no layout", 42, []any{}, functions.ErrInvalidValueType},
		{"string with a layout", "next week", []any{"Y-m-d"}, functions.ErrInvalidParamType},
		{"string with no layout", "next week", []any{}, functions.ErrInvalidValueType},
		{"empty string", "", []any{}, functions.ErrInvalidValueType},
		{"map", map[string]any{}, []any{"Y-m-d"}, functions.ErrInvalidParamType},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := dateModifier(tt.value, tt.params)
			assert.ErrorIs(t, err, tt.target)
			assert.Equal(t, nil, actual)
			if errors.Is(err, functions.ErrAllowsDefaultFunc) {
				t.Fatalf("errors.Is(err, ErrAllowsDefaultFunc) = true, want a terminal error for a value that is not a date")
			}
		})
	}
}
