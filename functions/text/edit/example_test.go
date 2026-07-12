package edit_test

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/text/edit"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(edit.Modifiers()).Render(tpl, vars)
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

// ExampleConcat joins the value and every part into one string, in order, with
// no separator.
func ExampleConcat() {
	fmt.Println(render(`{{ name | concat:'-','01','.txt' }}`, map[string]any{
		"name": "report",
	}))
	// Output: report-01.txt
}

// ExampleConcat_scalarValue takes a number value as its string form, so it can
// start a joined string.
func ExampleConcat_scalarValue() {
	fmt.Println(render(`{{ id | concat:'-draft' }}`, map[string]any{
		"id": 42,
	}))
	// Output: 42-draft
}

// ExampleConcat_numberPart accepts a number part and appends its string form.
func ExampleConcat_numberPart() {
	fmt.Println(render(`{{ base | concat:'/',2024 }}`, map[string]any{
		"base": "logs",
	}))
	// Output: logs/2024
}

// ExampleConcat_noParts returns the value unchanged when no parts are given.
func ExampleConcat_noParts() {
	fmt.Println(render(`{{ word | concat }}`, map[string]any{
		"word": "solo",
	}))
	// Output: solo
}

// ExampleWrap nests the value inside a new single-entry map under the given key.
func ExampleWrap() {
	fmt.Println(render(`{{ value | wrap:'name' }}`, map[string]any{
		"value": "Ada",
	}))
	// Output:
	// {
	//   "name": "Ada"
	// }
}

// ExampleWrap_number nests any value type, including a number, under the key.
func ExampleWrap_number() {
	fmt.Println(render(`{{ count | wrap:'total' }}`, map[string]any{
		"count": 3,
	}))
	// Output:
	// {
	//   "total": 3
	// }
}

// ExampleWrap_slice nests a slice under the key, keeping it as structured data
// rather than a string.
func ExampleWrap_slice() {
	fmt.Println(render(`{{ items | wrap:'list' }}`, map[string]any{
		"items": []any{"a", "b"},
	}))
	// Output:
	// {
	//   "list": [
	//     "a",
	//     "b"
	//   ]
	// }
}

// ExampleReplace replaces every occurrence of the old substring with the
// replacement.
func ExampleReplace() {
	fmt.Println(render(`{{ text | replace:'world','everyone' }}`, map[string]any{
		"text": "Hello world",
	}))
	// Output: Hello everyone
}

// ExampleReplace_allOccurrences replaces every match, not just the first.
func ExampleReplace_allOccurrences() {
	fmt.Println(render(`{{ text | replace:'a','o' }}`, map[string]any{
		"text": "banana",
	}))
	// Output: bonono
}

// ExampleReplace_remove deletes the substring by replacing it with an empty
// string.
func ExampleReplace_remove() {
	fmt.Println(render(`{{ text | replace:' ','' }}`, map[string]any{
		"text": "a b c",
	}))
	// Output: abc
}

// ExampleReplace_noMatch leaves the value unchanged when the old substring is
// absent.
func ExampleReplace_noMatch() {
	fmt.Println(render(`{{ text | replace:'xyz','abc' }}`, map[string]any{
		"text": "Hello",
	}))
	// Output: Hello
}

// ExampleReplacePattern replaces every match of an RE2 pattern, and the
// replacement may reference capture groups with $1 and $2.
func ExampleReplacePattern() {
	fmt.Println(render(`{{ name | replace_pattern:'(\w+), (\w+)','$2 $1' }}`, map[string]any{
		"name": "Doe, Jane",
	}))
	// Output: Jane Doe
}

// ExampleReplacePattern_delete removes every match when the replacement is
// empty.
func ExampleReplacePattern_delete() {
	fmt.Println(render(`{{ code | replace_pattern:'\d+','' }}`, map[string]any{
		"code": "abc123def",
	}))
	// Output: abcdef
}

// ExampleReplacePattern_collapse rewrites each run of whitespace to a single
// underscore.
func ExampleReplacePattern_collapse() {
	fmt.Println(render(`{{ text | replace_pattern:'\s+','_' }}`, map[string]any{
		"text": "a b  c",
	}))
	// Output: a_b_c
}

// ExampleReverse reverses the value by rune, so multi-byte characters stay
// intact.
func ExampleReverse() {
	fmt.Println(render(`{{ word | reverse }}`, map[string]any{
		"word": "café",
	}))
	// Output: éfac
}

// ExampleReverse_ascii reverses a plain ASCII string character by character.
func ExampleReverse_ascii() {
	fmt.Println(render(`{{ word | reverse }}`, map[string]any{
		"word": "abc",
	}))
	// Output: cba
}

// ExampleReverse_number reverses a number by taking its string form first.
func ExampleReverse_number() {
	fmt.Println(render(`{{ n | reverse }}`, map[string]any{
		"n": 123,
	}))
	// Output: 321
}

// ExampleShorten truncates the value to at most the given number of bytes.
func ExampleShorten() {
	fmt.Println(render(`{{ text | shorten:5 }}`, map[string]any{
		"text": "Hello world",
	}))
	// Output: Hello
}

// ExampleShorten_unchanged returns the value as-is when it is already within the
// limit.
func ExampleShorten_unchanged() {
	fmt.Println(render(`{{ text | shorten:20 }}`, map[string]any{
		"text": "Hi",
	}))
	// Output: Hi
}

// ExampleShorten_one keeps just the first byte when the limit is one.
func ExampleShorten_one() {
	fmt.Println(render(`{{ text | shorten:1 }}`, map[string]any{
		"text": "Hello",
	}))
	// Output: H
}

// ExampleShorten_scalar truncates a number by taking its string form first.
func ExampleShorten_scalar() {
	fmt.Println(render(`{{ n | shorten:3 }}`, map[string]any{
		"n": 123456,
	}))
	// Output: 123
}

// ExampleShortenParse accepts the length as a numeric string such as
// shorten:'5'.
func ExampleShortenParse() {
	fmt.Println(render(`{{ text | shorten:'5' }}`, map[string]any{
		"text": "Hello world",
	}))
	// Output: Hello
}

// ExampleShortenParse_unchanged leaves the value as-is when it is shorter than a
// numeric-string limit.
func ExampleShortenParse_unchanged() {
	fmt.Println(render(`{{ text | shorten:'20' }}`, map[string]any{
		"text": "Hi",
	}))
	// Output: Hi
}
