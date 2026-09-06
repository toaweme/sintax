package sintax

import (
	"testing"
	"time"

	"github.com/toaweme/sintax/assert"
	"github.com/toaweme/sintax/functions"
)

// today stands in for the now special variable, which the engine itself does
// not define, so the cases read the way a real template does while staying off
// the wall clock.
var today = time.Date(2024, 3, 14, 9, 30, 5, 0, time.UTC)

// Test_E2E_DateArithmetic renders the cases that sent apps storing a reminder
// date on the row, where the warning date is worked out where it is needed.
func Test_E2E_DateArithmetic(t *testing.T) {
	tests := []struct {
		name     string
		template string
		vars     map[string]any
		expected any
	}{
		{
			name:     "warn a fortnight ahead",
			template: `{{ today | add_days:14 | date:'Y-m-d' }}`,
			vars:     map[string]any{"today": today},
			expected: "2024-03-28",
		},
		{
			name:     "warn a number of days read off the row",
			template: `{{ today | add_days:notice | date:'Y-m-d' }}`,
			vars:     map[string]any{"today": today, "notice": 30},
			expected: "2024-04-13",
		},
		{
			name:     "look back a week",
			template: `{{ today | add_days:-7 | date:'Y-m-d' }}`,
			vars:     map[string]any{"today": today},
			expected: "2024-03-07",
		},
		{
			name:     "cross a month boundary",
			template: `{{ at | add_days:5 | date:'Y-m-d' }}`,
			vars:     map[string]any{"at": time.Date(2024, 4, 28, 0, 0, 0, 0, time.UTC)},
			expected: "2024-05-03",
		},
		{
			name:     "a renewal date clamps onto a short month",
			template: `{{ at | add_months:1 | date:'Y-m-d' }}`,
			vars:     map[string]any{"at": time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC)},
			expected: "2024-02-29",
		},
		{
			name:     "a leap day a year on",
			template: `{{ at | add_years:1 | date:'Y-m-d' }}`,
			vars:     map[string]any{"at": time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)},
			expected: "2025-02-28",
		},
		{
			name:     "a warranty term read off the row",
			template: `{{ bought_on | add_years:years | date:'F j, Y' }}`,
			vars:     map[string]any{"bought_on": today, "years": 3},
			expected: "March 14, 2027",
		},
		{
			name:     "a printed date is read before it is shifted",
			template: `{{ printed | from_date:'d/m/Y' | add_days:1 | date:'Y-m-d' }}`,
			vars:     map[string]any{"printed": "31/12/2024"},
			expected: "2025-01-01",
		},
		{
			name:     "days left is a number a notice can read",
			template: `expires in {{ today | days_between:expires_at }} days`,
			vars:     map[string]any{"today": today, "expires_at": time.Date(2024, 3, 26, 0, 0, 0, 0, time.UTC)},
			expected: "expires in 12 days",
		},
		{
			name:     "expiring inside the warning window",
			template: `{{ today | days_between:expires_at | lte:14 }}`,
			vars:     map[string]any{"today": today, "expires_at": time.Date(2024, 3, 26, 0, 0, 0, 0, time.UTC)},
			expected: true,
		},
		{
			name:     "still outside the warning window",
			template: `{{ today | days_between:expires_at | lte:14 }}`,
			vars:     map[string]any{"today": today, "expires_at": time.Date(2024, 6, 26, 0, 0, 0, 0, time.UTC)},
			expected: false,
		},
		{
			name:     "a deadline still ahead",
			template: `{{ expires_at | gt:today }}`,
			vars:     map[string]any{"today": today, "expires_at": time.Date(2024, 3, 26, 0, 0, 0, 0, time.UTC)},
			expected: true,
		},
		{
			name:     "a deadline already passed",
			template: `{{ expires_at | lt:today }}`,
			vars:     map[string]any{"today": today, "expires_at": time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC)},
			expected: true,
		},
	}

	s := New(builtins())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := s.Render(tt.template, tt.vars)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_E2E_DateArithmetic_RefusesWrongTypes asserts a value that never became a
// date is refused where it goes wrong, the same stance the date modifier takes.
func Test_E2E_DateArithmetic_RefusesWrongTypes(t *testing.T) {
	tests := []struct {
		name     string
		template string
		vars     map[string]any
	}{
		{"a printed date is not a date", `{{ printed | add_days:1 }}`, map[string]any{"printed": "2024-03-14"}},
		{"a number is not a date", `{{ n | add_months:1 }}`, map[string]any{"n": 20240314}},
		{"an absent field is not a date", `{{ missing | add_days:1 }}`, map[string]any{}},
		{"a written amount is not a number", `{{ today | add_days:'14' }}`, map[string]any{"today": today}},
		{"days between needs two dates", `{{ today | days_between:'2024-03-26' }}`, map[string]any{"today": today}},
	}

	s := New(builtins())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.Render(tt.template, tt.vars)
			assert.Error(t, err)
		})
	}
}

// Test_E2E_Arithmetic renders the number cases, where a total is worked out in
// the template rather than in the query that fetched the rows.
func Test_E2E_Arithmetic(t *testing.T) {
	tests := []struct {
		name     string
		template string
		vars     map[string]any
		expected any
	}{
		{
			name:     "a line total",
			template: `{{ price | multiply:qty | decimal:2 }}`,
			vars:     map[string]any{"price": 12.5, "qty": 3},
			expected: "37.50",
		},
		{
			name:     "tax on a column summed out of rows",
			template: `{{ rows | sum:'amount' | multiply:0.21 | decimal:2 }}`,
			vars: map[string]any{"rows": []any{
				map[string]any{"amount": 100},
				map[string]any{"amount": "50"},
			}},
			expected: "31.50",
		},
		{
			name:     "what is left of an allowance",
			template: `{{ seats | subtract:booked }}`,
			vars:     map[string]any{"seats": 120, "booked": 97},
			expected: float64(23),
		},
		{
			name:     "an average over a count",
			template: `{{ total | divide:orders | decimal:2 }}`,
			vars:     map[string]any{"total": 1250, "orders": 8},
			expected: "156.25",
		},
		{
			name:     "a number stored as text still adds",
			template: `{{ subtotal | add:shipping | decimal:2 }}`,
			vars:     map[string]any{"subtotal": "42.50", "shipping": 4.95},
			expected: "47.45",
		},
	}

	s := New(builtins())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := s.Render(tt.template, tt.vars)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}

// Test_E2E_Arithmetic_RefusesWrongTypes asserts a value that is not a number is
// refused rather than counted as zero, and that an empty count does not divide.
func Test_E2E_Arithmetic_RefusesWrongTypes(t *testing.T) {
	tests := []struct {
		name     string
		template string
		vars     map[string]any
		target   error
	}{
		{"a word does not add", `{{ note | add:1 }}`, map[string]any{"note": "abc"}, functions.ErrInvalidValueType},
		{"a word is not an amount", `{{ n | subtract:note }}`, map[string]any{"n": 1, "note": "abc"}, functions.ErrInvalidParamType},
		{"an empty count does not divide", `{{ total | divide:orders }}`, map[string]any{"total": 10, "orders": 0}, functions.ErrInvalidParamValue},
		{"nothing to divide by does not divide", `{{ total | divide:orders }}`, map[string]any{"total": 10, "orders": nil}, functions.ErrInvalidParamValue},
	}

	s := New(builtins())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.Render(tt.template, tt.vars)
			assert.ErrorIs(t, err, tt.target)
		})
	}
}
