package edit

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_Shorten(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		params   []any
		expected string
	}{
		{"clip description", "Hand-picked single-origin coffee, slow-roasted in small batches.", []any{30}, "Hand-picked single-origin coff"},
		{"clip name", "Alexandra Christine Whitehead", []any{12}, "Alexandra Ch"},
		{"shorter than limit unchanged", "OK", []any{10}, "OK"},
		{"equal to limit unchanged", "hello", []any{5}, "hello"},
		{"zero length yields empty", "hello", []any{0}, ""},
		{"empty value", "", []any{5}, ""},
		{"numeric string param", "hello", []any{"3"}, "hel"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Shorten(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_Shorten_BytesNotRunes locks in the documented gotcha: the limit counts
// bytes, so cutting inside a multi-byte character leaves a broken trailing
// byte. "café" is 5 bytes (é is two), so a 4-byte cut yields "caf" plus the
// first byte of é.
func Test_Shorten_BytesNotRunes(t *testing.T) {
	out, err := Shorten("café", []any{4})
	assert.NoError(t, err)
	assert.Equal(t, "caf\xc3", out)
}

func Test_Shorten_Errors(t *testing.T) {
	tests := []struct {
		name   string
		value  any
		params []any
	}{
		{"non-string value", 42, []any{5}},
		{"nil value", nil, []any{5}},
		{"no params", "hello", nil},
		{"too many params", "hello", []any{5, 6}},
		{"non-numeric param", "hello", []any{"abc"}},
		{"float param string", "hello", []any{"3.5"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Shorten(tt.value, tt.params)
			assert.Error(t, err)
		})
	}
}
