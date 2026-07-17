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
		expected any
	}{
		{"single line", "hello", "0. hello"},
		{
			name:     "multiple lines",
			value:    "Buy milk\nWalk the dog\nPay rent",
			expected: "0. Buy milk\n1. Walk the dog\n2. Pay rent",
		},
		{
			name:     "trailing newline yields an empty numbered line",
			value:    "a\n",
			expected: "0. a\n1. ",
		},
		{"nil passes through as nil", nil, nil},
		{"empty string passes through as nil", "", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := lineNumbers(tt.value, nil)
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
