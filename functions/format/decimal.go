package format

import (
	"fmt"
	"strconv"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameDecimal is the template name for the Decimal modifier.
const ModifierNameDecimal functions.ModifierName = "decimal"

// DecimalDefault formats a number as a string with two decimal places, the
// clause reached when no precision is given.
func DecimalDefault(value any) (string, error) {
	return decimalFormat(value, 2)
}

// DecimalPlaces formats a number with the given number of decimal places,
// rounding to the nearest value at that precision.
func DecimalPlaces(value any, places int) (string, error) {
	return decimalFormat(value, places)
}

// decimalFormat parses value as a number (a numeric string is accepted and nil
// counts as zero) and renders it at the given precision.
func decimalFormat(value any, places int) (string, error) {
	f, err := functions.ParseNumber(value)
	if err != nil {
		return "", fmt.Errorf("failed to read decimal value: %w", err)
	}
	return strconv.FormatFloat(f, 'f', places, 64), nil
}
