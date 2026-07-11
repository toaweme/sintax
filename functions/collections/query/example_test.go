package query_test

import (
	"fmt"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/collections/query"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(query.Modifiers()).Render(tpl, vars)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return fmt.Sprintf("%v", out)
}

// ExampleFilter returns the items of a slice whose named field equals the search
// value.
func ExampleFilter() {
	fmt.Println(render(`{{ items | filter:'role','admin' }}`, map[string]any{
		"items": []any{
			map[string]any{"name": "Alice", "role": "admin"},
			map[string]any{"name": "Bob", "role": "viewer"},
			map[string]any{"name": "Carol", "role": "admin"},
		},
	}))
	// Output: [map[name:Alice role:admin] map[name:Carol role:admin]]
}

// ExampleHas reports whether a plain slice contains a given element.
func ExampleHas() {
	fmt.Println(render(`{{ tags | has:'featured' }}`, map[string]any{
		"tags": []any{"featured", "sale", "new"},
	}))
	// Output: true
}

// ExampleHas_map shows that for a map with a single parameter Has tests only
// whether the key exists, so a key mapped to false still counts as present.
func ExampleHas_map() {
	fmt.Println(render(`{{ config | has:'debug_mode' }}`, map[string]any{
		"config": map[string]any{"debug_mode": false, "region": "eu-west-1"},
	}))
	// Output: true
}

// ExampleIs reports whether the value equals any one of the given candidates.
func ExampleIs() {
	fmt.Println(render(`{{ status | is:'active','pending' }}`, map[string]any{
		"status": "pending",
	}))
	// Output: true
}
