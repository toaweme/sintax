package boolean

import (
	"github.com/toaweme/sintax/functions"
)

// ModifierNameNot is the template name for the Not modifier.
const ModifierNameNot functions.ModifierName = "not"

// Not inverts the truthiness of the value. Truthiness follows the same rules as
// a template `if`: booleans by their value, numbers are true when greater than
// zero, strings are true when non-empty (except the literal "false"), and
// collections are true when non-empty. nil and any unrecognised type are
// falsey, so Not returns true for them.
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
//
// example: a zero count is falsey, so Not is true
// in:  count = 0
// tpl: {{ count | not }}
// out: true
var Not = func(value any, _ []any) (any, error) {
	valueIsTrue := functions.ConditionIsTrue(value)
	return !valueIsTrue, nil
}
