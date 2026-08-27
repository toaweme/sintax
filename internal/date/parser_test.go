package date

import (
	"errors"
	"testing"
	"time"
)

func Test_Parser_Parse(t *testing.T) {
	tz := time.FixedZone("CET", 2*60*60)

	tests := []struct {
		name   string
		value  string
		format string
		loc    *time.Location
		want   time.Time
	}{
		{"long month name", "June 26, 2026", "F j, Y", time.UTC, time.Date(2026, time.June, 26, 0, 0, 0, 0, time.UTC)},
		{"short month name", "Mar 14, 2024", "M j, Y", time.UTC, time.Date(2024, time.March, 14, 0, 0, 0, 0, time.UTC)},
		{"iso date", "2024-03-14", "Y-m-d", time.UTC, time.Date(2024, time.March, 14, 0, 0, 0, 0, time.UTC)},
		{"default format", "2024-03-14 09:05:06", DefaultFormat, time.UTC, time.Date(2024, time.March, 14, 9, 5, 6, 0, time.UTC)},
		{"day slash date and time", "14/03/2024 09:05", "d/m/Y H:i", time.UTC, time.Date(2024, time.March, 14, 9, 5, 0, 0, time.UTC)},
		{"weekday is matched and discarded", "Thursday, March 14, 2024", "l, F j, Y", time.UTC, time.Date(2024, time.March, 14, 0, 0, 0, 0, time.UTC)},
		{"two digit year", "14-03-24", "d-m-y", time.UTC, time.Date(2024, time.March, 14, 0, 0, 0, 0, time.UTC)},
		{"twelve hour with meridiem", "2024-03-14 03:05 pm", "Y-m-d h:i a", time.UTC, time.Date(2024, time.March, 14, 15, 5, 0, 0, time.UTC)},
		{"offset in the value wins over the location", "2024-03-14 09:05:06 +02:00", "Y-m-d H:i:s P", time.UTC, time.Date(2024, time.March, 14, 9, 5, 6, 0, tz)},
		{"nil location falls back to UTC", "2024-03-14", "Y-m-d", nil, time.Date(2024, time.March, 14, 0, 0, 0, 0, time.UTC)},
		{"given location applies", "2024-03-14", "Y-m-d", tz, time.Date(2024, time.March, 14, 0, 0, 0, 0, tz)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewParser(DefaultMapping).Parse(tt.value, tt.format, tt.loc)
			if err != nil {
				t.Fatalf("Parse(%q, %q) returned an error: %v", tt.value, tt.format, err)
			}
			if !got.Equal(tt.want) {
				t.Fatalf("Parse(%q, %q) = %v, want %v", tt.value, tt.format, got, tt.want)
			}
		})
	}
}

func Test_Parser_Parse_Errors(t *testing.T) {
	tests := []struct {
		name   string
		value  string
		format string
	}{
		{"month name against a numeric layout", "June 26, 2026", "Y-m-d"},
		{"empty value", "", "Y-m-d"},
		{"empty format", "2024-03-14", ""},
		{"trailing text", "2024-03-14 leftovers", "Y-m-d"},
		{"impossible day", "2024-02-31", "Y-m-d"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewParser(DefaultMapping).Parse(tt.value, tt.format, time.UTC)
			if err == nil {
				t.Fatalf("Parse(%q, %q) = %v, want an error", tt.value, tt.format, got)
			}
			var parseErr *time.ParseError
			if !errors.As(err, &parseErr) {
				t.Fatalf("Parse(%q, %q) error does not wrap *time.ParseError: %v", tt.value, tt.format, err)
			}
			if !got.IsZero() {
				t.Fatalf("Parse(%q, %q) = %v, want the zero time alongside the error", tt.value, tt.format, got)
			}
		})
	}
}

func Test_Parser_Parse_CustomMapping(t *testing.T) {
	got, err := NewParser(map[string]string{"Y": "2006"}).Parse("2024-m-d", "Y-m-d", time.UTC)
	if err != nil {
		t.Fatalf("Parse() returned an error: %v", err)
	}
	if want := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC); !got.Equal(want) {
		t.Fatalf("Parse() = %v, want %v", got, want)
	}
}

// Test_Parser_RoundTrip asserts the two halves of the pair read the same layout
// language, so anything Formatter writes Parser reads back.
func Test_Parser_RoundTrip(t *testing.T) {
	moment := time.Date(2024, time.March, 14, 9, 5, 6, 0, time.UTC)
	formats := []string{"Y-m-d", "d/m/Y H:i", "F j, Y", "M j, Y H:i:s", DefaultFormat}

	for _, format := range formats {
		t.Run(format, func(t *testing.T) {
			rendered := NewFormatter(DefaultMapping).Format(moment, format)
			got, err := NewParser(DefaultMapping).Parse(rendered, format, time.UTC)
			if err != nil {
				t.Fatalf("Parse(%q, %q) returned an error: %v", rendered, format, err)
			}
			want := NewFormatter(DefaultMapping).Format(got, format)
			if want != rendered {
				t.Fatalf("round trip of %q gave %q, want %q", format, want, rendered)
			}
		})
	}
}
