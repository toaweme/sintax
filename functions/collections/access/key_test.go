package access

import (
	"testing"

	"github.com/toaweme/sintax/assert"
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

// Test_Key_ForgivingNil documents that Key never returns an error: every failing
// lookup renders as nil so templates can rely on the default modifier instead of
// erroring the whole render.
func Test_Key_ForgivingNil(t *testing.T) {
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
		{"no params", map[string]any{"a": 1}, []any{}},
		{"scalar value", 42, []any{"x"}},
		{"non-numeric slice index", []any{"a", "b"}, []any{"abc"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Key(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, nil, out)
		})
	}
}
