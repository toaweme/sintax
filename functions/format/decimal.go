package format

import (
	"fmt"
	"strconv"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameDecimal is the template name for the Decimal modifier.
const ModifierNameDecimal functions.ModifierName = "decimal"

// Decimal formats a number as a string with a fixed number of decimal places,
// rounding to the nearest value at that precision. A numeric string is parsed
// first, and a nil value is treated as zero. The precision parameter is optional
// and defaults to 2.
//
// value: float, int, string
// param:0?: int - number of decimal places to keep (default 2)
// returns: string
//
// example: format a price with two decimals
// in:  amount = 19.5
// tpl: {{ amount | decimal:2 }}
// out: 19.50
//
// example: round a weight to one decimal
// in:  weight = 4.872
// tpl: {{ weight | decimal:1 }}
// out: 4.9
//
// example: use the default of two decimals
// in:  total = 7
// tpl: {{ total | decimal }}
// out: 7.00
//
// example: parse and format a numeric string to whole units
// in:  reading = "3.14159"
// tpl: {{ reading | decimal:0 }}
// out: 3
func Decimal(value any, params []any) (any, error) {
	places := 2
	if len(params) > 0 {
		switch p := params[0].(type) {
		case int:
			places = p
		case int64:
			places = int(p)
		case float64:
			places = int(p)
		default:
			return nil, fmt.Errorf("decimal: places must be an int, got %T", params[0])
		}
	}
	f, err := toFloat(value)
	if err != nil {
		return nil, fmt.Errorf("decimal: %w", err)
	}
	return strconv.FormatFloat(f, 'f', places, 64), nil
}

func toFloat(v any) (float64, error) {
	if v == nil {
		return 0, nil
	}
	switch x := v.(type) {
	case float64:
		return x, nil
	case float32:
		return float64(x), nil
	case int:
		return float64(x), nil
	case int8:
		return float64(x), nil
	case int16:
		return float64(x), nil
	case int32:
		return float64(x), nil
	case int64:
		return float64(x), nil
	case uint:
		return float64(x), nil
	case uint8:
		return float64(x), nil
	case uint16:
		return float64(x), nil
	case uint32:
		return float64(x), nil
	case uint64:
		return float64(x), nil
	case string:
		f, err := strconv.ParseFloat(x, 64)
		if err != nil {
			return 0, fmt.Errorf("cannot parse %q as number", x)
		}
		return f, nil
	}
	return 0, fmt.Errorf("cannot convert %T to number", v)
}
