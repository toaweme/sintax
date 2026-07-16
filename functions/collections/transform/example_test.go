package transform_test

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/collections/transform"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(sintax.WithModifiers(transform.Modifiers())).Render(tpl, vars)
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

// ExampleFlatten flattens a slice by exactly one level, spreading any element
// that is itself a slice into the result.
func ExampleFlatten() {
	fmt.Println(render(`{{ groups | flatten }}`, map[string]any{
		"groups": []any{
			[]any{"a", "b"},
			"c",
			[]any{"d"},
		},
	}))
	// Output:
	// [
	//   "a",
	//   "b",
	//   "c",
	//   "d"
	// ]
}

// ExampleFlatten_alreadyFlat leaves a slice with no nested slices unchanged.
func ExampleFlatten_alreadyFlat() {
	fmt.Println(render(`{{ tags | flatten }}`, map[string]any{
		"tags": []any{"go", "sintax", "docs"},
	}))
	// Output:
	// [
	//   "go",
	//   "sintax",
	//   "docs"
	// ]
}

// ExampleFlatten_deeplyNested removes only one level, so a slice nested two deep
// stays a slice in the result.
func ExampleFlatten_deeplyNested() {
	fmt.Println(render(`{{ groups | flatten }}`, map[string]any{
		"groups": []any{
			[]any{[]any{"x"}},
			[]any{"y"},
		},
	}))
	// Output:
	// [
	//   [
	//     "x"
	//   ],
	//   "y"
	// ]
}

// ExampleFlatten_empty flattens an empty slice to an empty slice.
func ExampleFlatten_empty() {
	fmt.Println(render(`{{ groups | flatten }}`, map[string]any{
		"groups": []any{},
	}))
	// Output: []
}

// ExampleSortAsc sorts a slice ascending, the default direction when none is
// given.
func ExampleSortAsc() {
	fmt.Println(render(`{{ names | sort }}`, map[string]any{
		"names": []any{"Charlie", "Alice", "Bob"},
	}))
	// Output:
	// [
	//   "Alice",
	//   "Bob",
	//   "Charlie"
	// ]
}

// ExampleSortAsc_numbers sorts numbers numerically rather than as text.
func ExampleSortAsc_numbers() {
	fmt.Println(render(`{{ scores | sort }}`, map[string]any{
		"scores": []any{30, 4, 200, 1},
	}))
	// Output:
	// [
	//   1,
	//   4,
	//   30,
	//   200
	// ]
}

// ExampleSortDir sorts a copy of a slice in the named direction, here descending.
func ExampleSortDir() {
	fmt.Println(render(`{{ scores | sort:'desc' }}`, map[string]any{
		"scores": []any{72, 95, 88},
	}))
	// Output:
	// [
	//   95,
	//   88,
	//   72
	// ]
}

// ExampleSortDir_strings sorts strings in reverse alphabetical order.
func ExampleSortDir_strings() {
	fmt.Println(render(`{{ names | sort:'desc' }}`, map[string]any{
		"names": []any{"Alice", "Charlie", "Bob"},
	}))
	// Output:
	// [
	//   "Charlie",
	//   "Bob",
	//   "Alice"
	// ]
}

// ExampleMap converts a slice of string-keyed maps into a single map keyed by the
// named field's value.
func ExampleMap() {
	fmt.Println(render(`{{ users | map:'id' }}`, map[string]any{
		"users": []map[string]any{
			{"id": "u1", "name": "Alice"},
			{"id": "u2", "name": "Bob"},
		},
	}))
	// Output:
	// {
	//   "u1": {
	//     "id": "u1",
	//     "name": "Alice"
	//   },
	//   "u2": {
	//     "id": "u2",
	//     "name": "Bob"
	//   }
	// }
}

// ExampleMap_missingField skips elements that lack the named field.
func ExampleMap_missingField() {
	fmt.Println(render(`{{ users | map:'id' }}`, map[string]any{
		"users": []map[string]any{
			{"id": "u1", "name": "Alice"},
			{"name": "Bob"},
		},
	}))
	// Output:
	// {
	//   "u1": {
	//     "id": "u1",
	//     "name": "Alice"
	//   }
	// }
}

// ExampleMap_duplicateKey keeps the later element when two share a key value.
func ExampleMap_duplicateKey() {
	fmt.Println(render(`{{ users | map:'id' }}`, map[string]any{
		"users": []map[string]any{
			{"id": "u1", "name": "Alice"},
			{"id": "u1", "name": "Bob"},
		},
	}))
	// Output:
	// {
	//   "u1": {
	//     "id": "u1",
	//     "name": "Bob"
	//   }
	// }
}

