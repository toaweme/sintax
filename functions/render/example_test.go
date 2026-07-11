package render_test

import (
	"fmt"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/render"
)

// template is a contextual modifier auto-wired into every renderer, so it is
// available without registering any modifier map. The reference keeps the
// package import live and links the example to the exported symbol.
var _ = render.Template

func renderTpl(tpl string, vars map[string]any) string {
	out, err := sintax.New(nil).Render(tpl, vars)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return fmt.Sprintf("%v", out)
}

// ExampleTemplate renders a string value as a nested template that inherits the
// parent variables, so the partial sees everything the outer template sees.
func ExampleTemplate() {
	fmt.Println(renderTpl(`{{ tpl | template }}`, map[string]any{
		"tpl":  "Hello {{ name }}",
		"name": "World",
	}))
	// Output: Hello World
}

// ExampleTemplate_scope renders with an isolated scope, where the nested
// template sees only the passed map and the parent variables are hidden.
func ExampleTemplate_scope() {
	fmt.Println(renderTpl(`{{ tpl | template:extra }}`, map[string]any{
		"tpl":   "{{ city }}",
		"name":  "parent",
		"extra": map[string]any{"city": "Vilnius"},
	}))
	// Output: Vilnius
}
