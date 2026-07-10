package control

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

// Test_Default covers the whole substitution rule. The fallback replaces only a
// nil value (absent or null variable) or an empty string. Every other value,
// including an empty list, an empty object, zero, and false, is a real value and
// passes through unchanged.
func Test_Default(t *testing.T) {
	def := defaultModifier
	fallback := "anonymous"

	tests := []struct {
		name     string
		value    any
		params   []any
		expected any
	}{
		{"nil uses fallback", nil, []any{fallback}, fallback},
		{"empty string uses fallback", "", []any{fallback}, fallback},
		{"nil uses zero fallback", nil, []any{0}, 0},
		{"nil uses empty-list fallback", nil, []any{[]any{}}, []any{}},
		{"nil uses empty-object fallback", nil, []any{map[string]any{}}, map[string]any{}},

		{"present string is kept", "Ada", []any{fallback}, "Ada"},
		{"whitespace string is kept", " ", []any{fallback}, " "},
		{"zero is a real value and kept", 0, []any{5}, 0},
		{"false is a real value and kept", false, []any{true}, false},
		{"empty slice is a real value and kept", []any{}, []any{fallback}, []any{}},
		{"empty map is a real value and kept", map[string]any{}, []any{fallback}, map[string]any{}},
		{"non-empty slice is kept", []any{1, 2}, []any{fallback}, []any{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := def(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_Default_MissingParam locks in the guard for a `default` with no fallback.
// It is a usage error, so the modifier returns ErrMissingParam rather than nil.
func Test_Default_MissingParam(t *testing.T) {
	def := defaultModifier
	_, err := def("value", nil)
	assert.ErrorIs(t, err, functions.ErrMissingParam)

	_, err = def(nil, []any{})
	assert.ErrorIs(t, err, functions.ErrMissingParam)
}

// Test_Default_EmptyCollectionFallback shows the real-world find + default pipe.
// find reports "no match" softly, the engine turns that into nil, and default
// substitutes an iterable stand-in so a downstream loop has something to range.
func Test_Default_EmptyCollectionFallback(t *testing.T) {
	def := defaultModifier
	out, err := def(nil, []any{[]any{}})
	assert.NoError(t, err)
	assert.Len(t, out, 0)

	out, err = def(nil, []any{map[string]any{}})
	assert.NoError(t, err)
	assert.Equal(t, map[string]any{}, out)
}
