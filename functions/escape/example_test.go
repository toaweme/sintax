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

// ExampleHTML_apostrophe escapes a single quote so the value is safe inside a single-quoted attribute.
func ExampleHTML_apostrophe() {
	fmt.Println(render(`{{ label | escape_html }}`, map[string]any{
		"label": "it's a trap",
	}))
	// Output: it&#39;s a trap
}

// ExampleHTML_unicode leaves multi-byte characters untouched and only defuses HTML metacharacters.
func ExampleHTML_unicode() {
	fmt.Println(render(`{{ note | escape_html }}`, map[string]any{
		"note": "café ☕ 5 < 6",
	}))
	// Output: café ☕ 5 &lt; 6
}

// ExampleHTML_number coerces a numeric value to its string form before escaping.
func ExampleHTML_number() {
	fmt.Println(render(`{{ count | escape_html }}`, map[string]any{
		"count": 42,
	}))
	// Output: 42
}

// ExampleURL escapes text for use as a query-string value.
func ExampleURL() {
	fmt.Println(render(`{{ q | escape_url }}`, map[string]any{
		"q": "a b&c=d",
	}))
	// Output: a+b%26c%3Dd
}

// ExampleURL_slash percent-encodes a slash so a value stays a single query parameter.
func ExampleURL_slash() {
	fmt.Println(render(`{{ path | escape_url }}`, map[string]any{
		"path": "docs/getting-started",
	}))
	// Output: docs%2Fgetting-started
}

// ExampleURL_unicode percent-encodes each byte of a multi-byte character.
func ExampleURL_unicode() {
	fmt.Println(render(`{{ term | escape_url }}`, map[string]any{
		"term": "café",
	}))
	// Output: caf%C3%A9
}

// ExampleURL_percent percent-encodes a literal percent sign and space so the value survives a query string.
func ExampleURL_percent() {
	fmt.Println(render(`{{ promo | escape_url }}`, map[string]any{
		"promo": "50% off!",
	}))
	// Output: 50%25+off%21
}

// ExampleJS escapes text for embedding inside a quoted JavaScript string.
func ExampleJS() {
	fmt.Println(render(`{{ code | escape_js }}`, map[string]any{
		"code": `x = "y";` + "\n",
	}))
	// Output: x = \"y\";\n
}

// ExampleJS_scriptTag escapes the angle brackets and slash so the value cannot close a script element.
func ExampleJS_scriptTag() {
	fmt.Println(render(`{{ payload | escape_js }}`, map[string]any{
		"payload": "</script>",
	}))
	// Output: \u003C\/script\u003E
}

// ExampleJS_apostrophe escapes a single quote so the value is safe inside a single-quoted string literal.
func ExampleJS_apostrophe() {
	fmt.Println(render(`{{ msg | escape_js }}`, map[string]any{
		"msg": "it's fine",
	}))
	// Output: it\'s fine
}

// ExampleJS_unicode leaves ordinary multi-byte characters untouched inside a string literal.
func ExampleJS_unicode() {
	fmt.Println(render(`{{ msg | escape_js }}`, map[string]any{
		"msg": "café",
	}))
	// Output: café
}
