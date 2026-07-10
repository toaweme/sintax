package edit

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_Concat(t *testing.T) {
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
			out, err := Concat(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_Concat_Errors covers the string-only contract: a non-string value or a
// non-string parameter is rejected rather than coerced.
func Test_Concat_Errors(t *testing.T) {
	tests := []struct {
		name   string
		value  any
		params []any
	}{
		{"non-string value", 42, []any{"!"}},
		{"nil value", nil, []any{"!"}},
		{"non-string param", "x", []any{7}},
		{"non-string param among strings", "x", []any{"a", true, "b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Concat(tt.value, tt.params)
			assert.Error(t, err)
		})
	}
}
