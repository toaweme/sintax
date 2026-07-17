package control_test

import (
	"fmt"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/control"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(sintax.WithModifiers(control.Modifiers())).Render(tpl, vars)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return fmt.Sprintf("%v", out)
}

// ExampleDefault swaps in the fallback when the piped value is absent, so a
// missing variable never renders as nothing.
func ExampleDefault() {
	fmt.Println(render(`{{ name | default:'anonymous' }}`, map[string]any{}))
	// Output: anonymous
}

// ExampleDefault_emptyString treats an empty string as "nothing there" and
// reaches for the fallback.
func ExampleDefault_emptyString() {
	fmt.Println(render(`{{ nickname | default:'anonymous' }}`, map[string]any{
		"nickname": "",
	}))
	// Output: anonymous
}

// ExampleDefault_present passes a real value straight through and never touches
// the fallback.
func ExampleDefault_present() {
	fmt.Println(render(`{{ name | default:'anonymous' }}`, map[string]any{
		"name": "Ada",
	}))
	// Output: Ada
}

// ExampleDefault_keepsZero shows that zero is a real value, so it is kept rather
// than replaced by the fallback.
func ExampleDefault_keepsZero() {
	fmt.Println(render(`{{ count | default:'n/a' }}`, map[string]any{
		"count": 0,
	}))
	// Output: 0
}
