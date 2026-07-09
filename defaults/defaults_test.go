package defaults_test

import (
	"testing"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/defaults"
	"github.com/toaweme/sintax/functions"
)

func Test_Defaults_New_RendersFullBattery(t *testing.T) {
	s := sintax.New(defaults.New())

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
	out, err := sintax.New(defaults.New()).Render(`{{ name | upper }}`, map[string]any{"name": "alice"})
	if err != nil {
		t.Fatalf("failed to render: %v", err)
	}
	if out != "ALICE" {
		t.Fatalf("got %q, want %q", out, "ALICE")
	}
}

func Test_Defaults_With_OverridesBuiltin(t *testing.T) {
	funcs := defaults.NewWith(map[string]functions.GlobalModifier{
		"upper": func(any, []any) (any, error) { return "OVERRIDDEN", nil },
	})
	if len(funcs) == 0 {
		t.Fatal("expected a populated modifier map")
	}
	if _, ok := funcs["lower"]; !ok {
		t.Fatal("expected built-ins to remain present alongside overrides")
	}
}
