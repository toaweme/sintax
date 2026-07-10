package boolean

import (
	"errors"
	"fmt"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameGt is the template name for the Gt modifier.
const ModifierNameGt functions.ModifierName = "gt"

// Gt returns true if the numeric value is greater than the threshold. Both the
// value and the threshold are coerced to numbers across the int and float
// kinds, and nil counts as zero. A non-numeric value such as a string produces
// an error rather than a false result.
//
// value: int, float
// param:0: number
// returns: bool
//
// example: check if a basket has items
// in:  items_in_cart = 3
// tpl: {{ items_in_cart | gt:0 }}
// out: true
//
// example: check if a price exceeds free-shipping threshold
// in:  total = 49.99
// tpl: {{ total | gt:50 }}
// out: false
//
// example: an integer value beats a float threshold
// in:  score = 91
// tpl: {{ score | gt:90.5 }}
// out: true
var Gt = func(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, errors.New("gt function requires at least one parameter")
	}

	val, err := functions.ValueNumber(value)
	if err != nil {
		return nil, fmt.Errorf("gt function expected a number, got %T", value)
	}

	than, err := functions.ValueNumber(params[0])
	if err != nil {
		return nil, fmt.Errorf("gt function expected a number, got %T", params[0])
	}

	return val > than, nil
}
