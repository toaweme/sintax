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
		value any
	}{
		{"html string", "<h1>Welcome</h1>"},
		{"empty string", ""},
		{"nil", nil},
		{"non-string", 42},
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

// Test_Markdown_Dispatch proves the registered modifier surfaces the stub error
// too.
func Test_Markdown_Dispatch(t *testing.T) {
	mdMod := markdownModifier
	out, err := mdMod("<h1>Welcome</h1>", nil)
	assert.Error(t, err)
	assert.Equal(t, "", out)
	assert.Equal(t, "markdown function needs to be injected", err.Error())
}
