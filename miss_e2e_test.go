package sintax

import (
	"errors"
	"testing"

	"github.com/toaweme/sintax/functions"
)

// A miss travels as nil until something answers it, so the default that does can
// sit anywhere downstream. Before this, a caught miss handed nil to the next
// modifier, whose type rejection became terminal and outran the default, so a
// fallback only worked when it sat immediately after the modifier that missed.
// That is not a rule anyone could infer from reading a template.
func Test_Render_Miss_ReachesADistantDefault(t *testing.T) {
	testCases := []struct {
		name     string
		template string
		want     any
	}{
		{
			name:     "default immediately after the miss",
			template: `{{ empty | first | default:'fallback' }}`,
			want:     "fallback",
		},
		{
			name:     "a modifier between the miss and the default",
			template: `{{ empty | first | upper | default:'fallback' }}`,
			want:     "fallback",
		},
		{
			name:     "several modifiers between the miss and the default",
			template: `{{ cfg | key:'nope' | upper | trim | default:'fallback' }}`,
			want:     "fallback",
		},
		{
			name:     "an absent variable is a miss like any other",
			template: `{{ nosuchvar | upper | default:'fallback' }}`,
			want:     "fallback",
		},
		{
			name:     "a caught miss resumes the pipeline",
			template: `{{ cfg | key:'nope' | default:'fallback' | upper }}`,
			want:     "FALLBACK",
		},
		{
			name:     "pluck of an absent field falls back to an empty slice",
			template: `{{ rows | pluck:'nope' | default:[] }}`,
			want:     nil, // compared by length below, a slice is not comparable with !=
		},
	}

	vars := map[string]any{
		"empty": []any{},
		"cfg":   map[string]any{"host": "db"},
		"rows":  []any{map[string]any{"id": 1}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := New(builtins())

			out, err := s.Render(tt.template, vars)
			if err != nil {
				t.Fatalf("failed to render a miss caught by a downstream default: %v", err)
			}
			if tt.want == nil {
				got, ok := out.([]any)
				if !ok || len(got) != 0 {
					t.Fatalf("got %#v, want an empty slice", out)
				}
				return
			}
			if out != tt.want {
				t.Fatalf("got %v, want %v", out, tt.want)
			}
		})
	}
}

// An if condition and a for iterable answer a miss on their own, so neither
// needs a default to say what absent data means. Writing `| default:false`
// inside a condition to make it work would be noise.
func Test_Render_Miss_AnsweredByIfAndFor(t *testing.T) {
	testCases := []struct {
		name     string
		template string
		want     any
	}{
		{
			name:     "a missing key is false",
			template: `{{ if cfg | key:'nope' }}yes{{ else }}no{{ endif }}`,
			want:     "no",
		},
		{
			name:     "a present key is true",
			template: `{{ if cfg | key:'host' }}yes{{ else }}no{{ endif }}`,
			want:     "yes",
		},
		{
			name:     "first of an empty slice is false",
			template: `{{ if empty | first }}yes{{ else }}no{{ endif }}`,
			want:     "no",
		},
		{
			name:     "a negated miss is true",
			template: `{{ if cfg | key:'nope' | not }}yes{{ else }}no{{ endif }}`,
			want:     "yes",
		},
		{
			name:     "a missing iterable yields no iterations",
			template: `{{ for x in cfg | key:'nope' }}{{ x }}{{ endfor }}done`,
			want:     "done",
		},
	}

	vars := map[string]any{
		"empty": []any{},
		"cfg":   map[string]any{"host": "db"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := New(builtins())

			out, err := s.Render(tt.template, vars)
			if err != nil {
				t.Fatalf("failed to render a miss in a position that answers it: %v", err)
			}
			if out != tt.want {
				t.Fatalf("got %v, want %v", out, tt.want)
			}
		})
	}
}

// An uncaught miss in an output position fails. Rendering a field as empty
// because its key was misspelled is indistinguishable from data that was
// genuinely absent, and the template author is the only one who can tell them
// apart.
func Test_Render_Miss_LoudInOutputPosition(t *testing.T) {
	testCases := []struct {
		name     string
		template string
		modifier string
	}{
		{
			name:     "a missing key",
			template: `{{ cfg | key:'nope' }}`,
			modifier: "key",
		},
		{
			name:     "an out-of-range index",
			template: `{{ rows | key:99 }}`,
			modifier: "key",
		},
		{
			name:     "an absent pluck field",
			template: `{{ rows | pluck:'nope' }}`,
			modifier: "pluck",
		},
		{
			name:     "an absent sum column",
			template: `{{ rows | sum:'nope' }}`,
			modifier: "sum",
		},
	}

	vars := map[string]any{
		"cfg":  map[string]any{"host": "db"},
		"rows": []any{map[string]any{"id": 1}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := New(builtins())

			_, err := s.Render(tt.template, vars)
			if err == nil {
				t.Fatal("expected an uncaught miss to fail rather than render empty")
			}
			if !errors.Is(err, functions.ErrAllowsDefaultFunc) {
				t.Errorf("got %v, want the miss marker to stay reachable", err)
			}

			var modErr *ModifierError
			if !errors.As(err, &modErr) {
				t.Fatalf("got %v, want a *ModifierError naming the modifier that missed", err)
			}
			if modErr.Modifier != tt.modifier {
				t.Errorf("got modifier %q, want %q", modErr.Modifier, tt.modifier)
			}
		})
	}
}

