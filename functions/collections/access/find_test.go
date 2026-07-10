package access

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

func Test_Find(t *testing.T) {
	find := findModifier
	users := []any{
		map[string]any{"id": 7, "name": "Bob"},
		map[string]any{"id": 42, "name": "Alice"},
	}
	tests := []struct {
		name     string
		value    any
		params   []any
		expected any
	}{
		{
			name:     "match int field in slice",
			value:    users,
			params:   []any{"id", 42},
			expected: map[string]any{"id": 42, "name": "Alice"},
		},
		{
			name:     "returns first match in order",
			value:    users,
			params:   []any{"id", 7},
			expected: map[string]any{"id": 7, "name": "Bob"},
		},
		{
			name:     "match string field",
			value:    []any{map[string]any{"slug": "hat"}, map[string]any{"slug": "mug"}},
			params:   []any{"slug", "mug"},
			expected: map[string]any{"slug": "mug"},
		},
		{
			name:     "value is a single matching map",
			value:    map[string]any{"id": 1, "name": "Solo"},
			params:   []any{"id", 1},
			expected: map[string]any{"id": 1, "name": "Solo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := find(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_Find_NotFound covers the not-found paths, which return the non-fatal
// sentinel so the default modifier can supply a fallback.
func Test_Find_NotFound(t *testing.T) {
	find := findModifier
	tests := []struct {
		name   string
		value  any
		params []any
	}{
		{"no matching element in slice", []any{map[string]any{"slug": "hat"}}, []any{"slug", "scarf"}},
		{"single map field does not match", map[string]any{"id": 1}, []any{"id", 2}},
		{"type mismatch string vs int", []any{map[string]any{"id": 42}}, []any{"id", "42"}},
		{"key absent from every element", []any{map[string]any{"id": 1}}, []any{"missing", 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := find(tt.value, tt.params)
			assert.ErrorIs(t, err, functions.ErrAllowsDefaultFunc)
		})
	}
}

func Test_Find_Errors(t *testing.T) {
	find := findModifier
	t.Run("fewer than two params", func(t *testing.T) {
		_, err := find([]any{map[string]any{"id": 1}}, []any{"id"})
		assert.ErrorIs(t, err, functions.ErrMissingParam)
	})
	t.Run("non-string key param", func(t *testing.T) {
		_, err := find([]any{map[string]any{"id": 1}}, []any{42, 1})
		assert.ErrorIs(t, err, functions.ErrInvalidParamType)
	})
	t.Run("value is neither slice nor map", func(t *testing.T) {
		_, err := find("scalar", []any{"id", 1})
		assert.ErrorIs(t, err, functions.ErrInvalidValueType)
	})
}
