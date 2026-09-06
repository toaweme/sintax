package datetime_test

import (
	"fmt"
	"maps"
	"time"

	"github.com/toaweme/sintax"
	"github.com/toaweme/sintax/functions"
	"github.com/toaweme/sintax/functions/datetime"
	"github.com/toaweme/sintax/functions/format"
)

// render runs a template through the date arithmetic modifiers and the format
// ones, since a shifted date is almost always written out with date.
func render(tpl string, vars map[string]any) string {
	all := make(map[string]functions.GlobalModifier)
	maps.Copy(all, datetime.Modifiers())
	maps.Copy(all, format.Modifiers())
	out, err := sintax.New(sintax.WithModifiers(all)).Render(tpl, vars)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return fmt.Sprintf("%v", out)
}

// moment is a fixed timestamp so the examples never depend on the wall clock.
var moment = time.Date(2024, 3, 14, 9, 30, 5, 0, time.UTC)

// ExampleAddDays shifts a date forward and writes the result out, which is how
// a reminder date is worked out at the point it is needed rather than stored on
// the row.
func ExampleAddDays() {
	fmt.Println(render(`{{ today | add_days:14 | date:'Y-m-d' }}`, map[string]any{
		"today": moment,
	}))
	// Output: 2024-03-28
}

// ExampleAddDays_back shifts a date back, since the amount is signed.
func ExampleAddDays_back() {
	fmt.Println(render(`{{ today | add_days:-7 | date:'Y-m-d' }}`, map[string]any{
		"today": moment,
	}))
	// Output: 2024-03-07
}

// ExampleAddMonths clamps a day the target month does not have, so a
// subscription charging on the 31st charges on the last day of February.
func ExampleAddMonths() {
	fmt.Println(render(`{{ charged_on | add_months:1 | date:'Y-m-d' }}`, map[string]any{
		"charged_on": time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC),
	}))
	// Output: 2024-02-29
}

// ExampleAddYears shifts by a warranty term read off the row.
func ExampleAddYears() {
	fmt.Println(render(`{{ bought_on | add_years:years | date:'Y-m-d' }}`, map[string]any{
		"bought_on": moment,
		"years":     2,
	}))
	// Output: 2026-03-14
}

// ExampleDaysBetween counts the days left, the number a notice reads out.
func ExampleDaysBetween() {
	fmt.Println(render(`expires in {{ today | days_between:expires_at }} days`, map[string]any{
		"today":      moment,
		"expires_at": time.Date(2024, 3, 26, 0, 0, 0, 0, time.UTC),
	}))
	// Output: expires in 12 days
}
