package boolean

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_Not(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		expected any
	}{
		{"true inverts to false", true, false},
		{"false inverts to true", false, true},
		{"positive number is truthy", 3, false},
		{"zero is falsey", 0, true},
		{"negative number is falsey", -1, true},
		{"positive float is truthy", 1.5, false},
		{"non-empty string is truthy", "hi", false},
		{"empty string is falsey", "", true},
		{`literal "false" string is falsey`, "false", true},
		{`literal "true" string is truthy`, "true", false},
		{"non-empty slice is truthy", []any{1}, false},
		{"empty slice is falsey", []any{}, true},
		{"non-empty map is truthy", map[string]any{"k": 1}, false},
		{"empty map is falsey", map[string]any{}, true},
		{"nil is falsey", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Not(tt.value, nil)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}
