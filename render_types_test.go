package sintax

import (
	"reflect"
	"testing"
)

// Render answers a single-token template with the value's own type rather than
// its text. These assertions go through Render on purpose. The passthrough lives
// in renderRange, so a test that calls renderVariable proves nothing about it,
// and that gap is what let a bool reach callers as "true" while every test
// stayed green.
func Test_Render_PassesValueTypesThrough(t *testing.T) {
	testCases := []struct {
		name     string
		template string
		vars     map[string]any
		want     any
	}{
		{
			name:     "int keeps its type",
			template: `{{ i }}`,
			vars:     map[string]any{"i": 7},
			want:     7,
		},
		{
			name:     "int64 keeps its type",
			template: `{{ i }}`,
			vars:     map[string]any{"i": int64(7)},
			want:     int64(7),
		},
		{
			name:     "float64 keeps its type",
			template: `{{ f }}`,
			vars:     map[string]any{"f": 7.5},
			want:     7.5,
		},
		{
			name:     "true keeps its type",
			template: `{{ b }}`,
			vars:     map[string]any{"b": true},
			want:     true,
		},
		{
			name:     "false keeps its type, and is not the string \"false\"",
			template: `{{ b }}`,
			vars:     map[string]any{"b": false},
			want:     false,
		},
		{
			name:     "a string is still a string",
			template: `{{ s }}`,
			vars:     map[string]any{"s": "ada"},
			want:     "ada",
		},
		{
			name:     "a slice passes through whole",
			template: `{{ xs }}`,
			vars:     map[string]any{"xs": []any{1, 2}},
			want:     []any{1, 2},
		},
		{
			name:     "a map passes through whole",
			template: `{{ m }}`,
			vars:     map[string]any{"m": map[string]any{"a": 1}},
			want:     map[string]any{"a": 1},
		},
		{
			name:     "a pipeline does not change the rule for int",
			template: `{{ i | default:0 }}`,
			vars:     map[string]any{"i": 7},
			want:     7,
		},
		{
			name:     "eq answers with a bool",
			template: `{{ i | eq:1 }}`,
			vars:     map[string]any{"i": 1},
			want:     true,
		},
		{
			name:     "gt answers with a bool",
			template: `{{ i | gt:9 }}`,
			vars:     map[string]any{"i": 1},
			want:     false,
		},
		{
			name:     "not answers with a bool",
			template: `{{ b | not }}`,
			vars:     map[string]any{"b": true},
			want:     false,
		},
		{
			name:     "length answers with a number",
			template: `{{ xs | length }}`,
			vars:     map[string]any{"xs": []any{1, 2, 3}},
			want:     3,
		},
		{
			name:     "a default fallback keeps the fallback's type",
			template: `{{ nosuchvar | default:true }}`,
			vars:     map[string]any{},
			want:     true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := New(builtins())

			out, err := s.Render(tt.template, tt.vars)
			if err != nil {
				t.Fatalf("failed to render a value-typed template: %v", err)
			}
			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %#v (%T), want %#v (%T)", out, out, tt.want, tt.want)
			}
		})
	}
}

// A value interpolated among text still reads naturally, which is what makes the
// passthrough safe to put first. fmt formatting renders a bool and an int the
// same way a type-specific branch would, so only the single-token case has a
// type to preserve and the text case is unaffected.
func Test_Render_StringifiesValuesAmongText(t *testing.T) {
	testCases := []struct {
		name     string
		template string
		vars     map[string]any
		want     string
	}{
		{
			name:     "a bool among text",
			template: `active={{ b }}!`,
			vars:     map[string]any{"b": true},
			want:     "active=true!",
		},
		{
			name:     "a false bool among text",
			template: `active={{ b }}!`,
			vars:     map[string]any{"b": false},
			want:     "active=false!",
		},
		{
			name:     "an int among text",
			template: `n={{ i }}`,
			vars:     map[string]any{"i": 7},
			want:     "n=7",
		},
		{
			name:     "a float among text",
			template: `n={{ f }}`,
			vars:     map[string]any{"f": 7.5},
			want:     "n=7.5",
		},
		{
			name:     "two tokens stringify rather than pass through",
			template: `{{ a }}{{ b }}`,
			vars:     map[string]any{"a": 1, "b": true},
			want:     "1true",
		},
		{
			name:     "a modifier result among text",
			template: `over={{ i | gt:0 }}`,
			vars:     map[string]any{"i": 5},
			want:     "over=true",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := New(builtins())

			out, err := s.Render(tt.template, tt.vars)
			if err != nil {
				t.Fatalf("failed to render a value among text: %v", err)
			}
			if out != tt.want {
				t.Fatalf("got %#v, want %q", out, tt.want)
			}
		})
	}
}
