package serialize_test

import (
	"fmt"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/convert/serialize"
)

// The yaml and markdown modifiers ship as injectable stubs that return an error
// until a consumer supplies an implementation, so they have no runnable example.
// Only json produces deterministic output through serialize.Modifiers().

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(serialize.Modifiers()).Render(tpl, vars)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return fmt.Sprintf("%v", out)
}

// ExampleJSON serializes a value to compact JSON with keys sorted alphabetically,
// so the output stays stable regardless of the input map's insertion order.
func ExampleJSON() {
	fmt.Println(render(`{{ user | json }}`, map[string]any{
		"user": map[string]any{"role": "admin", "name": "Alice"},
	}))
	// Output: {"name":"Alice","role":"admin"}
}

// ExampleJSONMode selects indented output with the literal 'pretty' mode, using
// two spaces per level.
func ExampleJSONMode() {
	fmt.Println(render(`{{ user | json:'pretty' }}`, map[string]any{
		"user": map[string]any{"name": "Alice"},
	}))
	// Output: {
	//   "name": "Alice"
	// }
}
