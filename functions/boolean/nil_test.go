package boolean

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

// A modifier that turns its value into a decision refuses nil, because a missing
// key and a key holding an explicit null both arrive as nil and neither one is an
// answer. The template says what absence means with a default ahead of the
// decision.
func Test_Boolean_RefusesNil(t *testing.T) {
	tests := []struct {
		name     string
		modifier functions.GlobalModifier
		params   []any
	}{
		{"not", notModifier, nil},
		{"and", andModifier, []any{true}},
		{"or", orModifier, []any{false}},
		{"eq", eqModifier, []any{"published"}},
		{"neq", neqModifier, []any{"published"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.modifier(nil, tt.params)
			assert.ErrorIs(t, err, functions.ErrNilDecision)
		})
	}
}

// Only the piped value is refused. A nil argument came from data the template
// named, so a real value is simply unequal to it.
func Test_Boolean_NilParamStillCompares(t *testing.T) {
	tests := []struct {
		name     string
		modifier functions.GlobalModifier
		expected any
	}{
		{"eq", eqModifier, false},
		{"neq", neqModifier, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := tt.modifier(0, []any{nil})
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}
