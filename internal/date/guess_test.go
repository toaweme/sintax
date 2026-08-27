package date

import (
	"errors"
	"testing"
	"time"
)

func Test_Guess(t *testing.T) {
	berlin := time.FixedZone("CEST", 2*60*60)

	tests := []struct {
		name  string
		value string
		loc   *time.Location
		want  time.Time
	}{
		{"rfc 3339 with nanoseconds", "2026-06-26T14:30:00.123456789Z", nil, time.Date(2026, time.June, 26, 14, 30, 0, 123456789, time.UTC)},
		{"rfc 3339", "2026-06-26T14:30:00Z", nil, time.Date(2026, time.June, 26, 14, 30, 0, 0, time.UTC)},
		{"rfc 3339 with an offset", "2026-06-26T14:30:00+02:00", nil, time.Date(2026, time.June, 26, 14, 30, 0, 0, berlin)},
		{"iso date and time, no zone", "2026-06-26T14:30:00", nil, time.Date(2026, time.June, 26, 14, 30, 0, 0, time.UTC)},
		{"iso date and minutes, no zone", "2026-06-26T14:30", nil, time.Date(2026, time.June, 26, 14, 30, 0, 0, time.UTC)},
		{"space separated with an offset", "2026-06-26 14:30:00+02:00", nil, time.Date(2026, time.June, 26, 14, 30, 0, 0, berlin)},
		{"the default output layout", "2024-03-14 09:30:05", nil, time.Date(2024, time.March, 14, 9, 30, 5, 0, time.UTC)},
		{"space separated minutes", "2024-03-14 09:30", nil, time.Date(2024, time.March, 14, 9, 30, 0, 0, time.UTC)},
		{"iso date only", "2024-03-14", nil, time.Date(2024, time.March, 14, 0, 0, 0, 0, time.UTC)},
		{"rfc 1123 with an offset", "Fri, 26 Jun 2026 14:30:00 +0200", nil, time.Date(2026, time.June, 26, 14, 30, 0, 0, berlin)},
		{"rfc 1123", "Fri, 26 Jun 2026 14:30:00 UTC", nil, time.Date(2026, time.June, 26, 14, 30, 0, 0, time.UTC)},
		{"weekday and long month", "Friday, June 26, 2026", nil, time.Date(2026, time.June, 26, 0, 0, 0, 0, time.UTC)},
		{"weekday, long month and a time", "Friday, June 26, 2026 14:30", nil, time.Date(2026, time.June, 26, 14, 30, 0, 0, time.UTC)},
		{"long month first", "June 26, 2026", nil, time.Date(2026, time.June, 26, 0, 0, 0, 0, time.UTC)},
		{"long month first with a time", "June 26, 2026 14:30", nil, time.Date(2026, time.June, 26, 14, 30, 0, 0, time.UTC)},
		{"long month first with seconds", "June 26, 2026 14:30:05", nil, time.Date(2026, time.June, 26, 14, 30, 5, 0, time.UTC)},
		{"long day first", "26 June 2026", nil, time.Date(2026, time.June, 26, 0, 0, 0, 0, time.UTC)},
		{"long day first with a time", "26 June 2026 14:30", nil, time.Date(2026, time.June, 26, 14, 30, 0, 0, time.UTC)},
		{"short month first", "Jun 26, 2026", nil, time.Date(2026, time.June, 26, 0, 0, 0, 0, time.UTC)},
		{"short month first with a time", "Jun 26, 2026 14:30", nil, time.Date(2026, time.June, 26, 14, 30, 0, 0, time.UTC)},
		{"short day first", "26 Jun 2026", nil, time.Date(2026, time.June, 26, 0, 0, 0, 0, time.UTC)},
		{"short day first with a time", "26 Jun 2026 14:30", nil, time.Date(2026, time.June, 26, 14, 30, 0, 0, time.UTC)},
		{"year first slashes", "2026/06/26", nil, time.Date(2026, time.June, 26, 0, 0, 0, 0, time.UTC)},
		{"year first slashes with a time", "2026/06/26 14:30", nil, time.Date(2026, time.June, 26, 14, 30, 0, 0, time.UTC)},
		{"year first slashes with seconds", "2026/06/26 14:30:05", nil, time.Date(2026, time.June, 26, 14, 30, 5, 0, time.UTC)},
		{"day beyond twelve reads day first", "25/12/2026", nil, time.Date(2026, time.December, 25, 0, 0, 0, 0, time.UTC)},
		{"month first when the day cannot be", "12/25/2026", nil, time.Date(2026, time.December, 25, 0, 0, 0, 0, time.UTC)},
		{"unpadded day beyond twelve", "25/6/2026", nil, time.Date(2026, time.June, 25, 0, 0, 0, 0, time.UTC)},
		{"numeric with a time", "25/12/2026 14:30", nil, time.Date(2026, time.December, 25, 14, 30, 0, 0, time.UTC)},
		{"numeric with seconds", "25/12/2026 14:30:05", nil, time.Date(2026, time.December, 25, 14, 30, 5, 0, time.UTC)},
		{"dotted", "25.12.2026", nil, time.Date(2026, time.December, 25, 0, 0, 0, 0, time.UTC)},
		{"dotted with a time", "25.12.2026 14:30", nil, time.Date(2026, time.December, 25, 14, 30, 0, 0, time.UTC)},
		{"hyphenated day first", "25-12-2026", nil, time.Date(2026, time.December, 25, 0, 0, 0, 0, time.UTC)},
		{"hyphenated with a time", "25-12-2026 14:30", nil, time.Date(2026, time.December, 25, 14, 30, 0, 0, time.UTC)},
		{"both readings agree", "05/05/2026", nil, time.Date(2026, time.May, 5, 0, 0, 0, 0, time.UTC)},
		{"a value with no zone takes the one given", "2024-03-14 09:30:05", berlin, time.Date(2024, time.March, 14, 9, 30, 5, 0, berlin)},
		{"a value with its own offset keeps it", "2026-06-26T14:30:00Z", berlin, time.Date(2026, time.June, 26, 14, 30, 0, 0, time.UTC)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Guess(tt.value, tt.loc)
			if err != nil {
				t.Fatalf("Guess(%q) returned an error: %v", tt.value, err)
			}
			if !got.Equal(tt.want) {
				t.Fatalf("Guess(%q) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}

// Test_Guess_Ambiguous asserts a numeric value that reads as a real date with
// the day and the month either way round is refused. Picking one would write a
// date that is wrong half the time and says nothing about it.
func Test_Guess_Ambiguous(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{"slashes", "03/04/2026"},
		{"slashes unpadded", "3/4/2026"},
		{"slashes with a time", "03/04/2026 14:30"},
		{"slashes with seconds", "03/04/2026 14:30:05"},
		{"dots", "03.04.2026"},
		{"dots with a time", "03.04.2026 14:30"},
		{"hyphens", "03-04-2026"},
		{"hyphens with a time", "03-04-2026 14:30"},
		{"first of the month either way", "01/02/2026"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Guess(tt.value, time.UTC)
			if !errors.Is(err, ErrAmbiguousDate) {
				t.Fatalf("Guess(%q) error = %v, want ErrAmbiguousDate", tt.value, err)
			}
			if !got.IsZero() {
				t.Fatalf("Guess(%q) = %v, want the zero time alongside the error", tt.value, got)
			}
		})
	}
}

// Test_Guess_Unknown asserts a value written in no layout Guess knows is a
// terminal error rather than a nearby date it half fits.
func Test_Guess_Unknown(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{"empty value", ""},
		{"prose", "next week"},
		{"trailing text", "2024-03-14 leftovers"},
		{"leading text", "on 2024-03-14"},
		{"impossible day", "2024-02-31"},
		{"month beyond twelve either way", "13/14/2026"},
		{"two digit year", "14/03/24"},
		{"a time on its own", "14:30"},
		{"a year on its own", "2026"},
		{"mixed separators", "25/12.2026"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Guess(tt.value, time.UTC)
			if !errors.Is(err, ErrUnknownLayout) {
				t.Fatalf("Guess(%q) error = %v, want ErrUnknownLayout", tt.value, err)
			}
			if !got.IsZero() {
				t.Fatalf("Guess(%q) = %v, want the zero time alongside the error", tt.value, got)
			}
		})
	}
}

// Test_Guess_RoundTripsWithTheDefaultLayout asserts a value the Formatter
// writes with the default layout reads back with nothing given, so the two
// no-argument defaults meet.
func Test_Guess_RoundTripsWithTheDefaultLayout(t *testing.T) {
	moment := time.Date(2024, 3, 14, 9, 30, 5, 0, time.UTC)
	rendered := NewFormatter(DefaultMapping).Format(moment, DefaultFormat)

	got, err := Guess(rendered, time.UTC)
	if err != nil {
		t.Fatalf("Guess(%q) returned an error: %v", rendered, err)
	}
	if !got.Equal(moment) {
		t.Fatalf("Guess(%q) = %v, want %v", rendered, got, moment)
	}
}
