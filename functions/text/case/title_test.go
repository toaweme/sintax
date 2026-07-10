package casing

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

func Test_Title(t *testing.T) {
	tests := []struct {
		title       string
		input       string
		expected    string
		expectedErr string
	}{
		{
			title:    "Test Slug",
			input:    "hello-world",
			expected: "Hello World",
		},
		{
			title:    "Test Slug with Acronym",
			input:    "gpt-4.5-preview",
			expected: "GPT 4.5 Preview",
		},
		{
			title:    "Test Slug with Special Characters",
			input:    "hello-world-2023",
			expected: "Hello World 2023",
		},
		{
			title:    "test acronyms",
			input:    "gpt-Oss-120b",
			expected: "GPT OSS 120b",
		},
		{
			title:    "test acronyms",
			input:    "openai/gpt-oss-120b:exacto",
			expected: "Openai/gpt OSS 120b:exacto",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			result, err := Title(tt.input, nil)
			if tt.expectedErr != "" {
				assert.Equal(t, tt.expectedErr, err.Error())
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Test_Title_Params covers the optional extra-acronym parameters, which are
// merged with the built-in acronym list before title-casing.
func Test_Title_Params(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		params   []any
		expected string
	}{
		{
			name:     "extra acronyms uppercased",
			input:    "seo-and-cta-tips",
			params:   []any{"seo", "cta"},
			expected: "SEO And CTA Tips",
		},
		{
			name:     "built-in acronyms still apply with params",
			input:    "the-api-and-seo-guide",
			params:   []any{"seo"},
			expected: "The API And SEO Guide",
		},
		{
			name:     "params are case-insensitive",
			input:    "my-cta-button",
			params:   []any{"CTA"},
			expected: "My CTA Button",
		},
		{
			name:     "empty input",
			input:    "",
			params:   nil,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Title(tt.input, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Test_Title_NonString proves Title rejects non-string values with the shared
// ErrInvalidValueType sentinel.
func Test_Title_NonString(t *testing.T) {
	for _, v := range []any{42, 3.14, true, nil, []int{1, 2}} {
		_, err := Title(v, nil)
		assert.ErrorIs(t, err, functions.ErrInvalidValueType)
	}
}
