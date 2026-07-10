package transform

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

func Test_Map(t *testing.T) {
	users := []map[string]any{
		{"id": "u1", "name": "Alice"},
		{"id": "u2", "name": "Bob"},
	}
	tests := []struct {
		name     string
		value    any
		params   []any
		expected any
	}{
		{
			name:   "index by id",
			value:  users,
			params: []any{"id"},
			expected: map[string]map[string]any{
				"u1": {"id": "u1", "name": "Alice"},
				"u2": {"id": "u2", "name": "Bob"},
			},
		},
		{
			name: "index by slug",
			value: []map[string]any{
				{"slug": "coffee", "label": "Coffee"},
				{"slug": "tea", "label": "Tea"},
			},
			params: []any{"slug"},
			expected: map[string]map[string]any{
				"coffee": {"slug": "coffee", "label": "Coffee"},
				"tea":    {"slug": "tea", "label": "Tea"},
			},
		},
		{
			name:     "empty slice yields empty map",
			value:    []map[string]any{},
			params:   []any{"id"},
			expected: map[string]map[string]any{},
		},
		{
			name:   "elements missing the key are skipped",
			value:  []map[string]any{{"id": "u1"}, {"name": "no-id"}},
			params: []any{"id"},
			expected: map[string]map[string]any{
				"u1": {"id": "u1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Map(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_Map_MissingParam proves the field name is required.
func Test_Map_MissingParam(t *testing.T) {
	_, err := Map([]map[string]any{{"id": "u1"}}, nil)
	assert.Error(t, err)
}

// Test_Map_NonStringParam proves a non-string field parameter is rejected with
// the shared ErrInvalidParamType sentinel.
func Test_Map_NonStringParam(t *testing.T) {
	_, err := Map([]map[string]any{{"id": "u1"}}, []any{42})
	assert.ErrorIs(t, err, functions.ErrInvalidParamType)
}

// Test_Map_WrongValueType proves the value must be a slice of maps.
func Test_Map_WrongValueType(t *testing.T) {
	_, err := Map([]any{"not a map"}, []any{"id"})
	assert.Error(t, err)
}

// Test_Map_NonStringKeyValue proves the field being keyed on must hold a string.
func Test_Map_NonStringKeyValue(t *testing.T) {
	_, err := Map([]map[string]any{{"id": 1}}, []any{"id"})
	assert.Error(t, err)
}
