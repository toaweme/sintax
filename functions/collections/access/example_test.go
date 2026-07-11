package access_test

import (
	"fmt"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/collections/access"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(access.Modifiers()).Render(tpl, vars)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return fmt.Sprintf("%v", out)
}

// ExamplePluck reads one named field from every element of a slice of maps and
// returns the collected values in order.
func ExamplePluck() {
	fmt.Println(render(`{{ users | pluck:'name' }}`, map[string]any{
		"users": []any{
			map[string]any{"name": "Alice", "role": "admin"},
			map[string]any{"name": "Bob", "role": "viewer"},
		},
	}))
	// Output: [Alice Bob]
}

// ExampleFirstString returns the leading byte of a string as a one-byte string.
func ExampleFirstString() {
	fmt.Println(render(`{{ word | first }}`, map[string]any{
		"word": "hello",
	}))
	// Output: h
}

// ExampleFirstSlice returns the first element of a slice as-is, keeping its type.
func ExampleFirstSlice() {
	fmt.Println(render(`{{ items | first }}`, map[string]any{
		"items": []any{"espresso", "latte", "macchiato"},
	}))
	// Output: espresso
}

// ExampleLastString returns the trailing byte of a string as a one-byte string.
func ExampleLastString() {
	fmt.Println(render(`{{ word | last }}`, map[string]any{
		"word": "hello",
	}))
	// Output: o
}

// ExampleLastSlice returns the last element of a slice as-is.
func ExampleLastSlice() {
	fmt.Println(render(`{{ items | last }}`, map[string]any{
		"items": []any{"espresso", "latte", "macchiato"},
	}))
	// Output: macchiato
}

// ExampleKey reads one value out of a map by key, walking nested maps with dot
// notation.
func ExampleKey() {
	fmt.Println(render(`{{ config | key:'database.host' }}`, map[string]any{
		"config": map[string]any{
			"database": map[string]any{"host": "db.local", "port": 5432},
		},
	}))
	// Output: db.local
}

// ExampleKey_missing shows that Key is forgiving, so a lookup that finds nothing
// renders as nil rather than raising a template error.
func ExampleKey_missing() {
	fmt.Println(render(`{{ user | key:'phone' }}`, map[string]any{
		"user": map[string]any{"name": "Alice"},
	}))
	// Output: <nil>
}

// ExampleFindSlice returns the first map in a slice whose key field equals the
// wanted value.
func ExampleFindSlice() {
	fmt.Println(render(`{{ items | find:'status','active' }}`, map[string]any{
		"items": []any{
			map[string]any{"name": "Coffee", "status": "sold-out"},
			map[string]any{"name": "Tea", "status": "active"},
		},
	}))
	// Output: map[name:Tea status:active]
}

// ExampleFindMap returns the map itself when its key field equals the wanted
// value.
func ExampleFindMap() {
	fmt.Println(render(`{{ item | find:'status','active' }}`, map[string]any{
		"item": map[string]any{"name": "Tea", "status": "active"},
	}))
	// Output: map[name:Tea status:active]
}
