package boolean_test

import (
	"fmt"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/boolean"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(boolean.Modifiers()).Render(tpl, vars)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return fmt.Sprintf("%v", out)
}

// ExampleNot inverts truthiness the same way a template `if` reads it, so a
// non-empty string is truthy and not makes it false.
func ExampleNot() {
	fmt.Println(render(`{{ name | not }}`, map[string]any{
		"name": "Ada",
	}))
	// Output: false
}

// ExampleGt reports whether the value is greater than the operand, comparing
// numerically across the int and float kinds.
func ExampleGt() {
	fmt.Println(render(`{{ score | gt:70 }}`, map[string]any{
		"score": 82,
	}))
	// Output: true
}

// ExampleGte reports whether the value is greater than or equal to the operand,
// so an exact match still holds.
func ExampleGte() {
	fmt.Println(render(`{{ score | gte:70 }}`, map[string]any{
		"score": 70,
	}))
	// Output: true
}

// ExampleEqNumber reports numeric equality across the int and float kinds, so 5
// equals 5.0.
func ExampleEqNumber() {
	fmt.Println(render(`{{ total | eq:5.0 }}`, map[string]any{
		"total": 5,
	}))
	// Output: true
}

// ExampleEqString reports verbatim string equality, the clause reached when
// neither operand is numeric.
func ExampleEqString() {
	fmt.Println(render(`{{ status | eq:'active' }}`, map[string]any{
		"status": "active",
	}))
	// Output: true
}

// ExampleEqAny falls back to Go equality for operands that are neither numeric
// nor both strings, so a number and its string form are never equal.
func ExampleEqAny() {
	fmt.Println(render(`{{ code | eq:'5' }}`, map[string]any{
		"code": 5,
	}))
	// Output: false
}
