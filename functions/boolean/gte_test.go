package boolean

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_Gte(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		params   []any
		expected any
	}{
		{"equal passes", 1, []any{1}, true},
		{"greater passes", 91, []any{90}, true},
		{"below fails", 89, []any{90}, false},
		{"float equal passes", 2.5, []any{2.5}, true},
		{"int meets float threshold", 91, []any{90.0}, true},
		{"nil equals zero threshold", nil, []any{0}, true},
		{"nil below positive threshold", nil, []any{1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Gte(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_Gte_Errors(t *testing.T) {
	t.Run("missing param", func(t *testing.T) {
		out, err := Gte(1, nil)
		assert.Error(t, err)
		assert.Equal(t, nil, out)
		// note: the ad-hoc message reads "gt" due to a copy-paste in the source.
		assert.Equal(t, "gt function requires at least one parameter", err.Error())
	})
	t.Run("non-numeric value", func(t *testing.T) {
		out, err := Gte("abc", []any{0})
		assert.Error(t, err)
		assert.Equal(t, nil, out)
		assert.Equal(t, "gte function expected a number, got string", err.Error())
	})
	t.Run("non-numeric param", func(t *testing.T) {
		out, err := Gte(1, []any{"abc"})
		assert.Error(t, err)
		assert.Equal(t, nil, out)
		assert.Equal(t, "gte function expected a number, got string", err.Error())
	})
}
