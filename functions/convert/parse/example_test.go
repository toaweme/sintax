package parse_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions"
	"github.com/toaweme/sintax/functions/convert/parse"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(sintax.WithModifiers(parse.Modifiers())).Render(tpl, vars)
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

// ExampleFromJSON parses a JSON object string into a map, so a serialized
// payload becomes data that later template steps can index into.
func ExampleFromJSON() {
	fmt.Println(render(`{{ body | from_json }}`, map[string]any{
		"body": `{"name": "Alice", "role": "admin"}`,
	}))
	// Output:
	// {
	//   "name": "Alice",
	//   "role": "admin"
	// }
}

// ExampleFromJSON_numbers shows that JSON numbers decode to native int64 and
// float64, so a value with a decimal point stays a float and downstream
// numeric modifiers see real numbers.
func ExampleFromJSON_numbers() {
	fmt.Println(render(`{{ body | from_json }}`, map[string]any{
		"body": `{"count": 3, "ratio": 1.5}`,
	}))
	// Output:
	// {
	//   "count": 3,
	//   "ratio": 1.5
	// }
}

// ExampleFromJSON_nested parses a JSON object that contains a nested object and
// an array, keeping the structure intact for later indexing.
func ExampleFromJSON_nested() {
	fmt.Println(render(`{{ body | from_json }}`, map[string]any{
		"body": `{"user": {"id": 7}, "scores": [1, 2.5]}`,
	}))
	// Output:
	// {
	//   "scores": [
	//     1,
	//     2.5
	//   ],
	//   "user": {
	//     "id": 7
	//   }
	// }
}

// ExampleFromCSV parses a CSV string into a list of rows keyed by the header,
// treating the first record as the header row.
func ExampleFromCSV() {
	fmt.Println(render(`{{ body | from_csv }}`, map[string]any{
		"body": "name,age\nAlice,30\nBob,25",
	}))
	// Output:
	// [
	//   {
	//     "age": "30",
	//     "name": "Alice"
	//   },
	//   {
	//     "age": "25",
	//     "name": "Bob"
	//   }
	// ]
}

// ExampleFromCSV_headerOnly shows that a CSV with only a header row and no data
// rows parses into an empty list.
func ExampleFromCSV_headerOnly() {
	fmt.Println(render(`{{ body | from_csv }}`, map[string]any{
		"body": "name,age\n",
	}))
	// Output: []
}

// ExampleFromCSV_shortRow shows that a row with fewer cells than the header
// pads the missing columns with an empty string.
func ExampleFromCSV_shortRow() {
	fmt.Println(render(`{{ body | from_csv }}`, map[string]any{
		"body": "a,b,c\n1,2",
	}))
	// Output:
	// [
	//   {
	//     "a": "1",
	//     "b": "2",
	//     "c": ""
	//   }
	// ]
}

// renderInjected renders against the modifier set with one entry replaced, the
// way an application wires up from_yaml. It is the runnable form of the setup
// that modifier's doc comment describes, so the example below documents a bound
// codec rather than the "needs to be injected" error, which would teach a reader
// nothing about what the modifier does.
//
// The parser it binds is a deliberately tiny stand-in. sintax ships no
// third-party dependencies, and that is the whole reason from_yaml is a stub, so
// an example cannot import a real codec to prove the wiring. Only the wiring is
// the point, and a real build swaps in gopkg.in/yaml.v3 at exactly this seam.
func renderInjected(name functions.ModifierName, impl functions.GlobalModifier, tpl string, vars map[string]any) string {
	mods := parse.Modifiers()
	mods[string(name)] = impl
	out, err := sintax.New(sintax.WithModifiers(mods)).Render(tpl, vars)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return show(out)
}

// miniYAML parses flat `key: value` lines, enough to show the modifier's shape.
// A real codec handles nesting, lists, types, and quoting this does not.
func miniYAML(doc string) (map[string]any, error) {
	out := map[string]any{}
	for _, line := range strings.Split(strings.TrimSpace(doc), "\n") {
		k, v, ok := strings.Cut(line, ":")
		if !ok {
			return nil, fmt.Errorf("miniYAML wants key: value, got %q", line)
		}
		out[strings.TrimSpace(k)] = strings.TrimSpace(v)
	}
	return out, nil
}

// ExampleFromYAML parses a YAML document into a map once a codec is injected, so
// a config file becomes data that later template steps can index into.
func ExampleFromYAML() {
	fmt.Println(renderInjected(parse.ModifierNameFromYAML, functions.Wrap(miniYAML),
		`{{ body | from_yaml }}`, map[string]any{
			"body": "host: localhost\nregion: eu-west-1",
		}))
	// Output:
	// {
	//   "host": "localhost",
	//   "region": "eu-west-1"
	// }
}

// ExampleFromYAML_notInjected shows what the modifier does before a codec is
// bound. The stub errors rather than silently parsing nothing, so a missing
// injection surfaces at render time.
func ExampleFromYAML_notInjected() {
	fmt.Println(render(`{{ body | from_yaml }}`, map[string]any{
		"body": "host: localhost",
	}))
	// Output: error: failed to render template: failed to render variable token 'body': function failed to apply: from_yaml function needs to be injected
}
