package boolean

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

// Test_And reads truthiness the same way a template `if` does, so the interesting
// cases are the values that look empty but are not, and absent data.
func Test_And(t *testing.T) {
	and := andModifier

	tests := []struct {
		name     string
		value    any
		params   []any
		expected any
	}{
		{"both true", true, []any{true}, true},
		{"value false", false, []any{true}, false},
		{"param false", true, []any{false}, false},
		{"both false", false, []any{false}, false},
		{"non-empty strings", "yes", []any{"also"}, true},
		{"empty string is falsey", "", []any{"also"}, false},
		{"the literal false string is falsey", "false", []any{true}, false},
		{"zero is falsey", 0, []any{true}, false},
		{"a positive number is truthy", 2, []any{"yes"}, true},
		{"an empty slice is falsey", []any{}, []any{true}, false},
		{"a filled slice is truthy", []any{1}, []any{true}, true},
		{"nil param is falsey", true, []any{nil}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := and(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_And_MissingParam(t *testing.T) {
	and := andModifier
	_, err := and(true, nil)
	assert.ErrorIs(t, err, functions.ErrMissingParam)
}
