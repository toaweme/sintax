package boolean

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

// Test_Or reads truthiness the same way a template `if` does, and always answers
// with a bool rather than with whichever operand was truthy.
func Test_Or(t *testing.T) {
	or := orModifier

	tests := []struct {
		name     string
		value    any
		params   []any
		expected any
	}{
		{"both true", true, []any{true}, true},
		{"value true", true, []any{false}, true},
		{"param true", false, []any{true}, true},
		{"both false", false, []any{false}, false},
		{"empty strings", "", []any{""}, false},
		{"a non-empty string answers for the pair", "yes", []any{""}, true},
		{"the literal false string is falsey", "false", []any{""}, false},
		{"zero and an empty slice", 0, []any{[]any{}}, false},
		{"a filled slice is truthy", []any{1}, []any{0}, true},
		{"nil value is falsey", nil, []any{false}, false},
		{"nil param is falsey", false, []any{nil}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := or(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_Or_MissingParam(t *testing.T) {
	or := orModifier
	_, err := or(true, nil)
	assert.ErrorIs(t, err, functions.ErrMissingParam)
}
