package sintax

import (
	"errors"
	"strings"
	"testing"

	"github.com/toaweme/sintax/functions"
)

// A failing modifier must be identifiable without reading the message, so these
// assert the field rather than the rendered text. The chain case is the one that
// matters: four modifiers, one of them wrong, and nothing else in the error says
// which.
func Test_Render_ModifierError_NamesTheModifier(t *testing.T) {
	testCases := []struct {
		name     string
		template string
		vars     map[string]any
		modifier string
		variable string
	}{
		{
			name:     "single modifier",
			template: `{{ name | upper:'x' }}`,
			vars:     map[string]any{"name": "ada"},
			modifier: "upper",
			variable: "name",
		},
		{
			name:     "failing link of a chain",
			template: `{{ text | trim | upper:'z' | lower }}`,
			vars:     map[string]any{"text": "  hi  "},
			modifier: "upper",
			variable: "text",
		},
		{
			name:     "overloaded modifier rejecting every clause",
			template: `{{ items | first:9 }}`,
			vars:     map[string]any{"items": []any{1, 2}},
			modifier: "first",
			variable: "items",
		},
		{
			name:     "contextual modifier",
			template: `{{ self | template }}`,
			vars:     map[string]any{"self": "{{ self | template }}"},
			modifier: "template",
			variable: "self",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := New(builtins())

			_, err := s.Render(tt.template, tt.vars)
			if err == nil {
				t.Fatalf("expected %s to fail", tt.template)
			}

			var modErr *ModifierError
			if !errors.As(err, &modErr) {
				t.Fatalf("got %v, want a *ModifierError in the chain", err)
			}
			if modErr.Modifier != tt.modifier {
				t.Errorf("got modifier %q, want %q", modErr.Modifier, tt.modifier)
			}
			if modErr.Variable != tt.variable {
				t.Errorf("got variable %q, want %q", modErr.Variable, tt.variable)
			}
			if !strings.Contains(err.Error(), tt.modifier) {
				t.Errorf("got %q, want the modifier name in the message too", err)
			}
		})
	}
}

// The typed error must not cost callers the sentinels they already match on.
func Test_Render_ModifierError_KeepsSentinelChain(t *testing.T) {
	s := New(builtins())

	_, err := s.Render(`{{ name | upper:'x' }}`, map[string]any{"name": "ada"})
	if !errors.Is(err, ErrFunctionApplyFailed) {
		t.Errorf("got %v, want ErrFunctionApplyFailed to stay in the chain", err)
	}
	if !errors.Is(err, functions.ErrInvalidParamType) {
		t.Errorf("got %v, want the modifier's own sentinel to stay reachable", err)
	}

	_, err = s.Render(`{{ self | template }}`, map[string]any{"self": "{{ self | template }}"})
	if !errors.Is(err, ErrMaxDepthExceeded) {
		t.Errorf("got %v, want ErrMaxDepthExceeded to stay in the chain", err)
	}
}

// A non-fatal miss that `default` catches is not a terminal failure, so it must
// not surface as a ModifierError. These are the shapes a template author
// actually writes, and the filter/first/default chain is the reason the empty
// collection has to count as a miss rather than a type error.
func Test_Render_ModifierError_NotRaisedBehindDefault(t *testing.T) {
	testCases := []struct {
		name     string
		template string
		want     any
	}{
		{
			name:     "find misses",
			template: `{{ items | find:'id',99 | default:'fallback' }}`,
			want:     "fallback",
		},
		{
			name:     "filter matches nothing then first",
			template: `{{ items | filter:'id',99 | first | default:'fallback' }}`,
			want:     "fallback",
		},
		{
			name:     "first of an empty slice",
			template: `{{ empty | first | default:'fallback' }}`,
			want:     "fallback",
		},
		{
			name:     "last of an empty slice",
			template: `{{ empty | last | default:'fallback' }}`,
			want:     "fallback",
		},
		{
			name:     "first of an empty string",
			template: `{{ blank | first | default:'fallback' }}`,
			want:     "fallback",
		},
	}

	vars := map[string]any{
		"items": []any{map[string]any{"id": 1, "name": "ada"}},
		"empty": []any{},
		"blank": "",
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := New(builtins())

			out, err := s.Render(tt.template, vars)
			if err != nil {
				t.Fatalf("failed to render a caught miss with a default fallback: %v", err)
			}
			if out != tt.want {
				t.Fatalf("got %v, want %v", out, tt.want)
			}
		})
	}
}

// Without a default in the chain a miss is still an error, so an empty
// collection never passes silently as nil.
func Test_Render_Miss_ErrorsWithoutDefault(t *testing.T) {
	s := New(builtins())

	_, err := s.Render(`{{ empty | first }}`, map[string]any{"empty": []any{}})
	if err == nil {
		t.Fatal("expected first of an empty slice to fail without a default")
	}
	if !errors.Is(err, functions.ErrAllowsDefaultFunc) {
		t.Errorf("got %v, want the miss marker to stay reachable", err)
	}
	// the marker is for the engine, not for whoever reads the message
	if strings.Contains(err.Error(), "non-fatal error") {
		t.Errorf("got %q, want the ErrAllowsDefaultFunc marker kept out of the text", err)
	}
}

// A terminal failure is still terminal with a `default` in the chain, since
// default only catches ErrAllowsDefaultFunc. The ModifierError must survive that
// path and still name the modifier.
func Test_Render_ModifierError_RaisedDespiteDefault(t *testing.T) {
	s := New(builtins())

	_, err := s.Render(`{{ name | upper:'x' | default:'fallback' }}`, map[string]any{"name": "ada"})
	if err == nil {
		t.Fatal("expected a terminal modifier failure to survive a default in the chain")
	}

	var modErr *ModifierError
	if !errors.As(err, &modErr) {
		t.Fatalf("got %v, want a *ModifierError in the chain", err)
	}
	if modErr.Modifier != "upper" {
		t.Errorf("got modifier %q, want %q", modErr.Modifier, "upper")
	}
}
