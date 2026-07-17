package query_test

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/collections/query"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(sintax.WithModifiers(query.Modifiers())).Render(tpl, vars)
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

// ExampleFilter returns the items of a slice whose named field equals the search
// value.
func ExampleFilter() {
	fmt.Println(render(`{{ items | filter:'role','admin' }}`, map[string]any{
		"items": []any{
			map[string]any{"name": "Alice", "role": "admin"},
			map[string]any{"name": "Bob", "role": "viewer"},
			map[string]any{"name": "Carol", "role": "admin"},
		},
	}))
	// Output:
	// [
	//   {
	//     "name": "Alice",
	//     "role": "admin"
	//   },
	//   {
	//     "name": "Carol",
	//     "role": "admin"
	//   }
	// ]
}

// ExampleFilter_nested filters on a nested field reached with dot notation.
func ExampleFilter_nested() {
	fmt.Println(render(`{{ posts | filter:'meta.published',true }}`, map[string]any{
		"posts": []any{
			map[string]any{"title": "Draft", "meta": map[string]any{"published": false}},
			map[string]any{"title": "Live", "meta": map[string]any{"published": true}},
		},
	}))
	// Output:
	// [
	//   {
	//     "meta": {
	//       "published": true
	//     },
	//     "title": "Live"
	//   }
	// ]
}

// ExampleFilter_number matches numbers by value, so an integer field equals a
// whole-number search.
func ExampleFilter_number() {
	fmt.Println(render(`{{ items | filter:'score',10 }}`, map[string]any{
		"items": []any{
			map[string]any{"name": "A", "score": 10},
			map[string]any{"name": "B", "score": 20},
		},
	}))
	// Output:
	// [
	//   {
	//     "name": "A",
	//     "score": 10
	//   }
	// ]
}

// ExampleFilter_noMatch returns an empty slice when nothing matches rather than
// an error.
func ExampleFilter_noMatch() {
	fmt.Println(render(`{{ items | filter:'role','owner' }}`, map[string]any{
		"items": []any{
			map[string]any{"name": "Alice", "role": "admin"},
			map[string]any{"name": "Bob", "role": "viewer"},
		},
	}))
	// Output: null
}

// ExampleHas reports whether a plain slice contains a given element.
func ExampleHas() {
	fmt.Println(render(`{{ tags | has:'featured' }}`, map[string]any{
		"tags": []any{"featured", "sale", "new"},
	}))
	// Output: true
}

// ExampleHas_map shows that for a map with a single parameter Has tests only
// whether the key exists, so a key mapped to false still counts as present.
func ExampleHas_map() {
	fmt.Println(render(`{{ config | has:'debug_mode' }}`, map[string]any{
		"config": map[string]any{"debug_mode": false, "region": "eu-west-1"},
	}))
	// Output: true
}

// ExampleHas_sliceOfMaps reports whether any item in a slice of maps has the
// named field set to the given value.
func ExampleHas_sliceOfMaps() {
	fmt.Println(render(`{{ items | has:'status','active' }}`, map[string]any{
		"items": []any{
			map[string]any{"name": "Coffee", "status": "sold-out"},
			map[string]any{"name": "Tea", "status": "active"},
		},
	}))
	// Output: true
}

// ExampleHas_anyOf passes several candidate values after the field key, so a
// slice of maps matches when any item's field equals any one of them.
func ExampleHas_anyOf() {
	fmt.Println(render(`{{ items | has:'status','active','pending' }}`, map[string]any{
		"items": []any{
			map[string]any{"name": "Coffee", "status": "sold-out"},
			map[string]any{"name": "Tea", "status": "pending"},
		},
	}))
	// Output: true
}

// ExampleHas_mapValue tests the stored value when given both a key and a value,
// returning false when the key holds a different value.
func ExampleHas_mapValue() {
	fmt.Println(render(`{{ config | has:'region','us-east-1' }}`, map[string]any{
		"config": map[string]any{"region": "eu-west-1"},
	}))
	// Output: false
}

// ExampleIs reports whether the value equals any one of the given candidates.
func ExampleIs() {
	fmt.Println(render(`{{ status | is:'active','pending' }}`, map[string]any{
		"status": "pending",
	}))
	// Output: true
}

// ExampleIs_noMatch returns false when the value equals none of the candidates.
func ExampleIs_noMatch() {
	fmt.Println(render(`{{ status | is:'active','pending' }}`, map[string]any{
		"status": "archived",
	}))
	// Output: false
}

// ExampleIs_single tests the value against a single candidate.
func ExampleIs_single() {
	fmt.Println(render(`{{ role | is:'admin' }}`, map[string]any{
		"role": "admin",
	}))
	// Output: true
}
