package escape_test

import (
	"fmt"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/escape"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(escape.Modifiers()).Render(tpl, vars)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return fmt.Sprintf("%v", out)
}

// ExampleHTML escapes text for an HTML text node or a quoted attribute value.
func ExampleHTML() {
	fmt.Println(render(`{{ name | escape_html }}`, map[string]any{
		"name": `Tom & "Jerry" <b>`,
	}))
	// Output: Tom &amp; &#34;Jerry&#34; &lt;b&gt;
}

// ExampleURL escapes text for use as a query-string value.
func ExampleURL() {
	fmt.Println(render(`{{ q | escape_url }}`, map[string]any{
		"q": "a b&c=d",
	}))
	// Output: a+b%26c%3Dd
}

// ExampleJS escapes text for embedding inside a quoted JavaScript string.
func ExampleJS() {
	fmt.Println(render(`{{ code | escape_js }}`, map[string]any{
		"code": `x = "y";` + "\n",
	}))
	// Output: x = \"y\";\n
}
