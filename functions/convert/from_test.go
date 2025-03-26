package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_From(t *testing.T) {
	tests := []struct {
		name        string
		value       any
		params      []any
		expected    map[string]any
		expectedErr error
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := From(tt.value, tt.params)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, tt.expectedErr, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
