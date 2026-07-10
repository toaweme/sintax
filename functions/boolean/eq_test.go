package boolean

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_Eq(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		params   []any
		expected any
	}{
		{"equal strings", "active", []any{"active"}, true},
		{"unequal strings", "active", []any{"archived"}, false},
		{"empty strings equal", "", []any{""}, true},
		{"equal ints", 3, []any{3}, true},
		{"unequal ints", 3, []any{4}, false},
		{"zero equals zero", 0, []any{0}, true},
		{"equal floats", 1.5, []any{1.5}, true},
		{"int equals float across kinds", 5, []any{5.0}, true},
		{"float equals int across kinds", 5.0, []any{5}, true},
		{"equal bools", true, []any{true}, true},
		{"unequal bools", true, []any{false}, false},
		{"number and its string form differ", 5, []any{"5"}, false},
		{"nil equals nil", nil, []any{nil}, true},
		{"nil differs from zero", nil, []any{0}, false},
		{"value differs from nil param", 0, []any{nil}, false},
		{"unicode strings equal", "café", []any{"café"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Eq(tt.value, tt.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_Eq_MissingParam(t *testing.T) {
	out, err := Eq("active", nil)
	assert.Error(t, err)
	assert.Equal(t, nil, out)
	assert.Equal(t, "eq function requires at least one parameter", err.Error())
}
