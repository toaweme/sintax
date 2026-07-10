package transform

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_Sum(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		params   []any
		expected float64
	}{
		{
			name:     "floats without field",
			value:    []any{12.50, 8.00, 4.25},
			params:   nil,
			expected: 24.75,
		},
		{
			name:     "integers without field",
			value:    []any{1, 2, 3, 4},
			params:   nil,
			expected: 10,
		},
		{
			name:     "numeric strings are parsed",
			value:    []any{"1.5", "2.5"},
			params:   nil,
			expected: 4.0,
		},
		{
			name:     "nil elements count as zero",
			value:    []any{1, nil, 2},
			params:   nil,
			expected: 3,
		},
		{
			name:     "typed int slice",
			value:    []int{5, 10, 15},
			params:   nil,
			expected: 30,
		},
		{
			name:     "sum a field over maps",
			value:    []map[string]any{{"price": 12}, {"price": 3}, {"price": 5}},
			params:   []any{"price"},
			expected: 20,
		},
		{
			name:     "field mixes int and float",
			value:    []map[string]any{{"price": 12}, {"price": 3.5}},
			params:   []any{"price"},
			expected: 15.5,
		},
		{
			name:     "empty slice sums to zero",
			value:    []any{},
			params:   nil,
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Sum(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_Sum_NilInput proves an untyped nil value is rejected: only a typed nil
// pointer or interface short-circuits to zero, a bare nil does not.
func Test_Sum_NilInput(t *testing.T) {
	_, err := Sum(nil, nil)
	assert.Error(t, err)
}

// Test_Sum_WrongType proves a non-slice value is rejected.
func Test_Sum_WrongType(t *testing.T) {
	_, err := Sum("not a slice", nil)
	assert.Error(t, err)
}

// Test_Sum_NonNumericElement proves an element that is neither a number nor a
// numeric string is rejected.
func Test_Sum_NonNumericElement(t *testing.T) {
	_, err := Sum([]any{1, "abc"}, nil)
	assert.Error(t, err)
}

// Test_Sum_NonStringFieldParam proves the field parameter must be a string.
func Test_Sum_NonStringFieldParam(t *testing.T) {
	_, err := Sum([]map[string]any{{"price": 1}}, []any{42})
	assert.Error(t, err)
}

// Test_Sum_FieldNotFound proves summing a field absent from an element errors.
func Test_Sum_FieldNotFound(t *testing.T) {
	_, err := Sum([]map[string]any{{"price": 1}}, []any{"missing"})
	assert.Error(t, err)
}

// Test_Sum_FieldOnNonMap proves a field parameter over non-map elements errors.
func Test_Sum_FieldOnNonMap(t *testing.T) {
	_, err := Sum([]any{1, 2}, []any{"price"})
	assert.Error(t, err)
}
