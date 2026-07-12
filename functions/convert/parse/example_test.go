package parse_test

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/convert/parse"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(parse.Modifiers()).Render(tpl, vars)
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

// ExampleFrom parses a JSON object string into a map, so a serialized payload
// becomes data that later template steps can index into.
func ExampleFrom() {
	fmt.Println(render(`{{ body | from:'json' }}`, map[string]any{
		"body": `{"name": "Alice", "role": "admin"}`,
	}))
	// Output:
	// {
	//   "name": "Alice",
	//   "role": "admin"
	// }
}

// ExampleFrom_numbers shows that JSON numbers decode to native int64 and
// float64, so a value with a decimal point stays a float and downstream
// numeric modifiers see real numbers.
func ExampleFrom_numbers() {
	fmt.Println(render(`{{ body | from:'json' }}`, map[string]any{
		"body": `{"count": 3, "ratio": 1.5}`,
	}))
	// Output:
	// {
	//   "count": 3,
	//   "ratio": 1.5
	// }
}

// ExampleFrom_nested parses a JSON object that contains a nested object and an
// array, keeping the structure intact for later indexing.
func ExampleFrom_nested() {
	fmt.Println(render(`{{ body | from:'json' }}`, map[string]any{
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

// ExampleFrom_csv parses a CSV string into a list of rows keyed by the header,
// treating the first record as the header row.
func ExampleFrom_csv() {
	fmt.Println(render(`{{ body | from:'csv' }}`, map[string]any{
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

// ExampleFrom_csvHeaderOnly shows that a CSV with only a header row and no data
// rows parses into an empty list.
func ExampleFrom_csvHeaderOnly() {
	fmt.Println(render(`{{ body | from:'csv' }}`, map[string]any{
		"body": "name,age\n",
	}))
	// Output: []
}

// ExampleFrom_csvShortRow shows that a row with fewer cells than the header pads
// the missing columns with an empty string.
func ExampleFrom_csvShortRow() {
	fmt.Println(render(`{{ body | from:'csv' }}`, map[string]any{
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
