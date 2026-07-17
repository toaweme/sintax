package boolean

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

func Test_Gte(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		than     float64
		expected bool
	}{
		{"equal passes", 1, 1, true},
		{"greater passes", 91, 90, true},
		{"below fails", 89, 90, false},
		{"float equal passes", 2.5, 2.5, true},
		{"value meets float threshold", 91, 90.0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Gte(tt.value, tt.than)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_Gte_Modifier exercises the registered modifier. It also confirms the old
// "gt function requires..." copy-paste message is gone: a missing param now
// surfaces the ErrMissingParam sentinel from WrapOne.
func Test_Gte_Modifier(t *testing.T) {
	gte := gteModifier

	t.Run("nil equals zero threshold", func(t *testing.T) {
		out, err := gte(nil, []any{0})
		assert.NoError(t, err)
		assert.Equal(t, true, out)
	})
	t.Run("nil below positive threshold", func(t *testing.T) {
		out, err := gte(nil, []any{1})
		assert.NoError(t, err)
		assert.Equal(t, false, out)
	})
	t.Run("missing param", func(t *testing.T) {
		_, err := gte(1, nil)
		assert.ErrorIs(t, err, functions.ErrMissingParam)
	})
	t.Run("non-numeric value", func(t *testing.T) {
		_, err := gte("abc", []any{0})
		assert.ErrorIs(t, err, functions.ErrInvalidValueType)
	})
	t.Run("non-numeric param", func(t *testing.T) {
		_, err := gte(1, []any{"abc"})
		assert.ErrorIs(t, err, functions.ErrInvalidParamType)
	})
}
