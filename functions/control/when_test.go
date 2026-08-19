package control

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

// Test_When drives the registered modifier, since the arity check and the param
// coercion are as much of when's contract as the choice itself.
func Test_When(t *testing.T) {
	when := whenModifier

	tests := []struct {
		name     string
		value    any
		params   []any
		expected any
	}{
		{"true takes the first branch", true, []any{"live", "draft"}, "live"},
		{"false takes the second branch", false, []any{"live", "draft"}, "draft"},
		{"the literal false string is falsey", "false", []any{"live", "draft"}, "draft"},
		{"the literal true string is truthy", "true", []any{"live", "draft"}, "live"},
		{"a non-empty string is truthy", "anything", []any{"live", "draft"}, "live"},
		{"an empty string is falsey", "", []any{"live", "draft"}, "draft"},
		{"zero is falsey", 0, []any{"live", "draft"}, "draft"},
		{"a positive number is truthy", 3, []any{"live", "draft"}, "live"},
		{"an empty slice is falsey", []any{}, []any{"live", "draft"}, "draft"},
		{"a filled slice is truthy", []any{1}, []any{"live", "draft"}, "live"},
		{"a miss arrives as nil and is falsey", nil, []any{"live", "draft"}, "draft"},
		{"branches keep their own types", true, []any{1, 0}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := when(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// A choice with one branch is a choice somebody forgot to finish, so when
// refuses it rather than treating the absent branch as an empty string.
func Test_When_RequiresBothBranches(t *testing.T) {
	when := whenModifier

	_, err := when(true, []any{"live"})
	assert.ErrorIs(t, err, functions.ErrMissingParam)

	_, err = when(true, nil)
	assert.ErrorIs(t, err, functions.ErrMissingParam)
}
