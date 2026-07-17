package serialize

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

// Test_Markdown_StubReturnsError proves the shipped stub returns an error until
// a consumer injects a real HTML-to-Markdown converter. The value is irrelevant
// to the stub, so a range of inputs all take the same path.
func Test_Markdown_StubReturnsError(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{"html string", "<h1>Welcome</h1>"},
		{"empty string", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Markdown(tt.value)
			assert.Error(t, err)
			assert.Equal(t, "", out)
			assert.Equal(t, "markdown function needs to be injected", err.Error())
		})
	}
}

// Test_Markdown_RejectsNonString proves the modifier turns a value that is not
// HTML text away before any converter sees it. The clause takes a string, so
// Wrap coerces strictly, and a piped map fails here rather than reaching a
// converter stringified as "map[a:1]" and being converted as if it were markup.
func Test_Markdown_RejectsNonString(t *testing.T) {
	tests := []struct {
		name  string
		value any
	}{
		{"map", map[string]any{"a": 1}},
		{"number", 42},
		{"nil", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := markdownModifier(tt.value, nil)
			assert.Error(t, err)
		})
	}
}

// Test_Markdown_Dispatch proves the registered modifier surfaces the stub error
// too.
func Test_Markdown_Dispatch(t *testing.T) {
	mdMod := markdownModifier
	out, err := mdMod("<h1>Welcome</h1>", nil)
	assert.Error(t, err)
	assert.Equal(t, "", out)
	assert.Equal(t, "markdown function needs to be injected", err.Error())
}
