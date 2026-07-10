package access

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_Pluck(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		params   []any
		expected []any
	}{
		{
			name:     "collect ids",
			value:    []any{map[string]any{"id": 1, "name": "Alice"}, map[string]any{"id": 2, "name": "Bob"}},
			params:   []any{"id"},
			expected: []any{1, 2},
		},
		{
			name:     "gather names",
			value:    []any{map[string]any{"name": "Mug", "price": 12}, map[string]any{"name": "Pen", "price": 3}},
			params:   []any{"name"},
			expected: []any{"Mug", "Pen"},
		},
		{
			name:     "single element",
			value:    []any{map[string]any{"id": 99}},
			params:   []any{"id"},
			expected: []any{99},
		},
		{
			name:     "empty slice yields empty slice",
			value:    []any{},
			params:   []any{"id"},
			expected: []any{},
		},
		{
			name:     "typed nil slice yields empty slice",
			value:    []any(nil),
			params:   []any{"id"},
			expected: []any{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Pluck(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_Pluck_Errors(t *testing.T) {
	tests := []struct {
		name   string
		value  any
		params []any
	}{
		{"missing field param", []any{map[string]any{"id": 1}}, []any{}},
		{"non-string field param", []any{map[string]any{"id": 1}}, []any{42}},
		{"field absent from an element", []any{map[string]any{"id": 1}, map[string]any{"name": "x"}}, []any{"id"}},
		{"element is not a map", []any{map[string]any{"id": 1}, "scalar"}, []any{"id"}},
		{"value is not a slice", map[string]any{"id": 1}, []any{"id"}},
		{"nil value is not a slice", nil, []any{"id"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Pluck(tt.value, tt.params)
			assert.Error(t, err)
		})
	}
}
