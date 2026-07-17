package defaults_test

import (
	"errors"
	"testing"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/defaults"
	"github.com/toaweme/sintax/functions"
)

func Test_Defaults_New_RendersFullBattery(t *testing.T) {
	s := sintax.New(defaults.All())

	// a pipeline that crosses several regrouped modifier groups: path/query,
	// casing, text/edit and convert/serialize.
	out, err := s.Render(`{{ p | filename | upper | wrap:'name' | json }}`, map[string]any{
		"p": "downloads/report.pdf",
	})
	if err != nil {
		t.Fatalf("failed to render pipeline: %v", err)
	}

	want := `{"name":"REPORT.PDF"}`
	if got, ok := out.(string); !ok || got != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

func Test_Defaults_New_UpperModifier(t *testing.T) {
	out, err := sintax.New(defaults.All()).Render(`{{ name | upper }}`, map[string]any{"name": "alice"})
	if err != nil {
		t.Fatalf("failed to render: %v", err)
	}
	if out != "ALICE" {
		t.Fatalf("got %q, want %q", out, "ALICE")
	}
}

func Test_Defaults_All_OverridesBuiltin(t *testing.T) {
	s := sintax.New(defaults.All(), sintax.WithModifiers(map[string]functions.GlobalModifier{
		"upper": func(any, []any) (any, error) { return "OVERRIDDEN", nil },
	}))

	out, err := s.Render(`{{ name | upper }}`, map[string]any{"name": "alice"})
	if err != nil {
		t.Fatalf("failed to render: %v", err)
	}
	if out != "OVERRIDDEN" {
		t.Fatalf("got %q, want %q, so the later option did not win", out, "OVERRIDDEN")
	}

	// the overridden name is replaced, not the whole set
	out, err = s.Render(`{{ name | lower }}`, map[string]any{"name": "ALICE"})
	if err != nil {
		t.Fatalf("failed to render: %v", err)
	}
	if out != "alice" {
		t.Fatalf("got %q, want %q, so an override dropped the other built-ins", out, "alice")
	}
}

// All bundles the contextual modifiers too, so `template` resolves without the
// caller wiring functions/render itself.
func Test_Defaults_All_WiresContextualModifiers(t *testing.T) {
	out, err := sintax.New(defaults.All()).Render(`{{ tpl | template }}`, map[string]any{
		"tpl":  "Hi {{ name }}",
		"name": "Bob",
	})
	if err != nil {
		t.Fatalf("failed to render nested template: %v", err)
	}
	if out != "Hi Bob" {
		t.Fatalf("got %q, want %q", out, "Hi Bob")
	}
}

// An engine built from global modifiers alone knows no contextual modifier, so
// `template` is unresolved rather than silently available.
func Test_Defaults_New_OmitsContextualModifiers(t *testing.T) {
	_, err := sintax.New(sintax.WithModifiers(defaults.New())).Render(`{{ tpl | template }}`, map[string]any{
		"tpl": "Hi",
	})
	if !errors.Is(err, sintax.ErrFunctionNotFound) {
		t.Fatalf("got %v, want ErrFunctionNotFound", err)
	}
}
