package format

import (
	"errors"
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

func Test_Currency(t *testing.T) {
	tests := []struct {
		name        string
		value       any
		params      []any
		expected    any
		expectedErr error
	}{
		{
			name:     "integer value",
			value:    100,
			params:   []any{1, 100},
			expected: int(10000),
		},
		{
			name:     "float value",
			value:    123.45,
			params:   []any{1, 100},
			expected: int(12345),
		},
		// string
		{
			name:     "string value",
			value:    "100",
			params:   []any{1, 100},
			expected: int(10000),
		},
		{
			name:     "dollars to cents",
			value:    9,
			params:   []any{1, 100},
			expected: int(900),
		},
		{
			name:     "cents to dollars truncates remainder",
			value:    1299,
			params:   []any{100, 1},
			expected: int(12),
		},
		{
			name:     "string with dollar symbol",
			value:    "$9.99",
			params:   []any{1, 100},
			expected: int(999),
		},
		{
			name:     "string with euro symbol",
			value:    "€5.50",
			params:   []any{1, 100},
			expected: int(550),
		},
		{
			name:        "missing params",
			value:       100,
			params:      []any{},
			expectedErr: errors.New("currency requires 2 parameters"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := Currency(tt.value, tt.params)
			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr.Error(), err.Error())
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

// Test_Currency_ParamType asserts a non-int parameter is rejected with the
// shared ErrInvalidParamType sentinel from ParamInt.
func Test_Currency_ParamType(t *testing.T) {
	_, err := Currency(100, []any{"1", 100})
	assert.ErrorIs(t, err, functions.ErrInvalidParamType)

	_, err = Currency(100, []any{1, "100"})
	assert.ErrorIs(t, err, functions.ErrInvalidParamType)
}

// Test_Currency_BadString asserts an unparseable string value surfaces an error
// rather than silently converting to zero.
func Test_Currency_BadString(t *testing.T) {
	_, err := Currency("not a price", []any{1, 100})
	assert.Error(t, err)
}
