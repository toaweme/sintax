package boolean

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

// Test_Neq drives the registered modifier, since neq's behavior is the Overload
// dispatch across its numeric, string, and any clauses (plus the nil guard). The
// cases mirror Test_Eq's so the two stay readable as each other's negation.
func Test_Neq(t *testing.T) {
	neq := neqModifier

	tests := []struct {
		name     string
		value    any
		params   []any
		expected any
	}{
		{"equal strings", "active", []any{"active"}, false},
		{"unequal strings", "active", []any{"archived"}, true},
		{"empty strings equal", "", []any{""}, false},
		{"equal ints", 3, []any{3}, false},
		{"unequal ints", 3, []any{4}, true},
		{"zero equals zero", 0, []any{0}, false},
		{"int equals float across kinds", 5, []any{5.0}, false},
		{"float equals int across kinds", 5.0, []any{5}, false},
		{"equal bools", true, []any{true}, false},
		{"unequal bools", true, []any{false}, true},
		{"number and its string form differ", 5, []any{"5"}, true},
		{"nil equals nil", nil, []any{nil}, false},
		{"nil differs from zero", nil, []any{0}, true},
		{"value differs from nil param", 0, []any{nil}, true},
		{"unicode strings equal", "café", []any{"café"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := neq(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_Neq_MissingParam(t *testing.T) {
	neq := neqModifier
	_, err := neq("active", nil)
	assert.ErrorIs(t, err, functions.ErrMissingParam)
}
