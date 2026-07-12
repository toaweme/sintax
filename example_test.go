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
	engine := sintax.New(defaults.New())

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
		defaults.New(),
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
	engine := sintax.New(defaults.New())

	out, err := engine.Render(
		`{{ response | from:'json' | key:'orders' | filter:'status','paid' | pluck:'total' | sum | decimal:2 }}`,
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
	engine := sintax.New(defaults.New())

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

// ExampleNewWith registers a custom modifier alongside the defaults, so a
// template can call it by name like any built-in.
func ExampleNewWith() {
	shout := func(value any, _ []any) (any, error) {
		s, ok := value.(string)
		if !ok {
			return value, nil
		}
		return strings.ToUpper(s) + "!", nil
	}

	engine := sintax.New(defaults.NewWith(map[string]sintax.GlobalModifier{
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
	engine := sintax.New(defaults.New())

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
	engine := sintax.New(defaults.New())

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
