package splitjoin

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_Split(t *testing.T) {
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
			out, err := Split(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_Split_Errors(t *testing.T) {
	tests := []struct {
		name   string
		value  any
		params []any
	}{
		{"no params", "a,b,c", nil},
		{"empty params", "a,b,c", []any{}},
		{"non-string value", 42, []any{","}},
		{"nil value", nil, []any{","}},
		{"non-string separator", "a,b,c", []any{42}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Split(tt.value, tt.params)
			assert.Error(t, err)
		})
	}
}
