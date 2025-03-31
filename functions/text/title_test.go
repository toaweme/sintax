package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			result, err := Title(tt.input, nil)
			if tt.expectedErr != "" {
				assert.Equal(t, err.Error(), tt.expectedErr)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}
