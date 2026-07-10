package format

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_Length(t *testing.T) {
	length := lengthModifier
	tests := []struct {
		name     string
		value    any
		expected any
	}{
		{"ascii string", "Alice", 5},
		{"empty string", "", 0},
		{"multi-byte counts bytes", "café", 5},
		{"byte slice", []byte("hello"), 5},
		{"string slice", []string{"mug", "pen", "pad"}, 3},
		{"int slice", []int{1, 2, 3, 4}, 4},
		{"any slice", []any{1, "two", 3.0}, 3},
		{"array", [2]int{1, 2}, 2},
		{"map", map[string]int{"a": 1, "b": 2}, 2},
		{"empty slice", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := length(tt.value, nil)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

// Test_Length_PointerAndInterface asserts the reflect path dereferences a
// pointer to a collection and counts the underlying elements.
func Test_Length_PointerAndInterface(t *testing.T) {
	length := lengthModifier
	s := []int{1, 2, 3}
	actual, err := length(&s, nil)
	assert.NoError(t, err)
	assert.Equal(t, 3, actual)

	var nilPtr *[]int
	actual, err = length(nilPtr, nil)
	assert.NoError(t, err)
	assert.Equal(t, 0, actual)
}

// Test_Length_Unsupported asserts a scalar with no meaningful length is rejected.
func Test_Length_Unsupported(t *testing.T) {
	length := lengthModifier
	_, err := length(42, nil)
	assert.Error(t, err)

	_, err = length(true, nil)
	assert.Error(t, err)

	// a bare nil (no concrete type) has no length and is rejected.
	_, err = length(nil, nil)
	assert.Error(t, err)
}
