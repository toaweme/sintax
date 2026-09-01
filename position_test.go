package sintax

import (
	"errors"
	"strings"
	"testing"

	"github.com/toaweme/sintax/functions"
)

// positionEngine registers the two modifiers these tests need, a failing one and
// one that re-enters the engine, rather than importing defaults, which would
// close an import cycle on an internal test.
func positionEngine() Sintax {
	return New(
		WithModifiers(map[string]GlobalModifier{
			"upper": func(any, []any) (any, error) { return nil, errBoom },
		}),
		WithContextualModifiers(map[string]ContextualModifier{
			"template": func(render func(string, map[string]any) (any, error), vars map[string]any, value any, _ []any) (any, error) {
				text, _ := value.(string)

				return render(text, vars)
			},
		}),
	)
}

var errBoom = errors.New("boom")

func Test_ModifierError_Position(t *testing.T) {
	tests := []struct {
		name       string
		template   string
		wantLine   int
		wantColumn int
		wantText   string
		wantInMsg  bool
	}{
		{
			name:       "a single line names its column and stays out of the message",
			template:   "{{ name | upper:'x' }}",
			wantLine:   1,
			wantColumn: 1,
			wantText:   "{{ name | upper:'x' }}",
		},
		{
			name:       "a later line is counted from one",
			template:   "first\nsecond\nthe {{ name | upper:'x' }} one\n",
			wantLine:   3,
			wantColumn: 5,
			wantText:   "the {{ name | upper:'x' }} one",
			wantInMsg:  true,
		},
		{
			name:       "a carriage return is not part of the line",
			template:   "first\r\n{{ name | upper:'x' }}\r\n",
			wantLine:   2,
			wantColumn: 1,
			wantText:   "{{ name | upper:'x' }}",
			wantInMsg:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := positionEngine().RenderString(tt.template, map[string]any{"name": "ada"})
			if err == nil {
				t.Fatal("rendering succeeded, want a modifier failure")
			}

			var failure *ModifierError
			if !errors.As(err, &failure) {
				t.Fatalf("error = %v, want a *ModifierError", err)
			}
			if failure.Position.Line != tt.wantLine || failure.Position.Column != tt.wantColumn {
				t.Errorf("position = %d:%d, want %d:%d",
					failure.Position.Line, failure.Position.Column, tt.wantLine, tt.wantColumn)
			}
			if failure.Position.Text != tt.wantText {
				t.Errorf("line text = %q, want %q", failure.Position.Text, tt.wantText)
			}
			if failure.Source != tt.template {
				t.Errorf("source = %q, want %q", failure.Source, tt.template)
			}
			if got := strings.Contains(failure.Error(), "at "); got != tt.wantInMsg {
				t.Errorf("message names a position = %v, want %v (%s)", got, tt.wantInMsg, failure.Error())
			}
		})
	}
}

// Test_ModifierError_NestedTemplatePosition checks that a modifier failing
// inside a template the engine re-entered reports the nested source, not the
// document that named it.
func Test_ModifierError_NestedTemplatePosition(t *testing.T) {
	const nested = "one\ntwo {{ name | upper:'x' }}"

	_, err := positionEngine().RenderString("{{ body | template }}", map[string]any{
		"body": nested,
		"name": "ada",
	})
	if err == nil {
		t.Fatal("rendering succeeded, want a modifier failure")
	}

	var failure *ModifierError
	if !errors.As(err, &failure) {
		t.Fatalf("error = %v, want a *ModifierError", err)
	}
	if failure.Source != nested {
		t.Errorf("source = %q, want the nested template %q", failure.Source, nested)
	}
	if failure.Position.Line != 2 || failure.Position.Column != 5 {
		t.Errorf("position = %d:%d, want 2:5", failure.Position.Line, failure.Position.Column)
	}
}

// Test_ModifierError_UnknownPosition checks that a token the engine built
// rather than parsed reports no position instead of pointing at byte zero.
func Test_ModifierError_UnknownPosition(t *testing.T) {
	token, err := parseExpr("name | upper:'x'")
	if err != nil {
		t.Fatalf("failed to parse expression: %v", err)
	}

	failure := modifierFailure("upper", token, functions.ErrAllowsDefaultFunc)

	var modifier *ModifierError
	if !errors.As(failure, &modifier) {
		t.Fatalf("error = %v, want a *ModifierError", failure)
	}
	if modifier.Position.Known() {
		t.Errorf("position = %+v, want it unknown", modifier.Position)
	}
	if strings.Contains(modifier.Error(), "at ") {
		t.Errorf("message = %q, want no position in it", modifier.Error())
	}
}
