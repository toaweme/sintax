package format_test

import (
	"fmt"
	"time"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/format"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(format.Modifiers()).Render(tpl, vars)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return fmt.Sprintf("%v", out)
}

// moment is a fixed timestamp so the date examples never depend on the wall clock.
var moment = time.Date(2024, 3, 14, 9, 30, 5, 0, time.UTC)

// ExampleFormatTime renders a time.Time using a PHP-style date layout, where Y is
// the 4-digit year, m the zero-padded month, and d the day.
func ExampleFormatTime() {
	fmt.Println(render(`{{ at | format:'Y-m-d' }}`, map[string]any{
		"at": moment,
	}))
	// Output: 2024-03-14
}

// ExampleFormatTimeDefault renders a time.Time with the default "Y-m-d H:i:s"
// layout, the clause reached when no layout is given.
func ExampleFormatTimeDefault() {
	fmt.Println(render(`{{ at | format }}`, map[string]any{
		"at": moment,
	}))
	// Output: 2024-03-14 09:30:05
}

// ExampleLengthString returns the number of UTF-8 bytes in a string, so a
// multi-byte character counts as more than one.
func ExampleLengthString() {
	fmt.Println(render(`{{ name | length }}`, map[string]any{
		"name": "Ada",
	}))
	// Output: 3
}

// ExampleLengthReflect counts the elements of a slice, array, or map, the
// fallback clause reached for a non-string value.
func ExampleLengthReflect() {
	fmt.Println(render(`{{ items | length }}`, map[string]any{
		"items": []any{"a", "b", "c", "d"},
	}))
	// Output: 4
}

// ExampleDecimalDefault formats a number with two decimal places, the clause
// reached when no precision is given.
func ExampleDecimalDefault() {
	fmt.Println(render(`{{ price | decimal }}`, map[string]any{
		"price": 3.5,
	}))
	// Output: 3.50
}

// ExampleDecimalPlaces formats a number with the given number of decimal places,
// rounding to the nearest value at that precision.
func ExampleDecimalPlaces() {
	fmt.Println(render(`{{ ratio | decimal:4 }}`, map[string]any{
		"ratio": 1.23456,
	}))
	// Output: 1.2346
}

// ExampleLineNumbers prepends each line of a string with its zero-based line
// number.
func ExampleLineNumbers() {
	fmt.Println(render(`{{ body | line_numbers }}`, map[string]any{
		"body": "first\nsecond\nthird",
	}))
	// Output: 0. first
	// 1. second
	// 2. third
}

// ExampleCurrency converts a numeric value between currency units by scaling it
// with a ratio of unit sizes, here dollars into cents.
func ExampleCurrency() {
	fmt.Println(render(`{{ price | currency:1,100 }}`, map[string]any{
		"price": 9,
	}))
	// Output: 900
}
