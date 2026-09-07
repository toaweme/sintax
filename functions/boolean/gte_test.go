package boolean

import (
	"testing"
	"time"

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

// Test_GteTime asserts a date is compared as a date, where equality is the same
// instant rather than the same calendar day.
func Test_GteTime(t *testing.T) {
	tests := []struct {
		name     string
		value    time.Time
		than     time.Time
		expected bool
	}{
		{"later is on or after", later, moment, true},
		{"the same instant is on or after", moment, moment, true},
		{"earlier is not", earlier, moment, false},
		{"a second earlier is not", moment.Add(-time.Second), moment, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := GteTime(tt.value, tt.than)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_Gte_Dates_Modifier asserts the registered modifier reaches the date
// clause and still takes numbers.
func Test_Gte_Dates_Modifier(t *testing.T) {
	t.Run("two dates compare", func(t *testing.T) {
		out, err := gteModifier(moment, []any{moment})
		assert.NoError(t, err)
		assert.Equal(t, true, out)
	})
	t.Run("numbers still compare", func(t *testing.T) {
		out, err := gteModifier(1, []any{1})
		assert.NoError(t, err)
		assert.Equal(t, true, out)
	})
	t.Run("a date against a number", func(t *testing.T) {
		_, err := gteModifier(moment, []any{5})
		assert.ErrorIs(t, err, functions.ErrInvalidParamType)
	})
}
