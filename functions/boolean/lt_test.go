package boolean

import (
	"testing"
	"time"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

// moment is a fixed timestamp so every date case is deterministic and never
// depends on the wall clock.
var (
	moment  = time.Date(2024, 3, 14, 9, 30, 5, 0, time.UTC)
	earlier = time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
	later   = time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC)
)

func Test_Lt(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		than     float64
		expected bool
	}{
		{"below", 2, 5, true},
		{"equal is not below", 3, 3, false},
		{"above", 5, 2, false},
		{"float under threshold", 49.99, 50, true},
		{"negative below zero", -1, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Lt(tt.value, tt.than)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_LtTime(t *testing.T) {
	tests := []struct {
		name     string
		value    time.Time
		than     time.Time
		expected bool
	}{
		{"earlier is before", earlier, moment, true},
		{"later is not before", later, moment, false},
		{"the same instant is not before", moment, moment, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := LtTime(tt.value, tt.than)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_Lt_Modifier exercises the registered modifier over both clauses.
func Test_Lt_Modifier(t *testing.T) {
	t.Run("numbers compare", func(t *testing.T) {
		out, err := ltModifier(2, []any{5})
		assert.NoError(t, err)
		assert.Equal(t, true, out)
	})
	t.Run("dates compare", func(t *testing.T) {
		out, err := ltModifier(earlier, []any{moment})
		assert.NoError(t, err)
		assert.Equal(t, true, out)
	})
	t.Run("nil counts as zero", func(t *testing.T) {
		out, err := ltModifier(nil, []any{1})
		assert.NoError(t, err)
		assert.Equal(t, true, out)
	})
	t.Run("missing param", func(t *testing.T) {
		_, err := ltModifier(3, nil)
		assert.ErrorIs(t, err, functions.ErrMissingParam)
	})
	t.Run("a printed date is refused", func(t *testing.T) {
		_, err := ltModifier("2024-03-01", []any{moment})
		assert.ErrorIs(t, err, functions.ErrInvalidValueType)
	})
}
