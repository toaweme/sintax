package transform_test

import (
	"fmt"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/collections/transform"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(transform.Modifiers()).Render(tpl, vars)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return fmt.Sprintf("%v", out)
}

// ExampleFlatten flattens a slice by exactly one level, spreading any element
// that is itself a slice into the result.
func ExampleFlatten() {
	fmt.Println(render(`{{ groups | flatten }}`, map[string]any{
		"groups": []any{
			[]any{"a", "b"},
			"c",
			[]any{"d"},
		},
	}))
	// Output: [a b c d]
}

// ExampleSortAsc sorts a slice ascending, the default direction when none is
// given.
func ExampleSortAsc() {
	fmt.Println(render(`{{ names | sort }}`, map[string]any{
		"names": []any{"Charlie", "Alice", "Bob"},
	}))
	// Output: [Alice Bob Charlie]
}

// ExampleSortDir sorts a copy of a slice in the named direction, here descending.
func ExampleSortDir() {
	fmt.Println(render(`{{ scores | sort:'desc' }}`, map[string]any{
		"scores": []any{72, 95, 88},
	}))
	// Output: [95 88 72]
}

// ExampleMap converts a slice of string-keyed maps into a single map keyed by the
// named field's value.
func ExampleMap() {
	fmt.Println(render(`{{ users | map:'id' }}`, map[string]any{
		"users": []map[string]any{
			{"id": "u1", "name": "Alice"},
			{"id": "u2", "name": "Bob"},
		},
	}))
	// Output: map[u1:map[id:u1 name:Alice] u2:map[id:u2 name:Bob]]
}

// ExampleMerge keys a slice of maps by the named field exactly as Map does.
func ExampleMerge() {
	fmt.Println(render(`{{ users | merge:'id' }}`, map[string]any{
		"users": []map[string]any{
			{"id": "u1", "name": "Alice"},
			{"id": "u2", "name": "Bob"},
		},
	}))
	// Output: map[u1:map[id:u1 name:Alice] u2:map[id:u2 name:Bob]]
}

// ExampleSumElements adds up the elements of a slice, returning a number.
func ExampleSumElements() {
	fmt.Println(render(`{{ prices | sum }}`, map[string]any{
		"prices": []any{10, 8, 4},
	}))
	// Output: 22
}

// ExampleSumField totals the named field across a slice of maps.
func ExampleSumField() {
	fmt.Println(render(`{{ items | sum:'price' }}`, map[string]any{
		"items": []any{
			map[string]any{"name": "Mug", "price": 10},
			map[string]any{"name": "Cup", "price": 8},
		},
	}))
	// Output: 18
}
