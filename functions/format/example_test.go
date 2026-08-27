package format_test

import (
	"fmt"
	"time"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions/format"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(sintax.WithModifiers(format.Modifiers())).Render(tpl, vars)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return fmt.Sprintf("%v", out)
}

// moment is a fixed timestamp so the date examples never depend on the wall clock.
var moment = time.Date(2024, 3, 14, 9, 30, 5, 0, time.UTC)

// ExampleDate renders a time.Time using a PHP-style date layout, where Y is
// the 4-digit year, m the zero-padded month, and d the day.
func ExampleDate() {
	fmt.Println(render(`{{ at | date:'Y-m-d' }}`, map[string]any{
		"at": moment,
	}))
	// Output: 2024-03-14
}

// ExampleDateDefault renders a time.Time with the default "Y-m-d H:i:s"
// layout, the clause reached when no layout is given.
func ExampleDateDefault() {
	fmt.Println(render(`{{ at | date }}`, map[string]any{
		"at": moment,
	}))
	// Output: 2024-03-14 09:30:05
}

// ExampleDate_dateName renders a time.Time with named day and month parts,
// where l is the full weekday, F the full month, and j the day without a
// leading zero.
func ExampleDate_dateName() {
	fmt.Println(render(`{{ at | date:'l, F j, Y' }}`, map[string]any{
		"at": moment,
	}))
	// Output: Thursday, March 14, 2024
}

// ExampleDate_timeOnly renders just the hour and minute, where H is the
// zero-padded 24-hour hour and i the zero-padded minute.
func ExampleDate_timeOnly() {
	fmt.Println(render(`{{ at | date:'H:i' }}`, map[string]any{
		"at": moment,
	}))
	// Output: 09:30
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

// ExampleLineNumbers prepends each line of a string with its one-based line
// number.
func ExampleLineNumbers() {
	fmt.Println(render(`{{ body | line_numbers }}`, map[string]any{
		"body": "first\nsecond\nthird",
	}))
	// Output: 1. first
	// 2. second
	// 3. third
}

// ExampleLineNumbers_single numbers a single line, which becomes line one.
func ExampleLineNumbers_single() {
	fmt.Println(render(`{{ body | line_numbers }}`, map[string]any{
		"body": "only one line",
	}))
	// Output: 1. only one line
}

// ExampleLineNumbers_steps numbers each line of a block from one, handy for
// turning a list of steps or a code snippet into a numbered listing.
func ExampleLineNumbers_steps() {
	fmt.Println(render(`{{ body | line_numbers }}`, map[string]any{
		"body": "clone the repo\nrun the build\nrun the tests",
	}))
	// Output: 1. clone the repo
	// 2. run the build
	// 3. run the tests
}

// ExampleLineNumbersFrom starts the numbering at a given line, so a snippet
// lifted from the middle of a file lines up with its real position.
func ExampleLineNumbersFrom() {
	fmt.Println(render(`{{ body | line_numbers:6 }}`, map[string]any{
		"body": "return nil\n}",
	}))
	// Output: 6. return nil
	// 7. }
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

// ExampleFromDate reads a date out of the string a store printed, using the
// same layout language date writes with, so a month name becomes a real date
// rather than a chain of replacements.
func ExampleFromDate() {
	fmt.Println(render(`{{ printed | from_date:'F j, Y' | date:'Y-m-d' }}`, map[string]any{
		"printed": "June 26, 2026",
	}))
	// Output: 2026-06-26
}

// ExampleFromDate_dateAndTime reads a date and a time together, where d is the
// zero-padded day, m the zero-padded month, H the 24-hour hour, and i the minute.
func ExampleFromDate_dateAndTime() {
	fmt.Println(render(`{{ printed | from_date:'d/m/Y H:i' | date }}`, map[string]any{
		"printed": "14/03/2024 09:30",
	}))
	// Output: 2024-03-14 09:30:00
}

// ExampleFromDateGuess reads a date with no layout given, working through the
// forms a printed date usually arrives in.
func ExampleFromDateGuess() {
	fmt.Println(render(`{{ printed | from_date | date:'Y-m-d' }}`, map[string]any{
		"printed": "26 June 2026",
	}))
	// Output: 2026-06-26
}

// ExampleFromDateGuess_dayBeyondTwelve reads a numeric date that only one
// reading fits, since 25 can only be the day.
func ExampleFromDateGuess_dayBeyondTwelve() {
	fmt.Println(render(`{{ printed | from_date | date:'l, F j, Y' }}`, map[string]any{
		"printed": "25/12/2026",
	}))
	// Output: Friday, December 25, 2026
}

// ExampleFromDateGuess_ambiguous refuses a numeric date that reads as a real
// date either way round, since a layout says which one it is.
func ExampleFromDateGuess_ambiguous() {
	fmt.Println(render(`{{ printed | from_date | date:'Y-m-d' }}`, map[string]any{
		"printed": "03/04/2026",
	}))
	fmt.Println(render(`{{ printed | from_date:'d/m/Y' | date:'Y-m-d' }}`, map[string]any{
		"printed": "03/04/2026",
	}))
	// Output:
	// error: failed to render template: failed to render variable token 'printed': modifier "from_date": function failed to apply: failed to read date "03/04/2026" without a layout (pass one such as 'd/m/Y'): day and month could be either way round
	// 2026-04-03
}
