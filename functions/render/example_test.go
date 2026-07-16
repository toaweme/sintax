package render_test

import (
	"errors"
	"fmt"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/fs"
	"github.com/toaweme/sintax/functions/render"
)

func renderTpl(tpl string, vars map[string]any) string {
	out, err := sintax.New(
		sintax.WithContextualModifiers(render.ContextualModifiers()),
	).Render(tpl, vars)
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

// ExampleTemplate_multipleVars expands a partial with several placeholders, each
// resolved against the inherited parent variables in a single second pass.
func ExampleTemplate_multipleVars() {
	fmt.Println(renderTpl(`{{ tpl | template }}`, map[string]any{
		"tpl":      "{{ greeting }}, {{ name }}!",
		"greeting": "Hi",
		"name":     "Ada",
	}))
	// Output: Hi, Ada!
}

// ExampleTemplate_plainText passes a string with no markup straight through, so
// a partial that happens to contain no placeholders renders unchanged.
func ExampleTemplate_plainText() {
	fmt.Println(renderTpl(`{{ tpl | template }}`, map[string]any{
		"tpl": "no markup here",
	}))
	// Output: no markup here
}

// ExampleTemplate_depthGuard shows a template that renders itself. Each pass
// re-enters the engine, so the chain is stopped at the maximum nesting depth
// and the guard error surfaces instead of recursing forever.
func ExampleTemplate_depthGuard() {
	_, err := sintax.New(
		sintax.WithContextualModifiers(render.ContextualModifiers()),
	).Render(`{{ self | template }}`, map[string]any{
		"self": "{{ self | template }}",
	})
	fmt.Println(errors.Is(err, sintax.ErrMaxDepthExceeded))
	// Output: true
}

// ExampleTemplate_fromFile reads a partial from disk with the file modifier and
// pipes it into template, so the file's own markup is rendered against the
// parent's variables. This is the include primitive, and it is a composition
// rather than a feature. file brings the text in, and template expands it.
func ExampleTemplate_fromFile() {
	out, err := sintax.New(
		sintax.WithModifiers(fs.Modifiers([]string{"testdata"})),
		sintax.WithContextualModifiers(render.ContextualModifiers()),
	).Render(
		`{{ p | file | template }}`,
		map[string]any{"p": "partial.tpl", "name": "Alice", "count": 3},
	)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Println(out)
	// Output: Hello, Alice! You have 3 new messages.
}
