package parse

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_FromJSON(t *testing.T) {
	tests := []struct {
		name        string
		value       string
		expected    any
		expectedErr bool
	}{
		{
			name:  "object with native numbers",
			value: `{"key": "value", "number": 123, "float": 1.23}`,
			expected: map[string]any{
				"key":    "value",
				"number": int64(123),
				"float":  1.23,
			},
		},
		{
			name:  "nested map and array preserve native numbers",
			value: `{"user": {"id": 7}, "scores": [1, 2.5]}`,
			expected: map[string]any{
				"user":   map[string]any{"id": int64(7)},
				"scores": []any{int64(1), 2.5},
			},
		},
		{
			name:     "empty object",
			value:    `{}`,
			expected: map[string]any{},
		},
		{
			name:  "array of objects",
			value: `[{"id": 1, "name": "a"}, {"id": 2, "name": "b"}]`,
			expected: []any{
				map[string]any{"id": int64(1), "name": "a"},
				map[string]any{"id": int64(2), "name": "b"},
			},
		},
		{
			name:     "array of numbers",
			value:    `[1, 2.5, 3]`,
			expected: []any{int64(1), 2.5, int64(3)},
		},
		{
			name:     "string",
			value:    `"hello"`,
			expected: "hello",
		},
		{
			name:     "integer",
			value:    `42`,
			expected: int64(42),
		},
		{
			name:     "float",
			value:    `4.2`,
			expected: 4.2,
		},
		{
			name:     "bool",
			value:    `true`,
			expected: true,
		},
		{
			name:     "null",
			value:    `null`,
			expected: nil,
		},
		{
			name:        "malformed input",
			value:       `not json`,
			expectedErr: true,
		},
		{
			name:        "unterminated array",
			value:       `[1, 2`,
			expectedErr: true,
		},
		{
			name:        "empty string",
			value:       "",
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := FromJSON(tt.value)
			if tt.expectedErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
