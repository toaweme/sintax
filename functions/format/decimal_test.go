package format

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_Decimal(t *testing.T) {
	decimal := decimalModifier
	tests := []struct {
		name     string
		value    any
		params   []any
		expected any
	}{
		{"two decimals", 19.5, []any{2}, "19.50"},
		{"round to one decimal", 4.872, []any{1}, "4.9"},
		{"default precision", 7, []any{}, "7.00"},
		{"zero places rounds", 3.14159, []any{0}, "3"},
		{"zero places rounds up", 3.6, []any{0}, "4"},
		{"int value", 42, []any{2}, "42.00"},
		{"int64 value", int64(5), []any{2}, "5.00"},
		{"float32 value", float32(1.5), []any{2}, "1.50"},
		{"string value", "3.14159", []any{2}, "3.14"},
		{"string value default", "8", []any{}, "8.00"},
		{"nil is zero", nil, []any{2}, "0.00"},
		{"negative value", -2.5, []any{1}, "-2.5"},
		{"int64 precision param", 12.3456, []any{int64(3)}, "12.346"},
		{"float precision param", 12.3456, []any{float64(2)}, "12.35"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := decimal(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

// Test_Decimal_BadPrecision asserts a non-numeric precision parameter is rejected.
func Test_Decimal_BadPrecision(t *testing.T) {
	decimal := decimalModifier
	_, err := decimal(1.0, []any{"2"})
	assert.Error(t, err)
}

// Test_Decimal_BadValue asserts an unparseable string value is rejected.
func Test_Decimal_BadValue(t *testing.T) {
	decimal := decimalModifier
	_, err := decimal("not a number", []any{2})
	assert.Error(t, err)

	_, err = decimal([]int{1, 2}, []any{2})
	assert.Error(t, err)
}
