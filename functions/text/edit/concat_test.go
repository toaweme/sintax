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
		// scalars stringify, in the value and in params
		{"number value", 42, []any{" items"}, "42 items"},
		{"number param", "id-", []any{7}, "id-7"},
		{"bool param", "active: ", []any{true}, "active: true"},
		{"mixed scalar parts", "v", []any{1, ".", 2}, "v1.2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := concat(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_Concat_Errors covers the text contract: scalars are accepted, but a
// composite value or param (a slice or map) is rejected.
func Test_Concat_Errors(t *testing.T) {
	concat := concatModifier
	tests := []struct {
		name    string
		value   any
		params  []any
		wrapped error
	}{
		{"composite value", []any{1, 2}, []any{"!"}, functions.ErrInvalidValueType},
		{"nil value", nil, []any{"!"}, functions.ErrInvalidValueType},
		{"composite param", "x", []any{[]any{1}}, functions.ErrInvalidParamType},
		{"composite param among scalars", "x", []any{"a", map[string]any{}, "b"}, functions.ErrInvalidParamType},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := concat(tt.value, tt.params)
			assert.ErrorIs(t, err, tt.wrapped)
		})
	}
}
