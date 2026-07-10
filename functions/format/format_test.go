package format

import (
	"testing"
	"time"

	"github.com/toaweme/sintax/assert"
)

// moment is a fixed timestamp so every case is deterministic and never depends
// on the wall clock.
var moment = time.Date(2024, 3, 14, 9, 30, 5, 0, time.UTC)

func Test_Format(t *testing.T) {
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
		{"default format when no param", moment, []any{}, "2024-03-14 09:30:05"},
		{"literal separators pass through", moment, []any{"Y . H"}, "2024 . 09"},
		{"string passes through unchanged", "next week", []any{"Y-m-d"}, "next week"},
		{"empty string passes through", "", []any{"Y-m-d"}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := Format(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

// Test_Format_BadParamType asserts a non-string format parameter is rejected.
func Test_Format_BadParamType(t *testing.T) {
	_, err := Format(moment, []any{123})
	assert.Error(t, err)
}

// Test_Format_BadValueType asserts a value that is neither a string nor a
// time.Time is rejected.
func Test_Format_BadValueType(t *testing.T) {
	_, err := Format(42, []any{"Y-m-d"})
	assert.Error(t, err)
}
