package date

import (
	"fmt"
	"strings"
	"time"
)

type Date struct {
	time.Time
}

func NewDate(t time.Time) *Date {
	return &Date{t}
}

// Format maps PHP date format constants to Go's and returns the formatted string.
// Unsupported format characters will cause an error.
func (d *Date) Format(format string) (string, error) {
	var result strings.Builder

	for i := 0; i < len(format); i++ {
		layout, ok := Format[string(format[i])]
		if !ok {
			result.WriteString(string(format[i]))
			continue
		}
		if layout == "" {
			result.WriteString(string(format[i]))
			continue
		}

		// Handle custom cases or return an error if unsupported.
		switch format[i] {
		// Example of custom handling: "N" - ISO-8601 numeric representation of the day of the week
		case 'N':
			day := int(d.Time.Weekday())
			if day == 0 { // Go's Weekday starts from Sunday as 0
				day = 7
			}
			result.WriteString(fmt.Sprintf("%d", day))
		// Additional cases need to be implemented here.
		default:
			return "", fmt.Errorf("unsupported format character: '%s' in '%s'", string(format[i]), format)
		}
	}

	return result.String(), nil
}
