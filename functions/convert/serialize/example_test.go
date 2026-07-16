package serialize_test

import (
	"fmt"
	"sort"
	"strings"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions"
	"github.com/toaweme/sintax/functions/convert/serialize"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(sintax.WithModifiers(serialize.Modifiers())).Render(tpl, vars)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return fmt.Sprintf("%v", out)
}

// renderInjected renders against the modifier set with one entry replaced, the
// way an application wires up yaml or markdown. It is the runnable form of the
// setup each stub's doc comment describes, so the examples below document a
// bound codec rather than the "needs to be injected" error, which teaches a
// reader nothing about what the modifier does.
//
// The encoders it binds are deliberately tiny stand-ins. sintax ships no
// third-party dependencies, and that is the whole reason these two modifiers are
// stubs, so an example cannot import a real codec to prove the wiring. Only the
// wiring is the point, and a real build swaps in gopkg.in/yaml.v3 or an
// HTML-to-Markdown library at exactly this seam.
func renderInjected(name functions.ModifierName, impl functions.GlobalModifier, tpl string, vars map[string]any) string {
	mods := serialize.Modifiers()
	mods[string(name)] = impl
	out, err := sintax.New(sintax.WithModifiers(mods)).Render(tpl, vars)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return fmt.Sprintf("%v", out)
}

// miniYAML encodes a flat map as YAML, enough to show the modifier's shape. A
// real codec handles nesting, quoting, and types this does not.
func miniYAML(v any) (string, error) {
	m, ok := v.(map[string]any)
	if !ok {
		return "", fmt.Errorf("miniYAML wants a map, got %T", v)
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var b strings.Builder
	for i, k := range keys {
		if i > 0 {
			b.WriteByte('\n')
		}
		fmt.Fprintf(&b, "%s: %v", k, m[k])
	}
	return b.String(), nil
}

// miniMarkdown converts a couple of HTML tags to Markdown, enough to show the
// modifier's shape. A real converter parses HTML properly.
func miniMarkdown(html string) (string, error) {
	for _, r := range []struct{ from, to string }{
		{"<h1>", "# "}, {"</h1>", "\n"},
		{"<strong>", "**"}, {"</strong>", "**"},
		{"<p>", ""}, {"</p>", ""},
	} {
		html = strings.ReplaceAll(html, r.from, r.to)
	}
	return strings.TrimSpace(html), nil
}

// ExampleJSON serializes a value to compact JSON with keys sorted alphabetically,
// so the output stays stable regardless of the input map's insertion order.
func ExampleJSON() {
	fmt.Println(render(`{{ user | json }}`, map[string]any{
		"user": map[string]any{"role": "admin", "name": "Alice"},
	}))
	// Output: {"name":"Alice","role":"admin"}
}

// ExampleJSON_slice serializes a list, preserving its order in a JSON array.
func ExampleJSON_slice() {
	fmt.Println(render(`{{ tags | json }}`, map[string]any{
		"tags": []any{"go", "rust"},
	}))
	// Output: ["go","rust"]
}

// ExampleJSON_string wraps a bare string in JSON quotes.
func ExampleJSON_string() {
	fmt.Println(render(`{{ name | json }}`, map[string]any{
		"name": "Alice",
	}))
	// Output: "Alice"
}

// ExampleJSON_null renders a nil value as the JSON null literal.
func ExampleJSON_null() {
	fmt.Println(render(`{{ missing | json }}`, map[string]any{
		"missing": nil,
	}))
	// Output: null
}

// ExampleJSONMode selects indented output with the literal 'pretty' mode, using
// two spaces per level.
func ExampleJSONMode() {
	fmt.Println(render(`{{ user | json:'pretty' }}`, map[string]any{
		"user": map[string]any{"name": "Alice"},
	}))
	// Output: {
	//   "name": "Alice"
	// }
}

// ExampleJSONMode_slice indents a JSON array, placing one element per line.
func ExampleJSONMode_slice() {
	fmt.Println(render(`{{ scores | json:'pretty' }}`, map[string]any{
		"scores": []any{1, 2},
	}))
	// Output: [
	//   1,
	//   2
	// ]
}

// ExampleJSONMode_compactFallback shows that any mode other than 'pretty' falls
// back to compact output.
func ExampleJSONMode_compactFallback() {
	fmt.Println(render(`{{ scores | json:'inline' }}`, map[string]any{
		"scores": []any{1, 2},
	}))
	// Output: [1,2]
}

// ExampleYAML serializes a map as YAML once a codec is injected, so a map
// assembled in a template can be written out as a config file.
func ExampleYAML() {
	fmt.Println(renderInjected(serialize.ModifierNameYAML, functions.Wrap(miniYAML),
		`{{ config | yaml }}`, map[string]any{
			"config": map[string]any{"host": "localhost", "port": 8080},
		}))
	// Output:
	// host: localhost
	// port: 8080
}

// ExampleYAML_notInjected shows what the modifier does before a codec is bound.
// The stub errors rather than guessing a format, so a missing injection surfaces
// at render time instead of silently emitting nothing.
func ExampleYAML_notInjected() {
	fmt.Println(render(`{{ config | yaml }}`, map[string]any{
		"config": map[string]any{"host": "localhost"},
	}))
	// Output: error: failed to render template: failed to render variable token 'config': function failed to apply: yaml function needs to be injected
}

// ExampleMarkdown converts HTML to Markdown once a converter is injected, so a
// rich-text field can be reduced to plain prose.
func ExampleMarkdown() {
	fmt.Println(renderInjected(serialize.ModifierNameMarkdown, functions.Wrap(miniMarkdown),
		`{{ body | markdown }}`, map[string]any{
			"body": "<h1>Title</h1><p>Some <strong>bold</strong> prose.</p>",
		}))
	// Output:
	// # Title
	// Some **bold** prose.
}
