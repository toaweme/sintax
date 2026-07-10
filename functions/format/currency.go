package format

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameCurrency is the template name for the Currency modifier.
const ModifierNameCurrency functions.ModifierName = "currency"

// Currency converts a numeric value between currency units by scaling it with a
// ratio of unit sizes: the result is value * toUnits / fromUnits, truncated to a
// whole integer. A string value may carry a leading currency symbol ($, EUR,
// GBP, JPY), which is stripped before parsing. A value that is neither a number
// nor a string counts as zero.
func Currency(value any, fromUnits, toUnits int) (int, error) {
	var val float64
	switch v := value.(type) {
	case int:
		val = float64(v)
	case float64:
		val = v
	case string:
		f, err := cleanCurrency(v)
		if err != nil {
			return 0, fmt.Errorf("failed to clean currency string: %w", err)
		}
		val = f
	}
	return int(val * float64(toUnits) / float64(fromUnits)), nil
}

func cleanCurrency(value string) (float64, error) {
	for _, symbol := range []string{"$", "€", "£", "¥"} {
		value = strings.ReplaceAll(value, symbol, "")
	}
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to convert %q to float: %w", value, err)
	}
	return floatValue, nil
}