// A template that cannot mean anything stays terminal no matter what sits
// downstream. A default is an answer to absent data, not a way to make a broken
// template render.
func Test_Render_AuthorError_NotCatchableByDefault(t *testing.T) {
	testCases := []struct {
		name     string
		template string
	}{
		{
			name:     "key on a value that is neither map nor slice",
			template: `{{ "str" | key:'x' | default:'fallback' }}`,
		},
		{
			name:     "key with no parameter",
			template: `{{ cfg | key | default:'fallback' }}`,
		},
		{
			name:     "key with a non-string map key",
			template: `{{ cfg | key:42 | default:'fallback' }}`,
		},
		{
			name:     "pluck over a slice of scalars",
			template: `{{ nums | pluck:'id' | default:[] }}`,
		},
		{
			name:     "a modifier given too many params",
			template: `{{ cfg | key:'host' | upper:'x' | default:'fallback' }}`,
		},
	}

	vars := map[string]any{
		"cfg":  map[string]any{"host": "db"},
		"nums": []any{1, 2},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := New(builtins())

			_, err := s.Render(tt.template, vars)
			if err == nil {
				t.Fatal("expected a broken template to fail despite a default in the chain")
			}
			if errors.Is(err, functions.ErrAllowsDefaultFunc) {
				t.Errorf("got %v, want a broken template to stay out of default's reach", err)
			}
		})
	}
}

// A param is written in the template and does not depend on the data, so a
// mistake in one is real whether or not the value arrived. Without this, a typo
// would hide behind a default for exactly as long as the data stayed absent and
// then surface later, when the data showed up, looking like the data broke it.
func Test_Render_ParamError_LoudEvenWhileDataIsAbsent(t *testing.T) {
	testCases := []struct {
		name     string
		template string
		vars     map[string]any
	}{
		{
			name:     "a miss in flight does not excuse a bad param",
			template: `{{ cfg | key:'nope' | upper:'x' | default:'fallback' }}`,
			vars:     map[string]any{"cfg": map[string]any{"host": "db"}},
		},
		{
			name:     "the same template with the data present",
			template: `{{ cfg | key:'host' | upper:'x' | default:'fallback' }}`,
			vars:     map[string]any{"cfg": map[string]any{"host": "db"}},
		},
		{
			name:     "an absent variable does not excuse a bad param either",
			template: `{{ nosuchvar | upper:'x' | default:'fallback' }}`,
			vars:     map[string]any{},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := New(builtins())

			_, err := s.Render(tt.template, tt.vars)
			if err == nil {
				t.Fatal("expected a bad param to fail whether or not the data was there")
			}
			if !errors.Is(err, functions.ErrInvalidParamType) {
				t.Errorf("got %v, want the param rejection to survive", err)
			}
		})
	}
}

// The counterpart to the above. A modifier rejecting the nil that stands in for
// absent data says nothing about the template, so the miss keeps traveling. A
// rejection carrying no sentinel at all (length) must stay catchable too, rather
// than being promoted to a hard failure on a guess.
func Test_Render_ValueRejection_LeavesTheMissTraveling(t *testing.T) {
	testCases := []struct {
		name     string
		template string
		want     any
	}{
		{
			name:     "a typed value rejection",
			template: `{{ cfg | key:'nope' | upper | default:'fallback' }}`,
			want:     "fallback",
		},
		{
			name:     "several modifiers declining in a row",
			template: `{{ cfg | key:'nope' | upper | trim | default:'fallback' }}`,
			want:     "fallback",
		},
		{
			name:     "a rejection carrying no sentinel",
			template: `{{ nosuchvar | length | default:'fallback' }}`,
			want:     "fallback",
		},
	}

	vars := map[string]any{"cfg": map[string]any{"host": "db"}}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := New(builtins())

			out, err := s.Render(tt.template, vars)
			if err != nil {
				t.Fatalf("failed to keep a miss catchable past a modifier declining it: %v", err)
			}
			if out != tt.want {
				t.Fatalf("got %v, want %v", out, tt.want)
			}
		})
	}
}

// A miss carries its cause, so a consumer can still tell a not-found file from
// any other miss without parsing the message.
func Test_Miss_CarriesItsCause(t *testing.T) {
	sentinel := errors.New("underlying cause")
	err := functions.Miss("failed to do the thing: %w", sentinel)

	if !errors.Is(err, functions.ErrAllowsDefaultFunc) {
		t.Errorf("got %v, want the miss marker reachable", err)
	}
	if !errors.Is(err, sentinel) {
		t.Errorf("got %v, want the wrapped cause reachable", err)
	}
	if err.Error() != "failed to do the thing: underlying cause" {
		t.Errorf("got %q, want the marker kept out of the message text", err)
	}
}
