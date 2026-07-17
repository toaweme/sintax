package format

import (
	"testing"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

func Test_Currency(t *testing.T) {
	currency := currencyModifier
	tests := []struct {
		name     string
		value    any
		params   []any
		expected any
	}{
		{"integer value", 100, []any{1, 100}, int(10000)},
		{"float value", 123.45, []any{1, 100}, int(12345)},
		{"string value", "100", []any{1, 100}, int(10000)},
		{"dollars to cents", 9, []any{1, 100}, int(900)},
		{"cents to dollars truncates remainder", 1299, []any{100, 1}, int(12)},
		{"string with dollar symbol", "$9.99", []any{1, 100}, int(999)},
		{"string with euro symbol", "€5.50", []any{1, 100}, int(550)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := currency(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

// Test_Currency_MissingParams asserts both units are required.
func Test_Currency_MissingParams(t *testing.T) {
	currency := currencyModifier
	_, err := currency(100, []any{})
	assert.ErrorIs(t, err, functions.ErrMissingParam)
}

// Test_Currency_ParamType asserts a non-int parameter is rejected with the
// shared ErrInvalidParamType sentinel.
func Test_Currency_ParamType(t *testing.T) {
	currency := currencyModifier
	_, err := currency(100, []any{"1", 100})
	assert.ErrorIs(t, err, functions.ErrInvalidParamType)

	_, err = currency(100, []any{1, "100"})
	assert.ErrorIs(t, err, functions.ErrInvalidParamType)
}

// Test_Currency_BadString asserts an unparseable string value surfaces an error
// rather than silently converting to zero.
func Test_Currency_BadString(t *testing.T) {
	currency := currencyModifier
	_, err := currency("not a price", []any{1, 100})
	assert.Error(t, err)
}
