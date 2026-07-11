package parse_test

import (
	"fmt"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/convert/parse"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(parse.Modifiers()).Render(tpl, vars)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return fmt.Sprintf("%v", out)
}

// ExampleFrom parses a JSON object string into a map, so a serialized payload
// becomes data that later template steps can index into.
func ExampleFrom() {
	fmt.Println(render(`{{ body | from:'json' }}`, map[string]any{
		"body": `{"name": "Alice", "role": "admin"}`,
	}))
	// Output: map[name:Alice role:admin]
}

// ExampleFrom_csv parses a CSV string into a list of rows keyed by the header,
// treating the first record as the header row.
func ExampleFrom_csv() {
	fmt.Println(render(`{{ body | from:'csv' }}`, map[string]any{
		"body": "name,age\nAlice,30\nBob,25",
	}))
	// Output: [map[age:30 name:Alice] map[age:25 name:Bob]]
}
