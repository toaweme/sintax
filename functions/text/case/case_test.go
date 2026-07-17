package casing

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
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
			result, err := ToLower(tt.input)
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
			result, err := ToUpper(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Test_ToLower_TextLenient proves the registered lower modifier (wrapped in
// AsText) stringifies a scalar value and rejects only a composite or nil.
func Test_ToLower_TextLenient(t *testing.T) {
	lower := lowerModifier
	for _, v := range []any{42, 3.14, true} {
		if _, err := lower(v, nil); err != nil {
			t.Fatalf("expected scalar %v accepted, got %v", v, err)
		}
	}
	for _, v := range []any{nil, []int{1, 2}, map[string]any{"a": 1}} {
		_, err := lower(v, nil)
		assert.ErrorIs(t, err, functions.ErrInvalidValueType)
	}
}

// Test_ToUpper_TextLenient proves the registered upper modifier stringifies a
// scalar value and rejects only a composite or nil.
func Test_ToUpper_TextLenient(t *testing.T) {
	upper := upperModifier
	for _, v := range []any{42, 3.14, true} {
		if _, err := upper(v, nil); err != nil {
			t.Fatalf("expected scalar %v accepted, got %v", v, err)
		}
	}
	for _, v := range []any{nil, []int{1, 2}, map[string]any{"a": 1}} {
		_, err := upper(v, nil)
		assert.ErrorIs(t, err, functions.ErrInvalidValueType)
	}
}
