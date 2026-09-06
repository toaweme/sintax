package boolean_test

import (
	"fmt"
	"time"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/boolean"
	"github.com/toaweme/sintax/functions/control"
)

// render registers the control modifiers alongside the boolean ones, since a
// boolean modifier refuses absent data and default is how a template answers it.
func render(tpl string, vars map[string]any) string {
	mods := boolean.Modifiers()
	for name, mod := range control.Modifiers() {
		mods[name] = mod
	}
	out, err := sintax.New(sintax.WithModifiers(mods)).Render(tpl, vars)
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

// ExampleNot_emptyString treats an empty string as falsey, so not reports it as
// true.
func ExampleNot_emptyString() {
	fmt.Println(render(`{{ name | not }}`, map[string]any{
		"name": "",
	}))
	// Output: true
}

// ExampleNot_zero treats zero as falsey, so not reports it as true.
func ExampleNot_zero() {
	fmt.Println(render(`{{ count | not }}`, map[string]any{
		"count": 0,
	}))
	// Output: true
}

// ExampleNot_falseString treats the literal string "false" as falsey, matching
// how a template `if` reads it, so not returns true.
func ExampleNot_falseString() {
	fmt.Println(render(`{{ flag | not }}`, map[string]any{
		"flag": "false",
	}))
	// Output: true
}

// ExampleGt reports whether the value is greater than the operand, comparing
// numerically across the int and float kinds.
func ExampleGt() {
	fmt.Println(render(`{{ score | gt:70 }}`, map[string]any{
		"score": 82,
	}))
	// Output: true
}

// ExampleGt_below returns false when the value does not exceed the operand.
func ExampleGt_below() {
	fmt.Println(render(`{{ score | gt:70 }}`, map[string]any{
		"score": 50,
	}))
	// Output: false
}

// ExampleGt_float compares fractional values numerically, so 4.5 exceeds 4.
func ExampleGt_float() {
	fmt.Println(render(`{{ rating | gt:4 }}`, map[string]any{
		"rating": 4.5,
	}))
	// Output: true
}

// ExampleGt_negative shows a negative value is not greater than zero.
func ExampleGt_negative() {
	fmt.Println(render(`{{ balance | gt:0 }}`, map[string]any{
		"balance": -5,
	}))
	// Output: false
}

// ExampleGte reports whether the value is greater than or equal to the operand,
// so an exact match still holds.
func ExampleGte() {
	fmt.Println(render(`{{ score | gte:70 }}`, map[string]any{
		"score": 70,
	}))
	// Output: true
}

// ExampleGte_float compares as numbers, so a fractional value clears a whole
// operand.
func ExampleGte_float() {
	fmt.Println(render(`{{ rating | gte:4 }}`, map[string]any{
		"rating": 4.5,
	}))
	// Output: true
}

// ExampleGte_below returns false when the value falls short of the operand.
func ExampleGte_below() {
	fmt.Println(render(`{{ score | gte:70 }}`, map[string]any{
		"score": 60,
	}))
	// Output: false
}

// ExampleEqNumber reports numeric equality across the int and float kinds, so 5
// equals 5.0.
func ExampleEqNumber() {
	fmt.Println(render(`{{ total | eq:5.0 }}`, map[string]any{
		"total": 5,
	}))
	// Output: true
}

// ExampleEqNumber_false returns false when two numbers differ.
func ExampleEqNumber_false() {
	fmt.Println(render(`{{ total | eq:5 }}`, map[string]any{
		"total": 3,
	}))
	// Output: false
}

// ExampleEqNumber_float compares fractional values numerically.
func ExampleEqNumber_float() {
	fmt.Println(render(`{{ price | eq:2.5 }}`, map[string]any{
		"price": 2.5,
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

// ExampleEqString_false returns false when the strings differ.
func ExampleEqString_false() {
	fmt.Println(render(`{{ status | eq:'active' }}`, map[string]any{
		"status": "pending",
	}))
	// Output: false
}

// ExampleEqString_unicode compares multi-byte strings verbatim, so an accented
// word matches its identical operand.
func ExampleEqString_unicode() {
	fmt.Println(render(`{{ city | eq:'café' }}`, map[string]any{
		"city": "café",
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

// ExampleEqAny_bool compares two booleans by value, so a true flag equals true.
func ExampleEqAny_bool() {
	fmt.Println(render(`{{ enabled | eq:true }}`, map[string]any{
		"enabled": true,
	}))
	// Output: true
}

// ExampleNeqString reports inequality, the negation of eq, so a status that differs
// from the one named is true.
func ExampleNeqString() {
	fmt.Println(render(`{{ status | neq:'archived' }}`, map[string]any{
		"status": "active",
	}))
	// Output: true
}

// ExampleNeqNumber compares numerically across the int and float
// kinds, so 5 and 5.0 are equal and neq reports false.
func ExampleNeqNumber() {
	fmt.Println(render(`{{ count | neq:5.0 }}`, map[string]any{
		"count": 5,
	}))
	// Output: false
}

// ExampleNeqAny_nil treats nil as equal only to nil, so an absent value differs
// from zero rather than reading as it.
func ExampleNeqAny_nil() {
	fmt.Println(render(`{{ missing | default:0 | neq:0 }}`, map[string]any{}))
	// Output: false
}

// ExampleAnd combines two conditions into one expression, so a template needs no
// nested if blocks to ask whether both hold.
func ExampleAnd() {
	fmt.Println(render(`{{ enabled | and:has_rows }}`, map[string]any{
		"enabled":  true,
		"has_rows": true,
	}))
	// Output: true
}

// ExampleAnd_absent reads a value that was never set as false, so the condition
// answers rather than failing the render.
func ExampleAnd_absent() {
	fmt.Println(render(`{{ enabled | default:false | and:has_rows }}`, map[string]any{
		"has_rows": true,
	}))
	// Output: false
}

// ExampleOr answers whether either side is set, and composes with not.
func ExampleOr() {
	fmt.Println(render(`{{ draft | not | or:published }}`, map[string]any{
		"draft":     true,
		"published": true,
	}))
	// Output: true
}

// ExampleOr_bool answers with a bool rather than with the truthy operand, so it
// reads as a condition and not as a fallback value.
func ExampleOr_bool() {
	fmt.Println(render(`{{ name | or:fallback }}`, map[string]any{
		"name":     "",
		"fallback": "anonymous",
	}))
	// Output: true
}

// ExampleLt reports whether a number is below a threshold.
func ExampleLt() {
	fmt.Println(render(`{{ stock | lt:5 }}`, map[string]any{
		"stock": 2,
	}))
	// Output: true
}

// ExampleLte reads a warning window off a day count, which is how a template
// asks whether something is expiring soon.
func ExampleLte() {
	fmt.Println(render(`{{ days_left | lte:14 }}`, map[string]any{
		"days_left": 12,
	}))
	// Output: true
}

// ExampleGtTime compares two dates, so a template can ask whether a deadline is
// still ahead. Both sides have to be real dates rather than printed ones.
func ExampleGtTime() {
	fmt.Println(render(`{{ expires_at | gt:today }}`, map[string]any{
		"expires_at": time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
		"today":      time.Date(2024, 3, 14, 0, 0, 0, 0, time.UTC),
	}))
	// Output: true
}

// ExampleLtTime asks whether a date has passed.
func ExampleLtTime() {
	fmt.Println(render(`{{ starts_at | lt:today }}`, map[string]any{
		"starts_at": time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC),
		"today":     time.Date(2024, 3, 14, 0, 0, 0, 0, time.UTC),
	}))
	// Output: true
}
