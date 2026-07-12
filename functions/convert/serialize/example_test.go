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

// ExampleJSON_slice serializes a list, preserving its order in a JSON array.
func ExampleJSON_slice() {
	fmt.Println(render(`{{ tags | json }}`, map[string]any{
		"tags": []any{"go", "rust"},
	}))
	// Output: ["go","rust"]
}

// ExampleJSON_string wraps a bare string in JSON quotes.
func ExampleJSON_string() {
	fmt.Println(render(`{{ name | json }}`, map[string]any{
		"name": "Alice",
	}))
	// Output: "Alice"
}

// ExampleJSON_null renders a nil value as the JSON null literal.
func ExampleJSON_null() {
	fmt.Println(render(`{{ missing | json }}`, map[string]any{
		"missing": nil,
	}))
	// Output: null
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

// ExampleJSONMode_slice indents a JSON array, placing one element per line.
func ExampleJSONMode_slice() {
	fmt.Println(render(`{{ scores | json:'pretty' }}`, map[string]any{
		"scores": []any{1, 2},
	}))
	// Output: [
	//   1,
	//   2
	// ]
}

// ExampleJSONMode_compactFallback shows that any mode other than 'pretty' falls
// back to compact output.
func ExampleJSONMode_compactFallback() {
	fmt.Println(render(`{{ scores | json:'inline' }}`, map[string]any{
		"scores": []any{1, 2},
	}))
	// Output: [1,2]
}
