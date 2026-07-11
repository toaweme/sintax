package edit_test

import (
	"fmt"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/text/edit"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(edit.Modifiers()).Render(tpl, vars)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return fmt.Sprintf("%v", out)
}

// ExampleConcat joins the value and every part into one string, in order, with
// no separator.
func ExampleConcat() {
	fmt.Println(render(`{{ name | concat:'-','01','.txt' }}`, map[string]any{
		"name": "report",
	}))
	// Output: report-01.txt
}

// ExampleWrap nests the value inside a new single-entry map under the given key.
func ExampleWrap() {
	fmt.Println(render(`{{ value | wrap:'name' }}`, map[string]any{
		"value": "Ada",
	}))
	// Output: map[name:Ada]
}

// ExampleReplace replaces every occurrence of the old substring with the
// replacement.
func ExampleReplace() {
	fmt.Println(render(`{{ text | replace:'world','everyone' }}`, map[string]any{
		"text": "Hello world",
	}))
	// Output: Hello everyone
}

// ExampleReplacePattern replaces every match of an RE2 pattern, and the
// replacement may reference capture groups with $1 and $2.
func ExampleReplacePattern() {
	fmt.Println(render(`{{ name | replace_pattern:'(\w+), (\w+)','$2 $1' }}`, map[string]any{
		"name": "Doe, Jane",
	}))
	// Output: Jane Doe
}

// ExampleReverse reverses the value by rune, so multi-byte characters stay
// intact.
func ExampleReverse() {
	fmt.Println(render(`{{ word | reverse }}`, map[string]any{
		"word": "café",
	}))
	// Output: éfac
}

// ExampleShorten truncates the value to at most the given number of bytes.
func ExampleShorten() {
	fmt.Println(render(`{{ text | shorten:5 }}`, map[string]any{
		"text": "Hello world",
	}))
	// Output: Hello
}
