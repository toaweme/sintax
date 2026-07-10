package transform

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

func Test_Sort(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		params   []any
		expected any
	}{
		{
			name:     "strings default ascending",
			value:    []any{"Charlie", "Alice", "Bob"},
			params:   nil,
			expected: []any{"Alice", "Bob", "Charlie"},
		},
		{
			name:     "strings explicit ascending",
			value:    []any{"Charlie", "Alice", "Bob"},
			params:   []any{"asc"},
			expected: []any{"Alice", "Bob", "Charlie"},
		},
		{
			name:     "integers descending",
			value:    []any{72, 95, 88},
			params:   []any{"desc"},
			expected: []any{95, 88, 72},
		},
		{
			name:     "floats ascending",
			value:    []any{9.99, 4.50, 14.00},
			params:   []any{"asc"},
			expected: []any{4.50, 9.99, 14.00},
		},
		{
			name:     "booleans false before true",
			value:    []any{true, false, true},
			params:   []any{"asc"},
			expected: []any{false, true, true},
		},
		{
			name:     "booleans descending true first",
			value:    []any{false, true, false},
			params:   []any{"desc"},
			expected: []any{true, false, false},
		},
		{
			name:     "mixed types grouped by type name",
			value:    []any{"b", 2, "a", 1},
			params:   []any{"asc"},
			expected: []any{1, 2, "a", "b"},
		},
		{
			name:     "empty slice",
			value:    []any{},
			params:   nil,
			expected: []any{},
		},
		{
			name:     "single element",
			value:    []any{"only"},
			params:   nil,
			expected: []any{"only"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Sort(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_Sort_NilInput proves a nil value passes through as nil rather than
// erroring, so an absent variable renders as nothing.
func Test_Sort_NilInput(t *testing.T) {
	out, err := Sort(nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, nil, out)
}

// Test_Sort_InvalidDirection proves a direction other than asc or desc is
// rejected.
func Test_Sort_InvalidDirection(t *testing.T) {
	_, err := Sort([]any{"a", "b"}, []any{"sideways"})
	assert.Error(t, err)
}

// Test_Sort_WrongType proves a non-slice value is rejected with the shared
// ErrInvalidValueType sentinel.
func Test_Sort_WrongType(t *testing.T) {
	_, err := Sort("not a slice", nil)
	assert.ErrorIs(t, err, functions.ErrInvalidValueType)
}
