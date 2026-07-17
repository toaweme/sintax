package parse

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_FromJSON(t *testing.T) {
	tests := []struct {
		name        string
		value       string
		expected    map[string]any
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
			name:        "top-level array is not an object",
			value:       `[1, 2, 3]`,
			expectedErr: true,
		},
		{
			name:        "malformed input",
			value:       `not json`,
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
