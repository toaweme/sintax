package splitjoin

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

func Test_Split(t *testing.T) {
	split := splitModifier
	tests := []struct {
		name     string
		value    any
		params   []any
		expected []string
	}{
		{"csv fields", "Alice,42,admin", []any{","}, []string{"Alice", "42", "admin"}},
		{"path segments", "/var/log/app/server.log", []any{"/"}, []string{"", "var", "log", "app", "server.log"}},
		{"words on space", "coffee tea espresso", []any{" "}, []string{"coffee", "tea", "espresso"}},
		{"separator not present", "hello", []any{"|"}, []string{"hello"}},
		{"single element", "hello", []any{","}, []string{"hello"}},
		{"empty input", "", []any{","}, []string{""}},
		{"multi-char separator", "a::b::c", []any{"::"}, []string{"a", "b", "c"}},
		{"empty separator splits into runes", "abc", []any{""}, []string{"a", "b", "c"}},
		{"trailing separator yields empty tail", "a,b,", []any{","}, []string{"a", "b", ""}},
		{"unicode content", "café☕tea", []any{"☕"}, []string{"café", "tea"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := split(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_Split_Errors(t *testing.T) {
	split := splitModifier
	tests := []struct {
		name    string
		value   any
		params  []any
		wrapped error
	}{
		{"no params", "a,b,c", nil, functions.ErrMissingParam},
		{"empty params", "a,b,c", []any{}, functions.ErrMissingParam},
		{"non-string value", 42, []any{","}, functions.ErrInvalidValueType},
		{"nil value", nil, []any{","}, functions.ErrInvalidValueType},
		{"non-string separator", "a,b,c", []any{42}, functions.ErrInvalidParamType},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := split(tt.value, tt.params)
			assert.ErrorIs(t, err, tt.wrapped)
		})
	}
}
