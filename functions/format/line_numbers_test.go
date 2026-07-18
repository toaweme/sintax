package format

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_LineNumbers(t *testing.T) {
	lineNumbers := lineNumbersModifier
	tests := []struct {
		name     string
		value    any
		params   []any
		expected any
	}{
		{name: "single line", value: "hello", expected: "1. hello"},
		{
			name:     "multiple lines",
			value:    "Buy milk\nWalk the dog\nPay rent",
			expected: "1. Buy milk\n2. Walk the dog\n3. Pay rent",
		},
		{
			name:     "trailing newline yields an empty numbered line",
			value:    "a\n",
			expected: "1. a\n2. ",
		},
		{
			name:     "custom start",
			value:    "Buy milk\nWalk the dog\nPay rent",
			params:   []any{6},
			expected: "6. Buy milk\n7. Walk the dog\n8. Pay rent",
		},
		{
			name:     "zero start",
			value:    "first\nsecond",
			params:   []any{0},
			expected: "0. first\n1. second",
		},
		{name: "nil passes through as nil", value: nil, expected: nil},
		{name: "empty string passes through as nil", value: "", expected: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := lineNumbers(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

// Test_LineNumbers_NonString asserts a non-string value is rejected.
func Test_LineNumbers_NonString(t *testing.T) {
	lineNumbers := lineNumbersModifier
	_, err := lineNumbers(42, nil)
	assert.Error(t, err)
}
