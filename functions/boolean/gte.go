package boolean

import (
	"errors"
	"fmt"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameGte is the template name for the Gte modifier.
const ModifierNameGte functions.ModifierName = "gte"

// Gte returns true if the numeric value is greater than or equal to the
// threshold. Both the value and the threshold are coerced to numbers across the
// int and float kinds, and nil counts as zero. A non-numeric value such as a
// string produces an error rather than a false result.
//
// value: int, float
// param:0: number
// returns: bool
//
// example: check minimum order quantity
// in:  qty = 1
// tpl: {{ qty | gte:1 }}
// out: true
//
// example: qualify for an A grade
// in:  score = 90
// tpl: {{ score | gte:90 }}
// out: true
//
// example: a value below the threshold fails
// in:  score = 89
// tpl: {{ score | gte:90 }}
// out: false
var Gte = func(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, errors.New("gt function requires at least one parameter")
	}

	val, err := functions.ValueNumber(value)
	if err != nil {
		return nil, fmt.Errorf("gte function expected a number, got %T", value)
	}

	than, err := functions.ValueNumber(params[0])
	if err != nil {
		return nil, fmt.Errorf("gte function expected a number, got %T", params[0])
	}

	return val >= than, nil
}
