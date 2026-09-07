package datetime

import (
	"time"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameAddMonths is the template name for the AddMonths modifier.
const ModifierNameAddMonths functions.ModifierName = "add_months"

// AddMonths shifts a date by a whole number of calendar months, forward on a
// positive amount and back on a negative one. It earns a name of its own
// because a renewal, a billing period and a notice period are all written in
// months, and a month is not a fixed number of days, so a template that said
// add_days:30 would drift a day or two every time it ran.
//
// A day that the target month does not have lands on that month's last day, so
// the 31st of January plus one month is the 28th of February (the 29th in a
// leap year) rather than the 2nd or 3rd of March. Overflowing into the next
// month is the arithmetic answer, and it is wrong for every case that asks the
// question, since a subscription that charges on the 31st charges in February.
//
// The clock time and the zone come through untouched.
func AddMonths(t time.Time, months int) (time.Time, error) {
	return shiftMonths(t, months), nil
}

// shiftMonths moves a date by whole months, clamping the day to the last day of
// the target month rather than letting it spill into the following one.
func shiftMonths(t time.Time, months int) time.Time {
	year, month, day := t.Date()
	total := int(month) - 1 + months
	targetYear := year + floorDiv(total, 12)
	targetMonth := time.Month(floorMod(total, 12) + 1)
	if last := daysInMonth(targetYear, targetMonth); day > last {
		day = last
	}
	hour, minute, second := t.Clock()
	return time.Date(targetYear, targetMonth, day, hour, minute, second, t.Nanosecond(), t.Location())
}

// daysInMonth reports how many days the given month has, reading day zero of
// the following month, which time.Date normalizes to the last day of this one.
func daysInMonth(year int, month time.Month) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

// floorDiv divides rounding towards negative infinity, which is what a month
// count needs. Go's own division truncates towards zero, so January minus one
// month would come out as month zero of the same year.
func floorDiv(a, b int) int {
	q := a / b
	if a%b != 0 && (a < 0) != (b < 0) {
		q--
	}
	return q
}

// floorMod is the remainder that pairs with floorDiv, always in [0, b).
func floorMod(a, b int) int {
	return a - floorDiv(a, b)*b
}
