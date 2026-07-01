// Package money provides template modifiers for currency formatting.
package money

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameCurrency is the template name for the Currency modifier.
const ModifierNameCurrency functions.ModifierName = "currency"

// Currency converts a numeric value between currency units by applying a unit multiplier ratio.
// Useful for converting between minor (cents) and major (dollars) units.
//
// value: int, float, string
// param:0: int
// param:1: int
// returns: int
//
// example: convert dollars to cents
// in:  price = 9
// tpl: {{ price | currency:1,100 }}
// out: 900
//
// example: convert cents back to dollars
// in:  cents = 1299
// tpl: {{ cents | currency:100,1 }}
// out: 12
func Currency(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, fmt.Errorf("currency requires 2 parameters")
	}

	fromUnits, err := functions.ParamInt(params, 0)
	if err != nil {
		return nil, err
	}

	toUnits, err := functions.ParamInt(params, 1)
	if err != nil {
		return nil, err
	}
	var val float64

	switch v := value.(type) {
	case int:
		val = float64(v)
	case float64:
		val = v
	case string:
		val, err = cleanCurrency(v)
		if err != nil {
			return nil, fmt.Errorf("failed to clean currency string: %w", err)
		}
	}

	// convert value from given units to target units
	convertedValue := val * float64(toUnits) / float64(fromUnits)

	// convert float into int
	intValue := int(convertedValue)

	return intValue, nil
}

func cleanCurrency(value string) (float64, error) {
	// remove dollar, euro, etc. symbols
	symbols := []string{"$", "€", "£", "¥"}
	for _, symbol := range symbols {
		value = strings.ReplaceAll(value, symbol, "")
	}

	// convert into float
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to convert %q to float: %w", value, err)
	}

	return floatValue, nil
}
