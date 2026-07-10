package transform

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		expected any
	}{
		{
			name:     "slice of slices",
			value:    []any{[]any{"Mon", "Tue"}, []any{"Wed", "Thu"}},
			expected: []any{"Mon", "Tue", "Wed", "Thu"},
		},
		{
			name:     "plucked field lists",
			value:    []any{[]any{"coffee", "tea"}, []any{"cookie", "muffin"}},
			expected: []any{"coffee", "tea", "cookie", "muffin"},
		},
		{
			name:     "non-slice elements pass through",
			value:    []any{[]any{"a", "b"}, "c", []any{"d"}},
			expected: []any{"a", "b", "c", "d"},
		},
		{
			name:     "flat slice is unchanged",
			value:    []any{"a", "b", "c"},
			expected: []any{"a", "b", "c"},
		},
		{
			name:     "typed slice of ints",
			value:    []int{1, 2, 3},
			expected: []any{1, 2, 3},
		},
		{
			name:     "empty slice",
			value:    []any{},
			expected: []any{},
		},
		{
			name:     "nil elements are dropped",
			value:    []any{nil, "a", nil},
			expected: []any{"a"},
		},
		{
			name:     "only one level is flattened",
			value:    []any{[]any{[]any{"deep"}}, []any{"shallow"}},
			expected: []any{[]any{"deep"}, "shallow"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Flatten(tt.value, nil)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_Flatten_NilInput proves an untyped nil value is rejected: only typed nil
// pointers or interfaces short-circuit to an empty slice, a bare nil does not.
func Test_Flatten_NilInput(t *testing.T) {
	_, err := Flatten(nil, nil)
	assert.Error(t, err)
}

// Test_Flatten_WrongType proves a non-slice value is rejected.
func Test_Flatten_WrongType(t *testing.T) {
	_, err := Flatten("not a slice", nil)
	assert.Error(t, err)
}
