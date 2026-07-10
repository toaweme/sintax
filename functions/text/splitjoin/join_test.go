package splitjoin

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_Join(t *testing.T) {
	join := joinModifier
	tests := []struct {
		name     string
		value    any
		params   []any
		expected string
	}{
		{"string slice with comma", []string{"coffee", "sale", "new"}, []any{","}, "coffee,sale,new"},
		{"any slice with comma", []any{"coffee", "sale", "new"}, []any{","}, "coffee,sale,new"},
		{"default separator is newline", []string{"a", "b", "c"}, nil, "a\nb\nc"},
		{"custom multi-char separator", []string{"sign in", "verify email"}, []any{" | "}, "sign in | verify email"},
		{"single element", []string{"only"}, []any{","}, "only"},
		{"empty string slice", []string{}, []any{","}, ""},
		{"empty any slice", []any{}, []any{","}, ""},
		{"single element default separator", []string{"solo"}, nil, "solo"},
		{"non-string separator falls back to newline", []string{"a", "b"}, []any{42}, "a\nb"},
		{"unicode separator", []string{"a", "b"}, []any{"☕"}, "a☕b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := join(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_Join_Errors(t *testing.T) {
	join := joinModifier
	tests := []struct {
		name   string
		value  any
		params []any
	}{
		{"non-slice value", "not a slice", []any{","}},
		{"nil value", nil, []any{","}},
		{"any slice with non-string item", []any{"ok", 42}, []any{","}},
		{"integer value", 42, []any{","}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := join(tt.value, tt.params)
			assert.Error(t, err)
		})
	}
}
