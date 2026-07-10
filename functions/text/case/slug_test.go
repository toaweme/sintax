package casing

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

func Test_Slug(t *testing.T) {
	tests := []struct {
		title       string
		input       string
		expected    string
		expectedErr string
	}{
		{
			title:    "Test Slug",
			input:    "Hello World",
			expected: "hello-world",
		},
		{
			title:    "Test Slug with Acronym",
			input:    "GPT-4.5 Preview",
			expected: "gpt-4.5-preview",
		},
		{
			title:    "Test Slug with Special Characters",
			input:    "Hello, World! @2023",
			expected: "hello-world-2023",
		},
		{
			title:    "keeps dots between digits",
			input:    "Version 1.2.3 Release",
			expected: "version-1.2.3-release",
		},
		{
			title:    "collapses extra whitespace and hyphens",
			input:    "  Hello -- World  ",
			expected: "hello-world",
		},
		{
			title:    "strips non-ASCII letters",
			input:    "Café Münchën",
			expected: "caf-m-nch-n",
		},
		{
			title:    "empty input",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			result, err := Slug(tt.input)
			if tt.expectedErr != "" {
				assert.Equal(t, tt.expectedErr, err.Error())
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Test_Slug_TextLenient proves the registered slug modifier (wrapped in AsText)
// stringifies a scalar value and rejects only a composite or nil.
func Test_Slug_TextLenient(t *testing.T) {
	slug := slugModifier
	for _, v := range []any{42, 3.14, true} {
		if _, err := slug(v, nil); err != nil {
			t.Fatalf("expected scalar %v accepted, got %v", v, err)
		}
	}
	for _, v := range []any{nil, []int{1, 2}, map[string]any{"a": 1}} {
		_, err := slug(v, nil)
		assert.ErrorIs(t, err, functions.ErrInvalidValueType)
	}
}
