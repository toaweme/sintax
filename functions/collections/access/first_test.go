package access

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_First(t *testing.T) {
	first := firstModifier
	tests := []struct {
		name     string
		value    any
		expected any
	}{
		{"slice of strings", []any{"espresso", "latte", "macchiato"}, "espresso"},
		{"slice of ints", []any{1, 2, 3}, 1},
		{"single element slice", []any{"only"}, "only"},
		{"slice of maps keeps shape", []any{map[string]any{"name": "Alice"}, map[string]any{"name": "Bob"}}, map[string]any{"name": "Alice"}},
		{"typed int slice", []int{7, 8, 9}, 7},
		{"ascii string first byte", "Alice", "A"},
		{"single char string", "x", "x"},
		{"leading multibyte byte", "café", "c"},
		{"byte buffer reads as text", []byte("Alice"), "A"},
		{"single byte buffer", []byte("x"), "x"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := first(tt.value, nil)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_First_Errors(t *testing.T) {
	first := firstModifier
	tests := []struct {
		name  string
		value any
	}{
		{"empty slice", []any{}},
		{"empty string", ""},
		{"empty byte buffer", []byte{}},
		{"nil", nil},
		{"int is not a collection", 42},
		{"bool is not a collection", true},
		{"map is not indexable by first", map[string]any{"a": 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := first(tt.value, nil)
			assert.Error(t, err)
		})
	}
}

func Test_Last(t *testing.T) {
	last := lastModifier
	tests := []struct {
		name     string
		value    any
		expected any
	}{
		{"slice of strings", []any{"espresso", "latte", "macchiato"}, "macchiato"},
		{"slice of ints", []any{1, 2, 3}, 3},
		{"single element slice", []any{"only"}, "only"},
		{"slice of maps keeps shape", []any{map[string]any{"name": "Alice"}, map[string]any{"name": "Bob"}}, map[string]any{"name": "Bob"}},
		{"typed int slice", []int{7, 8, 9}, 9},
		{"ascii string last byte", "Hello", "o"},
		{"single char string", "x", "x"},
		{"byte buffer reads as text", []byte("Hello"), "o"},
		{"single byte buffer", []byte("x"), "x"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := last(tt.value, nil)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func Test_Last_Errors(t *testing.T) {
	last := lastModifier
	tests := []struct {
		name  string
		value any
	}{
		{"empty slice", []any{}},
		{"empty string", ""},
		{"empty byte buffer", []byte{}},
		{"nil", nil},
		{"int is not a collection", 42},
		{"map is not indexable by last", map[string]any{"a": 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := last(tt.value, nil)
			assert.Error(t, err)
		})
	}
}
