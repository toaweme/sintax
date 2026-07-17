package access

import (
	"errors"
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

func Test_Key(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		params   []any
		expected any
	}{
		{
			name:     "top level map key",
			value:    map[string]any{"name": "Alice", "email": "alice@example.com"},
			params:   []any{"name"},
			expected: "Alice",
		},
		{
			name:     "nested map with dot path",
			value:    map[string]any{"database": map[string]any{"host": "db.local", "port": 5432}},
			params:   []any{"database.host"},
			expected: "db.local",
		},
		{
			name:     "nested map returns non-string leaf",
			value:    map[string]any{"database": map[string]any{"host": "db.local", "port": 5432}},
			params:   []any{"database.port"},
			expected: 5432,
		},
		{
			name:     "slice index",
			value:    []any{"espresso", "latte", "macchiato"},
			params:   []any{0},
			expected: "espresso",
		},
		{
			name:     "slice last index",
			value:    []any{"espresso", "latte", "macchiato"},
			params:   []any{2},
			expected: "macchiato",
		},
		{
			name:     "slice index from string param",
			value:    []any{"a", "b", "c"},
			params:   []any{"1"},
			expected: "b",
		},
		{
			name:     "int keyed map",
			value:    map[int]any{1: "one", 2: "two"},
			params:   []any{"2"},
			expected: "two",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Key(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_Key_Miss documents the lookups that find nothing. Each reports a miss,
// which default can catch and an if reads as false, rather than rendering as a
// silent nil that hides a misspelled key.
func Test_Key_Miss(t *testing.T) {
	tests := []struct {
		name   string
		value  any
		params []any
	}{
		{"missing map key", map[string]any{"name": "Alice"}, []any{"phone"}},
		{"missing nested segment", map[string]any{"a": map[string]any{"b": 1}}, []any{"a.z"}},
		{"path into non-map leaf", map[string]any{"name": "Alice"}, []any{"name.first"}},
		{"index out of range", []any{"a", "b"}, []any{5}},
		{"negative index", []any{"a", "b"}, []any{-1}},
		{"nil value", nil, []any{"x"}},
		{"key cannot exist in an int keyed map", map[int]any{1: "one"}, []any{"abc"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Key(tt.value, tt.params)
			assert.ErrorIs(t, err, functions.ErrAllowsDefaultFunc)
		})
	}
}

// Test_Key_AuthorError documents the inputs that mean the template itself is
// wrong. These stay terminal, so no default rescues them. Silently rendering
// nothing would leave a broken template looking like absent data.
func Test_Key_AuthorError(t *testing.T) {
	tests := []struct {
		name    string
		value   any
		params  []any
		wantErr error
	}{
		{"no params", map[string]any{"a": 1}, []any{}, functions.ErrMissingParam},
		{"scalar value", 42, []any{"x"}, functions.ErrInvalidValueType},
		{"non-string map key", map[string]any{"a": 1}, []any{42}, functions.ErrInvalidParamType},
		{"non-numeric slice index", []any{"a", "b"}, []any{"abc"}, functions.ErrInvalidParamType},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Key(tt.value, tt.params)
			assert.ErrorIs(t, err, tt.wantErr)
			if errors.Is(err, functions.ErrAllowsDefaultFunc) {
				t.Fatalf("a broken template must not be catchable by default: %v", err)
			}
		})
	}
}
