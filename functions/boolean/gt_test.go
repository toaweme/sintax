package boolean

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

func Test_Gt(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		than     float64
		expected bool
	}{
		{"greater", 3, 0, true},
		{"equal is not greater", 3, 3, false},
		{"smaller", 2, 5, false},
		{"float below threshold", 49.99, 50, false},
		{"value beats float threshold", 91, 90.5, true},
		{"negative below zero", -1, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Gt(tt.value, tt.than)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_Gt_Modifier exercises the registered modifier, where WrapOne coerces the
// untyped value and param before Gt runs.
func Test_Gt_Modifier(t *testing.T) {
	gt := gtModifier

	t.Run("int operands coerce", func(t *testing.T) {
		out, err := gt(91, []any{90.5})
		assert.NoError(t, err)
		assert.Equal(t, true, out)
	})
	t.Run("nil counts as zero", func(t *testing.T) {
		out, err := gt(nil, []any{-1})
		assert.NoError(t, err)
		assert.Equal(t, true, out)
	})
	t.Run("nil is not greater than zero", func(t *testing.T) {
		out, err := gt(nil, []any{0})
		assert.NoError(t, err)
		assert.Equal(t, false, out)
	})
	t.Run("missing param", func(t *testing.T) {
		_, err := gt(3, nil)
		assert.ErrorIs(t, err, functions.ErrMissingParam)
	})
	t.Run("non-numeric value", func(t *testing.T) {
		_, err := gt("abc", []any{0})
		assert.ErrorIs(t, err, functions.ErrInvalidValueType)
	})
	t.Run("non-numeric param", func(t *testing.T) {
		_, err := gt(3, []any{"abc"})
		assert.ErrorIs(t, err, functions.ErrInvalidParamType)
	})
}
