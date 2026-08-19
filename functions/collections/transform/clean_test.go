package transform

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

// Test_Clean drives the registered modifier, since the nil clause is part of
// what clean promises and only the Overload reaches it.
func Test_Clean(t *testing.T) {
	clean := cleanModifier
	tests := []struct {
		name     string
		value    any
		expected any
	}{
		{
			name:     "nils are dropped",
			value:    []any{"a", nil, "b"},
			expected: []any{"a", "b"},
		},
		{
			name:     "empty strings are dropped",
			value:    []any{"a", "", "b"},
			expected: []any{"a", "b"},
		},
		{
			name:     "empty lists and objects are dropped",
			value:    []any{[]any{}, map[string]any{}, "a"},
			expected: []any{"a"},
		},
		{
			name:     "zero and false are real values and stay",
			value:    []any{0, false, ""},
			expected: []any{0, false},
		},
		{
			name:     "filled collections stay",
			value:    []any{[]any{1}, map[string]any{"a": 1}},
			expected: []any{[]any{1}, map[string]any{"a": 1}},
		},
		{
			name:     "order is kept",
			value:    []any{"", "c", nil, "a", "", "b"},
			expected: []any{"c", "a", "b"},
		},
		{
			name:     "everything empty leaves an empty list",
			value:    []any{nil, "", []any{}},
			expected: []any{},
		},
		{
			name:     "an already clean list is unchanged",
			value:    []any{"a", "b"},
			expected: []any{"a", "b"},
		},
		{
			name:     "a typed slice is cleaned like any other",
			value:    []string{"a", "", "b"},
			expected: []any{"a", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := clean(tt.value, nil)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// A list that is not there at all is absent data rather than a broken template,
// so clean misses and a downstream default answers it.
func Test_Clean_NilIsAMiss(t *testing.T) {
	clean := cleanModifier
	_, err := clean(nil, nil)
	assert.ErrorIs(t, err, functions.ErrAllowsDefaultFunc)
}

// Being handed something that is not a list at all is the template being wrong,
// which no default rescues.
func Test_Clean_NonSliceIsTerminal(t *testing.T) {
	clean := cleanModifier
	_, err := clean(map[string]any{"a": 1}, nil)
	assert.ErrorIs(t, err, functions.ErrInvalidValueType)
}
