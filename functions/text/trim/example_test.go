package trim_test

import (
	"fmt"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/text/trim"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(trim.Modifiers()).Render(tpl, vars)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return fmt.Sprintf("%v", out)
}

// ExampleTrimString strips leading and trailing whitespace when trim is called
// without a cutset.
func ExampleTrimString() {
	fmt.Println(render(`{{ input | trim }}`, map[string]any{
		"input": "  hello  ",
	}))
	// Output: hello
}

// ExampleTrimStringSet strips any of the given cutset characters from both ends.
func ExampleTrimStringSet() {
	fmt.Println(render(`{{ input | trim:'/' }}`, map[string]any{
		"input": "/api/users/",
	}))
	// Output: api/users
}

// ExampleTrimPrefixStringArg removes the given prefix once from the start.
func ExampleTrimPrefixStringArg() {
	fmt.Println(render(`{{ path | trim_prefix:'/api' }}`, map[string]any{
		"path": "/api/users",
	}))
	// Output: /users
}

// ExampleTrimSuffixStringArg removes the given suffix once from the end.
func ExampleTrimSuffixStringArg() {
	fmt.Println(render(`{{ file | trim_suffix:'.txt' }}`, map[string]any{
		"file": "report.txt",
	}))
	// Output: report
}
