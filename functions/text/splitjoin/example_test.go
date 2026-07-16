package splitjoin_test

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/text/splitjoin"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(sintax.WithModifiers(splitjoin.Modifiers())).Render(tpl, vars)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return show(out)
}

// show renders a modifier result the way the docs display it. Composite values
// (maps and slices) are rendered as indented JSON so the generated docs ingest
// structured output plainly instead of Go's map[...] form, and scalars keep
// their plain string form.
func show(out any) string {
	if out != nil {
		v := reflect.ValueOf(out)
		k := v.Kind()
		isBytes := (k == reflect.Slice || k == reflect.Array) && v.Type().Elem().Kind() == reflect.Uint8
		if (k == reflect.Map || k == reflect.Slice || k == reflect.Array) && !isBytes {
			if b, err := json.MarshalIndent(out, "", "  "); err == nil {
				return string(b)
			}
		}
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

// ExampleJoinAny_space joins the elements of a slice with a multi-character
// separator.
func ExampleJoinAny_space() {
	fmt.Println(render(`{{ words | join:' - ' }}`, map[string]any{
		"words": []string{"one", "two", "three"},
	}))
	// Output: one - two - three
}

// ExampleJoinAny_pipe joins the elements with a pipe separator, a common way to
// render a slice as a single delimited field.
func ExampleJoinAny_pipe() {
	fmt.Println(render(`{{ parts | join:' | ' }}`, map[string]any{
		"parts": []string{"foo", "bar", "baz"},
	}))
	// Output: foo | bar | baz
}

// ExampleJoinAny_nonString reports an error when a slice element is not a
// string.
func ExampleJoinAny_nonString() {
	fmt.Println(render(`{{ nums | join:',' }}`, map[string]any{
		"nums": []any{"1", 2, "3"},
	}))
	// Output: error: failed to render template: failed to render variable token 'nums': modifier "join": function failed to apply: join expected an array of strings, got int at index 1
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

// ExampleJoinDefault_single joins a single-element slice, leaving it as just
// that element with no separator added.
func ExampleJoinDefault_single() {
	fmt.Println(render(`{{ items | join }}`, map[string]any{
		"items": []string{"only"},
	}))
	// Output: only
}

// ExampleJoinDefault_pair joins a two-element slice, placing each element on its
// own line since the default separator is a newline.
func ExampleJoinDefault_pair() {
	fmt.Println(render(`{{ items | join }}`, map[string]any{
		"items": []string{"line one", "line two"},
	}))
	// Output: line one
	// line two
}

// ExampleSplit splits the value on every occurrence of the separator.
func ExampleSplit() {
	fmt.Println(render(`{{ csv | split:',' }}`, map[string]any{
		"csv": "Alice,42,admin",
	}))
	// Output: [
	//   "Alice",
	//   "42",
	//   "admin"
	// ]
}

// ExampleSplit_emptySeparator splits the value into one element per UTF-8 rune
// when the separator is empty.
func ExampleSplit_emptySeparator() {
	fmt.Println(render(`{{ word | split:'' }}`, map[string]any{
		"word": "café",
	}))
	// Output: [
	//   "c",
	//   "a",
	//   "f",
	//   "é"
	// ]
}

// ExampleSplit_noMatch returns the value as a single element when the separator
// never appears.
func ExampleSplit_noMatch() {
	fmt.Println(render(`{{ value | split:',' }}`, map[string]any{
		"value": "no-commas-here",
	}))
	// Output: [
	//   "no-commas-here"
	// ]
}

// ExampleSplit_trailing yields an empty trailing element when the value ends
// with the separator.
func ExampleSplit_trailing() {
	fmt.Println(render(`{{ path | split:'/' }}`, map[string]any{
		"path": "a/b/",
	}))
	// Output: [
	//   "a",
	//   "b",
	//   ""
	// ]
}

// ExampleLinesString splits the value into lines on "\n".
func ExampleLinesString() {
	fmt.Println(render(`{{ body | lines }}`, map[string]any{
		"body": "first\nsecond\nthird",
	}))
	// Output: [
	//   "first",
	//   "second",
	//   "third"
	// ]
}

// ExampleLinesString_single returns a single-element slice when the value has
// no newline.
func ExampleLinesString_single() {
	fmt.Println(render(`{{ body | lines }}`, map[string]any{
		"body": "just one line",
	}))
	// Output: [
	//   "just one line"
	// ]
}

// ExampleLinesString_trailingNewline yields a final empty element when the
// value ends with a newline.
func ExampleLinesString_trailingNewline() {
	fmt.Println(render(`{{ body | lines }}`, map[string]any{
		"body": "a\nb\n",
	}))
	// Output: [
	//   "a",
	//   "b",
	//   ""
	// ]
}

// ExampleLinesString_blankLines keeps blank lines as empty elements, so the
// line count matches the source.
func ExampleLinesString_blankLines() {
	fmt.Println(render(`{{ body | lines }}`, map[string]any{
		"body": "alpha\n\ngamma",
	}))
	// Output: [
	//   "alpha",
	//   "",
	//   "gamma"
	// ]
}
