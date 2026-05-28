package convert

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_From(t *testing.T) {
	tests := []struct {
		name        string
		value       any
		params      []any
		expected    any
		expectedErr bool
	}{
		{
			name:   "json",
			value:  `{"key": "value", "number": 123, "float": 1.23}`,
			params: []any{"json"},
			expected: map[string]any{
				"key":    "value",
				"number": int64(123),
				"float":  1.23,
			},
		},
		{
			name:   "csv basic",
			value:  "name,age\nAlice,30\nBob,25",
			params: []any{"csv"},
			expected: []map[string]any{
				{"name": "Alice", "age": "30"},
				{"name": "Bob", "age": "25"},
			},
		},
		{
			name:   "csv with trailing newline and blank line",
			value:  "name,age\nAlice,30\n\nBob,25\n",
			params: []any{"csv"},
			expected: []map[string]any{
				{"name": "Alice", "age": "30"},
				{"name": "Bob", "age": "25"},
			},
		},
		{
			name:     "csv only headers",
			value:    "name,age\n",
			params:   []any{"csv"},
			expected: []map[string]any{},
		},
		{
			name:   "csv shorter row pads with empty string",
			value:  "a,b,c\n1,2",
			params: []any{"csv"},
			expected: []map[string]any{
				{"a": "1", "b": "2", "c": ""},
			},
		},
		{
			name:        "missing format",
			value:       "irrelevant",
			params:      []any{},
			expectedErr: true,
		},
		{
			name:        "unsupported format",
			value:       "irrelevant",
			params:      []any{"xml"},
			expectedErr: true,
		},
		{
			name:        "csv non-string input",
			value:       123,
			params:      []any{"csv"},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := From(tt.value, tt.params)
			if tt.expectedErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
