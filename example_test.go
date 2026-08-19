package sintax_test

import (
	"fmt"
	"strings"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/defaults"
)

// ExampleNew builds an engine with the batteries-included modifier set and
// renders a template against a set of variables.
func ExampleNew() {
	engine := sintax.New(defaults.All())

	out, err := engine.Render(`{{ name | upper }}`, map[string]any{
		"name": "ada",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: ADA
}

// ExampleRender renders a template in one call without holding onto an engine,
// passing the modifier set directly.
func ExampleRender() {
	out, err := sintax.Render(
		`{{ greeting | default:'hello' }}, {{ name }}`,
		map[string]any{"name": "world"},
		defaults.All(),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: hello, world
}

// ExampleNew_pipeline chains modifiers to turn a raw JSON response into a single
// formatted number, without leaving the template.
func ExampleNew_pipeline() {
	engine := sintax.New(defaults.All())

	out, err := engine.Render(
		`{{ response | from_json | key:'orders' | filter:'status','paid' | pluck:'total' | sum | decimal:2 }}`,
		map[string]any{
			"response": `{"orders":[
				{"total":10.5,"status":"paid"},
				{"total":4.25,"status":"pending"},
				{"total":15,"status":"paid"}
			]}`,
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 25.50
}

// ExampleNew_returnsValue shows that Render returns any, so a template resolving
// to a slice hands back a real slice with its element types intact, not a
// stringified one.
func ExampleNew_returnsValue() {
	engine := sintax.New(defaults.All())

	out, err := engine.Render(`{{ tags | split:',' }}`, map[string]any{
		"tags": "go,templates,data",
	})
	if err != nil {
		panic(err)
	}
	parts := out.([]string)
	fmt.Printf("%d tags, first is %q\n", len(parts), parts[0])
	// Output: 3 tags, first is "go"
}

// ExampleWithModifiers registers a custom modifier alongside the defaults, so a
// template can call it by name like any built-in. Options merge in order, so a
// modifier registered here under a built-in's name would replace it.
func ExampleWithModifiers() {
	shout := func(value any, _ []any) (any, error) {
		s, ok := value.(string)
		if !ok {
			return value, nil
		}
		return strings.ToUpper(s) + "!", nil
	}

	engine := sintax.New(defaults.All(), sintax.WithModifiers(map[string]sintax.GlobalModifier{
		"shout": shout,
	}))

	out, err := engine.Render(`{{ word | shout }}`, map[string]any{
		"word": "ship it",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: SHIP IT!
}

// ExampleNew_conditional renders an if/else block, choosing a branch from the
// truthiness of a variable.
func ExampleNew_conditional() {
	engine := sintax.New(defaults.All())

	out, err := engine.Render(
		`{{ if admin }}full access{{ else }}read only{{ endif }}`,
		map[string]any{"admin": false},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: read only
}

// ExampleNew_loop iterates a slice with a for block, binding the loop index to
// number each item.
func ExampleNew_loop() {
	engine := sintax.New(defaults.All())

	out, err := engine.Render(
		`{{ for i, name in names }}{{ i }}. {{ name }}
{{ endfor }}`,
		map[string]any{"names": []any{"first", "second", "third"}},
	)
	if err != nil {
		panic(err)
	}
	fmt.Print(out)
	// Output: 0. first
	// 1. second
	// 2. third
}

// ExampleNew_dottedPath reads a dotted name as a path into the value a shorter
// key holds, so a nested field needs no key modifier to reach it.
func ExampleNew_dottedPath() {
	engine := sintax.New(defaults.All())

	out, err := engine.Render(
		`{{ file_row.record.project_id }}`,
		map[string]any{
			"file_row.record": map[string]any{"project_id": "proj-1"},
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: proj-1
}

// ExampleNew_dottedPathFlatKeyWins shows the flat key answering ahead of the
// walk, so a producer that publishes names containing dots keeps deciding what
// they mean.
func ExampleNew_dottedPathFlatKeyWins() {
	engine := sintax.New(defaults.All())

	out, err := engine.Render(
		`{{ report.total }}`,
		map[string]any{
			"report.total": 42,
			"report":       map[string]any{"total": 7},
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 42
}

// ExampleNew_dottedPathMissing shows a walk that runs out arriving as a miss, so
// a default answers it exactly as it answers an absent variable.
func ExampleNew_dottedPathMissing() {
	engine := sintax.New(defaults.All())

	out, err := engine.Render(
		`{{ file_row.record.language | default:'en' }}`,
		map[string]any{
			"file_row.record": map[string]any{"project_id": "proj-1"},
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: en
}
