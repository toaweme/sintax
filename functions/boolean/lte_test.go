package boolean

import (
	"testing"
	"time"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

func Test_Lte(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		than     float64
		expected bool
	}{
		{"below", 2, 5, true},
		{"equal is at the limit", 3, 3, true},
		{"above", 5, 2, false},
		{"a fortnight of warning", 12, 14, true},
		{"past the warning window", 20, 14, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Lte(tt.value, tt.than)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_LteTime(t *testing.T) {
	tests := []struct {
		name     string
		value    time.Time
		than     time.Time
		expected bool
	}{
		{"earlier is on or before", earlier, moment, true},
		{"the same instant is on or before", moment, moment, true},
		{"later is not", later, moment, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := LteTime(tt.value, tt.than)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_Lte_Modifier exercises the registered modifier over both clauses.
func Test_Lte_Modifier(t *testing.T) {
	t.Run("numbers compare", func(t *testing.T) {
		out, err := lteModifier(14, []any{14})
		assert.NoError(t, err)
		assert.Equal(t, true, out)
	})
	t.Run("dates compare", func(t *testing.T) {
		out, err := lteModifier(moment, []any{moment})
		assert.NoError(t, err)
		assert.Equal(t, true, out)
	})
	t.Run("non-numeric value", func(t *testing.T) {
		_, err := lteModifier("abc", []any{0})
		assert.ErrorIs(t, err, functions.ErrInvalidValueType)
	})
	t.Run("non-numeric param", func(t *testing.T) {
		_, err := lteModifier(3, []any{"abc"})
		assert.ErrorIs(t, err, functions.ErrInvalidParamType)
	})
}
