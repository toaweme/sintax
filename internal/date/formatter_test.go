package date

import (
	"testing"
	"time"
)

func Test_Formatter_Format(t *testing.T) {
	moment := time.Date(2024, time.March, 14, 9, 5, 6, 0, time.UTC)
	afternoon := time.Date(2024, time.March, 14, 15, 5, 6, 0, time.UTC)
	tz := time.FixedZone("CET", 2*60*60)
	inTZ := time.Date(2024, time.March, 14, 9, 5, 6, 0, tz)

	tests := []struct {
		name   string
		time   time.Time
		format string
		want   string
	}{
		{"default format", moment, DefaultFormat, "2024-03-14 09:05:06"},
		{"year full", moment, "Y", "2024"},
		{"year short", moment, "y", "24"},
		{"month zero padded", moment, "m", "03"},
		{"month no padding", moment, "n", "3"},
		{"month short name", moment, "M", "Mar"},
		{"month full name", moment, "F", "March"},
		{"day zero padded", moment, "d", "14"},
		{"day no padding", moment, "j", "14"},
		{"day short name", moment, "D", "Thu"},
		{"day full name", moment, "l", "Thursday"},
		{"hour 24 zero padded", moment, "H", "09"},
		{"hour 24 no padding token still pads", moment, "G", "09"},
		{"hour 24 afternoon", afternoon, "H", "15"},
		{"hour 12 zero padded", afternoon, "h", "03"},
		{"hour 12 no padding", afternoon, "g", "3"},
		{"minute", moment, "i", "05"},
		{"second", moment, "s", "06"},
		{"meridiem lowercase", afternoon, "a", "pm"},
		{"meridiem uppercase", afternoon, "A", "PM"},
		{"timezone abbreviation", inTZ, "T", "CET"},
		{"timezone offset no colon", inTZ, "O", "+0200"},
		{"timezone offset with colon", inTZ, "P", "+02:00"},
		{"composite date and time", moment, "Y-m-d H:i:s", "2024-03-14 09:05:06"},
		{"composite with slashes", moment, "d/m/Y H:i", "14/03/2024 09:05"},
		{"unmapped characters pass through", moment, "Q Y", "Q 2024"},
		{"empty format", moment, "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFormatter(DefaultMapping)
			got := f.Format(tt.time, tt.format)
			if got != tt.want {
				t.Fatalf("Format(%q) = %q, want %q", tt.format, got, tt.want)
			}
		})
	}
}

func Test_Formatter_Format_CustomMapping(t *testing.T) {
	f := NewFormatter(map[string]string{"Y": "2006"})
	moment := time.Date(2024, time.March, 14, 0, 0, 0, 0, time.UTC)

	got := f.Format(moment, "Y-m-d")
	want := "2024-m-d"
	if got != want {
		t.Fatalf("Format() = %q, want %q", got, want)
	}
}
