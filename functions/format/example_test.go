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

// ExampleFormatTime_dateName renders a time.Time with named day and month
// parts, where l is the full weekday, F the full month, and j the day without a
// leading zero.
func ExampleFormatTime_dateName() {
	fmt.Println(render(`{{ at | format:'l, F j, Y' }}`, map[string]any{
		"at": moment,
	}))
	// Output: Thursday, March 14, 2024
}

// ExampleFormatTime_timeOnly renders just the hour and minute, where H is the
// zero-padded 24-hour hour and i the zero-padded minute.
func ExampleFormatTime_timeOnly() {
	fmt.Println(render(`{{ at | format:'H:i' }}`, map[string]any{
		"at": moment,
	}))
	// Output: 09:30
}

// ExampleFormatTime_passthrough returns a string value unchanged, so format is
// safe to apply to a field that is already text.
func ExampleFormatTime_passthrough() {
	fmt.Println(render(`{{ label | format }}`, map[string]any{
		"label": "Q1 2024",
	}))
	// Output: Q1 2024
}

// ExampleLengthString returns the number of UTF-8 bytes in a string, so a
// multi-byte character counts as more than one.
func ExampleLengthString() {
	fmt.Println(render(`{{ name | length }}`, map[string]any{
		"name": "Ada",
	}))
	// Output: 3
}

// ExampleLengthString_unicode counts UTF-8 bytes rather than runes, so the
// accented "é" adds two to the total.
func ExampleLengthString_unicode() {
	fmt.Println(render(`{{ name | length }}`, map[string]any{
		"name": "café",
	}))
	// Output: 5
}

// ExampleLengthBytes returns the number of bytes in a byte-slice value.
func ExampleLengthBytes() {
	fmt.Println(render(`{{ raw | length }}`, map[string]any{
		"raw": []byte("hello"),
	}))
	// Output: 5
}

// ExampleLengthReflect_map counts the entries of a map.
func ExampleLengthReflect_map() {
	fmt.Println(render(`{{ headers | length }}`, map[string]any{
		"headers": map[string]any{"a": 1, "b": 2},
	}))
	// Output: 2
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

// ExampleDecimalDefault_string parses a numeric string before formatting it, so
// a value that arrives as text still renders with two decimal places.
func ExampleDecimalDefault_string() {
	fmt.Println(render(`{{ amount | decimal }}`, map[string]any{
		"amount": "7.5",
	}))
	// Output: 7.50
}

// ExampleDecimalPlaces formats a number with the given number of decimal places,
// rounding to the nearest value at that precision.
func ExampleDecimalPlaces() {
	fmt.Println(render(`{{ ratio | decimal:4 }}`, map[string]any{
		"ratio": 1.23456,
	}))
	// Output: 1.2346
}

// ExampleDecimalPlaces_zero rounds to a whole number when zero decimal places
// are requested.
func ExampleDecimalPlaces_zero() {
	fmt.Println(render(`{{ n | decimal:0 }}`, map[string]any{
		"n": 42.7,
	}))
	// Output: 43
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

// ExampleLineNumbers_single numbers a single line, which becomes line zero.
func ExampleLineNumbers_single() {
	fmt.Println(render(`{{ body | line_numbers }}`, map[string]any{
		"body": "only one line",
	}))
	// Output: 0. only one line
}

// ExampleLineNumbers_steps numbers each line of a block from zero, handy for
// turning a list of steps or a code snippet into a numbered listing.
func ExampleLineNumbers_steps() {
	fmt.Println(render(`{{ body | line_numbers }}`, map[string]any{
		"body": "clone the repo\nrun the build\nrun the tests",
	}))
	// Output: 0. clone the repo
	// 1. run the build
	// 2. run the tests
}

// ExampleCurrency converts a numeric value between currency units by scaling it
// with a ratio of unit sizes, here dollars into cents.
func ExampleCurrency() {
	fmt.Println(render(`{{ price | currency:1,100 }}`, map[string]any{
		"price": 9,
	}))
	// Output: 900
}

// ExampleCurrency_toDollars scales cents back down to whole dollars by giving
// the larger unit size first.
func ExampleCurrency_toDollars() {
	fmt.Println(render(`{{ price | currency:100,1 }}`, map[string]any{
		"price": 900,
	}))
	// Output: 9
}

// ExampleCurrency_symbol strips a leading currency symbol from a string value
// before scaling it.
func ExampleCurrency_symbol() {
	fmt.Println(render(`{{ price | currency:1,100 }}`, map[string]any{
		"price": "$4.50",
	}))
	// Output: 450
}

// ExampleCurrency_truncates drops the fractional part of the result rather than
// rounding it.
func ExampleCurrency_truncates() {
	fmt.Println(render(`{{ price | currency:1,1 }}`, map[string]any{
		"price": 1.99,
	}))
	// Output: 1
}
