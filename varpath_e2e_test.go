package sintax

import (
	"errors"
	"testing"

	"github.com/toaweme/sintax/functions"
)

// A dotted name is read as a path only when nothing is stored under the whole
// name, so a producer that publishes flat keys containing dots keeps deciding
// what those names mean. A pipeline stores a step's outputs as "step.field",
// and that key has to keep answering ahead of any walk.
func Test_Render_DottedPath_ResolvesValues(t *testing.T) {
	testCases := []struct {
		name     string
		template string
		vars     map[string]any
		want     any
	}{
		{
			name:     "the flat key wins over walking the same name",
			template: `{{ a.b }}`,
			vars: map[string]any{
				"a.b": "flat",
				"a":   map[string]any{"b": "walked"},
			},
			want: "flat",
		},
		{
			name:     "a name with no flat key walks into the value a prefix holds",
			template: `{{ file_row.record.project_id }}`,
			vars: map[string]any{
				"file_row.record": map[string]any{"project_id": "p1"},
			},
			want: "p1",
		},
		{
			name:     "the walk continues for as many segments as are left",
			template: `{{ step.out.a.b.c }}`,
			vars: map[string]any{
				"step.out": map[string]any{
					"a": map[string]any{"b": map[string]any{"c": "deep"}},
				},
			},
			want: "deep",
		},
		{
			name:     "the longest prefix that hits owns the answer",
			template: `{{ a.b.c }}`,
			vars: map[string]any{
				"a":   map[string]any{"b": map[string]any{"c": "short prefix"}},
				"a.b": map[string]any{"c": "long prefix"},
			},
			want: "long prefix",
		},
		{
			name:     "a numeric segment indexes into a slice",
			template: `{{ row.items.1 }}`,
			vars: map[string]any{
				"row": map[string]any{"items": []any{"first", "second"}},
			},
			want: "second",
		},
		{
			name:     "a walk that runs out is a miss a default answers",
			template: `{{ file_row.record.nope | default:'fallback' }}`,
			vars: map[string]any{
				"file_row.record": map[string]any{"project_id": "p1"},
			},
			want: "fallback",
		},
		{
			name:     "an if reads a walk that ran out as false",
			template: `{{ if row.nope }}yes{{ else }}no{{ endif }}`,
			vars: map[string]any{
				"row": map[string]any{"id": 1},
			},
			want: "no",
		},
		{
			name:     "a for iterates nothing when the walk runs out",
			template: `{{ for x in row.nope }}{{ x }}{{ endfor }}done`,
			vars: map[string]any{
				"row": map[string]any{"id": 1},
			},
			want: "done",
		},
		{
			name:     "a modifier argument walks the same way",
			template: `{{ word | eq:item.a.value }}`,
			vars: map[string]any{
				"word": "v",
				"item": map[string]any{"a": map[string]any{"value": "v"}},
			},
			want: true,
		},
		{
			name:     "a modifier argument takes the flat key first",
			template: `{{ word | eq:item.a.value }}`,
			vars: map[string]any{
				"word":   "flat",
				"item.a": map[string]any{"value": "flat"},
				"item":   map[string]any{"a": map[string]any{"value": "walked"}},
			},
			want: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := New(builtins())

			out, err := s.Render(tt.template, tt.vars)
			if err != nil {
				t.Fatalf("failed to render a dotted variable path: %v", err)
			}
			if out != tt.want {
				t.Fatalf("got %v, want %v", out, tt.want)
			}
		})
	}
}

// Walking into something that holds no paths at all is the template being wrong
// rather than the data being absent, so it stays terminal and no default rescues
// it. A modifier argument that misses fails the render too, because an argument
// has no pipeline of its own to carry a miss down.
func Test_Render_DottedPath_Failures(t *testing.T) {
	testCases := []struct {
		name      string
		template  string
		vars      map[string]any
		want      error
		catchable bool
	}{
		{
			name:     "walking into a string is terminal",
			template: `{{ row.id.nope | default:'fallback' }}`,
			vars: map[string]any{
				"row": map[string]any{"id": "text"},
			},
			want: functions.ErrInvalidValueType,
		},
		{
			name:      "a bare token whose walk runs out is a miss nothing answered",
			template:  `{{ row.nope }}`,
			vars:      map[string]any{"row": map[string]any{"id": 1}},
			want:      functions.ErrAllowsDefaultFunc,
			catchable: true,
		},
		{
			name:     "walking a modifier argument into a string is terminal",
			template: `{{ word | eq:item.a.b }}`,
			vars: map[string]any{
				"word": "v",
				"item": map[string]any{"a": "text"},
			},
			want: functions.ErrInvalidValueType,
		},
		{
			name:      "a modifier argument whose walk runs out fails the render",
			template:  `{{ word | eq:item.nope | default:'fallback' }}`,
			vars:      map[string]any{"word": "v", "item": map[string]any{"a": 1}},
			want:      functions.ErrAllowsDefaultFunc,
			catchable: true,
		},
		{
			name:     "a name no prefix answers is still an unknown variable",
			template: `{{ nothing.here.at.all }}`,
			vars:     map[string]any{},
			want:     ErrVariableNotFound,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := New(builtins())

			_, err := s.Render(tt.template, tt.vars)
			if err == nil {
				t.Fatalf("got no error, want %v", tt.want)
			}
			if !errors.Is(err, tt.want) {
				t.Fatalf("got %v, want %v", err, tt.want)
			}
			if !tt.catchable && errors.Is(err, functions.ErrAllowsDefaultFunc) {
				t.Fatalf("got a catchable %v, want a terminal failure", err)
			}
		})
	}
}
