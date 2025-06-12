package money

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Currency(t *testing.T) {
	tests := []struct {
		name        string
		value       any
		params      []any
		expected    float64
		expectedErr error
	}{
		{
			name:     "integer value",
			value:    100,
			params:   []any{1, 100},
			expected: 10000,
		},
		{
			name:     "float value",
			value:    123.45,
			params:   []any{1, 100},
			expected: 12345,
		},
		// string
		{
			name:     "string value",
			value:    "100",
			params:   []any{1, 100},
			expected: 10000,
		},
		{
			name:        "missing params",
			value:       100,
			params:      []any{},
			expectedErr: fmt.Errorf("currency requires 2 parameters"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := Currency(tt.value, tt.params)
			if tt.expectedErr != nil {
				assert.Equal(t, err.Error(), tt.expectedErr.Error())
				return
			}

			assert.NoError(t, err)
			assert.EqualValues(t, tt.expected, actual)
		})
	}
}
