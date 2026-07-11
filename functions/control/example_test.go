package control_test

import (
	"fmt"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/control"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(control.Modifiers()).Render(tpl, vars)
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
