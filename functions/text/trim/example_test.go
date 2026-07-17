package trim_test

import (
	"fmt"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/text/trim"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(sintax.WithModifiers(trim.Modifiers())).Render(tpl, vars)
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

// ExampleTrimString_noWhitespaceUnchanged leaves a value that has no surrounding
// whitespace exactly as it is.
func ExampleTrimString_noWhitespaceUnchanged() {
	fmt.Println(render(`{{ input | trim }}`, map[string]any{
		"input": "Alice",
	}))
	// Output: Alice
}

// ExampleTrimString_innerWhitespaceKept trims only the ends and keeps whitespace
// between words.
func ExampleTrimString_innerWhitespaceKept() {
	fmt.Println(render(`{{ input | trim }}`, map[string]any{
		"input": "  a b  ",
	}))
	// Output: a b
}

// ExampleTrimStringSet strips any of the given cutset characters from both ends.
func ExampleTrimStringSet() {
	fmt.Println(render(`{{ input | trim:'/' }}`, map[string]any{
		"input": "/api/users/",
	}))
	// Output: api/users
}

// ExampleTrimStringSet_setSemantics treats the cutset as a set of characters, not
// a fixed prefix or suffix, so any of them are stripped from either end.
func ExampleTrimStringSet_setSemantics() {
	fmt.Println(render(`{{ input | trim:'xy' }}`, map[string]any{
		"input": "xy-hello-yx",
	}))
	// Output: -hello-
}

// ExampleTrimPrefixString strips leading whitespace when trim_prefix is called
// without an argument, leaving trailing whitespace in place.
func ExampleTrimPrefixString() {
	fmt.Println(render(`{{ input | trim_prefix }}`, map[string]any{
		"input": "   Welcome aboard.",
	}))
	// Output: Welcome aboard.
}

// ExampleTrimPrefixStringArg removes the given prefix once from the start.
func ExampleTrimPrefixStringArg() {
	fmt.Println(render(`{{ path | trim_prefix:'/api' }}`, map[string]any{
		"path": "/api/users",
	}))
	// Output: /users
}

// ExampleTrimPrefixStringArg_matchedOnce removes the prefix a single time, so a
// repeated leading segment is only stripped once.
func ExampleTrimPrefixStringArg_matchedOnce() {
	fmt.Println(render(`{{ path | trim_prefix:'/' }}`, map[string]any{
		"path": "//api",
	}))
	// Output: /api
}

// ExampleTrimPrefixStringArg_absent leaves the value unchanged when it does not
// start with the given prefix.
func ExampleTrimPrefixStringArg_absent() {
	fmt.Println(render(`{{ path | trim_prefix:'/' }}`, map[string]any{
		"path": "api/v1/users",
	}))
	// Output: api/v1/users
}

// ExampleTrimSuffixString strips trailing whitespace when trim_suffix is called
// without an argument, leaving leading whitespace in place.
func ExampleTrimSuffixString() {
	fmt.Println(render(`{{ input | trim_suffix }}`, map[string]any{
		"input": "Welcome aboard.   ",
	}))
	// Output: Welcome aboard.
}

// ExampleTrimSuffixStringArg removes the given suffix once from the end.
func ExampleTrimSuffixStringArg() {
	fmt.Println(render(`{{ file | trim_suffix:'.txt' }}`, map[string]any{
		"file": "report.txt",
	}))
	// Output: report
}

// ExampleTrimSuffixStringArg_absent leaves the value unchanged when it does not
// end with the given suffix.
func ExampleTrimSuffixStringArg_absent() {
	fmt.Println(render(`{{ file | trim_suffix:'.txt' }}`, map[string]any{
		"file": "report.md",
	}))
	// Output: report.md
}

// ExampleTrimSuffixStringArg_chained strips a leading and a trailing slash by
// piping trim_prefix into trim_suffix.
func ExampleTrimSuffixStringArg_chained() {
	fmt.Println(render(`{{ path | trim_prefix:'/' | trim_suffix:'/' }}`, map[string]any{
		"path": "/users/",
	}))
	// Output: users
}
