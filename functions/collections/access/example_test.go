package access_test

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/collections/access"
	"github.com/toaweme/sintax/functions/control"
)

func render(tpl string, vars map[string]any) string {
	// the access examples also register default so they can show the documented
	// "lookup then fall back" pattern (find with no match, a missing key).
	mods := access.Modifiers()
	for name, m := range control.Modifiers() {
		mods[name] = m
	}
	out, err := sintax.New(sintax.WithModifiers(mods)).Render(tpl, vars)
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

// ExamplePluck reads one named field from every element of a slice of maps and
// returns the collected values in order.
func ExamplePluck() {
	fmt.Println(render(`{{ users | pluck:'name' }}`, map[string]any{
		"users": []any{
			map[string]any{"name": "Alice", "role": "admin"},
			map[string]any{"name": "Bob", "role": "viewer"},
		},
	}))
	// Output:
	// [
	//   "Alice",
	//   "Bob"
	// ]
}

// ExamplePluck_numbers collects a numeric field from every element, keeping the
// values as numbers.
func ExamplePluck_numbers() {
	fmt.Println(render(`{{ orders | pluck:'price' }}`, map[string]any{
		"orders": []any{
			map[string]any{"item": "Coffee", "price": 3},
			map[string]any{"item": "Cake", "price": 5},
		},
	}))
	// Output:
	// [
	//   3,
	//   5
	// ]
}

// ExamplePluck_empty returns an empty slice when there is nothing to read from,
// so the result length always tracks the input.
func ExamplePluck_empty() {
	fmt.Println(render(`{{ users | pluck:'name' }}`, map[string]any{
		"users": []any{},
	}))
	// Output: []
}

// ExamplePluck_chained collects a field and then reads the first value out of
// the result, chaining two access modifiers in one pipeline.
func ExamplePluck_chained() {
	fmt.Println(render(`{{ users | pluck:'name' | first }}`, map[string]any{
		"users": []any{
			map[string]any{"name": "Alice"},
			map[string]any{"name": "Bob"},
		},
	}))
	// Output: Alice
}

// ExampleFirstString returns the leading byte of a string as a one-byte string.
func ExampleFirstString() {
	fmt.Println(render(`{{ word | first }}`, map[string]any{
		"word": "hello",
	}))
	// Output: h
}

// ExampleFirstSlice returns the first element of a slice as-is, keeping its type.
func ExampleFirstSlice() {
	fmt.Println(render(`{{ items | first }}`, map[string]any{
		"items": []any{"espresso", "latte", "macchiato"},
	}))
	// Output: espresso
}

// ExampleFirstSlice_numbers returns the first element of a slice of numbers,
// leaving it a number rather than turning it into text.
func ExampleFirstSlice_numbers() {
	fmt.Println(render(`{{ scores | first }}`, map[string]any{
		"scores": []any{3, 1, 4, 1, 5},
	}))
	// Output: 3
}

// ExampleFirstBytes reads the leading byte of a []byte buffer as text, so a file
// read or an HTTP body behaves the same as a string rather than yielding a raw
// byte number from the slice clause.
func ExampleFirstBytes() {
	fmt.Println(render(`{{ buffer | first }}`, map[string]any{
		"buffer": []byte("hello"),
	}))
	// Output: h
}

// ExampleFirstString_empty shows that first has nothing to return for an empty
// string, so it reports an error.
func ExampleFirstString_empty() {
	fmt.Println(render(`{{ word | first }}`, map[string]any{
		"word": "",
	}))
	// Output: error: failed to render template: failed to render variable token 'word': function failed to apply: first expected a non-empty string
}

// ExampleLastString returns the trailing byte of a string as a one-byte string.
func ExampleLastString() {
	fmt.Println(render(`{{ word | last }}`, map[string]any{
		"word": "hello",
	}))
	// Output: o
}

// ExampleLastSlice returns the last element of a slice as-is.
func ExampleLastSlice() {
	fmt.Println(render(`{{ items | last }}`, map[string]any{
		"items": []any{"espresso", "latte", "macchiato"},
	}))
	// Output: macchiato
}

// ExampleLastSlice_numbers returns the last element of a slice of numbers,
// keeping it a number.
func ExampleLastSlice_numbers() {
	fmt.Println(render(`{{ scores | last }}`, map[string]any{
		"scores": []any{3, 1, 4, 1, 5},
	}))
	// Output: 5
}

// ExampleLastBytes reads the trailing byte of a []byte buffer as text, so a file
// read or an HTTP body behaves the same as a string rather than yielding a raw
// byte number from the slice clause.
func ExampleLastBytes() {
	fmt.Println(render(`{{ buffer | last }}`, map[string]any{
		"buffer": []byte("hello"),
	}))
	// Output: o
}

// ExampleLastString_empty shows that last has nothing to return for an empty
// string, so it reports an error.
func ExampleLastString_empty() {
	fmt.Println(render(`{{ word | last }}`, map[string]any{
		"word": "",
	}))
	// Output: error: failed to render template: failed to render variable token 'word': function failed to apply: last expected a non-empty string
}

// ExampleKey reads one value out of a map by key, walking nested maps with dot
// notation.
func ExampleKey() {
	fmt.Println(render(`{{ config | key:'database.host' }}`, map[string]any{
		"config": map[string]any{
			"database": map[string]any{"host": "db.local", "port": 5432},
		},
	}))
	// Output: db.local
}

// ExampleKey_default shows the forgiving lookup paired with default, so a key
// that is absent falls back to a supplied value instead of rendering nothing.
func ExampleKey_default() {
	fmt.Println(render(`{{ user | key:'phone' | default:'no phone on file' }}`, map[string]any{
		"user": map[string]any{"name": "Alice"},
	}))
	// Output: no phone on file
}

// ExampleKey_index reads one element out of a list by passing its position as a
// number instead of a name.
func ExampleKey_index() {
	fmt.Println(render(`{{ items | key:0 }}`, map[string]any{
		"items": []any{"espresso", "latte", "macchiato"},
	}))
	// Output: espresso
}

// ExampleKey_nestedNumber walks a dot path into nested maps and returns the
// value it lands on, keeping its number type.
func ExampleKey_nestedNumber() {
	fmt.Println(render(`{{ config | key:'database.port' }}`, map[string]any{
		"config": map[string]any{
			"database": map[string]any{"host": "db.local", "port": 5432},
		},
	}))
	// Output: 5432
}

// ExampleFindSlice returns the first map in a slice whose key field equals the
// wanted value.
func ExampleFindSlice() {
	fmt.Println(render(`{{ items | find:'status','active' }}`, map[string]any{
		"items": []any{
			map[string]any{"name": "Coffee", "status": "sold-out"},
			map[string]any{"name": "Tea", "status": "active"},
		},
	}))
	// Output:
	// {
	//   "name": "Tea",
	//   "status": "active"
	// }
}

// ExampleFindSlice_typed matches on an exact value and type, so a numeric id is
// compared as a number rather than as its text form.
func ExampleFindSlice_typed() {
	fmt.Println(render(`{{ users | find:'id',2 }}`, map[string]any{
		"users": []any{
			map[string]any{"id": 1, "name": "Ada"},
			map[string]any{"id": 2, "name": "Bo"},
		},
	}))
	// Output:
	// {
	//   "id": 2,
	//   "name": "Bo"
	// }
}

// ExampleFindSlice_default pairs find with default, so a search that matches
// nothing falls back to a supplied value instead of failing the render.
func ExampleFindSlice_default() {
	fmt.Println(render(`{{ items | find:'status','active' | default:'none in stock' }}`, map[string]any{
		"items": []any{
			map[string]any{"name": "Coffee", "status": "sold-out"},
		},
	}))
	// Output: none in stock
}

// ExampleFindMap returns the map itself when its key field equals the wanted
// value.
func ExampleFindMap() {
	fmt.Println(render(`{{ item | find:'status','active' }}`, map[string]any{
		"item": map[string]any{"name": "Tea", "status": "active"},
	}))
	// Output:
	// {
	//   "name": "Tea",
	//   "status": "active"
	// }
}

// ExampleFindMap_default falls back to a supplied value when a single map's
// field does not match the wanted value, instead of failing the render.
func ExampleFindMap_default() {
	fmt.Println(render(`{{ item | find:'status','active' | default:'unavailable' }}`, map[string]any{
		"item": map[string]any{"name": "Tea", "status": "brewing"},
	}))
	// Output: unavailable
}
