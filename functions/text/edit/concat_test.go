package edit

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

func Test_Concat(t *testing.T) {
	concat := concatModifier
	tests := []struct {
		name     string
		value    any
		params   []any
		expected string
	}{
		{"single suffix", "Hello", []any{"!"}, "Hello!"},
		{"multiple parts", "file", []any{"-", "01", ".txt"}, "file-01.txt"},
		{"join two words", "user", []any{"profile"}, "userprofile"},
		{"no params returns value", "solo", nil, "solo"},
		{"empty params slice", "solo", []any{}, "solo"},
		{"empty value with suffix", "", []any{"tail"}, "tail"},
		{"empty suffix", "head", []any{""}, "head"},
		{"all empty", "", []any{""}, ""},
		{"unicode parts", "café", []any{" ", "☕"}, "café ☕"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := concat(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_Concat_Errors covers the string-only contract: a non-string value or a
// non-string parameter is rejected rather than coerced.
func Test_Concat_Errors(t *testing.T) {
	concat := concatModifier
	tests := []struct {
		name    string
		value   any
		params  []any
		wrapped error
	}{
		{"non-string value", 42, []any{"!"}, functions.ErrInvalidValueType},
		{"nil value", nil, []any{"!"}, functions.ErrInvalidValueType},
		{"non-string param", "x", []any{7}, functions.ErrInvalidParamType},
		{"non-string param among strings", "x", []any{"a", true, "b"}, functions.ErrInvalidParamType},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := concat(tt.value, tt.params)
			assert.ErrorIs(t, err, tt.wrapped)
		})
	}
}
