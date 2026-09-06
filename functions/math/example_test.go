package math_test

import (
	"fmt"
	"maps"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions"
	"github.com/toaweme/sintax/functions/format"
	"github.com/toaweme/sintax/functions/math"
)

// render runs a template through the arithmetic modifiers and the format ones,
// since a computed number is usually written out at a fixed precision.
func render(tpl string, vars map[string]any) string {
	all := make(map[string]functions.GlobalModifier)
	maps.Copy(all, math.Modifiers())
	maps.Copy(all, format.Modifiers())
	out, err := sintax.New(sintax.WithModifiers(all)).Render(tpl, vars)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return fmt.Sprintf("%v", out)
}

// ExampleAdd adds a field to the value it is piped.
func ExampleAdd() {
	fmt.Println(render(`{{ subtotal | add:shipping | decimal:2 }}`, map[string]any{
		"subtotal": 42.5,
		"shipping": 4.95,
	}))
	// Output: 47.45
}

// ExampleSubtract works out what is left of a total.
func ExampleSubtract() {
	fmt.Println(render(`{{ seats | subtract:booked }} left`, map[string]any{
		"seats":  120,
		"booked": 97,
	}))
	// Output: 23 left
}

// ExampleMultiply scales an amount by a rate.
func ExampleMultiply() {
	fmt.Println(render(`{{ net | multiply:1.21 | decimal:2 }}`, map[string]any{
		"net": "100",
	}))
	// Output: 121.00
}

// ExampleDivide averages a total over a count.
func ExampleDivide() {
	fmt.Println(render(`{{ total | divide:orders | decimal:2 }}`, map[string]any{
		"total":  1250,
		"orders": 8,
	}))
	// Output: 156.25
}
