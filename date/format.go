package date

import (
	"fmt"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
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
	log.Trace().Msgf("formatting time: %s", d.Time)
	for i := 0; i < len(format); i++ {
		log.Trace().Msgf("formatting character: %s", string(format[i]))
		layout, ok := Format[string(format[i])]
		if !ok {
			err := fmt.Errorf("skipping character '%s' invalid date format attribute", string(format[i]))
			log.Warn().Err(err).Msg("failed to apply format filter")
			result.WriteString(string(format[i]))
			continue
		}
		if layout == "" {
			err := fmt.Errorf("character '%s' does not have a direct mapping to a date format attribute", string(format[i]))

			log.Err(err).Msg("failed to apply format filter")
			result.WriteString(string(format[i]))
			continue
		}

		result.WriteString(layout)
	}

	return result.String(), nil
}
