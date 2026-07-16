package parse

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_FromCSV(t *testing.T) {
	tests := []struct {
		name        string
		value       string
		expected    []map[string]any
		expectedErr bool
	}{
		{
			name:  "header row keys each following row",
			value: "name,age\nAlice,30\nBob,25",
			expected: []map[string]any{
				{"name": "Alice", "age": "30"},
				{"name": "Bob", "age": "25"},
			},
		},
		{
			name:  "trailing newline and blank line are skipped",
			value: "name,age\nAlice,30\n\nBob,25\n",
			expected: []map[string]any{
				{"name": "Alice", "age": "30"},
				{"name": "Bob", "age": "25"},
			},
		},
		{
			name:     "only headers yields no rows",
			value:    "name,age\n",
			expected: []map[string]any{},
		},
		{
			name:  "shorter row pads with empty string",
			value: "a,b,c\n1,2",
			expected: []map[string]any{
				{"a": "1", "b": "2", "c": ""},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := FromCSV(tt.value)
			if tt.expectedErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
