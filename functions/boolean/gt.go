package boolean

import (
	"fmt"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameGt is the template name for the Gt modifier.
const ModifierNameGt functions.ModifierName = "gt"

// Gt returns true if the numeric value is greater than the threshold.
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
var Gt = func(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, fmt.Errorf("gt function requires at least one parameter")
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
