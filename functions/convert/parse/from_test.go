package parse

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
			name:   "json nested map and array preserve native numbers",
			value:  `{"user": {"id": 7}, "scores": [1, 2.5]}`,
			params: []any{"json"},
			expected: map[string]any{
				"user":   map[string]any{"id": int64(7)},
				"scores": []any{int64(1), 2.5},
			},
		},
		{
			name:     "json empty object",
			value:    `{}`,
			params:   []any{"json"},
			expected: map[string]any{},
		},
		{
			name:        "json top-level array is not an object",
			value:       `[1, 2, 3]`,
			params:      []any{"json"},
			expectedErr: true,
		},
		{
			name:        "json malformed input",
			value:       `not json`,
			params:      []any{"json"},
			expectedErr: true,
		},
		{
			name:        "json empty string",
			value:       "",
			params:      []any{"json"},
			expectedErr: true,
		},
		{
			name:        "json non-string input",
			value:       123,
			params:      []any{"json"},
			expectedErr: true,
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
			name:        "unsupported format yaml",
			value:       "key: value",
			params:      []any{"yaml"},
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
