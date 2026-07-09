package sintax

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

// Test_Escape_Engine drives the escape modifiers through the full parse+render
// pipeline, so it also covers chaining with other modifiers and use inside a
// loop.
func Test_Escape_Engine(t *testing.T) {
	tests := []struct {
		name     string
		template string
		vars     map[string]any
		expected any
	}{
		{
			name:     "html",
			template: "{{ v | escape_html }}",
			vars:     map[string]any{"v": "<b>&</b>"},
			expected: "&lt;b&gt;&amp;&lt;/b&gt;",
		},
		{
			name:     "html quotes",
			template: "{{ v | escape_html }}",
			vars:     map[string]any{"v": `"x"`},
			expected: "&#34;x&#34;",
		},
		{
			name:     "url",
			template: "https://e.com/s?q={{ v | escape_url }}",
			vars:     map[string]any{"v": "tea & coffee"},
			expected: "https://e.com/s?q=tea+%26+coffee",
		},
		{
			name:     "js inside a script string",
			template: `<script>var n = "{{ v | escape_js }}";</script>`,
			vars:     map[string]any{"v": `a"; </script>`},
			expected: "<script>var n = \"a\\\"; \\u003C\\/script\\u003E\";</script>",
		},
		{
			name:     "chains after default",
			template: "{{ v | default:'<none>' | escape_html }}",
			vars:     map[string]any{"v": nil},
			expected: "&lt;none&gt;",
		},
		{
			name:     "escapes a number",
			template: "{{ n | escape_html }}",
			vars:     map[string]any{"n": 42},
			expected: "42",
		},
		{
			name:     "inside a loop",
			template: "{{ for c in comments }}{{ c | escape_html }} {{ endfor }}",
			vars:     map[string]any{"comments": []any{"<a>", "b&c"}},
			expected: "&lt;a&gt; b&amp;c ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(BuiltinFunctions(nil, nil))
			out, err := s.Render(tt.template, tt.vars)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}
