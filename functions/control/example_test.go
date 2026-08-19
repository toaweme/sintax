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

// ExampleWhen renders a flag as the word it means, so a template says what a
// boolean stands for in one expression.
func ExampleWhen() {
	fmt.Println(render(`{{ published | when:'live','draft' }}`, map[string]any{
		"published": true,
	}))
	// Output: live
}

// ExampleWhen_falsy takes the second branch for a value a template `if` would
// read as false, an empty string included.
func ExampleWhen_falsy() {
	fmt.Println(render(`{{ title | when:'named','untitled' }}`, map[string]any{
		"title": "",
	}))
	// Output: untitled
}

// ExampleWhen_missing shows a miss reaching when as falsey and answered there,
// so the falsy branch stands in for absent data and nothing downstream sees the
// miss.
func ExampleWhen_missing() {
	fmt.Println(render(`{{ enabled | when:'on','off' }}`, map[string]any{}))
	// Output: off
}
