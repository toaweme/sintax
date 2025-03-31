package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			result, err := Slug(tt.input, nil)
			if tt.expectedErr != "" {
				assert.Equal(t, err.Error(), tt.expectedErr)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}
