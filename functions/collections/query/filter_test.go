package query

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

func Test_Filter(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		params   []any
		expected any
	}{
		{
			name: "keep only active items",
			value: []any{
				map[string]any{"name": "Coffee", "status": "active"},
				map[string]any{"name": "Tea", "status": "sold-out"},
			},
			params:   []any{"status", "active"},
			expected: []any{map[string]any{"name": "Coffee", "status": "active"}},
		},
		{
			name: "keep only admin users",
			value: []any{
				map[string]any{"name": "Alice", "role": "admin"},
				map[string]any{"name": "Bob", "role": "viewer"},
			},
			params:   []any{"role", "admin"},
			expected: []any{map[string]any{"name": "Alice", "role": "admin"}},
		},
		{
			name: "match multiple items",
			value: []any{
				map[string]any{"name": "Alice", "role": "admin"},
				map[string]any{"name": "Bob", "role": "viewer"},
				map[string]any{"name": "Carol", "role": "admin"},
			},
			params: []any{"role", "admin"},
			expected: []any{
				map[string]any{"name": "Alice", "role": "admin"},
				map[string]any{"name": "Carol", "role": "admin"},
			},
		},
		{
			name: "nested field via dot notation",
			value: []any{
				map[string]any{"title": "Hello", "meta": map[string]any{"published": true}},
				map[string]any{"title": "Draft", "meta": map[string]any{"published": false}},
			},
			params:   []any{"meta.published", true},
			expected: []any{map[string]any{"title": "Hello", "meta": map[string]any{"published": true}}},
		},
		{
			name: "numeric int matches int",
			value: []any{
				map[string]any{"name": "Mug", "price": 10},
				map[string]any{"name": "Cup", "price": 8},
			},
			params:   []any{"price", 10},
			expected: []any{map[string]any{"name": "Mug", "price": 10}},
		},
		{
			name: "numeric int value matches float search",
			value: []any{
				map[string]any{"name": "Mug", "price": 10},
				map[string]any{"name": "Cup", "price": 8},
			},
			params:   []any{"price", 10.0},
			expected: []any{map[string]any{"name": "Mug", "price": 10}},
		},
		{
			name: "boolean field",
			value: []any{
				map[string]any{"name": "Alice", "active": true},
				map[string]any{"name": "Bob", "active": false},
			},
			params:   []any{"active", false},
			expected: []any{map[string]any{"name": "Bob", "active": false}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Filter(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_Filter_Empty covers the cases where the result must come back empty
// rather than as an error: an empty input slice, no matching value, and a field
// that does not exist on any item.
func Test_Filter_Empty(t *testing.T) {
	tests := []struct {
		name   string
		value  any
		params []any
	}{
		{
			name:   "empty input slice",
			value:  []any{},
			params: []any{"role", "admin"},
		},
		{
			name: "no item matches",
			value: []any{
				map[string]any{"name": "Alice", "role": "admin"},
				map[string]any{"name": "Bob", "role": "viewer"},
			},
			params: []any{"role", "owner"},
		},
		{
			name: "field missing on every item",
			value: []any{
				map[string]any{"name": "Alice"},
				map[string]any{"name": "Bob"},
			},
			params: []any{"role", "admin"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Filter(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Empty(t, out)
		})
	}
}

// Test_Filter_Errors covers the input-validation error paths. A non-slice value
// yields ErrInvalidValueType, too few parameters yields a plain error, and a
// non-string key parameter yields ErrInvalidParamType.
func Test_Filter_Errors(t *testing.T) {
	t.Run("value is not a slice", func(t *testing.T) {
		_, err := Filter("not a slice", []any{"role", "admin"})
		assert.ErrorIs(t, err, functions.ErrInvalidValueType)
	})

	t.Run("nil value", func(t *testing.T) {
		_, err := Filter(nil, []any{"role", "admin"})
		assert.ErrorIs(t, err, functions.ErrInvalidValueType)
	})

	t.Run("missing both parameters", func(t *testing.T) {
		_, err := Filter([]any{}, nil)
		assert.Error(t, err)
	})

	t.Run("only one parameter", func(t *testing.T) {
		_, err := Filter([]any{}, []any{"role"})
		assert.Error(t, err)
	})

	t.Run("key parameter is not a string", func(t *testing.T) {
		_, err := Filter([]any{map[string]any{"role": "admin"}}, []any{5, "admin"})
		assert.ErrorIs(t, err, functions.ErrInvalidParamType)
	})
}
