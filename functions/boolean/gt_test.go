package boolean

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_Gt(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		params   []any
		expected any
	}{
		{"greater int", 3, []any{0}, true},
		{"equal is not greater", 3, []any{3}, false},
		{"smaller int", 2, []any{5}, false},
		{"float below threshold", 49.99, []any{50}, false},
		{"int beats float threshold", 91, []any{90.5}, true},
		{"float beats int threshold", 90.6, []any{90}, true},
		{"negative below zero", -1, []any{0}, false},
		{"nil counts as zero", nil, []any{-1}, true},
		{"nil is not greater than zero", nil, []any{0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Gt(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_Gt_Errors(t *testing.T) {
	t.Run("missing param", func(t *testing.T) {
		out, err := Gt(3, nil)
		assert.Error(t, err)
		assert.Equal(t, nil, out)
		assert.Equal(t, "gt function requires at least one parameter", err.Error())
	})
	t.Run("non-numeric value", func(t *testing.T) {
		out, err := Gt("abc", []any{0})
		assert.Error(t, err)
		assert.Equal(t, nil, out)
		assert.Equal(t, "gt function expected a number, got string", err.Error())
	})
	t.Run("non-numeric param", func(t *testing.T) {
		out, err := Gt(3, []any{"abc"})
		assert.Error(t, err)
		assert.Equal(t, nil, out)
		assert.Equal(t, "gt function expected a number, got string", err.Error())
	})
}
