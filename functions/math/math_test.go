package math

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

func Test_Arithmetic(t *testing.T) {
	tests := []struct {
		name     string
		op       func(any, any) (float64, error)
		value    any
		param    any
		expected float64
	}{
		{"add two ints", Add, 2, 3, 5},
		{"add a float", Add, 2.5, 0.25, 2.75},
		{"add a negative subtracts", Add, 10, -4, 6},
		{"add a numeric string", Add, "12.50", 1.5, 14},
		{"add to nothing", Add, nil, 3, 3},
		{"subtract", Subtract, 10, 4, 6},
		{"subtract past zero", Subtract, 4, 10, -6},
		{"subtract a numeric string", Subtract, "10", "2.5", 7.5},
		{"multiply", Multiply, 3, 4, 12},
		{"multiply by a rate", Multiply, 100, 1.21, 121},
		{"multiply by zero", Multiply, 3, 0, 0},
		{"divide", Divide, 12, 4, 3},
		{"divide into a fraction", Divide, 1, 8, 0.125},
		{"divide a numeric string", Divide, "10", 4, 2.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := tt.op(tt.value, tt.param)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_Divide_ByZero asserts an empty count is an error rather than an
// infinity, since a template printing "+Inf" reports absent data as a result.
func Test_Divide_ByZero(t *testing.T) {
	tests := []struct {
		name    string
		divisor any
	}{
		{"a zero literal", 0},
		{"a zero float", 0.0},
		{"a zero string", "0"},
		{"nothing at all", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Divide(10, tt.divisor)
			assert.ErrorIs(t, err, functions.ErrInvalidParamValue)
		})
	}
}

// Test_Arithmetic_RefusesNonNumbers asserts a value or an argument that is not
// a number is refused rather than counted as zero, and that the two rejections
// stay apart, since only a bad argument says the template itself is wrong.
func Test_Arithmetic_RefusesNonNumbers(t *testing.T) {
	tests := []struct {
		name    string
		op      func(any, any) (float64, error)
		value   any
		param   any
		target  error
		isParam bool
	}{
		{"a word as the value", Add, "abc", 1, functions.ErrInvalidValueType, false},
		{"a map as the value", Multiply, map[string]any{}, 2, functions.ErrInvalidValueType, false},
		{"a slice as the value", Subtract, []any{1, 2}, 2, functions.ErrInvalidValueType, false},
		{"a word as the argument", Add, 1, "abc", functions.ErrInvalidParamType, true},
		{"a map as the argument", Divide, 1, map[string]any{}, functions.ErrInvalidParamType, true},
		{"a bool as the argument", Multiply, 1, true, functions.ErrInvalidParamType, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := tt.op(tt.value, tt.param)
			assert.ErrorIs(t, err, tt.target)
			assert.Equal(t, float64(0), out)
			if functions.IsParamError(err) != tt.isParam {
				t.Fatalf("IsParamError(err) = %v, want %v", functions.IsParamError(err), tt.isParam)
			}
		})
	}
}

// Test_Arithmetic_Modifier exercises the registered modifiers, where WrapOne
// hands the untyped value and argument through.
func Test_Arithmetic_Modifier(t *testing.T) {
	t.Run("adds through the registered modifier", func(t *testing.T) {
		out, err := addModifier(2, []any{3})
		assert.NoError(t, err)
		assert.Equal(t, float64(5), out)
	})
	t.Run("missing argument", func(t *testing.T) {
		for name, modifier := range Modifiers() {
			_, err := modifier(2, nil)
			assert.ErrorIs(t, err, functions.ErrMissingParam)
			if err == nil {
				t.Fatalf("%s took no argument", name)
			}
		}
	})
	t.Run("too many arguments", func(t *testing.T) {
		_, err := multiplyModifier(2, []any{3, 4})
		assert.ErrorIs(t, err, functions.ErrInvalidParamType)
	})
}
