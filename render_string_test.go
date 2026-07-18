package sintax

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_RenderString(t *testing.T) {
	tests := []struct {
		name     string
		template string
		vars     map[string]any
		expected string
	}{
		{
			name:     "text with an interpolated variable",
			template: "Hello, {{ name }}!",
			vars:     map[string]any{"name": "Ada"},
			expected: "Hello, Ada!",
		},
		{
			name:     "a lone string expression",
			template: "{{ name | upper }}",
			vars:     map[string]any{"name": "ada"},
			expected: "ADA",
		},
		{
			name:     "a lone bool renders like it does inline",
			template: "{{ flag }}",
			vars:     map[string]any{"flag": true},
			expected: "true",
		},
		{
			name:     "a lone int renders in base 10",
			template: "{{ n }}",
			vars:     map[string]any{"n": 42},
			expected: "42",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(builtins())
			out, err := s.RenderString(tt.template, tt.vars)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_RenderString_MatchesInterpolation pins the contract that RenderString of
// a bare expression yields the same text as that expression embedded among
// other text, so the two stringifying paths can never drift apart.
func Test_RenderString_MatchesInterpolation(t *testing.T) {
	s := New(builtins())
	vars := map[string]any{"flag": false, "n": 7}

	for _, expr := range []string{"{{ flag }}", "{{ n }}"} {
		lone, err := s.RenderString(expr, vars)
		assert.NoError(t, err)
		embedded, err := s.RenderString("x"+expr+"y", vars)
		assert.NoError(t, err)
		assert.Equal(t, "x"+lone+"y", embedded)
	}
}

// Test_RenderString_Error surfaces the render error rather than a coerced value.
func Test_RenderString_Error(t *testing.T) {
	s := New(builtins())
	_, err := s.RenderString("{{ missing }}", map[string]any{})
	assert.Error(t, err)
}
