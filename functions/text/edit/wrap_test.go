package edit

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

func Test_Wrap(t *testing.T) {
	wrap := wrapModifier
	tests := []struct {
		name     string
		value    any
		params   []any
		expected map[string]any
	}{
		{"string under key", "Alice", []any{"user"}, map[string]any{"user": "Alice"}},
		{"list under key", []any{"mug", "pen", "pad"}, []any{"data"}, map[string]any{"data": []any{"mug", "pen", "pad"}}},
		{"number value", 42, []any{"count"}, map[string]any{"count": 42}},
		{"nil value", nil, []any{"empty"}, map[string]any{"empty": nil}},
		{"empty string key", "x", []any{""}, map[string]any{"": "x"}},
		{"non-string key coerces to empty", "x", []any{7}, map[string]any{"": "x"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := wrap(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_Wrap_Errors(t *testing.T) {
	wrap := wrapModifier
	t.Run("no params", func(t *testing.T) {
		_, err := wrap("x", nil)
		assert.ErrorIs(t, err, functions.ErrMissingParam)
	})
	t.Run("empty params slice", func(t *testing.T) {
		_, err := wrap("x", []any{})
		assert.ErrorIs(t, err, functions.ErrMissingParam)
	})
}
