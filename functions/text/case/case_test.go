package casing

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_ToLower(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "already lowercase", input: "hello", expected: "hello"},
		{name: "all uppercase", input: "HELLO", expected: "hello"},
		{name: "mixed case", input: "HeLLo World", expected: "hello world"},
		{name: "email normalization", input: "Alice@Example.COM", expected: "alice@example.com"},
		{name: "digits and punctuation untouched", input: "ABC-123!", expected: "abc-123!"},
		{name: "latin accents", input: "ÀÉÎ", expected: "àéî"},
		{name: "greek", input: "ΑΒΓ", expected: "αβγ"},
		{name: "cyrillic", input: "ПРИВЕТ", expected: "привет"},
		{name: "empty", input: "", expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ToLower(tt.input, nil)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func Test_ToUpper(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "already uppercase", input: "HELLO", expected: "HELLO"},
		{name: "all lowercase", input: "hello", expected: "HELLO"},
		{name: "mixed case", input: "HeLLo World", expected: "HELLO WORLD"},
		{name: "country code", input: "us", expected: "US"},
		{name: "digits and punctuation untouched", input: "abc-123!", expected: "ABC-123!"},
		{name: "latin accents", input: "àéî", expected: "ÀÉÎ"},
		{name: "greek", input: "αβγ", expected: "ΑΒΓ"},
		{name: "cyrillic", input: "привет", expected: "ПРИВЕТ"},
		{name: "empty", input: "", expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ToUpper(tt.input, nil)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Test_ToLower_NonString proves ToLower rejects non-string values with a
// descriptive error naming the offending Go type.
func Test_ToLower_NonString(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{name: "int", input: 42, expected: "tolower function expected a string, got int"},
		{name: "float", input: 3.14, expected: "tolower function expected a string, got float64"},
		{name: "bool", input: true, expected: "tolower function expected a string, got bool"},
		{name: "nil", input: nil, expected: "tolower function expected a string, got <nil>"},
		{name: "slice", input: []int{1, 2}, expected: "tolower function expected a string, got []int"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ToLower(tt.input, nil)
			assert.Error(t, err)
			assert.Equal(t, tt.expected, err.Error())
			assert.Equal(t, nil, result)
		})
	}
}

// Test_ToUpper_NonString proves ToUpper rejects non-string values with a
// descriptive error naming the offending Go type.
func Test_ToUpper_NonString(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{name: "int", input: 42, expected: "toupper function expected a string, got int"},
		{name: "float", input: 3.14, expected: "toupper function expected a string, got float64"},
		{name: "bool", input: true, expected: "toupper function expected a string, got bool"},
		{name: "nil", input: nil, expected: "toupper function expected a string, got <nil>"},
		{name: "slice", input: []int{1, 2}, expected: "toupper function expected a string, got []int"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ToUpper(tt.input, nil)
			assert.Error(t, err)
			assert.Equal(t, tt.expected, err.Error())
			assert.Equal(t, nil, result)
		})
	}
}
