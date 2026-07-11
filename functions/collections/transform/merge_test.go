package transform

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

func Test_Merge(t *testing.T) {
	merge := mergeModifier
	tests := []struct {
		name     string
		value    any
		params   []any
		expected any
	}{
		{
			name: "index users by id",
			value: []map[string]any{
				{"id": "u1", "name": "Alice"},
				{"id": "u2", "name": "Bob"},
			},
			params: []any{"id"},
			expected: map[string]map[string]any{
				"u1": {"id": "u1", "name": "Alice"},
				"u2": {"id": "u2", "name": "Bob"},
			},
		},
		{
			name: "index records by name",
			value: []map[string]any{
				{"name": "draft", "value": 1},
				{"name": "final", "value": 2},
			},
			params: []any{"name"},
			expected: map[string]map[string]any{
				"draft": {"name": "draft", "value": 1},
				"final": {"name": "final", "value": 2},
			},
		},
		{
			name:     "empty slice yields empty map",
			value:    []map[string]any{},
			params:   []any{"id"},
			expected: map[string]map[string]any{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := merge(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_Merge_MatchesMap proves merge is a behavioral alias of map.
func Test_Merge_MatchesMap(t *testing.T) {
	value := []map[string]any{{"id": "a", "n": 1}, {"id": "b", "n": 2}}
	mapped, err := Map(value, "id")
	assert.NoError(t, err)
	merged, err := Merge(value, "id")
	assert.NoError(t, err)
	assert.Equal(t, mapped, merged)
}

// Test_Merge_MissingParam proves the field name is required.
func Test_Merge_MissingParam(t *testing.T) {
	merge := mergeModifier
	_, err := merge([]map[string]any{{"id": "u1"}}, nil)
	assert.ErrorIs(t, err, functions.ErrMissingParam)
}

// Test_Merge_NonStringParam proves a non-string field parameter is rejected.
func Test_Merge_NonStringParam(t *testing.T) {
	merge := mergeModifier
	_, err := merge([]map[string]any{{"id": "u1"}}, []any{42})
	assert.ErrorIs(t, err, functions.ErrInvalidParamType)
}

// Test_Merge_WrongValueType proves the value must be a slice of maps.
func Test_Merge_WrongValueType(t *testing.T) {
	merge := mergeModifier
	_, err := merge([]any{"not a map"}, []any{"id"})
	assert.ErrorIs(t, err, functions.ErrInvalidValueType)
}
