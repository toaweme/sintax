package boolean

import (
	"github.com/toaweme/sintax/functions"
)

// ModifierNameNot is the template name for the Not modifier.
const ModifierNameNot functions.ModifierName = "not"

// Not inverts the truthiness of the value.
//
// value: any
// returns: bool
//
// example: invert an account status
// in:  is_active = true
// tpl: {{ is_active | not }}
// out: false
//
// example: detect a clean run with no errors
// in:  errors = []
// tpl: {{ errors | not }}
// out: true
var Not = func(value any, _ []any) (any, error) {
	valueIsTrue := functions.ConditionIsTrue(value)
	return !valueIsTrue, nil
}
