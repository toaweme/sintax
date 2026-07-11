package splitjoin_test

import (
	"fmt"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/text/splitjoin"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(splitjoin.Modifiers()).Render(tpl, vars)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return fmt.Sprintf("%v", out)
}

// ExampleJoinAny joins the elements of a slice into one string with the given
// separator.
func ExampleJoinAny() {
	fmt.Println(render(`{{ tags | join:',' }}`, map[string]any{
		"tags": []string{"coffee", "sale", "new"},
	}))
	// Output: coffee,sale,new
}

// ExampleJoinDefault joins the elements of a slice on a newline when no
// separator is given.
func ExampleJoinDefault() {
	fmt.Println(render(`{{ items | join }}`, map[string]any{
		"items": []string{"a", "b", "c"},
	}))
	// Output: a
	// b
	// c
}

// ExampleSplit splits the value on every occurrence of the separator.
func ExampleSplit() {
	fmt.Println(render(`{{ csv | split:',' }}`, map[string]any{
		"csv": "Alice,42,admin",
	}))
	// Output: [Alice 42 admin]
}

// ExampleLinesString splits the value into lines on "\n".
func ExampleLinesString() {
	fmt.Println(render(`{{ body | lines }}`, map[string]any{
		"body": "first\nsecond\nthird",
	}))
	// Output: [first second third]
}