// ExampleMap_byName keys the slice by any field, here the name instead of the id.
func ExampleMap_byName() {
	fmt.Println(render(`{{ users | map:'name' }}`, map[string]any{
		"users": []map[string]any{
			{"id": "u1", "name": "Alice"},
			{"id": "u2", "name": "Bob"},
		},
	}))
	// Output:
	// {
	//   "Alice": {
	//     "id": "u1",
	//     "name": "Alice"
	//   },
	//   "Bob": {
	//     "id": "u2",
	//     "name": "Bob"
	//   }
	// }
}

// ExampleMerge keys a slice of maps by the named field exactly as Map does.
func ExampleMerge() {
	fmt.Println(render(`{{ users | merge:'id' }}`, map[string]any{
		"users": []map[string]any{
			{"id": "u1", "name": "Alice"},
			{"id": "u2", "name": "Bob"},
		},
	}))
	// Output:
	// {
	//   "u1": {
	//     "id": "u1",
	//     "name": "Alice"
	//   },
	//   "u2": {
	//     "id": "u2",
	//     "name": "Bob"
	//   }
	// }
}

// ExampleMerge_bySku turns a list of products into a table keyed by sku.
func ExampleMerge_bySku() {
	fmt.Println(render(`{{ products | merge:'sku' }}`, map[string]any{
		"products": []map[string]any{
			{"sku": "A1", "name": "Bolt"},
			{"sku": "B2", "name": "Nut"},
		},
	}))
	// Output:
	// {
	//   "A1": {
	//     "name": "Bolt",
	//     "sku": "A1"
	//   },
	//   "B2": {
	//     "name": "Nut",
	//     "sku": "B2"
	//   }
	// }
}

// ExampleMerge_missing skips any element without the key field.
func ExampleMerge_missing() {
	fmt.Println(render(`{{ products | merge:'sku' }}`, map[string]any{
		"products": []map[string]any{
			{"sku": "A1", "name": "Bolt"},
			{"name": "Loose"},
		},
	}))
	// Output:
	// {
	//   "A1": {
	//     "name": "Bolt",
	//     "sku": "A1"
	//   }
	// }
}

// ExampleMerge_overwrite keeps the last element when a key repeats.
func ExampleMerge_overwrite() {
	fmt.Println(render(`{{ products | merge:'sku' }}`, map[string]any{
		"products": []map[string]any{
			{"sku": "A1", "name": "Old"},
			{"sku": "A1", "name": "New"},
		},
	}))
	// Output:
	// {
	//   "A1": {
	//     "name": "New",
	//     "sku": "A1"
	//   }
	// }
}

// ExampleSumElements adds up the elements of a slice, returning a number.
func ExampleSumElements() {
	fmt.Println(render(`{{ prices | sum }}`, map[string]any{
		"prices": []any{10, 8, 4},
	}))
	// Output: 22
}

// ExampleSumElements_strings sums numeric strings, parsing each before adding.
func ExampleSumElements_strings() {
	fmt.Println(render(`{{ amounts | sum }}`, map[string]any{
		"amounts": []any{"1.5", "2.5", "6"},
	}))
	// Output: 10
}

// ExampleSumElements_empty sums an empty slice to zero.
func ExampleSumElements_empty() {
	fmt.Println(render(`{{ amounts | sum }}`, map[string]any{
		"amounts": []any{},
	}))
	// Output: 0
}

// ExampleSumField totals the named field across a slice of maps.
func ExampleSumField() {
	fmt.Println(render(`{{ items | sum:'price' }}`, map[string]any{
		"items": []any{
			map[string]any{"name": "Mug", "price": 10},
			map[string]any{"name": "Cup", "price": 8},
		},
	}))
	// Output: 18
}

// ExampleSumField_decimals totals a field holding fractional amounts.
func ExampleSumField_decimals() {
	fmt.Println(render(`{{ items | sum:'price' }}`, map[string]any{
		"items": []any{
			map[string]any{"name": "Mug", "price": 10.5},
			map[string]any{"name": "Cup", "price": 4.25},
		},
	}))
	// Output: 14.75
}

// ExampleSumField_negative totals a field that mixes positive and negative
// numbers.
func ExampleSumField_negative() {
	fmt.Println(render(`{{ entries | sum:'amount' }}`, map[string]any{
		"entries": []any{
			map[string]any{"amount": 10},
			map[string]any{"amount": -3},
		},
	}))
	// Output: 7
}
