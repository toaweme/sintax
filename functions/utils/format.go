package utils

import (
	"fmt"
	"time"

	"github.com/toaweme/date"
)

func Format(value any, params []any) (any, error) {
	switch timeValue := value.(type) {
	case string:
		return timeValue, nil
	case time.Time:
		d := date.NewFormatter(timeValue, date.DefaultMapping)
		format := date.DefaultFormat
		if len(params) > 0 {
			if _, ok := params[0].(string); !ok {
				return nil, fmt.Errorf("format function expected string, got %T", params[0])
			}
			format = params[0].(string)
		}

		goDateFormat, err := d.Render(format)
		if err != nil {
			return nil, fmt.Errorf("failed to apply format filter '%s': %w", params[0], err)
		}
		return timeValue.Format(goDateFormat), nil
	}

	return nil, fmt.Errorf("format function expected string or time.Time, got %T", value)
}
